Init project:
    $ go mod init furniture_shop

Install dependencies:
    $ go get github.com/lib/pq
    $ go get -u github.com/goloop/env
    $ go get github.com/Masterminds/squirrel
    $ go get github.com/nfnt/resize

Install dlv:
    $ go install github.com/go-delve/delve/cmd/dlv@v1.24.1


Install missing dependencies:
    $ go mod tidy

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
    $ goose up      --- up
    $ goose down    --- down
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
    https://freshman.tech/file-upload-golang/
    https://www.calhoun.io/intro-to-templates-p3-functions/
    https://gist.github.com/obalunenko/8aea9d81901211891dd956379096617c --- template


Print struct to console:
    fmt.Printf("Structure: %+v\n", my_structure)


// ===========================================
    Operation templage

eq - Returns the boolean truth of arg1 == arg2
ne - Returns the boolean truth of arg1 != arg2
lt - Returns the boolean truth of arg1 < arg2
le - Returns the boolean truth of arg1 <= arg2
gt - Returns the boolean truth of arg1 > arg2
ge - Returns the boolean truth of arg1 >= arg2

{{if (ge .Usage .Limit)}}

{{else if (gt .Usage .Warning)}}

{{else if (eq .Usage 0)}}

{{else}}

{{end}}