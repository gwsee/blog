<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929;text-align: center" >
      <a-form :model="formState" :label-col="labelCol"   class="blog-card-edit" :wrapper-col="wrapperCol">
        <a-form-item label="Title">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="Description">
          <a-textarea v-model:value="formState.desc" />
        </a-form-item>
        <a-form-item label="IsHidden">
          <a-switch v-model:checked="formState.delivery" />
        </a-form-item>
        <a-form-item label="Cover">
          <a-upload action="/upload.do" :maxCount="1" list-type="picture-card">
            <div>
              <PlusOutlined />
              <div style="margin-top: 8px">Upload</div>
            </div>
          </a-upload>
        </a-form-item>

        <a-form-item label="Resources">
          <div style="border: 1px solid #ccc">
            <Toolbar
                style="border-bottom: 1px solid #ccc"
                :editor="editorRef"
                :defaultConfig="toolbarConfig"
                :mode="mode"
            />
            <Editor
                style="height: 500px; overflow-y: hidden;"
                v-model="valueHtml"
                :defaultConfig="editorConfig"
                :mode="mode"
                @onCreated="handleCreated"
            />
          </div>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 4 }" style="text-align: center">
          <a-button type="primary" @click="onSubmit">Create</a-button>
          <a-button style="margin-left: 10px">Cancel</a-button>
        </a-form-item>
      </a-form>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  style="background-color: #1a1a1a" >
      3
    </a-col>
  </a-row>
</template>

<script setup>
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { onBeforeUnmount,  shallowRef, onMounted ,reactive, toRaw, ref } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
// 内容 HTML
const valueHtml = ref('<p>hello</p>')
const editorRef = shallowRef()
// 模拟 ajax 异步获取内容
onMounted(() => {
  setTimeout(() => {
    valueHtml.value = '<p>模拟 Ajax 异步设置内容</p>'
  }, 1500)
})
const mode = ref("default")
const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}
const formState = reactive({
  name: '',
  delivery: false,
  type: [],
  resource: '',
  desc: '<p>Hello from CKEditor 5 in Vue!</p>',
});
const onSubmit = () => {
  console.log('submit!', toRaw(formState));
};
const labelCol = { style: { width: '120px' } };
const wrapperCol = { span: 20 };
</script>

<style scoped>
.blog-content-eclipses{
  flex: 1; /* 右边部分自适应 */
  margin: 15px 15px 15px 0;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 显示3行 */
  overflow: hidden;
  font-size: 15px;
  text-overflow: ellipsis;
}

.blog-card-edit{
  width: 95%;
  margin-top: 20px;
  text-align: left;
  margin-left: auto;
  margin-right: auto;
}
</style>
