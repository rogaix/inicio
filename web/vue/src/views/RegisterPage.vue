<script setup>
import { useForm } from '@/composables/useRegistrationForm.ts'

const { formData, errors, submitForm } = useForm()
</script>

<template>
  <div class="flex justify-center items-center h-screen bg-gray-100">
    <form @submit.prevent="submitForm">
      <div class="flex flex-col w-[600px] max-w-[80%] space-y-4">
        <h2 class="text-2xl font-semibold mb-4">
          Sign up an account!
          <div class="text-base mt-1 font-normal">
            Already have an account?
            <span class="underline text-bold">
                <a href="/login">Sign In</a>
              </span>
          </div>
        </h2>

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
              name="email"
              id="email"
              placeholder="E-Mail Address"
              v-model="formData.email"
              class="border-gray-300 border-1 p-5 rounded-lg shadow-md outline-none text-base"
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

        <div class="flex flex-col space-y-2">
          <span
              v-if="errors.confirmPassword"
              class="text-red-700"
          >
            * {{ errors.confirmPassword }}
          </span>
          <label for="confirm-password">
            Confirm Password
          </label>
          <input
              type="password"
              id="confirm-password"
              v-model="formData.confirmPassword"
              class="border-gray-300 border-1 p-5 rounded-lg shadow-md outline-none text-base"
          >
        </div>
        <button
            type="submit"
            class="rounded bg-blue-300 p-4 min-w-5"
        >
          Sign Up
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