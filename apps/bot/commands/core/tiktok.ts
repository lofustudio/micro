import { createAudioResource, joinVoiceChannel } from "@discordjs/voice";
import { DiscordCommand } from "../../types/command";
import axios from "axios";
import { Readable } from "stream";

function handleStatusError(status_code: any) {
    switch (status_code) {
        case 1:
            throw new Error(`Your TikTok session id might be invalid or expired. Try getting a new one. status_code: ${status_code}`);
        case 2:
            throw new Error(`The provided text is too long. status_code: ${status_code}`);
        case 4:
            throw new Error(`Invalid speaker, please check the list of valid speaker values. status_code: ${status_code}`);
        case 5:
            throw new Error(`No session id found. status_code: ${status_code}`);
    }
}

export const command: DiscordCommand = {
    name: "tiktok",
    category: "core",
    description: "Make VEGA speak with the tiktok voice.",
    run: async (client, message, args) => {
        if (!message.member?.voice.channel) return message.reply("You need to be in a voice channel to use this command.");
        if (!args[0]) return message.reply("You need to provide a sentence to speak.");

        let sentence = args.join(" ");

        sentence = sentence.replace("+", "plus");
        sentence = sentence.replace(/\s/g, '+');
        sentence = sentence.replace('&', 'and');

        let base = "https://api16-normal-v6.tiktokv.com/media/api/text/speech/invoke";
        let url = `${base}/?text_speaker=en_us_002&req_text=${sentence}&speaker_map_type=0&aid=1233`;
        const headers = {
            'User-Agent': 'com.zhiliaoapp.musically/2022600030 (Linux; U; Android 7.1.2; es_ES; SM-G988N; Build/NRD90M;tt-ok/3.12.13.1)',
            'Cookie': `sessionid="1e6150e05b0e7e26930e9d0d3a66d4a8"`,
            'Accept-Encoding': 'gzip,deflate,compress'
        };

        try {
            const result = await axios.post(url, null, { headers: headers });

            const status_code = result?.data?.status_code;

            if (status_code !== 0) return handleStatusError(status_code);
            const encoded_voice = result?.data?.data?.v_str;

            const buf = Buffer.from(encoded_voice, 'base64');
            const stream = Readable.from(buf);

            createAudioResource(stream);

            client.audio.player.play(createAudioResource(stream));

            if (!client.audio.connection) {
                client.audio.connection = joinVoiceChannel({
                    channelId: message.member?.voice.channel.id as string,
                    guildId: message.guild?.id as string,
                    adapterCreator: message.guild?.voiceAdapterCreator as any
                });

                client.audio.connection.subscribe(client.audio.player);
            };
        } catch (err) {
            throw new Error(`tiktok-tts ${err}`);
        }
    }
}