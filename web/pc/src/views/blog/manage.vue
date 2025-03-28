<template>
  <a-row style="background-color: transparent;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24"  >
      <span/>
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  class="rounded-xl p-4 sm:p-6  mb-12"  style="text-align: center;    background: linear-gradient(rgb(241 241 241), rgb(173 215 197));" >
      <a-form :label-col="labelCol"  :wrapper-col="wrapperCol"
              ref="formBlogRef" :model="formState"
              class="blog-card-edit">
        <a-form-item label="title" name="title" :rules="[{ required: true, message: '请输入博客标题' }]">
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
            <a-form-item label="tags"  name="tags"  :rules="[{ required: true, type: 'array', message: '请输入至少一个标签' }]">
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
          <a-upload  :max-count="1"
                     :multiple="false"
                     list-type="picture-card"
                     v-model:fileList="formState.coverList"
                    :customRequest="$customRequest"
                    :before-upload="beforeUpload"
                    @change="handleChangeCover"
                    @preview="handlePreview"

          >
            <div v-if="formState.coverList<1">
              <PlusOutlined />
              <div style="margin-top: 8px">上传封面</div>
            </div>
          </a-upload>
        </a-form-item>

        <a-form-item label="content"  name="content"  :rules="[{ required: true, message: '请输入博客内容' }]">
          <div style="border: 1px solid #ccc;background-color: white;min-height: 40vh" ref="quillEditor" />
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 5 }" style="text-align: center">
          <a-button type="primary" @click="onSubmit" >保存</a-button>
          <a-button style="margin-left: 10px" @click="toRoute('/blog')">取消</a-button>
        </a-form-item>
      </a-form>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"   >
      <span/>
    </a-col>
    <!-- 图片预览模态框 -->
    <a-modal :open="previewVisible" :footer="null" @cancel="handleCancel">
      <img alt="example" style="width: 100%" :src="previewImage" />
    </a-modal>
  </a-row>
</template>

<script setup>
import { onBeforeUnmount, shallowRef, onMounted, reactive, toRaw, ref, getCurrentInstance} from 'vue'
import {customRequest} from "@/utils/util.js";
import { useRouter,useRoute } from 'vue-router';
import {message} from "ant-design-vue";
import { PlusOutlined } from '@ant-design/icons-vue'
import { blogUpdate,blogCreate, blogGet} from "@/api/blog";
import Quill from 'quill';
import "quill/dist/quill.snow.css";
import "quill/dist/quill.core.css";
import "quill/dist/quill.bubble.css";

let quill = null;
const quillEditor = ref(null);
// const editorRef = shallowRef()
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
const router = useRouter();
const formState = reactive({
  id: 0,
  title: '',
  description: '',
  isHidden: 0,
  coverList:[],
  tags:[],
  content: '',
  cover: '',
  files: [],
});
const formBlogRef = ref(null)

const toRoute=(path)=> {
  router.push(path)
}
const handleChange = (value) => {
  console.log(`selected ${value}`);
};

const uploadFile = async (file) => {
  // 这里应该是您的实际文件上传逻辑
  return await customRequest({file})
};

const insertImage = async (file) => {
  const range = quill.getSelection(true);
  const insertIndex = range.index;
  const loadingText = '正在上传图片...';
  quill.insertText(range.index,loadingText);
  try {
   const {fileUrl}= await uploadFile(file);
    quill.deleteText(insertIndex, loadingText.length);
    quill.insertEmbed(insertIndex, 'image', fileUrl);
    quill.setSelection(insertIndex + 1);
  } catch (error) {
    console.error('图片上传失败', error);
    quill.deleteText(insertIndex, loadingText.length);
    const errorText = '图片上传失败';
    quill.insertText(insertIndex, errorText, 'bold', true);
    setTimeout(() => {
      quill.deleteText(insertIndex, errorText.length);
    }, 3000);
  }
};

const insertVideo = async (file) => {
  const range = quill.getSelection(true);
  quill.insertText(range.index, '正在上传视频...');
  try {
    const { fileUrl } = await uploadFile(file);
    quill.deleteText(range.index, '正在上传视频...'.length);
    quill.insertEmbed(range.index, 'video', fileUrl);
    // quill.clipboard.dangerouslyPasteHTML(range.index, `<video src="`+fileUrl+`" controls style="max-width: 100%;"></video>`);
    quill.setSelection(range.index + 1);
  } catch (error) {
    console.error('视频上传失败', error);
    quill.deleteText(range.index, '正在上传视频...'.length);
  }
};
const insertFile = async (file) => {
  const range = quill.getSelection(true);
  quill.insertText(range.index, '正在上传文件...');
  try {
    const { fileName, fileUrl } = await uploadFile(file);
    quill.deleteText(range.index, '正在上传文件...'.length);
    const fileLink = `<a href="${fileUrl}" target="_blank">${fileName}</a>`;
    quill.clipboard.dangerouslyPasteHTML(range.index, fileLink);
    quill.setSelection(range.index + 1);
  } catch (error) {
    console.error('文件上传失败', error);
    quill.deleteText(range.index, '正在上传文件...'.length);
  }
};

