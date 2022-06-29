import { connect, Schema, Types } from "mongoose";
import logger from "../logger";
import DbProduct from "../types/dbProduct";

const productSchema = new Schema({
   title: String,
   average_discount: Number,
   average_price: Number,
   appearances: Number,
});

export class Api {
   db: any;

   constructor(connectionString: string) {
      this.db = connect(connectionString, {
         // @ts-ignore
         useNewUrlParser: true,
         useFindAndModify: false,
         useUnifiedTopology: true,
      });

      this.db.on("error", logger.error("Error connecting to database"));

      this.db.once("open", function () {
         logger.info("Connected to data base");
      });
   }

   /**
    * Creates a connection to the database
    * @param connectionString - the connection string to the database
    */
   // async createConnection(connectionString: string): Promise<void> {
   // const db = createConnection(connectionString, {
   // // @ts-ignore
   // useNewUrlParser: true,
   // useUnifiedTopology: true,
   // });

   // db.once("error", (err) => {
   // throw new Error(err);
   // });

   // db.once("open", () => {
   // this.db = db.model("Open_box", productSchema, "open_box");
   // console.log("Api#createConnection#(anon) this: %s", this.db); // __AUTO_GENERATED_PRINT_VAR__
   // console.log("Connected to data base");
   // Promise.resolve();
   // });
   // }

   // return a product object by name exact name
   async findByName(searchTerm: string) {
      // return this.open_box.findOne(name).exec();
      return new Promise((resolve, reject) => {
         this.db.find(searchTerm, (err: any, data: any) => {
            if (err) {
               reject(err);
            }
            resolve(data);
         });
      });
   }

   /**
    * search through the database by regex search string
    * @param title - the search string
    * @returns an array of products from data base
    */
   async findNameByRegex(title: string | RegExp): Promise<Array<DbProduct>> {
      return new Promise((resolve, reject) => {
         // convert title to case insenitive regular expression
         title = new RegExp(title, "i");
         this.db.find(
            {
               title: title,
            },
            (error: any, data: any) => {
               if (error) {
                  reject(error);
               }
               resolve(data);
            }
         );
      });
   }
}

// async function main() {
// let api = new Api();
// try {
// // @ts-ignore
// await api.init(process.env.key);
// let obj = {
// _id: new Types.ObjectId("61dceb6228b23db27260d4e0"),
// title: "Play Money by Nick Diffatte (Instant Download)",
// average_discount: 33.333333333333336,
// average_price: 3.3000000000000003,
// appearances: 3,
// };

// let data = await api.findNameByRegex("nick diffatte");
// console.log(data);
// } catch (e) {
// console.log(`ERROR: ${e}`);
// }
// }

// // main();
