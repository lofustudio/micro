import { GatewayIntentBits } from "discord.js";

import Cookie from "./src";
import Web from "./web";

async function _() {
    let bot = new Cookie({
        intents: [
            GatewayIntentBits.Guilds,
            GatewayIntentBits.GuildMessages,
            GatewayIntentBits.GuildMembers,
            GatewayIntentBits.GuildModeration,
            GatewayIntentBits.MessageContent
        ]
    });

    await bot.start().then(async () => {
        await Web.start();
    });
}

_();