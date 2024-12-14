<script setup>
import Login from "@/components/Login.vue";
import avatar from "@/assets/image/default-avatar.png"
import logo from "@/assets/image/logo.png"
import {onMounted, ref, watch} from "vue"
import { useRouter } from 'vue-router';
import { useAuthStore  } from '@/store/auth.js'
const { isLoggedIn, state, setLoginData, getLoginToken } = useAuthStore();
const router = useRouter();
const loginRef = ref(null);

const showLoginDialog =()=>{
  loginRef.value.show()
}

onMounted(function () {
  if(isLoggedIn){
    const tk = getLoginToken()
    if(!tk||tk=='undefined'){
     return
    }
    setLoginData(tk)
  }
})

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
              :src="logo"
          />
        </a-col>
        <a-col  :md="8" :sm="24" :xs="24">
        </a-col>
        <a-col  :md="8" :sm="24" :xs="24">
          <a-avatar :size="44" class="layout-content-menu-item" @click="showLoginDialog" v-if="!isLoggedIn" >登陆</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" :src="avatar" v-else @click="goToAbout('/about')"></a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/picture')">照片墙</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/travel')" >旅行</a-avatar>
          <a-avatar :size="44" class="layout-content-menu-item" @click="goToAbout('/blog')" >日记</a-avatar>
        </a-col>
      </a-row>
    </div>
    <router-view class="layout-content" />
    <div class="layout-footer">
      <span/>
    </div>
    <Login ref="loginRef"/>
  </div>
</template>

<style scoped>
.layout {
  height: 100%;
  width: 100%;
  overflow: auto;
  background: transparent  url("/src/assets/image/bg02.gif");
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
