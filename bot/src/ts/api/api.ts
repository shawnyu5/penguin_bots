import { createConnection, Schema, Types } from "mongoose";
require("dotenv").config();

interface IProduct {
   _id: string;
   title: string;
   average_price: number;
   average_discount: number;
}

const productSchema = new Schema({
   title: String,
   average_discount: Number,
   average_price: Number,
   appearances: Number,
});

export class Api {
   open_box: any = null;

   constructor() {
      this.open_box = null;
   }

   async init(connectionString: string) {
      return new Promise<void | string>((resolve, reject) => {
         // @ts-ignore
         const db = createConnection(connectionString, {
            // @ts-ignore
            useNewUrlParser: true,
            useUnifiedTopology: true,
         });

         db.once("error", (err) => {
            reject(err);
         });
         db.once("open", () => {
            this.open_box = db.model("Open_box", productSchema, "open_box");
            console.log("Connected to data base");
            resolve();
         });
      });
   }

   // return a product object by name exact name
   async findByName(searchTerm: string) {
      // return this.open_box.findOne(name).exec();
      return new Promise((resolve, reject) => {
         this.open_box.find(searchTerm, (err: any, data: any) => {
            if (err) {
               reject(err);
            }
            resolve(data);
         });
      });
   }

   // return a product object by regex matching
   async findNameByRegex(title: string | RegExp): Promise<Array<IProduct>> {
      return new Promise((resolve, reject) => {
         // convert title to case insenitive regular expression
         title = new RegExp(title, "i");
         this.open_box.find(
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

async function main() {
   let api = new Api();
   try {
      // @ts-ignore
      await api.init(process.env.key);
      let obj = {
         _id: new Types.ObjectId("61dceb6228b23db27260d4e0"),
         title: "Play Money by Nick Diffatte (Instant Download)",
         average_discount: 33.333333333333336,
         average_price: 3.3000000000000003,
         appearances: 3,
      };

      let data = await api.findNameByRegex("nick diffatte");
      console.log(data);
   } catch (e) {
      console.log(`ERROR: ${e}`);
   }
}

// main();
