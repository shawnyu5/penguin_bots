import {
   SlashCommandBuilder,
   SlashCommandStringOption,
} from "@discordjs/builders";
import { CommandInteraction, MessageEmbed } from "discord.js";
import { OnStart } from "../deploy-commands";
import { ICommand } from "../types/command";
import { IHelpDocs } from "../types/helpDocs";

export let obj: ICommand = {
   data: new SlashCommandBuilder()
      .setName("help")
      .setDescription("help command")
      .addStringOption((option: SlashCommandStringOption) =>
         option
            .setName("command")
            .setDescription("name of command to get help page of")
      ),

   async execute(interaction: CommandInteraction) {
      let userInput = String(interaction).split(":")[1];
      const onStart = new OnStart();
      if (userInput) {
         let helpDocs: Array<IHelpDocs> = onStart.readAllHelpDocs();
         console.log("execute helpDocs: %s", JSON.stringify(helpDocs)); // __AUTO_GENERATED_PRINT_VAR__
         helpDocs.forEach((doc) => {
            if (doc && doc.name == userInput) {
               let reply = new MessageEmbed()
                  .setColor("RANDOM")
                  .setTitle("Help").setDescription(`
                                  Command name: ${doc.name}
                                  Description: ${doc.Description}
                                  Usage: ${doc.usage}
                                  `);
               interaction.reply({ embeds: [reply] });
            }
         });
      } else {
         interaction.reply("Fuck you, google it");
      }
   },

   help: {
      name: "help",
      description: "A help page for this bot",
      usage: "/help (command: <command name>)",
   },
};
