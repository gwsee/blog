<template>
  <div class="min-h-screen bg-gray-50/50">
    <!-- 主要内容区域 -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- 标题区域 -->
      <div class="mb-8" v-if="false">
        <h1 class="text-4xl font-bold text-gray-900 mb-4 font-sans tracking-tight">{{ travel.title }}</h1>
        <div class="flex items-center space-x-6">
          <div class="flex items-center space-x-2">
            <a-avatar :src="travel.author.avatar" />
            <span class="text-gray-700 font-medium">{{ travel.author.name }}</span>
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
      <div class="relative mb-8 rounded-2xl overflow-hidden shadow-lg ">
        <a-carousel ref="carouselRef">
          <div v-for="(image, index) in travel.images" :key="index"
               class="h-full relative group cursor-pointer"
               @click="showLightbox(index)">
            <img :src="image.url"
                 :alt="image.description"
                 class="w-full h-full object-cover object-center" />
            <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-all duration-300">
              <div class="absolute bottom-8 left-8 right-8">
                <h3 class="text-white text-2xl font-semibold mb-2">{{ image.title }}</h3>
                <p class="text-white/90 text-lg">{{ image.description }}</p>
              </div>
            </div>
          </div>
        </a-carousel>
        <!-- 自定义导航按钮 -->

        <button @click="prevSlide" class="absolute left-4 top-1/2 transform -translate-y-1/2 bg-white/90 hover:bg-white text-gray-800 p-3 rounded-full transition-all duration-300 shadow-lg backdrop-blur-sm">
          <left-outlined class="text-xl" />
        </button>
        <button @click="nextSlide" class="absolute right-4 top-1/2 transform -translate-y-1/2 bg-white/90 hover:bg-white text-gray-800 p-3 rounded-full transition-all duration-300 shadow-lg backdrop-blur-sm">
          <right-outlined class="text-xl" />
        </button>
      </div>

      <!-- 内容区域 -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- 左侧主要内容 -->
        <div class="lg:col-span-2 space-y-8">
          <!-- 简介 -->
          <a-card class="shadow-sm rounded-2xl border-0">
            <template #title>
              <div class="flex items-center space-x-2">
                <info-circle-outlined />
                <span class="font-semibold">旅行简介</span>
              </div>
            </template>
            <div class="prose max-w-none">
              {{ travel.description }}
            </div>
          </a-card>

          <!-- 视频展示 -->
          <a-card class="shadow-sm rounded-2xl border-0">
            <template #title>
              <div class="flex items-center space-x-2">
                <video-camera-outlined />
                <span class="font-semibold">精彩视频</span>
              </div>
            </template>
            <div class="aspect-video rounded-xl overflow-hidden">
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
          <!-- 作者信息 -->
          <a-card class="shadow-sm rounded-2xl border-0">
            <template #title>
              <div class="flex items-center space-x-2">
                <user-outlined />
                <span class="font-semibold">关于作者</span>
              </div>
            </template>
            <div class="text-center">
              <a-avatar
                  :src="travel.author.avatar"
                  :size="64"
                  class="mb-4" />
              <h3 class="text-lg font-semibold mb-2">{{ travel.author.name }}</h3>
              <p class="text-gray-500 mb-4">{{ travel.author.bio }}</p>
              <a-button type="primary" block class="rounded-full h-10">
                关注作者
              </a-button>
            </div>
          </a-card>

          <!-- 统计信息 -->
          <a-card class="shadow-sm rounded-2xl border-0">
            <template #title>
              <div class="flex items-center space-x-2">
                <bar-chart-outlined />
                <span class="font-semibold">数据统计</span>
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

    <!-- 悬浮返回按钮 -->
    <a-button
        type="primary"
        shape="circle"
        size="large"
        @click="router.back()"
        class="fixed bottom-8 right-8 w-14 h-14 flex items-center justify-center shadow-lg hover:scale-110 transition-transform duration-300"
    >
      <template #icon><left-outlined class="text-xl" /></template>
    </a-button>

    <!-- 图片查看器 -->
    <vue-easy-lightbox
        :visible="visibleLightbox"
        :imgs="travel.images.map(img => ({ url: img.url, title: img.title, description: img.description }))"
        :index="lightboxIndex"
        @hide="hideLightbox"
    ></vue-easy-lightbox>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import VueEasyLightbox from 'vue-easy-lightbox'
