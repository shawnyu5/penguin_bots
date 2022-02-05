"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const discord_js_1 = require("discord.js");
const { SlashCommandBuilder } = require("@discordjs/builders");
module.exports = {
    data: new SlashCommandBuilder()
        .setName("average")
        .setDescription("Replies the average product for a price")
        .addStringOption((option) => option
        .setName("category")
        .setDescription("The gif category")
        .setRequired(false)),
    async execute(interaction) {
        let userMessage = interaction.options._hoistedOptions[0].value;
        let message = new discord_js_1.MessageEmbed()
            .setTitle(`Average price for ${userMessage}`)
            .setColor("RANDOM");
        console.log(userMessage);
        await interaction.reply({ embeds: [message] });
    },
};
