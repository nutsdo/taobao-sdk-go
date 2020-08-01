package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

//

type OrderDetailsGetRequest struct {
	*services.BaseRequest
}

type OrderDetailsGetResponse struct {
	TbkOrderDetailsGetResponse TbkOrderDetailsGetResponse `json:"tbk_order_details_get_response""`
}

func (r *OrderDetailsGetRequest) Method() string {
	return "taobao.tbk.order.details.get"
}

//	查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询
func (r *OrderDetailsGetRequest) SetQueryType(query_type string) {
	r.SetValue("query_type", query_type)
}

//	位点，除第一页之外，都需要传递；前端原样返回。
func (r *OrderDetailsGetRequest) SetPositionIndex(position_index string) {
	r.SetValue("position_index", position_index)
}

//页大小，默认20，1~100
func (r *OrderDetailsGetRequest) SetPageSize(page_size string) {
	r.SetValue("page_size", page_size)
}

//推广者角色类型,2:二方，3:三方，不传，表示所有角色
func (r *OrderDetailsGetRequest) SetMemberType(member_type string) {
	r.SetValue("member_type", member_type)
}

//	淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
func (r *OrderDetailsGetRequest) SetTkStatus(tk_status string) {
	r.SetValue("tk_status", tk_status)
}

//订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，
// 但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，
// 调用时请务必注意时间段的选择，以保证亲能正常调用！
func (r *OrderDetailsGetRequest) SetEndTime(end_time string) {
	r.SetValue("end_time", end_time)
}

//订单查询开始时间
func (r *OrderDetailsGetRequest) SetStartTime(start_time string) {
	r.SetValue("start_time", start_time)
}

//跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
func (r *OrderDetailsGetRequest) SetJumpType(jump_type string) {
	r.SetValue("jump_type", jump_type)
}

//第几页，默认1，1~100
func (r *OrderDetailsGetRequest) SetPageNo(page_no string) {
	r.SetValue("page_no", page_no)
}

//场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
func (r *OrderDetailsGetRequest) SetOrderScene(order_scene string) {
	r.SetValue("order_scene", order_scene)
}

func (c *TbkClient) OrderDetailsGet(request services.TaoBaoRequest) (response *OrderDetailsGetResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	response = &OrderDetailsGetResponse{}

	err = services.Byte2Response(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewOrderDetailsGetRequest() *OrderDetailsGetRequest {
	return &OrderDetailsGetRequest{&services.BaseRequest{}}
}

type TbkOrderDetailsGetResponse struct {
	Data struct {
		Results struct {
			PublisherOrderDto []struct {
				TbPaidTime                         string `json:"tb_paid_time"`
				TkPaidTime                         string `json:"tk_paid_time"`
				PayPrice                           string `json:"pay_price"`
				PubShareFee                        string `json:"pub_share_fee"`
				TradeId                            string `json:"trade_id"`
				TkOrderRole                        int64  `json:"tk_order_role"`
				TkEarningTime                      string `json:"tk_earning_time"`
				AdzoneId                           int64  `json:"adzone_id"`
				PubShareRate                       string `json:"pub_share_rate"`
				Unid                               string `json:"unid"`
				RefundTag                          int64  `json:"refund_tag"`
				SubsidyRate                        string `json:"subsidy_rate"`
				TkTotalRate                        string `json:"tk_total_rate"`
				ItemCategoryName                   string `json:"item_category_name"`
				SellerNick                         string `json:"seller_nick"`
				PubId                              int64  `json:"pub_id"`
				AlimamaRate                        string `json:"alimama_rate"`
				SubsidyType                        string `json:"subsidy_type"`
				ItemImg                            string `json:"item_img"`
				PubSharePreFee                     string `json:"pub_share_pre_fee"`
				AlipayTotalPrice                   string `json:"alipay_total_price"`
				ItemTitle                          string `json:"item_title"`
				SiteName                           string `json:"site_name"`
				ItemNum                            int64  `json:"item_num"`
				SubsidyFee                         string `json:"subsidy_fee"`
				AlimamaShareFee                    string `json:"alimama_share_fee"`
				TradeParentId                      string `json:"trade_parent_id"`
				OrderType                          string `json:"order_type"`
				TkCreateTime                       string `json:"tk_create_time"`
				FlowSource                         string `json:"flow_source"`
				TerminalType                       string `json:"terminal_type"`
				ClickTime                          string `json:"click_time"`
				TkStatus                           int64  `json:"tk_status"`
				ItemPrice                          string `json:"item_price"`
				ItemId                             int64  `json:"item_id"`
				AdzoneName                         string `json:"adzone_name"`
				TotalCommissionRate                string `json:"total_commission_rate"`
				ItemLink                           string `json:"item_link"`
				SiteId                             int64  `json:"site_id"`
				SellerShopTitle                    string `json:"seller_shop_title"`
				IncomeRate                         string `json:"income_rate"`
				TotalCommissionFee                 string `json:"total_commission_fee"`
				TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
				TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
				TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
				SpecialId                          int64  `json:"special_id"`
				RelationId                         int64  `json:"relation_id"`
				DepositPrice                       string `json:"deposit_price"`
				TbDepositTime                      string `json:"tb_deposit_time"`
				TkDepositTime                      string `json:"tk_deposit_time"`
				AlscId                             string `json:"alsc_id"`
				AlscPid                            string `json:"alsc_pid"`
				ServiceFeeDtoList                  struct {
					ServiceFeeDto []struct {
						TkShareRoleType   int64  `json:"tk_share_role_type"`
						ShareRelativeRate string `json:"share_relative_rate"`
						ShareFee          string `json:"share_fee"`
						SharePreFee       string `json:"share_pre_fee"`
					}
				} `json:"service_fee_dto_list"`
			} `json:"publisher_order_dto"`
		} `json:"results"`
	} `json:"data"`
	HasPre        bool   `json:"has_pre"`
	PositionIndex string `json:"position_index"`
	HasNext       bool   `json:"has_next"`
	PageNo        int64  `json:"page_no"`
	PageSize      int64  `json:"page_size"`
	RequestId     string `json:"request_id"`
}
