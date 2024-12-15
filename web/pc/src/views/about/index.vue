<template>
  <div class="min-h-screen bg-gradient-to-b ">
    <!-- Hero Section -->
    <section class="relative py-12 overflow-hidden">
      <div class="container mx-auto px-4">
        <div class="flex flex-col md:flex-row items-center gap-8">
          <!-- Profile Image -->
          <div class="md:w-1/3">
            <div class="relative w-48 h-48 mx-auto rounded-full overflow-hidden border-4 border-white shadow-lg">
              <img
                  src="/src/assets/image/bg.jpg?height=200&width=200"
                  alt="Profile"
                  class="object-cover w-full h-full"
              />
            </div>
          </div>
          <!-- Text Content -->
          <div class="md:w-2/3 text-center md:text-left">
            <h1 class="text-3xl md:text-4xl font-bold mb-4">
              John Doe
            </h1>
            <p class="text-lg text-gray-600 mb-6">
              Experienced Mobile App Developer specializing in creating innovative solutions with cutting-edge technology.
            </p>
            <div class="flex justify-center md:justify-start gap-4">
              <a-button type="primary" size="large">View Projects</a-button>
              <a-button size="large">Contact Me</a-button>
              <a-button size="large" @click="toRoute('/about/profile/manage')">Manage</a-button>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Projects Grid -->
    <section class="py-20">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold mb-12">Featured Projects</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="project in projects" :key="project.id" class="group relative">
            <div class="relative aspect-[4/3] rounded-lg overflow-hidden bg-white shadow-lg">
              <img
                  :src="project.image"
                  :alt="project.title"
                  class="w-full h-full object-cover"
              />
              <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                <div class="absolute bottom-0 left-0 right-0 p-4">
                  <h3 class="text-white font-medium">{{ project.title }}</h3>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Experience Timeline -->
    <section class="py-20 ">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold mb-12 text-center">Work Experience</h2>
        <div class="relative">
          <!-- Experience Cards Container -->
          <a-timeline mode="alternate">
            <a-timeline-item v-for="experience in experiences" :key="experience.id"   :class="index % 2 === 0 ? 'self-start' : 'self-start'">
              <template #dot>
                <AppstoreOutlined style="font-size: 16px;" />
              </template>
              <a-card
                  hoverable
                  class="experience-card overflow-hidden"
                  style="text-align: left"
                  @click="toRoute(`/about/experience/1`)"
              >
                <template #cover>
                  <div class="h-32 bg-gradient-to-r from-blue-500 to-purple-500"></div>
                </template>
                <a-card-meta :title="experience.role">
                  <template #description>
                    <div>
                      <p class="text-gray-600">{{ experience.company }}</p>
                      <p class="text-sm text-gray-500">{{ experience.period }}</p>
                    </div>
                  </template>
                </a-card-meta>
                <div class="mt-4 experience-details opacity-0 transition-opacity duration-300">
                  <p class="text-sm text-gray-600">{{ experience.description }}</p>
                </div>
              </a-card>
            </a-timeline-item>
          </a-timeline>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import {userGet,projectList,experienceList} from "@/api/user";
import {ref} from 'vue'
import {useRouter} from 'vue-router'
const router = useRouter()
import {
  UserOutlined,
  MailOutlined,
  CodeOutlined,
  AppstoreOutlined,
} from '@ant-design/icons-vue'

const projects = ref([
  {
    id: 1,
    title: 'E-commerce App',
    image: '/src/assets/image/bg.jpg?height=400&width=600',
  },
  {
    id: 2,
    title: 'Social Media Platform',
    image: '/src/assets/image/bg.jpg?height=400&width=600',
  },
  {
    id: 3,
    title: 'Fitness Tracker',
    image: '/src/assets/image/bg.jpg?height=400&width=600',
  },
  {
    id: 4,
    title: 'Task Management',
    image: '/src/assets/image/bg.jpg?height=400&width=600',
  },
])

const experiences = ref([
  {
    id: 1,
    role: 'Senior Mobile Developer',
    company: 'Tech Corp',
    period: '2020 - Present',
    description: 'Led the development of multiple high-profile mobile applications, managing a team of developers and implementing cutting-edge technologies.'
  },
  {
    id: 2,
    role: 'Mobile Developer',
    company: 'App Studio',
    period: '2018 - 2020',
    description: 'Developed and maintained various mobile applications for iOS and Android platforms, focusing on performance optimization and user experience.'
  },
  {
    id: 3,
    role: 'Junior Developer',
    company: 'Startup Inc',
    period: '2016 - 2018',
    description: 'Assisted in the development of mobile applications, learned best practices in mobile development, and contributed to the growth of a fast-paced startup.'
  },
])

const toRoute = (path) => {
  router.push(path)
}
</script>

<style scoped>
/* ... existing styles ... */

.experience-card {
  transform: rotate(-5deg);
}

.experience-card:hover {
  transform: rotate(0deg);
}

.experience-card:hover .experience-details {
  opacity: 1;
}
</style>

