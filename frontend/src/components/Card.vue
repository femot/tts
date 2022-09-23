<template>
  <div class="card">
    <div class="card-header">{{ title }}</div>

    <div class="card-body">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Card',
  props: {
    title: String
  }
}
</script>

<style lang="sass" scoped>
@import @/assets/sass/base

.card
  display: flex
  flex-direction: column
  width: 400px
  border-radius: 2px
  overflow: hidden
  box-shadow: rgb(0 0 0 / 20%) 0px 3px 3px -2px, rgb(0 0 0 / 14%) 0px 3px 4px 0px, rgb(0 0 0 / 12%) 0px 1px 8px 0px

  .card-body
    background: $secondary-shade-50
    flex: 1

  .card-header
    border-top: none
    border-bottom: none
    display: flex
    background: $primary-shade-20
    justify-content: space-between
    align-items: center
    padding: $lg
  
</style>
