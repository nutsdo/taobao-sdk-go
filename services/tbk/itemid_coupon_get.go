package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ItemidCouponGetRequest struct {
	*services.BaseRequest
}

type ItemidCouponGetResponse struct {
	TbkItemidCouponGetResponse TbkItemidCouponGetResponse `json:"tbk_itemid_coupon_get_response"`
}

func (r *ItemidCouponGetRequest) Method() string {
	return "taobao.tbk.itemid.coupon.get"
}

func (r *ItemidCouponGetRequest) SetNumIids(num_iids string) {
	r.SetValue("num_iids", num_iids)
}

func (r *ItemidCouponGetRequest) SetPlatform(platform string) {
	r.SetValue("platform", platform)
}

func (r *ItemidCouponGetRequest) SetPid(pid string) {
	r.SetValue("pid", pid)
}

func (c *TbkClient) ItemidCouponGet(request services.TaoBaoRequest) (response *ItemidCouponGetResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	response = &ItemidCouponGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkItemidCouponGetResponse struct {
	Results struct {
		TbkCoupon []struct{
			SmallImages []struct{
				String string
			}
			CouponTotalCount int64 `json:"coupon_total_count"`
			CommissionRate string `json:"commission_rate"`
			UserType int64 `json:"user_type"`
			Title           string `json:"title"`
			CouponInfo string `json:"coupon_info"`
			ShopTitle string `json:"shop_title"`
			Category int64 `json:"category"`
			NumIid int64 `json:"num_iid"`
			CouponRemainCount int64 `json:"coupon_remain_count"`
			ZkFinalPrice string `json:"zk_final_price"`
			CouponStartTime string `json:"coupon_start_time"`
			Nick string `json:"nick"`
			SellerId int64 `json:"seller_id"`
			Volume int64 `json:"volume"`
			CouponEndTime string `json:"coupon_end_time"`
			CouponClickUrl string `json:"coupon_click_url"`
			PictUrl string `json:"pict_url"`
			ItemUrl string `json:"item_url"`
		}
	} `json:"results"`
}

func NewItemidCouponGetRequest() *ItemidCouponGetRequest {
	return &ItemidCouponGetRequest{&services.BaseRequest{}}
}
