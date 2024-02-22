package main

import (
	"bufio"
	"fmt"
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

	// Password
	val, ok := os.LookupEnv("DB_PASSWORD_FILE")
	if !ok {
		db_conf.db_password = os.Getenv("DB_PASSWORD")
	} else {
		file, err := os.Open(val)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		db_conf.db_password = scanner.Text()
	}

	return db_conf
}
