<template>
  <a-modal v-model:open="open" :title="isRegister?`注册`:`登录`" >
    <a-form
        style="margin: 32px 4px"
        ref="formRef"
        :model="formState"
        :wrapper-col="{ span: 20, offset:2 }"
        autocomplete="off"
    >
      <a-form-item
          name="Account"
          :rules="[{ required: true, message: '请输入账户名称' }]"
      >
        <a-input placeholder="请输入账户名称" v-model:value="formState.Account" />
      </a-form-item>

      <a-form-item
          name="Password"
          :rules="[{ required: true, message: isRegister?'请输入新密码':'请输入密码' }]"
      >
        <a-input-password  :placeholder="isRegister?'请输入新密码':'请输入密码'" v-model:value="formState.Password" />
      </a-form-item>
      <a-form-item
          v-if="isRegister"
          name="Confirm"
          :rules="[{ required: true,   validator: validatePassConfirm }]"
      >
        <a-input-password  placeholder="请再次确认密码" v-model:value="formState.Confirm" />
      </a-form-item>
    </a-form>
    <template #footer>
      <a-button type="primary" key="back" @click="()=>{isRegister=!isRegister}" style="float: left">{{ isRegister?'去登录':'去注册' }}</a-button>
      <a-button key="back" @click="close">取消</a-button>
      <a-button type="primary" :loading="confirmLoading" @click="handleOk">{{ isRegister?'提交':'确定' }}</a-button>
    </template>
  </a-modal>
</template>

<script setup>
import { login, register, resetPass} from "@/api/account.js";
import { setToken } from '@/utils/auth'
import { ref,reactive } from 'vue';
import { useAuthStore  } from '@/store/auth.js'
const { setLoginShow } = useAuthStore();

const open = ref(false);
const formRef = ref(null)
const confirmLoading = ref(false);
const isRegister = ref(false);
const formState = reactive({
  Account: '',
  Password: '',
  Confirm: '',
});
const validatePassConfirm= async (_rule, value) => {
  if (!isRegister.value){
    return Promise.resolve();
  }
  if (value === '') {
    return Promise.reject('请再次确认密码');
  } else if (value !== formState.Password) {
    return Promise.reject("两次密码输入不一致!");
  } else {
    return Promise.resolve();
  }
};


const close = ()=>{
  setLoginShow(false);
  open.value = false
}

const handleLogin =(data)=>{
  confirmLoading.value = true;
  login(data).then(res=>{
    if(res.code===0){
      setToken(res.data.Token)
      confirmLoading.value = false;
      close()
    }
  }).finally(()=>{
    confirmLoading.value = false;
  });
}
const handleRegister =(data)=>{
  confirmLoading.value = true;
  register(data).then(res=>{
    isRegister.value = false;
  }).finally(()=>{
    confirmLoading.value = false;
  });
}
const handleReset =(data)=>{
  confirmLoading.value = true;
  register(data).then(res=>{

  }).finally(()=>{
    confirmLoading.value = false;
  });
}
const handleOk = () => {
    formRef.value
    .validate()
    .then(() => {
      console.log(isRegister.value);
      if(isRegister.value){
        handleRegister(formState)
      }else{
        handleLogin(formState)
      }
    })
    .catch(error => {
      console.log('error', error);
    });
};

// 通过storage来控制弹出的显隐 就不用show 来暴露
const show = () => {
  open.value = true;
  if(formRef.value){
    formRef.value.resetFields();
  }
};

defineExpose({
  show
})
</script>

<style scoped>

</style>
