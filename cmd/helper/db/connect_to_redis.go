package db

import (
	"com.lapangan.cuy/cmd/common"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func ConnectToRedis() *redis.Client {
	// address
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	address := fmt.Sprintf("%s:%s", host, port)

	// password
	password := os.Getenv("REDIS_PASSWORD")

	// db
	dbString := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbString)
	common.PanicOnError(err, "error convert string to int")

	// create new client
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return client

}
