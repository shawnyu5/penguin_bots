export interface IConfig {
   // discord bot token
   TOKEN: string;
   GUILDID: string;
   // discord bot client id
   CLIENTID: string;
   COIN_PRODUCT_ALERT_USERS: Array<string>;
   // mongodb connection string
   MONGOOSE_KEY: string;
   // api address for coin product
   API_ADDRESS: string;
   DEVELOPMENT: "true" | "false";
   LOG_LEVEL: "info" | "debug" | "error" | "warn" | "trace";
}
