import { AudioPlayer, AudioResource, VoiceConnection } from "@discordjs/voice";
import { Client, Collection, VoiceBasedChannel, VoiceChannel } from "discord.js";
import { DiscordCommand } from "../types/command";
import { DiscordEvent } from "../types/event";
import { readdirSync } from "fs";
import path from "path";

export default class VEGA extends Client {
    public commands: Collection<string, DiscordCommand> = new Collection();
    public events: Collection<string, DiscordEvent<any>> = new Collection();

    public audio: {
        player: AudioPlayer,
        connection: VoiceConnection | null,
        channel: VoiceBasedChannel | VoiceChannel | null
    } = {
            player: new AudioPlayer(),
            connection: null,
            channel: null
        };

    public tts: {
        channel: null | string,
        connection: null | VoiceConnection,
        lastUser: null,
        queue: Array<AudioResource<{ title: string; }>>
    } = {
            channel: null,
            connection: null,
            lastUser: null,
            queue: []
        }

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

        readdirSync(commandsPath).forEach((dir) => {
            const commandsList = readdirSync(`${commandsPath}/${dir}`).filter((file) => file.endsWith(".ts") || file.endsWith(".js"));

            for (const file of commandsList) {
                const { command } = require(`${commandsPath}/${dir}/${file}`);
                console.log(`Loaded command ${command.name}`);
                this.commands.set(command.name, command);
            }
        });

        await this.login(process.env.TOKEN).then(async () => {
            console.log("Connected to Discord.");
        });
    }
};