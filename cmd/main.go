package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Captain-Leftovers/beekeepers_log_go_htmx/view/layout"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	_ = os.Getenv("JWT_SECRET")

	port := os.Getenv("PORT")

	portInt, err := strconv.Atoi(port)

	if err != nil {
		log.Fatalf("Port must be a number of a PORT but Got -> %v", port)
	}

	mainRouter := chi.NewRouter()

	mainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mainRouter.Handle("/", templ.Handler(layout.Base("Hello World!")))

	fmt.Println("Server is running on http://localhost:" + port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", portInt), mainRouter)

	if err != nil {
		fmt.Println("Error while starting the server")
		fmt.Println(err)
	}

}
