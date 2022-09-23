<template>
  <div class="home">
    <Card title="Connect your channel">
      <div class="padding-md">
        <form class="channel mb-lg" @submit.prevent="submit">
          <Field v-model="channel" label="Channel name"></Field>

          <Action>Connect</Action>
        </form>
        
        <p class="help text-grey-50">By clicking connect you agree that you voluntarily subject yourself to mental torture via any TTS spam. The creators of this tool do not take any responsibilty for potential emotional damage caused.</p>
      </div>
    </Card>
  </div>
</template>


<script>
import Card from '@/components/Card.vue'
import Field from '@/components/Field.vue'
import Action from '@/components/Action.vue'

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
  methods: {
    submit () {
      if (!this.channel) {
        return
      }

      this.$router.push(`/${this.channel}`)
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

.help
  font-size: $text-xs
  text-align: justify
</style>
  