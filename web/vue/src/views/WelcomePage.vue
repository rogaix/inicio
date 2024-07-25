<script setup>
import LoginPage from "@/views/LoginPage.vue"
import RegisterPage from "@/views/RegisterPage.vue"
import { useState } from '@/composables/useState'

const state = useState()
</script>

<template>
  <div v-if="state.isLoggedIn">
    You're logged in
  </div>
  <div v-else="">
    <LoginPage v-if="true" />
    <RegisterPage v-else />
  </div>

</template>

<style scoped>
</style>
