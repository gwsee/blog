<template>
  <div class="min-h-screen bg-gradient-to-br ">
    <!-- 头部展示区 -->
    <section class="container mx-auto px-4 py-12 md:py-24">
      <div class="flex flex-col md:flex-row items-center gap-8">
        <div class="relative w-32 h-32 md:w-48 md:h-48">
          <img
              :src="userInfo.avatar?($fileFull(userInfo.avatar)):avatar"
              alt="Profile"
              class="rounded-full w-full h-full object-cover shadow-xl"
          />
        </div>
        <div class="text-center md:text-left">
          <h1 class="text-4xl md:text-6xl font-bold text-gray-800 dark:text-white mb-2">
            {{ userInfo.name || '天天开心' }}
          </h1>
          <h2 class="text-xl md:text-2xl text-gray-600 dark:text-gray-400 mb-6">
            {{
              userInfo.description || '我稍微偷点懒哈!'
            }}
          </h2>
          <div class="flex flex-wrap gap-4 justify-center md:justify-start">
            <button class="px-6 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors">
              View Projects
            </button>
            <button class="px-6 py-2 bg-white text-gray-800 rounded-md hover:bg-gray-100 transition-colors">
              Contact Me
            </button>
            <button class="px-6 py-2 bg-white text-gray-800 rounded-md hover:bg-gray-100 transition-colors" @click="toRoute('/about/profile/manage')">
              Manage
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- 项目展示区 -->
    <section class="text-white py-12">
      <nav class="container mx-auto px-4 mb-12" v-if="false">
        <ul class="flex justify-center md:justify-start gap-8 text-sm uppercase tracking-wider">
          <li class="hover:text-gray-400 cursor-pointer">Hi</li>
          <li class="hover:text-gray-400 cursor-pointer">Work</li>
          <li class="hover:text-gray-400 cursor-pointer">Profile</li>
          <li class="hover:text-gray-400 cursor-pointer">Contact</li>
        </ul>
      </nav>

      <div class="relative overflow-hidden">
        <div class="container mx-auto px-4">
          <div class="flex gap-6 overflow-x-auto pb-8 snap-x snap-mandatory scrollbar-hide">
            <div
                v-for="(experience,key) in experiences"
                :key="experience.id"
                class="relative flex-shrink-0 w-[280px] h-[400px] md:w-[320px] md:h-[500px]
                     rounded-lg overflow-hidden snap-start cursor-pointer group"
                @mouseenter="experience.isHovered = true"
                @mouseleave="experience.isHovered = false"
            >
              <img
                  :src="experience.image||'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400'"
                  :alt="experience.company"
                  class="absolute inset-0 w-full h-full object-cover"
              />
              <div :class="`absolute inset-0 bg-gradient-to-b ${colors[key%5]} opacity-75`"></div>
              <div class="absolute inset-0 p-6 flex flex-col justify-between transform transition-transform duration-500 group-hover:scale-105">
                <div>
                  <span class="text-4xl font-bold opacity-50">{{ experience.company }}</span>
                  <div class="mt-2 text-sm opacity-75">{{ $formatDate(experience.start,'{y}-{m}-{d}')  }} - {{experience.end?$formatDate(experience.end,'{y}-{m}-{d}'):'至今'}}</div>
                </div>
                <div>
                  <h3 class="text-2xl font-bold mb-2">{{ experience.role }}</h3>
                  <p class="text-sm opacity-75">{{ experience.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
      <div class=" bottom-0 left-0 w-full flex justify-center pb-8" v-if="experiencesMore">
        <router-link
            to="/experience"
            class="px-8 py-3 bg-white/10 backdrop-blur-sm rounded-full
                 hover:bg-white/20 transition-all duration-300
                 flex items-center gap-2 group"
        >
          View All Projects
          <!--          <ArrowRightIcon class="w-5 h-5 transform group-hover:translate-x-1 transition-transform" />-->
        </router-link>
      </div>
    </section>
  </div>
</template>

<script setup>
import {userGet,projectList,experienceList} from "@/api/user";
const colors = ref(['from-blue-600 to-blue-900','from-teal-600 to-teal-900',
  'from-red-600 to-red-900','from-pink-600 to-pink-900',
  'from-orange-600 to-orange-900'])
const projects = ref([
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-blue-600 to-blue-900',
    date: '2024-01-08',
    description: 'Energy Management System',
    isHovered: false
  },
  {
    id: 2,
    number: '02',
    title: 'CIC',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-teal-600 to-teal-900',
    date: '2024-01-07',
    description: 'Banking Solutions',
    isHovered: false
  },
  {
    id: 3,
    number: '03',
    title: 'CGI',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-red-600 to-red-900',
    date: '2024-01-06',
    description: 'Digital Transformation',
    isHovered: false
  },
  {
    id: 4,
    number: '04',
    title: 'Art-Spire',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-pink-600 to-pink-900',
    date: '2024-01-05',
    description: 'Creative Design Platform',
    isHovered: false
  },
  {
    id: 5,
    number: '05',
    title: 'Tokkun',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-orange-600 to-orange-900',
    date: '2024-01-04',
    description: 'Learning Management System',
    isHovered: false
  }
])
import {ref,onMounted} from 'vue'
import {useRouter} from 'vue-router'
import avatar from "@/assets/image/default-avatar.png"
const router = useRouter()
import {
  UserOutlined,
  MailOutlined,
  CodeOutlined,
  AppstoreOutlined,
} from '@ant-design/icons-vue'

// const projects = ref([
//   // {
//   //   id: 1,
//   //   title: 'E-commerce App',
//   //   image: '/src/assets/image/bg.jpg?height=400&width=600',
//   // },
//   // {
//   //   id: 2,
//   //   title: 'Social Media Platform',
//   //   image: '/src/assets/image/bg.jpg?height=400&width=600',
//   // },
//   // {
//   //   id: 3,
//   //   title: 'Fitness Tracker',
//   //   image: '/src/assets/image/bg.jpg?height=400&width=600',
//   // },
//   // {
//   //   id: 4,
//   //   title: 'Task Management',
//   //   image: '/src/assets/image/bg.jpg?height=400&width=600',
//   // },
// ])
const projectsMore = ref(false)
const experiences = ref([

])
const experiencesMore = ref(false)
const userInfo = ref({})

onMounted(()=>{
  userGet({}).then(res=>{
    if(res&&res.code===200){
      userInfo.value =res.data
    }
  })
  projectList({}).then(res=>{
    if(res&&res.code===200){
      let data =res.data
      // projects.value = data.list || []
      projectsMore.value = projects.value.length<(data.total || 0)
    }
  })
  experienceList({}).then(res=>{
    if(res&&res.code===200){
      let data =res.data
      experiences.value = data.list || []
      experiencesMore.value = experiences.value.length<(data.total || 0)
    }
  })
})

const toRoute = (path) => {
  router.push(path)
}
</script>

<style>
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}

@keyframes tilt {
  0%, 100% { transform: rotate3d(0); }
  25% { transform: rotate3d(1, 1, 0, 15deg); }
  75% { transform: rotate3d(-1, 1, 0, 15deg); }
}

.hover\:tilt:hover {
  animation: tilt 5s infinite;
}
</style>
