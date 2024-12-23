import {fileUpload,filePrefix} from "@/api/tool.js";

export function parseTime(time, pattern) {
    if (arguments.length === 0 || !time) {
        return null
    }
    const format = pattern || '{y}-{m}-{d} {h}:{i}:{s}'
    let date
    if (typeof time === 'object') {
        date = time
    } else {
        if ((typeof time === 'string') && (/^[0-9]+$/.test(time))) {
            time = parseInt(time)
        } else if (typeof time === 'string') {
            time = time.replace(new RegExp(/-/gm), '/').replace('T', ' ').replace(new RegExp(/\.[\d]{3}/gm), '');
        }
        if ((typeof time === 'number') && (time.toString().length === 10)) {
            time = time * 1000
        }
        date = new Date(time)
    }
    const formatObj = {
        y: date.getFullYear(),
        m: date.getMonth() + 1,
        d: date.getDate(),
        h: date.getHours(),
        i: date.getMinutes(),
        s: date.getSeconds(),
        a: date.getDay()
    }
    const time_str = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, key) => {
        let value = formatObj[key]
        // Note: getDay() returns 0 on Sunday
        if (key === 'a') { return ['日', '一', '二', '三', '四', '五', '六'][value] }
        if (result.length > 0 && value < 10) {
            value = '0' + value
        }
        return value || 0
    })
    return time_str
}
export function fileFull(uuid) {
    if (!uuid) {
        return uuid
    }
    if(uuid.indexOf('http') !== -1) {
        return uuid
    }
    return filePrefix+uuid
}
export function  customRequest ({file, onSuccess,onError})  {
    let formData = new FormData();
    formData.append('file',file)
    fileUpload(formData).then(res=>{
        if(res&&res.code===200){
            const item = {uuid:res.data.uuid,name:file.name,url:fileFull(res.data.uuid)}
            onSuccess(item)
        }else{
            onError(res.message||'上传失败')
        }
    }).finally(()=>{
    }).catch((e)=>{
        onError(e)
    })
}