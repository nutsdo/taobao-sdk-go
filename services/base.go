package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type BaseRequest struct {
	baseRequest *http.Request
	values      url.Values
}

type TaoBaoRequest interface {
	Method() string
	GetValues() url.Values
}

type BaseResponse struct {
	httpStatus         int
	httpHeaders        map[string][]string
	httpContentString  string
	httpContentBytes   []byte
	originHttpResponse *http.Response
}

type TaoBaoResponse interface {
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpHeaders() map[string][]string
	GetHttpContentString() string
	GetHttpContentBytes() []byte
	GetOriginHttpResponse() *http.Response
	parseFromHttpResponse(httpResponse *http.Response) error
}

//func (r *BaseRequest) GetRequestHost() string {
//	return r.baseRequest.Host
//}

func (r *BaseRequest) SetValue(key, value string) {

	if r.values == nil {
		r.values = url.Values{}
	}
	r.values.Set(key, value)
}

func (r *BaseRequest) GetValues() url.Values {

	return r.values
}

func (br *BaseResponse) GetHttpStatus() int {
	return br.httpStatus
}

func (br *BaseResponse) GetHttpHeaders() map[string][]string {
	return br.httpHeaders
}

func (br *BaseResponse) GetHttpContentString() string {
	return br.httpContentString
}

func (br *BaseResponse) GetHttpContentBytes() []byte {
	return br.httpContentBytes
}

func (br *BaseResponse) GetOriginHttpResponse() *http.Response {
	return br.originHttpResponse
}

func (br *BaseResponse) IsSuccess() bool {

	if br.GetHttpStatus() >= 200 && br.GetHttpStatus() < 300 {
		return true
	}

	return false
}

func Byte2Response(data []byte, resp interface{}) error {

	fmt.Println(string(data))
	if strings.Contains(string(data),"error_response") {

		return errors.New(string(data))
	}
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}

	return nil
}

// Unmarshal object from http response body to target Response
func Unmarshal(response TaoBaoResponse, httpResponse *http.Response, format string) (err error) {
	err = response.parseFromHttpResponse(httpResponse)
	if err != nil {
		return
	}
	if !response.IsSuccess() {
		err = errors.New("123123")
		return
	}

	if len(response.GetHttpContentBytes()) == 0 {
		return
	}

	if strings.ToUpper(format) == "JSON" {
		err = json.Unmarshal(response.GetHttpContentBytes(), response)
		if err != nil {

		}
	} else if strings.ToUpper(format) == "XML" {
		err = xml.Unmarshal(response.GetHttpContentBytes(), response)
	}
	return
}


func (br *BaseResponse) parseFromHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	br.httpStatus = httpResponse.StatusCode
	br.httpHeaders = httpResponse.Header
	br.httpContentBytes = body
	br.httpContentString = string(body)
	br.originHttpResponse = httpResponse
	return
}

func (br *BaseResponse) String() string {
	resultBuilder := bytes.Buffer{}
	// statusCode
	// resultBuilder.WriteString("\n")
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", br.originHttpResponse.Proto, br.originHttpResponse.Status))
	// httpHeaders
	//resultBuilder.WriteString("Headers:\n")
	for key, value := range br.httpHeaders {
		resultBuilder.WriteString(key + ": " + strings.Join(value, ";") + "\n")
	}
	resultBuilder.WriteString("\n")
	// content
	//resultBuilder.WriteString("Content:\n")
	resultBuilder.WriteString(br.httpContentString + "\n")
	return resultBuilder.String()
}

type BaseErrorResponse struct {
	ErrorResponse *ErrorResponse `json:"error_response"`
}

type BaseError struct {
	BaseCode int64
	BaseMsg string
}

type ErrorResponse struct{
	SubMsg  string `json:"sub_msg"`
	Code    int64  `json:"code"`
	SubCode string `json:"sub_code"`
	Msg     string `json:"msg"`
}

func MarshalErrorResponse(data []byte) *BaseErrorResponse {
	//将结果转换成map，判断是否错误响应
	respMap := make(map[string]interface{})

	errresp := &BaseErrorResponse{}

	if err := json.Unmarshal(data, &respMap); err == nil {

		fmt.Println("转换map后:", respMap)
		if _,ok := respMap["error_response"]; ok {

			err :=json.Unmarshal(data, errresp)
			if err != nil {
				fmt.Println("解析错误:", err)
			}
			fmt.Println("错误的响应:", errresp)
			return errresp
		}
	}
	return errresp
}