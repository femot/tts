<template>
  <div class="tts">
    <Card :title="`Connected to ${$route.params.channel}`">
      <div class="padding-md">
        <Action v-if="!active" @click="active = true">Activate</Action>

        <p v-else-if="!queue || queue.length === 0" class="text-grey-30">The TTS queue is currently empty, awaiting requests</p>

        <div v-else>
          Requested messages

          <ul class="list">
            <li v-for="(message, key) of queue.reverse()" :key="key">
              {{ message }}
            </li>
          </ul>
        </div>
      </div>
    </Card>
  </div>
</template>


<script>
import Card from '@/components/Card.vue'
import Action from '@/components/Action.vue'
import twitch from '@/scripts/twitch'
import tts from '@/scripts/tts'

export default {
  name: 'Tts',
  components: {
    Card,
    Action,
  },
  data () {
    return {
      active: false,
      queue: [],
    }
  },
  mounted () {
    twitch.init(this.$route.params.channel)

    twitch.onMessage((_, context, message) => {
      const isTts = /^!tts .+/.test(message)
      const isSkipTts = /^!skiptts$/.test(message)
      const canSkipTts = Object.keys(context.badges).some(key => ['broadcaster', 'moderator'].includes(key))

      if (isTts) {
        const text= message.replace('!tts ', '')

        this.queue.push(text)

        return tts.say(text);
      }

      if (isSkipTts && canSkipTts) {
        return tts.skip();
      }
    })
  }
}
</script>

<style lang="sass" scoped>
@import @/assets/sass/base

.tts
  height: 100vh
  display: flex
  align-items: center
  justify-content: center

.list
  margin-top: 16px
</style>
  