import { defineConfig } from 'vite'
import {resolve} from 'path'
import vue from '@vitejs/plugin-vue'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  css: {
    devSourcemap: true,
    postcss: {
      plugins: [
        tailwindcss,
        autoprefixer,
      ],
    },
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, "./src"),
    }
  },
  server:{
    hmr: true,
    https:false,
    host: '127.0.0.1',
    open: true,
    proxy:{
      "/api":{
        target: `http://127.0.0.1:1000/`,
        // target: `http://blog.gwsee.com:8080/`,
        changeOrigin:true,
        rewrite:(path)=>path.replace(/^\/api/,"")
      }
    }
  }
})
