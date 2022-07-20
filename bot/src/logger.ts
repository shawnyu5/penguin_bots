import pino from "pino";
import { environment } from "./enviroments/enviroment";

const logger = pino({
   transport: { target: "pino-pretty" },
   options: { colorize: true },
   level: environment.LOG_LEVEL || "info",
});

if (environment.DEVELOPMENT) {
   logger.level = "debug";
}

export default logger;
