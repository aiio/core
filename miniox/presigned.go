package miniox

import (
	"context"
	"net/url"
	"time"
)

// PresignedGetObject 预获取签名
func (s *Minio) PresignedGetObject(bucketName string, objectName string, expires time.Duration, reqParams url.Values) (*url.URL, error) {
	return s.client.PresignedGetObject(context.Background(), bucketName, objectName, expires, reqParams)
}

// PresignedPutObject 预上传签名
func (s *Minio) PresignedPutObject(bucketName string, objectName string, expires time.Duration) (*url.URL, error) {
	return s.client.PresignedPutObject(context.Background(), bucketName, objectName, expires)
}

// PresignedHeadObject 返回一个预签名 URL来访问
func (s *Minio) PresignedHeadObject(bucketName string, objectName string, expires time.Duration, reqParams url.Values) (*url.URL, error) {
	return s.client.PresignedHeadObject(context.Background(), bucketName, objectName, expires, reqParams)
}
