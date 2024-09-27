<script setup>
// import logo from './assets/image/Sara11712577366099546.jpg';
// const bg = logo
</script>

<template>
  <a-layout style='min-height: 100%;background-color: transparent'>
    <a-layout-header class="layout-header">
      <div class="logo">

      </div>
      <div class="title">

      </div>
      <div class="menu">
        <a-avatar :size="44" style="border: 1px solid red;" >登陆</a-avatar>
        <a-avatar :size="44" style="border: 1px solid red;" src="https://www.antdv.com/assets/logo.1ef800a8.svg" >USD</a-avatar>
      </div>

    </a-layout-header>
    <a-layout-content class="layout-content">
<!--      <a-image :src='bg' />-->
      center
    </a-layout-content>
    <a-layout-footer class="layout-footer">Footer</a-layout-footer>
  </a-layout>
</template>

<style scoped>
.layout-header {
  color: #d47777;
  height: 64px;
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
