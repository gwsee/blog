<template>
  <PhotoSphere :photos="photos" />
</template>

<script setup>
import PhotoSphere from './photosphere.vue'
import {onMounted, ref} from "vue";
import {photosList} from "@/api/photo.js";
import {filePrefix} from "@/api/tool.js";
// Example photos array
const photos = ref([])
onMounted(()=>{
  photosList({pageSize:66}).then(res=>{
    if(res&&res.data){
      let data = []
      const images = res.data.images||[]
      for(let i=0;i<images.length;i++){
        data.push({url:filePrefix+images[i],title:"",description:""})
      }

      photos.value = data
      console.log(photos)
    }
  })
})
</script>
