package helper

import (
	"com.lapangan.cuy/cmd/common"
	"database/sql"
	"fmt"
)

func ConnectToTurso(dbName string, dbAuthToken string) *sql.DB {
	dbUrl := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", dbName, dbAuthToken)
	db, err := sql.Open("libsql", dbUrl)
	common.FailOnError(err, "failed to connect to db")

	return db
}
