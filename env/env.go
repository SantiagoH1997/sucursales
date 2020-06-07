package env

import (
	"os"
)

func init() {
	os.Setenv("TEST_DB_HOST", "localhost")
	os.Setenv("TEST_DB_PORT", "27017")
	os.Setenv("TEST_DB_NAME", "challenge_test")
	os.Setenv("TEST_DB_COLLECTION", "sucursales")
}
