const PROXY_CONFIG = [
  {
    context: [
      "/weatherforecast",
    ],
    target: "https://localhost:7054",
    secure: false
  }
]

module.exports = PROXY_CONFIG;
