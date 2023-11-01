import type { ClientEvents } from "discord.js";
import Cookie from "../modules/client";

export interface DiscordEvent<T extends keyof ClientEvents> {
    name: T;
    run: (client: Cookie, ...args: ClientEvents[T]) => void | PromiseLike<void>;
}