package main

import (
	"github.com/JonasBak/ioh/hub/ioh_config"
	"github.com/JonasBak/ioh/hub/mqtt"
	"github.com/JonasBak/ioh/hub/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func applyMiddleware(handler http.Handler) http.Handler {
	// Auth
	if os.Getenv("DISABLE_AUTH") != "true" {
		authMiddleware := server.AuthMiddleware()
		handler = authMiddleware.Handler(handler)
	}

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	handler = handlers.CORS(originsOk, headersOk, methodsOk)(handler)

	// Logging
	handler = handlers.LoggingHandler(os.Stdout, handler)

	return handler
}

func main() {
	go mqtt.ConnectAndListen()

	router := mux.NewRouter()
	config := ioh_config.GetConfig()
	publisher := mqtt.GetPublisher()

	rootHandler := applyMiddleware(server.GQLHandler(config, publisher))
	queryHandler := applyMiddleware(server.QueryHandler(config, publisher))

	router.Handle("/", rootHandler)
	router.Handle("/query", queryHandler)

	http.ListenAndServe(":5151", (router))
}
