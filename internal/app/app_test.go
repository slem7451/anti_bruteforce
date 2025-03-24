package app_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/slem7451/anti_bruteforce/internal/app"
	"github.com/slem7451/anti_bruteforce/internal/entity/request"
	appmock "github.com/slem7451/anti_bruteforce/internal/mocks/storage"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	ctx         context.Context
	ipLim       int
	loginLim    int
	passwordLim int
}

func (s *Suite) SetupTest() {
	os.Setenv("MAX_LOGIN", "10")
	os.Setenv("MAX_IP", "10")
	os.Setenv("MAX_PASSWORD", "10")
	s.ctx = context.Background()

	s.ipLim = 10
	s.loginLim = 10
	s.passwordLim = 10
}

func (s *Suite) TestLimits() {
	mLimiter := appmock.NewLimiter(s.T())
	mList := appmock.NewList(s.T())

	mList.EXPECT().IsIPInBlacklist(s.ctx, mock.AnythingOfType("string")).Return(false, nil)
	mList.EXPECT().IsIPInWhitelist(s.ctx, mock.AnythingOfType("string")).Return(false, nil)

	for i := 0; i < 10; i++ {
		mLimiter.EXPECT().IsIPInLimit(s.ctx, fmt.Sprintf("%dip", i), s.ipLim).Return(false, nil)
		mLimiter.EXPECT().IsLoginInLimit(s.ctx, fmt.Sprintf("%dlogin", i), s.loginLim).Return(false, nil)
		mLimiter.EXPECT().IsPasswordInLimit(s.ctx, fmt.Sprintf("%dpassword", i), s.passwordLim).Return(false, nil)

		mLimiter.EXPECT().IsIPInLimit(s.ctx, "ip", s.ipLim).Return(false, nil).Once()
		mLimiter.EXPECT().IsLoginInLimit(s.ctx, "login", s.loginLim).Return(false, nil).Once()
		mLimiter.EXPECT().IsPasswordInLimit(s.ctx, "password", s.passwordLim).Return(false, nil).Once()
	}

	mLimiter.EXPECT().IsIPInLimit(s.ctx, "ip", s.ipLim).Return(true, nil).Once()
	mLimiter.EXPECT().IsLoginInLimit(s.ctx, "login", s.loginLim).Return(true, nil).Once()
	mLimiter.EXPECT().IsPasswordInLimit(s.ctx, "password", s.passwordLim).Return(true, nil).Once()

	app, err := app.NewApp(mLimiter, mList)
	require.NoError(s.T(), err)

	for i := 0; i < 10; i++ {
		res, err := app.ValidateAuth(s.ctx, request.Credits{
			IP:       "ip",
			Login:    fmt.Sprintf("%dlogin", i),
			Password: fmt.Sprintf("%dpassword", i),
		})
		require.NoError(s.T(), err)
		require.Equal(s.T(), true, res)

		res, err = app.ValidateAuth(s.ctx, request.Credits{
			IP:       fmt.Sprintf("%dip", i),
			Login:    "login",
			Password: fmt.Sprintf("%dpassword", i),
		})
		require.NoError(s.T(), err)
		require.Equal(s.T(), true, res)

		res, err = app.ValidateAuth(s.ctx, request.Credits{
			IP:       fmt.Sprintf("%dip", i),
			Login:    fmt.Sprintf("%dlogin", i),
			Password: "password",
		})
		require.NoError(s.T(), err)
		require.Equal(s.T(), true, res)
	}

	res, err := app.ValidateAuth(s.ctx, request.Credits{IP: "ip", Login: "1login", Password: "1password"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), false, res)

	res, err = app.ValidateAuth(s.ctx, request.Credits{IP: "1ip", Login: "login", Password: "1password"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), false, res)

	res, err = app.ValidateAuth(s.ctx, request.Credits{IP: "1ip", Login: "1login", Password: "password"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), false, res)
}

func (s *Suite) TestLists() {
	mLimiter := appmock.NewLimiter(s.T())
	mList := appmock.NewList(s.T())

	mLimiter.EXPECT().IsIPInLimit(s.ctx, mock.AnythingOfType("string"), s.ipLim).Return(true, nil).Once()
	mLimiter.EXPECT().IsLoginInLimit(s.ctx, mock.AnythingOfType("string"), s.loginLim).Return(true, nil).Once()
	mLimiter.EXPECT().IsPasswordInLimit(s.ctx, mock.AnythingOfType("string"), s.passwordLim).Return(true, nil).Once()

	mLimiter.EXPECT().IsIPInLimit(s.ctx, mock.AnythingOfType("string"), s.ipLim).Return(false, nil)
	mLimiter.EXPECT().IsLoginInLimit(s.ctx, mock.AnythingOfType("string"), s.loginLim).Return(false, nil)
	mLimiter.EXPECT().IsPasswordInLimit(s.ctx, mock.AnythingOfType("string"), s.passwordLim).Return(false, nil)

	mList.EXPECT().IsIPInWhitelist(s.ctx, "ip_w").Return(true, nil).Once()
	mList.EXPECT().IsIPInBlacklist(s.ctx, "ip_w").Return(false, nil).Once()

	mList.EXPECT().IsIPInBlacklist(s.ctx, "ip_b").Return(true, nil).Once()

	mList.EXPECT().IsIPInBlacklist(s.ctx, "ip").Return(false, nil)
	mList.EXPECT().IsIPInWhitelist(s.ctx, "ip").Return(false, nil)

	app, err := app.NewApp(mLimiter, mList)
	require.NoError(s.T(), err)

	res, err := app.ValidateAuth(s.ctx, request.Credits{IP: "ip_w", Login: "l", Password: "p"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), true, res)

	res, err = app.ValidateAuth(s.ctx, request.Credits{IP: "ip_b", Login: "l", Password: "p"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), false, res)

	res, err = app.ValidateAuth(s.ctx, request.Credits{IP: "ip", Login: "l", Password: "p"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), false, res)

	res, err = app.ValidateAuth(s.ctx, request.Credits{IP: "ip", Login: "l", Password: "p"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), true, res)
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
