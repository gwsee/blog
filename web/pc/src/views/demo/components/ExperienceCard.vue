<template>
  <div
      class="flex-shrink-0 w-[280px] h-[400px] md:w-[320px] md:h-[500px]
           rounded-lg overflow-hidden cursor-pointer group relative"
  >
    <img
        :src="image"
        :alt="company"
        class="absolute inset-0 w-full h-full object-cover"
    />
    <div :class="`absolute inset-0 bg-gradient-to-b ${color} opacity-75`"></div>
    <div class="absolute inset-0 p-6 flex flex-col justify-between transform transition-all duration-500 group-hover:scale-105">
      <div>
        <span class="text-4xl font-bold opacity-50">{{ number }}</span>
        <div class="mt-2 text-sm opacity-75">{{ period }}</div>
      </div>
      <div>
        <h3 class="text-2xl font-bold mb-2">{{ company }}</h3>
        <h4 class="text-lg mb-2">{{ role }}</h4>
        <p class="text-sm opacity-75">{{ description }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  id: Number,
  number: String,
  company: String,
  role: String,
  period: String,
  image: String,
  color: String,
  description: String
})
</script>

