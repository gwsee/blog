<template>
  <div class="min-h-screen ">
    <!-- 主要内容区域 -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <!-- 标题区域 -->
      <div class="mb-6">
        <h1 class="text-4xl font-bold text-gray-900 mb-4 font-sans tracking-tight">{{ travel.title }}</h1>
        <div class="flex items-center space-x-6">
          <div class="flex items-center space-x-2">
            <a-avatar :src="travel.account.avatar||avatar" />
            <span class="text-gray-700 font-medium">{{ travel.account.nickname }}</span>
          </div>
          <div class="flex items-center space-x-4 text-gray-500 text-sm">
            <span class="flex items-center"  @click.stop="travelCollectAc(travel)">
              <star-outlined class="mr-1" v-if="!travel.isCollect"  />
              <star-filled class="mr-1  text-red-500" v-else  />
               {{ travel.collectNum>999?'999+': (travel.collectNum||0)}}
            </span>
            <span class="flex items-center"   @click.stop="travelThumbAc(travel)">
              <heart-outlined class="mr-1"  v-if="!travel.isThumb" />
              <HeartFilled class="mr-1 text-red-500"  v-else/>
              {{ travel.thumbNum>999?'999+': (travel.thumbNum||0)}}
            </span>
            <span class="flex items-center">
              <eye-outlined class="mr-1" />
                {{ travel.browseNum>999?'999+': (travel.browseNum||0)}}
            </span>
          </div>
        </div>
      </div>

      <!-- 图片轮播 -->
      <div class="relative mb-6 rounded-3xl overflow-hidden shadow-xl bg-black/5 backdrop-blur-sm">
        <a-carousel class="w-full" :autoplay="false" ref="carouselRef">
          <div v-for="(image, index) in travel.photos" :key="index"
               class="relative group cursor-zoom-in"
               @click="showLightbox(index)">
            <div class="w-full aspect-[16/9] relative">
              <img :src="filePrefix+image"
                   class="w-full h-full object-contain bg-black/30" />
              <!-- 渐变遮罩 -->
              <div class="absolute inset-0 bg-gradient-to-b from-black/10 to-black/60 opacity-0 group-hover:opacity-100 transition-all duration-300">
                <div class="absolute bottom-8 left-8 right-8 transform translate-y-4 group-hover:translate-y-0 transition-transform duration-300" v-if="false">
                  <h3 class="text-white text-2xl font-semibold mb-2">{{ image.title }}</h3>
                  <p class="text-white/90 text-lg leading-relaxed">{{ image.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </a-carousel>

        <!-- 自定义导航按钮 -->
        <div class="absolute bottom-6 left-1/2 transform -translate-x-1/2 flex space-x-3 z-10">
          <button v-for="(_, index) in travel.images"
                  :key="index"
                  @click="goToSlide(index)"
                  class="w-2.5 h-2.5 rounded-full transition-all duration-300 bg-white/40 hover:bg-white/80"
                  :class="[currentSlide === index ? 'w-8 bg-white' : '']">
          </button>
        </div>

        <!-- 导航按钮 -->
        <button @click="prevSlide"
                class="absolute left-4 top-1/2 transform -translate-y-1/2 w-12 h-12 flex items-center justify-center rounded-full bg-white/10 backdrop-blur-md hover:bg-white/20 transition-all duration-300 group">
          <left-outlined class="text-white  opacity-60 group-hover:opacity-100" />
        </button>
        <button @click="nextSlide"
                class="absolute right-4 top-1/2 transform -translate-y-1/2 w-12 h-12 flex items-center justify-center rounded-full bg-white/10 backdrop-blur-md hover:bg-white/20 transition-all duration-300 group">
          <right-outlined class="text-white  opacity-60 group-hover:opacity-100" />
        </button>
      </div>

      <!-- 内容区域 -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- 左侧主要内容 -->
        <div class="lg:col-span-2 space-y-6">
          <!-- 简介 -->
          <a-card class="shadow-lg rounded-2xl border-0 bg-white/80 backdrop-blur-sm">
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
          <a-card class="shadow-lg rounded-2xl border-0 bg-white/80 backdrop-blur-sm" v-if="travel.video">
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
                  :poster="filePrefix+travel.photos[0]">
                <source :src="filePrefix+travel.video" type="video/mp4">
                您的浏览器不支持视频播放。
              </video>
            </div>
          </a-card>
        </div>

        <!-- 右侧信息栏 -->
        <div class="space-y-6">
          <!-- 作者信息 -->
          <a-card class="shadow-lg rounded-2xl border-0 bg-white/80 backdrop-blur-sm">
            <template #title>
              <div class="flex items-center space-x-2">
                <user-outlined />
                <span class="font-semibold">关于作者</span>
              </div>
            </template>
            <div class="text-center">
              <a-avatar
                  :src="travel.account.avatar||avatar"
                  :size="64"
                  class="mb-4 ring-4 ring-white shadow-lg" />
              <h3 class="text-lg font-semibold mb-2">{{ travel.account.nickname }}</h3>
              <p class="text-gray-500 mb-4">{{ travel.account.description ||'这个人好懒哦,简介都不写个' }}</p>
              <a-button type="primary" block class="rounded-full h-10 shadow-md hover:shadow-lg transition-shadow" v-if="false">
                关注作者
              </a-button>
            </div>
          </a-card>
          <!-- 统计信息 -->
          <a-card class="shadow-lg rounded-2xl border-0 bg-white/80 backdrop-blur-sm" v-if="false">
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

    <div  class="fixed bottom-8 right-8 w-14 h-14 flex items-center justify-center transition-all duration-300">
      <a-button
          shape="circle"
          size="large"
          v-if="state.user&&state.user.id === travel.account.id"
          class="travel-info-handle-button"
          style="background-color: gray"
          @click.stop="editTravel(travel)"
      >
        <template #icon>
          <EditOutlined class="text-white " />
        </template>
      </a-button>
      <a-button
          shape="circle"
          size="large"
          class="travel-info-handle-button"
          @click="router.back()"
          style="background-color: gray"
      >
        <template #icon><left-outlined class="text-white " /></template>
      </a-button>
    </div>


    <!-- 图片查看器 -->
    <vue-easy-lightbox
        :visible="visibleLightbox"
        :imgs="travel.photos.map(img => (filePrefix+img))"
        :index="lightboxIndex"
        @hide="hideLightbox"
        moveDisabled
        class="lightbox-custom"
    ></vue-easy-lightbox>
  </div>
</template>

<script setup>
import avatar from "@/assets/image/default-avatar.png"
import {travelGet, travelCollect, travelThumb} from "@/api/travel";
import {ref, onMounted, watch, reactive} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import { useAuthStore  } from '@/store/auth'
const { state } = useAuthStore();
import {filePrefix} from "@/api/tool";
import VueEasyLightbox from 'vue-easy-lightbox'
import {
  EditOutlined,
  LeftOutlined,
  RightOutlined,
  ShareAltOutlined,
  HeartOutlined,
  EyeOutlined,
  LikeOutlined,
  StarOutlined,
  StarFilled,
  LikeFilled,
  HeartFilled,
  InfoCircleOutlined,
  VideoCameraOutlined,
  UserOutlined,
  BarChartOutlined,
  LeftCircleOutlined, RightCircleOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const carouselRef = ref(null)
const currentSlide = ref(0)
const visibleLightbox = ref(false)
const lightboxIndex = ref(0)
const prevSlide = () => {
  carouselRef.value?.prev()
}

const nextSlide = () => {
  carouselRef.value?.next()
}

const goToSlide = (index) => {
  carouselRef.value?.goTo(index)
}
const images = ref([
  'https://via.placeholder.com/600x400.png?text=Image+1',
  'https://via.placeholder.com/600x400.png?text=Image+2',
  'https://via.placeholder.com/600x400.png?text=Image+3',
])
const editTravel = (travel) => {
  router.push(`/travel/manage/${travel.id}`)
}
// 示例数据
const travel = ref({
  photos:[],
  account:{},
})


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
const formState = reactive({});
const loading = ref(false);
onMounted(() => {
  currentSlide.value = carouselRef.value?.current || 0
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    router.push("/travel")
    return
  }
  loading.value = true;
  travelGet({id:id}).then(res=>{
    if(res&&res.code===200){
      let obj = res.data
      travel.value = obj
    }
  }).finally(()=>{
    loading.value = false;
  })
})
const travelCollectAc = async (travel) => {
  try {
    travel.collectNum=travel.collectNum||0
    travel.isCollect = !travel.isCollect
    travel.collectNum +=  travel.isCollect? 1 : -1
    await travelCollect({id:travel.id,do:travel.isCollect})
  } catch (error) {
    console.error('Error toggling favorite:', error)
  }
}

const travelThumbAc = async (travel) => {
  try {
    travel.thumbNum=  travel.thumbNum||0
    travel.isThumb = !travel.isThumb
    travel.thumbNum += travel.isThumb ? 1 : -1
    await travelThumb({id:travel.id,do:travel.isThumb})

  } catch (error) {
    console.error('Error toggling like:', error)
  }
}

</script>

<style scoped>
.prose {
  max-width: 65ch;
  line-height: 1.75;
  font-size: 1.125rem;
}

.prose p {
  margin-bottom: 1.5em;
  color: #374151;
}

/* 自定义轮播样式 */
:deep(.ant-carousel) {
  width: 100%;
}

:deep(.ant-carousel .slick-slide) {
  overflow: hidden;
}

:deep(.ant-carousel .slick-slide > div) {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 卡片样式优化 */
:deep(.ant-card) {
  transition: all 0.3s ease;
}

:deep(.ant-card:hover) {
  transform: translateY(-2px);
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
}

/* 图片查看器自定义样式 */
:deep(.vel-modal) {
  background-color: rgba(0, 0, 0, 0.95);
}

:deep(.vel-img-title) {
  font-size: 1.5rem;
  margin-bottom: 0.75rem;
  color: white;
}

:deep(.vel-img-description) {
  font-size: 1.125rem;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.6;
}

:deep(.vel-toolbar) {
  background: transparent !important;
}

:deep(.vel-toolbar .btn-close),
:deep(.vel-toolbar .btn-prev),
:deep(.vel-toolbar .btn-next) {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(8px);
  border-radius: 9999px;
  padding: 12px;
  transition: all 0.3s ease;
}

:deep(.vel-toolbar .btn-close:hover),
:deep(.vel-toolbar .btn-prev:hover),
:deep(.vel-toolbar .btn-next:hover) {
  background: rgba(255, 255, 255, 0.2);
}
.travel-info-handle-button{
  background-color: gray;
  display: flex;
  margin-right: 10px;
  flex-direction: column;
  flex-wrap: wrap;
  align-content: space-around;
  justify-content: center;
}
</style>

