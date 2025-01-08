<template>
  <div class="min-h-screen bg-black text-white p-8">
    <div class="max-w-7xl mx-auto">
      <div class="mb-12">
        <h1 class="text-5xl md:text-7xl font-bold">Experience</h1>
        <p class="text-xl text-gray-400 mt-4">Places I've worked and grown</p>
      </div>

      <div class="space-y-24">
        <div
            v-for="(experience, index) in experiences"
            :key="experience.id"
            class="relative"
            :class="index % 2 === 0 ? 'text-left' : 'text-right'"
        >
          <!-- Background number -->
          <div
              class="absolute -top-20 opacity-10 text-[12rem] font-bold"
              :class="index % 2 === 0 ? '-left-8' : '-right-8'"
          >
            {{ experience.number }}
          </div>

          <!-- Content -->
          <div class="relative z-10">
            <div class="flex items-center justify-between mb-8">
              <h2 class="text-4xl font-bold" :class="index % 2 === 1 ? 'order-2' : ''">
                {{ experience.company }}
              </h2>
              <span class="text-gray-400">{{ experience.period }}</span>
            </div>

            <div class="grid md:grid-cols-2 gap-8" :class="index % 2 === 1 ? 'md:text-right' : ''">
              <div :class="index % 2 === 1 ? 'md:order-2' : ''">
                <h3 class="text-2xl font-bold mb-4">{{ experience.role }}</h3>
                <p class="text-gray-400 leading-relaxed mb-6">{{ experience.description }}</p>

                <!-- Project Cards -->
                <div class="grid grid-cols-2 gap-4 mt-6">
                  <div
                      v-for="project in experience.projects"
                      :key="project.id"
                      class="group relative bg-gray-800 rounded-lg overflow-hidden cursor-pointer"
                  >
                    <div class="aspect-square">
                      <img
                          :src="project.image"
                          :alt="project.title"
                          class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                      />
                      <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
                      <div class="absolute inset-0 p-4 flex flex-col justify-end">
                        <span class="text-sm font-medium">{{ project.title }}</span>
                        <span class="text-xs text-gray-400">{{ project.date }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div
                  class="relative aspect-square rounded-lg overflow-hidden"
                  :class="index % 2 === 1 ? 'md:order-1' : ''"
              >
                <img
                    :src="experience.image"
                    :alt="experience.company"
                    class="absolute inset-0 w-full h-full object-cover"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const experiences = ref([
  {
    id: 1,
    number: '01',
    company: 'Aris Digital',
    role: 'Senior Developer',
    period: '2022 - Present',
    description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects.',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    projects: [
      {
        id: 1,
        title: 'E-commerce Platform',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 2,
        title: 'Banking App',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 3,
        title: 'CRM System',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 4,
        title: 'Analytics Dashboard',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      }
    ]
  },{
    id: 1,
    number: '01',
    company: 'Aris Digital',
    role: 'Senior Developer',
    period: '2022 - Present',
    description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects.',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    projects: [
      {
        id: 1,
        title: 'E-commerce Platform',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 2,
        title: 'Banking App',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 3,
        title: 'CRM System',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 4,
        title: 'Analytics Dashboard',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      }
    ]
  },{
    id: 1,
    number: '01',
    company: 'Aris Digital',
    role: 'Senior Developer',
    period: '2022 - Present',
    description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects.',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    projects: [
      {
        id: 1,
        title: 'E-commerce Platform',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 2,
        title: 'Banking App',
        date: '2023',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 3,
        title: 'CRM System',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 4,
        title: 'Analytics Dashboard',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      }
    ]
  },
  // ... more experiences
])
</script>