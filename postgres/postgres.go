package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "user"
// 	password = "123qazWSX"
// 	dbname   = "projectdb"
// )

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type DB struct {
	Pool *pgxpool.Pool
	cfg  Config
}

func New(cfg Config) (db DB) {
	db.cfg = cfg
	return db
}

func (db *DB) Open(db DB, err error) {
	db.cfg = cfg
	pool, err := pgxpool.Connect(DB.cfg)
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")
	db.pool = pool
	return db, nil
}

func (db *DB) Close() {

}

// func (db *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {

// 	var id string

// 	sqlStatement := `SELECT id FROM users WHERE username=$1 and password=$2;`
// 	fmt.Println("Db in internal:", Db)

// 	row := Db.QueryRow(sqlStatement, username, password)

// 	switch err := row.Scan(&id); err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")

// 	case nil:
// 		fmt.Println(id)

// 	default:
// 		panic(err)
// 	}

// 	return row
// }
