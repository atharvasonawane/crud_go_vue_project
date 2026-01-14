package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/sessions"
)

var DB *sql.DB
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func ConnectDB() {

	var err error

	dsn := "root:mysql@atharva04@tcp(localhost:3306)/student_db?parseTime=true"

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database successfully")

}

func InitSession(){
	Store.Options = &sessions.Options{
		Path: "/",
		MaxAge: 3600,
		HttpOnly: true,
	}
}
