<template>
  <a-row style="background-color: transparent;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24"  >
      <span/>
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="text-align: center;    background: linear-gradient(rgb(241 241 241), rgb(173 215 197));" >
      <a-form :label-col="labelCol"  :wrapper-col="wrapperCol"  ref="formBlogRef" :model="formState"   autocomplete="off"  class="blog-card-edit">
        <a-form-item label="title" :rules="[{ required: true, message: '请输入博客标题' }]">
          <a-input  v-model:value="formState.title"   placeholder="请输入博客标题"/>
        </a-form-item>
        <a-form-item label="description">
          <a-textarea v-model:value="formState.description"   placeholder="请输入博客简介"/>
        </a-form-item>
        <a-row  >
          <a-col :md="12" :sm="24" :xs="24"   >
            <a-form-item label="isHidden">
              <a-switch :checkedValue="1" :unCheckedValue="0" v-model:checked="formState.isHidden" />
            </a-form-item>
          </a-col>
          <a-col :md="12" :sm="24" :xs="24"   >
            <a-form-item label="tags"  :rules="[{ required: true, message: '请输入至少一个标签' }]">
              <a-select
                  v-model:value="formState.tags"
                  mode="tags"
                  class="tag-class"
                  placeholder="请输入标签,方便查询"
                  @change="handleChange"
              ></a-select>
            </a-form-item>
          </a-col>
        </a-row>
        <a-form-item label="cover">
          <a-upload action="/upload.do" :maxCount="1" list-type="picture-card" >
            <div>
              <PlusOutlined />
              <div style="margin-top: 8px">Upload</div>
            </div>
          </a-upload>
        </a-form-item>

        <a-form-item label="content"  :rules="[{ required: true, message: '请输入博客内容' }]">
          <div style="border: 1px solid #ccc;background-color: white;min-height: 40vh" id="editor" >
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
import "quill/dist/quill.snow.css";
import "quill/dist/quill.core.css";
import "quill/dist/quill.bubble.css";
import { onBeforeUnmount,  shallowRef, onMounted ,reactive, toRaw, ref } from 'vue'
import Quill from 'quill';
import { blogUpdate,blogCreate, blogGet} from "@/api/blog";
const editorRef = shallowRef()
import { useRouter,useRoute } from 'vue-router';
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}



const handleChange = (value) => {
  console.log(`selected ${value}`);
};



const formState = reactive({
  id: 0,
  title: '',
  description: '',
  isHidden: 0,
  tags:[],
  content: '',
  cover: '',
});
const formBlogRef = ref(null)
const toolbarOptions = [
  ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
  ['blockquote', 'code-block'],
  ['link', 'image', 'video', 'formula'],

  [{ 'header': 1 }, { 'header': 2 }],               // custom button values
  [{ 'list': 'ordered'}, { 'list': 'bullet' }, { 'list': 'check' }],
  [{ 'script': 'sub'}, { 'script': 'super' }],      // superscript/subscript
  [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent
  [{ 'direction': 'rtl' }],                         // text direction

  [{ 'size': ['small', false, 'large', 'huge'] }],  // custom dropdown
  [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

  [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
  [{ 'font': [] }],
  [{ 'align': [] }],

  ['clean']                                         // remove formatting button
];

const options = {
  modules: {
    toolbar: toolbarOptions,
  },
  placeholder: '请输入内容',
  theme: 'snow'
};

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editorRef.value = null
})

onMounted(function (){
  const quill = new Quill('#editor',options);  // First matching element will be used
  if(formBlogRef.value){
    formBlogRef.value.resetFields();
  }
  editorRef.value = quill
  quill.on('editor-change', (eventName, ...args) => {
    if (eventName === 'text-change') {
      formState.content = quill.getSemanticHTML()
    } else if (eventName === 'selection-change') {
      // args[0] will be old range
      formState.content= quill.getSemanticHTML()
    }
  });
  // quill.on('text-change', (delta, oldDelta, source) => {
  //   if (source == 'api') {
  //     console.log('An API call triggered this change.');
  //     formState.Content = delta
  //   } else if (source == 'user') {
  //     console.log('A user action triggered this change.');
  //     formState.Content = delta
  //   }
  // });
  console.log("here, on Mounted")
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  blogGet({ID:id}).then(res=>{
    if(res&&res.code===200){
      formState.content = res.data.content
      quill.clipboard.dangerouslyPasteHTML(res.data.content);
      const obj = res.data.header
      formState.cover = obj.cover
      formState.tags = obj.tags
      formState.id = id
      formState.title = obj.title
      formState.description = obj.description
      formState.isHidden = obj.isHidden-0
    }
  })
})
const confirmLoading = ref(false);
const onSubmit = () => {
  formBlogRef.value
      .validate()
      .then(() => {
        confirmLoading.value = true
       if(formState.id>0){
         blogUpdate(formState).then(res=>{
             if(res&&res.code===200){
             toRoute('/blog')
           }
         }).finally(()=>{
           confirmLoading.value = false
         })
       }else{
         blogCreate(formState).then(res=>{
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
const labelCol = { style: { width: '80px' } };
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
