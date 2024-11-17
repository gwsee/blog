<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929;text-align: center;position: relative" >
      <a-input-group compact  style="margin-top:30px;">
        <a-input v-model:value="value" placeholder="you can search blog title from here!"
                 style="width: calc(90% - 140px);height: 40px;background-color: transparent" />
        <a-button type="default" style="height: 40px;width: 70px" @click="onQuery">Query</a-button>
        <a-button type="default" style="height: 40px;width: 70px" @click="toRoute('/blog/edit/0')">New</a-button>
      </a-input-group>
      <a-card :title="item.Title" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}" v-for="(item,key) in data">
        <template #extra><a @click="toRoute('/blog/edit/'+item.Id)">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;" :src="item.cover||defaultCover">
          </a-avatar>
          <div class="blog-content-eclipses">
           &nbsp; &nbsp;  {{item.Description||'这个人很烂什么都没写'}}
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <div class="blog-card" style="bottom: 0;position: absolute">
        <a-pagination style="float: right" v-model:current="current" :total="total" />
      </div>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  style="background-color: #1a1a1a" >
      3
    </a-col>
  </a-row>
</template>

<script setup>
import {UserOutlined, HeartOutlined, LikeOutlined, DislikeOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router';
import {blogList} from "@/api/blog";
import {onMounted, ref} from "vue";
import defaultCover from "@/assets/image/default-cover.jpg"
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}
const current = ref(1)
const total = ref(0)
const data = ref([])
const loading = ref(false)

onMounted( ()=>{
  list()
})
const list =()=>{
  loading.value=true
  blogList({"Page":{"PageNum":1,"PageSize":1}}).then(res=>{
    if(res&&res.code===0){
     data.value = res.data.List
     total.value = res.data.Total-0
    }
  }).finally(()=>{
    loading.value = false
  })
}
</script>

<style scoped>
.blog-content-eclipses{
  flex: 1; /* 右边部分自适应 */
  margin: 15px 15px 15px 0;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 显示3行 */
  overflow: hidden;
  font-size: 15px;
  text-overflow: ellipsis;
}

.blog-card{
  width: 90%;
  margin-top: 20px;
  text-align: left;
  margin-left: auto;
  margin-right: auto;
}
</style>
