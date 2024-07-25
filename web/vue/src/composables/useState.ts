import { computed } from 'vue'
import useApi from '@/composables/useApi'

const api = useApi()

export function useState() {
    const isLoggedIn = computed(() => {
        return api.hasToken()
    })

    return {
        isLoggedIn
    }
}