import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "help",
    description: "List all commands.",
    category: "core",
    run: async (client, message, args) => {
        const commands = client.commands.map(command => {
            return {
                name: command.name,
                description: command.description,
                category: command.category
            }
        });

        await message.reply(JSON.stringify(commands, null, 4));
    }
}