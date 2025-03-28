import request from '@/utils/request'
// 列表
export function fileUpload(data) {
    return request({
        url: '/v1/upload/file',
        method: 'post',
        data: data,
        IsAllowedRepeat:true,
        headers: {
            'Content-type': 'multipart/form-data',
        },
    })
}
export function fileOssToken(data) {
    return request({
        url: '/v1/upload/oss/token',
        method: 'post',
        data: data,
    })
}
export function fileOssSave(data) {
    return request({
        url: '/v1/upload/oss/save',
        method: 'post',
        data: data,
    })
}
export const  filePrefix = "/api/v1/file/"
