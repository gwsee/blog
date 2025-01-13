<template>
  <div class="min-h-screen">
    <div class="container mx-auto px-4 rounded-xl p-4 sm:p-6">
      <h1 class="text-3xl font-bold mb-8 text-center"  v-if="false">Add Work Experience</h1>
      <a-form
          :label-col="labelCol"  :wrapper-col="wrapperCol"
          :model="formState"
          ref="formRef"
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
          <a-form-item label="图片">
            <a-upload  :max-count="1"
                       :multiple="false"
                       list-type="picture-card"
                       v-model:fileList="formState.imageList"
                       :customRequest="$customRequest"
                       :before-upload="beforeUpload"
                       @change="handleChange"
                       @preview="handlePreview"

            >
              <div v-if="formState.imageList<1">
                <PlusOutlined />
                <div style="margin-top: 8px">上传封面</div>
              </div>
            </a-upload>
          </a-form-item>
          <a-row  >
            <a-col :md="11" :sm="24" :xs="24"   >
              <a-form-item name="start" label="在职时间" :rules="[{ required: true }]">
                  <a-date-picker placeholder="Start Date" format="YYYY-MM-DD" value-format="X" v-model:value="formState.start">
                  </a-date-picker>
              </a-form-item>
          </a-col>
            <a-col :md="11" :sm="24" :xs="24"   >
              <a-form-item name="start" label="" :rules="[{ required: false }]">
                <a-date-picker placeholder="至今" format="YYYY-MM-DD" value-format="X" v-model:value="formState.end">
                </a-date-picker>
              </a-form-item>
            </a-col>
          </a-row>
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
            ></a-select>
          </a-form-item>
        </a-card>
        <a-card title="项目经历" v-if="formState.id>0">
          <template #extra>
            <a-button @click="showProject(0,formState.id)">add more</a-button>
          </template>
          <a-list
              :data-source="formState.projects"
          >
            <template #renderItem="{ item }">
              <a-list-item class="a-list-item-action">
                <template #actions >
                  <a key="list-loadmore-edit" @click="showProject(item.id,item.experienceId)">edit</a>&nbsp;
                  <a key="list-loadmore-more" v-if="false">more</a>
                </template>
                <a-skeleton avatar :title="false"  :loading="false" class="relative">
                  <a-list-item-meta
                    :description="`${$formatDate(item.start,'{y}-{m}-{d}')}  | ${item.end&&$formatDate(item.end,'{y}-{m}-{d}')||'至今'} `"
                  >
                    <template #title>
                      <a :href="item.link">{{ item.title }}</a>
                    </template>
                  </a-list-item-meta>
                </a-skeleton>
                <div class=" content-wrapper truncate overflow-hidden whitespace-nowrap">
                  {{item.description}}
                </div>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
        <a-form-item style="text-align: center;padding: 5px 5px">
          <a-button type="primary" html-type="submit" block>Save Experience</a-button>
        </a-form-item>
      </a-form>
    </div>
    <!-- 图片预览模态框 -->
    <a-modal :open="previewVisible" :footer="null" @cancel="handleCancel">
      <img alt="example" style="width: 100%" :src="previewImage" />
    </a-modal>
    <project ref="projectRef"  />
  </div>
</template>

<script setup>
import {experienceGet, experienceSave, experienceList, projectList} from "@/api/user";
import project from "./project.vue"
import { PlusOutlined } from '@ant-design/icons-vue'
import {reactive, ref, onMounted, getCurrentInstance} from 'vue'
import {useRoute, useRouter} from 'vue-router'
const router = useRouter()
const projectRef = ref(null);
import { message } from 'ant-design-vue'
const labelCol = { style: { width: '90px' } };
const wrapperCol = { span: 24 };
const formRef = ref(null)
const showProject = (id,experienceId)=>{
  projectRef.value?.show(id,experienceId)
}
const predefinedSkills = []
const formState = reactive({
  id: 0,
  company: '',
  role: '',
  location: '',
  start:0,
  end:0,
  description: '',
  responsibilities: '',
  achievements: '',
  skills: [],
  image:"",
  imageList:[],
  period: [],
  projects:[],
})
const loading = ref(false);
onMounted(()=>{
  if(formRef.value){
    formRef.value.resetFields();
  }
  const route = useRoute();
  let id = route.params.id;
  id = id - 0
  if(!id){
    return
  }
  const { appContext } = getCurrentInstance();
  const $fileFull = appContext.config.globalProperties.$fileFull;
  experienceGet({id:id}).then((res) => {
    if(res){
      const data = res.data || {}
      formState.id = data.id;
      formState.company = data.company;
      formState.role = data.role;
      formState.location = data.location;
      formState.start = data.start;
      formState.end = data.end;
      formState.description = data.description;
      formState.skills = data.skills;
      formState.responsibilities = data.responsibilities;
      formState.achievements = data.achievements;
      formState.period = data.period;
      if(data.image!==""){
        formState.imageList = [{uuid:data.image,url:$fileFull(data.image)}]
      }
    }
  })
  projectList({experienceId: id}).then((res) => {
      if(res){
        formState.projects = res.data&&res.data.list || [];
      }
  })
})

const onFinish = (values) => {
  loading.value = true;
  if(formState.imageList.length>0){
    formState.image = formState.imageList[0].uuid;
  }
  experienceSave(formState).then((res) => {
    if(res){
      message.success('saved successfully!')
      router.back()
    }
  }).finally(() => {
    loading.value = false;
  })
  // Convert period to string format
  values.period = `${values.period[0].format('YYYY-MM')} - ${values.period[1].format('YYYY-MM')}`
  // Split responsibilities and achievements into arrays
  values.responsibilities = values.responsibilities.split('\n').filter(item => item.trim() !== '')
  values.achievements = values.achievements.split('\n').filter(item => item.trim() !== '')

  console.log('Success:', values)
  message.success('Experience saved successfully!')
  // Here you would typically send the data to your backend
}

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
  formState.imageList = resFileList
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

