import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { CommandInteraction, Interaction, User } from "discord.js";
import { writeFileSync } from "fs";
import { environment } from "../enviroments/enviroment";
import { IConfig } from "../types/config";
import logger from "../logger";

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

   async execute(interaction: CommandInteraction) {
      let userChoice = String(interaction).split(":")[1];
      let user = interaction.user; // get the user that sent the command

      // TODO: fix this to use roles instead of storing users in config
      await interaction.reply("Renvating command... Try again later");
      // if (userChoice == "on") {
      // // let newconfig: IConfig = addUser(user);
      // // writeFileSync("./config.json", JSON.stringify(newconfig));

      // logger.info(`${user.username} has opted into product alerts`);
      // await interaction.reply(`${user} recorded`);
      // } else {
      // let newConfig = deleteUser(user);
      // writeFileSync("./config.json", JSON.stringify(newConfig));
      // logger.info(`${user.username} has opted out of product alerts`);
      // await interaction.reply(`${user} removed from notifications list`);
      // }
   },
   help: {
      name: "alert",
      Description: "Chose to opt in or out of coin product alerts",
      usage: "/alert notification: on | off",
   },
   // addUser: addUser,
   // deleteUser: deleteUser,
};
/**
 * Adds a user id to config object and return the new modified object
 *
 * @param user - user object
 * @returns new json config object with the user added
 */
// export function addUser(user: User): IConfig {
// let updatedConfig = environment;

// let users: Array<string> = updatedConfig.COIN_PRODUCT_ALERT_USERS;

// // look for the user id passes in in config.json array
// let found = users.findIndex((element) => {
// return element == user.id;
// });

// // only add user if it is not recorded
// if (found === -1) {
// users.push(user.id);
// console.log(`User ${user.username} successfully added`);
// }
// updatedConfig = { ...updatedConfig, COIN_PRODUCT_ALERT_USERS: users };
// // @ts-ignore
// return updatedConfig;
// }

/**
 * delete a user from config.json and return the newly modified object
 * @param user - user object to be deleted
 * @returns the updated config object with the user removed
 */
// export function deleteUser(user: User): IConfig {
// let updatedConfig = environment;

// let index = updatedConfig.COIN_PRODUCT_ALERT_USERS.findIndex((element) => {
// return element == user.id;
// });

// if (index >= 0) {
// updatedConfig.COIN_PRODUCT_ALERT_USERS.splice(index, 1);
// console.log(`User ${user.username} successfully removed`);
// }
// // @ts-ignore
// return updatedConfig;
// }
