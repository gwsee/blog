import './index.css'
import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import router from "@/router"
import App from './app.vue';
import 'ant-design-vue/dist/reset.css';
import {parseTime} from "@/utils/util";

const app = createApp(App);
app.config.globalProperties.$formatDate = parseTime;
app.use(Antd).use(router).mount('#app');
