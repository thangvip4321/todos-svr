package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"github.com/letung3105/todos-svr/handler"
)

var port int
var host string

func init() {
	flag.IntVar(&port, "p", 3000, " port flag define which port the server will use")
	flag.StringVar(&host, "h", "127.0.0.1", "the address of your server (default to localhost")
}

func main() {
	cfg, err := mysql.ParseDSN("thang:060901ttvt@tcp(127.0.0.1)/tasks")
	cfg.ParseTime = true
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	error := db.Ping()
	if error != nil {
		log.Fatal(error)
	}
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: handler.MainHandler(db),
	}
	server.ListenAndServe()
	defer db.Close()
}
