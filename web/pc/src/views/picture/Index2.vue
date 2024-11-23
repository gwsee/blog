<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <a-button @click="router.back()" type="text" size="large">
            <template #icon><ArrowLeftOutlined /></template>
            Back
          </a-button>
          <div class="flex items-center gap-4">
            <a-button type="primary" shape="round">
              <template #icon><DownloadOutlined /></template>
              Download CV
            </a-button>
            <a-button shape="round">
              <template #icon><ShareAltOutlined /></template>
              Share
            </a-button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-12">
      <div class="max-w-4xl mx-auto">
        <!-- Basic Info -->
        <a-card class="mb-8 overflow-hidden">
          <div class="flex flex-col md:flex-row items-center md:items-start gap-6">
            <a-avatar :size="120" :src="experience.avatar" />
            <div class="flex-1 text-center md:text-left">
              <h1 class="text-3xl font-bold mb-2">{{ experience.role }}</h1>
              <p class="text-xl text-gray-600 mb-4">{{ experience.company }}</p>
              <div class="flex flex-wrap justify-center md:justify-start gap-4 text-sm text-gray-500">
                <span class="flex items-center gap-2">
                  <CalendarOutlined />
                  {{ experience.period }}
                </span>
                <span class="flex items-center gap-2">
                  <EnvironmentOutlined />
                  {{ experience.location }}
                </span>
              </div>
            </div>
          </div>
        </a-card>

        <!-- Details -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Left Sidebar -->
          <div class="space-y-8">
            <!-- Skills -->
            <a-card title="Skills" :bodyStyle="{ padding: '24px' }">
              <div class="space-y-4">
                <div v-for="skill in experience.skills" :key="skill.name">
                  <div class="flex justify-between mb-1">
                    <span>{{ skill.name }}</span>
                    <span>{{ skill.level }}%</span>
                  </div>
                  <a-progress
                      :percent="skill.level"
                      :showInfo="false"
                      :strokeColor="{
                      '0%': '#108ee9',
                      '100%': '#87d068',
                    }"
                  />
                </div>
              </div>
            </a-card>

            <!-- Tools -->
            <a-card title="Tools & Technologies" :bodyStyle="{ padding: '24px' }">
              <div class="flex flex-wrap gap-2">
                <a-tag v-for="tool in experience.tools" :key="tool" color="blue">
                  {{ tool }}
                </a-tag>
              </div>
            </a-card>
          </div>

          <!-- Main Content -->
          <div class="md:col-span-2 space-y-8">
            <!-- About -->
            <a-card title="About Role" :bodyStyle="{ padding: '24px' }">
              <p class="text-gray-600 whitespace-pre-line">{{ experience.description }}</p>
            </a-card>

            <!-- Key Responsibilities -->
            <a-card title="Key Responsibilities" :bodyStyle="{ padding: '24px' }">
              <a-list :dataSource="experience.responsibilities" size="small">
                <template #renderItem="{ item }">
                  <a-list-item>
                    <template #actions>
                      <a-tag color="green">Completed</a-tag>
                    </template>
                    <a-list-item-meta>
                      <template #description>
                        <span class="text-gray-600">{{ item }}</span>
                      </template>
                    </a-list-item-meta>
                  </a-list-item>
                </template>
              </a-list>
            </a-card>

            <!-- Achievements -->
            <a-card title="Key Achievements" :bodyStyle="{ padding: '24px' }">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <a-card
                    v-for="(achievement, index) in experience.achievements"
                    :key="index"
                    size="small"
                    :bodyStyle="{ padding: '16px' }"
                >
                  <template #title>
                    <div class="flex items-center gap-2">
                      <TrophyOutlined style="color: #faad14;" />
                      <span>{{ achievement.title }}</span>
                    </div>
                  </template>
                  <p class="text-sm text-gray-600">{{ achievement.description }}</p>
                </a-card>
              </div>
            </a-card>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  ArrowLeftOutlined,
  CalendarOutlined,
  DownloadOutlined,
  EnvironmentOutlined,
  ShareAltOutlined,
  TrophyOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()

const experience = ref({
  avatar: '/placeholder.svg?height=120&width=120',
  role: 'Senior Mobile Developer',
  company: 'Tech Corp',
  period: '2020 - Present',
  location: 'San Francisco, CA',
  description: `Led the mobile development team in creating innovative solutions for enterprise clients.
  Responsible for architecture decisions and implementing best practices across mobile projects.
  Collaborated with cross-functional teams to deliver high-quality mobile applications on time and within budget.`,
  skills: [
    { name: 'React Native', level: 95 },
    { name: 'iOS Development', level: 85 },
    { name: 'Android Development', level: 80 },
    { name: 'UI/UX Design', level: 75 },
  ],
  tools: [
    'React Native',
    'Swift',
    'Kotlin',
    'Firebase',
    'Git',
    'Jira',
    'Figma',
  ],
  responsibilities: [
    'Led development of multiple mobile applications from concept to deployment',
    'Managed team of 5 developers and coordinated with cross-functional teams',
    'Implemented CI/CD pipelines and automated testing procedures',
    'Conducted code reviews and maintained code quality standards',
    'Mentored junior developers and facilitated knowledge sharing sessions',
  ],
  achievements: [
    {
      title: 'App Store Feature',
      description: 'Main app featured on App Store\'s main page for 2 weeks, resulting in a 500% increase in downloads.'
    },
    {
      title: 'Performance Improvement',
      description: 'Reduced app load time by 40% through optimization techniques, significantly improving user experience.'
    },
    {
      title: 'Team Growth',
      description: 'Grew mobile team from 2 to 5 developers, increasing productivity and expanding project capabilities.'
    },
    {
      title: 'Process Improvement',
      description: 'Implemented new workflow reducing deployment time by 60%, allowing for more frequent updates.'
    },
  ],
})

onMounted(() => {
  // In a real app, you would fetch the experience data based on the route.params.id
  console.log('Experience ID:', route.params.id)
})
</script>

<style scoped>
.ant-card {
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
</style>

