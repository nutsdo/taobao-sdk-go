package tbk

import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ItemClickExtractRequest struct {
	*services.BaseRequest
}

type ItemClickExtractResponse struct {
	TbkItemClickExtractResponse TbkItemClickExtractResponse `json:"tbk_item_click_extract_response"`
}

func (r *ItemClickExtractRequest) Method() string {
	return "taobao.tbk.item.click.extract"
}

func (r *ItemClickExtractRequest) SetClickUrl(click_url string) {
	r.SetValue("click_url", click_url)
}

func (c *TbkClient) ItemClickExtract(request services.TaoBaoRequest) (response *ItemClickExtractResponse, err error) {

	resp, err := c.DoRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	response = &ItemClickExtractResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkItemClickExtractResponse struct {
	ItemId  string `json:"item_id"`
	OpenIid string `json:"open_iid"`
}

func NewItemClickExtractRequest() *ItemClickExtractRequest {
	return &ItemClickExtractRequest{&services.BaseRequest{}}
}
