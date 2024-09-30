import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import App from './App.vue';
import 'ant-design-vue/dist/reset.css';
import router from "@/router"

// import Login from "@/components/Login.vue"
// const dialogInstance = createApp(Login).mount(document.createElement('div'));
// 将 dialog 组件添加到 DOM 中

// Vue.prototype.$login = dialogInstance
const app = createApp(App);

//.directive("$login",Login)
app.use(Antd).use(router).mount('#app');
