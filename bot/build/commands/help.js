"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const builders_1 = require("@discordjs/builders");
const discord_js_1 = require("discord.js");
const deploy_commands_1 = require("../deploy-commands");
module.exports = {
    data: new builders_1.SlashCommandBuilder()
        .setName("help")
        .setDescription("help command")
        .addStringOption((option) => option
        .setName("command")
        .setDescription("name of command to get help page of")),
    async execute(interaction) {
        let userInput = String(interaction).split(":")[1];
        const onStart = new deploy_commands_1.OnStart();
        if (userInput) {
            let helpDocs = onStart.readAllHelpDocs();
            console.log("execute helpDocs: %s", JSON.stringify(helpDocs)); // __AUTO_GENERATED_PRINT_VAR__
            helpDocs.forEach((doc) => {
                if (doc && doc.name == userInput) {
                    let reply = new discord_js_1.MessageEmbed()
                        .setColor("RANDOM")
                        .setTitle("Help").setDescription(`
                                  Command name: ${doc.name}
                                  Description: ${doc.Description}
                                  Usage: ${doc.usage}
                                  `);
                    interaction.reply({ embeds: [reply] });
                }
            });
        }
        else {
            interaction.reply("Fuck you, google it");
        }
    },
    help: {
        name: "help",
        Description: "A help page for this bot",
        usage: "/help (command: <command name>)",
    },
};
