import {createWebHistory,createRouter} from  'vue-router'
/* BaseLayout */
import BaseLayout from '@/layout/BaseLayout.vue'
// 公共路由
export const constantRoutes = [
    {
        path: '',
        component: BaseLayout,
        redirect: 'index',
        children: [
            {
                path: 'index',
                component: () => import('@/views/Index.vue'),
                name: 'Index',
                meta: { title: '首页',  affix: true }
            },
            {
                path: 'about',
                component: () => import('@/views/about/Index.vue'),
                name: 'about',
                meta: { title: '关于我',  affix: true }
            },
            {
                path: 'blog',
                component: () => import('@/views/blog/Index.vue'),
                name: 'blog',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'blog/detail/:id(\\d+)',
                component: () => import('@/views/blog/Detail.vue'),
                name: 'blog-detail',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'blog/edit/:id(\\d+)',
                component: () => import('@/views/blog/Edit.vue'),
                name: 'blog-edit',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'user',
                component: () => import('@/views/user/Index.vue'),
                name: 'user',
                meta: { title: '我',  affix: true }
            },
            {
                path: 'travel',
                component: () => import('@/views/travel/Index.vue'),
                name: 'travel',
                meta: { title: '旅行',  affix: true }
            },
        ]
    },
]

const router = createRouter({
    history:createWebHistory(),
    routes:constantRoutes
})

export default router
