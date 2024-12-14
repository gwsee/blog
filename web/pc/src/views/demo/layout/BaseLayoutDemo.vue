<template>
  <div class="w-full">
    <!-- 主导航栏 -->
    <div class="sticky top-0 z-50 w-full bg-white border-b">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 items-center justify-between">
          <!-- 左侧区域：移动端菜单按钮和Logo -->
          <div class="flex items-center gap-4">
            <!-- 移动端菜单按钮 -->
            <a-button
                type="text"
                class="md:hidden"
                @click="showMobileMenu = true"
            >
              <template #icon><MenuOutlined /></template>
            </a-button>

            <!-- Logo -->
            <router-link to="/" class="flex items-center">
              <img src="@/assets/image/logo.png" alt="Logo" class="h-8 w-auto" />
            </router-link>
          </div>

          <!-- 中间区域：搜索框（桌面端） -->
          <div class="hidden md:flex flex-1 justify-center max-w-2xl mx-8">
            <div class="relative w-full">
              <input
                  type="search"
                  placeholder="搜索..."
                  class="w-full rounded-lg border border-gray-200 pl-4 pr-24 py-2 bg-gray-50/50 focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
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
                    class="flex items-center justify-center bg-pink-500 hover:bg-pink-600 border-0 mr-2"
                >
                  <template #icon><SearchOutlined /></template>
                </a-button>
              </div>
            </div>
          </div>

          <!-- 右侧区域：导航菜单和头像 -->
          <div class="flex items-center gap-6">
            <!-- 移动端搜索按钮 -->
            <a-button
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
                  class="text-gray-600 hover:text-gray-900 transition-colors"
              >
                {{ item.label }}
              </router-link>
            </nav>

            <!-- 头像 -->
            <a-avatar
                :src="'/avatar.jpg'"
                :size="40"
                class="cursor-pointer ring-2 ring-primary ring-offset-2 hover:ring-primary/80"
            />
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
              class="w-full rounded-lg border border-gray-200 pl-4 pr-24 py-2 bg-gray-50/50 focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
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
                class="flex items-center justify-center bg-pink-500 hover:bg-pink-600 border-0 mr-2"
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
        :visible="showMobileMenu"
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
            <img src="@/assets/image/logo.png" alt="Logo" class="h-8 w-auto" />
          </div>
          <a-button type="primary" class="rounded-full">
            登录
          </a-button>
        </div>

        <!-- 菜单项 -->
        <div class="flex-1 space-y-6">
          <div v-for="item in navItems" :key="item.key">
            <div class="flex items-center justify-between">
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
    <router-view class="layout-content" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
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
  { key: 'diary', label: '日记', href: '/diary' },
  { key: 'travel', label: '旅行', href: '/travel' },
  { key: 'photo', label: '图片墙', href: '/photos' },
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
}

/* 移动端搜索抽屉样式 */
:deep(.ant-drawer-top.ant-drawer-open .ant-drawer-content-wrapper) {
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
}
</style>