import {
  LeftCircleOutlined,
  RightCircleOutlined,
  LeftOutlined,
  RightOutlined,
  ShareAltOutlined,
  HeartOutlined,
  EyeOutlined,
  LikeOutlined,
  StarOutlined,
  InfoCircleOutlined,
  VideoCameraOutlined,
  UserOutlined,
  BarChartOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const carouselRef = ref(0)
const currentSlide = ref(0)
const visibleLightbox = ref(false)
const lightboxIndex = ref(0)

// 示例数据
const travel = ref({
  title: '探索巴厘岛的神秘海滩',
  description: `巴厘岛是印度尼西亚最受欢迎的旅游目的地之一，以其独特的文化、美丽的海滩和壮观的自然风光而闻名。
  在这次探索之旅中，我们将带您深入发现巴厘岛最神秘的海滩，感受原始自然的魅力。从隐藏的海湾到壮丽的悬崖，
  每一处都是摄影爱好者的天堂。我们不仅会分享这些地方的美，还会为您介绍当地的文化习俗和美食推荐。`,
  images: [
    {
      url: 'https://gwsee-blog.oss-cn-chengdu.aliyuncs.com/image/png/20241211/10002.png?x-oss-credential=LTAI5t5j88DF7KKibWGDMesb%2F20241211%2Fcn-chengdu%2Foss%2Faliyun_v4_request&x-oss-date=20241211T031518Z&x-oss-expires=86400&x-oss-signature=a5c076c4d7d411bc889d0d5785b2cd2bbe6727dec8101da8a2d6a08c28768e6f&x-oss-signature-version=OSS4-HMAC-SHA256',
      title: '金巴兰海滩日落',
      description: '日落时分的金巴兰海滩，天空被染成金色，海浪轻轻拍打着沙滩。'
    },
    {
      url: 'https://gwsee-blog.oss-cn-chengdu.aliyuncs.com/image/png/20241211/10002.png?x-oss-credential=LTAI5t5j88DF7KKibWGDMesb%2F20241211%2Fcn-chengdu%2Foss%2Faliyun_v4_request&x-oss-date=20241211T031518Z&x-oss-expires=86400&x-oss-signature=a5c076c4d7d411bc889d0d5785b2cd2bbe6727dec8101da8a2d6a08c28768e6f&x-oss-signature-version=OSS4-HMAC-SHA256',
      title: '努沙杜瓦清晨',
      description: '清晨的努沙杜瓦海滩，晨曦中的椰林和细腻的白沙。'
    },
    {
      url: 'https://gwsee-blog.oss-cn-chengdu.aliyuncs.com/image/png/20241211/10002.png?x-oss-credential=LTAI5t5j88DF7KKibWGDMesb%2F20241211%2Fcn-chengdu%2Foss%2Faliyun_v4_request&x-oss-date=20241211T031518Z&x-oss-expires=86400&x-oss-signature=a5c076c4d7d411bc889d0d5785b2cd2bbe6727dec8101da8a2d6a08c28768e6f&x-oss-signature-version=OSS4-HMAC-SHA256',
      title: '乌鲁瓦图悬崖',
      description: '壮观的乌鲁瓦图悬崖，印度洋的波涛在脚下翻滚。'
    },
    {
      url: 'https://gwsee-blog.oss-cn-chengdu.aliyuncs.com/image/png/20241211/10002.png?x-oss-credential=LTAI5t5j88DF7KKibWGDMesb%2F20241211%2Fcn-chengdu%2Foss%2Faliyun_v4_request&x-oss-date=20241211T031518Z&x-oss-expires=86400&x-oss-signature=a5c076c4d7d411bc889d0d5785b2cd2bbe6727dec8101da8a2d6a08c28768e6f&x-oss-signature-version=OSS4-HMAC-SHA256',
      title: '蓝梦岛水晶湾',
      description: '蓝梦岛的水晶湾，清澈的海水和丰富的海洋生态。'
    }
  ],
  videoUrl: '/videos/bali-highlights.mp4',
  videoPoster: '/images/video-cover.jpg',
  author: {
    name: '旅行摄影师小王',
    avatar: '/images/avatar.jpg',
    bio: '专业旅行摄影师，足迹遍布30多个国家，热爱分享旅行故事和摄影技巧。'
  },
  views: 12580,
  likes: 2390,
  favorites: 1678
})

const prevSlide = () => {
  carouselRef.value =  carouselRef.value-1
}

const nextSlide = () => {
  carouselRef.value =  carouselRef.value+1
}

const goToSlide = (index) => {
  carouselRef.value?.goTo(index)
}

const showLightbox = (index) => {
  lightboxIndex.value = index
  visibleLightbox.value = true
}

const hideLightbox = () => {
  visibleLightbox.value = false
}

// 监听轮播图变化
watch(() => carouselRef.value?.current, (newVal) => {
  if (newVal !== undefined) {
    currentSlide.value = newVal
  }
})

onMounted(() => {
  currentSlide.value = carouselRef.value?.current || 0
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
  border-radius: 1rem;
  overflow: hidden;
}

/* 响应式调整 */
@media (max-width: 768px) {
  :deep(.ant-carousel .slick-slide) {
    height: 50vh !important;
  }
}

/* 卡片样式优化 */
:deep(.ant-card) {
  border-radius: 1rem;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

:deep(.ant-card:hover) {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}

:deep(.ant-card-head) {
  border-bottom: none;
  padding: 1.25rem 1.5rem;
}

:deep(.ant-card-body) {
  padding: 1.5rem;
}

/* 按钮样式优化 */
:deep(.ant-btn-primary) {
  background: #1890ff;
  border-color: #1890ff;
  box-shadow: 0 2px 0 rgba(24, 144, 255, 0.1);
}

:deep(.ant-btn-primary:hover) {
  background: #40a9ff;
  border-color: #40a9ff;
  transform: translateY(-1px);
}

/* 图片查看器自定义样式 */
:deep(.vel-modal) {
  background-color: rgba(0, 0, 0, 0.85);
}

:deep(.vel-img-title) {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
}

:deep(.vel-img-description) {
  font-size:rem;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.5;
}
</style>

