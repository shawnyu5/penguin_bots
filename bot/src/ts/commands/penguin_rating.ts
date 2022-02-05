import { ApplicationCommandType } from "discord-api-types";
import { MessageEmbed } from "discord.js";
const { SlashCommandBuilder } = require("@discordjs/builders");

module.exports = {
   data: new SlashCommandBuilder()
      .setName("average")
      .setDescription("Replies the average product for a price")
      .addStringOption((option: any) =>
         option
            .setName("category")
            .setDescription("The gif category")
            .setRequired(false)
      ),

   async execute(interaction: any) {
      let userMessage = interaction.options._hoistedOptions[0].value;
      let message = new MessageEmbed()
         .setTitle(`Average price for ${userMessage}`)
         .setColor("RANDOM");
      console.log(userMessage);
      await interaction.reply({ embeds: [message] });
   },
};
