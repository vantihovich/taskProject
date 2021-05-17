package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "123qazWSX"
	dbname   = "projectdb"
)

var DB struct {
}

func (DB) Open() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func (DB) Close() {

}

func (DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	var e bool
	var id string

	sqlStatement := `SELECT id FROM users WHERE username=$1 and password=$2;`
	fmt.Println("Db in internal:", Db)

	row := Db.QueryRow(sqlStatement, username, password)

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		e = false
	case nil:
		fmt.Println(id)
		e = true
	default:
		panic(err)
	}

	return e
}
