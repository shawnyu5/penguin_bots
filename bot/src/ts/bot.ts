import { Client, Collection, Intents } from "discord.js";
require("dotenv").config();
const fs = require("fs");
require("./deploy-commands");

const client = new Client({
   intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES],
});

client.commands = new Collection();

const commandFiles = fs
   .readdirSync(__dirname + "/commands")
   .filter((file: string) => file.endsWith(".js"));

for (const file of commandFiles) {
   const command = require(`./commands/${file}`);
   // Set a new item in the Collection
   // With the key as the command name and the value as the exported module
   client.commands.set(command.data.name, command);
}

client.on("ready", () => {
   // @ts-ignore
   console.log(`${client.user.tag} logged in`);
});

client.on("interactionCreate", async (interaction) => {
   if (!interaction.isCommand()) return;
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

// console.log(require("../../config.json").token);
client.login(require("../../config.json").token);
