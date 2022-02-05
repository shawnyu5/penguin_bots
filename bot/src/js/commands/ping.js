"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const discord_js_1 = require("discord.js");
const { SlashCommandBuilder } = require("@discordjs/builders");
module.exports = {
    data: new SlashCommandBuilder()
        .setName("ping")
        .setDescription("Replies with Pong!"),
    async execute(interaction) {
        let embededMessage = new discord_js_1.MessageEmbed()
            .setColor("#0099ff")
            .setTitle("apple!")
            .setDescription("Pong!!!")
            .setTitle("PONG");
        console.log(interaction.channel);
        await interaction.reply({ embeds: [embededMessage] });
    },
};
