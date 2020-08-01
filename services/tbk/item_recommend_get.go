package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ItemRecommendRequest struct {
	*services.BaseRequest
}

func (r *ItemRecommendRequest) Method() string {
	return "taobao.tbk.item.recommend.get"
}

func (r *ItemRecommendRequest) SetFields(fields string) *ItemRecommendRequest {
	r.SetValue("fields", fields)
	return r
}

func (r *ItemRecommendRequest) SetNumIid(num_iid string) *ItemRecommendRequest {
	r.SetValue("num_iid", num_iid)
	return r
}

func (r *ItemRecommendRequest) SetCount(count string) *ItemRecommendRequest {
	r.SetValue("count", count)
	return r
}

func (r *ItemRecommendRequest) SetPlatform(platform string) *ItemRecommendRequest {
	r.SetValue("platform", platform)
	return r
}

func (c *TbkClient) ItemRecommendGet(request services.TaoBaoRequest) (response *ItemRecommendGetResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	response = NewItemRecommendGetResponse()

	err = services.Byte2Response(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ItemRecommendGetResponse struct {
	TbkItemRecommendGetResponse TbkItemRecommendGetResponse `json:"tbk_item_recommend_get_response"`
}

type TbkItemRecommendGetResponse struct {
	Results TbkItemRecommendGetResults `json:"results"`
}

type TbkItemRecommendGetResults struct {
	NTbkItem []TbkItemRecommendGetNTbkItem `json:"n_tbk_item"`
}

type TbkItemRecommendGetNTbkItem struct {
	NumIid       int64    `json:"num_iid"`
	Title        string   `json:"title"`
	PictUrl      string   `json:"pict_url"`
	SmallImages  []string `json:"small_images"`
	ReservePrice string   `json:"reserve_price"`
	ZkFinalPrice string   `json:"zk_final_price"`
	UserType     int64    `json:"user_type"`
	Provcity     string   `json:"provcity"`
	ItemUrl      string   `json:"item_url"`
	Nick         string   `json:"nick"`
	SellerId     int64    `json:"seller_id"`
	Volume       int64    `json:"volume"`
}

func NewItemRecommendRequest() *ItemRecommendRequest {
	return &ItemRecommendRequest{&services.BaseRequest{}}
}

func NewItemRecommendGetResponse() *ItemRecommendGetResponse {
	return &ItemRecommendGetResponse{}
}
