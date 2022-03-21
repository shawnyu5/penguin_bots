"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildMessage = exports.getChannelByName = exports.checkCoinProduct = void 0;
// const exec = require("child_process").execSync;
const child_process_1 = require("child_process");
const config_json_1 = __importDefault(require("../config.json"));
/*
 * @param {Client} client the client
 * @param {string} channelName the channel name to send the message
 * @param {string} message the message to send
 * @return {boolean} if the message was successfully send
 */
/**
 * @param client - discord client
 * @param channelName - name of channel to search for
 * @returns channel information with the name passed in. If not found. undefined
 */
function getChannelByName(client, channelName) {
    const channel = client.channels.cache.find((ch) => {
        // @ts-ignore
        return ch.name == channelName;
    });
    return channel;
}
exports.getChannelByName = getChannelByName;
/**
 * @returns return json message from python script. Other wise return null
 */
function checkCoinProduct() {
    let result = (0, child_process_1.execSync)("python3 ../coin_products/coin_products.py");
    return JSON.parse(String(result));
}
exports.checkCoinProduct = checkCoinProduct;
/**
 * @param coinProduct - the coin product
 * @returns a message pining all users in config.json about the coinProduct
 */
function buildMessage(coinProduct) {
    let message = "";
    config_json_1.default.coin_product_alert_users.forEach((user) => {
        message += `<@${user}>
      `;
    });
    message += `title: ${coinProduct.title}
   url: ${coinProduct.url}`;
    return message;
}
exports.buildMessage = buildMessage;
