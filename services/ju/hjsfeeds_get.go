package ju

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type HjsFeedsGetRequest struct {
	*services.BaseRequest
}

func (r *HjsFeedsGetRequest) Method() string {
	return "taobao.tbk.activity.info.get"
}

func (r *HjsFeedsGetRequest) SetActivityMaterialId(activity_material_id string) *ActivityInfoGetRequest {
	r.SetValue("activity_material_id", activity_material_id)
	return r
}

func (r *HjsFeedsGetRequest) SetAdzoneId(adzone_id string) *ActivityInfoGetRequest {
	r.SetValue("adzone_id", adzone_id)
	return r
}

func (r *HjsFeedsGetRequest) SetSubPid(sub_pid string) *ActivityInfoGetRequest {
	r.SetValue("sub_pid", sub_pid)
	return r
}

func (r *HjsFeedsGetRequest) SetRelationId(relation_id string) *ActivityInfoGetRequest {
	r.SetValue("relation_id", relation_id)
	return r
}

func (c *TbClient) ActivityInfoGet(request services.TaoBaoRequest) (response *ActivityInfoGetResponse, err error) {

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
	response = &HjsFeedsGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

func NewActivityInfoGetRequest() *HjsFeedsGetRequest {
	return &HjsFeedsGetRequest{&services.BaseRequest{}}
}

type JuHjsFeedsGetResponse struct {
	HjsFeedsGetResponse HjsFeedsGetResponse `json:"hjs_feeds_get_response"`
}

type HjsFeedsGetResponse struct {
	CurPage int64 `json:"cur_page"`
	TotalPage int64 `json:"total_page"`
	HasMore bool `json:"has_more"`
	PageSize int64 `json:"page_size"`
	Items []Items
}

type Items struct {
	RemindNum string
	OriginalPrice string
	ItemTitle string
	SoldCount string
	ActivityPrice string
	ItemPic string
	ItemStatus string
	BenefitInfo string
	CommAmount string
	BrandName string
	Index int64
	HjsItemLabel string
	SalePoint string
	ItemId string
	SellStatus int64
	JuId string
	SupplyList string
	ShopTitle string
	BenefitBO struct{
		BenefitName string
		StartFee int64
		Discount string
	}
}
