<template>
  <div class="w-full layout">
    <!-- 主导航栏 -->
    <div class="sticky top-0 z-50 w-full layout-header bg-white/5 backdrop-blur-[2px]hover:bg-white/20 transition-colors duration-300">
      <div class="mx-auto max-w-[1920px] px-4">
        <div class="flex h-16 items-center justify-between">
          <!-- 左侧区域：移动端菜单按钮和Logo -->
          <div class="flex items-center gap-4">
            <!-- 移动端菜单按钮 -->
            <a-button
                type="text"
                class="md:hidden -ml-2"
                @click="showMobileMenu = true"
            >
              <template #icon><MenuOutlined /></template>
            </a-button>

            <!-- Logo -->
            <router-link to="/" class="flex items-center -ml-3 md:ml-0">
              <img src="@/assets/image/logo.png" alt="Logo" class="h-20 w-auto" />
            </router-link>
          </div>

          <!-- 中间区域：搜索框 -->
          <div class="hidden md:flex flex-1 justify-center max-w-3xl mx-8" v-if="false">
            <div class="relative w-full group">
              <input
                  type="search"
                  placeholder="搜索..."
                  class="w-full rounded-full border-0 pl-4 pr-24 py-2.5 bg-[#98FB98]/10 hover:bg-[#98FB98]/20 focus:bg-[#98FB98]/30 transition-colors duration-300 focus:outline-none focus:ring-1 focus:ring-green-200/30"
              />
              <div class="absolute right-0 top-0 bottom-0 flex items-center opacity-60 group-hover:opacity-100 transition-opacity duration-300">
                <a-dropdown :trigger="['click']">
                  <a-button type="text" class="text-gray-500 hover:text-gray-700 mr-1">
                    Shots <down-outlined />
                  </a-button>
                  <template #overlay>
                    <a-menu>
                      <a-menu-item key="1">所有内容</a-menu-item>
                      <a-menu-item key="2">图片</a-menu-item>
                      <a-menu-item key="3">视频</a-menu-item>
                    </a-menu>
                  </template>
                </a-dropdown>
                <a-button
                    type="primary"
                    shape="circle"
                    class="flex items-center justify-center bg-pink-500/80 hover:bg-pink-500 border-0 mr-2 transition-colors duration-300"
                >
                  <template #icon><SearchOutlined /></template>
                </a-button>
              </div>
            </div>
          </div>

          <!-- 右侧区域：导航菜单和头像 -->
          <div class="flex items-center gap-8">
            <!-- 移动端搜索按钮 -->
            <a-button
                v-if="false"
                type="text"
                class="md:hidden"
                @click="showMobileSearch = true"
            >
              <template #icon><SearchOutlined /></template>
            </a-button>

            <!-- 桌面端导航菜单 -->
            <nav class="hidden md:flex items-center gap-6">
              <router-link
                  v-for="item in navItems"
                  :key="item.key"
                  :to="item.href"
                  v-show="isLoggedIn||!item.login"
                  class="text-gray-500/60 hover:text-gray-900 transition-colors duration-300"
              >
                <a-avatar :size="48" class="layout-content-menu-item" :src="item.img"  > {{ item.label }}</a-avatar>
              </router-link>
            </nav>

            <a-avatar :size="48" class="layout-content-menu-item" @click="showLoginDialog" v-if="!isLoggedIn" >登录</a-avatar>
            <a-avatar :size="48" class="layout-content-menu-item" @click="showAccountDialog" :src="avatar" v-else ></a-avatar>
          </div>
        </div>
      </div>
    </div>

    <!-- 移动端搜索抽屉 -->
    <a-drawer
        :visible="showMobileSearch"
        placement="top"
        height="auto"
        :closable="false"
        @close="showMobileSearch = false"
    >
      <div class="py-4">
        <div class="relative">
          <input
              type="search"
              placeholder="搜索..."
              class="w-full rounded-full border-0 pl-4 pr-24 py-2.5 bg-[#98FB98]/10 hover:bg-[#98FB98]/20 focus:bg-[#98FB98]/30 transition-colors duration-300 focus:outline-none focus:ring-1 focus:ring-green-200/30"
          />
          <div class="absolute right-0 top-0 bottom-0 flex items-center">
            <a-dropdown :trigger="['click']">
              <a-button type="text" class="text-gray-500 hover:text-gray-700 mr-1">
                Shots <down-outlined />
              </a-button>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">所有内容</a-menu-item>
                  <a-menu-item key="2">图片</a-menu-item>
                  <a-menu-item key="3">视频</a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
            <a-button
                type="primary"
                shape="circle"
                class="flex items-center justify-center bg-pink-500/80 hover:bg-pink-500 border-0 mr-2"
                @click="showMobileSearch = false"
            >
              <template #icon><SearchOutlined /></template>
            </a-button>
          </div>
        </div>
      </div>
    </a-drawer>

    <!-- 移动端侧边菜单 -->
    <a-drawer
        :open="showMobileMenu"
        placement="left"
        :closable="false"
        @close="showMobileMenu = false"
        class="mobile-menu"
    >
      <div class="flex flex-col h-full">
        <!-- 顶部：关闭按钮和Logo -->
        <div class="flex justify-between items-center mb-8">
          <div class="flex items-center">
            <a-button
                type="text"
                @click="showMobileMenu = false"
                class="!text-gray-600 mr-4"
            >
              <template #icon><CloseOutlined /></template>
            </a-button>
            <img src="@/assets/image/logo.png" alt="Logo" class="h-12 w-auto" />
          </div>
        </div>

        <!-- 菜单项 -->
        <div class="flex-1 space-y-6">
          <div v-for="item in navItems" :key="item.key" >
            <div class="flex items-center justify-between" v-show="isLoggedIn||!item.login">
              <router-link
                  :to="item.href"
                  class="text-lg font-medium text-gray-900 hover:text-primary transition-colors"
                  @click="showMobileMenu = false"
              >
                {{ item.label }}
              </router-link>
              <right-outlined class="text-gray-400" />
            </div>
            <div class="mt-2 border-b border-gray-100"></div>
          </div>
        </div>
      </div>
    </a-drawer>

    <!-- 内容区域 -->
    <router-view  />
    <!-- 登录 -->
    <Login ref="loginRef"/>
    <Account ref="accountRef"/>
  </div>
