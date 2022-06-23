import { SlashCommandBuilder } from "@discordjs/builders";

module.exports = {
   data: new SlashCommandBuilder()
      .setName("sub_to_carmen")
      .setDescription("Subscribes to Carmen's ramblings for free!"),
};
