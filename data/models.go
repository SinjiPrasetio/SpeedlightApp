package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	// Any models inserted here (and in the New function)
	// Are easily accessible throughout the entire app
}

func local() *time.Location {
	location := os.Getenv("TIME_LOCALTIME")
	if location == "" {
		location = "UTC"
	}
	local, err := time.LoadLocation(location)
	if err != nil {
		return nil
	}
	return local
}

func New(databasePool *sql.DB) Models {

	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)
	case "postgresql", "postgres":
		upper, _ = postgresql.New(databasePool)
	default:
		// do nothing
	}

	return Models{}
}

func getInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
