"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const discord_js_1 = require("discord.js");
require("dotenv").config();
const fs_1 = __importDefault(require("fs"));
const deploy_commands_1 = require("./deploy-commands");
const utils_1 = require("./utils");
const config_json_1 = __importDefault(require("../config.json"));
const client = new discord_js_1.Client({
    intents: [discord_js_1.Intents.FLAGS.GUILDS, discord_js_1.Intents.FLAGS.GUILD_MESSAGES],
});
const onStart = new deploy_commands_1.OnStart();
//@ts-ignore
client.commands = new discord_js_1.Collection();
const commandFiles = fs_1.default
    .readdirSync(__dirname + "/commands")
    .filter((file) => file.endsWith(".js"));
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
    // let allCommands = onStart.readAllCommands();
    // onStart.registerCommands(config.clientID, guild.id, allCommands);
    let allCommands = onStart.readAllCommands();
    client.guilds.cache.forEach((guild) => {
        onStart.registerCommands(config_json_1.default.clientID, guild.id, allCommands);
    });
    setInterval(() => {
        let coinProduct = (0, utils_1.checkCoinProduct)();
        let message = (0, utils_1.buildMessage)(coinProduct);
        console.log("(anon) message: %s", message); // __AUTO_GENERATED_PRINT_VAR__
        let channel = (0, utils_1.getChannelByName)(client, "notifications");
        if (channel) {
            let embed = new discord_js_1.MessageEmbed()
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
client.on("guildCreate", function (guild) {
    let allCommands = onStart.readAllCommands();
    onStart.registerCommands(config_json_1.default.clientID, guild.id, allCommands);
});
client.login(require("../config.json").token);
