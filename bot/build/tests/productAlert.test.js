"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const productAlert_1 = require("../commands/productAlert");
describe("Product alert function testing", () => {
    it("Should add a user to config.json", () => {
        const config = require("../../config.json");
        let user = {
            id: "d",
            username: "d",
            accentColor: 10,
        };
        let updatedConfig = (0, productAlert_1.addUser)(user);
        expect(updatedConfig.coin_product_alert_users).toContain("d");
    });
    it("should not add repeat users to config.json", () => {
        let user = {
            id: "d",
            username: "d",
        };
        let updatedConfig = (0, productAlert_1.addUser)(user);
        updatedConfig = (0, productAlert_1.addUser)(user);
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
        let user = {
            id: "d",
            username: "d",
            accentColor: 10,
        };
        const updatedConfig = (0, productAlert_1.deleteUser)(user);
        expect(updatedConfig.coin_product_alert_users).not.toContain("d");
    });
});
