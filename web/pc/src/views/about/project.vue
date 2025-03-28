<template>
  <div class="min-h-screen text-white p-8">
    <div class="max-w-7xl mx-auto">
      <div class="mb-8 md:mb-12 backdrop-blur-md  rounded-xl p-6 shadow-lg">
        <h1 class="text-4xl md:text-6xl font-bold">Projects</h1>
        <p class="text-lg text-white/70 mt-4" v-if="false">Places I've worked and grown</p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
            v-for="(project,key) in projects"
            :key="project.id"
            @click="toRoute('/about/project/'+project.id)"
            class="group relative aspect-[3/4] overflow-hidden rounded-lg bg-gray-900"
        >
          <img
              :src="$fileFull(project.photo)||defaultCover"
              :alt="project.title"
              class="absolute inset-0 w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>

          <div class="absolute inset-0 p-6 flex flex-col justify-between">
            <div class="flex justify-between items-start">
              <span class="text-5xl font-bold opacity-25">{{ key+1 }}</span>
              <span class="text-sm bg-white/10 px-3 py-1 rounded-full">{{ $formatDate(project.start,'{y}-{m}-{d}') }} -{{ $formatDate(project.end,'{y}-{m}-{d}') }} </span>
            </div>

            <div>
              <h3 class="text-2xl font-bold mb-2">{{ project.title }}</h3>
              <p class="text-gray-400 mb-4" v-html="project.description"></p>
              <button class="px-4 py-2 bg-white/10 rounded-full hover:bg-white/20 transition-colors">
                View Details
              </button>
            </div>
          </div>
        </div>
      </div>
      <div ref="observerTarget" class="h-10 mt-4"></div>
      <div v-if="loading" class="text-center py-8">
        <div class="inline-block animate-spin rounded-full h-8 w-8 border-t-2 border-white"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {userGet,projectList,experienceList} from "@/api/user";
import {onMounted, onUnmounted, ref} from 'vue'
import {useRouter} from "vue-router";
import defaultCover from "@/assets/image/default-cover.jpg";
const router = useRouter()
const toRoute = (path) => {
  router.push(path)
}
const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
}
const projects = ref([
  // {
  //   id: 1,
  //   number: '01',
  //   title: 'Engie',
  //   image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  //   date: '2024-01-08',
  //   description: 'Energy Management System with innovative solutions for sustainable future.'
  // },
  // {
  //   id: 1,
  //   number: '01',
  //   title: 'Engie',
  //   image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  //   date: '2024-01-08',
  //   description: 'Energy Management System with innovative solutions for sustainable future.'
  // },
  // {
  //   id: 1,
  //   number: '01',
  //   title: 'Engie',
  //   image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  //   date: '2024-01-08',
  //   description: 'Energy Management System with innovative solutions for sustainable future.'
  // },
  // {
  //   id: 1,
  //   number: '01',
  //   title: 'Engie',
  //   image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  //   date: '2024-01-08',
  //   description: 'Energy Management System with innovative solutions for sustainable future.'
  // },
  // {
  //   id: 1,
  //   number: '01',
  //   title: 'Engie',
  //   image: 'https://images.unsplash.com/photo-1537996194471-e657df975ab4?height=600&width=400',
  //   date: '2024-01-08',
  //   description: 'Energy Management System with innovative solutions for sustainable future.'
  // },
  // ... more projects
])
const loading = ref(false)
const current = ref(1)
const total = ref(0)
const pageSize = ref(6)
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
const list =()=>{
  if (loading.value || !hasMore.value) return;
  loading.value = true;
  projectList(Object.assign({},params.value,{"page":{"pageNum":current.value,"pageSize":pageSize.value}})).then(res=>{
    if(res&&res.code===200){
      let data = res.data.list || []
      projects.value = [...projects.value, ...data]; // res.data.list||[]
      total.value = res.data.total||0
      hasMore.value = data.length === pageSize.value;
      current.value++
    }
  }).finally(()=>{
    loading.value = false
  })
}
// onMounted(()=>{
//   projectList({}).then(res=>{
//     if(res&&res.code===200){
//       let data =res.data
//       projects.value =[ ...projects.value,...data.list]
//       projectsMore.value = projects.value.length<(data.total || 0)
//     }
//   })
//
// })
</script>
