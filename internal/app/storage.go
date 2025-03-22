package app

import "context"

type list interface {
	IsIPInBlacklist(context.Context, string) (bool, error)
	IsIPInWhitelist(context.Context, string) (bool, error)
	AddToWhitelist(context.Context, string) error
	AddToBlacklist(context.Context, string) error
	DeleteFromBlacklist(context.Context, string) error
	DeleteFromWhitelist(context.Context, string) error
}

type limiter interface {
	IsIPInLimit(context.Context, string, int) (bool, error)
	IsLoginInLimit(context.Context, string, int) (bool, error)
	IsPasswordInLimit(context.Context, string, int) (bool, error)
	RemoveLimit(context.Context, string, string) error
}
