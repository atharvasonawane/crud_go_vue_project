package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var DB *sql.DB
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func ConnectDB(dsn string) {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging DB: ", err)
	}

	fmt.Println("Connected to Database")
}

func InitSession() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}