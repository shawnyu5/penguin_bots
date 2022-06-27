import { CommandInteraction, MessageEmbed } from "discord.js";
import { SlashCommandBuilder } from "@discordjs/builders";
import { Api } from "../api/api";
import DbProduct from "../types/dbProduct";
import { enviroment } from "../enviroments/enviroment";

interface IProduct {
   _id: string;
   title: string;
   average_price: number;
   average_discount: number;
   appearances: number;
}

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
      // let api = new Api(process.env.MONGOOSE_KEY);

      // let response;
      // try {
      // await getProductDetail(userMessage as string, api);
      // } catch (error) {
      // await getProductDetail(userMessage as string, api);
      // }

      // let message = new MessageEmbed()
      // .setTitle(`Search term: ${userMessage}`)
      // .setDescription(response)
      // .setColor("RANDOM");
      // await interaction.editReply({ embeds: [message] });
      await interaction.editReply("NOT IMPLEMENTED");
   },

   help: {
      name: "average",
      Description: "Retrieves the average price based on a search keyword",
      usage: "/average keyword: <search word>",
   },
};

/**
 * searches the data base based on a search string
 * @param keyword - the search string
 * @returns the search result from data base
 */
async function getProductDetail(keyword: string, api: Api) {
   let productData: Array<DbProduct> = await api.findNameByRegex(keyword);
   let response: string = "";

   // if a single product is found
   if (productData.length == 1) {
      // get the first index of array
      let product: IProduct = productData[0];
      response = `title: ${product.title}
      average price: ${product.average_price}
      average discount: ${product.average_discount}
      appearances: ${product.appearances}`;
   }
   // no product is found
   else if (productData.length == 0) {
      response = "No product found";
   }
   // if an array of product is found
   else {
      productData.forEach((element) => {
         let currentResponse = `title: ${element.title}
         average discount: ${element.average_discount}
         average price: ${element.average_price}
         appearances: ${element.appearances}

         `;
         response = response.concat(" ", currentResponse);
      });
   }
   return response;
}
