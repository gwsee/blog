import {createWebHistory,createRouter} from  'vue-router'
/* BaseLayout */
import BaseLayout from '@/layout/BaseLayout.vue'
// 公共路由
export const constantRoutes = [
    {
        path: '',
        component: BaseLayout,
        redirect: '/index',
        children: [
            {
                path: 'index',
                component: () => import('@/views/index.vue'),
                name: 'index',
                meta: { title: '首页',  affix: true }
            },
            {
                path: 'about',
                component: () => import('@/views/about/index.vue'),
                name: 'about',
                meta: { title: '关于我',  affix: true }
            },
            {
                path: 'demo',
                component: () => import('@/views/demo/index.vue'),
                name: 'demo',
                meta: { title: 'demo',  affix: true }
            },
            {
                path: 'about/experience/:id(\\d+)',
                component: () => import('@/views/about/experience.vue'),
                name: 'about-experience',
                meta: { title: '工作经历管理',  affix: true }
            },
            {
                path: 'about/experience/manage/:id(\\d+)',
                component: () => import('@/views/about/manage/experience.vue'),
                name: 'about-experience-manage',
                meta: { title: '工作经历管理',  affix: true }
            },
            {
                path: 'about/profile/manage',
                component: () => import('@/views/about/manage/profile.vue'),
                name: 'about-profile-manage',
                meta: { title: '个人简介管理',  affix: true }
            },
            {
                path: 'about/project/manage/:id(\\d+)',
                component: () => import('@/views/about/manage/project.vue'),
                name: 'about-project-manage',
                meta: { title: '工作项目管理',  affix: true }
            },
            {
                path: 'blog',
                component: () => import('@/views/blog/index.vue'),
                name: 'blog',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'blog/detail/:id(\\d+)',
                component: () => import('@/views/blog/detail.vue'),
                name: 'blog-detail',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'blog/manage/:id(\\d+)',
                component: () => import('@/views/blog/manage.vue'),
                name: 'blog-edit',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'user',
                component: () => import('@/views/user/index.vue'),
                name: 'user',
                meta: { title: '我',  affix: true }
            },
            {
                path: 'travel',
                component: () => import('@/views/travel/index.vue'),
                name: 'travel',
                meta: { title: '旅行',  affix: true }
            },
            {
                path: 'travel/manage/:id(\\d+)',
                component: () => import('@/views/travel/manage.vue'),
                name: 'travel-manage',
                meta: { title: '旅行管理',  affix: true }
            },
            {
                path: 'travel/info/:id(\\d+)',
                component: () => import('@/views/travel/info.vue'),
                name: 'travel-info',
                meta: { title: '日记',  affix: true }
            },
            {
                path: 'picture',
                component: () => import('@/views/picture/index.vue'),
                name: 'picture',
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
