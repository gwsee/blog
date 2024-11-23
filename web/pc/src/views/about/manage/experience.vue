<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="container mx-auto px-4">
      <h1 class="text-3xl font-bold mb-8 text-center">Add Work Experience</h1>
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          @finish="onFinish"
          class="max-w-2xl mx-auto bg-white p-8 rounded-lg shadow-md"
      >
        <a-form-item name="role" label="Job Title" :rules="[{ required: true }]">
          <a-input v-model:value="formState.role" />
        </a-form-item>

        <a-form-item name="company" label="Company" :rules="[{ required: true }]">
          <a-input v-model:value="formState.company" />
        </a-form-item>

        <a-form-item name="location" label="Location">
          <a-input v-model:value="formState.location" />
        </a-form-item>

        <a-form-item name="period" label="Employment Period" :rules="[{ required: true }]">
          <a-range-picker v-model:value="formState.period" />
        </a-form-item>

        <a-form-item name="description" label="Job Description" :rules="[{ required: true }]">
          <a-textarea v-model:value="formState.description" :rows="4" />
        </a-form-item>

        <a-form-item name="responsibilities" label="Key Responsibilities">
          <a-textarea v-model:value="formState.responsibilities" :rows="4" placeholder="Enter each responsibility on a new line" />
        </a-form-item>

        <a-form-item name="achievements" label="Key Achievements">
          <a-textarea v-model:value="formState.achievements" :rows="4" placeholder="Enter each achievement on a new line" />
        </a-form-item>

        <a-form-item name="skills" label="Skills Used">
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

        <a-form-item>
          <a-button type="primary" html-type="submit" block>Save Experience</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { message } from 'ant-design-vue'
const labelCol = { style: { width: '80px' } };
const wrapperCol = { span: 24 };
const formState = reactive({
  role: '',
  company: '',
  location: '',
  period: [],
  description: '',
  responsibilities: '',
  achievements: '',
  skills: [],
})

const predefinedSkills = [
  'JavaScript', 'React Native', 'iOS Development', 'Android Development',
  'Node.js', 'Python', 'UI/UX Design', 'Project Management'
]

const onFinish = (values) => {
  // Convert period to string format
  values.period = `${values.period[0].format('YYYY-MM')} - ${values.period[1].format('YYYY-MM')}`

  // Split responsibilities and achievements into arrays
  values.responsibilities = values.responsibilities.split('\n').filter(item => item.trim() !== '')
  values.achievements = values.achievements.split('\n').filter(item => item.trim() !== '')

  console.log('Success:', values)
  message.success('Experience saved successfully!')
  // Here you would typically send the data to your backend
}
</script>

