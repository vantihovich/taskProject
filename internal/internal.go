package check

import (
	"database/sql"
	"fmt"

	dbc "github.com/vantihovich/taskProject/dbConn"
)

func Check(username string, password string) (combExists bool) {

	var e bool
	sqlStatement := `SELECT id FROM users WHERE username=$1 and password=$2;`

	var id string

	row := dbc.Db.QueryRow(sqlStatement, username, password)

	switch err := row.Scan(&id, &username, &password); err {
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
