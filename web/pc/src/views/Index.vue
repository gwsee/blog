<script setup>
import {ref} from "vue"
const value = ref('');
const onQuery = () => {
  console.log('query!', value.value);
};
</script>

<template>
  <div class="main-container search-content"  >
    <a-input-group compact style="margin-top: -66px">
      <a-input v-model:value="value" placeholder="you can search blog title from here!"
               style="width: calc(60% - 100px);height: 40px;background-color: transparent" />
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
