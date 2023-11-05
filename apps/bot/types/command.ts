import { Message } from "discord.js";
import Cookie from "../modules/client";

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
    category: string;
    run: Run;
}