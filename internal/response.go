/*
 * SPDX-License-Identifier: Apache-2.0 License
 * Author: cnbattle  <qiaicn@gmail.com>
 * Copyright (c) 2022.
 */

package internal

import (
	"encoding/json"
	"log"
)

type HttpResult struct {
	Message string
}

func (r *HttpResult) ToBytes() (bytes []byte) {
	var err error
	bytes, err = json.Marshal(r)
	if err != nil {
		log.Fatal("数据转换失败")
	}
	return
}
