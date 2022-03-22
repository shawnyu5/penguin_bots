import { ApplicationCommandType } from "discord-api-types";
import { MessageEmbed } from "discord.js";
const { SlashCommandBuilder } = require("@discordjs/builders");
import { Api } from "../api/api";
let config = require("../../config.json");

interface IProduct {
   _id: string;
   title: string;
   average_price: number;
   average_discount: number;
   appearances: number;
}

let api: any;
async function init() {
   api = new Api();
   await api.init(process.env.key);
}
init();

async function getProductDetail(keyword: string) {
   let productData: Array<IProduct> = await api.findNameByRegex(keyword);
   // console.log("getProductDetail productData: %s", JSON.stringify(productData)); // __AUTO_GENERATED_PRINT_VAR__
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

module.exports = {
   data: new SlashCommandBuilder()
      .setName("average")
      .setDescription("Replies the average product for a price")
      .addStringOption((option: any) =>
         option.setName("keyword").setDescription("A string").setRequired(true)
      ),

   async execute(interaction: any) {
      let userMessage = interaction.options._hoistedOptions[0].value;
      let api = new Api();

      // @ts-ignore
      await api.init(process.env.key);
      let response = await getProductDetail(userMessage);

      let message = new MessageEmbed()
         .setTitle(`Search term: ${userMessage}`)
         .setDescription(response)
         .setColor("RANDOM");
      await interaction.reply({ embeds: [message] });
   },

   help: {
      name: "average",
      Description: "Reteieves the average price based on a search keyword",
      usage: "/average keyword: <search word>",
   },
};
