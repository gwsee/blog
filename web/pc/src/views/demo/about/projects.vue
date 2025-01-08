<template>
  <div class="min-h-screen bg-black text-white p-8">
    <div class="max-w-7xl mx-auto">
      <h1 class="text-5xl md:text-7xl font-bold mb-12">Projects</h1>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
            v-for="project in projects"
            :key="project.id"
            class="group relative aspect-[3/4] overflow-hidden rounded-lg bg-gray-900"
        >
          <img
              :src="project.image"
              :alt="project.title"
              class="absolute inset-0 w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>

          <div class="absolute inset-0 p-6 flex flex-col justify-between">
            <div class="flex justify-between items-start">
              <span class="text-5xl font-bold opacity-25">{{ project.number }}</span>
              <span class="text-sm bg-white/10 px-3 py-1 rounded-full">{{ project.date }}</span>
            </div>

            <div>
              <h3 class="text-2xl font-bold mb-2">{{ project.title }}</h3>
              <p class="text-gray-400 mb-4">{{ project.description }}</p>
              <button class="px-4 py-2 bg-white/10 rounded-full hover:bg-white/20 transition-colors">
                View Details
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const projects = ref([
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    date: '2024-01-08',
    description: 'Energy Management System with innovative solutions for sustainable future.'
  },
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    date: '2024-01-08',
    description: 'Energy Management System with innovative solutions for sustainable future.'
  },
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    date: '2024-01-08',
    description: 'Energy Management System with innovative solutions for sustainable future.'
  },
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    date: '2024-01-08',
    description: 'Energy Management System with innovative solutions for sustainable future.'
  },
  {
    id: 1,
    number: '01',
    title: 'Engie',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    date: '2024-01-08',
    description: 'Energy Management System with innovative solutions for sustainable future.'
  },
  // ... more projects
])
</script>