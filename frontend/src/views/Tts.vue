<template>
  <div class="home">
    <Card title="Hi">
      <div class="padding-md">
        <form class="channel" @submit.prevent="submit">
          <Field v-model="channel" label="Channel name" ></Field>

          <Action>Connect</Action>
        </form>
      </div>
    </Card>
  </div>
</template>


<script>
import Card from '@/components/Card.vue'
import Field from '@/components/Field.vue'
import Action from '@/components/Action.vue'
import twitch from '@/scripts/twitch'
import tts from '@/scripts/tts'

export default {
  name: 'Home',
  components: {
    Card,
    Field,
    Action,
  },
  data () {
    return {
      channel: ''
    }
  },
  mounted () {
    twitch.init()

    twitch.onMessage((_, context, message) => {
      console.log(context, message)

      const isTts = /^!tts .+/.test(message)
      const isSkipTts = /^!skiptts$/.test(message)
      const canSkipTts = Object.keys(context.badges).some(key => ['broadcaster', 'moderator'].includes(key))

      if (isTts) {
        return tts.say(message.replace('!tts ', ''));
      }

      if (isSkipTts && canSkipTts) {
        return tts.skip();
      }
    })
  },
  methods: {
    submit () {
      console.log('aa')
    }
  }
}
</script>

<style lang="sass" scoped>
@import @/assets/sass/base

.home
  height: 100vh
  display: flex
  align-items: center
  justify-content: center

.channel
  display: flex
  justify-content: space-between
  align-items: center
</style>
  