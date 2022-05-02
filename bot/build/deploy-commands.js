"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.OnStart = void 0;
const rest_1 = require("@discordjs/rest");
const v9_1 = require("discord-api-types/v9");
const config_json_1 = require("../config.json");
const fs_1 = __importDefault(require("fs"));
class OnStart {
    /**
     * @returns all commands contained in `/commands`
     */
    readAllCommands() {
        const commands = [];
        const commandFiles = fs_1.default
            .readdirSync(__dirname + "/commands")
            .filter((file) => file.endsWith(".js"));
        for (const file of commandFiles) {
            const command = require(`${__dirname}/commands/${file}`);
            commands.push(command.data.toJSON());
        }
        return commands;
    }
    /**
     * read all help docs from command modules and store in array
     * @returns json array of help docs
     */
    readAllHelpDocs() {
        const helpDocs = [];
        const commandFiles = fs_1.default
            .readdirSync(__dirname + "/commands")
            .filter((file) => file.endsWith(".js"));
        for (const file of commandFiles) {
            const command = require(`${__dirname}/commands/${file}`);
            helpDocs.push(command.help);
        }
        // console.log(
        // "OnStart#readAllHelpDocs helpDocs: %s",
        // JSON.stringify(helpDocs)
        // ); // __AUTO_GENERATED_PRINT_VAR__
        return helpDocs;
    }
    /**
     * @param clientID - ClientID
     * @param guild - guildID
     * @param commands - array of commands
     */
    registerCommands(clientID, guild, commands) {
        const rest = new rest_1.REST({ version: "9" }).setToken(config_json_1.token);
        (async () => {
            try {
                console.log(`Started refreshing application (/) commands for ${guild.name}`);
                if (!global) {
                    await rest.put(v9_1.Routes.applicationGuildCommands(clientID, guild.id), {
                        body: commands,
                    });
                }
                else {
                    await rest.put(v9_1.Routes.applicationCommands(clientID), {
                        body: commands,
                    });
                }
                console.log(`Successfully reloaded application (/) commands for ${guild.name}`);
            }
            catch (error) {
                console.error(error);
            }
        })();
    }
}
exports.OnStart = OnStart;
