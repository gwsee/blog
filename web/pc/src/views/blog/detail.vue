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

          <a-typography-paragraph class="post-summary" v-if="formState.description" >
          <div v-html="formattedContent(formState.description)"></div>
          </a-typography-paragraph>

          <a-divider  v-if="formState.description"/>
          <div class="preview-wrapper">
            <div class="post-content ql-snow">
              <div class="ql-editor" v-html="formState.content"></div>
            </div>
          </div>

        </a-layout-content>
      </a-layout>
      </a-spin>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  >
     <span/>
    </a-col>
  </a-row>
</template>

<script setup>
import "quill/dist/quill.snow.css";
import "quill/dist/quill.core.css";
import "quill/dist/quill.bubble.css";
import { onBeforeUnmount,  shallowRef, onMounted ,reactive, ref } from 'vue'
import {  blogGet} from "@/api/blog";
const editorRef = shallowRef()
import { useRoute } from 'vue-router';
const loading = ref(false);
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const formattedContent=(data)=> {
  if(!data){
    return data
  }
  return data.replace(/\n/g, '<br/>');
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


@media (max-width: 768px) {
  .blog-post {
    padding: 20px;
    margin: 10px;
  }
}
.preview-wrapper {
  border: 1px solid #ccc;
  border-radius: 4px;
}
/* 确保预览区域的 ql-editor 不会被 Quill 的默认高度限制 */
.preview-wrapper .ql-editor {
  min-height: auto;
  max-height: none;
}
</style>
