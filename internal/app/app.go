package app

import (
	"context"
	"os"
	"strconv"

	"github.com/slem7451/anti_bruteforce/internal/entity/request"
)

type App struct {
	rdb limiter
	db list
	loginLim int
	ipLim int
	passwordLim int
}

func NewApp(rdb limiter, db list) (*App, error) {
	loginLim, err := strconv.Atoi(os.Getenv("MAX_LOGIN"))
	if err != nil {
		return nil, err
	}

	ipLim, err := strconv.Atoi(os.Getenv("MAX_IP"))
	if err != nil {
		return nil, err
	}

	passwordLim, err := strconv.Atoi(os.Getenv("MAX_PASSWORD"))
	if err != nil {
		return nil, err
	}

	return &App{
		rdb: rdb,
		db: db,
		loginLim: loginLim,
		ipLim: ipLim,
		passwordLim: passwordLim,
	}, nil
}

func (a *App)ValidateAuth(ctx context.Context, req request.Credits) (bool, error) {
	isBlack, err := a.db.IsIPInBlacklist(ctx, req.IP)
	if err != nil {
		return false, err
	}

	if isBlack {
		return false, nil
	}

	isWhite, err := a.db.IsIPInWhitelist(ctx, req.IP)
	if err != nil {
		return false, err
	}

	if isWhite {
		return true, nil
	}

	isIPLimited, err := a.rdb.IsIPInLimit(ctx, req.IP, a.ipLim)
	if err != nil {
		return false, err
	}

	isLoginLimited, err := a.rdb.IsLoginInLimit(ctx, req.Login, a.loginLim)
	if err != nil {
		return false, err
	}

	isPasswordLimited, err := a.rdb.IsPasswordInLimit(ctx, req.Password, a.passwordLim)
	if err != nil {
		return false, err
	}

	return !isIPLimited && !isLoginLimited && !isPasswordLimited, nil
}

func (a *App)AddToBlacklist(ctx context.Context, subnet string) error {
	return a.db.AddToBlacklist(ctx, subnet)
}

func (a *App)DeleteFromBlacklist(ctx context.Context, subnet string) error {
	return a.db.DeleteFromBlacklist(ctx, subnet)
}

func (a *App)AddToWhitelist(ctx context.Context, subnet string) error {
	return a.db.AddToWhitelist(ctx, subnet)
}

func (a *App)DeleteFromWhitelist(ctx context.Context, subnet string) error {
	return a.db.DeleteFromWhitelist(ctx, subnet)
}