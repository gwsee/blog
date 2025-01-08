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
      images.forEach(item=>{
        item.url = filePrefix+item.url
      })
      photos.value = images
      console.log(photos)
    }
  })
})
</script>
