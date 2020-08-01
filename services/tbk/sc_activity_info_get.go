package tbk

//淘宝客-服务商-官方活动信息获取

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ScActivityInfoGetRequest struct {
	*services.BaseRequest
}

func (r *ScActivityInfoGetRequest) Method() string {
	return "taobao.tbk.sc.activity.info.get"
}

func (r *ScActivityInfoGetRequest) SetActivityMaterialId(activity_material_id string) *ScActivityInfoGetRequest {
	r.SetValue("activity_material_id", activity_material_id)
	return r
}

func (r *ScActivityInfoGetRequest) SetAdzoneId(adzone_id string) *ScActivityInfoGetRequest {
	r.SetValue("adzone_id", adzone_id)
	return r
}

func (r *ScActivityInfoGetRequest) SetSiteId(site_id string) *ScActivityInfoGetRequest {
	r.SetValue("site_id", site_id)
	return r
}

func (r *ScActivityInfoGetRequest) SetRelationId(relation_id string) *ScActivityInfoGetRequest {
	r.SetValue("relation_id", relation_id)
	return r
}

func (c *TbkClient) ScActivityInfoGet(request services.TaoBaoRequest) (response *ScActivityInfoGetResponse, err error) {

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
	response = &ScActivityInfoGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewScActivityInfoGetRequest() *ScActivityInfoGetRequest {
	return &ScActivityInfoGetRequest{&services.BaseRequest{}}
}

type ScActivityInfoGetResponse struct {
	TbkScActivityInfoGetResponse TbkScActivityInfoGetResponse `json:"tbk_sc_activity_info_get_response"`
}

type TbkScActivityInfoGetResponse struct {
	Data struct {
		WxQrcodeUrl    string `json:"wx_qrcode_url"`
		ClickUrl       string `json:"click_url"`
		ShortClickUrl  string `json:"short_click_url"`
		TerminalType   string `json:"terminal_type"`
		MaterialOssUrl string `json:"material_oss_url"`
		PageName       string `json:"page_name"`
		PageStartTime  string `json:"page_start_time"`
		PageEndTime    string `json:"page_end_time"`
	} `json:"data"`
}
