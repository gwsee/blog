<template>
  <div class="min-h-screen  text-white" >
    <!-- Hero Section -->
    <div  :style="{ backgroundImage: `url(${$fileFull(experience.image)})`,backgroundSize:`cover` }">
    <div class="relative h-[30vh] sm:h-[30vh] overflow-hidden max-w-6xl mx-auto px-4 sm:px-8 py-8 sm:py-16">
      <div class="absolute inset-0  from-black via-black/60 to-transparent"></div>
      <!-- Back Button -->
      <div class="absolute top-4 left-4 sm:top-8 sm:left-8 z-10">
        <button
            @click="$router.back()"
            class="flex items-center gap-2 text-shadow-lg hover:text-white transition-colors"
        >
          <ArrowLeftIcon class="w-5 h-5" />
          <span>Back</span>
        </button>
      </div>
      <!-- Content -->
      <div class="absolute bottom-0 left-0 w-full p-4 sm:p-8 ">
        <div class="max-w-6xl mx-auto">
          <div class="flex flex-col items-start sm:items-end sm:flex-row justify-between gap-4 sm:gap-6">
            <div>
              <h1 class="text-2xl sm:text-4xl md:text-6xl font-bold mt-1 sm:mt-2">{{ experience.company }}</h1>
              <h2 class="text-lg sm:text-xl md:text-2xl text-gray-300 mt-1 sm:mt-2">{{ experience.role }}</h2>
              <span class="text-xs sm:text-sm text-shadow-lg">{{ $formatDate(experience.start,'{y}-{m}-{d}') }} - {{ $formatDate(experience.end,'{y}-{m}-{d}') }}</span>
            </div>
            <div class="flex gap-2 sm:gap-4 mt-4 sm:mt-0">
              <a
                  v-if="experience.website"
                  :href="experience.website"
                  target="_blank"
                  class="px-4 sm:px-6 py-2 text-sm sm:text-base bg-white/10 backdrop-blur rounded-full hover:bg-white/20 transition-colors"
              >
                Visit Website
              </a>
              <button
                  class="px-4 sm:px-6 py-2 text-sm sm:text-base bg-white/10 backdrop-blur rounded-full hover:bg-white/20 transition-colors"
              >
                Share
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Main Content -->
    <div class="max-w-6xl mx-auto px-4 sm:px-8 py-8 sm:py-16">
      <div class="grid md:grid-cols-3 gap-8 md:gap-16">
        <!-- Left Column - Details -->
        <div class="md:col-span-2 space-y-8 sm:space-y-12">
          <!-- Overview -->
          <div class="bg-gray-800/50 rounded-xl p-4 sm:p-6">
            <h3 class="text-xl sm:text-2xl font-bold mb-4 sm:mb-6">Description</h3>
            <div class="prose prose-invert">
              <p class="text-sm sm:text-base text-gray-300 leading-relaxed" v-html="formattedContent(experience.description)"></p>
            </div>
            <h3 class="text-xl sm:text-2xl font-bold mb-4 sm:mb-6">Responsibilities</h3>
            <div class="prose prose-invert">
              <p class="text-sm sm:text-base text-gray-300 leading-relaxed" v-html="formattedContent(experience.responsibilities)"></p>
            </div>
          </div>
          <!-- Technologies -->
          <div class="bg-gray-800/50 rounded-xl p-4 sm:p-6">
            <h3 class="text-xl sm:text-2xl font-bold mb-4 sm:mb-6">Technologies & Tools</h3>
            <div class="flex flex-wrap gap-2 sm:gap-3">
              <div
                  v-for="(tech,key) in experience.skills"
                  :key="tech"
                  class="group flex items-center gap-2 px-3 sm:px-4 py-2 bg-gray-800 rounded-lg hover:bg-gray-700 transition-colors"
              >
                <component
                    :is="skillsIcons&&skillsIcons[key%4]"
                    class="w-4 h-4 sm:w-5 sm:h-5 text-gray-400 group-hover:text-white transition-colors"
                />
                <span class="text-sm sm:text-base">{{ tech }}</span>
              </div>
            </div>
          </div>
          <!-- Projects -->
          <div class="bg-gray-800/50 rounded-xl p-4 sm:p-6">
            <h3 class="text-xl sm:text-2xl font-bold mb-4 sm:mb-6">Key Projects</h3>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3 sm:gap-4 md:gap-6">
              <div
                  v-for="project in projects"
                  :key="project.id"
                  @click="toRoute('/about/project/'+project.id)"
                  class="group relative bg-gray-800 rounded-lg overflow-hidden cursor-pointer h-[200px] sm:h-[220px] md:h-[240px]"
              >
                <img
                    :src="$fileFull(project.photo)"
                    :alt="project.title"
                    class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/60 to-transparent"></div>
                <div class="absolute inset-0 p-4 flex flex-col justify-end">
                  <div class="space-y-2">
                    <div class="flex items-center justify-between">
                      <span class="text-base font-medium line-clamp-1" :title="project.title">{{ project.title }}</span>

                    </div>
                    <p class="text-xs text-gray-400">
                        {{ $formatDate(project.start,'{y}-{m}') }} - {{ $formatDate(project.end,'{y}-{m}') }}
                     </p>
                    <p class="text-sm text-gray-400 line-clamp-2">{{ project.description }}</p>
                    <div class="flex flex-wrap gap-1.5 mt-2">
                      <span
                          v-for="tech in project.skills"
                          :key="tech"
                          class="px-2 py-0.5 text-xs bg-white/10 rounded-full"
                      >
                        {{ tech }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
              <!-- View More Card -->
<!--              <div-->
<!--                  v-if="experience.projects.length > 0"-->
<!--                  class="relative bg-gray-800/50 rounded-lg overflow-hidden cursor-pointer h-[200px] sm:h-[220px] md:h-[240px] hover:bg-gray-800/70 transition-colors"-->
<!--              >-->
<!--                <div class="absolute inset-0 flex items-center justify-center">-->
<!--                  <div class="text-center">-->
<!--                    <p class="text-base font-medium text-gray-400">View More</p>-->
<!--                    <p class="text-sm text-gray-500">{{ experience.projects.length }} Projects</p>-->
<!--                  </div>-->
<!--                </div>-->
<!--              </div>-->
            </div>
          </div>
        </div>

        <!-- Right Column - Additional Info -->
        <div class="space-y-6 sm:space-y-8">
          <!-- Quick Facts -->
          <div class="bg-gray-800/50 rounded-xl p-4 sm:p-6">
            <h4 class="text-base sm:text-lg font-medium mb-3 sm:mb-4">Quick Facts</h4>
            <div class="space-y-3 sm:space-y-4">
              <div>
                <span class="text-xs sm:text-sm text-gray-400">Location</span>
                <p class="text-sm sm:text-base">{{ experience.location }}</p>
              </div>
              <div>
                <span class="text-xs sm:text-sm text-gray-400">Team Size</span>
                <p class="text-sm sm:text-base">{{ experience.role }}</p>
              </div>
              <div>
                <span class="text-xs sm:text-sm text-gray-400">Industry</span>
                <p class="text-sm sm:text-base">{{ experience.industry }}</p>
              </div>
            </div>
          </div>

          <!-- Achievements -->
          <div class="bg-gray-800/50 rounded-xl p-4 sm:p-6">
            <h4 class="text-base sm:text-lg font-medium mb-3 sm:mb-4">Key Achievements</h4>
            <div class="space-y-3 sm:space-y-4"  v-html="formattedContent(experience.achievements)">
            </div>
          </div>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script setup>
import {experienceGet, experienceSave, experienceList, projectList} from "@/api/user";
import {onMounted, ref} from 'vue'
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
import {useRoute} from "vue-router";
const skillsIcons = ref[CodeBracketIcon,ServerIcon,CommandLineIcon,CpuChipIcon]
// Sample data
const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
}
const experience = ref({
  id: 1,
  company: 'Aris Digital',
  role: 'Senior Developer',
  period: '2022 - Present',
  location: 'Shanghai, China',
  teamSize: '15-20 members',
  industry: 'Technology & Software',
  website: 'https://example.com',
  achievements:"",
  image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects. Worked on various challenging projects that pushed the boundaries of web technology.',
  highlights: [
    'Led a team of 5 developers in building a new e-commerce platform',
    'Implemented CI/CD pipelines reducing deployment time by 60%',
    'Mentored 3 junior developers who were promoted to mid-level positions',
    'Introduced new testing practices that increased code coverage by 40%'
  ],
  projects: [

  ],
  technologies: [
    { name: 'Frontend Development', icon: CodeBracketIcon },
    { name: 'Backend Systems', icon: ServerIcon },
    { name: 'DevOps', icon: CommandLineIcon },
    { name: 'Microservices', icon: CpuChipIcon }
  ],
})
const projects = ref([{
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
  }])
onMounted(()=>{
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  experienceGet({id:id}).then((res) => {
    if(res){
      const data = res.data || {}
      experience.value = data
    }
  })
  projectList({experienceId: id,"page":{"pageNum":1,"pageSize":9999}}).then((res) => {
    if(res){
      projects.value=res.data&&res.data.list||[]
    }
  })
})
import {useRouter} from "vue-router";
const router = useRouter()
const toRoute = (path) => {
  router.push(path)
}
</script>

<style>
.prose {
  max-width: 65ch;
}

/* Add smooth scrolling */
html {
  scroll-behavior: smooth;
}

/* Improve image loading */
img {
  transition: opacity 0.3s ease;
}

img[loading] {
  opacity: 0;
}

img.loaded {
  opacity: 1;
}
</style>

