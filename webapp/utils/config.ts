export const BASE_URL =
  process.env.NODE_ENV === "production"
    ? "https://ioh.jbakken.com"
    : "http://localhost:3000";

export const HUB_BASE_URL =
  process.env.NODE_ENV === "production"
    ? "https://ioh-api.jbakken.com"
    : "http://localhost:5151";

export const AUTH_CONFIG = {
  domain: "jbakken.eu.auth0.com",
  clientId: "hr2ELdP18wXdnkUTsu0y9jjurB1QNNiI",
  redirectURI: BASE_URL,
  audience: "ioh-api.jbakken.com"
};
