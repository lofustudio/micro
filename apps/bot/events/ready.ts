import { DiscordEvent } from "../types/event";
import { bgBlue } from "colorette";

export const event: DiscordEvent<"ready"> = {
    name: "ready",
    run: async (client) => {
        console.log(bgBlue("[Bot]") + " Ready.");
    }
}