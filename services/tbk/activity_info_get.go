package tbk

//淘宝客-服务商-官方活动信息获取
//目前只支持饿了么，2020-04月底全部支持
import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)


type ActivityInfoGetRequest struct {
	*services.BaseRequest
}

func (r *ActivityInfoGetRequest) Method() string {
	return "taobao.tbk.activity.info.get"
}

func (r *ActivityInfoGetRequest) SetActivityMaterialId(activity_material_id string) *ActivityInfoGetRequest {
	r.SetValue("activity_material_id", activity_material_id)
	return  r
}

func (r *ActivityInfoGetRequest) SetAdzoneId(adzone_id string) *ActivityInfoGetRequest {
	r.SetValue("adzone_id", adzone_id)
	return  r
}

func (r *ActivityInfoGetRequest) SetSubPid(sub_pid string) *ActivityInfoGetRequest {
	r.SetValue("sub_pid", sub_pid)
	return  r
}

func (r *ActivityInfoGetRequest) SetRelationId(relation_id string) *ActivityInfoGetRequest {
	r.SetValue("relation_id", relation_id)
	return  r
}

func (c *TbkClient) ActivityInfoGet(request services.TaoBaoRequest) (response *ActivityInfoGetResponse, err error) {

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
	response = &ActivityInfoGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewActivityInfoGetRequest() *ActivityInfoGetRequest {
	return &ActivityInfoGetRequest{&services.BaseRequest{}}
}

type ActivityInfoGetResponse struct {
	TbkActivityInfoGetResponse TbkActivityInfoGetResponse `json:"tbk_activity_info_get_response"`
}

type TbkActivityInfoGetResponse struct {
	Data struct{
		WxQrcodeUrl string `json:"wx_qrcode_url"`
		ClickUrl string `json:"click_url"`
		ShortClickUrl string `json:"short_click_url"`
		TerminalType string `json:"terminal_type"`
		MaterialOssUrl string `json:"material_oss_url"`
		PageName string `json:"page_name"`
		PageStartTime string `json:"page_start_time"`
		PageEndTime string `json:"page_end_time"`
	} `json:"data"`
}