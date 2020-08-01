package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ItemInfoRequest struct {
	*services.BaseRequest
}

type ItemInfoResponse struct {
	TbkItemInfoGetResponse TbkItemInfoGetResponse `json:"tbk_item_info_get_response"`
}

func (r *ItemInfoRequest) Method() string {
	return "taobao.tbk.item.info.get"
}

func (r *ItemInfoRequest) SetNumIids(num_iids string) {
	r.SetValue("num_iids", num_iids)
}

func (r *ItemInfoRequest) SetPlatform(platform string) {
	r.SetValue("platform", platform)
}

func (r *ItemInfoRequest) SetIp(ip string) {
	r.SetValue("ip", ip)
}

func (c *TbkClient) GetItemInfo(request services.TaoBaoRequest) (response *ItemInfoResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	response = &ItemInfoResponse{}

	err = services.Byte2Response(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewItemInfoRequest() *ItemInfoRequest {
	return &ItemInfoRequest{&services.BaseRequest{}}
}

type TbkItemInfoGetResponse struct {
	Results struct {
		NTbkItem []struct {
			CatName     string `json:"cat_name"`
			NumIid      int64  `json:"num_iid"`
			Title       string `json:"title"`
			PictUrl     string `json:"pict_url"`
			SmallImages struct {
				string []string
			} `json:"small_images"`
			ReservePrice               string
			ZkFinalPrice               string `json:"zk_final_price"`
			UserType                   int64  `json:"user_type"`
			Provcity                   string `json:"provcity"`
			ItemURL                    string `json:"item_url"`
			SellerID                   int64  `json:"seller_id"`
			Volume                     int64  `json:"volume"`
			Nick                       string `json:"nick"`
			CatLeafName                string `json:"cat_leaf_name"`
			IsPrepay                   bool   `json:"is_prepay"`
			ShopDsr                    int64  `json:"shop_dsr"`
			Ratesum                    int64  `json:"ratesum"`
			IRfdRate                   bool   `json:"i_rfd_rate"`
			HGoodRate                  bool   `json:"h_good_rate"`
			HPayRate30                 bool   `json:"h_pay_rate30"`
			FreeShipment               bool   `json:"free_shipment"`
			MaterialLibType            string `json:"material_lib_type"`
			PresaleDiscountFeeText     string `json:"presale_discount_fee_text"`
			PresaleTailEndTime         int64  `json:"presale_tail_end_time"`
			PresaleTailStartTime       int64  `json:"presale_tail_start_time"`
			PresaleEndTime             int64  `json:"presale_end_time"`
			PresaleStartTime           int64  `json:"presale_start_time"`
			PresaleDeposit             string `json:"presale_deposit"`
			JuPlayEndTime              int64  `json:"ju_play_end_time"`
			JuPlayStartTime            int64  `json:"ju_play_start_time"`
			PlayInfo                   string `json:"play_info"`
			TmallPlayActivityEndTime   int64  `json:"tmall_play_activity_end_time"`
			TmallPlayActivityStartTime int64  `json:"tmall_play_activity_start_time"`
		} `json:"n_tbk_item"`
	} `json:"results"`
	RequestId string `json:"request_id"`
}
