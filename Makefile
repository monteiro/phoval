DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=verif2fa

bin: 
	[ -d bin ] || mkdir bin

build: bin
	go get -u -d github.com/golang-migrate/migrate/cli
	cd $(GOPATH)/src/github.com/golang-migrate/migrate/cli
	dep ensure	
	go build -tags 'mysql' -o bin/migrate github.com/golang-migrate/migrate/cli

#
# do all migrations
#
migrate: 
	bin/migrate -source file://migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

#
# usage: make create-migration MIGRATION_NAME=create_user
#
create-migration: build
	bin/migrate create -ext sql -dir migrations $(MIGRATION_NAME)

clean:
	rm -rf bin

tests:
    go test ./...
