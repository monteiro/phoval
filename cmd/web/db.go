package main

import (
	"database/sql"
	"fmt"
)

func createDbConn(userDB string, passwordDB string, hostDB string, nameDB string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", userDB, passwordDB, hostDB, nameDB))
	if err != nil {
		return nil, err
	}

	return db, nil
}
