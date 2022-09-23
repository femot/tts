<template>
  <button @click="$emit('click')" class="button">
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: 'Action',
}
</script>

<style scoped lang="sass">
@import @/assets/sass/base

.button
  padding: $sm $lg
  font-size: $text-sm
  color: $light
  border: 0
  background: $primary
  border-radius: 2px
  cursor: pointer
  outline: none

  &:hover
    background: $primary-tint-10
  &:active
    background: $primary-shade-10
</style>
