import request from '@/utils/request'

export function userSave(data) {
    return request({
        url: '/v1/user/save',
        method: 'post',
        data: data
    })
}
export function userGet(data) {
    return request({
        url: '/v1/user/get',
        method: 'post',
        data: data
    })
}
// 列表
export function projectList(data) {
    return request({
        url: '/v1/project/list',
        method: 'post',
        data: data
    })
}
export function projectDel(data) {
    return request({
        url: '/v1/project/del',
        method: 'post',
        data: data
    })
}
export function projectGet(data) {
    return request({
        url: '/v1/project/get',
        method: 'post',
        data: data
    })
}
export function projectSave(data) {
    return request({
        url: '/v1/project/save',
        method: 'post',
        data: data
    })
}

export function experienceList(data) {
    return request({
        url: '/v1/experience/list',
        method: 'post',
        data: data
    })
}
export function experienceSave(data) {
    return request({
        url: '/v1/experience/save',
        method: 'post',
        data: data
    })
}
export function experienceGet(data) {
    return request({
        url: '/v1/experience/get',
        method: 'post',
        data: data
    })
}
export function experienceDel(data) {
    return request({
        url: '/v1/experience/del',
        method: 'post',
        data: data
    })
}
