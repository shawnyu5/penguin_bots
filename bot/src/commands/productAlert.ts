// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { Interaction, User } from "discord.js";
import { writeFile } from "fs";
import config from "../../config.json";
const exec = require("child_process").exec;

interface Iconfig {
   token: string;
   guildID: string;
   clientID: string;
   coin_product_alert_users: Array<string>;
}

function updateUsers(user: any) {
   let updatedConfig: Iconfig = config;

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
            .setName("notification")
            .setDescription(
               "Choose weather to opt in or out of coin product notifications"
            )
            .setRequired(true)
            .addChoice("on", "on")
            .addChoice("off", "off")
      ),

   async execute(interaction: Interaction) {
      let userChoice = String(interaction).split(":")[1];
      let user: User = interaction.user; // get the user that sent the command
      let newConfig: Iconfig = updateUsers(user); // TODO: update user should be based on user selection
      console.log("execute newConfig: %s", newConfig.coin_product_alert_users); // __AUTO_GENERATED_PRINT_VAR__

      await interaction.reply(
         `<@${newConfig.coin_product_alert_users}> recorded`
      );
   },
};