</template>

<script setup>
import Login from "@/components/Login.vue";
import Account from "@/components/Account.vue"
import avatar from "@/assets/image/default-avatar.png"
import logo from "@/assets/image/logo.png"
import {onMounted, ref, watch} from "vue"
import { useRouter } from 'vue-router';
const router = useRouter();
import { useAuthStore  } from '@/store/auth.js'
const { isLoggedIn, state, setLoginData, getLoginToken } = useAuthStore();

const loginRef = ref(null);
const accountRef = ref(null);
const showLoginDialog =()=>{
  loginRef.value.show()
}
const showAccountDialog = ()=>{
  accountRef.value.show()
}
onMounted(function () {
  if(isLoggedIn){
    const tk = getLoginToken()
    if(!tk||tk=='undefined'){
      return
    }
    setLoginData(tk)
  }
})

import {
  MenuOutlined,
  CloseOutlined,
  SearchOutlined,
  DownOutlined,
  RightOutlined
} from '@ant-design/icons-vue'

const showMobileMenu = ref(false)
const showMobileSearch = ref(false)
const navItems = [
  {key: 'diary', label: '日记', href: '/blog',login:false},
  {key: 'travel', label: '旅行', href: '/travel',login:false},
  {key: 'photo', label: '墙', href: '/picture',login:true},
  // {key: 'demo', label: 'DEMO', href: '/demo',login:true},
  {key: 'about', label: '关于', href: '/about',login:true},
]


</script>

<style scoped>
/* 自定义抽屉样式 */
:deep(.ant-drawer-content-wrapper) {
  max-width: 80vw;
}

:deep(.ant-drawer-body) {
  padding: 24px;
}
.layout {
  height: 100%;
  width: 100%;
  overflow: auto;
  background: transparent  url("/src/assets/image/bg02.gif");
  background-size: cover;
  display: flex;
  flex-direction: column;
}
.layout .layout-header {
  transition: background-color 2s;
}
.layout .layout-header:hover {
  background-color: #fffdfd;
}
/* 移动端菜单链接样式 */
.mobile-menu a {
  @apply font-medium;
}

.mobile-menu a:active {
  @apply bg-gray-50;
}

/* 搜索按钮样式 */
:deep(.ant-btn-primary) {
  @apply shadow-none;
}

:deep(.ant-btn-primary:hover) {
  @apply shadow-lg;
}

/* 下拉菜单样式 */
:deep(.ant-dropdown-menu) {
  @apply rounded-lg shadow-lg border border-gray-100;
}

:deep(.ant-dropdown-menu-item) {
  @apply py-2 px-4;
}

/* 内容区域样式 */
.layout-content {
  min-height: calc(100vh - 4rem);
  max-height: max-content;
}

/* 移动端搜索抽屉样式 */
:deep(.ant-drawer-top.ant-drawer-open .ant-drawer-content-wrapper) {
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
}
</style>

