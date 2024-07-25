import { ref } from 'vue'
import useApi from './useApi'

export function useForm() {
    const { request, setToken, clearToken } = useApi()

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
                const response = await request({
                    method: 'post',
                    url: '/login',
                    data: formData.value
                });

                if(response.token) {
                    setToken(response.token)
                    console.log('Token saved')
                } else {
                    console.log('No token in response')
                }
                console.log(response)
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
