import request from '@/utils/request'
// 列表
export function blogList(data) {
    return request({
        url: '/v1/blogs/list',
        method: 'post',
        data: data
    })
}
