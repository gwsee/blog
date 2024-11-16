import {reactive, computed, ref} from 'vue'

const state = reactive({
    user: null,
    token: null,
})

const isLoggedIn = computed(() => !!getLoginToken())

const showLogin = ref(false)

const setLoginData = (userData, token) => {
    state.user = userData
    state.token = token
    localStorage.setItem('token', token)
}

const setLoginShow=(flag)=>{
    if(!flag){
        showLogin.value = false
        return
    }
    if(isLoggedIn.value){
        logout()
    }
    showLogin.value = true
}

const logout = () => {
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
