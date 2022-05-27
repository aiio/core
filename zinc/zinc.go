package zinc

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Zinc struct {
	Host     string
	UserId   string
	Password string
}

// NewZinc
// https://docs.zincsearch.com/API%20Reference/
func NewZinc(host string, userId string, password string) *Zinc {
	return &Zinc{Host: host, UserId: userId, Password: password}
}

// ListIndexes List existing indexes
//func (s *Zinc) ListIndexes() {
//	response, err := s.request().Get(s.Host + "/api/index")
//	log.Println(response.String())
//	log.Println(err)
//}

// UpdateDocumentWithId 根据ID创建/更新文档
func (s *Zinc) UpdateDocumentWithId(target, id string, body interface{}) error {
	response, err := s.request().SetBody(body).
		Put(s.Host + fmt.Sprintf("/api/%v/_doc/%v", target, id))
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return errors.New(response.String())
	}
	return nil
}

// UpdateDocument 创建/更新文档
func (s *Zinc) UpdateDocument(target string, body interface{}) error {
	response, err := s.request().SetBody(body).
		Put(s.Host + fmt.Sprintf("/api/%v/document", target))
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return errors.New(response.String())
	}
	return nil
}

// UpdateDocumentsBulk 批量上传文档
func (s *Zinc) UpdateDocumentsBulk(body interface{}) error {
	response, err := s.request().SetBody(body).
		Post(s.Host + fmt.Sprintf("/api/_bulk"))
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return errors.New(response.String())
	}
	return nil
}

// Search 搜索
func (s *Zinc) Search(target string, body interface{}) (string, error) {
	response, err := s.request().SetBody(body).
		Post(s.Host + fmt.Sprintf("/api/%v/_search", target))
	return response.String(), err
}

// DeleteDocument 删除一个文档
func (s *Zinc) DeleteDocument(target, id string) error {
	response, err := s.request().
		Delete(s.Host + fmt.Sprintf("/api/%v/_doc/%v", target, id))
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return errors.New(response.String())
	}
	return nil
}

// Version Get current version of ZincSearch
func (s *Zinc) Version() (Version, error) {
	var resp Version
	response, err := s.request().SetResult(&resp).
		Get(s.Host + fmt.Sprintf("/version"))
	if err != nil {
		return Version{}, err
	}
	if response.StatusCode() != 200 {
		return Version{}, errors.New(response.String())
	}
	return resp, nil
}

func (s *Zinc) request() *resty.Request {
	token := base64.StdEncoding.EncodeToString([]byte(s.UserId + ":" + s.Password))
	client := resty.New()
	return client.R().SetAuthToken(token).SetAuthScheme("Basic")
}

//CreateIndex
//DeleteIndex
//ListIndexes
//Search
//UpdateDocumentWithId
//UpdateDocument
//DeleteDocument
//UpdateDocumentsBulk
//UpdateIndexMappings
//GetIndexMappings
//Version
//Metrics
//CreateUpdateUser
//DeleteUser
