package config_test

import (
	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/config"
	"testing"
)

func TestConfig(t *testing.T)  {
	credential := &auth.Credentials{
		AppKey:"12312312",
		AppSecret:[]byte("123123adasds3"),
	}
	session := "1123213123"
	cfg := config.NewConfigWithCredential(credential, session)
	t.Log(cfg)
}