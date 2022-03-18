"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function sendMessage(client, channelName, message) {
    /* sends a message to a channel name
     * @param {Client} client the client
     * @param {string} channelName the channel name to send the message
     * @param {string} message the message to send
     * @return {boolean} if the message was successfully send
     */
    const channel = client.channels.cache.find((ch) => {
        // console.log("(anon)#(anon) ch.name: %s", ch.name); // __AUTO_GENERATED_PRINT_VAR__
        // @ts-ignore
        return ch.name == "development";
    });
    console.log("(anon)#(anon) channel: %s", channel); // __AUTO_GENERATED_PRINT_VAR__
    if (channel) {
        // @ts-ignore
        channel.send(message);
        return true;
    }
    return false;
}
