const require = require("require");

class Client {
  constructor(service) {
    this.service = service;
    this.baseURL = null;
  }

  setBaseURL(URL) {
    this.baseURL = URL;
  }

  info(message) {
    const msg = {
      info: message,
      level: "info",
      service: this.service
    };

    this.send(msg);
  }

  warn(message) {
    const msg = {
      warn: message,
      level: "warn",
      service: this.service
    };

    this.send(msg);
  }

  error(message) {
    const msg = {
      error: message,
      level: "error",
      service: this.service
    };

    this.send(msg);
  }

  send(msg) {
    request({
      url: this.baseURL,
      method: "POST",
      json: true,
      body: JSON.stringify(msg)
    });
  }
}
