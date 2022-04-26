package main

import (
	"fmt"

	"github.com/marcelo-tm/go-rest-api/internal/comment"
	"github.com/marcelo-tm/go-rest-api/internal/db"
	transportHttp "github.com/marcelo-tm/go-rest-api/internal/transport/http"
)

// Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println(("Failed to connect to the database"))
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println(("Failed to migrate the database"))
		return err
	}
	fmt.Println("Successfully connected to the database")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
