import { createAudioResource, joinVoiceChannel } from "@discordjs/voice";
import { DiscordCommand } from "../../types/command";
import VEGA from "../../modules/client";
import { Message } from "discord.js";
import * as tts from "google-tts-api";

export function ttsHandle(client: VEGA, message: Message, args: Array<string>) {
    if (message.channel.id !== client.tts.channel || message.author.bot || !message.member?.voice.channel) return;
    if (!message.content) return message.reply("TTS only works with text atm.");

    let sentence = message.content;

    const url = tts.getAudioUrl(sentence, {
        lang: "en-US",
        slow: false,
        host: 'https://translate.google.com'
    });


    if (!client.audio.connection) {
        client.audio.connection = joinVoiceChannel({
            channelId: message.member?.voice.channel.id as string,
            guildId: message.guild?.id as string,
            adapterCreator: message.guild?.voiceAdapterCreator as any
        });
    }

    const resource = createAudioResource(url, {
        metadata: {
            title: "tts"
        }
    });

    if (client.tts.queue.length > 0) {
        client.tts.queue.push(resource);
    } else {
        client.audio.player.play(resource);
    }
}

export const command: DiscordCommand = {
    name: "tts",
    category: "core",
    description: "TTS",
    run: async (client, message, args) => {
        if (!message.member?.voice.channel) return message.reply("You need to be in a voice channel to use this command.");

        client.tts.channel = message.channel.id;

        if (!client.tts.connection) {
            client.tts.connection = joinVoiceChannel({
                channelId: message.member?.voice.channel.id as string,
                guildId: message.guild?.id as string,
                adapterCreator: message.guild?.voiceAdapterCreator as any
            });
        }

        client.tts.connection.subscribe(client.audio.player);

        message.reply(`TTS enabled. Listening to the voice channel: <#${message.channel.id}>.`);
    }
}