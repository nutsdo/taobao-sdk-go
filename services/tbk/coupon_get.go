package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type CouponGetRequest struct {
	*services.BaseRequest
}

type CouponGetResponse struct {
	TbkCouponGetResponse TbkCouponGetResponse `json:"tbk_coupon_get_response"`
}

func (r *CouponGetRequest) Method() string {
	return "taobao.tbk.coupon.get"
}

func (r *CouponGetRequest) SetMe(me string) {
	r.SetValue("me", me)
}

func (r *CouponGetRequest) SetItemId(item_id string) {
	r.SetValue("item_id", item_id)
}

func (r *CouponGetRequest) SetActivityId(activity_id string) {
	r.SetValue("activity_id", activity_id)
}

func (c *TbkClient) CouponGet(request services.TaoBaoRequest) (response *CouponGetResponse, err error) {

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
	response = &CouponGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkCouponGetResponse struct {
	Data struct {
		CouponStartFee string `json:"coupon_start_fee"`
		CouponSemainCount int64 `json:"coupon_semain_count"`
		CouponTotalCount int64 `json:"coupon_total_count"`
		CouponEndTime string `json:"coupon_end_time"`
		CouponStartTime string `json:"coupon_start_time"`
		CouponAmount string `json:"coupon_amount"`
		CouponSrcScene int64 `json:"coupon_src_scene"`
		CouponType int64 `json:"coupon_type"`
		CouponActivityId string `json:"coupon_activity_id"`
	} `json:"data"`
}

func NewCouponGetRequest() *CouponGetRequest {
	return &CouponGetRequest{&services.BaseRequest{}}
}

