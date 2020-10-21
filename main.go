package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"todos-svr/handler"

	"github.com/fsnotify/fsnotify"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var port int
var host string

func init() {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	env := "development"
	if envName := os.Getenv("ENV"); envName != "" {
		env = strings.ToLower(envName)
	}
	viper.SetConfigName(env)
	viper.MergeInConfig()

	viper.SetEnvPrefix(fmt.Sprint("todo_", env))
	viper.BindEnv("secret")
	viper.BindEnv("db_pass")
	port = viper.GetInt("port")
	host = viper.GetString("host")
}

func main() {
	cfgStr := fmt.Sprintf("%v:%v@tcp(%v)/%v", viper.Get("db_user"), viper.Get("db_pass"), viper.Get("db_host"), viper.Get("db"))
	cfg, err := mysql.ParseDSN(cfgStr)
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
