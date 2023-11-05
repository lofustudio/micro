import 'dotenv/config';

import Cookie from "./modules/client";
import { GatewayIntentBits, Partials } from 'discord.js';

new Cookie({
    intents: [
        GatewayIntentBits.Guilds,
        GatewayIntentBits.GuildMessages,
        GatewayIntentBits.GuildVoiceStates,
        GatewayIntentBits.GuildModeration,
        GatewayIntentBits.MessageContent
    ],
    partials: [
        Partials.Message,
        Partials.Channel,
    ]
}).start();