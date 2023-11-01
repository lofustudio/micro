import { DiscordEvent } from "../types/event";

export const event: DiscordEvent<"messageCreate"> = {
    name: "messageCreate",
    run: async (client, message) => {
        const prefix = "?";

        if (message.author.bot || !message.guild || !message.content.startsWith(prefix)) return;

        const args = message.content.slice(prefix.length).trim().split(/ +/g);
        const cmd = args.shift()?.toLowerCase();
        if (!cmd) return;

        const command = client.commands.get(cmd);

        if (command) {
            command.run(client, message, args);
        }
    },
}