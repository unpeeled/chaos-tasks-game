package main

import (
	"os"
)

type DBConf struct {
	db_host string
	db_user string
	db_name string
}

type DBConfAll struct {
	conf        DBConf
	db_password string
}

// taken from https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDBconfig() DBConfAll {

	var db_conf DBConfAll

	db_conf.conf.db_host = getEnv("DB_HOST", "127.0.0.1")
	db_conf.conf.db_user = getEnv("DB_USER", "chaostasks")
	db_conf.conf.db_name = getEnv("DB_NAME", "chaostasks")
	db_conf.db_password = os.Getenv("DB_PASSWORD")

	return db_conf
}
