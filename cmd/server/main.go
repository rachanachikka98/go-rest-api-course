package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/rachanachikka98/go-rest-api-course/internal/transport"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
