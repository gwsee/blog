import request from '@/utils/request'

export function palacesCreate(data) {
    return request({
        url: '/v1/palaces/create',
        method: 'post',
        data: data
    })
}
export function palacesUpdate(data) {
    return request({
        url: '/v1/palaces/update',
        method: 'post',
        data: data
    })
}
export function palacesDelete(data) {
    return request({
        url: '/v1/palaces/delete',
        method: 'post',
        data: data
    })
}
export function palacesGet(data) {
    return request({
        url: '/v1/palaces/get',
        method: 'post',
        data: data
    })
}
export function palacesList(data) {
    return request({
        url: '/v1/palaces/list',
        method: 'post',
        data: data
    })
}
//memo相关
export function memoSave(data) {
    return request({
        url: '/v1/palaces/memo/save',
        method: 'post',
        data: data
    })
}
export function memoDone(data) {
    return request({
        url: '/v1/palaces/memo/done',
        method: 'post',
        data: data
    })
}
export function memoDelete(data) {
    return request({
        url: '/v1/palaces/memo/delete',
        method: 'post',
        data: data
    })
}
export function memoList(data) {
    return request({
        url: '/v1/palaces/memo/list',
        method: 'post',
        data: data
    })
}
//todo相关
export function todoSave(data) {
    return request({
        url: '/v1/palaces/todo/save',
        method: 'post',
        data: data
    })
}
export function todoDone(data) {
    return request({
        url: '/v1/palaces/todo/done',
        method: 'post',
        data: data
    })
}
export function todoDelete(data) {
    return request({
        url: '/v1/palaces/todo/delete',
        method: 'post',
        data: data
    })
}
export function todoDoneDelete(data) {
    return request({
        url: '/v1/palaces/todo-done/delete',
        method: 'post',
        data: data
    })
}
export function todoRecordDelete(data) {
    return request({
        url: '/v1/palaces/todo-record/delete',
        method: 'post',
        data: data
    })
}
export function todoList(data) {
    return request({
        url: '/v1/palaces/todo/list',
        method: 'post',
        data: data
    })
}
export function todoDoneList(data) {
    return request({
        url: '/v1/palaces/todo-done/list',
        method: 'post',
        data: data
    })
}
