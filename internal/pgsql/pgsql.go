package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"
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

var ErrSubnetIsAlreadyExist = errors.New("subnet is already exist in this list")

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

func (c *Client)isSubnetInList(ctx context.Context, subnet string, listType string) (bool, error) {
	row := c.conn.QueryRowContext(ctx, `select subnet from ips where type = $1 and subnet = $2`, listType, subnet)
	var searchedSubnet string

	err := row.Scan(&searchedSubnet)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return searchedSubnet != "", nil
}

func (c *Client)isIPInList(ctx context.Context, ip string, listType string) (bool, error) {
	rows, err := c.conn.QueryContext(ctx, `select subnet from ips where type = $1`, listType)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	
	for rows.Next() {
		var searchedSubnet string

		err := rows.Scan(&searchedSubnet)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, nil
			}

			return false, err
		}

		_, subnet, err := net.ParseCIDR(searchedSubnet)
		if err != nil {
			return false, err
		}

		pIP := net.ParseIP(ip)

		if subnet.Contains(pIP) {
			return true, nil
		}
	}

	if err := rows.Close(); err != nil {
		return false, err
	}

	return false, nil
}

func (c *Client)IsIPInBlacklist(ctx context.Context, ip string) (bool, error) {
	return c.isIPInList(ctx, ip, blacklistType)
}

func (c *Client)IsIPInWhitelist(ctx context.Context, ip string) (bool, error) {
	return c.isIPInList(ctx, ip, whitelistType)
}

func (c *Client)deleteFromList(ctx context.Context, subnet string, listType string) error {
	_, err := c.conn.ExecContext(ctx, `delete from ips where subnet = $1 and type = $2`, subnet, listType)
	return err
}

func (c *Client)addToList(ctx context.Context, subnet string, listType string) error {
	isInList, err := c.isSubnetInList(ctx, subnet, listType)
	if err != nil {
		return err
	}

	if isInList {
		return ErrSubnetIsAlreadyExist
	}

	_, err = c.conn.ExecContext(ctx, `insert into ips (subnet, type) values ($1, $2)`, subnet, listType)
	return err
}

func (c *Client)AddToWhitelist(ctx context.Context, subnet string) error {
	return c.addToList(ctx, subnet, whitelistType)
}

func (c *Client)AddToBlacklist(ctx context.Context, subnet string) error {
	return c.addToList(ctx, subnet, blacklistType)
}

func (c *Client)DeleteFromBlacklist(ctx context.Context, subnet string) error {
	return c.deleteFromList(ctx, subnet, blacklistType)
}

func (c *Client)DeleteFromWhitelist(ctx context.Context, subnet string) error {
	return c.deleteFromList(ctx, subnet, whitelistType)
}