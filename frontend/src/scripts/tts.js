import axios from "axios"
export default {
	audio: null,
	q: [],

	async say(message) {
		console.log("tts: " + message)

		this.q.push(message)
		console.log("added TTS to q. current queue length: " + this.q.length)

		if ((this.audio && this.audio.ended) || !this.audio) {
			this.play()
		}
	},
	async play() {
		const apiURL = "https://utils.idalon.com/v1/tts"

		if ((this.q && this.q.length < 1) || !this.q) {
			console.log("queue is empty")
			return
		}

		const message = this.q.shift()
		const resp = await axios.post(apiURL, {
			"voice": "Brian",
			"text": message
		})

		this.audio = new Audio(resp.data.speak_url)
		this.audio.onended = this.play.bind(this)
		this.audio.play()
	},
	skip() {
		if (this.audio) {
			this.audio.pause()
			this.play()
		}
	}
}
