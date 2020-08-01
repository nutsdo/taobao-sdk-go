package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

//

type DgMaterialOptionalRequest struct {
	*services.BaseRequest
}

type DgMaterialOptionalResponse struct {
	TbkDgMaterialOptionalResponse TbkDgMaterialOptionalResponse `json:"tbk_dg_material_optional_response"`
}

func (r *DgMaterialOptionalRequest) Method() string {
	return "taobao.tbk.dg.material.optional"
}

func (r *DgMaterialOptionalRequest) SetStartDsr(start_dsr string) {
	r.SetValue("start_dsr", start_dsr)
}

func (r *DgMaterialOptionalRequest) SetPageSize(page_size string) {
	r.SetValue("page_size", page_size)
}

func (r *DgMaterialOptionalRequest) SetPageNo(page_no string) {
	r.SetValue("page_no", page_no)
}

func (r *DgMaterialOptionalRequest) SetPlatform(platform string) {
	r.SetValue("platform", platform)
}

func (r *DgMaterialOptionalRequest) SetEndTkRate(end_tk_rate string) {
	r.SetValue("end_tk_rate", end_tk_rate)
}

func (r *DgMaterialOptionalRequest) SetStartTkRate(start_tk_rate string) {
	r.SetValue("start_tk_rate", start_tk_rate)
}

func (r *DgMaterialOptionalRequest) SetEndPrice(end_price string) {
	r.SetValue("end_price", end_price)
}

func (r *DgMaterialOptionalRequest) SetStartPrice(start_price string) {
	r.SetValue("start_price", start_price)
}

func (r *DgMaterialOptionalRequest) SetIsOverseas(is_overseas string) {
	r.SetValue("is_overseas", is_overseas)
}

func (r *DgMaterialOptionalRequest) SetIsTmall(is_tmall string) {
	r.SetValue("is_tmall", is_tmall)
}

func (r *DgMaterialOptionalRequest) SetSort(sort string) {
	r.SetValue("sort", sort)
}

func (r *DgMaterialOptionalRequest) SetItemloc(itemloc string) {
	r.SetValue("itemloc", itemloc)
}

func (r *DgMaterialOptionalRequest) SetCat(cat string) {
	r.SetValue("cat", cat)
}

func (r *DgMaterialOptionalRequest) SetQ(q string) {
	r.SetValue("q", q)
}

func (r *DgMaterialOptionalRequest) SetMaterialId(material_id string) {
	r.SetValue("material_id", material_id)
}

func (r *DgMaterialOptionalRequest) SetHasCoupon(has_coupon string) {
	r.SetValue("has_coupon", has_coupon)
}

func (r *DgMaterialOptionalRequest) SetIp(ip string) {
	r.SetValue("ip", ip)
}

func (r *DgMaterialOptionalRequest) SetAdzoneId(adzone_id string) {
	r.SetValue("adzone_id", adzone_id)
}

func (r *DgMaterialOptionalRequest) SetNeedFreeShipment(need_free_shipment string) {
	r.SetValue("need_free_shipment", need_free_shipment)
}

func (r *DgMaterialOptionalRequest) SetNeedPrepay(need_prepay string) {
	r.SetValue("need_prepay", need_prepay)
}

func (r *DgMaterialOptionalRequest) SetIncludePayRate30(include_pay_rate_30 string) {
	r.SetValue("include_pay_rate_30", include_pay_rate_30)
}

func (r *DgMaterialOptionalRequest) SetIncludeGoodRate(include_good_rate string) {
	r.SetValue("include_good_rate", include_good_rate)
}

func (r *DgMaterialOptionalRequest) SetIncludeRfdRate(include_rfd_rate string) {
	r.SetValue("include_rfd_rate", include_rfd_rate)
}

func (r *DgMaterialOptionalRequest) SetNpxLevel(npx_level string) {
	r.SetValue("npx_level", npx_level)
}

func (r *DgMaterialOptionalRequest) SetEndKaTkRate(end_ka_tk_rate string) {
	r.SetValue("end_ka_tk_rate", end_ka_tk_rate)
}

func (r *DgMaterialOptionalRequest) SetStartKaTkRate(start_ka_tk_rate string) {
	r.SetValue("start_ka_tk_rate", start_ka_tk_rate)
}

func (r *DgMaterialOptionalRequest) SetDeviceEncrypt(device_encrypt string) {
	r.SetValue("device_encrypt", device_encrypt)
}

func (r *DgMaterialOptionalRequest) SetDeviceValue(device_value string) {
	r.SetValue("device_value", device_value)
}

func (r *DgMaterialOptionalRequest) SetDeviceType(device_type string) {
	r.SetValue("device_type", device_type)
}

func (r *DgMaterialOptionalRequest) SetLockRateEndTime(lock_rate_end_time string) {
	r.SetValue("lock_rate_end_time", lock_rate_end_time)
}

func (r *DgMaterialOptionalRequest) SetLockRateStartTime(lock_rate_start_time string) {
	r.SetValue("lock_rate_start_time", lock_rate_start_time)
}

//2020-07-10更新
//本地化业务入参-LBS信息-经度
func (r *DgMaterialOptionalRequest) SetLongitude(longitude string) {
	r.SetValue("longitude", longitude)
}

//本地化业务入参-LBS信息-纬度
func (r *DgMaterialOptionalRequest) SetLatitude(latitude string) {
	r.SetValue("latitude", latitude)
}

//本地化业务入参-LBS信息-国际城市码，仅支持单个请求
func (r *DgMaterialOptionalRequest) SetCityCode(city_code string) {
	r.SetValue("city_code", city_code)
}

