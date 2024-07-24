import { ref } from 'vue'
import axios from 'axios'

export function useForm() {
    const formData = ref({
        password: '',
        confirmPassword: '',
        email: ''
    })

    const errors = ref<{ password: string | null, confirmPassword: string | null, email: string | null }>({
        password: null,
        confirmPassword: null,
        email: null
    })

    const validateForm = () => {
        let valid = true

        if (!formData.value.confirmPassword) {
            errors.value.confirmPassword = 'Password is required'
            valid = false
        } else if (formData.value.confirmPassword !== formData.value.password) {
            errors.value.confirmPassword = 'Passwords do not match'
            valid = false
        }

        if (!formData.value.password) {
            errors.value.password = 'Password is required'
            valid = false
        } else if (formData.value.password.length < 8) {
            errors.value.password = 'Password must be at least 8 characters long'
            valid = false
        } else if (!/[A-Z]/.test(formData.value.password) ||
            !/[0-9]/.test(formData.value.password) ||
            !/[!@#$%^&*]/.test(formData.value.password)) {
            errors.value.password = 'Password must include at least one uppercase letter, one number and one special character'
            valid = false
        } else {
            errors.value.password = null
        }

        if (!formData.value.email) {
            errors.value.email = 'Email is required'
            valid = false
        } else if (!/\S+@\S+\.\S+/.test(formData.value.email)) {
            errors.value.email = 'Email is invalid'
            valid = false
        } else {
            errors.value.email = null
        }

        return valid
    }

    const submitForm = async () => {
        if (validateForm()) {
            try {
                console.log(import.meta.env.VITE_API_URL + '/register')
                const response = await axios.post(import.meta.env.VITE_API_URL + '/register', formData.value)
                console.log(response.data)
            } catch (error) {
                console.error('Error submitting form:', error)
            }
        }
    }

    return {
        formData,
        errors,
        submitForm
    }
}
