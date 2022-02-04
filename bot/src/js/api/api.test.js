"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const mongoose_1 = require("mongoose");
const { Api } = require("./api");
describe("Find by name API", () => {
    test("Return name by exact match", async () => {
        try {
            let api = new Api();
            await api.init(process.env.key);
            let product = await api.findByName({
                title: "Play Money by Nick Diffatte (Instant Download)",
            });
            let obj = {
                _id: new mongoose_1.Types.ObjectId("61dceb6228b23db27260d4e0"),
                title: "Play Money by Nick Diffatte (Instant Download)",
                average_discount: 33.333333333333336,
                average_price: 3.3000000000000003,
                appearances: 3,
            };
            expect(product).toBe(obj);
        }
        catch (e) {
            console.log(`ERROR: ${e}`);
        }
    });
    test("return name by partial match", async () => {
        let api = new Api();
        try {
            await api.init(process.env.key);
            let obj = {
                _id: new mongoose_1.Types.ObjectId("61dceb6228b23db27260d4e0"),
                title: "Play Money by Nick Diffatte (Instant Download)",
                average_discount: 33.333333333333336,
                average_price: 3.3000000000000003,
                appearances: 3,
            };
            let data = await api.findNameByRegex({
                title: "Play Money by Nick Diffatte",
            });
            // expect(data).toBe(obj);
        }
        catch (e) {
            console.log(`ERROR: ${e}`);
        }
    });
});
