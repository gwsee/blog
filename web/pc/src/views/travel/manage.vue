<template>
  <div class="min-h-screen  py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto" style=" background: linear-gradient(rgb(241 241 241), rgb(173 215 197));min-height:100%">
      <h1 class="text-3xl font-bold text-gray-900 p-8 text-center">记录一次有趣的旅行</h1>
      <a-form

          :model="formState"
          @finish="onFinish"
          ref="formTravelRef"
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          style="width: 90%;margin-left: auto;margin-right:auto"
      >
        <!-- 标题 -->
        <a-form-item
            name="title"
            label="标题"
            :rules="[{ required: true, message: '请输入标题' }]"
        >
          <a-input v-model:value="formState.title" placeholder="请输入旅行标题" />
        </a-form-item>

        <!-- 简介 -->
        <a-form-item
            name="description"
            label="简介"
            :rules="[{ required: true, message: '请输入简介' }]"
        >
          <a-textarea
              v-model:value="formState.description"
              :rows="4"
              placeholder="请输入旅行简介"
          />
        </a-form-item>
        <a-form-item name="是否隐藏" label="isHidden">
          <a-switch :checkedValue="true" :unCheckedValue="false" v-model:checked="formState.isHidden" />
        </a-form-item>
        <!-- 照片上传 -->
        <a-form-item
            name="photoList"
            label="旅行照片"
            :rules="[{ required: true, message: '请上传至少一张照片' }]"
        >
          <a-upload
              v-model:fileList="formState.photoList"
              list-type="picture-card"
              :max-count="66"
              :multiple="true"
              :customRequest="customRequestPhoto"
              :before-upload="beforeUpload"
              @change="handleChangePhoto"
              @preview="handlePreview"
          >
            <div v-if="formState.photoList.length < 66">
              <plus-outlined />
              <div class="mt-2">上传照片</div>
            </div>
          </a-upload>
        </a-form-item>

        <!-- 视频上传 -->
        <a-form-item name="video" label="视频">
          <a-upload
              v-model:fileList="formState.videoList"
              :customRequest="customRequestVideo"
              :max-count="1"
              :multiple="false"
              :before-upload="beforeUploadVideo"
              @change="handleChangeVideo"
          >
            <a-button type="default"> 上传视频</a-button>
          </a-upload>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 5 }" style="text-align: center;margin-bottom: 10px">
          <a-button type="primary" html-type="submit" :loading="loading" :disabled="loading">保存</a-button>
          <a-button style="margin-left: 10px" @click="toRoute('/travel')" :loading="loading" :disabled="loading">取消</a-button>
        </a-form-item>
      </a-form>

      <!-- 图片预览模态框 -->
      <a-modal :open="previewVisible" :footer="null" @cancel="handleCancel">
        <img alt="example" style="width: 100%" :src="previewImage" />
      </a-modal>
    </div>
  </div>
</template>

<script setup>
import {ref, reactive, onMounted} from 'vue';
import {message} from 'ant-design-vue';
import {PlusOutlined, UploadOutlined} from '@ant-design/icons-vue';
import {travelSave,travelGet} from "@/api/travel";
import {fileUpload,filePrefix} from "@/api/tool";
const formState = reactive({
  id: 0,
  title: '',
  isHidden: false,
  description: '',
  photoList: [],
  photo:[],
  videoList: [],
  video:'',
});
const formTravelRef = ref(null);

onMounted(function (){
  if(formTravelRef.value){
    formTravelRef.value.resetFields();
  }
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  loading.value = true;
  travelGet({id:id}).then(res=>{
    if(res&&res.code===200){
      console.log(res.data)
      let obj = res.data
      formState.id = id
      formState.title = obj.title
      formState.description = obj.description
      formState.isHidden = obj.isHidden
      let photoList = []
      let videoList = []
      if(obj.video){
        videoList.push({uuid:obj.video,url:filePrefix+obj.video,name:obj.video})
      }
      obj.photos.forEach((item)=>{
        const obj = {uuid:item,url:filePrefix+item}
        photoList.push(obj)
      })
      formState.photoList = photoList
      formState.videoList = videoList
    }
  }).finally(()=>{
    loading.value = false;
  })
})

