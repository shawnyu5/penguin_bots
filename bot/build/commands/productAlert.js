"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
const builders_1 = require("@discordjs/builders");
const fs_1 = require("fs");
const config_json_1 = __importDefault(require("../../config.json"));
const exec = require("child_process").exec;
/**
 * Adds a user id to config object and return the new modified object
 *
 * @param user - user object
 * @returns new json config object with the user added
 */
function addUser(user) {
    let updatedConfig = config_json_1.default;
    let users = updatedConfig.coin_product_alert_users;
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
 * delete a user from config object and return the newly modified object
 * @param user - user object to be deleted
 * @returns the updated config object with the user removed
 */
function deleteUser(user) {
    let updatedConfig = config_json_1.default;
    let index = updatedConfig.coin_product_alert_users.findIndex((element) => element == user.id);
    if (index >= 0) {
        updatedConfig.coin_product_alert_users.splice(index, 1);
        console.log(`User ${user.username} successfully removed`);
    }
    return updatedConfig;
}
module.exports = {
    data: new builders_1.SlashCommandBuilder()
        .setName("alert")
        .setDescription("opt into coin product alerts")
        .addStringOption((option) => option
        .setName("notification")
        .setDescription("Choose weather to opt in or out of coin product notifications")
        .setRequired(true)
        .addChoice("on", "on")
        .addChoice("off", "off")),
    async execute(interaction) {
        let userChoice = String(interaction).split(":")[1];
        let user = interaction.user; // get the user that sent the command
        if (userChoice == "on") {
            let newConfig = addUser(user);
            (0, fs_1.writeFileSync)("./config.json", JSON.stringify(newConfig));
            await interaction.reply(`${user} recorded`);
        }
        else {
            let newConfig = deleteUser(user);
            (0, fs_1.writeFileSync)("./config.json", JSON.stringify(newConfig));
            await interaction.reply(`${user} removed from notifications list`);
        }
    },
};
