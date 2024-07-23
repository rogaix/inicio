<script setup>
import LoginPage from "@/components/LoginPage.vue"
import RegisterPage from "@/components/RegisterPage.vue"
</script>

<template>
  <LoginPage v-if="true" />
  <RegisterPage v-else />

</template>

<style scoped>
</style>
