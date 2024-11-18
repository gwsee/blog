<template>
  <a-row style="background-color: transparent;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24"  >
      <span/>
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="text-align: center" >
      <a-form :label-col="labelCol"   ref="formBlogRef" :model="formState"   autocomplete="off"  class="blog-card-edit" :wrapper-col="wrapperCol">
        <a-form-item label="Title" :rules="[{ required: true, message: '请输入博客标题' }]">
          <a-input  v-model:value="formState.Title"   placeholder="请输入博客标题"/>
        </a-form-item>
        <a-form-item label="Description">
          <a-textarea v-model:value="formState.Description"   placeholder="请输入博客简介"/>
        </a-form-item>
        <a-row  >
          <a-col :md="12" :sm="24" :xs="24"   >
            <a-form-item label="IsHidden">
              <a-switch :checkedValue="1" :unCheckedValue="0" v-model:checked="formState.IsHidden" />
            </a-form-item>
          </a-col>
          <a-col :md="12" :sm="24" :xs="24"   >
            <a-form-item label="Tags"  :rules="[{ required: true, message: '请输入至少一个标签' }]">
              <a-select
                  v-model:value="formState.Tags"
                  mode="tags"
                  class="tag-class"
                  placeholder="请输入标签,方便查询"
                  @change="handleChange"
              ></a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-form-item label="Cover">
          <a-upload action="/upload.do" :maxCount="1" list-type="picture-card">
            <div>
              <PlusOutlined />
              <div style="margin-top: 8px">Upload</div>
            </div>
          </a-upload>
        </a-form-item>

        <a-form-item label="Content"  :rules="[{ required: true, message: '请输入博客内容' }]">
          <div style="border: 1px solid #ccc">
            <Toolbar
                style="border-bottom: 1px solid #ccc"
                :editor="editorRef"
                :defaultConfig="toolbarConfig"
                :mode="mode"
            />
            <Editor

                style="height: 500px; overflow-y: hidden;"
                v-model="formState.Content"
                :defaultConfig="editorConfig"
                :mode="mode"
                @onCreated="handleCreated"
            />
          </div>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 5 }" style="text-align: center">
          <a-button type="primary" @click="onSubmit">保存</a-button>
          <a-button style="margin-left: 10px" @click="toRoute('/blog')">取消</a-button>
        </a-form-item>
      </a-form>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"   >
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
  Id: 0,
  Title: '',
  Description: '',
  IsHidden: 0,
  Tags:[],
  Content: '',
  Cover: '',
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
  blogGet({Id:id}).then(res=>{
    if(res&&res.code===0){
      formState.Content = res.data.Content
      const obj = res.data.Header
      formState.Cover = obj.Cover
      formState.Tags = obj.Tags
      formState.Id = id
      formState.Title = obj.Title
      formState.Description = obj.Description
      formState.IsHidden = obj.IsHidden-0
    }
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
           if(res&&res.code===0){
             toRoute('/blog')
           }
         }).finally(()=>{
           confirmLoading.value = false
         })
       }else{
         blogCreate(formState).then(res=>{
           console.log(res,"....")
           if(res&&res.code===0){
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

.blog-card-edit{
  width: 90%;
  margin-top: 20px;
  text-align: left;
  margin-left: auto;
  margin-right: auto;
}
.tag-class{
}
</style>
