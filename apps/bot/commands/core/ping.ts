import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "ping",
    description: "Ping Pong!",
    category: "core",
    run: async (client, message, args) => {
        await message.reply("Pong!");
    }
}