<template>
  <div
      class="min-h-screen bg-cover bg-center bg-fixed relative"
      :style="{ backgroundImage: `url(${backgroundImage})` }"
  >
    <!-- Dynamic background overlay -->
    <div
        class="absolute inset-0 bg-gradient-to-b"
        :class="[
        isLightBg ?
        'from-black/40 via-black/20 to-black/40' :
        'from-black/20 via-transparent to-black/20'
      ]"
    ></div>

    <!-- Main content -->
    <div class="relative z-10 min-h-screen">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <!-- Header section with enhanced readability -->
        <div class="space-y-4 mb-12">
          <h1 class="text-5xl md:text-7xl font-bold text-shadow-lg">
            Experience
          </h1>
          <p class="text-xl text-shadow-md">
            Places I've worked and grown
          </p>
        </div>

        <!-- Project title card with backdrop blur -->
        <div class="backdrop-blur-md bg-black/30 rounded-2xl p-8 mb-8 shadow-xl border border-white/10">
          <h2 class="text-4xl font-bold text-shadow-md mb-4">
            Enterprise E-commerce Platform
          </h2>
          <div class="flex items-center space-x-4 text-shadow-sm">
            <span class="flex items-center">
              <CalendarIcon class="w-5 h-5 mr-2" />
              2023-01-01 - 2023-12-31
            </span>
            <a href="#" class="text-blue-400 hover:text-blue-300 flex items-center">
              <LinkIcon class="w-5 h-5 mr-2" />
              项目地址
            </a>
          </div>
        </div>

        <!-- Technologies section -->
        <div class="mb-8">
          <h3 class="text-2xl font-bold mb-4 text-shadow-md flex items-center">
            <CodeBracketIcon class="w-6 h-6 mr-2" />
            Technologies & Tools
          </h3>
          <div class="backdrop-blur-md bg-black/30 rounded-xl p-6 border border-white/10">
            <div class="flex flex-wrap gap-2">
              <span
                  v-for="tech in technologies"
                  :key="tech"
                  class="px-3 py-1.5 rounded-full text-sm font-medium
                       bg-white/10 hover:bg-white/20 transition-colors
                       backdrop-blur-sm text-shadow-sm"
              >
                {{ tech }}
              </span>
            </div>
          </div>
        </div>

        <!-- Key Projects Grid -->
        <div class="grid md:grid-cols-2 gap-6">
          <div
              v-for="(project, index) in projects"
              :key="index"
              class="group relative rounded-xl overflow-hidden"
          >
            <img
                :src="project.image"
                :alt="project.title"
                class="w-full h-64 object-cover"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent
                        group-hover:from-black/90 transition-all duration-300">
              <div class="absolute bottom-0 p-6 w-full">
                <h4 class="text-xl font-bold mb-2 text-shadow-md">
                  {{ project.title }}
                </h4>
                <div class="flex flex-wrap gap-2">
                  <span
                      v-for="tech in project.technologies"
                      :key="tech"
                      class="px-2 py-1 rounded-full text-xs font-medium
                           bg-white/10 backdrop-blur-sm text-shadow-sm"
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
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
  CalendarIcon,
  LinkIcon,
  CodeBracketIcon
} from '@heroicons/vue/24/outline'

const backgroundImage = ref('https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=1080&width=1920')

// Determine if background is light (you may want to implement actual image analysis)
const isLightBg = computed(() => {
  return false // This should be dynamically determined based on background
})

const technologies = [
  'Vue.js',
  'TypeScript',
  'Tailwind CSS',
  'Node.js',
  'MongoDB',
  'Docker',
  'AWS',
  'Redis',
  'GraphQL'
]

const projects = [
  {
    title: 'Project 1',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=400&width=600',
    technologies: ['Vue.js', 'TypeScript']
  },
  {
    title: 'Project 2',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=400&width=600',
    technologies: ['Node.js', 'MongoDB']
  }
]
</script>

<style>
/* Text shadow utilities */
.text-shadow-sm {
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
}

.text-shadow-md {
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
}

.text-shadow-lg {
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.5);
}

/* Ensure text is always readable */
.text-shadow-contrast {
  text-shadow:
      -1px -1px 0 rgba(0,0,0,0.5),
      1px -1px 0 rgba(0,0,0,0.5),
      -1px 1px 0 rgba(0,0,0,0.5),
      1px 1px 0 rgba(0,0,0,0.5);
}
</style>
