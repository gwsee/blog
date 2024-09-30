<template>
  <a-modal v-model:open="open" title="Welcome" >
    <a-form
        style="margin: 32px 4px"
        ref="formRef"
        :model="formState"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
        autocomplete="off"
    >
      <a-form-item
          label="Username"
          name="username"
          :rules="[{ required: true, message: 'Please input your username!' }]"
      >
        <a-input placeholder="input your username" v-model:value="formState.username" />
      </a-form-item>

      <a-form-item
          label="Password"
          name="password"
          :rules="[{ required: true, message: 'Please input your password!' }]"
      >
        <a-input-password  placeholder="input your password" v-model:value="formState.password" />
      </a-form-item>
    </a-form>
    <template #footer>
      <a-button type="primary" key="back" @click="close" style="float: left">Registry</a-button>
      <a-button key="back" @click="close">Cancel</a-button>
      <a-button type="primary" :loading="confirmLoading" @click="handleOk">Submit</a-button>
    </template>
  </a-modal>
</template>

<script setup>
import { ref,reactive } from 'vue';
const modalText = ref('Content of the modal');
const open = ref(false);
const formRef = ref(null)
const confirmLoading = ref(false);
const show = () => {
  open.value = true;
  if(formRef.value){
    formRef.value.resetFields();
  }
};
const close = ()=>{
  open.value = false;
}

const formState = reactive({
  username: '',
  password: '',
});
const handleOk = () => {
    formRef.value
    .validate()
    .then(() => {
      console.log('values', formState, toRaw(formState));
    })
    .catch(error => {
      console.log('error', error);
    });
};
defineExpose({
  show
})
</script>

<style scoped>

</style>
