import { Client, Collection, Intents, TextChannel } from "discord.js";
import fs from "fs";
import { OnStart } from "./deploy-commands";
import { getChannelByName, buildMessage } from "./utils";
import { environment } from "./enviroments/enviroment";
import axios, { AxiosResponse } from "axios";
import logger from "./logger";
import { ICoinProduct } from "./types/coinProduct";
import { DataBase } from "./database/database";
import * as dotenv from "dotenv";
dotenv.config();

const client = new Client({
   intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES],
});

const onStart = new OnStart();

// @ts-ignore
client.commands = new Collection();

const commandFiles = fs
   .readdirSync(__dirname + "/commands")
   .filter((file: string) => file.endsWith(".js"));

for (const file of commandFiles) {
   const command = require(`./commands/${file}`);
   // Set a new item in the Collection
   // With the key as the command name and the value as the exported module

   // @ts-ignore
   client.commands.set(command.data.name, command);
}

client.on("ready", () => {
   logger.info(`${client.user?.tag} logged in`);
   let db = new DataBase(environment.MONGOOSE_KEY as string);

   let allCommands = onStart.readAllCommands();
   client.guilds.cache.forEach((guild) => {
      onStart.registerCommands(
         environment.CLIENTID as string,
         guild,
         allCommands
      );
   });

   let interval = 0;
   setInterval(async () => {
      // keep tack of execution count
      logger.debug(`Execution count: ${interval}`);
      interval++;
      let response: AxiosResponse<any>;
      try {
         response = await axios.get(`${environment.API_ADDRESS}/coinProduct`, {
            timeout: 5000,
         });
      } catch (e) {
         logger.error("Axios error: " + e);
         return;
      }

      let coinProduct: ICoinProduct = response.data;
      if (!coinProduct.IsValid) {
         logger.warn(
            `Product *${coinProduct.Title}* is not valid: ${coinProduct.Reason}`
         );
         return;
      }

      let message: string = buildMessage(coinProduct);
      let channel: TextChannel | undefined;
      if (environment.DEVELOPMENT) {
         channel = getChannelByName(client, "development");
      } else {
         channel = getChannelByName(client, "notifications");
      }
      if (channel) {
         channel.send(message);
      }
   }, 12000);
   // 120000 - 2 minutes in milliseconds
   // 300000 - 5 mins in milliseconds
});

client.on("interactionCreate", async (interaction) => {
   if (!interaction.isCommand()) return;
   // @ts-ignore
   const command = client.commands.get(interaction.commandName);

   if (!command) return;

   try {
      await command.execute(interaction);
   } catch (error) {
      logger.error(error);
      await interaction.reply({
         content: "There was an error while executing this command!",
         ephemeral: true,
      });
   }
});

client.on("guildCreate", function (guild) {
   let allCommands = onStart.readAllCommands();
   onStart.registerCommands(environment.CLIENTID as string, guild, allCommands);
});

client.login(environment.TOKEN);
