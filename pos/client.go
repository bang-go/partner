package pos

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	*oss.Config
}
type Options = oss.Options
type AppendOptions = oss.AppendOptions
type PutObjectRequest = oss.PutObjectRequest
type PutObjectResult = oss.PutObjectResult
type AppendObjectRequest = oss.AppendObjectRequest
type AppendObjectResult = oss.AppendObjectResult
type AppendOnlyFile = oss.AppendOnlyFile
type Client interface {
	PutObject(*PutObjectRequest, ...func(*Options)) (*PutObjectResult, error)
	PutObjectFromFile(string, *PutObjectRequest, ...func(*Options)) (*PutObjectResult, error)
	AppendObject(*AppendObjectRequest, ...func(*Options)) (*AppendObjectResult, error)
	AppendFile(string, string, ...func(*AppendOptions)) (*AppendOnlyFile, error)
}

type ClientEntity struct {
	*Config
	ossClient *oss.Client
}

func New(config *Config) Client {
	client := &ClientEntity{}
	config.CredentialsProvider = credentials.NewStaticCredentialsProvider(config.AccessKeyId, config.AccessKeySecret)
	client.Config = config
	client.ossClient = oss.NewClient(config.Config)
	return client
}

func (s *ClientEntity) PutObject(req *PutObjectRequest, optFns ...func(*Options)) (*PutObjectResult, error) {
	return s.ossClient.PutObject(context.TODO(), req, optFns...)
}

func (s *ClientEntity) PutObjectFromFile(localFile string, req *PutObjectRequest, optFns ...func(*Options)) (*PutObjectResult, error) {
	return s.ossClient.PutObjectFromFile(context.TODO(), req, localFile, optFns...)
}

func (s *ClientEntity) AppendObject(req *AppendObjectRequest, optFns ...func(*Options)) (*AppendObjectResult, error) {
	return s.ossClient.AppendObject(context.TODO(), req, optFns...)
}

func (s *ClientEntity) AppendFile(bucket string, key string, optFns ...func(*AppendOptions)) (*AppendOnlyFile, error) {
	return s.ossClient.AppendFile(context.TODO(), bucket, key, optFns...)
}
