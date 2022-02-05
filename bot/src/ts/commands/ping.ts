import { MessageEmbed } from "discord.js";

const { SlashCommandBuilder } = require("@discordjs/builders");

module.exports = {
   data: new SlashCommandBuilder()
      .setName("ping")
      .setDescription("Replies with Pong!"),
   async execute(interaction: any) {
      let embededMessage = new MessageEmbed()
         .setColor("#0099ff")
         .setTitle("apple!")
         .setDescription("Pong!!!")
         .setTitle("PONG");

      console.log(interaction.channel);
      await interaction.reply({ embeds: [embededMessage] });
   },
};
