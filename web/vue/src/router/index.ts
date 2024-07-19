import { createRouter, createWebHistory } from 'vue-router'
import AboutPage from '@/components/AboutPage.vue'
import WelcomePage from '@/components/WelcomePage.vue'
import PageNotFound from '@/components/PageNotFound.vue'

const routes = [
    { path: '/', component: WelcomePage },
    { path: '/about', component: AboutPage },
    { path: '/:pathMatch(.*)*', component: PageNotFound },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
