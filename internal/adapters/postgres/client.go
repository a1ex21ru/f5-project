package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Client struct {
	gormConn *gorm.DB
	sqlConn  *sql.DB
}

func (c *Client) GetGormConn() *gorm.DB {
	return c.gormConn
}

func NewClient(ctx context.Context) (*Client, error) {
	const fn = "NewClient"

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, database)
	connDB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	sqlConn, err := connDB.DB()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	if err = sqlConn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return &Client{
		gormConn: connDB,
		sqlConn:  sqlConn,
	}, nil
}

func (c *Client) Close() error {
	return c.sqlConn.Close()
}
