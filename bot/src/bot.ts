import { Client, Collection, Intents, MessageEmbed } from "discord.js";
import fs from "fs";
import { OnStart } from "./deploy-commands";
import { checkCoinProduct, getChannelByName, buildMessage } from "./utils";
import config from "./enviroments/config.json";
import axios, { AxiosResponse } from "axios";
import { ICoinProduct } from "./types/coinProduct";
import { enviroment } from "./enviroments/enviroment";

const client = new Client({
   intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES],
});

const onStart = new OnStart();

//@ts-ignore
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
   // @ts-ignore
   console.log(`${client.user.tag} logged in`);

   let allCommands = onStart.readAllCommands();
   client.guilds.cache.forEach((guild) => {
      onStart.registerCommands(config.clientID, guild, allCommands);
   });

   let interval = 0;
   setInterval(async () => {
      let response: AxiosResponse<any>;
      try {
         response = await axios.get(enviroment.api_address, {
            timeout: 5000,
         });
      } catch (e) {
         console.log("Axios error: " + e);
         return;
      }

      // @ts-ignore
      let coinProduct: CoinProduct = response.data;
      if (!coinProduct.IsValid) {
         console.error(`Product *${coinProduct.Title}* is not valid`);
         return;
      }

      // keep tack of execution count
      console.log(`Execution count: ${interval}`);
      interval++;

      let message: string = buildMessage(coinProduct);
      // let channel = getChannelByName(client, "notifications");
      let channel = getChannelByName(client, "development");
      if (channel) {
         channel.send(message);
      }
   }, 120000);
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
      console.error(error);
      await interaction.reply({
         content: "There was an error while executing this command!",
         ephemeral: true,
      });
   }
});

client.on("guildCreate", function (guild) {
   let allCommands = onStart.readAllCommands();
   onStart.registerCommands(config.clientID, guild, allCommands);
});

client.login(enviroment.token);
