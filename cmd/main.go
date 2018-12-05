package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"monteiro/phoval/pkg/notification"
	"monteiro/phoval/pkg/phoval"
	"monteiro/phoval/pkg/phoval/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	envProduction = "prod"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	userDB := flag.String("userdb", "root", "database user")
	passwordDB := flag.String("passworddb", "root", "database password")
	hostDB := flag.String("hostdb", "127.0.0.1", "database address")
	nameDB := flag.String("namedb", "verif2fa", "database name")
	env := flag.String("env", "dev", "environment (dev, prod, stag)")
	brand := flag.String("brand", "phoval", "brand to be used in the message recipient")

	flag.Parse()

	db, err := createDbConn(*userDB, *passwordDB, *hostDB, *nameDB)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv := phoval.NewHttpServer(*addr, &mysql.VerificationStorage{DB: db}, *brand, getVerificationNotifier(*env))
	log.Printf("Starting server on %s", *addr)
	log.Fatal(srv.ListenAndServe())
}

func createDbConn(userDB string, passwordDB string, hostDB string, nameDB string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", userDB, passwordDB, hostDB, nameDB))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getVerificationNotifier(env string) phoval.VerificationNotifier {
	if env == envProduction {
		return notification.AWSSESNotifier{}
	}

	return notification.LoggerSmsNotifier{}
}
