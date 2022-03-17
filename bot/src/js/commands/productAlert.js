"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
const builders_1 = require("@discordjs/builders");
const exec = require("child_process").exec;
module.exports = {
    data: new builders_1.SlashCommandBuilder()
        .setName("alert")
        .setDescription("opt into coin product alerts"),
    async execute(interaction) {
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
        let user = interaction.user;
        await interaction.reply(`${user} recorded`);
    },
};
