<script setup lang="ts">
import { ref, onMounted } from 'vue';

const message = ref('');

onMounted(async () => {
  try {
    const response = await fetch('/api/data');
    const data = await response.json();
    message.value = data.message;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
});
</script>

<template>
  <div>
    <h1>About Page</h1>
    <p>This is the about page.</p>
    <p>{{ message }}</p>
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
