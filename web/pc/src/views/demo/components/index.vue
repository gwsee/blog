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
    <section class="relative py-12 overflow-hidden">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold mb-8">Experience</h2>
        <div class="flex overflow-x-auto pb-8 -mx-4 px-4 scrollbar-hide">
          <div class="flex gap-6 flex-nowrap">
            <ExperienceCard
                v-for="experience in experiences"
                :key="experience.id"
                v-bind="experience"
            />
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
          <ArrowRightIcon class="w-5 h-5 transform group-hover:translate-x-1 transition-transform" />
        </router-link>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ArrowRightIcon } from '@heroicons/vue/24/outline'
import ExperienceCard from './components/ExperienceCard.vue'
const experiences = ref([
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
// const experiences = ref([
//   {
//     id: 1,
//     number: '01',
//     company: 'Aris Digital',
//     role: 'Senior Developer',
//     period: '2022 - Present',
//     image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
//     color: 'from-teal-600 to-teal-900',
//     description: 'Led development of enterprise-scale applications, mentored junior developers.'
//   },
//   {
//     id: 2,
//     number: '02',
//     company: 'Mon Studio',
//     role: 'Frontend Developer',
//     period: '2020 - 2022',
//     image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
//     color: 'from-pink-600 to-pink-900',
//     description: 'Developed innovative user interfaces and implemented responsive designs.'
//   },
//   // Add more experiences as needed
// ])
</script>

<style>
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>

