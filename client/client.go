package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var UserAgent = "Golang taobao/client package"
var DefaultClient = Client{&http.Client{Transport: http.DefaultTransport}}

// 用来打印调试信息
var DebugMode = false

// Client 负责发送HTTP请求到淘宝接口服务器
type Client struct {
	*http.Client
}

// TurnOnDebug 开启Debug模式
func TurnOnDebug() {
	DebugMode = true
}

func newRequest(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader) (req *http.Request,err error) {
	req,err = http.NewRequest(method, reqUrl, body)
	if err != nil {
		return
	}
	if headers == nil {
		headers = http.Header{}
	}

	req.Header = headers
	req = req.WithContext(ctx)

	//check access token

	return
}

func (r Client) DoRequest(ctx context.Context, method, reqUrl string, headers http.Header) (resp *http.Response, err error) {
	req, err := newRequest(ctx, method, reqUrl, headers, nil)
	if err !=nil {
		return
	}
	return r.Do(ctx, req)
}

func (r Client) DoRequestWith(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil{
		return
	}
	req.ContentLength = int64(bodyLength)
	return r.Do(ctx, req)
}

func (r Client) DoRequestWith64(ctx context.Context, method, reqUrl string, headers http.Header, body io.Reader,
	bodyLength int64) (resp *http.Response, err error) {

	req, err := newRequest(ctx, method, reqUrl, headers, body)
	if err != nil{
		return
	}
	req.ContentLength = bodyLength
	return r.Do(ctx, req)
}

func (r Client) DoRequestWithForm(ctx context.Context, method, reqUrl string, headers http.Header,
	data map[string][]string) (resp *http.Response, err error) {

	if headers == nil {
		headers = http.Header{}
	}
	headers.Add("Content-Type", "application/x-www-form-urlencoded")
	requestData := url.Values(data).Encode()

	if method == "GET" || method == "POST" || method == "DELETE"{
		if strings.ContainsRune(reqUrl,'?') {
			reqUrl += "&"
		} else {
			reqUrl += "?"
		}

		return r.DoRequest(ctx, method, reqUrl + requestData, headers)

	}

	return r.DoRequestWith(ctx, method, reqUrl, headers, strings.NewReader(requestData), len(requestData))
}

func (r Client) DoRequestWithJson(ctx context.Context, method, reqUrl string, headers http.Header,
	data interface{}) (resp *http.Response, err error) {

	reqBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	if headers == nil {
		headers = http.Header{}
	}

	headers.Add("Content-Type", "application/json")
	return r.DoRequestWith(ctx, method, reqUrl, headers, bytes.NewReader(reqBody), len(reqBody))
}

func (r Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", UserAgent)
	}

	transport := r.Transport

	if transport == nil {
		transport = http.DefaultTransport
	}

	select {
	case <- ctx.Done():
		err = ctx.Err()
		return
	default:
	}

	resp, err = r.Client.Do(req)
	return 
}