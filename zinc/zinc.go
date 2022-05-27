package zinc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"

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
func (s *Zinc) ListIndexes() {
	response, err := s.request().Get(s.Host + "/api/index")
	log.Println(response.String())
	log.Println(err)
}

// UpdateDocumentWithId
// Create/Update a document and index it for searches. Provide a doc Id
func (s *Zinc) UpdateDocumentWithId(target, id string, body interface{}) {
	response, err := s.request().SetBody(body).
		Put(s.Host + fmt.Sprintf("/api/%v/_doc/%v", target, id))
	log.Println(response.String())
	log.Println(err)
}

// UpdateDocument
// Create/Update a document and index it for searches
func (s *Zinc) UpdateDocument(target, body interface{}) {
	response, err := s.request().SetBody(body).
		Put(s.Host + fmt.Sprintf("/api/%v/document", target))
	log.Println(response.String())
	log.Println(err)
}

// Search 搜索
func (s *Zinc) Search(target, body interface{}) (string, error) {
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
	//parse := gjson.Parse(response.String())
	return nil
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
