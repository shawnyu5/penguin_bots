import { CommandInteraction, MessageEmbed } from "discord.js";
import { SlashCommandBuilder } from "@discordjs/builders";
import { DataBase } from "../database/database";
import IDbProduct from "../types/dbProduct";
import logger from "../logger";
import config from "../enviroments/config.json";

// interface IProduct {
// title: string;
// appeartitle: string;
// average_discount: number;
// average_price: number;
// created_date: Date;
// updated_date: Date;
// }

module.exports = {
   data: new SlashCommandBuilder()
      .setName("average")
      .setDescription("Replies the average product for a price")
      .addStringOption((option: any) =>
         option
            .setName("keyword")
            .setDescription("The product you want to search for")
            .setRequired(true)
      ),

   async execute(interaction: CommandInteraction) {
      await interaction.deferReply();
      let userMessage = interaction.options.getString("keyword");

      let db = new DataBase(config.MONGOOSE_KEY);

      let response = "";

      try {
         response = await getProductDetail(userMessage as string, db);
      } catch (error) {
         response = "Error: " + error;
      }

      let message = new MessageEmbed()
         .setTitle(`Search term: ${userMessage}`)
         .setDescription(response)
         .setColor("RANDOM");
      await interaction.editReply({ embeds: [message] });
      logger.info(`Replied to search term: ${userMessage}`);
      // await interaction.editReply("hello????");
   },

   help: {
      name: "average",
      Description: "Retrieves the average price based on a search keyword",
      usage: "/average keyword: <search word>",
   },
};

/**
 * searches the database based on a search string
 * @param keyword - the search string
 * @returns a string response of the search result from database
 */
async function getProductDetail(
   keyword: string,
   db: DataBase
): Promise<string> {
   let productData: Array<IDbProduct> = await db.findNameByRegex(keyword);
   let response: string = "";

   // if products are found
   if (productData) {
      // get the first index of array
      productData.forEach((product: IDbProduct) => {
         response += `\
**title**: ${product.title}
**average price**: ${product.average_price}
**average discount**: ${product.average_discount}
**appearances**: ${product.appearances}

`;
      });
   }
   return response;
}
