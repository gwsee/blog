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
        target: `http://127.0.0.1:7999/`,
        changeOrigin:true,
        rewrite:(path)=>path.replace(/^\/api/,"")
      }
    }
  }
})
