export interface IConfig {
   TOKEN: string;
   GUILDID: string;
   CLIENTID: string;
   COIN_PRODUCT_ALERT_USERS: Array<string>;
   MONGOOSE_KEY: string;
   API_ADDRESS: string;
   DEVELOPMENT: "true" | "false";
   LOG_LEVEL: "info" | "debug" | "error" | "warn" | "trace";
}
