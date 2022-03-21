import { Utils } from "discord-api-types";
import { Client, Collection, Intents, MessageEmbed } from "discord.js";
require("dotenv").config();
const fs = require("fs");
require("./deploy-commands");
import { checkCoinProduct, getChannelByName } from "./utils";
import { IConfig } from "./types/config";
import { execSync } from "child_process";
import config from "../config.json";
import { ICoinProduct } from "./types/coinProduct";

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

/**
 * @param coinProduct - the coin product
 * @returns a message pining all users in config.json about the coinProduct
 */
function buildMessage(coinProduct: ICoinProduct): string {
   let message: string = "";
   let users = config.coin_product_alert_users.forEach((user) => {
      message += `<@${user}>
      `;
   });
   message += `title: ${coinProduct.title}
   url: ${coinProduct.url}`;
   return message;
}

client.on("ready", () => {
   // @ts-ignore
   console.log(`${client.user.tag} logged in`);

   // // run python script every 5 minutes
   // let execution = 0;
   setInterval(() => {
      let coinProduct = checkCoinProduct();
      let message: string = buildMessage(coinProduct);
      console.log("(anon) message: %s", message); // __AUTO_GENERATED_PRINT_VAR__
      let channel = getChannelByName(client, "notifications");

      if (channel) {
         let embed = new MessageEmbed()
            .setColor("RANDOM")
            .setTitle("Coin product alert")
            .setDescription(message);

         channel.send({ embeds: [embed] });
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

client.login(require("../config.json").token);
