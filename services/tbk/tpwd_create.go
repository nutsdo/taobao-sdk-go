package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type TpwdCreateRequest struct {
	*services.BaseRequest
}

type TpwdCreateResponse struct {
	TbkTpwdCreateResponse TbkTpwdCreateResponse `json:"tbk_tpwd_create_response"`
}

func (r *TpwdCreateRequest) Method() string {
	return "taobao.tbk.tpwd.create"
}

func (r *TpwdCreateRequest) SetUserId(user_id string) {
	r.SetValue("user_id", user_id)
}

func (r *TpwdCreateRequest) SetText(text string) {
	r.SetValue("text", text)
}

func (r *TpwdCreateRequest) SetUrl(url string) {
	r.SetValue("url", url)
}

func (r *TpwdCreateRequest) SetLogo(logo string) {
	r.SetValue("logo", logo)
}

func (r *TpwdCreateRequest) SetExt(ext string) {
	r.SetValue("ext", ext)
}

func (c *TbkClient) TpwdCreate(request services.TaoBaoRequest) (response *TpwdCreateResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	response = &TpwdCreateResponse{}

	if err := services.Byte2Response(data,response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkTpwdCreateResponse struct {
	Data struct{
		Model string `json:"model"`
	} `json:"data"`
}

func NewTpwdCreateRequest() *TpwdCreateRequest {
	return &TpwdCreateRequest{&services.BaseRequest{}}
}
