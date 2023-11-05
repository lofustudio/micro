import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "eval",
    description: "Evaluate JavaScript code.",
    category: "core",
    run: async (client, message, args) => {
        if (message.author.id !== "944612635431821333") return;
        if (args[0] === "client.token") return message.reply("Nice try original gangster.");

        const code = args.join(" ");
        let evaled;

        try {
            evaled = await eval(code);
            message.channel.send(`\`\`\`js\n${evaled}\n\`\`\``);
        } catch (error) {
            message.channel.send(`\`\`\`js\n${error}\n\`\`\``);
        }
    }
}