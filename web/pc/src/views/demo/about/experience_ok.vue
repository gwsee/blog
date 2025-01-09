<template>
  <div class="min-h-screen bg-black text-white p-4 md:p-8">
    <div class="max-w-6xl mx-auto overflow-hidden">
      <div class="mb-8 md:mb-12">
        <h1 class="text-4xl md:text-6xl font-bold">Experience</h1>
        <p class="text-lg text-gray-400 mt-4">Places I've worked and grown</p>
      </div>

      <div class="space-y-24">
        <div
            v-for="(experience, index) in experiences"
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
            <!-- Header -->
            <div class="flex items-center justify-between mb-6">
              <div :class="index % 2 === 1 ? 'order-2' : ''">
                <h2 class="text-3xl font-bold">{{ experience.company }}</h2>
                <h3 class="text-xl text-gray-400 mt-1">{{ experience.role }}</h3>
              </div>
              <span class="text-gray-400 text-sm">{{ experience.period }}</span>
            </div>

            <!-- Main content -->
            <div
                class="grid md:grid-cols-2 gap-6 items-end"
                :class="[
                index % 2 === 1 ? 'md:text-right' : '',
                'relative w-full'
              ]"
            >
              <!-- Left column -->
              <div
                  :class="[
                  index % 2 === 1 ? 'md:order-2' : '',
                  'w-full'
                ]"
              >
                <p class="text-gray-400 leading-relaxed mb-6 text-sm md:text-base">{{ experience.description }}</p>

                <!-- Project Cards Grid -->
                <div class="grid grid-cols-2 gap-3 w-full">
                  <!-- First 3 Project Cards -->
                  <template v-for="(project, idx) in experience.projects.slice(0, 3)" :key="project.id">
                    <div
                        class="group relative bg-gray-800 rounded-lg overflow-hidden cursor-pointer"
                        :class="[
                        idx === 2 ? 'col-span-2 md:col-span-1' : '',
                        'w-full'
                      ]"
                    >
                      <div class="aspect-square">
                        <img
                            :src="project.image"
                            :alt="project.title"
                            class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                        />
                        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
                        <div class="absolute inset-0 p-3 flex flex-col justify-end">
                          <span class="text-sm font-medium">{{ project.title }}</span>
                          <span class="text-xs text-gray-400">{{ project.date }}</span>
                        </div>
                      </div>
                    </div>
                  </template>

                  <!-- View More Card -->
                  <div
                      v-if="experience.projects.length > 3"
                      class="group relative bg-gray-800/50 rounded-lg overflow-hidden cursor-pointer
                           border border-gray-700 hover:border-gray-500 transition-colors w-full"
                  >
                    <div class="aspect-square flex items-center justify-center">
                      <div class="text-center">
                        <div class="mb-1">
                          <PlusCircleIcon class="w-6 h-6 mx-auto text-gray-400 group-hover:text-white transition-colors" />
                        </div>
                        <span class="text-xs font-medium text-gray-400 group-hover:text-white transition-colors">
                          View {{ experience.projects.length - 3 }} More
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Right column / Featured Image -->
              <div
                  class="relative aspect-video md:aspect-square rounded-lg overflow-hidden"
                  :class="[
                  index % 2 === 1 ? 'md:order-1' : '',
                  'w-full'
                ]"
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
// import { PlusCircleIcon } from '@heroicons/vue/24/outline'

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
  },
  {
    id: 2,
    number: '02',
    company: 'Mon Studio',
    role: 'Frontend Developer',
    period: '2020 - 2022',
    description: 'Developed innovative user interfaces and implemented responsive designs for various client projects.',
    image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
    projects: [
      {
        id: 1,
        title: 'Portfolio Website',
        date: '2022',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 2,
        title: 'Design System',
        date: '2021',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      },
      {
        id: 3,
        title: 'Marketing Site',
        date: '2021',
        image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
      }
    ]
  }
])
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
</style>