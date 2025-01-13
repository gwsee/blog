<template>
  <a-modal v-model:open="open" :title="`项目管理`"  @cancel="close">
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          ref="formRef"
          @finish="onFinish"
          class=" py-8 rounded-xl p-4 sm:p-6"
      >
        <a-form-item name="title" label="项目名称" :rules="[{ required: true }]">
          <a-input v-model:value="formState.title" />
        </a-form-item>
        <a-form-item name="description" label="项目描述" :rules="[{ required: true }]">
          <a-textarea v-model:value="formState.description" :rows="4" />
        </a-form-item>
        <a-form-item name="skills" label="使用技能">
          <a-select
              v-model:value="formState.skills"
              mode="tags"
              style="width: 100%"
              placeholder="Enter skills"
          >
          </a-select>
        </a-form-item>

        <a-form-item name="start" label="开始日期">
          <a-date-picker placeholder="开始日期"  v-model:value="formState.start" value-format="X" />
        </a-form-item>

        <a-form-item name="end" label="结束日期">
          <a-date-picker placeholder="至今" v-model:value="formState.end" value-format="X" />
        </a-form-item>

        <a-form-item name="link" label="项目地址">
          <a-input v-model:value="formState.link" />
        </a-form-item>

        <a-form-item name="photosList" label="项目快照">
          <a-upload
              v-model:fileList="formState.photosList"
              name="file"
              list-type="picture-card"
              class="project-image-uploader"
              :multiple="false"
              :before-upload="beforeUpload"
              @change="handleChange"
              :customRequest="$customRequest"
              @preview="handlePreview"
          >
            <div >
              <plus-outlined />
              <div style="margin-top: 8px">Upload</div>
            </div>
          </a-upload>
        </a-form-item>
      </a-form>
    <template #footer>
      <a-button key="back" @click="close">取消</a-button>
      <a-button type="primary" :loading="confirmLoading" @click="handleOk">{{ '保存' }}</a-button>
    </template>
    <!-- 图片预览模态框 -->
    <a-modal :open="previewVisible" :footer="null" @cancel="handleCancel">
      <img alt="example" style="width: 100%" :src="previewImage" />
    </a-modal>
  </a-modal>
</template>

<script setup>
import {userGet, userSave, projectList, experienceList, projectGet, projectSave} from "@/api/user";
import {ref, reactive, onMounted, getCurrentInstance} from 'vue'
import { message } from 'ant-design-vue'
import {filePrefix} from "@/api/tool.js";
import { PlusOutlined } from '@ant-design/icons-vue'
const open = ref(false)
const confirmLoading = ref(false)
const handleOk = ()=>{
  formRef.
  value.
  validate().
  then(()=>{
    confirmLoading.value = true
    const photos = []
    if(formState.photosList.length>0){
      formState.photosList.forEach(item=>{
        photos.push(item.uuid)
      })
      formState.photos = photos
    }
    projectSave(formState).then((res)=>{
      if(res){
        confirmLoading.value = false
        open.value = false
        message.success('saved successfully!')
      }
    }).finally(()=>{
      confirmLoading.value = false
    })
  })
}
const onFinish = (values) => {
  console.log('Success:', values)
  message.success('Project saved successfully!')
  // Here you would typically send the data to your backend
}
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
// 通过storage来控制弹出的显隐 就不用show 来暴露
const formRef = ref(null)
const formState = reactive({
  id: 0,
  experienceId: 0,
  title: '',
  description: '',
  skills: [],
  start: 0,
  end: 0,
  link: '',
  photos: [],
  photosList: [],
})
const show = (id,experienceId) => {
  confirmLoading.value = false
  open.value = true;
  if(formRef.value){
    formRef.value.resetFields();
  }
  formState.experienceId = experienceId;
  if(!id){
    return
  }
  projectGet({id}).then((res) => {
    if(res){
      const data = res.data||{}
      formState.id = data.id;
      formState.title = data.title;
      formState.description = data.description;
      formState.skills = data.skills;
      formState.start = data.start;
      formState.end = data.end;
      formState.link = data.link;
      formState.photos = data.photos;
      if(formState.photos.length > 0){
        formState.photos.forEach(photo => {
          formState.photosList.push({uuid:photo,url:filePrefix+(photo)});
        })
      }
    }
  })
};
const close = ()=>{
  open.value = false
}
defineExpose({
  show
})
const handleChange = (info) => {
  if (info.file&&info.file.status === "uploading"){
    return false
  }
  let resFileList = [...info.fileList];
  const includes = [];
  resFileList = resFileList.filter(function (item) {
    if(item.status==='error'){
      return false
    }
    let uuid = item.uuid||item.response&&item.response.uuid||''
    let url = item.url||item.response&&item.response.url||''
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
  formState.photosList = resFileList
};
const previewVisible = ref(false);
const previewImage = ref('');
const handlePreview = async (file) => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64(file.originFileObj);
  }
  previewImage.value = file.url || file.preview;
  previewVisible.value = true;
};
const handleCancel = () => {
  previewVisible.value = false;
};
const getBase64 = (img, callback) => {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}

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
</script>

<style scoped>
.project-image-uploader > .ant-upload {
  width: 128px;
  height: 128px;
}
</style>

