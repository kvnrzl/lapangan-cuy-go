package helper

import (
	"com.lapangan.cuy/cmd/common"
	"github.com/joho/godotenv"
)

func LoadEnvDevFile() {
	err := godotenv.Load("../.env.dev")
	common.FailOnError(err, "failed to load env file")
}
