package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/vantihovich/taskProject/configuration"
)

type DB struct {
	pool *pgxpool.Pool
	cfg  string
}

func New(cfg config.App) (db DB) {
	db.cfg = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	fmt.Println("Added configs to DB")
	return db
}

func (db *DB) Open() error {
	pool, err := pgxpool.Connect(context.Background(), db.cfg)
	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
		return err
	}

	fmt.Println("Successfully connected to DB!")
	db.pool = pool
	return nil
}

func (db *DB) QueryRow(ctx, sql string, args DB) pgx.Row {
	return db.pool.QueryRow(context.Background(), sql, args)
}

func (db *DB) Query(ctx, sql string, args DB) (pgx.Row, error) {
	return db.pool.Query(context.Background(), sql, args)
}

func (db *DB) Exec(ctx, sql string, args DB) ([]byte, error) {
	return db.pool.Exec(context.Background(), sql, args)
}

func (db *DB) Close() {
	db.pool.Close()
}
