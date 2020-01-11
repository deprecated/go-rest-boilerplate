package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/deprecated/go-rest-boilerplate/config"
	"github.com/deprecated/go-rest-boilerplate/database"
	"github.com/deprecated/go-rest-boilerplate/cmd/server/routes"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// initialize DB connection
	db := database.InitDB(config.Config.Database.DBHost, config.Config.Database.DBPort, config.Config.Database.DBUser, config.Config.Database.DBPassword, config.Config.Database.DBName)
	defer db.Close()

	// inititalize all routes
	router := routes.InitRoutes(db, config.Config.JWT.AdminSecret, config.Config.JWT.UserSecret)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":3000", router))
}
