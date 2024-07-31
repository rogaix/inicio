<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import useApi from '@/composables/useApi'

const { checkSession, updateLastActivity } = useApi()
let sessionCheckInterval: NodeJS.Timeout | null = null

const startSessionCheck = () => {
  sessionCheckInterval = setInterval(async () => {
    const isActive = await checkSession()
    if (!isActive) {
      // alert('Your session has expired. Please log in again.')
      window.location.href = '/login'
    }
  }, 60000 * 5)
}

const handleUserActivity = () => {
  updateLastActivity()
}

onMounted(() => {
  document.addEventListener('mousemove', handleUserActivity)
  document.addEventListener('keydown', handleUserActivity)
  startSessionCheck()
})

onUnmounted(() => {
  if (sessionCheckInterval) {
    clearInterval(sessionCheckInterval)
  }
  document.removeEventListener('mousemove', handleUserActivity)
  document.removeEventListener('keydown', handleUserActivity)
})
</script>

<template>
  <router-view></router-view>
</template>
