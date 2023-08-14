package main

import (
	"net/http"

	"github.com/SteaceP/GO-BLOG/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// "github.com/rs/cors"
	"github.com/urfave/negroni"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/blog/{id:[0-9]+}", handlers.BlogHandler)
	http.Handle("/", router)
	// this is the same as http.ListenAndServe(":80", router);
	}
	func New() negroni.HandlerFunc{
		return cors.Default().HandlerWithConfig(
			cors.Options{
				AllowedOrigins: []string{"*"}, // allow all origins
				AllowCredentials: true,         // allow cookies to be sent cross domain
				},
				)
				}

}

func main() {
	router := mux.NewRouter()
	SetupRoutes(router)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(cors.Default())

	n.UseHandler(router)
	n.Run(":8080") // same as http.ListenAndServe
}
