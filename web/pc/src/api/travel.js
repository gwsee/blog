import request from '@/utils/request'
// 列表
export function travelList(data) {
    return request({
        url: '/v1/travel/list',
        method: 'post',
        data: data
    })
}
export function travelDel(data) {
    return request({
        url: '/v1/travel/del',
        method: 'post',
        data: data
    })
}
export function travelSave(data) {
    return request({
        url: '/v1/travel/save',
        method: 'post',
        data: data
    })
}
export function travelGet(data) {
    return request({
        url: '/v1/travel/get',
        method: 'post',
        data: data
    })
}
export function travelThumb(data) {
    return request({
        url: '/v1/travel/thumb',
        method: 'post',
        data: data
    })
}
export function travelCollect(data) {
    return request({
        url: '/v1/travel/collect',
        method: 'post',
        data: data
    })
}
export function toggleFavorite(){}
export function toggleLike(){}
