import pino from "pino";
import config from "./enviroments/config.json";

const logger = pino({
   transport: { target: "pino-pretty" },
   options: { colorize: true },
   level: config.LOG_LEVEL || "info",
});

if (config.DEVELOPMENT == "true") {
   logger.level = "debug";
}

export default logger;
