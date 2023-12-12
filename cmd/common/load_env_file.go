package common

import (
	"github.com/joho/godotenv"
)

func LoadEnvDevFile() {
	err := godotenv.Load("../configs/.env.dev")
	PanicOnError(err, "failed to load env file")
}
