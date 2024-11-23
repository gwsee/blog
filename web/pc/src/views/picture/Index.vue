<template>
  <PhotoSphere :photos="photos" />
</template>

<script setup>
import PhotoSphere from './photosphere.vue'
import {ref} from "vue";
// Example photos array
const photos = ref([
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
  {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  }, {   url: '/src/assets/image/bg.jpg?height=300&width=300',
    title: 'Photo 1',
    description: 'Description for photo 1'
  },
])

</script>
