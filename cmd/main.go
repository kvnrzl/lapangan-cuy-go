package main

import (
	"com.lapangan.cuy/cmd/helper"
	_ "github.com/libsql/libsql-client-go/libsql"
	"log"
	"os"
)

func main() {
	log.Println("bismillahirrohmanirrohim - start")

	helper.LoadEnvDevFile()

	dbName := os.Getenv("DB_NAME")
	dbAuthToken := os.Getenv("DB_AUTH_TOKEN")

	_ = helper.ConnectToTurso(dbName, dbAuthToken)

	log.Println("alhamdulillah - end")
}
