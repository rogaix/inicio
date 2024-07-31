import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import AboutPage from '@/views/AboutPage.vue'
import WelcomePage from '@/views/WelcomePage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import PageNotFound from '@/views/PageNotFound.vue'

const routes = [
    { path: '/', component: WelcomePage },
    {
        path: '/register',
        component: RegisterPage,
        meta: { requiresGuest: true }
    }, {
        path: '/login',
        component: LoginPage,
        meta: { requiresGuest: true }
    }, {
        path: '/about',
        component: AboutPage,
        meta: { requiresAuth: true }
    },
    { path: '/:pathMatch(.*)*', component: PageNotFound },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const { isAuthenticated } = useAuth()
    if (to.meta.requiresAuth && !isAuthenticated.value) {
        next('/login')
    } else if (to.meta.requiresGuest && isAuthenticated.value) {
        next('/')
    } else {
        next()
    }
})

export default router
