import pino from "pino";
import * as dotenv from "dotenv";
dotenv.config();

const logger = pino({
   transport: { target: "pino-pretty" },
   options: { colorize: true },
   level: process.env.LOG_LEVEL || "info",
});

if (process.env.DEVELOPMENT == "true") {
   logger.level = "debug";
}

export default logger;
