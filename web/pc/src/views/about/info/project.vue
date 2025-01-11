<template>
  <!--bg-[url('https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=1080&width=1920')]-->
  <div class="min-h-screen  bg-cover bg-center bg-fixed">
    <div class="min-h-screen backdrop-blur-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <!-- 返回按钮 -->
        <button
            @click="$router.back()"
            class="mb-8 flex items-center space-x-2 text-shadow-lg hover:text-black transition-colors duration-300"
        >
          <ArrowLeftIcon class="w-5 h-5" />
          <span>返回列表</span>
        </button>

        <div class="grid lg:grid-cols-3 gap-8">
          <!-- 主要内容区域 -->
          <div class="lg:col-span-2 space-y-8">
            <!-- 项目标题区域 -->
            <div class="bg-gray-900/80 backdrop-blur-md rounded-2xl p-8 border border-gray-800">
              <h1 class="text-4xl font-bold text-white mb-4">{{ project.title }}</h1>
              <div class="flex flex-wrap gap-4 items-center text-sm text-gray-300">
                <div class="flex items-center space-x-2" v-if="false">
                  <CalendarIcon class="w-4 h-4" />
                  <span>  {{ $formatDate(project.start,'{y}-{m}') }} - {{ $formatDate(project.end,'{y}-{m}') }}</span>
                </div>
                <a
                    v-if="project.link"
                    :href="project.link"
                    target="_blank"
                    class="flex items-center space-x-2 text-blue-400 hover:text-blue-300"
                >
                  <LinkIcon class="w-4 h-4" />
                  <span>项目地址</span>
                </a>
              </div>
            </div>
            <!-- Technologies & Tools -->
            <div class="space-y-4" v-if=" project.skills&& project.skills.length>0">
              <h2 class="text-2xl font-bold text-white flex items-center space-x-2">
                <CodeBracketIcon class="w-6 h-6" />
                <span>Technologies & Tools</span>
              </h2>
              <div class="grid gap-4">
                <div class="bg-gray-900/80 backdrop-blur-md rounded-xl p-6 border border-gray-800">
                  <div class="flex flex-wrap gap-3">
                    <span
                        v-for="skill in project.skills"
                        :key="skill"
                        class="px-4 py-2 bg-gray-800/60 rounded-full text-sm text-gray-300 hover:bg-gray-700/60 transition-colors duration-300"
                    >
                      {{ skill }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
            <!-- Key Projects -->
            <div class="space-y-4" v-if=" project.photos&& project.photos.length>0">
              <h2 class="text-2xl font-bold text-white flex items-center space-x-2">
                <DocumentIcon class="w-6 h-6" />
                <span>Snapshot</span>
              </h2>
              <div class="grid sm:grid-cols-2 gap-4">
                <div
                    v-for="(image, index) in project.photos"
                    :key="index"
                    class="group relative overflow-hidden rounded-xl cursor-pointer"
                    @click="openImage(index)"
                >
                  <img
                      :src="$fileFull(image)"
                      :alt="`Project image ${index + 1}`"
                      class="w-full h-48 object-cover transform transition duration-500 group-hover:scale-110"
                  />
                  <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent">
                    <div class="absolute bottom-0 p-4 w-full">
                      <h3 class="text-lg font-semibold text-white mb-1" v-if="false">Project {{ index + 1 }}</h3>
                      <div class="flex flex-wrap gap-2"  v-if="false">
                        <span
                            v-for="tech in project.skills"
                            :key="tech"
                            class="px-2 py-1 bg-gray-800/60 backdrop-blur-sm rounded text-xs text-gray-300"
                        >
                          {{ tech }}
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 右侧信息 -->
          <div class="space-y-6">
            <!-- Key Achievements -->
            <div class="bg-gray-900/80 backdrop-blur-md rounded-2xl p-6 border border-gray-800">
              <h2 class="text-2xl font-bold text-white mb-6 flex items-center space-x-2">
                <TrophyIcon class="w-6 h-6" />
                <span>Key Achievements</span>
              </h2>
              <div class="space-y-6">
                <div class="space-y-2">
                  <h3 class="text-lg font-semibold text-white">项目描述</h3>
                  <p class="text-gray-300 leading-relaxed" v-html="formattedContent(project.description)"></p>
                </div>
                <div class="space-y-2">
                  <h3 class="text-lg font-semibold text-white">项目周期</h3>
                  <div class="flex justify-between text-gray-300">
                    <span>开始时间:  {{ $formatDate(project.start,'{y}-{m}') }}</span>
                    <span>结束时间:  {{ $formatDate(project.end,'{y}-{m}') }}</span>
                  </div>
                </div>
                <div class="space-y-2">
                  <h3 class="text-lg font-semibold text-white">项目状态</h3>
                  <div class="flex items-center space-x-2 text-gray-300">
                    <span
                        class="w-2 h-2 rounded-full"
                        :class="!project.end ? 'bg-green-500' : 'bg-gray-500'"
                    ></span>
                    <span>{{ !project.end  ? '进行中' : '已完成' }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 图片预览模态框 -->
    <div
        v-if="selectedImage !== null"
        class="fixed inset-0 bg-black/95 z-50 flex items-center justify-center p-4"
        @click="selectedImage = null"
    >
      <div class="relative max-w-5xl w-full">
        <img
            :src="project.images[selectedImage]"
            :alt="`Project image ${selectedImage + 1}`"
            class="w-full h-auto rounded-lg"
        />
        <button
            class="absolute top-4 right-4 text-white hover:text-gray-300 transition-colors duration-300"
            @click.stop="selectedImage = null"
        >
          <XMarkIcon class="w-6 h-6" />
        </button>
        <!-- 导航按钮 -->
        <div class="absolute inset-y-0 left-0 right-0 flex justify-between items-center px-4">
          <button
              v-if="selectedImage > 0"
              @click.stop="selectedImage--"
              class="p-2 rounded-full bg-black/50 hover:bg-black/75 transition-opacity duration-300"
          >
            <ChevronLeftIcon class="w-6 h-6" />
          </button>
          <button
              v-if="selectedImage < project.images.length - 1"
              @click.stop="selectedImage++"
              class="p-2 rounded-full bg-black/50 hover:bg-black/75 transition-opacity duration-300"
          >
            <ChevronRightIcon class="w-6 h-6" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'
import {
  ArrowLeftIcon,
  CalendarIcon,
  LinkIcon,
  DocumentIcon,
  CodeBracketIcon,
  TrophyIcon,
  XMarkIcon,
  ChevronLeftIcon,
  ChevronRightIcon
} from '@heroicons/vue/24/outline'
import {useRoute} from "vue-router";
import {projectGet} from "@/api/user";
const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
}
const project = ref({
  name: "Enterprise E-commerce Platform",
  description: `A comprehensive e-commerce platform rebuild focused on improving user experience and system performance.
Key achievements:
• Redesigned and implemented entire frontend architecture
• Improved performance by 50%
• Implemented responsive design
• Integrated new payment and logistics systems`,
  startDate: "2023-01-01",
  endDate: "2023-12-31",
  status: "active",
  url: "https://example.com/project",
  skills: [
    "Vue.js",
    "TypeScript",
    "Tailwind CSS",
    "Node.js",
    "MongoDB",
    "Docker",
    "AWS",
    "CI/CD",
    "Redis",
    "GraphQL"
  ],
  images: [
    "https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=800",
    "https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=800",
    "https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=800",
    "https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=800"
  ]
})
onMounted(()=>{
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  projectGet({id:id}).then((res) => {
    if(res){
      const data = res.data || {}
      console.log(data)
      project.value = data
    }
  })
})
// 示例数据

const selectedImage = ref(null)

const openImage = (index) => {
  selectedImage.value = index
}

</script>
