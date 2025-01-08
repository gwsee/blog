<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-900 to-black text-white">
    <!-- Hero Section -->
    <section class="container mx-auto px-4 py-12 md:py-24">
      <div class="flex flex-col md:flex-row items-center gap-8">
        <div class="relative w-32 h-32 md:w-48 md:h-48">
          <img
              src="https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=200&width=200"
              alt="Profile"
              class="rounded-full w-full h-full object-cover shadow-xl"
          />
        </div>
        <div class="text-center md:text-left">
          <h1 class="text-4xl md:text-6xl font-bold mb-2">
            to be or not to be
          </h1>
          <h2 class="text-xl md:text-2xl text-gray-400 mb-6">
            HEIHEI
          </h2>
          <div class="flex flex-wrap gap-4 justify-center md:justify-start">
            <button class="px-6 py-2 bg-teal-500 text-white rounded-md hover:bg-teal-600 transition-colors">
              View Projects
            </button>
            <button class="px-6 py-2 bg-pink-500 text-white rounded-md hover:bg-pink-600 transition-colors">
              Contact Me
            </button>
            <button class="px-6 py-2 border border-white/20 text-white rounded-md hover:bg-white/10 transition-colors">
              Manage
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- Navigation -->
    <nav class="container mx-auto px-4 mb-12">
      <ul class="flex justify-center md:justify-start gap-8 text-sm uppercase tracking-wider">
        <li class="hover:text-teal-400 cursor-pointer transition-colors">Hi</li>
        <li class="hover:text-teal-400 cursor-pointer transition-colors">Work</li>
        <li class="hover:text-teal-400 cursor-pointer transition-colors">Profile</li>
        <li class="hover:text-teal-400 cursor-pointer transition-colors">Contact</li>
      </ul>
    </nav>

    <!-- Projects Section -->
    <section class="relative py-12">
      <div ref="scrollContainer" class="project-scroll-container" @mousedown="startDrag" @mousemove="onDrag" @mouseup="stopDrag" @mouseleave="stopDrag">
        <div class="flex gap-6 px-4 md:px-[10vw]" :style="{ transform: `translateX(${scrollPosition}px)` }">
          <div
              v-for="project in projects"
              :key="project.id"
              class="relative flex-shrink-0 w-[280px] h-[400px] md:w-[320px] md:h-[500px]
                   rounded-lg overflow-hidden cursor-pointer group"
          >
            <img
                :src="project.image"
                :alt="project.title"
                class="absolute inset-0 w-full h-full object-cover"
            />
            <div :class="`absolute inset-0 bg-gradient-to-b ${project.color} opacity-75`"></div>
            <div class="absolute inset-0 p-6 flex flex-col justify-between transform transition-all duration-500 group-hover:scale-105">
              <div>
                <span class="text-4xl font-bold opacity-50">{{ project.number }}</span>
                <div class="mt-2 text-sm opacity-75">{{ project.date }}</div>
              </div>
              <div>
                <h3 class="text-2xl font-bold mb-2">{{ project.title }}</h3>
                <p class="text-sm opacity-75">{{ project.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- View All Projects Button -->
      <div class="absolute bottom-0 left-0 w-full flex justify-center pb-8">
        <router-link
            to="/projects"
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
import { ref, onMounted, onUnmounted } from 'vue'
// import { ArrowRightIcon } from '@heroicons/vue/24/outline'

const projects = ref([
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-teal-600 to-teal-900',
    date: '2024-01-08',
    description: 'Energy Management System'
  },
  {
    id: 2,
    number: '02',
    title: 'CIC',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-pink-600 to-pink-900',
    date: '2024-01-07',
    description: 'Banking Solutions'
  },
  {
    id: 3,
    number: '03',
    title: 'CGI',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-purple-600 to-purple-900',
    date: '2024-01-06',
    description: 'Digital Transformation'
  },
  {
    id: 4,
    number: '04',
    title: 'Art-Spire',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-rose-600 to-rose-900',
    date: '2024-01-05',
    description: 'Creative Design Platform'
  },
  {
    id: 5,
    number: '05',
    title: 'Tokkun',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'from-amber-600 to-amber-900',
    date: '2024-01-04',
    description: 'Learning Management System'
  }
])

// Drag to scroll functionality
const scrollContainer = ref(null)
const scrollPosition = ref(0)
const isDragging = ref(false)
const startX = ref(0)
const scrollLeft = ref(0)

const startDrag = (e) => {
  isDragging.value = true
  startX.value = e.pageX - scrollPosition.value
  document.body.style.cursor = 'grabbing'
}

const stopDrag = () => {
  isDragging.value = false
  document.body.style.cursor = 'default'
}

const onDrag = (e) => {
  if (!isDragging.value) return
  e.preventDefault()
  const x = e.pageX - startX.value
  const walk = x
  scrollPosition.value = Math.min(0, Math.max(walk, -1500)) // Adjust max scroll value based on content
}

// Clean up
onUnmounted(() => {
  document.body.style.cursor = 'default'
})
</script>

<style>
.project-scroll-container {
  overflow: hidden;
  cursor: grab;
}

.project-scroll-container:active {
  cursor: grabbing;
}

/* Hide scrollbar */
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>