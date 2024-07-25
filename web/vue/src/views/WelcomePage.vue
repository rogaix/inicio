<script setup>
import LoginPage from "@/views/LoginPage.vue"
import RegisterPage from "@/views/RegisterPage.vue"
</script>

<template>
  <LoginPage v-if="true" />
  <RegisterPage v-else />

</template>

<style scoped>
</style>
