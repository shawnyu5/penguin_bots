import pino from "pino";
import * as dotenv from "dotenv";
dotenv.config();

const logger = pino({
   transport: { target: "pino-pretty" },
   options: { colorize: true },
   level: process.env.LOG_LEVEL || "info",
});

export default logger;
