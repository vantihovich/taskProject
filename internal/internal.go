package internal

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func Check(username string, password string) (combExists bool) {

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
