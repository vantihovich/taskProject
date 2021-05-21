package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type DB struct {
	pool *pgxpool.Pool
	cfg  string
}

func New(cfg string) (db DB) {
	db.cfg = cfg
	fmt.Println("Added configs to DB struct", db.cfg)
	return db
}

func (db *DB) Open() {
	fmt.Println("Trying to start method Open")
	pool, err := pgxpool.Connect(context.Background(), db.cfg)
	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Successfully connected!")
	db.pool = pool
	//return err
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
