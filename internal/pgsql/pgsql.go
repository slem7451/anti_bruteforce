package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib" //nolint:depguard
	"github.com/jmoiron/sqlx"       //nolint:depguard
	"github.com/pressly/goose/v3"
	"github.com/slem7451/anti_bruteforce/migrations"
)

const (
	blacklistType = "b"
	whitelistType = "w"
)

type Client struct {
	db   *sqlx.DB
	conn *sqlx.Conn
}

func NewClient(ctx context.Context) (*Client, error) {
	db, err := sqlx.Open("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(db.DB, "."); err != nil {
		return nil, err
	}

	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	log.Print("PostgreSQL: connected")

	return &Client{db: db, conn: conn}, nil
}

func (c *Client)isIPInList(ctx context.Context, ip string, listType string) (bool, error) {
	row := c.conn.QueryRowContext(ctx, `select ip from ips where ip = $1 and type = $2`, ip, listType)
	var searchedIP string

	err := row.Scan(&searchedIP)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return searchedIP != "", nil
}

func (c *Client)IsIPInBlacklist(ctx context.Context, ip string) (bool, error) {
	return c.isIPInList(ctx, ip, blacklistType)
}

func (c *Client)IsIPInWhitelist(ctx context.Context, ip string) (bool, error) {
	return c.isIPInList(ctx, ip, whitelistType)
}