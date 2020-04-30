// main.go

package handlers

import (
	productsApp "github.com/honnuanand/GoMuxPostgres/pkg/products/handlers"
	utils "github.com/honnuanand/GoMuxPostgres/pkg/utils"
)

func main() {
	// println(hello.Hello("HelloWorld"))

	a := productsApp.App{}
	a.Initialize(
		utils.GetEnv("postgres", "APP_DB_USERNAME"),
		utils.GetEnv("pgpass", "APP_DB_PASSWORD"),
		utils.GetEnv("postgres", "APP_DB_NAME"),
		utils.GetEnv("localhost", "APP_DB_HOST"))

	a.Run(":8010")
}
