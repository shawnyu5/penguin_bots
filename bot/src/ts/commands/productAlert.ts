// import { ApplicationCommandType } from "discord-api-types";
// import { MessageEmbed } from "discord.js";
import { SlashCommandBuilder } from "@discordjs/builders";
import { writeFile } from "fs";
import config from "../../../config.json";
const exec = require("child_process").exec;

module.exports = {
   data: new SlashCommandBuilder()
      .setName("alert")
      .setDescription("opt into coin product alerts"),

   async execute(interaction: any) {
      let output: string;
      // const channel = <client>.channels.cache.get('<id>');
      // channel.send("<content>");
      exec(
         "python3 /home/shawn/python/penguin_bots/coin_products/coin_products.py",
         (err: any, stdout: any, stderr: any) => {
            console.log("execute#(anon) err: %s", err.code); // __AUTO_GENERATED_PRINT_VAR__
            console.log("(anon) stdout: %s", stdout); // __AUTO_GENERATED_PRINT_VAR__

            // only record output if script exited successfull
            if (err.code == 0) {
               output = stdout;
            }
         }
      );

      let user: string = interaction.user;
      await interaction.reply(`${user} recorded`);
   },
};
