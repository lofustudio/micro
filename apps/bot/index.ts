import 'dotenv/config';
import Cookie from "./modules/client";
import { GatewayIntentBits, Partials } from 'discord.js';

new Cookie({
    intents: [
        GatewayIntentBits.Guilds,
        GatewayIntentBits.GuildMessages,
        GatewayIntentBits.MessageContent
    ],
    partials: [
        Partials.Message
    ]
}).start();