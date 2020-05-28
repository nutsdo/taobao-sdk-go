package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

//

type DgOptimusMaterialRequest struct {
	*services.BaseRequest
}

type DgOptimusMaterialResponse struct {
	TbkDgOptimusMaterialResponse TbkDgOptimusMaterialResponse `json:"tbk_dg_optimus_material_response""`
}

func (r *DgOptimusMaterialRequest) Method() string {
	return "taobao.tbk.dg.optimus.material"
}

func (r *DgOptimusMaterialRequest) SetPageSize(page_size string) {
	r.SetValue("page_size", page_size)
}


func (r *DgOptimusMaterialRequest) SetAdzoneId(adzone_id string) {
	r.SetValue("adzone_id", adzone_id)
}

func (r *DgOptimusMaterialRequest) SetPageNo(page_no string) {
	r.SetValue("page_no", page_no)
}

func (r *DgOptimusMaterialRequest) SetMaterialId(material_id string) {
	r.SetValue("material_id", material_id)
}

func (r *DgOptimusMaterialRequest) SetDeviceValue(device_value string) {
	r.SetValue("device_value", device_value)
}

func (r *DgOptimusMaterialRequest) SetDeviceEncrypt(device_encrypt string) {
	r.SetValue("device_encrypt", device_encrypt)
}

func (r *DgOptimusMaterialRequest) SetDeviceType(device_type string) {
	r.SetValue("device_type", device_type)
}

//内容专用-内容详情ID
func (r *DgOptimusMaterialRequest) SetContentId(content_id string) {
	r.SetValue("content_id", content_id)
}

//内容专用-内容渠道信息
func (r *DgOptimusMaterialRequest) SetContentSource(content_source string) {
	r.SetValue("content_source", content_source)
}

//商品ID，用于相似商品推荐
func (r *DgOptimusMaterialRequest) SetItemId(item_id string) {
	r.SetValue("item_id", item_id)
}

//选品库投放id
func (r *DgOptimusMaterialRequest) SetFavoritesId(favorites_id string) {
	r.SetValue("favorites_id", favorites_id)
}

