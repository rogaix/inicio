import { createRouter, createWebHistory } from 'vue-router'
import AboutPage from '@/components/AboutPage.vue'
import WelcomePage from '@/components/WelcomePage.vue'
import RegisterPage from '@/components/RegisterPage.vue'
import LoginPage from '@/components/LoginPage.vue'
import PageNotFound from '@/components/PageNotFound.vue'

const routes = [
    { path: '/', component: WelcomePage },
    { path: '/register', component: RegisterPage },
    { path: '/login', component: LoginPage },
    { path: '/about', component: AboutPage },
    { path: '/:pathMatch(.*)*', component: PageNotFound },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
