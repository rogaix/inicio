import { ref } from 'vue'
import axios, { AxiosRequestConfig } from 'axios'

export default function useApi() {
    const token = ref(localStorage.getItem('token') || '')
    const error = ref(null)

    const baseURL = import.meta.env.VITE_API_URL

    const instance = axios.create({
        baseURL
    })

    instance.interceptors.request.use((config) => {
        if (token.value) {
            config.headers.Authorization = `Bearer ${token.value}`
        }
        return config
    })

    instance.interceptors.response.use((response) => {
        if (response.data.token) {
            localStorage.setItem('token', response.data.token)
            token.value = response.data.token
        }
        return response
    }, (errorResponse) => {
        // if (errorResponse.response && errorResponse.response.status === 401) {
        //     refreshAuthToken()
        // }
        error.value = errorResponse
        return Promise.reject(errorResponse)
    })

    const request = async (options: AxiosRequestConfig<any>) => {
        try {
            error.value = null
            const response = await instance(options)
            return response.data
        } catch (err) {
            throw error.value
        }
    }

    const setToken = (newToken: string) => {
        localStorage.setItem('token', newToken)
        token.value = newToken
    }

    const clearToken = () => {
        localStorage.removeItem('token')
        token.value = ''
    };

    const hasToken = (): boolean => {
        const storedToken = localStorage.getItem('token')
        return Boolean(storedToken)
    }

    const refreshAuthToken = async () => {
        try {
            const response = await instance.post('/refreshToken')
            localStorage.setItem('token', response.data.token)
            token.value = response.data.token
            return response
        } catch (errorResponse) {
            console.error("Refresh auth token error: ", errorResponse)
            // @ts-ignore
            error.value = errorResponse
            return Promise.reject(errorResponse)
        }
    }

    return {
        request,
        setToken,
        clearToken,
        error,
        hasToken
    }
}