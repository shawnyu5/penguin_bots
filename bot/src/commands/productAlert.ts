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
import { IConfig } from "../types/config";

/**
 * Adds a user id to config object and return the new modified object
 *
 * @param user - user object
 * @returns new json config object with the user added
 */
function addUser(user: any): IConfig {
   let updatedConfig: IConfig = config;

   let users: Array<string> = updatedConfig.coin_product_alert_users;

   // look for the user id passes in in config.json array
   let found = users.find((element) => {
      element == user.id;
   });

   // only add user if it is not recorded
   if (!found) {
      users.push(user.id);
      console.log(`User ${user.username} successfully added`);
   }
   updatedConfig = { ...updatedConfig, coin_product_alert_users: users };
   return updatedConfig;
}

/**
 * @param user - user object to be deleted
 * @returns the updated config object with the user removed
 */
function deleteUser(user: any): IConfig {
   let updatedConfig: IConfig = config;
   let index = updatedConfig.coin_product_alert_users.findIndex(
      (element) => element == user.id
   );

   if (index >= 0) {
      updatedConfig.coin_product_alert_users.splice(index, 1);
      console.log(`User ${user.username} successfully removed`);
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
         writeFileSync("./config.json", JSON.stringify(newConfig));

         await interaction.reply(`${user} recorded`);
      } else {
         let newConfig: IConfig = deleteUser(user);
         writeFileSync("./config.json", JSON.stringify(newConfig));
         await interaction.reply(`${user} removed from notifications list`);
      }
   },
};
