package main

import (
	"log"
	"os"

	"github.com/judwhite/viceskeleton/pkg/database"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("skeleton", "A skeleton REST API that uses Postgres.")

	flagListen     = app.Flag("listen", "The HTTP listen address.").Default("127.0.0.1:8423").String()
	flagDBHost     = app.Flag("db-host", "The database host.").Default("127.0.0.1").String()
	flagDBPort     = app.Flag("db-port", "The database port.").Default("5434").Int()
	flagDBUser     = app.Flag("db-user", "The database user.").Default("postgres").String()
	flagDBPassword = app.Flag("db-password", "The database user's password.").Default("password").String()
	flagDBName     = app.Flag("db-name", "The database name.").Default("vicetestdb").String()
	flagDBSSL      = app.Flag("db-ssl", "The database SSL mode.").Default("disable").String()
)

// @title Vice Software Example API
// @version 1
// @BasePath /api/v1

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	dbSettings := database.Settings{
		Host:     *flagDBHost,
		Port:     *flagDBPort,
		User:     *flagDBUser,
		Password: *flagDBPassword,
		DBName:   *flagDBName,
		SSLMode:  *flagDBSSL,
	}

	db, err := database.New(dbSettings)
	if err != nil {
		log.Fatal(err)
	}

	ws := webserver{addr: *flagListen, db: db}
	ws.Start()
}
