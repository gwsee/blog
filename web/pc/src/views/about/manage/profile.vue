<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="container mx-auto px-4">
      <h1 class="text-3xl font-bold mb-8 text-center">Add Profile Information</h1>
      <a-form
          :model="formState"
          @finish="onFinish"
          class="max-w-2xl mx-auto bg-white p-8 rounded-lg shadow-md"
      >
        <a-form-item name="name" label="Full Name" :rules="[{ required: true }]">
          <a-input v-model:value="formState.name" />
        </a-form-item>

        <a-form-item name="title" label="Professional Title" :rules="[{ required: true }]">
          <a-input v-model:value="formState.title" />
        </a-form-item>

        <a-form-item name="bio" label="Bio" :rules="[{ required: true }]">
          <a-textarea v-model:value="formState.bio" :rows="4" />
        </a-form-item>

        <a-form-item name="email" label="Email" :rules="[{ required: true, type: 'email' }]">
          <a-input v-model:value="formState.email" />
        </a-form-item>

        <a-form-item name="location" label="Location">
          <a-input v-model:value="formState.location" />
        </a-form-item>

        <a-form-item name="skills" label="Skills">
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

        <a-form-item name="avatar" label="Profile Picture">
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

        <a-form-item>
          <a-button type="primary" html-type="submit" block>Save Profile</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'

const formState = reactive({
  name: '',
  title: '',
  bio: '',
  email: '',
  location: '',
  skills: [],
})

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
</style>

