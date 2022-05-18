package sms

import (
	"os"
	"testing"
)

func TestNewSmsClient(t *testing.T) {
	os.Setenv("SMS_PROVIDER", "Aliyun SMS")
	os.Setenv("SMS_ACCESS_ID", "LTAI4GHirunJfxuaZUDexNkf")
	os.Setenv("SMS_ACCESS_KEY", "SefoGyNkWOIckTXkxxGkfxXpzARwRQ")
	os.Setenv("SMS_SIGN", "天籁奇放")
	os.Setenv("SMS_TEMPLATE", "SMS_192280049")
	//os.Setenv("SMS_PROVIDER", "Aliyun SMS")
	//os.Setenv("SMS_ACCESS_ID", "LTAI4GHPhvCcWyzijNLhmWmi")
	//os.Setenv("SMS_ACCESS_KEY", "S34qdIWZQ9QKgzDglvJUL18TUnAVwJ")
	//os.Setenv("SMS_SIGN", "佳友荟")
	//os.Setenv("SMS_TEMPLATE", "SMS_217427205")
	got, err := NewSmsClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	err = got.SendMessage(map[string]string{
		"code": "111111",
	}, "17606905010")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNewSmsClient2(t *testing.T) {
	os.Setenv("SMS_PROVIDER", "Tencent Cloud SMS")
	os.Setenv("SMS_ACCESS_ID", "AKIDrRknWvnxby4tNevOr64lBKAkqfDXqRvj")
	os.Setenv("SMS_ACCESS_KEY", "fAPwkGVONBk00mHUO1KcbOzIs44mDUVG")
	os.Setenv("SMS_APP_ID", "1400080433")
	os.Setenv("SMS_SIGN", "城边")
	os.Setenv("SMS_TEMPLATE", "860471")

	got, err := NewSmsClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	err = got.SendMessage(map[string]string{
		"code": "666666",
	}, "17606905010")
	if err != nil {
		t.Errorf(err.Error())
	}
}
