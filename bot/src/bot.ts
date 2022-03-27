import { Client, Collection, Intents, MessageEmbed } from "discord.js";
require("dotenv").config();
import fs from "fs";
import { OnStart } from "./deploy-commands";
import { checkCoinProduct, getChannelByName, buildMessage } from "./utils";
import config from "../config.json";

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
      onStart.registerCommands(config.clientID, guild.id, allCommands);
   });

   let interval = 0;
   setInterval(() => {
      console.log(`Execution count: ${interval}`);
      interval++;

      let coinProduct = checkCoinProduct();
      console.log("(anon) coinProduct : %s", coinProduct); // __AUTO_GENERATED_PRINT_VAR__
      if (!coinProduct) {
         console.log("Invaid coin product");
         return;
      }

      let message: string = buildMessage(coinProduct);
      console.log("(anon) message: %s", message); // __AUTO_GENERATED_PRINT_VAR__
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
   onStart.registerCommands(config.clientID, guild.id, allCommands);
});

client.login(require("../config.json").token);
