package server

import (
	"context"

	"github.com/slem7451/anti_bruteforce/internal/entity/request"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type App interface {
	ValidateAuth(context.Context, request.Credits) (bool, error)
	RemoveLimit(context.Context, request.Credits) error
	AddToBlacklist(context.Context, string) error
	DeleteFromBlacklist(context.Context, string) error
	AddToWhitelist(context.Context, string) error
	DeleteFromWhitelist(context.Context, string) error
}