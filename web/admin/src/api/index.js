import request from '@/utils/request'

// 获取用户详细信息
export function getInfo() {
    return request({
        url: '/captchaImage',
        method: 'get'
    })
}
