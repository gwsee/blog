<template>
  <div class="min-h-screen p-4 md:p-8" style="background: transparent;" ref="backgroundElement">
    <div class="max-w-6xl mx-auto overflow-hidden">
      <div
          class="mb-8 md:mb-12 rounded-xl p-6 shadow-lg transition-colors duration-300"
          :class="isDarkBackground ? 'bg-black/50' : 'bg-white/50'"
      >
        <h1
            class="text-4xl md:text-6xl font-bold"
            :class="isDarkBackground ? 'text-white' : 'text-gray-800'"
        >Experience</h1>
        <p
            class="text-lg mt-4"
            :class="isDarkBackground ? 'text-gray-300' : 'text-gray-600'"
        >Places I've worked and grown</p>
      </div>

      <div class="space-y-24" ref="experienceContainer">
        <div
            v-for="(experience, index) in displayedExperiences"
            :key="experience.id"
            class="relative"
        >
          <!-- Background number -->
          <div
              class="absolute -top-16 opacity-10 text-8xl md:text-9xl font-bold"
              :class="[
                index % 2 === 0 ? '-left-4' : '-right-4',
                isDarkBackground ? 'text-white' : 'text-gray-800'
              ]"
          >
            {{ experience.number }}
          </div>

          <!-- Content -->
          <div class="relative z-10">
            <!-- Header -->
            <div
                class="flex rounded-xl p-4 sm:p-6 items-center justify-between mb-6 shadow-lg transition-colors duration-300"
                :class="isDarkBackground ? 'bg-black/50' : 'bg-white/50'"
            >
              <div style="width: 50%">
                <h2
                    class="text-3xl font-bold"
                    :class="isDarkBackground ? 'text-white' : 'text-gray-800'"
                >{{ experience.company }}</h2>
                <h3
                    class="mt-1 text-sm"
                    :class="isDarkBackground ? 'text-gray-300' : 'text-gray-600'"
                >{{ $formatDate(experience.start,'{y}-{m}-{d}') }} - {{ experience.end ? $formatDate(experience.end,'{y}-{m}-{d}') : '至今' }}</h3>
                <p
                    class="leading-relaxed mb-6 text-sm md:text-base text-left"
                    :class="isDarkBackground ? 'text-gray-200' : 'text-gray-700'"
                    v-html="formattedContent(experience.description)"
                ></p>
              </div>
              <span
                  class="text-3xl"
                  :class="isDarkBackground ? 'text-white' : 'text-gray-800'"
                  style="width: 50%"
              >{{ experience.role }}</span>
            </div>
            <!-- Main content -->
            <div
                class="flex flex-col md:grid md:grid-cols-2 gap-6 items-end"
                :class="[
                index % 2 === 1 ? 'md:text-right' : '',
                'relative w-full'
              ]"
            >
              <!-- Featured Image - Full width on mobile -->
              <div
                  @click="toRoute('/about/experience/'+experience.id)"
                  class="order-1 md:order-none relative aspect-video md:aspect-square rounded-lg overflow-hidden mb-6 md:mb-0 w-full shadow-lg"
                  :class="[
                  index % 2 === 1 ? 'md:order-1' : '',
                ]"
              >
                <img
                    :src="$fileFull(experience.image)||defaultCover"
                    :alt="experience.company"
                    class="absolute inset-0 w-full h-full object-cover"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
              </div>

              <!-- Projects Section -->
              <div
                  :class="[
                  index % 2 === 1 ? 'md:order-2' : '',
                  'order-2 w-full h-full'
                ]"
              >
                <!-- Project Cards Grid - Always 2 per row -->
                <div class="grid grid-cols-2 gap-3 w-full [&>*:only-child]:col-span-2" v-if="experience.project">
                  <!-- First 3 Project Cards -->
                  <template v-for="(project, idx) in experience.project.slice(0, 3)" :key="project.id">
                    <div
                        @click="toRoute('/about/project/'+project.id)"
                        class="group relative rounded-lg overflow-hidden cursor-pointer w-full shadow-lg transition-colors duration-300"
                        :class="isDarkBackground ? 'bg-black/60' : 'bg-white/60'"
                    >
                      <div class="aspect-square">
                        <img
                            :src="$fileFull(project.photo)||defaultCover"
                            :alt="project.title"
                            class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                        />
                        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
                        <div class="absolute inset-0 p-3 flex flex-col justify-end">
                          <span class="text-sm font-medium text-white">{{ project.title }}</span>
                          <span class="text-xs text-white/80">{{ $formatDate(project.start,'{y}-{m}-{d}') }} - {{ project.end ? $formatDate(project.end,'{y}-{m}-{d}') : '至今' }}</span>
                        </div>
                      </div>
                    </div>
                  </template>

                  <!-- View More Card -->
                  <div
                      v-if="experience.project.length > 3"
                      class="group relative rounded-lg overflow-hidden cursor-pointer w-full shadow-lg transition-colors duration-300"
                      :class="isDarkBackground ? 'bg-black/60 border border-white/20 hover:border-white/40' : 'bg-white/60 border border-gray-300/50 hover:border-gray-400'"
                  >
                    <div class="aspect-square flex items-center justify-center" @click="toRoute('/about/project?experience_id='+experience.id)">
                      <div class="text-center">
                        <div class="mb-1">
                          <PlusCircleIcon
                              class="w-6 h-6 mx-auto transition-colors"
                              :class="isDarkBackground ? 'text-gray-300 group-hover:text-white' : 'text-gray-600 group-hover:text-gray-800'"
                          />
                        </div>
                        <span
                            class="text-xs font-medium transition-colors"
                            :class="isDarkBackground ? 'text-gray-300 group-hover:text-white' : 'text-gray-600 group-hover:text-gray-800'"
                        >
                          View {{ experience.project.length - 3 }} More
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div ref="observerTarget" class="h-10 mt-4"></div>
      <!-- Loading indicator -->
      <div v-if="loading" class="text-center py-8">
        <div
            class="inline-block animate-spin rounded-full h-8 w-8 border-t-2"
            :class="isDarkBackground ? 'border-white' : 'border-gray-800'"
        ></div>
      </div>
    </div>

    <!-- Theme Toggle Button -->
    <div class="fixed bottom-6 right-6">
      <button
          @click="toggleTheme"
          class="p-3 rounded-full shadow-lg transition-colors"
          :class="isDarkBackground ? 'bg-white text-gray-800' : 'bg-gray-800 text-white'"
      >
        <svg v-if="isDarkBackground" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
        <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import {userGet,projectList,experienceList} from "@/api/user";
