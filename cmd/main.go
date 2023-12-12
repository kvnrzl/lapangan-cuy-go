package main

import (
	"com.lapangan.cuy/cmd/common"
	"com.lapangan.cuy/cmd/helper/db"
	_ "github.com/libsql/libsql-client-go/libsql"
	"log"
	"os"
)

func main() {
	log.Println("bismillahirrohmanirrohim - start")

	common.LoadEnvDevFile()

	dbName := os.Getenv("DB_NAME_TURSO")
	dbAuthToken := os.Getenv("DB_AUTH_TOKEN_TURSO")

	_ = db.ConnectToTurso(dbName, dbAuthToken)

	log.Println("alhamdulillah - end")
}
