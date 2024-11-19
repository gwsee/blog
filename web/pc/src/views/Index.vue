<script setup>
import {ref} from "vue"
const value = ref('');
const onQuery = () => {
  console.log('query!', value.value);
};
</script>

<template>
  <div class="main-container search-content"  >
    <a-input-group compact style="margin-top: -66px;">
      <a-input v-model:value="value" placeholder="you can search blog title from here!"
               style="width: calc(60% - 100px);height: 40px;background: linear-gradient(238deg, rgb(170 189 166 / 99%) 0%, rgb(241 235 235) 3%, rgb(187 243 184) 99%);" />
      <a-button type="default" style="height: 40px" @click="onQuery">Query</a-button>
    </a-input-group>
  </div>
</template>


<style scoped>
.search-content{
  align-content: space-around;
  text-align: center;
  justify-content: center;
}
</style>
