package main

import (
	"github.com/JonasBak/ioh/hub/ioh_config"
	"github.com/JonasBak/ioh/hub/mqtt"
	"github.com/JonasBak/ioh/hub/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	go mqtt.ConnectAndListen()

	router := mux.NewRouter()
	authMiddleware := server.AuthMiddleware()

	config := ioh_config.GetConfig()
	publisher := mqtt.GetPublisher()
	router.Handle("/", authMiddleware.Handler(server.GQLHandler(config, publisher)))
	router.Handle("/query", authMiddleware.Handler(server.QueryHandler(config, publisher)))

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":5151", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