const imageHandler = () => {
  const input = document.createElement('input');
  input.setAttribute('type', 'file');
  input.setAttribute('accept', 'image/*');
  input.click();

  input.onchange = () => {
    const file = input.files[0];
    if (file) {
      insertImage(file);
    }
  };
};
const videoHandler = () => {
  const input = document.createElement('input');
  input.setAttribute('type', 'file');
  input.setAttribute('accept', 'video/*');
  input.click();

  input.onchange = () => {
    const file = input.files[0];
    if (file) {
      insertVideo(file);
    }
  };
};
const fileHandler = () => {
  const input = document.createElement('input');
  input.setAttribute('type', 'file');
  input.click();

  input.onchange = () => {
    const file = input.files[0];
    if (file) {
      insertFile(file);
    }
  };
};
const pasteHandler = (e) => {
  const clipboard = e.clipboardData || window.clipboardData;
  if (clipboard && clipboard.items) {
    for (let i = 0; i < clipboard.items.length; i++) {
      const item = clipboard.items[i];
      if (item.type.indexOf('image') !== -1) {
        e.preventDefault();
        const file = item.getAsFile();
        insertImage(file);
        return;
      } else if (item.type.indexOf('video') !== -1) {
        e.preventDefault();
        const file = item.getAsFile();
        insertVideo(file);
        return;
      }
    }
  }
};
const onReadyQuill =  () => {
  const options = {
    modules: {
      toolbar: {
        container: [
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

          ['clean'],                                         // remove formatting
          ['custom-image-resize'],
        ],
        handlers: {
          'custom-image-resize': customImageResize,
          image: imageHandler,
          video: videoHandler,
          file: fileHandler
        },
        imageResize: {
          displaySize: true
        }
      },
    },
    placeholder: '请输入内容...',
    theme: 'snow'
  };
  quill = quill = new Quill(quillEditor.value,options);  // First matching element will be used
  const toolbar = quill.getModule('toolbar');
  toolbar.addHandler('custom-image-resize', customImageResize);
  quill.on('editor-change', (eventName, ...args) => {
    if (eventName === 'text-change') {
      formState.content =  quillEditor.value.getHTML()
    } else if (eventName === 'selection-change') {
      // args[0] will be old range
      formState.content=  quillEditor.value.getHTML()
    }
  });
  quill.root.addEventListener('paste', pasteHandler);
  quill.root.addEventListener('click', handleImageClick);
  return quill;
}
const customImageResize = () => {
  const range = quill.getSelection();
  if (range) {
    const [leaf] = quill.getLeaf(range.index);
    if (leaf.domNode && leaf.domNode.tagName === 'IMG') {
      const img = leaf.domNode;
      const newWidth = prompt('Enter new width (in pixels):', img.width);
      if (newWidth && !isNaN(newWidth)) {
        img.width = newWidth;
        img.height = 'auto'; // Maintain aspect ratio
      }
    }
  }
};

const handleImageClick = (event) => {
  if (event.target && event.target.tagName === 'IMG') {
    const img = event.target;
    const newWidth = prompt('Enter new width (in pixels):', img.width);
    if (newWidth && !isNaN(newWidth)) {
      img.width = newWidth;
      img.height = 'auto'; // Maintain aspect ratio
    }
  }
};

onMounted(function (){
  if(formBlogRef.value){
    formBlogRef.value.resetFields();
  }
  onReadyQuill()
  const route = useRoute();
  let id = route.params&&route.params.id;
  id = id - 0
  if(!id){
    return
  }
  const { appContext } = getCurrentInstance();
  const $fileFull = appContext.config.globalProperties.$fileFull;
  blogGet({id:id}).then(res=>{
    if(res&&res.code===200){
      formState.content = res.data.content
      quill.clipboard.dangerouslyPasteHTML(res.data.content);
      //quill.root.innerHTML = formState.content//文档提供的 有bug
      const obj = res.data.header
      formState.cover = obj.cover
      formState.tags = obj.tags
      formState.id = id
      formState.title = obj.title
      if(obj.cover!==""){
        formState.coverList = [{uuid:obj.cover,url:$fileFull(obj.cover)}]
      }
      formState.description = obj.description
      formState.isHidden = obj.isHidden-0
    }
  })
})
onBeforeUnmount(() => {
  if (quill) {
    quill.off('text-change');
    quill.root.removeEventListener('paste', pasteHandler);
    quill.root.removeEventListener('click', handleImageClick);
    quill = null;
  }
});
const confirmLoading = ref(false);
const onFinish = (values) => {
  console.log(values)
}
const onSubmit = () => {
  formBlogRef.value
      .validate()
      .then(() => {
        if(formState.coverList.length>0){
          formState.cover = formState.coverList[0].uuid;
        }
        confirmLoading.value = true
        console.log(formState,"formState....")
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


const handlePreview = async (file) => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64(file.originFileObj);
  }
  previewImage.value = file.url || file.preview;
  previewVisible.value = true;
};
const previewVisible = ref(false);
const previewImage = ref('');
const handleCancel = () => {
  previewVisible.value = false;
};
const beforeUpload = (file) => {
  const isImage = file.type.split("/")[0]==='image'
  if (!isImage) {
    console.log(file)
    message.error('只能上传 图片 文件!');
    return false
  }
  const isLt20M = file.size / 1024 / 1024 < 20;
  if (!isLt20M) {
    message.error('图片必须小于 20MB!');
    return false
  }
  return true;
};
const handleChangeCover = (info) => {
  if (info.file&&info.file.status === "uploading"){
    return false
  }
  let resFileList = [...info.fileList];
  const includes = [];
  resFileList = resFileList.filter(function (item) {
    if(item.status==='error'){
      return false
    }
    let uuid = item.uuid||item.response.uuid
    let url = item.url||item.response.url
    if(!uuid){
      return false
    }
    if(!includes.includes(uuid)){
      item.uuid = uuid
      item.url = url
      includes.push(item.uuid)
      return true
    }
    return false
  })
  if(resFileList.length>0){
    formState.coverList = [resFileList[resFileList.length-1]]
    return
  }
  formState.coverList = resFileList
};

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

:deep(.ql-custom-image-resize) {
  &::after {
    content: "Resize";
  }
}

:deep(.ql-editor img) {
  cursor: pointer;
}
</style>
