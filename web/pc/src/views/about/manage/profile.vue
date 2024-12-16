<template>
  <div class="min-h-screen  px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto  py-8"  style=" background: linear-gradient(rgb(241 241 241), rgb(173 215 197));">
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          @finish="onFinish"
          class="max-w-2xl mx-auto bg-white  rounded-lg shadow-md"
      >
        <a-card title="基本信息">
          <template #extra>
            <a href="#">编辑</a>
          </template>
          <a-form-item name="name" label="姓名" :rules="[{ required: true }]">
            <a-input v-model:value="formState.name" />
          </a-form-item>
          <a-form-item name="email" label="邮箱" :rules="[{ required: true, type: 'email' }]">
            <a-input v-model:value="formState.email" />
          </a-form-item>
          <a-form-item name="avatar" label="头像">
            <a-upload
                v-model:fileList="fileList"
                name="avatar"
                list-type="picture-card"
                class="avatar-uploader"
                :show-upload-list="false"
                action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                :before-upload="beforeUpload"
                @change="handleChange"
            >
              <img v-if="imageUrl" :src="imageUrl" alt="avatar" />
              <div v-else>
                <plus-outlined />
                <div style="margin-top: 8px">Upload</div>
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
              <a-select-option v-for="skill in predefinedSkills" :key="skill" :value="skill">
                {{ skill }}
              </a-select-option>
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
                  <a key="list-loadmore-edit">edit</a>&nbsp;
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
                <div class=" content-wrapper truncate overflow-hidden whitespace-nowrap">contentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontent</div>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
        <a-form-item style="text-align: center;padding: 5px 5px">
          <a-button type="primary" html-type="submit" block>保    存</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import {useRouter} from 'vue-router'
const router = useRouter()
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
const formState = reactive({
  name: '',
  title: '',
  bio: '',
  email: '',
  location: '',
  skills: [],
  experiences: [],
})
const toRoute = (path) => {
  router.push(path)
}
const maxVisibleProjects = ref(3)
const visibleProjects = (projects)=> {
  return projects.slice(0, maxVisibleProjects);
}
const allProjectNames=(projects)=>{
  return projects.map(project => project).join(', ');
}
const predefinedSkills = [
  'JavaScript', 'React Native', 'iOS Development', 'Android Development',
  'Node.js', 'Python', 'UI/UX Design', 'Project Management'
]

const fileList = ref([])
const imageUrl = ref('')

const beforeUpload = (file) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isJpgOrPng) {
    message.error('You can only upload JPG or PNG files!')
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    message.error('Image must be smaller than 2MB!')
  }
  return isJpgOrPng && isLt2M
}

const handleChange = (info) => {
  if (info.file.status === 'uploading') {
    return
  }
  if (info.file.status === 'done') {
    // Get this url from response in real world.
    getBase64(info.file.originFileObj, (url) => {
      imageUrl.value = url
    })
  }
}

const getBase64 = (img, callback) => {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}

const onFinish = (values) => {
  console.log('Success:', values)
  message.success('Profile saved successfully!')
  // Here you would typically send the data to your backend
}
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

