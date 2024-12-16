<template>
  <div class="min-h-screen  py-12">
    <div class="container mx-auto px-4">
      <h1 class="text-3xl font-bold mb-8 text-center"  v-if="false">Add Work Experience</h1>
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          @finish="onFinish"
          class="max-w-2xl mx-auto bg-white  rounded-lg shadow-md"
      >
        <a-card title="工作经历">
          <template #extra>
            <a-button >编辑</a-button>
          </template>
          <a-form-item name="company" label="公司名称" :rules="[{ required: true }]">
            <a-input v-model:value="formState.company" />
          </a-form-item>
          <a-form-item name="role" label="职位名称" :rules="[{ required: true }]">
            <a-input v-model:value="formState.role" />
          </a-form-item>
          <a-form-item name="location" label="公司地址">
            <a-input v-model:value="formState.location" />
          </a-form-item>
          <a-form-item name="period" label="在职时间" :rules="[{ required: true }]">
            <a-range-picker v-model:value="formState.period" />
          </a-form-item>
          <a-form-item name="description" label="职位描述" :rules="[{ required: true }]">
            <a-textarea v-model:value="formState.description" :rows="4" />
          </a-form-item>
          <a-form-item name="responsibilities" label="主要职责">
            <a-textarea v-model:value="formState.responsibilities" :rows="4" placeholder="Enter each responsibility on a new line" />
          </a-form-item>
          <a-form-item name="achievements" label="工作成就">
            <a-textarea v-model:value="formState.achievements" :rows="4" placeholder="Enter each achievement on a new line" />
          </a-form-item>
          <a-form-item name="skills" label="使用技能">
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
        </a-card>
        <a-card title="项目经历">
          <template #extra>
            <a-button @click="showProject">add more</a-button>
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
          <a-button type="primary" html-type="submit" block>Save Experience</a-button>
        </a-form-item>
      </a-form>
    </div>
    <project ref="projectRef"  />
  </div>
</template>

<script setup>
import project from "./project.vue"
import {reactive, ref} from 'vue'
import {useRouter} from 'vue-router'
const router = useRouter()
const projectRef = ref(null);
import { message } from 'ant-design-vue'
const labelCol = { style: { width: '90px' } };
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
  experiences:[],
})
const showProject = ()=>{
  projectRef.value?.show()
}
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

