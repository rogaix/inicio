import { ref } from 'vue'
import axios, { AxiosRequestConfig } from 'axios'

export default function useApi() {
    const token = ref(localStorage.getItem('token') || '')
    const error = ref(null)
    const baseURL = import.meta.env.VITE_API_URL
    const lastActivity = ref(Date.now())

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
        if (response.headers.authorization) {
            const newToken = response.headers.authorization.split(' ')[1]
            localStorage.setItem('token', newToken)
            token.value = newToken
        }
        return response
    }, (errorResponse) => {
        if (errorResponse.response && errorResponse.response.status === 401) {
            if (hasToken()) {
                clearToken()
                deleteToken()
            }
        }
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
    }

    const deleteToken = async () => {
        if (token.value) {
            try {
                await request({
                    method: 'post',
                    url: '/deleteSession'
                })
            } catch (error) {
                console.error("Failed to delete session from database:", error)
            }
        }
    }

    const hasToken = (): boolean => {
        const storedToken = localStorage.getItem('token')
        return Boolean(storedToken)
    }

    const checkSession = async () => {
        const now = Date.now()
        const inactivityPeriod = now - lastActivity.value

        if (!hasToken()) {
            return true
        }

        if (inactivityPeriod > 30 * 60 * 1000) { // 30 Minutes
            clearToken()
            await deleteToken()
            return false
        }

        try {
            const response = await request({
                method: 'post',
                url: '/updateSession'
            })
            return response.active
        } catch (errorResponse) {
            console.error("Check session error: ", errorResponse)
            clearToken()
            return false
        }
    }

    const updateLastActivity = () => {
        lastActivity.value = Date.now()
    }

    return {
        request,
        setToken,
        clearToken,
        error,
        hasToken,
        checkSession,
        updateLastActivity
    }
}