<script setup>
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <div class="flex flex-col w-80 space-y-4">
      <h2 class="text-2xl font-semibold">
        Sign in to your account!
        <div class="text-base mt-1 font-normal">
          No account yet?
          <span class="underline text-bold">
              <a href="/register">Sign up!</a>
            </span>
        </div>
      </h2>

      <div class="flex flex-col space-y-2">
        <label for="email">
          Email Address
        </label>
        <input
            type="text"
            name="email"
            id="email"
            placeholder="E-Mail Address"
            class="border-gray-300 border-1 p-3 rounded-md shadow-md outline-none text-sm"
        >
      </div>

      <div class="flex flex-col space-y-2">
        <label for="password">
          Password
        </label>
        <input
            type="password"
            name="password"
            id="password"
            class="border-gray-300 border-1 p-3 rounded-md shadow-md outline-none text-sm"
        >
      </div>
    </div>
  </div>
</template>

<style scoped>
h1 {
  color: #42b983;
}
</style>