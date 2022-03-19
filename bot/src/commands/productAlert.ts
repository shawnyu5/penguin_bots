// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { Interaction, User } from "discord.js";
import { writeFileSync, readFileSync, WriteFileOptions } from "fs";
import { Error } from "mongoose";
import config from "../../config.json";
const exec = require("child_process").exec;

interface IConfig {
   token: string;
   guildID: string;
   clientID: string;
   coin_product_alert_users: Array<string>;
}

/**
 * Adds a user id to config object and return the new modified object
 *
 * @param user - id of user to be added to alert
 * @returns new json config object with the user added
 */
function addUser(user: any): IConfig {
   let updatedConfig: IConfig = config;

   let users: Array<string> = updatedConfig.coin_product_alert_users;

   let found = users.find((element) => {
      element == user.id;
   });

   // only add user if it is not recorded right now
   if (!found) {
      users.push(user.id);
   }
   updatedConfig = { ...updatedConfig, coin_product_alert_users: users };
   return updatedConfig;
}

/**
 * @param user - the user to be deleted
 * @returns the updated config object with the user removed
 */
function deleteUser(user: any): IConfig {
   let updatedConfig: IConfig = config;
   console.log(updatedConfig); // __AUTO_GENERATED_PRINT_VAR__
   let index = updatedConfig.coin_product_alert_users.findIndex(
      (element) => element == user.id
   );

   if (index >= 0) {
      updatedConfig.coin_product_alert_users.splice(index, 1);
   }
   return updatedConfig;
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

      if (userChoice == "on") {
         let newConfig: IConfig = addUser(user);

         console.log(
            "execute#if JSON.stringify(newConfig): %s",
            JSON.stringify(newConfig)
         ); // __AUTO_GENERATED_PRINT_VAR__
         writeFileSync(
            "./config.json",
            JSON.stringify(newConfig),
            (err: any) => {
               if (!err) {
                  console.log("Config.json updated");
               } else {
                  console.log("Config.json failed up date");
               }
            }
         );

         await interaction.reply(`${user} recorded`);
      } else {
         let newConfig = deleteUser(user);
         writeFileSync(
            "./config.json",
            JSON.stringify(newConfig),
            (err: WriteFileOptions | undefined) => {
               if (!err) {
                  console.log("Config.json updated");
               } else {
                  console.log("Config.json failed up date");
               }
            }
         );
         await interaction.reply("Nothing to do!");
      }
   },
};
