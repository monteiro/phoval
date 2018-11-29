package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"phoval"
	"phoval/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	userDB := flag.String("userdb", "root", "database user")
	passwordDB := flag.String("passworddb", "root", "database password")
	hostDB := flag.String("hostdb", "127.0.0.1", "database address")
	nameDB := flag.String("namedb", "verif2fa", "database name")

	db, err := createDbConn(*userDB, *passwordDB, *hostDB, *nameDB)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv := phoval.NewHttpServer(*addr, &mysql.Database{db})
	log.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func createDbConn(userDB string, passwordDB string, hostDB string, nameDB string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", userDB, passwordDB, hostDB, nameDB))
	if err != nil {
		return nil, err
	}

	return db, nil
}
