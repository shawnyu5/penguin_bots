"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const discord_js_1 = require("discord.js");
require("dotenv").config();
const fs = require("fs");
require("./deploy-commands");
const utils_1 = require("./utils");
const config_json_1 = __importDefault(require("../config.json"));
const client = new discord_js_1.Client({
    intents: [discord_js_1.Intents.FLAGS.GUILDS, discord_js_1.Intents.FLAGS.GUILD_MESSAGES],
});
//@ts-ignore
client.commands = new discord_js_1.Collection();
const commandFiles = fs
    .readdirSync(__dirname + "/commands")
    .filter((file) => file.endsWith(".js"));
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
function buildMessage(coinProduct) {
    let message = "";
    let users = config_json_1.default.coin_product_alert_users.forEach((user) => {
        message += `<@${user}> `;
    });
    message += coinProduct;
    return message;
}
client.on("ready", () => {
    // @ts-ignore
    console.log(`${client.user.tag} logged in`);
    // run python script every 5 minutes
    let execution = 0;
    setInterval(() => {
        try {
            let coinProduct = (0, utils_1.checkCoinProduct)();
            let message = buildMessage("This is a coin product");
            (0, utils_1.sendMessage)(client, "notifications", message);
        }
        catch (error) {
            console.log("ERROR: " + error);
        }
        console.log(`Execution ${execution}`);
        execution++;
    }, 120000);
    // 120000 - 2 minutes in milliseconds
    // 300000 - 5 mins in milliseconds
});
client.on("interactionCreate", async (interaction) => {
    if (!interaction.isCommand())
        return;
    // @ts-ignore
    const command = client.commands.get(interaction.commandName);
    if (!command)
        return;
    try {
        await command.execute(interaction);
    }
    catch (error) {
        console.error(error);
        await interaction.reply({
            content: "There was an error while executing this command!",
            ephemeral: true,
        });
    }
});
client.login(require("../config.json").token);
