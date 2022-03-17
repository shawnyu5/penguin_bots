// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
import { SlashCommandBuilder } from "@discordjs/builders";
import { writeFile } from "fs";
import config from "../../../config.json";

module.exports = {
   data: new SlashCommandBuilder()
      .setName("alert")
      .setDescription("opt into coin product alerts"),

   async execute(interaction: any) {
      console.log("execute interaction: %s", JSON.stringify(interaction.user)); // __AUTO_GENERATED_PRINT_VAR__
      // console.log(config);
      let user: string = interaction.user;
      await interaction.reply(`${user}`);
   },
};
