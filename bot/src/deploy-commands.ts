import { REST } from "@discordjs/rest";
import { Routes } from "discord-api-types/v9";
import fs from "fs";
import { IHelpDocs } from "./types/helpDocs";
import { Guild } from "discord.js";
import logger from "./logger";

class OnStart {
   /**
    * @returns all commands contained in `/commands`
    */
   readAllCommands() {
      const commands = [];
      const commandFiles = fs
         .readdirSync(__dirname + "/commands")
         .filter((file: string) => file.endsWith(".js"));
      for (const file of commandFiles) {
         const command = require(`${__dirname}/commands/${file}`);
         commands.push(command.data.toJSON());
      }
      return commands;
   }

   /**
    * read all help docs from command modules and store in array
    * @returns json array of help docs
    */
   readAllHelpDocs(): Array<IHelpDocs> {
      const helpDocs = [];
      const commandFiles = fs
         .readdirSync(__dirname + "/commands")
         .filter((file: string) => file.endsWith(".js"));
      for (const file of commandFiles) {
         const command = require(`${__dirname}/commands/${file}`);
         helpDocs.push(command.help);
      }
      // console.log(
      // "OnStart#readAllHelpDocs helpDocs: %s",
      // JSON.stringify(helpDocs)
      // ); // __AUTO_GENERATED_PRINT_VAR__
      return helpDocs;
   }
   /**
    * @param clientID - ClientID
    * @param guild - guildID
    * @param commands - array of commands
    */
   registerCommands(clientID: string, guild: Guild, commands: any): void {
      const rest = new REST({ version: "9" }).setToken(process.env.TOKEN);
      (async () => {
         try {
            logger.info(
               `Started refreshing application (/) commands for ${guild.name}`
            );

            if (!global) {
               await rest.put(
                  Routes.applicationGuildCommands(clientID, guild.id),
                  {
                     body: commands,
                  }
               );
            } else {
               await rest.put(Routes.applicationCommands(clientID), {
                  body: commands,
               });
            }

            logger.info(
               `Successfully reloaded application (/) commands for ${guild.name}`
            );
         } catch (error) {
            logger.error(error);
         }
      })();
   }
}

export { OnStart };
