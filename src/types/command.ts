import { Message } from "discord.js";
import Cookie from "..";

interface Run {
    (
        client: Cookie,
        message: Message,
        args: string[]
    ): unknown;
}

export interface DiscordCommand {
    name: string;
    description: string;
    run: Run;
}