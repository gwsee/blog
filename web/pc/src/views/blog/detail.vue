<template>
  <a-row style="background-color: transparent;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24"  >
      <span/>
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  class=" mb-12" style="text-align: unset;" :loading="loading">
      <a-spin :spinning="loading">
      <a-layout class="blog-post" style="background: linear-gradient(rgb(241 241 241), rgb(173 215 197));">
        <a-layout-content>
          <a-typography-title :level="1" class="blog-title">
            {{ formState.title }}
          </a-typography-title>

          <div class="post-meta">
            <a-space>
              <a-tag v-for="tag in formState.tags" :key="tag" color="cyan"  :bordered="false">{{ tag }}</a-tag>
            </a-space>
            <a-typography-text type="secondary">
              最后更新时间:  {{ $formatDate(formState.createdAt) }}
            </a-typography-text>
          </div>

          <a-divider />

          <div class="author-info" v-if="false">
            <a-space align="center">
              <a-avatar :src="formState.avatar||avatar" :size="64" />
              <div>
                <a-typography-title :level="4" :style="{ margin: 0 }">
                  {{ formState.author||'Gwsee'  }}
                </a-typography-title>
                <a-typography-text type="secondary">
                  {{ formState.author||'2024-10-12' }}
                </a-typography-text>
              </div>
            </a-space>
          </div>

          <a-divider  v-if="false" />

          <a-typography-paragraph class="post-summary" v-if="formState.description">
            &nbsp;&nbsp;&nbsp;&nbsp;{{ formState.description }}
          </a-typography-paragraph>

          <a-divider  v-if="formState.description"/>

          <div class="post-content" v-if="formState.content" v-html="formState.content"></div>
        </a-layout-content>
      </a-layout>
      <a-card hoverable v-if="false">
        <template #cover>
          <img
              alt="example"
              src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
          />
        </template>
        <template #actions>
          <setting-outlined key="setting" />
          <edit-outlined key="edit" />
          <ellipsis-outlined key="ellipsis" />
        </template>
        <a-card-meta :title="formState.title" :description="formState.description">
          <template #avatar>
            <a-avatar src="https://joeschmoe.io/api/v1/random" />
          </template>
        </a-card-meta>
        <div  v-html="formState.content" style="overflow-x: hidden">

        </div>
      </a-card>
      </a-spin>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  >
     <span/>
    </a-col>
  </a-row>
</template>

<script setup>
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { onBeforeUnmount,  shallowRef, onMounted ,reactive, toRaw, ref } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { blogUpdate,blogCreate, blogGet} from "@/api/blog";
const editorRef = shallowRef()
import { useRouter,useRoute } from 'vue-router';
const router = useRouter();

const toRoute=(path)=> {
  router.push(path)
}
const mode = ref("default")
const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }
const loading = ref(false);
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleChange = (value) => {
  console.log(`selected ${value}`);
};

const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

const formState = reactive({
  id: 0,
  title: '',
  description: '',
  isHidden: 0,
  tags:[],
  content: '',
  cover: '',
  createdAt:0,
});
const formBlogRef = ref(null)

onMounted(function (){
  if(formBlogRef.value){
    formBlogRef.value.resetFields();
  }
  console.log("here, on Mounted")
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  loading.value = true;
  blogGet({id:id}).then(res=>{
    if(res&&res.code===200){
      formState.content = res.data.content
      const obj = res.data.header
      formState.cover = obj.cover
      formState.tags = obj.tags
      formState.id = id
      formState.title = obj.title
      formState.description = obj.description
      formState.isHidden = obj.isHidden
      formState.createdAt = obj.createdAt
    }
  }).finally(()=>{
    loading.value = false;
  })
})
const confirmLoading = ref(false);
const onSubmit = () => {
  formBlogRef.value
      .validate()
      .then(() => {
        confirmLoading.value = true
        if(formState.Id>0){
          blogUpdate(formState).then(res=>{
            console.log(res,"....")
            if(res&&res.code===200){
              toRoute('/blog')
            }
          }).finally(()=>{
            confirmLoading.value = false
          })
        }else{
          blogCreate(formState).then(res=>{
            console.log(res,"....")
            if(res&&res.code===200){
              toRoute('/blog')
            }
          }).finally(()=>{
            confirmLoading.value = false
          })
        }
      })
      .catch(error => {
        console.log('error', error);
      });
};
const labelCol = { style: { width: '100px' } };
const wrapperCol = { span: 24 };

</script>

<style scoped>
.blog-post {
  max-width: 100%;
  width: fit-content;
  margin: 0 auto;
  padding: 24px;
  min-height: 50vh;
  min-width: 100%;
  background: linear-gradient(135deg, #f5f7fa 0%, #e0e6ed 100%);
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid #e8e8e8;
  position: relative;
  overflow: hidden;
}
.blog-post::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, rgba(233, 242, 248, 0.5) 0%, rgba(197, 239, 203, 0.5) 100%);
}

.blog-post:hover {
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
  transform: translateY(-5px);
}

.blog-title {
  margin-bottom: 16px !important;
  color: #2c3e50;
}

.post-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.author-info {
  margin: 24px 0;
  padding: 16px;
  background-color: rgba(255, 255, 255, 0.7);
  border-radius: 8px;
}

.post-summary {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 24px;
  padding: 16px;
  background-color: rgba(255, 255, 255, 0.7);
  border-radius: 8px;
  border-left: 4px solid rgba(79,172,254,0.3);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.post-content {
  font-size: 16px;
  line-height: 1.8;
  background-color: rgba(255, 255, 255, 0.9);
  padding: 24px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.post-content h2 {
  font-size: 24px;
  margin-top: 32px;
  margin-bottom: 16px;
  color: #34495e;
}

.post-content p {
  margin-bottom: 16px;
}

.post-content pre {
  background-color: #f8f8f8;
  padding: 16px;
  border-radius: 4px;
  overflow-x: auto;
  border: 1px solid #e8e8e8;
}

@media (max-width: 768px) {
  .blog-post {
    padding: 20px;
    margin: 10px;
  }
}

</style>
