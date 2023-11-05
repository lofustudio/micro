import { Message } from "discord.js";
import { ttsHandle } from "../commands/core/tts";
import VEGA from "../modules/client";
import { DiscordEvent } from "../types/event";

export const event: DiscordEvent<"messageCreate"> = {
    name: "messageCreate",
    run: async (client: VEGA, message: Message<boolean>): Promise<void> => {
        const prefix = process.env.PREFIX ?? ">";
        const args = message.content.slice(prefix.length).trim().split(/ +/g);

        if (!message.content.startsWith(prefix)) {
            ttsHandle(client, message, args);
            return;
        } else {
            if (message.author.bot || !message.guild) return;

            const cmd = args.shift()?.toLowerCase();
            if (!cmd) return;

            const command = client.commands.get(cmd);
            if (command) {
                await command.run(client, message, args);
            }
        }
    },
};