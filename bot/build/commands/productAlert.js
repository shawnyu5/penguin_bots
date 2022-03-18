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
 * @param user - id of user to be added to alert
 * @returns new json config object with the user added
 */
function addUser(user) {
    let updatedConfig = config_json_1.default;
    let users = updatedConfig.coin_product_alert_users;
    // only add user if it is not recorded right now
    let found = users.find((element) => {
        element == user.id;
    });
    if (found) {
        users.push(user.id);
    }
    updatedConfig = { ...updatedConfig, coin_product_alert_users: users };
    return updatedConfig;
}
// TODO: implement function to delete a user
function deleteUser(user) { }
function checkCoinProduct() {
    let output;
    // const channel = <client>.channels.cache.get('<id>');
    // channel.send("<content>");
    exec("python3 /home/shawn/python/penguin_bots/coin_products/coin_products.py", (err, stdout, stderr) => {
        console.log("execute#(anon) err: %s", err.code); // __AUTO_GENERATED_PRINT_VAR__
        console.log("(anon) stdout: %s", stdout); // __AUTO_GENERATED_PRINT_VAR__
        // only record output if script exited successfull
        if (err.code == 0) {
            output = stdout;
        }
    });
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
            // console.log("execute newConfig: %s", newConfig); // __AUTO_GENERATED_PRINT_VAR__
            console.log((0, fs_1.readFileSync)("./config.json", "utf8"));
            (0, fs_1.writeFileSync)("./config.json", JSON.stringify(newConfig), (err) => {
                if (!err) {
                    console.log("Config.json updated");
                }
                else {
                    console.log("Config.json failed up date");
                }
            });
            await interaction.reply(`<@${user}> recorded`);
        }
        else {
            await interaction.reply("Nothing to do!");
        }
    },
};
