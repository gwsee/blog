import request from '@/utils/request'
// 列表
export function photosList(data) {
    return request({
        url: '/v1/photos',
        method: 'post',
        data: data
    })
}
