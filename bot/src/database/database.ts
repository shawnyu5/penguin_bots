import { connect, connection, model, Schema } from "mongoose";
import logger from "../logger";
import IDbProduct from "../types/dbProduct";

const productSchema = new Schema({
   title: String,
   appeartitle: String,
   average_discount: Number,
   average_price: Number,
   created_date: Date,
   updated_date: Date,
});

// const openBoxModel = model("open_box", productSchema, "open_box");
export class DataBase {
   #openBoxModel = model("open_box", productSchema, "open_box");

   constructor(connectionString: string) {
      connect(connectionString, {
         // @ts-ignore
         useNewUrlParser: true,
         useUnifiedTopology: true,
      });

      connection.on("error", function () {
         logger.error("Error connecting to database");
         throw new Error("Error connecting to database");
      });

      connection.once("open", function () {
         logger.info("Connected to data base");
      });
   }

   // return a product object by name exact name
   async findByName(searchTerm: string) {
      this.#openBoxModel.find({ title: searchTerm }, (err: any, data: any) => {
         if (err) {
            Promise.reject(err);
         }
         Promise.resolve(data);
      });
   }

   /**
    * search through the database by regex search string
    * @param title - the search string
    * @returns an array of products from data base
    */
   async findNameByRegex(title: string | RegExp): Promise<IDbProduct[]> {
      // convert title to case insenitive regular expression
      title = new RegExp(title, "i");
      let response = await this.#openBoxModel
         .find({
            title: title,
         })
         .lean() // convert mongoose object to plain javascript object
         .exec();

      // logger.debug(JSON.stringify(response.slice(0, 10), null, 2));
      // logger.debug(`Response 0: ${response[0]}`);
      // logger.debug(response[0].average_price);

      return Promise.resolve(response.slice(0, 10));
   }
}
