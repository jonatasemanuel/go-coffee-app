package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonatasemanuel/coffee-server/db"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	// Models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is listening on port:", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// add router
	}

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		// TODO: add models later
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
