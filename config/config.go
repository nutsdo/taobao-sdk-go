package config

import (
	"fmt"
	"github.com/nutsdo/taobao-sdk-go/auth"
	"os"
	"time"
)

type Config struct {
	AppKey       string
	TargetAppKey string
	SignMethod   string
	//Session      string
	Timestamp    string
	Format       string
	V            string
	PartnerId    string
	Simplify     string
}

var DefaultConfig = Config{
	AppKey:     os.Getenv("APP_KEY"),
	SignMethod: "md5",
	Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
	Format:     "json",
	V:          "2.0",
}

func NewConfigWithCredential(credential *auth.Credentials, session string) *Config {
	config := &Config{
		AppKey:     credential.AppKey,
		SignMethod: "md5",
		//Session:    session,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Format:     "json",
		V:          "2.0",
	}
	fmt.Printf("输出配置:%#v", config)
	return config
}

func (c *Config) SetAppKey(appKey string) *Config {
	c.AppKey = appKey
	return c
}
