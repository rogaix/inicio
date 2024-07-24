import { ref } from 'vue'
import axios from 'axios'

export function useForm() {
    const formData = ref({
        password: '',
        email: ''
    })

    const errors = ref<{ password: string | null, email: string | null }>({
        password: null,
        email: null
    })

    const validateForm = () => {
        let valid = true

        if (!formData.value.password) {
            errors.value.password = 'Password is required'
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
                const response = await axios.post(import.meta.env.VITE_API_URL + '/api/login', formData.value)
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
