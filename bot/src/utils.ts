import { Client } from "discord.js";
import { writeFileSync } from "fs";
const exec = require("child_process").exec;

/* sends a message to a channel name
 * @param {Client} client the client
 * @param {string} channelName the channel name to send the message
 * @param {string} message the message to send
 * @return {boolean} if the message was successfully send
 */
function sendMessage(
   client: Client,
   channelName: string,
   message: string
): boolean {
   const channel = client.channels.cache.find((ch) => {
      // @ts-ignore
      return ch.name == "development";
   });

   if (channel) {
      // @ts-ignore
      channel.send(message);
      return true;
   }
   return false;
}

/*
 * check if the current product on penguin open box is a coin product
 * @return {boolean} whether the product is a coin product
 */
function checkCoinProduct(): string {
   let output: string | null = null;
   exec(
      "python3 /home/shawn/python/penguin_bots/coin_products/coin_products.py",
      (err: any, stdout: any, stderr: any) => {
         console.log("execute#(anon) err: %s", err.code); // __AUTO_GENERATED_PRINT_VAR__
         console.log("(anon) stdout: %s", stdout); // __AUTO_GENERATED_PRINT_VAR__
         // only record output if script exited successfull
         if (err.code == 0) {
            output = stdout;
         }
      }
   );

   if (output != null) {
      return output;
   } else {
      return "";
   }
}

export { checkCoinProduct, sendMessage };
