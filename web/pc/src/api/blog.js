import request from '@/utils/request'

// 列表
export function blogList(data) {
    return request({
        url: '/v1/blogs/list',
        method: 'post',
        data: data
    })
}
// 创建
export function blogCreate(data) {
    return request({
        url: '/v1/blogs/create',
        method: 'post',
        data: data
    })
}
// 编辑
export function blogUpdate(data) {
    return request({
        url: '/v1/blogs/update',
        method: 'post',
        data: data
    })
}
// 删除
export function blogDel(data) {
    return request({
        url: '/v1/blogs/del',
        method: 'post',
        data: data
    })
}
// 查询
export function blogGet(data) {
    return request({
        url: '/v1/blogs/get',
        method: 'post',
        data: data
    })
}

// 评论列表
export function blogCommentList(data) {
    return request({
        url: '/v1/blogs_comment/list',
        method: 'post',
        data: data
    })
}
// 评论创建
export function blogCommentCreate(data) {
    return request({
        url: '/v1/blogs_comment/create',
        method: 'post',
        data: data
    })
}
// 评论编辑
export function blogCommentUpdate(data) {
    return request({
        url: '/v1/blogs_comment/update',
        method: 'post',
        data: data
    })
}
// 评论删除
export function blogCommentDel(data) {
    return request({
        url: '/v1/blogs_comment/del',
        method: 'post',
        data: data
    })
}
// 评论查询
export function blogCommentGet(data) {
    return request({
        url: '/v1/blogs_comment/get',
        method: 'post',
        data: data
    })
}
