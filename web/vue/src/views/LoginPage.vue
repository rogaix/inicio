<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useForm } from '@/composables/useLoginForm.ts'

const { formData, errors, submitForm, loginError } = useForm()

</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <form
        @submit.prevent="submitForm"
        class="w-[600px] max-w-[80%] min-w-[300px]"
    >
      <div class="flex flex-col space-y-4">
        <h2 class="text-2xl font-semibold mb-4">
          Sign in to your account!
          <div class="text-base mt-1 font-normal">
            No account yet?
            <span class="underline text-bold">
                <a href="/register">Sign up!</a>
              </span>
          </div>
        </h2>

        <div
            v-if="loginError"
            class="text-red-700 m-2"
        >
          * {{ loginError }}
        </div>

        <div class="flex flex-col space-y-2">
          <span
              v-if="errors.email"
              class="text-red-700"
          >
            * {{ errors.email }}
          </span>
          <label for="email">
            Email Address
          </label>
          <input
              type="text"
              id="email"
              placeholder="E-Mail Address"
              v-model="formData.email"
              class="border-gray-300 border-1 p-5 rounded-lg shadow-md outline-none text-base focus:bg-white"
          >
        </div>

        <div class="flex flex-col space-y-2">
          <span
              v-if="errors.password"
              class="text-red-700"
          >
            * {{ errors.password }}
          </span>
          <label for="password">
            Password
          </label>
          <input
              type="password"
              id="password"
              v-model="formData.password"
              class="border-gray-300 border-1 p-5 rounded-lg shadow-md outline-none text-base"
          >
        </div>
        <button
            type="submit"
            class="rounded bg-blue-300 p-4 min-w-5"
        >
          Sign In
        </button>
      </div>
    </form>
  </div>
</template>

<style scoped>
h1 {
  color: #42b983;
}
</style>