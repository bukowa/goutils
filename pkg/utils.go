package pkg

import (
	"log"
	"os"
)

func GetEnvFatal(k string) (v string) {
	if v = os.Getenv(k); v == "" {
		log.Fatalf("Environment variable %s is empty", k)
	}
	return
}
