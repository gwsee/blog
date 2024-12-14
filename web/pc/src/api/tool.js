import request from '@/utils/request'
// 列表
export function fileUpload(data) {
    return request({
        url: '/v1/upload/file',
        method: 'post',
        data: data,
        headers: {
            'Content-type': 'multipart/form-data',
        },
    })
}
export const  filePrefix = "/api/v1/file/"
