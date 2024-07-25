import { ref } from 'vue'
import useApi from '@/composables/useApi'

const api = useApi()
const isAuthenticated = ref(api.hasToken())

export function useAuth() {
    return { isAuthenticated }
}