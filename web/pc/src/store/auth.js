import {reactive, computed, ref} from 'vue'
import {info} from "@/api/account";

const state = reactive({
    user: null,
    token: null,
})

const isLoggedIn = ref(false)

const showLogin = ref(false)

const setLoginData = (token) => {
    if(!token){
        state.user = null
        state.token = null
        isLoggedIn.value = false
        return
    }
    localStorage.setItem('token', token)
    info({}).then(res=>{
        if(res&&res.code===200){
            state.user = res.data
        }
    }).finally(()=>{
        isLoggedIn.value = !!token
        state.token = token
    })
}

const setLoginShow=(flag)=>{
    if(!flag){
        showLogin.value = false
        return
    }
    logout()
    showLogin.value = true
}

const logout = () => {
    isLoggedIn.value = false
    state.user = null
    state.token = null
    localStorage.removeItem('token')
}

const getLoginToken=()=>{
   return  localStorage.getItem('token')
}

const checkAuth = () => {
    const token = localStorage.getItem('token')
    if (token) {
        // 这里应该验证 token 是否有效
        // 如果有效，设置用户信息
        state.token = token
        isLoggedIn.value = true
        // 获取用户信息的逻辑
    }
}

export const useAuthStore = () => ({
    state,
    isLoggedIn,
    showLogin,
    setLoginShow,
    setLoginData,
    getLoginToken,
    logout,
    checkAuth,
})
