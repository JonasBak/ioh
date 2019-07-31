const path = require("path");

const URL_ENV = process.env.URL_ENV || process.env.NODE_ENV;

module.exports = {
  publicRuntimeConfig: {
    URLS:
      URL_ENV === "production"
        ? { BASE: "https://ioh-webapp.jonasbak.now.sh", HUB: "TODO" }
        : URL_ENV === "dev_compose"
        ? { BASE: "http://localhost:3000", HUB: "http://hub:5151" }
        : { BASE: "http://localhost:3000", HUB: "http://localhost:5151" }
  },
  webpack(config, options) {
    config.resolve.alias["components"] = path.join(__dirname, "components");
    config.resolve.alias["utils"] = path.join(__dirname, "utils");
    return config;
  }
};
