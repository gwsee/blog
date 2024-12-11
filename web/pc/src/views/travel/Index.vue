<template>
  <div class="min-h-screen bg-gray-50">
    <a-layout style="background-color: transparent">
      <a-layout-content class="p-4 md:p-8">
        <div class="relative max-w-7xl mx-auto">
          <!-- Grid Container -->
          <a-row :gutter="[16, 16]">
            <a-col
                v-for="project in projects"
                :key="project.id"
                :xs="24"
                :sm="24"
                :md="12"
                :lg="6"
                :xl="6"
            >
              <!-- Project Card -->
              <div
                  class="project-card group relative overflow-hidden rounded-lg"
                  :class="{'bg-black': project.darkMode}"
              >
                <!-- Project Image Container with Fixed Aspect Ratio -->
                <div class="relative w-full pb-[75%]">
                  <img
                      v-if="project.image"
                      :src="project.image"
                      :alt="project.title"
                      class="absolute inset-0 w-full h-full object-cover"
                  />
                  <div
                      v-else
                      class="absolute inset-0 bg-gray-100 flex items-center justify-center"
                  >
                    <div class="text-4xl text-gray-400">No Image</div>
                  </div>

                  <!-- Project Title -->
                  <div class="absolute bottom-0 left-0 right-0 p-4">
                    <h3 class="text-lg font-medium text-white truncate overflow-hidden whitespace-nowrap">{{ project.title }}</h3>
                  </div>

                  <!-- Hover Actions -->
                  <div
                      class="absolute top-0 left-0 right-0 bottom-0 bg-black/40 opacity-0
                           group-hover:opacity-100 transition-opacity duration-200 flex
                           items-center justify-center gap-2"
                  >
                    <a-button
                        class="action-button !bg-white/90 hover:!bg-white"
                        shape="circle"
                        size="large"
                    >
                      <template #icon><SaveOutlined /></template>
                    </a-button>
                    <a-button
                        class="action-button !bg-white/90 hover:!bg-white"
                        shape="circle"
                        size="large"
                    >
                      <template #icon><HeartOutlined /></template>
                    </a-button>
                  </div>
                </div>

                <!-- Project Info -->
                <div class="p-4 flex items-center justify-between">
                  <div class="flex items-center gap-2">
                    <a-avatar :src="project.author.avatar" :size="24" />
                    <span class="text-sm" :class="{'text-white': project.darkMode}">
                      {{ project.author.name }}
                    </span>
                    <a-tag
                        v-if="project.author.pro"
                        class="border-0 text-xs"
                        :class="project.darkMode ? 'bg-white text-black' : 'bg-black text-white'"
                    >
                      ME
                    </a-tag>
                  </div>
                  <div class="flex items-center gap-3">
                    <span class="flex items-center gap-1" :class="{'text-white': project.darkMode}">
                      <HeartOutlined />
                      {{ project.likes }}
                    </span>
                    <span class="flex items-center gap-1" :class="{'text-white': project.darkMode}">
                      <EyeOutlined />
                      {{ project.views }}
                    </span>
                  </div>
                </div>
              </div>
            </a-col>

          </a-row>
          <div
              class="sticky float-right bottom-2 right-2 w-12 h-12 bg-blue-500 text-white flex items-center justify-center rounded-full right"
          >
            <a-button
                class="action-button !bg-white/90 hover:!bg-white"
                shape="circle"
                size="large"
                @click="toRoute('/travel/manage/0')"
            >
             NEW
            </a-button>
          </div>
        </div>
      </a-layout-content>

    </a-layout>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  HeartOutlined,
  EyeOutlined,
  SaveOutlined
} from '@ant-design/icons-vue'
import avatar from "@/assets/image/default-avatar.png"
import {useRouter} from "vue-router";
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}
const projects = ref([
  {
    id: 1,
    title: 'Miller Design Build Logo',
    image: '/placeholder.svg?height=600&width=800',
    darkMode: false,
    author: {
      name: 'Dalius Stuoka',
      avatar: avatar,
      pro: true
    },
    views: '11k',
    likes: 93
  },
  {
    id: 2,
    title: 'Floral PatternFloral PatternFloral PatternFloral PatternFloral PatternFloral PatternFloral Pattern',
    image: avatar,
    darkMode: true,
    author: {
      name: 'Aleksandar Savić',
      avatar: avatar,
      pro: true
    },
    views: '4.3k',
    likes: 57
  },
  {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  },
  {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  },
  {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  },
  {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  }, {
    id: 3,
    title: 'Project without Image',
    image: null,
    darkMode: false,
    author: {
      name: 'John Doe',
      avatar: '/placeholder.svg?height=40&width=40',
      pro: false
    },
    views: '2.1k',
    likes: 42
  },
  // Add more projects as needed...
])
</script>

<style scoped>
.project-card {
  transition: transform 0.2s ease;
}

.project-card:hover {
  transform: translateY(-2px);
}

.action-button {
  transition: all 0.2s ease;
}

.action-button:hover {
  transform: scale(1.05);
}

/* Override Ant Design button styles */
v-deep::deep (.ant-btn-circle) {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
}
</style>

