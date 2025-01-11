<template>
  <div class="min-h-screen text-white p-4 md:p-8">
    <div class="max-w-6xl mx-auto overflow-hidden">
      <div class="mb-8 md:mb-12">
        <h1 class="text-4xl md:text-6xl font-bold">Experience</h1>
        <p class="text-lg text-gray-400 mt-4">Places I've worked and grown</p>
      </div>

      <div class="space-y-24" ref="experienceContainer">
        <div
            v-for="(experience, index) in displayedExperiences"
            :key="experience.id"
            class="relative"
        >
          <!-- Experience Grid Container -->
          <div class="experience-grid">
            <!-- Top Left - Date -->
            <div class="quadrant top-left">
              <span class="text-gray-400 text-sm">
                {{ $formatDate(experience.start,'{y}-{m}-{d}') }} -{{ $formatDate(experience.end,'{y}-{m}-{d}') }}
              </span>
            </div>

            <!-- Top Right - Company Info -->
            <div class="quadrant top-right">
              <h2 class="text-3xl font-bold">{{ experience.company }}</h2>
              <h3 class="text-xl text-gray-400">{{ experience.role }}</h3>
              <p class="text-gray-400 leading-relaxed mt-2 text-sm md:text-base"
                 v-html="formattedContent(experience.description)">
              </p>
            </div>

            <!-- Bottom Left - Featured Image -->
            <div class="quadrant bottom-left">
              <div
                  @click="toRoute('/about/experience/'+experience.id)"
                  class="featured-image-container"
              >
                <img
                    :src="$fileFull(experience.image)||'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'"
                    :alt="experience.company"
                    class="featured-image"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
              </div>
            </div>

            <!-- Bottom Right - Projects Grid -->
            <div class="quadrant bottom-right">
              <div v-if="experience.project && experience.project.length > 0"
                   class="projects-grid">
                <!-- Project Cards -->
                <template v-for="(project, idx) in experience.project.slice(0, 3)" :key="project.id">
                  <div
                      @click="toRoute('/about/project/'+project.id)"
                      class="project-card"
                  >
                    <div class="project-image-container">
                      <img
                          :src="$fileFull(project.photo)||'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'"
                          :alt="project.title"
                          class="project-image"
                      />
                      <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
                      <div class="absolute inset-0 p-3 flex flex-col justify-end">
                        <span class="text-sm font-medium">{{ project.title }}</span>
                        <span class="text-xs text-gray-400">
                          {{ $formatDate(project.start,'{y}-{m}-{d}') }} -{{ $formatDate(project.end,'{y}-{m}-{d}') }}
                        </span>
                      </div>
                    </div>
                  </div>
                </template>

                <!-- View More Card -->
                <div
                    v-if="experience.project.length > 3"
                    class="project-card view-more-card"
                >
                  <div class="project-image-container"
                       @click="toRoute('/about/project?experience_id='+experience.id)">
                    <div class="flex items-center justify-center h-full">
                      <div class="text-center">
                        <div class="mb-1">
                          <PlusCircleIcon class="w-6 h-6 mx-auto text-gray-400 group-hover:text-white transition-colors" />
                        </div>
                        <span class="text-xs font-medium text-gray-400 group-hover:text-white transition-colors">
                          View {{ experience.project.length - 3 }} More
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div ref="observerTarget" class="h-10 mt-4"></div>
      <!-- Loading indicator -->
      <div v-if="loading" class="text-center py-8">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-t-2 border-white"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {userGet,projectList,experienceList} from "@/api/user";
import { ref, onMounted, onUnmounted } from 'vue'

const displayedExperiences = ref([ ])
const loading = ref(false)
const current = ref(1)
const total = ref(0)
const pageSize = ref(16)
const experiencesMore = ref(false)
let observer = null;
const observerTarget = ref(null)
const hasMore = ref(true)
const params = ref({my:false,title:undefined,myCollect:false,sort:"",myThumb:false,description:""})
const setupIntersectionObserver = () => {
  observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting && !loading.value && hasMore.value) {
      list();
    }
  }, {
    root: null,
    rootMargin: '0px',
    threshold: 0.1
  });

  if (observerTarget.value) {
    observer.observe(observerTarget.value);
  }
}
onMounted(()=>{
  // list()
  setupIntersectionObserver();
})
onUnmounted(() => {
  if (observer) {
    observer.disconnect();
  }
})
const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
}
const list =()=>{
  if (loading.value || !hasMore.value) return;
  loading.value = true;
  experienceList(Object.assign({},params.value,{"page":{"pageNum":current.value,"pageSize":pageSize.value}})).then(res=>{
    if(res&&res.code===200){
      let data = res.data.list || []
      displayedExperiences.value = [...displayedExperiences.value, ...data]; // res.data.list||[]
      total.value = res.data.total||0
      hasMore.value = data.length === pageSize.value;
      current.value++
    }
  }).finally(()=>{
    loading.value = false
  })
}
// Initialize with first batch of experiences
import {useRouter} from "vue-router";
const router = useRouter()
const toRoute = (path) => {
  router.push(path)
}