import { ref, onMounted, onUnmounted } from 'vue'
import { PlusCircleIcon } from '@heroicons/vue/24/outline'

const displayedExperiences = ref([])
const loading = ref(false)
const current = ref(1)
const total = ref(0)
const pageSize = ref(16)
const experiencesMore = ref(false)
let observer = null;
const observerTarget = ref(null)
const hasMore = ref(true)
const params = ref({my:false,title:undefined,myCollect:false,sort:"",myThumb:false,description:""})

const backgroundElement = ref(null);
const isDarkBackground = ref(true);

// 检测背景颜色是否为深色
const detectBackgroundBrightness = () => {
  // 这里使用简单的方法，实际项目中可以使用更复杂的算法
  // 或者根据用户上传的背景图片的主色调来决定

  // 这里我们使用系统偏好作为初始值
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
  isDarkBackground.value = mediaQuery.matches;

  // 监听系统主题变化
  mediaQuery.addEventListener('change', (e) => {
    isDarkBackground.value = e.matches;
  });
};

// 切换主题
const toggleTheme = () => {
  isDarkBackground.value = !isDarkBackground.value;
};

const setupIntersectionObserver = () => {
  observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting && !loading.value && hasMore.value) {
      list();
    }
  }, {
    root: null,
    rootMargin: '0px',
    threshold: 0.1
  });

  if (observerTarget.value) {
    observer.observe(observerTarget.value);
  }
}

onMounted(()=>{
  setupIntersectionObserver();
  detectBackgroundBrightness();
})

onUnmounted(() => {
  if (observer) {
    observer.disconnect();
  }
})

const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
}

const list =()=>{
  if (loading.value || !hasMore.value) return;
  loading.value = true;
  experienceList(Object.assign({},params.value,{"page":{"pageNum":current.value,"pageSize":pageSize.value}})).then(res=>{
    if(res&&res.code===200){
      let data = res.data.list || []
      displayedExperiences.value = [...displayedExperiences.value, ...data];
      total.value = res.data.total||0
      hasMore.value = data.length === pageSize.value;
      current.value++
    }
  }).finally(()=>{
    loading.value = false
  })
}

import {useRouter} from "vue-router";
import defaultCover from "@/assets/image/default-cover.jpg";
const router = useRouter()
const toRoute = (path) => {
  router.push(path)
}
</script>

<style scoped>
/* Add these styles to ensure proper containment */
.max-w-6xl {
  max-width: 72rem;
  margin-left: auto;
  margin-right: auto;
  width: 100%;
}

/* Ensure grid items don't overflow */
.grid {
  width: 100%;
  max-width: 100%;
}
</style>

