import { Client, Collection, Intents } from "discord.js";
require("dotenv").config();
const fs = require("fs");
require("./deploy-commands");
import { checkCoinProduct } from "./utils";

const client = new Client({
   intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES],
});

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

   setInterval(() => {
      let coinProduct = checkCoinProduct();
      console.log("(anon)#(anon) coinProduct: %s", coinProduct); // __AUTO_GENERATED_PRINT_VAR__
   }, 5000);
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

client.login(require("../config.json").token);
