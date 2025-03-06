Init project:
    $ go mod init furniture_shop

Install dependencies:
    $ go get github.com/lib/pq
    $ go get -u github.com/goloop/env

Before run make sure that all migrations with goose done (read below).
Run:
    $ go run cmd/main.go

// =========================================
    Goose
    -------
Install goose for migrations:
    $ go install github.com/pressly/goose/v3/cmd/goose@latest
    check
    $ goose
    or
    $ ~/go/bin/goose
Then put credentials in .env file.
Create migration:
    -s flag to create sequestially
    $ goose create -s create_table_user sql

Run all migrations:
    $ goose up

// ======================================

	Postgres in docker with dbgate tool 
    (can connect ot any database, setting already created in docker-compose.yml file
    just uncomment what you need)
    -------------------------------------------------------
1. Build --- only for the first time
    $ docker-compoe bild
2. Run:
    $ docker-compose up
3. Stop
    $ docker-compose down
Open dbgate in browser;
    localhost:8087  --- credentials wrotten in docker-compose.yml

// =========================================
    Sources
    -------

http response in json:
    https://golangbyexample.com/json-response-body-http-go/