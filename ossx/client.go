package ossx

import (
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// 类型别名定义，简化导入
type Config = oss.Config
type Options = oss.Options
type AppendOptions = oss.AppendOptions
type PutObjectRequest = oss.PutObjectRequest
type PutObjectResult = oss.PutObjectResult
type AppendObjectRequest = oss.AppendObjectRequest
type AppendObjectResult = oss.AppendObjectResult
type AppendOnlyFile = oss.AppendOnlyFile

// Client 定义了OSS客户端的接口
type Client interface {
	// PutObject 上传对象到OSS
	PutObject(*PutObjectRequest, ...func(*Options)) (*PutObjectResult, error)
	// PutObjectFromFile 从本地文件上传对象到OSS
	PutObjectFromFile(string, *PutObjectRequest, ...func(*Options)) (*PutObjectResult, error)
	// AppendObject 追加对象到OSS
	AppendObject(*AppendObjectRequest, ...func(*Options)) (*AppendObjectResult, error)
	// AppendFile 追加文件到OSS
	AppendFile(string, string, ...func(*AppendOptions)) (*AppendOnlyFile, error)
}

// ClientEntity 实现了Client接口
type ClientEntity struct {
	*Config
	ossClient *oss.Client
}

func New(config *Config) Client {
	client := &ClientEntity{}
	client.Config = config
	client.ossClient = oss.NewClient(config)
	return client
}

// NewCredentialsProvider 创建静态凭据提供者
func NewCredentialsProvider(accessKeyId, accessKeySecret string) credentials.CredentialsProvider {
	return credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret)
}

func (c *ClientEntity) PutObject(req *PutObjectRequest, optFns ...func(*Options)) (*PutObjectResult, error) {
	return c.ossClient.PutObject(context.TODO(), req, optFns...)
}

func (c *ClientEntity) PutObjectFromFile(localFile string, req *PutObjectRequest, optFns ...func(*Options)) (*PutObjectResult, error) {
	return c.ossClient.PutObjectFromFile(context.TODO(), req, localFile, optFns...)
}

func (c *ClientEntity) AppendObject(req *AppendObjectRequest, optFns ...func(*Options)) (*AppendObjectResult, error) {
	return c.ossClient.AppendObject(context.TODO(), req, optFns...)
}

func (c *ClientEntity) AppendFile(bucket string, key string, optFns ...func(*AppendOptions)) (*AppendOnlyFile, error) {
	return c.ossClient.AppendFile(context.TODO(), bucket, key, optFns...)
}
