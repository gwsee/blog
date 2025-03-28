package oss

import (
	"blog/api/global"
	"bytes"
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/sts-20150401/v2/client"
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
	if strings.HasPrefix(path, "http:") {
		path = strings.Replace(path, "http:", "https:", 1)
	}
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

func (o *File) UploadToken(conf *global.Oss) (*client.AssumeRoleResponseBodyCredentials, error) {
	req := client.AssumeRoleRequest{
		DurationSeconds: oss.Ptr(conf.Expire),
		RoleArn:         oss.Ptr(conf.Arn),
		RoleSessionName: oss.Ptr("oss-client-upload"),
	}
	cli, err := client.NewClient(&openapi.Config{
		AccessKeyId:     &conf.Key,
		AccessKeySecret: &conf.Secret,
		RegionId:        &conf.Region,
	})
	if err != nil {
		return nil, err
	}
	res, err := cli.AssumeRole(&req)
	if err != nil {
		return nil, err
	}
	return res.Body.Credentials, nil
}
