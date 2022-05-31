// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sms

import (
	"fmt"

	"github.com/aiio/core/config"
)

const (
	Aliyun       = "Aliyun SMS"
	TencentCloud = "Tencent Cloud SMS"
	VolcEngine   = "Volc Engine SMS"
	Huyi         = "Huyi SMS"
)

type Client interface {
	SendMessage(param map[string]string, targetPhoneNumber ...string) error
}

func NewSmsClient() (Client, error) {
	switch config.V.SMS.Provider {
	case Aliyun:
		return NewAliyunClient(config.V.SMS.AccessID, config.V.SMS.AccessKey,
			config.V.SMS.Sign, config.V.SMS.Template)
	case TencentCloud:
		return NewTencentClient(config.V.SMS.AccessID, config.V.SMS.AccessKey,
			config.V.SMS.Sign, config.V.SMS.Template, config.V.SMS.AppID)
	case VolcEngine:
		return NewVolcClient(config.V.SMS.AccessID, config.V.SMS.AccessKey,
			config.V.SMS.Sign, config.V.SMS.Template, config.V.SMS.Account)
	case Huyi:
		return NewHuyiClient(config.V.SMS.AccessID, config.V.SMS.AccessKey, config.V.SMS.Template)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", config.V.SMS.Provider)
	}
}
