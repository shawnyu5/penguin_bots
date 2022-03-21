import { AnyChannel, Channel, Client } from "discord.js";
// const exec = require("child_process").execSync;
import { execSync } from "child_process";
import { ICoinProduct } from "./types/coinProduct";
import config from "../config.json";

/*
 * @param {Client} client the client
 * @param {string} channelName the channel name to send the message
 * @param {string} message the message to send
 * @return {boolean} if the message was successfully send
 */

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
 * @returns return json message from python script. Other wise return null
 */
function checkCoinProduct(): ICoinProduct {
   let result = execSync("python3 ../coin_products/coin_products.py");

   return JSON.parse(String(result));
}

/**
 * @param coinProduct - the coin product
 * @returns a message pining all users in config.json about the coinProduct
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
