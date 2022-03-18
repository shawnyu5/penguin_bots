// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { writeFile } from "fs";
import config from "../../../config.json";
const Iconfig = require("../../js/types/config.js");
const exec = require("child_process").exec;

interface Iconfig {
   token: string;
   guildID: string;
   clientID: string;
   coin_product_alert_users: Array<string>;
}

function updateUser(user: any) {
   console.log("execute user: %s", user); // __AUTO_GENERATED_PRINT_VAR__

   let updatedConfig: Iconfig = config;
   console.log("updateUser updatedConfig: %s", updatedConfig); // __AUTO_GENERATED_PRINT_VAR__

   let newUsers: Array<string> = updatedConfig.coin_product_alert_users;
   newUsers.push(user.id);
   updatedConfig = { ...updatedConfig, coin_product_alert_users: newUsers };
   return updatedConfig;
}

function checkCoinProduct() {
   let output: string;
   // const channel = <client>.channels.cache.get('<id>');
   // channel.send("<content>");
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
}
module.exports = {
   data: new SlashCommandBuilder()
      .setName("alert")
      .setDescription("opt into coin product alerts")
      .addStringOption((option: SlashCommandStringOption) =>
         option
            .setName("choice")
            .setDescription(
               "Choose weather to opt in or out of coin product notifications"
            )
            .setRequired(true)
            .addChoice("in", "in")
            .addChoice("out", "out")
      ),

   async execute(interaction: any) {
      let userChoice = String(interaction).split(":")[1];
      console.log("execute userChoice : %s", userChoice); // __AUTO_GENERATED_PRINT_VAR__

      // checkCoinProduct();

      let user: string = interaction.user; // get the user that sent the command
      let newConfig: Iconfig = updateUser(user);
      console.log("execute newConfig: %s", newConfig.coin_product_alert_users); // __AUTO_GENERATED_PRINT_VAR__

      await interaction.reply(
         `<@${newConfig.coin_product_alert_users}> recorded`
      );
   },
};
