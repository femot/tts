<template>
  <p class="field">
    <span v-if="label" class="label">{{ label }}</span>

    <input
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      maxlength="100"
      :placeholder="placeholder"
    />
  </p>
</template>

<script>
export default {
  name: 'Field',
  emits: ['update:modelValue'],
  props: {
    label: String,
    placeholder: String,
    modelValue: String
  }
}
</script>

<style scoped lang="sass">
@import @/assets/sass/base

.field
  margin: 0
  padding: 0
  display: flex
  width: 100%
  margin-right: $lg

  .suffix
    display: block
    padding: $sm
    font-size: $text-sm
    color: $grey-30
    border-bottom: 2px solid $secondary-tint-10

  .label
    display: block
    padding: $sm
    padding-left: 0
    font-size: $text-sm
    color: $grey-30
    border-bottom: 2px solid transparent

  input
    width: 100%
    flex: 1
    border: none
    background: none
    padding: $sm $md
    font-size: $text-sm
    font-family: $font-family
    color: $light
    outline: none
    border-bottom: 2px solid $secondary-tint-10

    &:hover
      &, + .suffix
        border-bottom: 2px solid $secondary-tint-20

    &:active, &:focus
      &, + .suffix
        border-bottom: 2px solid $primary
</style>
