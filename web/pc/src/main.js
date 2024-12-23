import './index.css'
import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import router from "@/router"
import App from './app.vue';
import 'ant-design-vue/dist/reset.css';
import {parseTime,fileFull,customRequest} from "@/utils/util";
const app = createApp(App);
app.config.globalProperties.$formatDate = parseTime;
app.config.globalProperties.$fileFull = fileFull;
app.config.globalProperties.$customRequest = customRequest;
app.use(Antd).use(router).mount('#app');
