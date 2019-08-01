export const QUERIES = {
  getClients: {
    operationName: "GetConfig",
    query: `
      query GetConfig{
        clients {
          id
          active
          config {
            plant
            water
          }
        }
      }
    `,
    variables: []
  },
  getConfig: {
    operationName: "GetConfig",
    query: `
      query GetConfig($clientId: ID!){
        config(clientId: $clientId) {
          plant
          water
        }
      }
    `,
    variables: ["clientId"]
  },
  setConfig: {
    operationName: "SetConfig",
    query: `
      mutation SetConfig($clientId: ID!, $plant: String!, $water: Int!){
        setConfig(config: {clientId: $clientId, plant: $plant, water: $water}) {
          plant
          water
        }
      }
    `,
    variables: ["clientId", "plant", "water"]
  }
};
