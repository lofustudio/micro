import { createAudioResource, joinVoiceChannel } from "@discordjs/voice";
import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "stream",
    description: "Play a OPUS stream.",
    category: "music",
    run: async (client, message, args) => {
        if (!message.member?.voice.channel) return message.reply("You need to be in a voice channel to use this command.");
        if (!args[0]) return message.reply("You need to provide a stream URL to play.");

        if (!client.audio.connection) {
            client.audio.connection = joinVoiceChannel({
                channelId: message.member.voice.channel.id,
                guildId: message.guild?.id as string,
                adapterCreator: message.guild?.voiceAdapterCreator as any
            });
        }

        const resource = createAudioResource(args[0], {
            metadata: {
                title: args
            }
        });

        client.audio.player.play(resource);

        client.audio.connection.subscribe(client.audio.player);

        message.reply("Playing stream.");
    }
}