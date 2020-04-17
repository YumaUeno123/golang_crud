package infrastructure

import (
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
	"os"
)

type SQLHandler struct {
	Conn *sqlx.DB
}

func NewSQLHandler() *SQLHandler {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	conn, err := sqlx.Open("mysql", "user:password@tcp(db:3306)/first_docker_project")
	if err != nil {
		log.Fatalf("error opening mysql. %s", err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}