//商家ID，饿了么专用
func (r *DgMaterialOptionalRequest) SetSellerIds(seller_ids string) {
	r.SetValue("seller_ids", seller_ids)
}

//会员运营ID
func (r *DgMaterialOptionalRequest) SetSpecialId(special_id string) {
	r.SetValue("special_id", special_id)
}

//淘宝客外部用户标记，如自身系统账户ID：微信ID等
func (r *DgMaterialOptionalRequest) SetExternalId(external_id string) {
	r.SetValue("external_id", external_id)
}

func (c *TbkClient) DgMaterialOptional(request services.TaoBaoRequest) (response *DgMaterialOptionalResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	response = &DgMaterialOptionalResponse{}

	err = services.Byte2Response(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewDgMaterialOptionalRequest() *DgMaterialOptionalRequest {
	return &DgMaterialOptionalRequest{&services.BaseRequest{}}
}

//2020-07-10新增字段
type TbkDgMaterialOptionalResponse struct {
	TotalResults int64 `json:"total_results"`
	ResultList   struct {
		MapData []struct {
			CouponStartTime string `json:"coupon_start_time"`
			CouponEndTime   string `json:"coupon_end_time"`
			InfoDxjh        string `json:"info_dxjh"`
			TkTotalSales    string `json:"tk_total_sales"`
			TkTotalCommi    string `json:"tk_total_commi"`
			CouponId        string `json:"coupon_id"`
			NumIid          int64  `json:"num_iid"`
			Title           string `json:"title"`
			PictUrl         string `json:"pict_url"`
			SmallImages     struct {
				String []string `json:"string"`
			} `json:"small_images"`
			ReservePrice           string `json:"reserve_price"`
			ZkFinalPrice           string `json:"zk_final_price"`
			UserType               int64  `json:"user_type"`
			Provcity               string `json:"provcity"`
			ItemUrl                string `json:"item_url"`
			IncludeMkt             string `json:"include_mkt"`
			IncludeDxjh            string `json:"include_dxjh"`
			CommissionRate         string `json:"commission_rate"`
			Volume                 int64  `json:"volume"`
			SellerId               int64  `json:"seller_id"`
			CouponTotalCount       int64  `json:"coupon_total_count"`
			CouponRemainCount      int64  `json:"coupon_remain_count"`
			CouponInfo             string `json:"coupon_info"`
			CommissionType         string `json:"commission_type"`
			ShopTitle              string `json:"shop_title"`
			ShopDsr                int64  `json:"shop_dsr"`
			CouponShareUrl         string `json:"coupon_share_url"`
			Url                    string `json:"url"`
			LevelOneCategoryName   string `json:"level_one_category_name"`
			LevelOneCategoryId     int64  `json:"level_one_category_id"`
			CategoryName           string `json:"category_name"`
			CategoryId             int64  `json:"category_id"`
			ShortTitle             string `json:"short_title"`
			WhiteImage             string `json:"white_image"`
			Oetime                 string `json:"oetime"`
			Ostime                 string `json:"ostime"`
			JddNum                 int64  `json:"jdd_num"`
			JddPrice               string `json:"jdd_price"`
			UvSumPreSale           int64  `json:"uv_sum_pre_sale"`
			XId                    string `json:"x_id"`
			CouponStartFee         string `json:"coupon_start_fee"`
			CouponAmount           string `json:"coupon_amount"`
			ItemDescription        string `json:"item_description"`
			Nick                   string `json:"nick"`
			OrigPrice              string `json:"orig_price"`
			TotalStock             int64  `json:"total_stock"`
			SellNum                int64  `json:"sell_num"`
			Stock                  int64  `json:"stock"`
			TmallPlayActivityInfo  string `json:"tmall_play_activity_info"`
			ItemId                 int64  `json:"item_id"`
			RealPostFee            string `json:"real_post_fee"`
			LockRate               string `json:"lock_rate"`
			LockRateEndTime        int64  `json:"lock_rate_end_time"`
			LockRateStartTime      int64  `json:"lock_rate_start_time"`
			PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
			PresaleTailEndTime     int64  `json:"presale_tail_end_time"`
			PresaleTailStartTime   int64  `json:"presale_tail_start_time"`
			PresaleEndTime         int64  `json:"presale_end_time"`
			PresaleStartTime       int64  `json:"presale_start_time"`
			PresaleDeposit         string `json:"presale_deposit"`
			YsylTljSendTime        string `json:"ysyl_tlj_send_time"`
			YsylCommissionRate     string `json:"ysyl_commission_rate"`
			YsylTljFace            string `json:"ysyl_tlj_face"`
			YsylClickUrl           string `json:"ysyl_click_url"`
			YsylTljUseEndTime      string `json:"ysyl_tlj_use_end_time"`
			YsylTljUseStartTime    string `json:"ysyl_tlj_use_start_time"`
			SaleBeginTime          string `json:"sale_begin_time"`        //2020-07-10新增
			SaleEndTime            string `json:"sale_end_time"`          //2020-07-10新增
			Distance               string `json:"distance"`               //2020-07-10新增
			UsableShopId           string `json:"usable_shop_id"`         //2020-07-10新增
			SalePrice              string `json:"sale_price"`             //2020-07-10新增
			KuadianPromotionInfo   string `json:"kuadian_promotion_info"` //2020-07-10新增
		} `json:"map_data"`
	} `json:"result_list"`
	RequestId string `json:"request_id"`
}
