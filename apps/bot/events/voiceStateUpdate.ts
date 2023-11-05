import { AudioPlayerStatus } from "@discordjs/voice";
import { DiscordEvent } from "../types/event";

export const event: DiscordEvent<"voiceStateUpdate"> = {
    name: "voiceStateUpdate",
    run: async (client, oldState, newState) => {
        client.audio.player.on("error", error => {
            console.error(`Error: ${error.message} with resource ${error.resource.metadata}`)
        });

        client.audio.player.on(AudioPlayerStatus.Playing, () => {

        });
    }
}