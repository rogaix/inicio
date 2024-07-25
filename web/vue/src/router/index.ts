import { createRouter, createWebHistory } from 'vue-router'
import AboutPage from '@/views/AboutPage.vue'
import WelcomePage from '@/views/WelcomePage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import PageNotFound from '@/views/PageNotFound.vue'
import { useState } from "@/composables/useState"

const routes = [
    { path: '/', component: WelcomePage },
    { path: '/register', component: RegisterPage },
    {
        path: '/login',
        component: LoginPage,
        beforeEnter: (to: any, from: any, next: any) => {
            const state = useState()
            console.log("state is ", state.isLoggedIn())
            if (state.isLoggedIn()) {
                console.log("state is true, redirecting to /")
                next({ path: '/' })
            } else {
                console.log("state is false, proceeding to /login")
                next()
            }
        },
    },
    { path: '/about', component: AboutPage },
    { path: '/:pathMatch(.*)*', component: PageNotFound },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
