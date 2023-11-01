import { DiscordCommand } from "../types/command";

export const command: DiscordCommand = {
    name: "ping",
    description: "Ping Pong!",
    run: async (client, message, args) => {
        await message.reply("Pong!");
    }
}