import { User } from "discord.js";
import { addUser, deleteUser } from "../commands/productAlert";

describe("Product alert function testing", () => {
   it("Should add a user to config.json", () => {
      const config = require("../../config.json");
      let user: User = {
         id: "d",
         username: "d",
         accentColor: 10,
      } as User;

      let updatedConfig = addUser(user);
      expect(updatedConfig.coin_product_alert_users).toContain("d");
   });

   it("should not add repeat users to config.json", () => {
      let user: User = {
         id: "d",
         username: "d",
      } as User;
      let updatedConfig = addUser(user);
      updatedConfig = addUser(user);

      let count = 0;
      // count number of times user with id d is in config.json
      updatedConfig.coin_product_alert_users.forEach((user) => {
         if (user === "d") {
            count++;
         }
      });

      expect(count).toBe(1); // should only be one user with id d
   });

   it("should delete a user", () => {
      let user: User = {
         id: "d",
         username: "d",
         accentColor: 10,
      } as User;
      const updatedConfig = deleteUser(user);
      expect(updatedConfig.coin_product_alert_users).not.toContain("d");
   });
});
