package tbk

//淘宝客-推广者-官方活动转链
import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ActivityLinkGetRequest struct {
	*services.BaseRequest
}

func (r *ActivityLinkGetRequest) Method() string {
	return "taobao.tbk.activitylink.get"
}

func (r *ActivityLinkGetRequest) SetPlatform(platform string) *ActivityLinkGetRequest {
	r.SetValue("platform", platform)
	return r
}

func (r *ActivityLinkGetRequest) SetUnionId(union_id string) *ActivityLinkGetRequest {
	r.SetValue("union_id", union_id)
	return r
}

func (r *ActivityLinkGetRequest) SetAdzoneId(adzone_id string) *ActivityLinkGetRequest {
	r.SetValue("adzone_id", adzone_id)
	return r
}

func (r *ActivityLinkGetRequest) SetPromotionSceneId(promotion_scene_id string) *ActivityLinkGetRequest {
	r.SetValue("promotion_scene_id", promotion_scene_id)
	return r
}

func (r *ActivityLinkGetRequest) SetSubPid(sub_pid string) *ActivityLinkGetRequest {
	r.SetValue("sub_pid", sub_pid)
	return r
}

func (r *ActivityLinkGetRequest) SetRelationId(relation_id string) *ActivityLinkGetRequest {
	r.SetValue("relation_id", relation_id)
	return r
}

func (c *TbkClient) ActivityLinkGet(request services.TaoBaoRequest) (response *ActivityLinkGetResponse, err error) {

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
	response = &ActivityLinkGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewActivityLinkGetRequest() *ActivityLinkGetRequest {
	return &ActivityLinkGetRequest{&services.BaseRequest{}}
}

type ActivityLinkGetResponse struct {
	TbkActivitylinkGetResponse TbkActivitylinkGetResponse `json:"tbk_activitylink_get_response"`
}

type TbkActivitylinkGetResponse struct {
	ResultMsg    string `json:"result_msg"`
	Data         string `json:"data"`
	ResultCode   int64  `json:"result_code"`
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int64  `json:"biz_error_code"`
}
