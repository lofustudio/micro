import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "stop",
    category: "music",
    description: "Stop playing music.",
    run: async (client, message, args) => {
        client.audio.player.stop();
        message.reply("Stopped playing stream.");
    }
}