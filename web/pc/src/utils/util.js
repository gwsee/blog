import SparkMD5 from "spark-md5";
import OSS from 'ali-oss'
import {filePrefix,fileOssToken,fileOssSave, fileUpload} from "@/api/tool.js";
export function parseTime(time, pattern) {
    if (arguments.length === 0 || !time) {
        return "至今"
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
//最开始的服务端上传 非必须就不使用了
export function  customRequestOld ({file, onSuccess,onError})  {
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
//oss客户端上传:适用于vue的action
export  function  customRequest ({file, onSuccess,onError})  {
    return new Promise((resolve,reject)=>{
        (async ()=>{
            try {
                //第一步获取md5数据
                const md5Obj = await processMd5File(file);
                if (md5Obj.error) {
                    onError&&onError(md5Obj.error)
                    reject(md5Obj.error)
                    return
                }
                const item = {uuid:md5Obj.id,name:md5Obj.name,fileName:md5Obj.name,url:fileFull(md5Obj.id),fileUrl:fileFull(md5Obj.id)}
                //第二步判断是否存在
                const tkObj = await checkFile(md5Obj.id)
                if (tkObj.error) {
                    onError&&onError(tkObj.error)
                    reject(tkObj.error)
                    return
                }
                if (tkObj.exist){
                    onSuccess&&onSuccess(item)
                    resolve(item)
                    return
                }
                //第三步上传oss
                const upObj = await uploadToOss(file,tkObj,md5Obj)
                if (upObj.error) {
                    onError&&onError(upObj.error)
                    reject(upObj.error)
                    return
                }
                //第四步保存数据
                const res = await saveFile(md5Obj)
                if (res.error){
                    onError&&onError(res.error)
                    reject(res.error)
                    return
                }
                onSuccess&&onSuccess(item)
                resolve(item)
            }catch (error){
                onError&&onError(error)
                reject(error)
            }
        })()
    })

}

async function saveFile(file){
    return await fileOssSave(file).then(res=>{
        if(res&&res.code===200){
            return res.data
        }else{
            return {error:res.message||"保存数据失败"}
        }
    }).catch((e)=>{
        return {error:e.toString()}
    })
}
async function uploadToOss(file,tkObj,md5Obj){
    const client = new OSS({
        // 将<YOUR_BUCKET>设置为OSS Bucket名称。
        bucket: tkObj.bucket,
        // 将<YOUR_REGION>设置为OSS Bucket所在地域，例如oss-cn-hangzhou。
        region: "oss-"+tkObj.region,
        // 如果这里使用的是本地获取的ststoken的参数，请将credentials.xxx手动替换为已获取到参数的具体数值
        accessKeyId: tkObj.key_id,
        accessKeySecret: tkObj.key_secret,
        stsToken: tkObj.token,
    });
    const result = await client.put(md5Obj.path, file);
    if(result&&result.url&&result.res&&result.res.status&&result.res.status===200){
        md5Obj.path=result.url;
        return md5Obj
    }else{
        return {error:result.toString()}
    }
}
async function  checkFile (uuid) {
    return await fileOssToken({id: uuid}).then((res) => {
        if(res&&res.code===200){
            return res.data
        }else{
            return {error:res.message||"获取token失败"}
        }
    }).catch((e)=>{
        return {error:e.toString()}
    })
}

/**
 * 计算文件 MD5，并获取文件信息
 * @param {File} file - 用户上传的文件对象
 * @returns {Promise<{ md5: string, filename: string, type: string, size: number, storagePath: string } | { error: string }>}
 */
async function processMd5File(file) {
    return new Promise((resolve) => {
        if (!file) {
            resolve({ error: "文件不能为空" });
            return;
        }

        const chunkSize = 2 * 1024 * 1024; // 2MB 分片
        const chunks = Math.ceil(file.size / chunkSize);
        const spark = new SparkMD5.ArrayBuffer();
        const reader = new FileReader();
        let currentChunk = 0;

        reader.onload = (event) => {
            spark.append(event.target.result);
            currentChunk++;

            if (currentChunk < chunks) {
                loadNext();
            } else {
                const md5 = spark.end();
                const fileExt = getFileExtension(file.name);
                const fileType = file.type.split("/")[0] || "unknown"; // 取主 MIME 类型，如 "image"
                const formattedDate = new Date().toISOString().slice(0, 10).replace(/-/g, ""); // YYYYMMDD
                const newFilename = `${md5}.${fileExt}`;
                const storagePath = formatPath(fileType, fileExt, formattedDate, newFilename);

                resolve({
                    id:md5,
                    name: file.name,
                    type: fileType,
                    size: file.size,
                    path: storagePath,
                });
            }
        };
        reader.onerror = () => resolve({ error: "文件读取失败" });
        function loadNext() {
            const start = currentChunk * chunkSize;
            const end = Math.min(start + chunkSize, file.size);
            reader.readAsArrayBuffer(file.slice(start, end));
        }
        loadNext();
    });
}
export async function calculateMD5(file){
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = async (event) => {
            try {
                const arrayBuffer = event.target.result;
                const hashBuffer = await crypto.subtle.digest("MD5", arrayBuffer);
                const hashArray = Array.from(new Uint8Array(hashBuffer));
                const hashHex = hashArray.map((b) => b.toString(16).padStart(2, "0")).join("");
                resolve(hashHex);
            } catch (error) {
                reject(error);
            }
        };
        reader.onerror = (error) => reject(error);
        reader.readAsArrayBuffer(file);
    });
}
const filePathTemplate = "{type}/{ext}/{date}/{filename}"

/**
 * 格式化存储路径
 * @param {string} template - 文件路径模板
 * @param {string} type - 文件类型
 * @param {string} ext - 文件扩展名
 * @param {string} date - 格式化日期 (YYYYMMDD)
 * @param {string} filename - 新的 MD5 文件名
 * @returns {string} - 生成的存储路径
 */
function formatPath(type, ext, date, filename) {
    return filePathTemplate
        .replace("{type}", type)
        .replace("{ext}", ext)
        .replace("{date}", date)
        .replace("{filename}", filename);
}
/**
 * 获取文件扩展名
 * @param {string} filename
 * @returns {string}
 */
function getFileExtension(filename) {
    const extMatch = filename.match(/\.([a-zA-Z0-9]+)$/);
    return extMatch ? extMatch[1] : "unknown";
}
