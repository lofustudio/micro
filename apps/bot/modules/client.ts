import { Client, Collection } from "discord.js";
import { DiscordCommand } from "../types/command";
import { readdirSync } from "fs";
import path from "path";
import { DiscordEvent } from "../types/event";

export default class Cookie extends Client {
    public commands: Collection<string, DiscordCommand> = new Collection();
    public events: Collection<string, DiscordEvent<any>> = new Collection();

    public async start() {
        if (!process.env.TOKEN) throw new Error("No token provided.");

        const eventsPath = path.join(__dirname, "..", "events");
        const commandsPath = path.join(__dirname, "..", "commands");

        readdirSync(eventsPath).forEach((file) => {
            const { event } = require(`${eventsPath}/${file}`);
            console.log(`Loaded event ${event.name}`);
            this.events.set(event.name, event);
            this.on(event.name, event.run.bind(null, this));
        });

        readdirSync(commandsPath).forEach((file) => {
            const { command } = require(`${commandsPath}/${file}`);
            console.log(`Loaded command ${command.name}`);
            this.commands.set(command.name, command);
        });

        await this.login(process.env.TOKEN).then(async () => {
            console.log("Connected to Discord.");
        });
    }
};