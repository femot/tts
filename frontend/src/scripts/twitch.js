import tmi from "tmi.js";

export default {
	client: null,
	init(channel) {
		const client = new tmi.Client({
			channels: [channel],
		});

		client.connect();

		this.client = client
	},
	onMessage(callback) {
		this.client.on("message", callback);
	}
}