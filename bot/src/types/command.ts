import { SlashCommandBuilder } from "@discordjs/builders";
import { CommandInteraction } from "discord.js";

export interface ICommand {
   // the slash command builder object
   data: Omit<SlashCommandBuilder, "addSubcommand" | "addSubcommandGroup">;
   execute(interaction: CommandInteraction): void;
   help: { name: string; description: string; usage: string };
}
