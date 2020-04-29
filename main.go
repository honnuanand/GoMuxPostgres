// main.go

package main

import (
	"os"
	// "github.com/honnuanand/GoMuxPostgres/pkg/hello"
)

func main() {
	// println(hello.Hello("HelloWorld"))

	a := App{}
	a.Initialize(
		GetEnv("postgres", "APP_DB_USERNAME"),
		GetEnv("pgpass", "APP_DB_PASSWORD"),
		GetEnv("postgres", "APP_DB_NAME"))

	a.Run(":8010")
}

//TODO : move ths to the GOStrings Module
func GetEnv(defaultString string, envName string) string {
	if value := os.Getenv(envName); value != "" {
		return value
	}
	return defaultString
}
