<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import useApi from '@/composables/useApi'

const { checkSession, updateLastActivity, updateSession } = useApi()
let sessionCheckInterval: NodeJS.Timeout | null = null
let sessionUpdateInterval: NodeJS.Timeout | null = null

const startSessionCheck = () => {
  sessionCheckInterval = setInterval(async () => {
    const isActive = await checkSession()
    if (!isActive) {
      window.location.href = '/login'
    }
  }, 60 * 1000)
}

const startSessionUpdate = () => {
  sessionUpdateInterval = setInterval(async () => {
    await updateSession()
  }, 5 * 60 * 1000)
}

const handleUserActivity = () => {
  updateLastActivity()
}

onMounted(() => {
  document.addEventListener('mousemove', handleUserActivity)
  document.addEventListener('keydown', handleUserActivity)
  startSessionCheck()
  startSessionUpdate()
})

onUnmounted(() => {
  if (sessionCheckInterval) {
    clearInterval(sessionCheckInterval)
  }
  if (sessionUpdateInterval) {
    clearInterval(sessionUpdateInterval)
  }
  document.removeEventListener('mousemove', handleUserActivity)
  document.removeEventListener('keydown', handleUserActivity)
})
</script>

<template>
  <router-view></router-view>
</template>
