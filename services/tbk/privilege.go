package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type PrivilegeGetRequest struct {
	*services.BaseRequest
}

type PrivilegeGetResponse struct {
	TbkPrivilegeGetResponse TbkPrivilegeGetResponse `json:"tbk_privilege_get_response"`
}

func (r *PrivilegeGetRequest) Method() string {
	return "taobao.tbk.privilege.get"
}

func (r *PrivilegeGetRequest) SetItemId(item_id string) {
	r.SetValue("item_id", item_id)
}

func (r *PrivilegeGetRequest) SetAdzoneId(adzone_id string) {
	r.SetValue("adzone_id", adzone_id)
}

func (r *PrivilegeGetRequest) SetPlatform(platform string) {
	r.SetValue("platform", platform)
}

func (r *PrivilegeGetRequest) SetSiteId(site_id string) {
	r.SetValue("site_id", site_id)
}

func (r *PrivilegeGetRequest) SetMe(me string) {
	r.SetValue("me", me)
}

func (r *PrivilegeGetRequest) SetRelationId(relation_id string) {
	r.SetValue("relation_id", relation_id)
}

//2020-07-10 新增入参
//会员运营ID
func (r *PrivilegeGetRequest) SetSpecialId(special_id string) {
	r.SetValue("special_id", special_id)
}

//淘宝客外部用户标记，如自身系统账户ID：微信ID等
func (r *PrivilegeGetRequest) SetExternalId(external_id string) {
	r.SetValue("external_id", external_id)
}

func (c *TbkClient) PrivilegeGet(request services.TaoBaoRequest) (response *PrivilegeGetResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	response = &PrivilegeGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewPrivilegeGetRequest() *PrivilegeGetRequest {
	return &PrivilegeGetRequest{&services.BaseRequest{}}
}

type TbkPrivilegeGetResponse struct {
	Result struct {
		Data struct {
			CategoryId          int64  `json:"category_id"`
			CouponClickUrl      string `json:"coupon_click_url"`
			CouponEndTime       string `json:"coupon_end_time"`
			CouponInfo          string `json:"coupon_info"`
			CouponStartTime     string `json:"coupon_start_time"`
			ItemId              int64  `json:"item_id"`
			MaxCommissionRate   string `json:"max_commission_rate"`
			CouponTotalCount    int64  `json:"coupon_total_count"`
			CouponRemainCount   int64  `json:"coupon_remain_count"`
			MmCouponRemainCount int64  `json:"mm_coupon_remain_count"`
			MmCouponTotalCount  int64  `json:"mm_coupon_total_count"`
			MmCouponClickUrl    string `json:"mm_coupon_click_url"`
			MmCouponEndTime     string `json:"mm_coupon_end_time"`
			MmCouponStartTime   string `json:"mm_coupon_start_time"`
			MmCouponInfo        string `json:"mm_coupon_info"`
			CouponType          int64  `json:"coupon_type"`
			ItemUrl             string `json:"item_url"`
			YsylClickUrl        string `json:"ysyl_click_url"`
			YsylTljFace         string `json:"ysyl_tlj_face"`
			YsylTljSendTime     string `json:"ysyl_tlj_send_time"`
			YsylCommissionRate  string `json:"ysyl_commission_rate"`
			YsylTljUseStartTime string `json:"ysyl_tlj_use_start_time"`
			YsylTljUseEndTime   string `json:"ysyl_tlj_use_end_time"`
			MinCommissionRate   string `json:"min_commission_rate"`
		} `json:"data"`
	} `json:"result"`
	RequestId string `json:"request_id"`
}
