import { joinVoiceChannel } from "@discordjs/voice";
import { DiscordCommand } from "../../types/command";

export const command: DiscordCommand = {
    name: "join",
    category: "music",
    description: "Joins the voice channel you are in.",
    run: async (client, message, args) => {
        const channel = message.member?.voice.channel;
        if (!channel) return message.reply("You must be in a voice channel to use this command.");

        if (client.audio.connection) return message.reply("I am already in a voice channel.");

        client.audio.channel = channel;
        client.audio.connection = joinVoiceChannel({
            channelId: channel.id,
            guildId: channel.guild.id,
            adapterCreator: channel.guild.voiceAdapterCreator
        });

        message.reply(`Joined ${channel.name}.`);
    }
}