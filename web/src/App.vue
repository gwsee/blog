<script setup>
</script>

<template>
  <router-view />
</template>

<style scoped>
.layout-header {
  color: #d47777;
  height: max-content;
  padding-inline: 10px;
  line-height: 50px;
  transition: background-color 2s;
  background-color: transparent;
  z-index: 99;
  display: flex;
  align-items: center; /* 垂直居中 */
}
.layout-header .menu{
  margin-left: auto
}
.layout-header:hover {
  background-color: #7dbcea;
}

.layout-content{
  text-align: center;
  line-height: 120px;
  color: #d47777;
  background-color: transparent;
}

.layout-footer{
  text-align: center;
  color: #d47777;
  background-color: transparent;
}
</style>
