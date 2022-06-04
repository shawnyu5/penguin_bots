import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { Interaction, User } from "discord.js";
import { writeFileSync } from "fs";
import config from "../enviroments/config.json";
import { IConfig } from "../types/config";

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
      console.log("execute user: %s", JSON.stringify(user)); // __AUTO_GENERATED_PRINT_VAR__

      if (userChoice == "on") {
         let newConfig: IConfig = addUser(user);
         writeFileSync("./config.json", JSON.stringify(newConfig));

         await interaction.reply(`${user} recorded`);
      } else {
         console.log("execute#if user: %s", user); // __AUTO_GENERATED_PRINT_VAR__
         let newConfig = deleteUser(user);
         writeFileSync("./config.json", JSON.stringify(newConfig));
         await interaction.reply(`${user} removed from notifications list`);
      }
   },
   help: {
      name: "alert",
      Description: "Chose to opt in or out of coin product alerts",
      usage: "/alert notification: on | off",
   },
   addUser: addUser,
   deleteUser: deleteUser,
};
/**
 * Adds a user id to config object and return the new modified object
 *
 * @param user - user object
 * @returns new json config object with the user added
 */
export function addUser(user: User): IConfig {
   let updatedConfig: IConfig = config;

   let users: Array<string> = updatedConfig.coin_product_alert_users;

   // look for the user id passes in in config.json array
   let found = users.findIndex((element) => {
      return element == user.id;
   });

   // only add user if it is not recorded
   if (found === -1) {
      users.push(user.id);
      console.log(`User ${user.username} successfully added`);
   }
   updatedConfig = { ...updatedConfig, coin_product_alert_users: users };
   return updatedConfig;
}

/**
 * delete a user from config.json and return the newly modified object
 * @param user - user object to be deleted
 * @returns the updated config object with the user removed
 */
export function deleteUser(user: User): IConfig {
   let updatedConfig: IConfig = config;

   let index = updatedConfig.coin_product_alert_users.findIndex((element) => {
      return element == user.id;
   });

   if (index >= 0) {
      updatedConfig.coin_product_alert_users.splice(index, 1);
      console.log(`User ${user.username} successfully removed`);
   }
   return updatedConfig;
}
