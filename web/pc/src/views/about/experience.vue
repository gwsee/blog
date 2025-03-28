<template>
  <div class="min-h-screen text-white p-4 md:p-8" style="background: transparent;">
    <div class="max-w-6xl mx-auto overflow-hidden">
      <div class="mb-8 md:mb-12 backdrop-blur-md  rounded-xl p-6 shadow-lg">
        <h1 class="text-4xl md:text-6xl font-bold">Experience</h1>
        <p class="text-lg text-white/70 mt-4">Places I've worked and grown</p>
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
              :class="index % 2 === 0 ? '-left-4' : '-right-4'"
          >
            {{ experience.number }}
          </div>

          <!-- Content -->
          <div class="relative z-10">
            <!-- 经验详情头部 - 响应式布局 -->
            <div class="backdrop-blur-md bg-black/30 border border-white/20 rounded-xl p-4 sm:p-6 mb-6 shadow-lg">
              <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
                <div class="w-full md:w-1/2">
                  <h2 class="text-3xl font-bold text-white">{{ experience.company }}</h2>
                  <h3 class="text-white/70 mt-1 text-sm">{{ $formatDate(experience.start,'{y}-{m}-{d}') }} - {{ experience.end ? $formatDate(experience.end,'{y}-{m}-{d}') : '至今' }}</h3>
                  <p class="text-white/80 leading-relaxed mt-2 mb-4 md:mb-0 text-sm md:text-base text-left" v-html="formattedContent(experience.description)"></p>
                </div>
                <div class="w-full md:w-1/2 text-center md:text-center">
                  <span class="text-2xl md:text-3xl text-white/80 font-semibold">{{ experience.role }}</span>
                </div>
              </div>
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
                        class="group relative backdrop-blur-md bg-black/30 border border-white/20 rounded-lg overflow-hidden cursor-pointer w-full shadow-lg"
                    >
                      <div class="aspect-square">
                        <img
                            :src="$fileFull(project.photo)||defaultCover"
                            :alt="project.title"
                            class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                        />
                        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
                        <div class="absolute inset-0 p-3 flex flex-col justify-end">
                          <span class="text-sm font-medium">{{ project.title }}</span>
                          <span class="text-xs text-white/70">{{ $formatDate(project.start,'{y}-{m}-{d}') }} - {{ project.end ? $formatDate(project.end,'{y}-{m}-{d}') : '至今' }}</span>
                        </div>
                      </div>
                    </div>
                  </template>

                  <!-- View More Card -->
                  <div
                      v-if="experience.project.length > 3"
                      class="group relative backdrop-blur-md bg-black/30 border border-white/20 rounded-lg overflow-hidden cursor-pointer
                           hover:border-white/40 transition-colors w-full shadow-lg"
                  >
                    <div class="aspect-square flex items-center justify-center" @click="toRoute('/about/project?experience_id='+experience.id)">
                      <div class="text-center">
                        <div class="mb-1">
                          <PlusCircleIcon class="w-6 h-6 mx-auto text-white/70 group-hover:text-white transition-colors" />
                        </div>
                        <span class="text-xs font-medium text-white/70 group-hover:text-white transition-colors">
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
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-t-2 border-white"></div>
      </div>
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

/* 添加额外的玻璃态效果样式 */
.backdrop-blur-md {
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
}
</style>

