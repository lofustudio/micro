import { Elysia } from "elysia";

const Web = {
    app: new Elysia(),
    async start() {
        this.app.get("/ping", () => "pong")
            .listen(3000);

        console.log(`Web has started at http://${this.app.server?.hostname}:${this.app.server?.port}.`);
    }
}

export default Web;