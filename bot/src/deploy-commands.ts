import { REST } from "@discordjs/rest";
import { Routes } from "discord-api-types/v9";
import { clientID, guildID, token } from "../config.json";
import fs from "fs";
import { IHelpDocs } from "./types/helpDocs";

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
    * @param guildID - guildID
    * @param commands - array of commands
    */
   registerCommands(clientID: string, guildID: string, commands: Array<any>) {
      const rest = new REST({ version: "9" }).setToken(token);
      rest
         .put(Routes.applicationGuildCommands(clientID, guildID), {
            body: commands,
         })
         .then(() =>
            console.log("Successfully registered application commands.")
         )
         .catch(console.error);
   }
}

export { OnStart };
