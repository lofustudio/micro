import { DiscordEvent } from "../types/event";

export const event: DiscordEvent<"ready"> = {
    name: "ready",
    run: async (client) => {
        console.log("Ready.");
    }
}