<template>
  <section class="relative py-12">
    <div class="container mx-auto px-4">
      <h2 class="text-4xl font-bold mb-12 text-white">Experience</h2>

      <div ref="scrollContainer" class="relative">
        <!-- Cards Container -->
        <div
            ref="cardsContainer"
            class="flex gap-6 transition-transform duration-300 ease-out"
            :style="{ transform: `translateX(${scrollPosition}px)` }"
            @mousedown="startDrag"
            @mousemove="onDrag"
            @mouseup="stopDrag"
            @mouseleave="stopDrag"
            @touchstart="startTouch"
            @touchmove="onTouch"
            @touchend="stopTouch"
        >
          <div
              v-for="project in projects"
              :key="project.id"
              class="relative flex-shrink-0 w-[300px] md:w-[400px] aspect-[3/4] rounded-lg overflow-hidden group"
          >
            <!-- Background Image -->
            <img
                :src="project.image"
                :alt="project.title"
                class="absolute inset-0 w-full h-full object-cover"
            />

            <!-- Colored Overlay -->
            <div
                class="absolute inset-0 opacity-90 transition-opacity duration-300 group-hover:opacity-75"
                :class="project.color"
            ></div>

            <!-- Content -->
            <div class="relative h-full p-8 flex flex-col justify-between text-white">
              <!-- Top Content -->
              <div>
                <span class="text-6xl font-bold opacity-50">{{ project.number }}</span>
                <div class="mt-2 text-sm opacity-75">{{ project.date }}</div>
              </div>

              <!-- Bottom Content -->
              <div>
                <h3 class="text-2xl font-bold mb-2">{{ project.title }}</h3>
                <h4 class="text-lg mb-2">{{ project.role }}</h4>
                <p class="text-sm opacity-75">{{ project.description }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Navigation Buttons -->
        <button
            v-show="canScrollLeft"
            @click="scroll('left')"
            class="absolute left-2 top-1/2 -translate-y-1/2 w-10 h-10 flex items-center justify-center
                 bg-black/30 backdrop-blur-sm rounded-full hover:bg-black/50 transition-colors z-10"
        >
          <ChevronLeftIcon class="w-6 h-6 text-white" />
        </button>

        <button
            v-show="canScrollRight"
            @click="scroll('right')"
            class="absolute right-2 top-1/2 -translate-y-1/2 w-10 h-10 flex items-center justify-center
                 bg-black/30 backdrop-blur-sm rounded-full hover:bg-black/50 transition-colors z-10"
        >
          <ChevronRightIcon class="w-6 h-6 text-white" />
        </button>

        <!-- View All Projects Button -->
        <div class="absolute -bottom-12 left-1/2 -translate-x-1/2">
          <router-link
              to="/projects"
              class="inline-flex items-center gap-2 px-6 py-2 bg-white/10 backdrop-blur-sm rounded-full
                   hover:bg-white/20 transition-all duration-300 text-white"
          >
            View All Projects
            <ArrowRightIcon class="w-5 h-5" />
          </router-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ChevronLeftIcon, ChevronRightIcon, ArrowRightIcon } from '@heroicons/vue/24/outline'

const scrollContainer = ref(null)
const cardsContainer = ref(null)
const scrollPosition = ref(0)
const isDragging = ref(false)
const startX = ref(0)
const currentX = ref(0)

// Scroll limits
const maxScroll = computed(() => {
  if (!cardsContainer.value || !scrollContainer.value) return 0
  return -(cardsContainer.value.scrollWidth - scrollContainer.value.clientWidth)
})

// Scroll indicators
const canScrollLeft = computed(() => scrollPosition.value < 0)
const canScrollRight = computed(() => scrollPosition.value > maxScroll.value)

// Mouse handlers
const startDrag = (e) => {
  isDragging.value = true
  startX.value = e.pageX - scrollPosition.value
  currentX.value = scrollPosition.value
  document.body.style.cursor = 'grabbing'
}

const onDrag = (e) => {
  if (!isDragging.value) return
  e.preventDefault()
  const x = e.pageX - startX.value
  scrollPosition.value = Math.max(maxScroll.value, Math.min(0, x))
}

const stopDrag = () => {
  isDragging.value = false
  document.body.style.cursor = 'default'
}

// Touch handlers
const startTouch = (e) => {
  isDragging.value = true
  startX.value = e.touches[0].pageX - scrollPosition.value
  currentX.value = scrollPosition.value
}

const onTouch = (e) => {
  if (!isDragging.value) return
  e.preventDefault()
  const x = e.touches[0].pageX - startX.value
  scrollPosition.value = Math.max(maxScroll.value, Math.min(0, x))
}

const stopTouch = () => {
  isDragging.value = false
}

// Button scroll
const scroll = (direction) => {
  const amount = direction === 'left' ? 400 : -400
  const newPosition = scrollPosition.value + amount
  scrollPosition.value = Math.max(maxScroll.value, Math.min(0, newPosition))
}

// Clean up
onUnmounted(() => {
  document.body.style.cursor = 'default'
})

// Sample data
const projects = ref([
  {
    id: 1,
    number: '01',
    title: 'Engie',
    role: 'Energy Management',
    date: '2024-01-08',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'bg-gradient-to-b from-blue-600 to-blue-900',
    description: 'Energy Management System'
  },
  {
    id: 2,
    number: '02',
    title: 'CIC',
    role: 'Banking Solutions',
    date: '2024-01-07',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'bg-gradient-to-b from-teal-600 to-teal-900',
    description: 'Banking Solutions Platform'
  },
  {
    id: 3,
    number: '03',
    title: 'CGI',
    role: 'Digital Transform',
    date: '2024-01-06',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'bg-gradient-to-b from-purple-600 to-purple-900',
    description: 'Digital Transformation'
  },
  {
    id: 4,
    number: '04',
    title: 'Art-Spire',
    role: 'Creative Design',
    date: '2024-01-05',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'bg-gradient-to-b from-pink-600 to-pink-900',
    description: 'Creative Design Platform'
  },
  {
    id: 5,
    number: '05',
    title: 'Tokkun',
    role: 'Learning System',
    date: '2024-01-04',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    color: 'bg-gradient-to-b from-orange-600 to-orange-900',
    description: 'Learning Management System'
  }
])
</script>

<style scoped>
.container {
  max-width: 1400px;
}
</style>