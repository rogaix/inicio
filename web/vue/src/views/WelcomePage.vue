<script setup>
import LoginPage from "@/views/LoginPage.vue"
import RegisterPage from "@/views/RegisterPage.vue"
import { useState } from '@/composables/useState'
import { useForm } from '@/composables/useLogoutForm'
import { ref } from "vue"

const state = useState()
const { submit } = useForm()

const isLoggedIn = ref(state.isLoggedIn)
</script>

<template>
  <div
      v-if="isLoggedIn"
      class="m-5"
  >
    <a
        @click="submit"
        class="rounded-md bg-blue-300 m-4 p-2 cursor-pointer"
    >
      Logout
    </a>
    You're logged in
  </div>
  <div v-else="">
    <LoginPage v-if="true" />
    <RegisterPage v-else />
  </div>

</template>

<style scoped>
</style>
