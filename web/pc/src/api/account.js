import request from '@/utils/request'

// 登录
export function login(data) {
    return request({
        url: '/v1/login',
        method: 'post',
        data: data
    })
}

// 注册
export function register(data) {
    return request({
        url: '/v1/register',
        method: 'post',
        data: data
    })
}

// 重置密码
export function resetPass(data) {
    return request({
        url: '/v1/reset_password',
        method: 'post',
        data: data
    })
}
