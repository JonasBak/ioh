package ioh_config

import (
	graphql "github.com/graph-gophers/graphql-go"
)

const Schema string = `
                schema {
                        query: Query
                        mutation: Mutation
                }
                type Query {
                        clients: [Client!]!
                        config(clientId: ID!): ClientConfig
                }
                type Client {
                        id: ID!
                        active: Boolean!
                        config: ClientConfig
                }
                type ClientConfig {
                        plant: String!
                        water: Int!
                }
                type Mutation {
                        setConfig(config: ClientConfigInput!): ClientConfig
                }
                input ClientConfigInput {
                        clientId: ID!
                        plant: String!
                        water: Int!
                }
        `

func (c Client) ID() graphql.ID {
	return graphql.ID(c.Id)
}

func (c ClientConfig) WATER() int32 {
	return int32(c.Water)
}

func (c Client) CONFIG() *ClientConfig {
	return c.config_ptr.GetConfig(c.Id)
}

type ClientConfigInput struct {
	ClientId graphql.ID
	Plant    string
	Water    int32
}

type Resolver struct {
	config IOHConfig
}

func NewResolver(config IOHConfig) Resolver {
	return Resolver{config}
}

func (r *Resolver) Clients() []Client {
	return r.config.GetClients()
}

func (r *Resolver) Config(args struct{ ClientId graphql.ID }) *ClientConfig {
	return r.config.GetConfig(string(args.ClientId))
}

func (r *Resolver) SetConfig(args struct {
	Config *ClientConfigInput
}) (*ClientConfig, error) {
	clientConfig := ClientConfig{Plant: args.Config.Plant, Water: int(args.Config.Water)}
	r.config.SetConfig(string(args.Config.ClientId), clientConfig)
	return r.config.GetConfig(string(args.Config.ClientId)), nil
}
