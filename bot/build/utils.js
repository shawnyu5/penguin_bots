"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.sendMessage = exports.checkCoinProduct = void 0;
// const exec = require("child_process").execSync;
const child_process_1 = require("child_process");
/*
 * @param {Client} client the client
 * @param {string} channelName the channel name to send the message
 * @param {string} message the message to send
 * @return {boolean} if the message was successfully send
 */
/**
 * @param client - the client
 * @param channelName - name of the channel to send the message
 * @param message - the message to be sent
 * @returns if message was send successfully
 */
function sendMessage(client, channelName, message) {
    const channel = client.channels.cache.find((ch) => {
        // @ts-ignore
        return ch.name == channelName;
    });
    if (channel) {
        // @ts-ignore
        channel.send(message);
        return true;
    }
    return false;
}
exports.sendMessage = sendMessage;
/**
 * @returns check if the current product from python script is a coin product. Other wise return null
 */
function checkCoinProduct() {
    let output = null;
    let result = (0, child_process_1.execSync)("python3 ../coin_products/coin_products.py");
    console.log("checkCoinProduct result: %s", result); // __AUTO_GENERATED_PRINT_VAR__
    return output;
}
exports.checkCoinProduct = checkCoinProduct;
