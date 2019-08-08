package server

import (
	"github.com/JonasBak/ioh/hub/ioh_config"
	"github.com/JonasBak/ioh/hub/mqtt"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

func QueryHandler(config ioh_config.IOHConfig, publisher mqtt.Publisher) http.HandlerFunc {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	resolver := ioh_config.NewResolver(config, func(id string, c ioh_config.ClientConfig) {
		publisher.UpdatedConfig(id, c)
	})
	schema := graphql.MustParseSchema(ioh_config.Schema, &resolver, opts...)
	handler := relay.Handler{Schema: schema}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}

func GQLHandler(config ioh_config.IOHConfig, publisher mqtt.Publisher) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphiql)
	})
}
