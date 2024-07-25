import { createRouter, createWebHistory } from 'vue-router'
import AboutPage from '@/views/AboutPage.vue'
import WelcomePage from '@/views/WelcomePage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import PageNotFound from '@/views/PageNotFound.vue'

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