// const allExperiences = [
//   {
//     id: 1,
//     number: '01',
//     company: 'Aris Digital',
//     role: 'Senior Developer',
//     period: '2022 - Present',
//     description: 'Led development of enterprise-scale applications, mentored junior developers, and implemented best practices across multiple projects.',
//     image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
//     projects: [
//       {
//         id: 1,
//         title: 'E-commerce Platform',
//         date: '2023',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 2,
//         title: 'Banking App',
//         date: '2023',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 3,
//         title: 'CRM System',
//         date: '2022',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 4,
//         title: 'Analytics Dashboard',
//         date: '2022',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 5,
//         title: 'Mobile App',
//         date: '2022',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       }
//     ]
//   },
//   {
//     id: 2,
//     number: '02',
//     company: 'Mon Studio',
//     role: 'Frontend Developer',
//     period: '2020 - 2022',
//     description: 'Developed innovative user interfaces and implemented responsive designs for various client projects.',
//     image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
//     projects: [
//       {
//         id: 1,
//         title: 'Portfolio Website',
//         date: '2022',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 2,
//         title: 'Design System',
//         date: '2021',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       },
//       {
//         id: 3,
//         title: 'Marketing Site',
//         date: '2021',
//         image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=300&width=300'
//       }
//     ]
//   }
// ]

// displayedExperiences.value = allExperiences.slice(0, 2)
// const experienceContainer = ref(null)
// // Intersection Observer setup
// const observerCallback = async (entries) => {
//   const target = entries[0]
//   if (target.isIntersecting && !loading.value) {
//     await loadMore()
//   }
// }
//
// const loadMore = async () => {
//   if (loading.value) return
//
//   loading.value = true
//
//   // Simulate API call with setTimeout
//   await new Promise(resolve => setTimeout(resolve, 1000))
//
//   const nextBatch = allExperiences.slice(
//       displayedExperiences.value.length,
//       displayedExperiences.value.length + 2
//   )
//
//   if (nextBatch.length > 0) {
//     displayedExperiences.value = [...displayedExperiences.value, ...nextBatch]
//     page.value++
//   }
//
//   loading.value = false
// }
//
// let observer
//
// onMounted(() => {
//   observer = new IntersectionObserver(observerCallback, {
//     root: null,
//     rootMargin: '100px',
//     threshold: 0.1
//   })
//
//   if (experienceContainer.value) {
//     observer.observe(experienceContainer.value)
//   }
// })
//
// onUnmounted(() => {
//   if (observer) {
//     observer.disconnect()
//   }
// })
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

.experience-grid {
  display: grid;
  grid-template-areas:
    "top-left top-right"
    "bottom-left bottom-right";
  grid-template-columns: 1fr 1fr;
  grid-template-rows: auto 1fr;
  gap: 2rem;
  min-height: 400px;
}

.quadrant {
  padding: 1rem;
}

.top-left {
  grid-area: top-left;
}

.top-right {
  grid-area: top-right;
}

.bottom-left {
  grid-area: bottom-left;
  padding: 0;
}

.bottom-right {
  grid-area: bottom-right;
  padding: 0;
}

.featured-image-container {
  position: relative;
  width: 100%;
  height: 0;
  padding-bottom: 100%;
  overflow: hidden;
  border-radius: 0.5rem;
}

.featured-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.projects-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 0.75rem;
  height: 100%;
}

.project-card {
  position: relative;
  background: rgb(31, 41, 55);
  border-radius: 0.5rem;
  overflow: hidden;
  cursor: pointer;
}

.project-image-container {
  position: relative;
  width: 100%;
  height: 0;
  padding-bottom: 100%;
}

.project-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s;
}

.project-card:hover .project-image {
  transform: scale(1.1);
}

.view-more-card {
  background: rgba(31, 41, 55, 0.5);
  border: 1px solid rgb(55, 65, 81);
}

.view-more-card:hover {
  border-color: rgb(107, 114, 128);
}

@media (max-width: 768px) {
  .experience-grid {
    gap: 1rem;
  }

  .projects-grid {
    gap: 0.5rem;
  }
}
</style>

