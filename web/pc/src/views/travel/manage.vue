<template>
  <div class="min-h-screen bg-gray-50 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto">
      <h1 class="text-3xl font-bold text-gray-900 mb-8 text-center">新增旅行内容</h1>
      <a-form
          :model="formState"
          @finish="onFinish"
          :label-col="{ span: 24 }"
          :wrapper-col="{ span: 24 }"
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

        <!-- 照片上传 -->
        <a-form-item
            name="photos"
            label="旅行照片"
            :rules="[{ required: true, message: '请上传至少一张照片' }]"
        >
          <a-upload
              v-model:fileList="formState.photos"
              list-type="picture-card"
              :customRequest="customRequest"
              :before-upload="beforeUpload"
              @preview="handlePreview"
          >
            <div v-if="formState.photos.length < 8">
              <plus-outlined />
              <div class="mt-2">上传照片</div>
            </div>
          </a-upload>
        </a-form-item>

        <!-- 视频上传 -->
        <a-form-item name="video" label="视频">
          <a-upload
              v-model:fileList="formState.video"
              :customRequest="customRequest"
              :before-upload="beforeUploadVideo"
              @change="handleVideoChange"
          >
            <a-button icon="<upload-outlined />">上传视频</a-button>
          </a-upload>
        </a-form-item>

        <!-- 提交按钮 -->
        <a-form-item :wrapper-col="{ span: 24 }">
          <a-button type="primary" html-type="submit" class="w-full">提交</a-button>
        </a-form-item>
      </a-form>

      <!-- 图片预览模态框 -->
      <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel">
        <img alt="example" style="width: 100%" :src="previewImage" />
      </a-modal>
    </div>
  </div>
</template>

<script setup>
import {ref, reactive} from 'vue';
import {message} from 'ant-design-vue';
import {PlusOutlined, UploadOutlined} from '@ant-design/icons-vue';

const formState = reactive({
  title: '',
  description: '',
  photos: [],
  video: []
});

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
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
  if (!isJpgOrPng) {
    message.error('只能上传 JPG/PNG 文件!');
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    message.error('图片必须小于 2MB!');
  }
  return isJpgOrPng && isLt2M;
};

const beforeUploadVideo = (file) => {
  const isMP4 = file.type === 'video/mp4';
  if (!isMP4) {
    message.error('只能上传 MP4 文件!');
  }
  const isLt50M = file.size / 1024 / 1024 < 50;
  if (!isLt50M) {
    message.error('视频必须小于 50MB!');
  }
  return isMP4 && isLt50M;
};

const customRequest = ({file, onSuccess}) => {
  // 这里应该是您的文件上传逻辑
  // 为了演示，我们只是模拟一个成功的上传
  setTimeout(() => {
    onSuccess("ok");
  }, 0);
};

const handleVideoChange = (info) => {
  if (info.file.status === 'done') {
    message.success(`${info.file.name} 文件上传成功`);
  } else if (info.file.status === 'error') {
    message.error(`${info.file.name} 文件上传失败`);
  }
};

const onFinish = (values) => {
  console.log('Success:', values);
  // 这里应该是您的表单提交逻辑
  message.success('表单提交成功！');
};

const getBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
};
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

