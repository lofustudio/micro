import { Client, Collection } from "discord.js";
import { DiscordCommand } from "./types/command";

class Cookie extends Client {
    public commands: Collection<string, DiscordCommand> = new Collection();

    public async start() {
        await this.login(process.env.TOKEN).then(() => {
            console.log("Connected to Discord.");
        });
    }
}

export default Cookie;