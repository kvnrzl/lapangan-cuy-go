package main

import (
	"com.lapangan.cuy/cmd/common"
	"com.lapangan.cuy/cmd/helper/db"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("bismillahirrohmanirrohim - start")

	common.LoadEnvDevFile()

	//dbName := os.Getenv("DB_NAME_TURSO")
	//dbAuthToken := os.Getenv("DB_AUTH_TOKEN_TURSO")

	//_ = db.ConnectToTurso(dbName, dbAuthToken)

	//log.Println("alhamdulillah - end")

	// ===== REDIS =====
	ctx := context.Background()

	client := db.ConnectToRedis()

	if err := client.Set(ctx, "name", "kevin", -1).Err(); err != nil {
		common.PanicOnError(err, "error set value on redis")
	}

	result, err := client.Get(ctx, "name").Result()
	if err != nil {
		common.PanicOnError(err, "error get value from redis")
	}

	fmt.Println("name : ", result)

	// ===== POSTGRES ======
	//connectionStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	//
	//conn, err := sql.Open("postgres", connectionStr)
	//common.PanicOnError(err, "can't connect to postgres")
	//
	//rows, err := conn.Query("SELECT version();")
	//if err != nil {
	//	panic(err)
	//}
	//
	//for rows.Next() {
	//	var version string
	//	rows.Scan(&version)
	//	fmt.Println(version)
	//}
	//
	//rows.Close()
	//
	//conn.Close()
}
