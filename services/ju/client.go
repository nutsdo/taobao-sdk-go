package ju

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/client"
	"github.com/nutsdo/taobao-sdk-go/config"
	"github.com/nutsdo/taobao-sdk-go/services"
)

const host = "https://eco.taobao.com/router/rest"

type TbClient struct {
	credential *auth.Credentials
	client     client.Client
	config     config.Config
	session    string
}

func NewTbClient() *TbClient {
	c := &TbClient{}
	return c
}

func (c *TbClient) SetSession(session string) {
	c.session = session
}

func NewClientWith(credential *auth.Credentials) *TbClient {

	return &TbClient{
		credential: credential,
		client:     client.DefaultClient,
		config:     config.DefaultConfig,
	}
}

func (c *TbClient) DoRequest(request services.TaoBaoRequest) (resp *http.Response, err error) {

	query := c.QueryBuilder(request)

	fmt.Println(query)
	reqUrl := host + "?" + query
	req, err := http.NewRequest("POST", reqUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err = c.client.Do(nil, req)
	if err != nil {
		return
	}
	return
}

func (c *TbClient) paramsBuilder(request services.TaoBaoRequest) map[string]string {

	params := make(map[string]string)

	//公共参数
	params["method"] = request.Method()
	params["app_key"] = c.credential.AppKey
	params["target_appkey"] = c.config.TargetAppKey
	params["sign_method"] = c.config.SignMethod
	params["timestamp"] = c.config.Timestamp
	params["format"] = c.config.Format
	params["v"] = c.config.V
	params["partner_id"] = c.config.PartnerId
	params["simplify"] = c.config.Simplify

	params["session"] = c.session
	//将业务参数添加到参数列表
	for k, _ := range request.GetValues() {
		params[k] = request.GetValues().Get(k)
	}
	return params
}

func (c *TbClient) QueryBuilder(request services.TaoBaoRequest) string {

	params := c.paramsBuilder(request)

	query := url.Values{}

	for k, v := range params {

		if v != "" {
			query.Add(k, v)
		}
	}
	query.Add("sign", c.sign(params))

	return query.Encode()
}

// sign appkey appsecret params config
func (c *TbClient) sign(params map[string]string) string {
	//
	var keys []string

	//参数key列表
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}
	//数组按ASCII字典升序排序
	sort.Strings(keys)
	//拼接字符串
	var query strings.Builder

	if params["sign_method"] == "md5" {

		query.WriteString(string(c.credential.AppSecret))
	}
	for _, key := range keys {
		query.WriteString(key)
		v, _ := params[key]
		if v != "" {
			query.WriteString(v)
		}
	}
	//fmt.Println(query.String())
	var sign string
	if c.config.SignMethod == "hmac" {
		//生成签名方式:HMAC_MD5
		sign = fmt.Sprintf("%x", hmac.New(md5.New, []byte(c.credential.AppSecret)))
	} else {
		//生成签名方式:MD5
		query.WriteString(string(c.credential.AppSecret))
		sign = fmt.Sprintf("%x", md5.Sum([]byte(query.String())))
	}
	//fmt.Println(sign)
	return strings.ToUpper(sign)

}
