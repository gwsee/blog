package oss

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"strings"
	"time"
)

func (o *File) UploadFile(filePath string, content []byte) (string, error) {
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(o.bucket),        // 存储空间名称
		Key:    oss.Ptr(filePath),        // 对象名称
		Body:   bytes.NewReader(content), // 要上传的内容
	}
	_, err := o.cli.PutObject(o.ctx, request)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s%s", o.site, filePath)
	return url, nil
}
func (o *File) Temp(path string, expireMinute int) (string, error) {
	key := strings.ReplaceAll(path, o.site, "")
	result, err := o.cli.Presign(context.TODO(), &oss.GetObjectRequest{
		Bucket: oss.Ptr(o.bucket),
		Key:    oss.Ptr(key),
	}, oss.PresignExpires(time.Duration(expireMinute)*time.Minute))
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
