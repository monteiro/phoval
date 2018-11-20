package main

import (
	"2fa-api/pkg/storage"
	"flag"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	userDB := flag.String("userdb", "root", "database user")
	passwordDB := flag.String("passworddb", "root", "database password")
	hostDB := flag.String("hostdb", "127.0.0.1", "database address")
	nameDB := flag.String("namedb", "verif2fa", "database name")

	db, err := createDbConn(*userDB, *passwordDB, *hostDB, *nameDB)
	checkError(err)

	app := &App{
		Addr: *addr,
		Database: storage.Database{
			DB: db,
		},
	}

	app.RunServer()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
