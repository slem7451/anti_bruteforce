package redis

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	db  *redis.Client
	ttl time.Duration
}

func NewClient(ctx context.Context) (*Client, error) {
	ttl, err := strconv.Atoi(os.Getenv("REDIS_TTL"))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Print("Redis: connected")

	return &Client{
		db:  client,
		ttl: time.Duration(ttl) * time.Second,
	}, nil
}

func (c *Client) isFieldInLimit(ctx context.Context, prefix, field string, lim int) (bool, error) {
	key := prefix + "_" + field

	res, err := c.db.Get(ctx, key).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}

	if errors.Is(err, redis.Nil) {
		err := c.db.Set(ctx, key, 1, c.ttl).Err()
		if err != nil {
			return false, err
		}

		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = c.db.Set(ctx, key, res+1, c.ttl).Err()
	if err != nil {
		return false, err
	}

	return res >= lim, nil
}

func (c *Client) IsIPInLimit(ctx context.Context, ip string, ipLim int) (bool, error) {
	return c.isFieldInLimit(ctx, "ip", ip, ipLim)
}

func (c *Client) IsLoginInLimit(ctx context.Context, login string, loginLim int) (bool, error) {
	return c.isFieldInLimit(ctx, "login", login, loginLim)
}

func (c *Client) IsPasswordInLimit(ctx context.Context, password string, passwordLim int) (bool, error) {
	return c.isFieldInLimit(ctx, "password", password, passwordLim)
}

func (c *Client) RemoveLimit(ctx context.Context, login, ip string) error {
	return c.db.Del(ctx, "ip_"+ip, "login_"+login).Err()
}
