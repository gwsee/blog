<script setup>
import Login from "@/components/Login.vue";
import {getInfo} from "@/api/index.js";
import {ref} from "vue"
import { useRouter } from 'vue-router';

const router = useRouter();
const loginRef = ref(null);

const openDialog = () => {
  if( loginRef.value){
    loginRef.value.show()
  }
};
const getInfoData =()=>{
  getInfo().then(res=>{
    console.log(res)
  })
}
const goToAbout=(path)=> {
  router.push(path)
}
</script>

<template>
  <!--  :style='{minHeight: "100%",backgroundColor: "transparent",backgroundImage: bg}'-->
  <a-layout style='min-height: 100%;background-color: transparent;background-image: url("src/assets/image/Sara11712577366099546.jpg");background-size: cover'>
    <a-layout-header class="layout-header">
      <a-row style="width: 100%"  >
        <a-col :sm="8" :xs="24" >
            <a-image
                @click="goToAbout('/index')"
                :preview="false"
                :width="200"
                :height="44"
                class="layout-content-logo-item"
                src="https://www.antdv.com/assets/logo.1ef800a8.svg"
            />
        </a-col>
        <a-col :sm="8" :xs="24">
        </a-col>
        <a-col :sm="8" :xs="24">
          <a-avatar :size="44" class="layout-content-menu-item" @click="openDialog" >登陆</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="getInfoData" >测试</a-avatar>
           <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/about')" >关于</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/travel')" >旅行</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/blog')" >日记</a-avatar>
        </a-col>
      </a-row>
    </a-layout-header>
    <a-layout-content class="layout-content" >
    <router-view/>
    </a-layout-content>
    <a-layout-footer class="layout-footer">
    </a-layout-footer>
    <Login ref="loginRef"/>
  </a-layout>
</template>

<style scoped>
.layout-header {
  color: #d47777;
  height: max-content;
  padding-inline: 0px;
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
  background-color: #fffdfd;
}

.layout-content{
  text-align: center;
  line-height: 120px;
  position: relative;
  color: #d47777;
  background-color: transparent;
}

.layout-footer{
  text-align: center;
  color: #d47777;
  background-color: transparent;
}

.layout-content-menu-item{
  border: 1px solid #e8d4d4;
  float: right;
  cursor: pointer;
  margin: 5px 6px;
}
.layout-content-logo-item{
  border: 1px solid #e8d4d4;
  cursor: pointer;
  margin: 20px 6px;
  width: 44px;
}
.main-container{
  height: 100%;
  width: 100%;
  position: absolute;
}
</style>
