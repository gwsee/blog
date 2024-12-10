package oss

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

type File struct {
	cli    *oss.Client
	bucket string
	ctx    context.Context
	site   string
}

func NewFileClient(ctx context.Context) *File {
	f := &File{ctx: ctx}
	return f
}
func (o *File) SetOSSClient(key, secret, region string) *File {
	o.cli = NewOSSClient(key, secret, region)
	return o
}
func (o *File) SetBucket(bucket string) *File {
	o.bucket = bucket
	return o
}
func (o *File) SetSite(site string) *File {
	o.site = site
	return o
}
func NewOSSClient(key, secret, region string) *oss.Client {
	provider := credentials.CredentialsProviderFunc(func(ctx context.Context) (credentials.Credentials, error) {
		// 返回长期凭证
		return credentials.Credentials{AccessKeyID: key, AccessKeySecret: secret}, nil
		// 返回临时凭证
		//return credentials.Credentials{AccessKeyID: "id", AccessKeySecret: "secret",    SecurityToken: "token"}, nil
	})
	//// 定义要上传的内容
	//credentials.NewAnonymousCredentialsProvider()
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(region)
	return oss.NewClient(cfg)
}
