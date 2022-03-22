import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { Interaction, MessageEmbed } from "discord.js";
import { OnStart } from "../deploy-commands";
import { IHelpDocs } from "../types/helpDocs";

module.exports = {
   data: new SlashCommandBuilder()
      .setName("help")
      .setDescription("help command")
      .addStringOption((option: SlashCommandStringOption) =>
         option
            .setName("command")
            .setDescription("name of command to get help page of")
      ),

   async execute(interaction: Interaction) {
      let userInput = String(interaction).split(":")[1];
      console.log("execute Interaction: %s", interaction); // __AUTO_GENERATED_PRINT_VAR__
      const onStart = new OnStart();
      if (userInput) {
         let helpDocs: Array<IHelpDocs> = onStart.readAllHelpDocs();
         console.log("execute helpDocs: %s", JSON.stringify(helpDocs)); // __AUTO_GENERATED_PRINT_VAR__
         helpDocs.forEach((doc) => {
            console.log("execute#if#(anon) doc: %s", doc); // __AUTO_GENERATED_PRINT_VAR__
            if (doc && doc.name == userInput) {
               let reply = new MessageEmbed()
                  .setColor("RANDOM")
                  .setTitle("Help").setDescription(`
                                  Command name: ${doc.name}
                                  Description: ${doc.Description}
                                  Usage: ${doc.usage}
                                  `);
               interaction.reply({ embeds: [reply] });
               return;
            }
         });
      }
      // interaction.reply("hi");
   },
};
