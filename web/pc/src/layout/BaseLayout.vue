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
  <div class="layout">
    <div class="layout-header">
      <a-row style="width: 100%"  >
        <a-col :md="8" :sm="24" :xs="24" >
          <a-image
              @click="goToAbout('/index')"
              :preview="false"
              :width="200"
              :height="44"
              class="layout-content-logo-item"
              src="https://www.antdv.com/assets/logo.1ef800a8.svg"
          />
        </a-col>
        <a-col  :md="8" :sm="24" :xs="24">
        </a-col>
        <a-col  :md="8" :sm="24" :xs="24">
          <a-avatar :size="44" class="layout-content-menu-item" @click="openDialog" >登陆</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="getInfoData" >杂谈</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="getInfoData" >留言</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/about')" >关于</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/travel')" >旅行</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/blog')" >日记</a-avatar>
        </a-col>
      </a-row>
    </div>
    <router-view class="layout-content" />
    <div class="layout-footer">
      我是底部
    </div>
    <Login ref="loginRef"/>
  </div>
</template>

<style scoped>
.layout {
  height: 100%;
  width: 100%;
  overflow: auto;
  background: transparent  url("src/assets/image/Sara11712577366099546.jpg");
  background-size: cover;
  display: flex;
  flex-direction: column;
}
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
  justify-content: center;
  color: #d47777;
  flex:1;
  /*min-height: calc(100vh - 104px);*/
  background-color: transparent;
}

.layout-footer{
  text-align: center;
  color: #d47777;
  height: 50px;
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
}
</style>
