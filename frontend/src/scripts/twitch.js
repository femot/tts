import tmi from "tmi.js";

export default {
	client: null,
	init() {
		const client = new tmi.Client({
			channels: ["kitsuxiu"],
		});

		client.connect();

		this.client = client
	},
	onMessage(callback) {
		this.client.on("message", callback);
	}
}