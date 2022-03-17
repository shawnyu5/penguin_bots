"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
const builders_1 = require("@discordjs/builders");
module.exports = {
    data: new builders_1.SlashCommandBuilder()
        .setName("alert")
        .setDescription("opt into coin product alerts"),
    async execute(interaction) {
        console.log("execute interaction: %s", JSON.stringify(interaction.user)); // __AUTO_GENERATED_PRINT_VAR__
        // console.log(config);
        let user = interaction.user;
        await interaction.reply(`${user}`);
    },
};
