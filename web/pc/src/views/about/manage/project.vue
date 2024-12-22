<template>
  <a-modal v-model:open="open" :title="`项目管理`"  @cancel="close">
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          ref="formRef"
          @finish="onFinish"
          class=" py-8 "
      >
        <a-form-item name="title" label="项目名称" :rules="[{ required: true }]">
          <a-input v-model:value="formState.title" />
        </a-form-item>
        <a-form-item name="description" label="项目描述" :rules="[{ required: true }]">
          <a-textarea v-model:value="formState.description" :rows="4" />
        </a-form-item>
        <a-form-item name="technologies" label="使用技能">
          <a-select
              v-model:value="formState.technologies"
              mode="tags"
              style="width: 100%"
              placeholder="Enter technologies"
          >
            <a-select-option v-for="tech in predefinedTechnologies" :key="tech" :value="tech">
              {{ tech }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item name="startDate" label="开始日期">
          <a-date-picker v-model:value="formState.startDate" />
        </a-form-item>

        <a-form-item name="endDate" label="结束日期">
          <a-date-picker v-model:value="formState.endDate" />
        </a-form-item>

        <a-form-item name="projectUrl" label="项目地址">
          <a-input v-model:value="formState.projectUrl" />
        </a-form-item>

        <a-form-item name="image" label="项目快照">
          <a-upload
              v-model:fileList="fileList"
              name="image"
              list-type="picture-card"
              class="project-image-uploader"
              :show-upload-list="false"
              action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              :before-upload="beforeUpload"
              @change="handleChange"
          >
            <img v-if="imageUrl" :src="imageUrl" alt="project" />
            <div v-else>
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
import {userGet,userSave,projectList,experienceList} from "@/api/user";
import {filePrefix} from "@/api/tool";
import { ref, reactive, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'
const open = ref(false)
const confirmLoading = ref(false)
const handleOk = ()=>{
  confirmLoading.value = true
}
const formState = reactive({
  title: '',
  description: '',
  technologies: [],
  startDate: null,
  endDate: null,
  projectUrl: '',
})

const predefinedTechnologies = [
  'JavaScript', 'React Native', 'iOS', 'Android', 'Node.js', 'Python',
  'Vue.js', 'Angular', 'Express', 'MongoDB', 'PostgreSQL', 'AWS'
]

const fileList = ref([])
const imageUrl = ref('')
const onFinish = (values) => {
  console.log('Success:', values)
  message.success('Project saved successfully!')
  // Here you would typically send the data to your backend
}
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
// 通过storage来控制弹出的显隐 就不用show 来暴露
const formRef = ref(null)
const show = () => {
  open.value = true;
  if(formRef.value){
    formRef.value.resetFields();
  }
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
  formState.avatarList = resFileList
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

