<template>
  <div class="min-h-screen bg-black text-white">
    <!-- Hero Section -->
    <div class="relative h-[60vh] overflow-hidden">
      <img
          :src="experience.image"
          :alt="experience.company"
          class="absolute inset-0 w-full h-full object-cover"
      />
      <div class="absolute inset-0 bg-gradient-to-t from-black via-black/60 to-transparent"></div>

      <!-- Back Button -->
      <div class="absolute top-8 left-8 z-10">
        <button
            @click="$router.back()"
            class="flex items-center gap-2 text-white/80 hover:text-white transition-colors"
        >
          <ArrowLeftIcon class="w-5 h-5" />
          <span>Back</span>
        </button>
      </div>

      <!-- Content -->
      <div class="absolute bottom-0 left-0 w-full p-8 md:p-16">
        <div class="max-w-6xl mx-auto">
          <div class="flex flex-col md:flex-row items-end justify-between gap-6">
            <div>
              <span class="text-sm text-gray-400">{{ experience.period }}</span>
              <h1 class="text-4xl md:text-6xl font-bold mt-2">{{ experience.company }}</h1>
              <h2 class="text-xl md:text-2xl text-gray-300 mt-2">{{ experience.role }}</h2>
            </div>
            <div class="flex gap-4">
              <a
                  v-if="experience.website"
                  :href="experience.website"
                  target="_blank"
                  class="px-6 py-2 bg-white/10 backdrop-blur rounded-full hover:bg-white/20 transition-colors"
              >
                Visit Website
              </a>
              <button
                  class="px-6 py-2 bg-white/10 backdrop-blur rounded-full hover:bg-white/20 transition-colors"
              >
                Share
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-6xl mx-auto px-8 py-16">
      <div class="grid md:grid-cols-3 gap-16">
        <!-- Left Column - Details -->
        <div class="md:col-span-2 space-y-12">
          <!-- Overview -->
          <section>
            <h3 class="text-2xl font-bold mb-6">Overview</h3>
            <div class="prose prose-invert">
              <p class="text-gray-300 leading-relaxed">{{ experience.description }}</p>
              <ul class="mt-6 space-y-2 text-gray-300">
                <li v-for="point in experience.highlights" :key="point">
                  • {{ point }}
                </li>
              </ul>
            </div>
          </section>

          <!-- Projects -->
          <section>
            <h3 class="text-2xl font-bold mb-6">Key Projects</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div
                  v-for="project in experience.projects"
                  :key="project.id"
                  class="group relative bg-gray-800 rounded-lg overflow-hidden cursor-pointer"
              >
                <div class="aspect-video">
                  <img
                      :src="project.image"
                      :alt="project.title"
                      class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                  />
                  <div class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/50 to-transparent"></div>
                  <div class="absolute inset-0 p-6 flex flex-col justify-end">
                    <span class="text-lg font-medium mb-2">{{ project.title }}</span>
                    <p class="text-sm text-gray-400 mb-4">{{ project.description }}</p>
                    <div class="flex flex-wrap gap-2">
                      <span
                          v-for="tech in project.technologies"
                          :key="tech"
                          class="px-3 py-1 text-xs bg-white/10 rounded-full"
                      >
                        {{ tech }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <!-- Technologies -->
          <section>
            <h3 class="text-2xl font-bold mb-6">Technologies & Tools</h3>
            <div class="flex flex-wrap gap-3">
              <div
                  v-for="tech in experience.technologies"
                  :key="tech.name"
                  class="group flex items-center gap-2 px-4 py-2 bg-gray-800 rounded-lg hover:bg-gray-700 transition-colors"
              >
                <component
                    :is="tech.icon"
                    class="w-5 h-5 text-gray-400 group-hover:text-white transition-colors"
                />
                <span>{{ tech.name }}</span>
              </div>
            </div>
          </section>
        </div>

        <!-- Right Column - Additional Info -->
        <div class="space-y-8">
          <!-- Quick Facts -->
          <div class="bg-gray-800/50 rounded-xl p-6">
            <h4 class="text-lg font-medium mb-4">Quick Facts</h4>
            <div class="space-y-4">
              <div>
                <span class="text-sm text-gray-400">Location</span>
                <p>{{ experience.location }}</p>
              </div>
              <div>
                <span class="text-sm text-gray-400">Team Size</span>
                <p>{{ experience.teamSize }}</p>
              </div>
              <div>
                <span class="text-sm text-gray-400">Industry</span>
                <p>{{ experience.industry }}</p>
              </div>
            </div>
          </div>

          <!-- Achievements -->
          <div class="bg-gray-800/50 rounded-xl p-6">
            <h4 class="text-lg font-medium mb-4">Key Achievements</h4>
            <div class="space-y-4">
              <div
                  v-for="achievement in experience.achievements"
                  :key="achievement.title"
                  class="flex gap-4"
              >
                <div class="flex-shrink-0">
                  <component
                      :is="achievement.icon"
                      class="w-5 h-5 text-gray-400"
                  />
                </div>
                <div>
                  <p class="font-medium">{{ achievement.title }}</p>
                  <p class="text-sm text-gray-400">{{ achievement.description }}</p>
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
import { ref } from 'vue'
import {
  ArrowLeftIcon,
  CodeBracketIcon,
  CommandLineIcon,
  CpuChipIcon,
  ServerIcon,
  TrophyIcon,
  UserGroupIcon,
  ChartBarIcon
} from '@heroicons/vue/24/outline'

// 模拟数据
const experience = ref({
  id: 1,
  company: 'Aris Digital',
  role: 'Senior Developer',
  period: '2022 - Present',
  location: 'Shanghai, China',
  teamSize: '15-20 members',
  industry: 'Technology & Software',
  website: 'https://example.com',
  image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects. Worked on various challenging projects that pushed the boundaries of web technology.',
  highlights: [
    'Led a team of 5 developers in building a new e-commerce platform',
    'Implemented CI/CD pipelines reducing deployment time by 60%',
    'Mentored 3 junior developers who were promoted to mid-level positions',
    'Introduced new testing practices that increased code coverage by 40%'
  ],
  projects: [
    {
      id: 1,
      title: 'E-commerce Platform',
      date: '2023',
      image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300',
      description: 'A modern e-commerce platform with real-time inventory management',
      technologies: ['Vue.js', 'Node.js', 'MongoDB']
    },
    {
      id: 2,
      title: 'Banking App',
      date: '2023',
      image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300',
      description: 'Secure banking application with real-time transactions',
      technologies: ['React', 'TypeScript', 'AWS']
    },
    {
      id: 3,
      title: 'CRM System',
      date: '2022',
      image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300',
      description: 'Customer relationship management system with analytics',
      technologies: ['Vue.js', 'Python', 'PostgreSQL']
    }
  ],
  technologies: [
    { name: 'Frontend Development', icon: CodeBracketIcon },
    { name: 'Backend Systems', icon: ServerIcon },
    { name: 'DevOps', icon: CommandLineIcon },
    { name: 'Microservices', icon: CpuChipIcon }
  ],
  achievements: [
    {
      icon: TrophyIcon,
      title: 'Best Performing Team 2023',
      description: 'Led team to achieve highest customer satisfaction scores'
    },
    {
      icon: UserGroupIcon,
      title: 'Mentorship Excellence',
      description: 'Successfully mentored 3 junior developers to promotion'
    },
    {
      icon: ChartBarIcon,
      title: 'Performance Improvement',
      description: 'Reduced system response time by 40%'
    }
  ]
})
</script>

<style>
.prose {
  max-width: 65ch;
}
</style>