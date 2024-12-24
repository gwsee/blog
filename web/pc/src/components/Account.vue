<template>
  <a-modal v-model:open="open"  @cancel="close" :closable="false">
    <a-card :title="isReset?'重置密码':'我的账户' ">
      <template #extra><a @click="clickManage">{{ isIngEdit?'取消':'去管理' }}</a></template>
      <a-form
          style="margin: 4px 4px"
          ref="formRef"
          :model="formState"
          autocomplete="off"
      >
        <a-form-item
            name="account"
            :rules="[{ required: true, message: '账户名称' }]"
        >
          <a-input placeholder="账户名称" v-model:value="formState.account" disabled />
        </a-form-item>
        <a-form-item
            v-if="!isReset"
            name="nickname"
            :rules="[{ required: true, message: '请输入昵称' }]"
        >
          <a-input placeholder="昵称" v-model:value="formState.nickname" :disabled="!isIngEdit" />
        </a-form-item>
        <a-form-item
            v-if="!isReset"
            name="email"
            :rules="[{ required: false, message: '请输入邮箱' , type: 'email'}]"
        >
          <a-input placeholder="邮箱" v-model:value="formState.email"  :disabled="!isIngEdit" />
        </a-form-item>
        <a-form-item
            v-if="!isReset"
            name="avatarList"
        >
          <a-upload  :max-count="1"
                     :multiple="false"
                     :disabled="!isIngEdit"
                     list-type="picture-card"
                     v-model:fileList="formState.avatarList"
                     :customRequest="$customRequest"
                     :before-upload="beforeUpload"
                     @change="handleChange"
                     @preview="handlePreview"

          >
            <div v-if="formState.avatarList<1">
              <PlusOutlined />
              <div style="margin-top: 8px">上传头像</div>
            </div>
          </a-upload>
        </a-form-item>
        <a-form-item
            v-if="!isReset"
            name="description"
        >
          <a-textarea placeholder="要不来个介绍?" v-model:value="formState.description" :disabled="!isIngEdit"  />
        </a-form-item>
        <a-form-item
            v-if="isReset&&isIngEdit"
            name="password"
            :rules="[{ required: true, message: isReset?'请输入新密码':'请输入密码' }]"
        >
          <a-input-password  :placeholder="isReset?'请输入新密码':'请输入密码'" v-model:value="formState.password" />
        </a-form-item>
        <a-form-item
            v-if="isReset&&isIngEdit"
            name="confirm"
            :rules="[{ required: true,   validator: validatePassconfirm }]"
        >
          <a-input-password  placeholder="请再次确认密码" v-model:value="formState.confirm" />
        </a-form-item>
      </a-form>
    </a-card>
    <template #footer>
      <a-button type="primary" key="reg" @click="()=>{isReset=!isReset}" style="float: left" v-if="isIngEdit">{{ isReset?'去账户管理':'去重置密码' }}</a-button>
      <a-button type="info" key="logout" @click="loginOut" style="float: left" v-else >{{ '退出登录' }}</a-button>
      <a-button key="back" @click="close"  v-if="!isIngEdit">{{ '关闭' }}</a-button>
      <a-button type="primary" :loading="confirmLoading" @click="handleOk" v-if="isIngEdit">{{ isReset?'确定':'保存' }}</a-button>
    </template>
    <!-- 图片预览模态框 -->
    <a-modal :open="previewVisible" :footer="null" @cancel="handleCancel">
      <img alt="example" style="width: 100%" :src="previewImage" />
    </a-modal>
  </a-modal>
</template>

<script setup>
import { PlusOutlined } from '@ant-design/icons-vue'
import { login, updateAccount, resetPass} from "@/api/account";
import {ref, reactive, watch} from 'vue';
import { useAuthStore  } from '@/store/auth'
import {filePrefix} from "@/api/tool.js";

const { showLogin,setLoginShow,logout, state } = useAuthStore();
const open = ref(false);
const formRef = ref(null)
const confirmLoading = ref(false);
const isReset = ref(false);
const isIngEdit = ref(false);
const formState = reactive({
  id:0,
  nickname: '',
  account: '',
  email: '',
  avatar: '',
  avatarList: [],
  description: '',
  password: '',
  confirm: '',
});
const labelCol = { style: { width: '90px' } };
const wrapperCol = { span: 24 };
const validatePassconfirm= async (_rule, value) => {
  if (!isReset.value){
    return Promise.resolve();
  }
  if (value === '') {
    return Promise.reject('请再次确认密码');
  } else if (value !== formState.password) {
    return Promise.reject("两次密码输入不一致!");
  } else {
    return Promise.resolve();
  }
};
import { useRouter } from 'vue-router';
import {message} from "ant-design-vue";
const router = useRouter();
const goToAbout=(path)=> {
  router.push(path)
}
const clickManage = ()=>{
  if(!isIngEdit.value){
    isReset.value = false;
    isIngEdit.value = true;
  }else{
    isReset.value = false;
    isIngEdit.value = false;
  }
}
const loginOut = ()=>{
  logout()
  close()
  goToAbout('/');
}
watch(showLogin,(value)=>{
  if(value){
    open.value = true
  }
})
const close = ()=>{
  setLoginShow(false);
  open.value = false
}

const handleSave =(data)=>{
  confirmLoading.value = true;
  if(data.avatarList.length>0){
    data.avatar=data.avatarList[0].uuid;
  }else{
    data.avatar="";
  }
  updateAccount(data).then(res=>{
    if(res&&res.code===200){
      isReset.value = false;
      isIngEdit.value = false;
      message.success("账户修改成功")
    }
  }).finally(()=>{
    confirmLoading.value = false;
  });
}

const handleReset =(data)=>{
  confirmLoading.value = true;
  resetPass(data).then(res=>{
    if(res&&res.code===200){
      isReset.value = false;
      isIngEdit.value = false;
      message.success("密码修改成功")
    }
  }).finally(()=>{
    confirmLoading.value = false;
  });
}
const handleOk = () => {
    formRef.value
    .validate()
    .then(() => {
      if(isReset.value){
        handleReset(formState)
      }else{
        handleSave(formState)
      }
    })
    .catch(error => {
      console.log('error', error);
    });
};

// 通过storage来控制弹出的显隐 就不用show 来暴露
const show = () => {
  isReset.value = false;
  isIngEdit.value = false;
  open.value = true;
  if(formRef.value){
    formRef.value.resetFields();
  }
  formState.account = state.user.account;
  formState.id = state.user.id;
  formState.nickname = state.user.nickname;
  formState.email = state.user.email;
  formState.avatar = state.user.avatar;
  if(formState.avatar&&formState.avatar!==""){
    formState.avatarList=[{uuid:formState.avatar,url:filePrefix+formState.avatar}];
  }
  formState.description = state.user.description;
};

defineExpose({
  show
})



const handlePreview = async (file) => {
  if (!file.url && !file.preview) {
    file.preview = await getBase64(file.originFileObj);
  }
  previewImage.value = file.url || file.preview;
  previewVisible.value = true;
};
const previewVisible = ref(false);
const previewImage = ref('');

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
      console.log(item,"哪里出现问题了")
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
  if(resFileList.length>0){
    formState.avatarList = [resFileList[resFileList.length-1]]
    return
  }
  formState.avatarList = resFileList
};

</script>

<style scoped>

</style>
