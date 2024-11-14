import { defineConfig } from 'vite'
import {resolve} from 'path'
import vue from '@vitejs/plugin-vue'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, "./src")
    }
  },
  server:{
    https:false,
    proxy:{
      "/api":{
        target: `http://47.100.14.57:8080/prod3101`,
        changeOrigin:true,
        rewrite:(path)=>path.replace(/^\/api/,"")
      }
    }
  }
})
