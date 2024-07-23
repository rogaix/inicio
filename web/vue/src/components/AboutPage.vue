<script setup lang="ts">
import { ref, onMounted } from 'vue'

const message = ref('')
const user = ref()

onMounted(async () => {
  try {
    const response = await fetch('/api/data')
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const data = await response.json()
    message.value = data.message
    user.value = data.user
    console.log("User data:", user.value)
  } catch (error) {
    console.error('Error fetching data:', error)
  }
})
</script>

<template>
  <div>
    <h1>About Page</h1>
    <p>This is the about page.</p>
    <p>{{ message }}</p>
    <div v-if="user">
      <p>{{ user.name }} ID: {{ user.id }}</p>
    </div>
    <div v-else>
      <p>Loading user data...</p>
    </div>
  </div>
</template>

<style scoped>
h1 {
  color: #42b983;
}
p {
  color: #666;
}
</style>
