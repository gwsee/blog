<template>
  <div class="min-h-screen bg-gray-900">
    <PhotoSphere :photos="photos" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import PhotoSphere from './photosphere_ok.vue'

// Sample photo data
const photos = ref([
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Mountain Landscape',
    description: 'Beautiful mountain vista at sunset'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Ocean View',
    description: 'Peaceful ocean waves at dawn'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'City Skyline',
    description: 'Modern city skyline at night'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Forest Path',
    description: 'Serene forest walking path'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Desert Dunes',
    description: 'Rolling sand dunes at sunrise'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Snow Peak',
    description: 'Snow-capped mountain peak'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Autumn Colors',
    description: 'Vibrant autumn forest colors'
  },
  {
    url: '/placeholder.svg?height=400&width=400',
    title: 'Tropical Beach',
    description: 'Paradise tropical beach scene'
  }
])
</script>
