declare namespace NodeJS {
   interface ProcessEnv {
      TOKEN: string;
      GUILDID: string;
      CLIENTID: string;
      MONGOOSE_KEY: string;
      API_ADDRESS: string;
      DEVELOPMENT: "true" | "false";
      LOG_LEVEL: "debug" | "info" | "warn" | "error";
   }
}
