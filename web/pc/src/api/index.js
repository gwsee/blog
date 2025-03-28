import request from '@/utils/request'

// 获取用户详细信息
export function getInfo() {
    return request({
        url: '/captchaImage',
        method: 'get'
    })
}
//名言警句 或者 通知
export function messages(data){
    return request({
        url: "/v1/messages",
        method: "post",
        data:data
    })
}

