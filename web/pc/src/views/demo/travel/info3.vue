<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 顶部导航栏 -->
    <a-affix>
      <nav class="bg-white/80 backdrop-blur-md border-b border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="flex justify-between h-16">
            <div class="flex items-center">
              <a-button type="link" @click="router.back()">
                <template #icon><left-outlined /></template>
                返回
              </a-button>
            </div>
            <div class="flex items-center space-x-4">
              <a-button shape="circle">
                <template #icon><share-alt-outlined /></template>
              </a-button>
              <a-button shape="circle">
                <template #icon><heart-outlined /></template>
              </a-button>
            </div>
          </div>
        </div>
      </nav>
    </a-affix>

    <!-- 主要内容区域 -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- 标题区域 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-gray-900 mb-4">{{ travel.title }}</h1>
        <div class="flex items-center space-x-6">
          <div class="flex items-center space-x-2">
            <a-avatar :src="travel.author.avatar" />
            <span class="text-gray-700">{{ travel.author.name }}</span>
          </div>
          <div class="flex items-center space-x-4 text-gray-500 text-sm">
            <span class="flex items-center">
              <eye-outlined class="mr-1" />
              {{ travel.views }}
            </span>
            <span class="flex items-center">
              <like-outlined class="mr-1" />
              {{ travel.likes }}
            </span>
            <span class="flex items-center">
              <star-outlined class="mr-1" />
              {{ travel.favorites }}
            </span>
          </div>
        </div>
      </div>

      <!-- 图片轮播 -->
      <div class="relative mb-8 rounded-2xl overflow-hidden">
        <a-carousel class="aspect-[16/9] h-[70vh]" :autoplay="false" :dots="false" ref="carousel">
          <div v-for="(image, index) in travel.images" :key="index"
               class="h-full relative group cursor-pointer"
               @click="showLightbox(index)">
            <img :src="image"
                 :alt="`Travel image ${index + 1}`"
                 class="w-full h-full object-cover object-center" />
            <div class="absolute inset-0 bg-black/30 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
              <div class="absolute bottom-4 left-4 text-white">
                <p class="text-lg font-semibold">{{ travel.imageDescriptions[index] }}</p>
              </div>
            </div>
          </div>
        </a-carousel>
        <div class="absolute bottom-4 right-4 bg-black/50 text-white px-3 py-1 rounded-full text-sm">
          {{ currentSlide + 1 }} / {{ travel.images.length }}
        </div>
        <button @click="prevSlide" class="absolute left-4 top-1/2 transform -translate-y-1/2 bg-white/80 hover:bg-white text-gray-800 p-2 rounded-full transition-colors duration-300">
          <left-outlined />
        </button>
        <button @click="nextSlide" class="absolute right-4 top-1/2 transform -translate-y-1/2 bg-white/80 hover:bg-white text-gray-800 p-2 rounded-full transition-colors duration-300">
          <right-outlined />
        </button>
      </div>

      <!-- 内容区域 -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- 左侧主要内容 -->
        <div class="lg:col-span-2 space-y-8">
          <!-- 简介 -->
          <a-card class="shadow-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <info-circle-outlined />
                <span>旅行简介</span>
              </div>
            </template>
            <div class="prose max-w-none">
              {{ travel.description }}
            </div>
          </a-card>

          <!-- 视频展示 -->
          <a-card class="shadow-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <video-camera-outlined />
                <span>精彩视频</span>
              </div>
            </template>
            <div class="aspect-video h-[50vh] rounded-lg overflow-hidden">
              <video
                  controls
                  class="w-full h-full object-cover object-center"
                  :poster="travel.videoPoster">
                <source :src="travel.videoUrl" type="video/mp4">
                您的浏览器不支持视频播放。
              </video>
            </div>
          </a-card>
        </div>

        <!-- 右侧信息栏 -->
        <div class="space-y-6">
          <!-- 时间信息 -->
          <a-card class="shadow-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <clock-circle-outlined />
                <span>时间信息</span>
              </div>
            </template>
            <div class="space-y-4">
              <div class="flex justify-between items-center">
                <span class="text-gray-500">创建时间</span>
                <span>{{ formatDate(travel.createdAt) }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-gray-500">更新时间</span>
                <span>{{ formatDate(travel.updatedAt) }}</span>
              </div>
            </div>
          </a-card>

          <!-- 作者信息 -->
          <a-card class="shadow-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <user-outlined />
                <span>关于作者</span>
              </div>
            </template>
            <div class="text-center">
              <a-avatar
                  :src="travel.author.avatar"
                  :size="64"
                  class="mb-4" />
              <h3 class="text-lg font-semibold mb-2">{{ travel.author.name }}</h3>
              <p class="text-gray-500 mb-4">{{ travel.author.bio }}</p>
              <a-button type="primary" block>
                关注作者
              </a-button>
            </div>
          </a-card>

          <!-- 统计信息 -->
          <a-card class="shadow-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <bar-chart-outlined />
                <span>数据统计</span>
              </div>
            </template>
            <div class="grid grid-cols-3 gap-4 text-center">
              <div>
                <p class="text-2xl font-semibold">{{ travel.views }}</p>
                <p class="text-gray-500">浏览</p>
              </div>
              <div>
                <p class="text-2xl font-semibold">{{ travel.likes }}</p>
                <p class="text-gray-500">点赞</p>
              </div>
              <div>
                <p class="text-2xl font-semibold">{{ travel.favorites }}</p>
                <p class="text-gray-500">收藏</p>
              </div>
            </div>
          </a-card>
        </div>
      </div>
    </main>

    <!-- 图片查看器 -->
    <vue-easy-lightbox
        :visible="visibleLightbox"
        :imgs="travel.images"
        :index="lightboxIndex"
        @hide="hideLightbox"
    ></vue-easy-lightbox>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import VueEasyLightbox from 'vue-easy-lightbox'
import {
  LeftOutlined,
  RightOutlined,
  ShareAltOutlined,
  HeartOutlined,
  EyeOutlined,
  LikeOutlined,
  StarOutlined,
  InfoCircleOutlined,
  VideoCameraOutlined,
  ClockCircleOutlined,
  UserOutlined,
  BarChartOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const carousel = ref(null)
const currentSlide = ref(0)
const visibleLightbox = ref(false)
const lightboxIndex = ref(0)

// 格式化日期的函数
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// 示例数据
const travel = ref({
  title: '探索巴厘岛的神秘海滩',
  description: `巴厘岛是印度尼西亚最受欢迎的旅游目的地之一，以其独特的文化、美丽的海滩和壮观的自然风光而闻名。
  在这次探索之旅中，我们将带您深入发现巴厘岛最神秘的海滩，感受原始自然的魅力。从隐藏的海湾到壮丽的悬崖，
  每一处都是摄影爱好者的天堂。我们不仅会分享这些地方的美，还会为您介绍当地的文化习俗和美食推荐。`,
  images: [
    'http://127.0.0.1:1000/v1/file/43837c94ef2ed8ba75ca1bcc6dd55c52',
    '/placeholder.svg?height=720&width=1280',
    '/placeholder.svg?height=720&width=1280',
    '/placeholder.svg?height=720&width=1280'
  ],
  imageDescriptions: [
    '日落时分的金巴兰海滩',
    '努沙杜瓦的清晨',
    '乌鲁瓦图悬崖',
    '蓝梦岛的水晶湾'
  ],
  videoUrl: 'https://example.com/video.mp4',
  videoPoster: '/placeholder.svg?height=720&width=1280',
  createdAt: '2024-01-15T08:00:00.000Z',
  updatedAt: '2024-01-16T10:30:00.000Z',
  author: {
    name: '旅行摄影师小王',
    avatar: '/placeholder.svg?height=64&width=64',
    bio: '专业旅行摄影师，足迹遍布30多个国家，热爱分享旅行故事和摄影技巧。'
  },
  views: 12580,
  likes: 2390,
  favorites: 1678
})

const prevSlide = () => {
  carousel.value.prev()
}

const nextSlide = () => {
  carousel.value.next()
}

const showLightbox = (index) => {
  lightboxIndex.value = index
  visibleLightbox.value = true
}

const hideLightbox = () => {
  visibleLightbox.value = false
}

onMounted(() => {
  if (carousel.value) {
    carousel.value.$on('change', (current) => {
      currentSlide.value = current
    })
  }
})
</script>

<style scoped>
.prose {
  max-width: 65ch;
  line-height: 1.75;
}

.prose p {
  margin-bottom: 1.5em;
}

/* 自定义轮播样式 */
:deep(.ant-carousel .slick-slide) {
  height: 70vh !important;
}

@media (max-width: 768px) {
  :deep(.ant-carousel .slick-slide) {
    height: 50vh !important;
  }
}

/* 轮播图导航按钮样式 */
:deep(.ant-carousel .slick-prev),
:deep(.ant-carousel .slick-next) {
  font-size: 20px;
  color: white;
}

:deep(.ant-carousel .slick-prev) {
  left: 10px;
  z-index: 2;
}

:deep(.ant-carousel .slick-next) {
  right: 10px;
  z-index: 2;
}
</style>

