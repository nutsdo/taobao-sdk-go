package tbk

//淘宝客-推广者-官方活动转链
import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ScActivityLinkToolGetRequest struct {
	*services.BaseRequest
}

func (r *ScActivityLinkToolGetRequest) Method() string {
	return "taobao.tbk.sc.activitylink.toolget"
}


func (r *ScActivityLinkToolGetRequest) SetPlatform(platform string) *ScActivityLinkToolGetRequest {
	r.SetValue("platform", platform)
	return  r
}

func (r *ScActivityLinkToolGetRequest) SetUnionId(union_id string) *ScActivityLinkToolGetRequest {
	r.SetValue("union_id", union_id)
	return  r
}

func (r *ScActivityLinkToolGetRequest) SetAdzoneId(adzone_id string) *ScActivityLinkToolGetRequest {
	r.SetValue("adzone_id", adzone_id)
	return  r
}

func (r *ScActivityLinkToolGetRequest) SetPromotionSceneId(promotion_scene_id string) *ScActivityLinkToolGetRequest {
	r.SetValue("promotion_scene_id", promotion_scene_id)
	return  r
}

func (r *ScActivityLinkToolGetRequest) SetSiteId(site_id string) *ScActivityLinkToolGetRequest {
	r.SetValue("site_id", site_id)
	return  r
}

func (r *ScActivityLinkToolGetRequest) SetRelationId(relation_id string) *ScActivityLinkToolGetRequest {
	r.SetValue("relation_id", relation_id)
	return  r
}

func (c *TbkClient) ActivityLinkToolGet(request services.TaoBaoRequest) (response *ActivityLinkToolGetResponse, err error) {

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
	response = &ActivityLinkToolGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewActivityLinkToolGetRequest() *ScActivityLinkToolGetRequest {
	return &ScActivityLinkToolGetRequest{&services.BaseRequest{}}
}

type ActivityLinkToolGetResponse struct {
	TbkActivitylinkToolGetResponse TbkActivitylinkToolGetResponse `json:"tbk_activitylink_toolget_response"`
}

type TbkActivitylinkToolGetResponse struct {
	ResultMsg string `json:"result_msg"`
	Data string `json:"data"`
	ResultCode int64 `json:"result_code"`
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int64 `json:"biz_error_code"`
}
