<script setup>
import { defineProps } from 'vue'

defineProps({
  msg: String
})
</script>

<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
  </div>
</template>

<style scoped>
h1 {
  font-weight: normal;
}
</style>
