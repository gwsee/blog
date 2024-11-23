<template>
  <a-row style="width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" >
      <span></span>
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="text-align: center;position: relative;    background: linear-gradient(rgb(241 241 241), rgb(173 215 197));" >
      <a-input-group compact  style="margin-top:30px;">
        <a-select   mode="tags"
                    max-tag-count="responsive"
                    :size="'large'" placeholder="请输入需要查询的标签!"
                    class="tag-class" v-model:value="Tags" style="width: 20%;height: 40px;" >
        </a-select>
        <a-input v-model:value="Title" placeholder="请输入需要查询的内容!"
                 :style="state.user&&state.user.ID?`width: calc(70% - 140px);height: 40px;`:`width: calc(70% - 70px);height: 40px;`" />
        <a-button type="default" :style="`height: 40px;width: 70px`" @click="onQuery">查询</a-button>
        <a-button type="default" :style="`height: 40px;width: 70px`" @click="toRoute('/blog/manage/0')" v-if="state.user&&state.user.ID" >新增</a-button>
      </a-input-group>
      <a-card :title="item.Title" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}" v-for="(item,key) in data">
        <template #extra><a @click="toRoute('/blog/manage/'+item.Id)" v-if="state.user&&state.user.ID === item.AccountId">编辑</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;" :src="item.cover||defaultCover">
          </a-avatar>
          <div class="blog-content-eclipses" @click="toRoute('/blog/detail/'+item.Id)">
           &nbsp; &nbsp;  {{item.Description||'这个人很烂什么都没写'}}
            <div style="bottom: 2px;position: absolute">
              &nbsp; &nbsp;  <a-tag v-for="tag in item.Tags" :key="tag" color="cyan"  :bordered="false">{{ tag }}</a-tag>
            </div>
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <div class="blog-card" v-if="total>pageSize">
        <a-pagination style="float: right" v-model:current="current" :total="total" :page-size-options="pageSizeOptions"  v-model:pageSize="pageSize" />
      </div>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"   >
      <span></span>
    </a-col>
  </a-row>
</template>

<script setup>
import {UserOutlined, HeartOutlined, LikeOutlined, DislikeOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router';
import {blogList} from "@/api/blog";
import {onMounted, ref,watch} from "vue";
import defaultCover from "@/assets/image/default-cover.jpg"
import { useAuthStore  } from '@/store/auth'
const { state } = useAuthStore();
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}
const pageSizeOptions = ref(['5','10', '20', '30', '40', '50']);
const current = ref(1)
const total = ref(0)
const pageSize = ref(5)
const data = ref([])
const loading = ref(false)
const Title = ref("")
const Tags = ref([])
onMounted( ()=>{
  list()
})

watch(pageSize, () => {
  if(current.value===1){
    list()
  }else{
    current.value = 1
  }
});
watch(current, () => {
  list()
});
const list =()=>{
  loading.value=true
  blogList({"Page":{"PageNum":current.value,"PageSize":pageSize.value},"Title":Title.value,"Tags":Tags.value}).then(res=>{
    if(res&&res.code===0){
     data.value = res.data.List
     total.value = res.data.Total-0
    }
  }).finally(()=>{
    loading.value = false
  })
}
const onQuery=()=>{
  if(current.value===1){
    list()
  }else{
    current.value =1
  }
}
</script>

<style scoped>
.blog-content-eclipses{
  flex: 1; /* 右边部分自适应 */
  margin: 15px 15px 15px 0;
  display: -webkit-box;
  color: gray;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 显示3行 */
  overflow: hidden;
  position: relative;
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
