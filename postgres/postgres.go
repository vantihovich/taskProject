package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	cnfg "github.com/vantihovich/taskProject/Conf"
)

type DB struct {
	pool *pgxpool.Pool
	cfg  cnfg.Config
}

func New(cfg cnfg.Config) (db DB) {
	db.cfg = cnfg.Configs()
	return db
}

func (db *DB) Open(err error) {
	pool, err := pgxpool.Connect(context.Background(), db.cfg.Database)
	if err != nil {
		fmt.Println("Unable to connection to database: %v\n", err)
	}

	fmt.Println("Successfully connected!")
	db.pool = pool

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
