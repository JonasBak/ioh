export const BASE_URL =
  process.env.NODE_ENV === "production"
    ? "https://ioh.jbakken.com"
    : "http://localhost:3001";
export const HUB_BASE_URL =
  process.env.NODE_ENV === "production"
    ? "http://hub:5151"
    : "http://localhost:5151";
