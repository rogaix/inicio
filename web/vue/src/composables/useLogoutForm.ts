import { ref } from 'vue'
import useApi from './useApi'

export function useForm() {
    const { request, clearToken } = useApi()

    const formData = ref({
        password: '',
        email: ''
    })

    const errors = ref<{ password: string | null, email: string | null }>({
        password: null,
        email: null
    })

    const submit = async () => {
        try {
            const response = await request({
                method: 'post',
                url: '/logout',
                data: formData.value
            })

            clearToken()
        } catch (error) {
            console.error('Error submitting form:', error)
        }
    }

    return {
        formData,
        errors,
        submit
    }
}