func (c *TbkClient) DgOptimusMaterial(request services.TaoBaoRequest) (response *DgOptimusMaterialResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	response = &DgOptimusMaterialResponse{}

	err = services.Byte2Response(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewDgOptimusMaterialRequest() *DgOptimusMaterialRequest {
	return &DgOptimusMaterialRequest{&services.BaseRequest{}}
}

type TbkDgOptimusMaterialResponse struct {
	TotalCount int64 `json:"total_count"`
	IsDefault string `json:"is_default"`
	ResultList   struct {
		MapData []struct {
			CouponAmount           int64 `json:"coupon_amount"`
			SmallImages     struct {
				String []string `json:"string"`
			} `json:"small_images"`
			ShopTitle              string `json:"shop_title"`
			CategoryId             int64  `json:"category_id"`
			CouponStartFee         string `json:"coupon_start_fee"`
			ItemId                 int64  `json:"item_id"`
			CouponTotalCount       int64  `json:"coupon_total_count"`
			UserType               int64  `json:"user_type"`
			ZkFinalPrice           string `json:"zk_final_price"`
			CouponRemainCount      int64  `json:"coupon_remain_count"`
			CommissionRate         string `json:"commission_rate"`
			CouponStartTime string `json:"coupon_start_time"`
			Title           string `json:"title"`
			ItemDescription        string `json:"item_description"`
			SellerId               int64  `json:"seller_id"`
			Volume                 int64  `json:"volume"`
			CouponEndTime   string `json:"coupon_end_time"`
			CouponClickUrl string `json:"coupon_click_url"`
			PictUrl         string `json:"pict_url"`
			ClickUrl           string `json:"click_url"`
			Stock                  int64  `json:"stock"`
			SellNum                int64  `json:"sell_num"`
			TotalStock             int64  `json:"total_stock"`
			Oetime                 string `json:"oetime"`
			Ostime                 string `json:"ostime"`
			JddNum                 int64  `json:"jdd_num"`
			JddPrice               string `json:"jdd_price"`
			OrigPrice              string `json:"orig_price"`
			LevelOneCategoryName   string `json:"level_one_category_name"`
			LevelOneCategoryId     int64  `json:"level_one_category_id"`
			CategoryName           string `json:"category_name"`
			WhiteImage             string `json:"white_image"`
			ShortTitle             string `json:"short_title"`
			WordList				struct{
				WordMapData []struct{
					Url string `json:"url"`
					Word string `json:"word"`
				} `json:"word_map_data"`
			} `json:"word_list"`
			TmallPlayActivityInfo  string `json:"tmall_play_activity_info"`
			UvSumPreSale           int64  `json:"uv_sum_pre_sale"`
			XId                    string `json:"x_id"`
			NewUserPrice           string `json:"new_user_price"`
			CouponInfo             string `json:"coupon_info"`
			CouponShareUrl         string `json:"coupon_share_url"`
			Nick                   string `json:"nick"`
			ReservePrice           string `json:"reserve_price"`
			JuOnlineEndTime string `json:"ju_online_end_time"`
			JuOnlineStartTime string `json:"ju_online_start_time"`
			MaochaoPlayEndTime string `json:"maochao_play_end_time"`
			MaochaoPlayStartTime string `json:"maochao_play_start_time"`
			MaochaoPlayConditions string `json:"maochao_play_conditions"`
			MaochaoPlayDiscounts string `json:"maochao_play_discounts"`
			MaochaoPlayDiscountType string `json:"maochao_play_discount_type"`
			MaochaoPlayFreePostFee string `json:"maochao_play_free_post_fee"`
			MultiCouponZkRate string `json:"multi_coupon_zk_rate"`
			PriceAfterMultiCoupon string `json:"price_after_multi_coupon"`
			MultiCouponItemCount string `json:"multi_coupon_item_count"`
			LockRate               string `json:"lock_rate"`
			LockRateEndTime        int64  `json:"lock_rate_end_time"`
			LockRateStartTime      int64  `json:"lock_rate_start_time"`
			PromotionType      string `json:"promotion_type"`
			PromotionInfo      string `json:"promotion_info"`
			PromotionDiscount  string `json:"promotion_discount"`
			PromotionCondition string `json:"promotion_condition"`
			PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
			PresaleTailEndTime     int64  `json:"presale_tail_end_time"`
			PresaleTailStartTime   int64  `json:"presale_tail_start_time"`
			PresaleEndTime         int64  `json:"presale_end_time"`
			PresaleStartTime       int64  `json:"presale_start_time"`
			PresaleDeposit         string `json:"presale_deposit"`
			RealPostFee            string `json:"real_post_fee"`
			YsylTljSendTime        string `json:"ysyl_tlj_send_time"`
			YsylCommissionRate     string `json:"ysyl_commission_rate"`
			YsylTljFace            string `json:"ysyl_tlj_face"`
			YsylClickUrl           string `json:"ysyl_click_url"`
			YsylTljUseEndTime      string `json:"ysyl_tlj_use_end_time"`
			YsylTljUseStartTime    string `json:"ysyl_tlj_use_start_time"`
			JuPlayEndTime int64 `json:"ju_play_end_time"`
			JuPlayStartTime int64 `json:"ju_play_start_time"`
			PlayInfo string `json:"play_info"`
			TmallPlayActivityEndTime int64 `json:"tmall_play_activity_end_time"`
			TmallPlayActivityStartTime int64 `json:"tmall_play_activity_start_time"`
			JuPreShowEndTime string `json:"ju_pre_show_end_time"`
			JuPreShowStartTime string `json:"ju_pre_show_start_time"`
			FavoritesInfo struct{
				TotalCount int64 `json:"total_count"`
				FavoritesList struct{
					FavoritesDetail []struct{
						FavoritesId int64 `json:"favorites_id"`
						FavoritesTitle string `json:"favorites_title"`
					} `json:"favorites_detail"`
				} `json:"favorites_list"`
			} `json:"favorites_info"`

		} `json:"map_data"`
	} `json:"result_list"`
	RequestId string `json:"request_id"`
}
