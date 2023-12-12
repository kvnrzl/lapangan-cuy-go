package db

import (
	"com.lapangan.cuy/cmd/common"
	"database/sql"
	"fmt"
)

func ConnectToTurso(dbName string, dbAuthToken string) *sql.DB {
	dbUrl := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", dbName, dbAuthToken)
	db, err := sql.Open("libsql", dbUrl)
	common.PanicOnError(err, "failed to connect to db")

	return db
}
