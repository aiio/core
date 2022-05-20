package config

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type configs struct {
	AppName   string
	AppDebug  bool
	HttpPort  string
	Database  Database
	Jwt       JWT
	Redis     Redis
	Mail      Mail
	MiniParam MiniParam
}

func (conf *configs) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

var V *configs

func init() {
	V = &configs{
		AppName:  GetDefaultEnv("APP_NAME", "cnbattle"),
		AppDebug: GetEnvToBool("APP_DEBUG"),
		HttpPort: GetDefaultEnv("HTTP_PORT", "1993"),
		Database: Database{
			Engine: GetDefaultEnv("MSQL_ENGINE", "mysql"),
			DSN:    GetDefaultEnv("MSQL_DSN", "root:123456(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"),
			Prefix: GetDefaultEnv("MSQL_PREFIX", ""),
		},
		Jwt: JWT{
			SecretKey: GetDefaultEnv("JWT_SECRET_KEY", "cnbattle"),
			Exp:       GetDefaultEnvToInt("JWT_SECRET_EXP", 7200),
		},
		Redis: Redis{
			Host: GetDefaultEnv("REDIS_HOST", "127.0.0.1:6379"),
			Pass: GetDefaultEnv("REDIS_PASSWORD", ""),
			DB:   GetDefaultEnvToInt("REDIS_DB", 0),
		},
		Mail: Mail{
			User:     GetDefaultEnv("MAIL_USER", "sc_7byq8v_test_eCxYAD"),
			Pass:     GetDefaultEnv("MAIL_PASS", "b144480c16c8ce505ef05000129ced22"),
			Host:     GetDefaultEnv("MAIL_HOST", "smtp.sendcloud.net"),
			Port:     GetDefaultEnvToInt("MAIL_PORT", 587),
			From:     GetDefaultEnv("MAIL_FROM", "test@RWjiOgcIFCvHxmbHCjXX4eQe33YAbrVB.sendcloud.org"),
			FromName: GetDefaultEnv("MAIL_FROM_NAME", "cnbattle"),
		},
		MiniParam: MiniParam{
			AppId:     GetDefaultEnv("MINI_APP_ID", ""),
			AppSecret: GetDefaultEnv("MINI_APP_SECRET", ""),
		},
	}
	if V.AppDebug {
		fmt.Printf("configs:\n %v\n", V.String())
	}
}
