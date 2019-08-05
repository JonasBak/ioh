export const HUB_BASE_URL =
  process.env.NODE_ENV === "production"
    ? "https://ioh-api.jbakken.com"
    : "http://localhost:5151";
