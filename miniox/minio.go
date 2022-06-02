package miniox

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Conf struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Token     string
	UseSSL    bool
}

type Minio struct {
	conf   *Conf
	client *minio.Client
}

// NewMinio 初始化 minio
func NewMinio(conf *Conf) (*Minio, error) {
	client, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, conf.Token),
		Secure: conf.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &Minio{conf: conf, client: client}, nil
}

// Client 获取minio client
func (s *Minio) Client() *minio.Client {
	return s.client
}

// PutObject 上传
func (s *Minio) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64) (minio.UploadInfo, error) {
	return s.client.PutObject(context.Background(), bucketName, objectName, reader, objectSize, minio.PutObjectOptions{})
}

// GetObject 获取
func (s *Minio) GetObject(bucketName, objectName string) (*minio.Object, error) {
	return s.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
}

// StatObject 查询对象状态
func (s *Minio) StatObject(bucketName, objectName string) (minio.ObjectInfo, error) {
	return s.client.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
}

// RemoveObject 删除对象
func (s *Minio) RemoveObject(bucketName, objectName string) error {
	return s.client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{
		GovernanceBypass: true,
	})
}

// RemoveIncompleteUpload 在完成上传删除
func (s *Minio) RemoveIncompleteUpload(bucketName, objectName string) error {
	return s.client.RemoveIncompleteUpload(context.Background(), bucketName, objectName)
}

// CopyObject 复制
func (s *Minio) CopyObject(dst minio.CopyDestOptions, src minio.CopySrcOptions) (minio.UploadInfo, error) {
	return s.client.CopyObject(context.Background(), dst, src)
}
