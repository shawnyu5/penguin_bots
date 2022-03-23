import { AnyChannel, Channel, Client } from "discord.js";
// const exec = require("child_process").execSync;
import { execSync } from "child_process";
import { ICoinProduct } from "./types/coinProduct";
import config from "../config.json";

// TODO: Unit test this stuff

/**
 * @param client - discord client
 * @param channelName - name of channel to search for
 * @returns channel information with the name passed in. If not found. undefined
 */
function getChannelByName(
   client: Client,
   channelName: string
): AnyChannel | undefined {
   const channel = client.channels.cache.find((ch) => {
      // @ts-ignore
      return ch.name == channelName;
   });
   return channel;
}

/**
 * @returns return json string from python script. Other wise return null
 */
function checkCoinProduct(): ICoinProduct | null {
   // let result = execSync("python3 ../coin_products/coin_products.py");
   try {
      let result = execSync(
         "python3 ../coin_products/coin_products.py"
      ).toString();
      // console.log("checkCoinProduct result.toString(): %s", result); // __AUTO_GENERATED_PRINT_VAR__
      result = result.split("{")[1];
      result = "{" + result;
      // console.log(JSON.parse(result)); // __AUTO_GENERATED_PRINT_VAR__
      return JSON.parse(result);
   } catch (error) {
      // console.log(error);
      return null;
   }
}

/**
 * generates a message pinging all users in config.json about the coinProduct
 * @param coinProduct - the coin product
 * @returns A message string
 */
function buildMessage(coinProduct: ICoinProduct): string {
   let message: string = "";
   config.coin_product_alert_users.forEach((user) => {
      message += `<@${user}>
      `;
   });
   message += `title: ${coinProduct.title}
   url: ${coinProduct.url}`;
   return message;
}

export { checkCoinProduct, getChannelByName, buildMessage };
