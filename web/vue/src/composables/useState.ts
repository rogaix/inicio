import useApi from '@/composables/useApi'

const api = useApi()


export function useState() {
    const isLoggedIn = (): boolean => {
        return api.hasToken()
    }

    return {
        isLoggedIn
    }
}