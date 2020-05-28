package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type TpwdConvertRequest struct {
	*services.BaseRequest
}

type TpwdConvertResponse struct {
	TbkTpwdConvertResponse TbkTpwdConvertResponse `json:"tbk_tpwd_convert_response"`
}

func (r *TpwdConvertRequest) Method() string {
	return "taobao.tbk.tpwd.convert"
}

func (r *TpwdConvertRequest) SetPasswordContent(password_content string) {
	r.SetValue("password_content", password_content)
}

func (r *TpwdConvertRequest) SetAdzoneId(adzone_id string) {
	r.SetValue("adzone_id", adzone_id)
}

func (r *TpwdCreateRequest) SetDx(dx string) {
	r.SetValue("dx", dx)
}

func (c *TbkClient) TpwdConvert(request services.TaoBaoRequest) (response *TpwdConvertResponse, err error) {

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
	response = &TpwdConvertResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkTpwdConvertResponse struct {
	Data struct {
		NumIid   string `json:"num_iid"`
		ClickUrl string `json:"click_url"`
	} `json:"data"`
}

func NewTpwdConvertRequest() *TpwdConvertRequest {
	return &TpwdConvertRequest{&services.BaseRequest{}}
}
