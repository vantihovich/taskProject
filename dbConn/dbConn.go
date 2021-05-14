package dbConn

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

func DbConn() (Db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	fmt.Println("Successfully connected!")

	return Db
}
