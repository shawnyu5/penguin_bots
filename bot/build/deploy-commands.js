"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// import { SlashCommandBuilder } from "@discordjs/builders";
const rest_1 = require("@discordjs/rest");
const v9_1 = require("discord-api-types/v9");
const { clientID, guildID, token } = require("../config.json");
const fs = require("fs");
// container for all our commands
const commands = [];
const commandFiles = fs
    .readdirSync(__dirname + "/commands")
    .filter((file) => file.endsWith(".js"));
for (const file of commandFiles) {
    const command = require(`${__dirname}/commands/${file}`);
    commands.push(command.data.toJSON());
}
const rest = new rest_1.REST({ version: "9" }).setToken(token);
rest
    .put(v9_1.Routes.applicationGuildCommands(clientID, guildID), { body: commands })
    .then(() => console.log("Successfully registered application commands."))
    .catch(console.error);