const loading = ref(false);
const previewVisible = ref(false);
const previewImage = ref('');

const handleCancel = () => {
  previewVisible.value = false;
};

const handlePreview = async (file) => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64(file.originFileObj);
  }
  previewImage.value = file.url || file.preview;
  previewVisible.value = true;
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

const beforeUploadVideo = (file) => {
  const isVideo = file.type.split("/")[0]==='video'
  if (!isVideo) {
    console.log(file)
    message.error('只能上传 视频 文件!');
    return false
  }
  const isLt500M = file.size / 1024 / 1024 < 500;
  if (!isLt500M) {
    message.error('视频必须小于 500MB!');
    return false
  }
  return true;
};
const handleChangePhoto = async  (info) => {
  if (info.file && info.file.status === "removed") {
    // Handle file removal
    formState.photoList = formState.photoList.filter(item => item.uid !== info.file.uid);
    return;
  }

  let fileList = [...info.fileList];
  const uuids = new Set();
  let hasUploading = false;

  fileList = fileList.filter(item => {
    if (item.status === 'uploading') {
      hasUploading = true;
      return true; // Keep uploading files in the list
    }
    if (item.status === 'error') {
      return false;
    }
    const uuid = item.uuid || (item.response && item.response.uuid) || '';
    const url = item.url || (item.response && item.response.url) || '';
    if (!uuid) {
      return false;
    }
    // Check for duplicates
    if (uuids.has(uuid)) {
      return false;
    }

    uuids.add(uuid);
    item.uuid = uuid;
    item.url = url;
    return true;
  });
  if (!hasUploading) {
    formState.photoList = [...fileList];
  }
};
const customRequestPhoto = ({file, onSuccess,onError}) => {
  let formData = new FormData();
  formData.append('file',file)
  fileUpload(formData).then(res=>{
    if(res&&res.code===200){
      const item = {uuid:res.data.uuid,name:file.name,url:filePrefix+res.data.uuid}
      onSuccess(item)
    }else{
      onError(res.message||'上传失败')
    }
  }).finally(()=>{
  }).catch((e)=>{
    onError(e)
  })
};
const customRequestVideo = ({file, onSuccess}) => {
  let formData = new FormData();
  formData.append('file',file)
  fileUpload(formData).then(res=>{
    if(res&&res.code===200){
      onSuccess({uuid:res.data.uuid,name:file.name,url:filePrefix+res.data.uuid})
    }else{
      onError(res.message||'上传失败')
    }
  }).finally(()=>{
  }).catch((e)=>{
    onError(e)
  })
};
const handleChangeVideo = (info) => {
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
      console.log(item,"哪里出现问题了")
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
    formState.videoList = [resFileList[resFileList.length-1]]
    return
  }
  formState.videoList = resFileList
};

import { useRouter,useRoute } from 'vue-router';
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}
const onFinish = (values) => {
  let form = Object.assign({},formState,values)
  form.photo = []
  for (let i in form.photoList){
    let item = form.photoList[i].uuid
    form.photo.push(item)
  }
  if(form.videoList.length>0){
    form.video = form.videoList[0].uuid
  }
  delete form['photoList']
  delete form['videoList']

   loading.value = true
   travelSave(form).then(res=>{
     if(res&&res.code==200){
       toRoute('/travel')
     }
   }).finally(()=>{
     loading.value = false
   })
};

const getBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
};
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
</script>

<style scoped>
/* 自定义上传按钮样式 */
:deep(.ant-upload-select-picture-card) {
  @apply bg-white border border-gray-300 rounded-lg;
}

:deep(.ant-upload-select-picture-card:hover) {
  @apply border-blue-500;
}

/* 响应式调整 */
@media (max-width: 640px) {
  :deep(.ant-form-item-label) {
    @apply pb-2;
  }
}
</style>

