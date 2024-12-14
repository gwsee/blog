<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-gray-100">
    <!-- Hero Section -->
    <section class="relative py-20 overflow-hidden">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row items-center gap-12">
          <!-- Profile Image -->
          <div class="md:w-1/3 relative">
            <div class="w-64 h-64 rounded-full overflow-hidden border-8 border-white shadow-2xl">
              <img
                  src="/src/assets/image/bg.jpg?height=300&width=300"
                  alt="Profile"
                  class="object-cover w-full h-full"
              />
            </div>
            <div class="absolute -bottom-4 -right-4 bg-blue-500 text-white rounded-full p-4">
              <CodeOutlined style="font-size: 2rem;" />
            </div>
          </div>
          <!-- Text Content -->
          <div class="md:w-2/3">
            <h1 class="text-4xl md:text-6xl font-bold mb-6 text-gray-800">
              John Doe
            </h1>
            <p class="text-xl text-gray-600 mb-8 leading-relaxed">
              Passionate Mobile App Developer with 5+ years of experience in creating innovative solutions.
              Specialized in React Native and iOS development, with a keen eye for user experience and performance optimization.
            </p>
            <div class="flex flex-wrap gap-4">
              <a-button type="primary" size="large" shape="round">
                <template #icon><UserOutlined /></template>
                View Profile
              </a-button>
              <a-button size="large" shape="round">
                <template #icon><MailOutlined /></template>
                Contact Me
              </a-button>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Projects Timeline -->
    <section class="py-20 bg-white">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold mb-12 text-center">Featured Projects</h2>
        <a-timeline mode="alternate">
          <a-timeline-item v-for="project in projects" :key="project.id">
            <template #dot>
              <AppstoreOutlined style="font-size: 16px;" />
            </template>
            <a-card hoverable class="max-w-md mx-auto">
              <template #cover>
                <img :alt="project.title" :src="project.image" />
              </template>
              <a-card-meta :title="project.title">
                <template #description>
                  <p>{{ project.description }}</p>
                  <div class="mt-4">
                    <a-tag v-for="tech in project.technologies" :key="tech" color="blue" class="mr-2 mb-2">
                      {{ tech }}
                    </a-tag>
                  </div>
                </template>
              </a-card-meta>
            </a-card>
          </a-timeline-item>
        </a-timeline>
      </div>
    </section>

    <!-- Experience Timeline -->
    <section class="py-20 bg-gray-50">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold mb-12 text-center">Work Experience</h2>
        <a-timeline mode="alternate">
          <a-timeline-item v-for="experience in experiences" :key="experience.id">
            <template #dot>
              <AppstoreOutlined style="font-size: 16px;" />
            </template>
            <a-card hoverable class="max-w-md mx-auto" @click="navigateToExperience(experience.id)">
              <a-card-meta :title="experience.role">
                <template #description>
                  <p class="font-semibold">{{ experience.company }}</p>
                  <p class="text-gray-500">{{ experience.period }}</p>
                  <p class="mt-2">{{ experience.description }}</p>
                </template>
              </a-card-meta>
            </a-card>
          </a-timeline-item>
        </a-timeline>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  UserOutlined,
  MailOutlined,
  CodeOutlined,
  AppstoreOutlined,
} from '@ant-design/icons-vue'

const router = useRouter()

const projects = ref([
  {
    id: 1,
    title: 'E-commerce App',
    image: '/src/assets/image/bg.jpg?height=200&width=300',
    description: 'A full-featured e-commerce mobile app with real-time inventory and personalized recommendations.',
    technologies: ['React Native', 'Node.js', 'MongoDB']
  },
  {
    id: 2,
    title: 'Social Media Platform',
    image: '/src/assets/image/bg.jpg?height=200&width=300',
    description: 'A modern social media platform focusing on privacy and user-generated content.',
    technologies: ['Swift', 'Firebase', 'WebRTC']
  },
  {
    id: 3,
    title: 'Fitness Tracker',
    image: '/src/assets/image/bg.jpg?height=200&width=300',
    description: 'An AI-powered fitness tracking app that provides personalized workout plans and nutrition advice.',
    technologies: ['Kotlin', 'TensorFlow', 'HealthKit']
  },
])

const experiences = ref([
  {
    id: 1,
    role: 'Senior Mobile Developer',
    company: 'Tech Corp',
    period: '2020 - Present',
    description: 'Leading a team of developers in creating cutting-edge mobile applications for enterprise clients.'
  },
  {
    id: 2,
    role: 'Mobile Developer',
    company: 'App Studio',
    period: '2018 - 2020',
    description: 'Developed and maintained various iOS and Android applications, focusing on performance and user experience.'
  },
  {
    id: 3,
    role: 'Junior Developer',
    company: 'Startup Inc',
    period: '2016 - 2018',
    description: 'Assisted in the development of mobile applications and learned best practices in a fast-paced startup environment.'
  },
])

const navigateToExperience = (id) => {
  router.push(`/experience/${id}`)
}
</script>

<style scoped>
.ant-timeline-item-content {
  width: 100%;
}
</style>

