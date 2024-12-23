<template>
  <div class="min-h-screen  px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto  py-8"  style=" background: linear-gradient(rgb(241 241 241), rgb(173 215 197));">
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          ref="formRef"
          @finish="onFinish"
          class="max-w-2xl mx-auto bg-white  rounded-lg shadow-md"
      >
        <a-card title="基本信息">
          <template #extra>
            <a-button  v-if="false">编辑</a-button>
          </template>
          <a-form-item name="name" label="姓名" :rules="[{ required: true }]">
            <a-input v-model:value="formState.name" />
          </a-form-item>
          <a-form-item name="email" label="邮箱" :rules="[{ required: true, type: 'email' }]">
            <a-input v-model:value="formState.email" />
          </a-form-item>
          <a-form-item name="avatar" label="头像">
            <a-upload
                v-model:fileList="formState.avatarList"
                name="avatar"
                :max-count="1"
                :multiple="false"
                list-type="picture-card"
                :customRequest="customRequest"
                :before-upload="beforeUpload"
                @change="handleChange"
                @preview="handlePreview"
            >
              <div v-if="formState.avatarList<1">
                <plus-outlined />
                <div style="margin-top: 8px">上传头像</div>
              </div>
            </a-upload>
          </a-form-item>

          <a-form-item name="professional" label="职称" :rules="[{ required: true }]">
            <a-input v-model:value="formState.professional" />
          </a-form-item>
          <a-form-item name="skills" label="技能">
            <a-select
                v-model:value="formState.skills"
                mode="tags"
                style="width: 100%"
                placeholder="Enter skills"
            >
            </a-select>
          </a-form-item>
          <a-form-item name="description" label="个人经历" :rules="[{ required: true }]">
            <a-textarea v-model:value="formState.description" :rows="4" />
          </a-form-item>
          <a-form-item name="address" label="地址">
            <a-input v-model:value="formState.address" />
          </a-form-item>

        </a-card>
        <a-card title="工作经历">
          <template #extra>
            <a-button @click="toRoute('/about/experience/manage/0')">add more</a-button>
          </template>
          <a-list
              :data-source="formState.experiences"
          >
            <template #renderItem="{ item }">
              <a-list-item class="a-list-item-action">
                <template #actions >
                  <a key="list-loadmore-edit" @click="toRoute('/about/experience/manage/'+item.id)">edit</a>&nbsp;
                  <a key="list-loadmore-more">more</a>
                </template>
                <a-skeleton avatar :title="false"  :loading="false" class="relative">
                  <a-list-item-meta
                      :description="`${item.company} | ${item.period}`"
                  >
                    <template #title>
                      <a href="https://www.antdv.com/">{{ item.role }}</a>
                    </template>
                  </a-list-item-meta>
                </a-skeleton>
                <div class=" content-wrapper truncate overflow-hidden whitespace-nowrap">ddd</div>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
        <a-form-item style="text-align: center;padding: 5px 5px">
          <a-button type="primary" html-type="submit" block>保    存</a-button>
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
import { ref, reactive, onMounted } from 'vue'
import {userGet,userSave,projectList,experienceList} from "@/api/user";
import {useRouter} from 'vue-router'
const router = useRouter()
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import {filePrefix,fileUpload} from "@/api/tool";
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
const toRoute = (path) => {
  router.push(path)
}
const maxVisibleProjects = ref(3)
const allProjectNames=(projects)=>{
  return projects.map(project => project).join(', ');
}
const predefinedSkills = []
const formState = reactive({
  avatarList: [],
  avatar: '',
  name: '',
  title: '',
  email: '',
  location: '',
  skills: [],
  experiences: [],
})

const formRef = ref(null)
onMounted(()=> {
  formState.experiences = [];
  if(formRef.value){
    formRef.value.resetFields();
  }
  userGet({}).then(res => {
    if (res && res.code === 200) {
      const data =res.data || {}
      formState.name  =data.name
      formState.email =data.email
      formState.avatar = data.avatar
      if(data.avatar){
        formState.avatarList = [{uuid:data.avatar,url:filePrefix+data.avatar}]
      }
      formState.professional = data.professional
      formState.skills = data.skills
      formState.description = data.description
      formState.address = data.address
    }
  })
  experienceList({"page":{"pageNum":1,"pageSize":10}}).then(res=>{
    if(res&&res.code===200){
      let data =res.data
      formState.experiences = data.list || []
      // experiencesMore.value = experiences.value.length<(data.total || 0)
    }
  })
})
const getBase64 = (img, callback) => {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}
const loading = ref(false)
const onFinish = (values) => {
  loading.value = true
  if(formState.avatarList.length>0){
    formState.avatar = formState.avatarList[0].uuid
  }
  console.log(values,formState)
  userSave(formState).then(res=>{
    if(res){
      message.success('saved successfully!')
    }
  })
}

// <!--图片部分-->
const customRequest = ({file, onSuccess,onError}) => {
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
.avatar-uploader > .ant-upload {
  width: 128px;
  height: 128px;
}
.ant-list-item {
  flex-direction: column;
  align-items: flex-start;
}

.ant-list-item-meta {
  margin-bottom: 16px;
  width: 100%;
}

.content-wrapper {
  width: 100%;
  margin-bottom: 16px;
}


::v-deep .a-list-item-action .ant-list-item-action {
  right: 0!important;
  position: absolute!important;
}

</style>

