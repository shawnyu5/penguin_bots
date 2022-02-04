import { Client, Intents } from "discord.js";
require("dotenv").config();

const client = new Client({
   intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MESSAGES],
});

client.on("ready", () => {
   // @ts-ignore
   console.log(`${client.user.tag} logged in`);
});

client.login(process.env.token);
