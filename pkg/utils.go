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

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func HomeOrWd() (dir string) {
	if dir = HomeDir(); dir == "" {
		dir, _ = os.Getwd()
	}
	return
}
