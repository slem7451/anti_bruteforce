package app

import (
	"context"
	"os"
	"strconv"

	"github.com/slem7451/anti_bruteforce/internal/pgsql"
	"github.com/slem7451/anti_bruteforce/internal/redis"
	"github.com/slem7451/anti_bruteforce/internal/request"
)

type App struct {
	rdb *redis.Client
	db *pgsql.Client
	loginLim int
	ipLim int
	passwordLim int
}

func NewApp(rdb *redis.Client, db *pgsql.Client) (*App, error) {
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

	if isIPLimited {
		return false, nil
	}

	isLoginLimited, err := a.rdb.IsLoginInLimit(ctx, req.Login, a.loginLim)
	if err != nil {
		return false, err
	}

	if isLoginLimited {
		return false, nil
	}

	isPasswordLimited, err := a.rdb.IsPasswordInLimit(ctx, req.Password, a.passwordLim)
	if err != nil {
		return false, err
	}

	if isPasswordLimited {
		return false, nil
	}

	return true, nil
}