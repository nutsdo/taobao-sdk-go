package tbk

// 淘宝客-推广者-图文内容输出
import (
	"fmt"
	"io/ioutil"

	"github.com/nutsdo/taobao-sdk-go/services"
)

type ContentGetRequest struct {
	*services.BaseRequest
}

type ContentGetResponse struct {
	TbkContentGetResponse TbkContentGetResponse `json:"tbk_content_get_response"`
}

func (r *ContentGetRequest) Method() string {
	return "taobao.tbk.content.get"
}

func (r *ContentGetRequest) SetAdzoneId(adzone_id string) {
	r.SetValue("adzone_id", adzone_id)
}

func (r *ContentGetRequest) SetType(tp string) {
	r.SetValue("type", tp)
}

func (r *ContentGetRequest) SetBeforeTimestamp(before_timestamp string) {
	r.SetValue("before_timestamp", before_timestamp)
}

func (r *ContentGetRequest) SetCount(count string) {
	r.SetValue("count", count)
}

func (r *ContentGetRequest) SetCid(cid string) {
	r.SetValue("cid", cid)
}

func (r *ContentGetRequest) SetImageWidth(image_width string) {
	r.SetValue("image_width", image_width)
}

func (r *ContentGetRequest) SetImageHeight(image_height string) {
	r.SetValue("image_height", image_height)
}

//默认可不传,内容库类型:1 优质内容
func (r *ContentGetRequest) SetContentSet(content_set string) {
	r.SetValue("content_set", content_set)
}

func (c *TbkClient) ContentGet(request services.TaoBaoRequest) (response *ContentGetResponse, err error) {

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
	response = &ContentGetResponse{}

	if err := services.Byte2Response(data, response); err != nil {
		return nil, err
	}

	return response, err
}

type TbkContentGetResponse struct {
	Result struct {
		Data struct{
			Count int64 `json:"count"`
			Contents struct{
				MapData []struct{
					Summary string `json:"summary"`
					Score int64 `json:"score"`
					Tags string `json:"tags"`
					Title string `json:"title"`
					AuthorId string `json:"author_id"`
					AuthorNick string `json:"author_nick"`
					AuthorAvatar string `json:"author_avatar"`
					Clink string `json:"clink"`
					Type string `json:"type"`
					UiStyle string `json:"ui_style"`
					Images struct{
						String []string `json:"string"`
					} `json:"images"`
					ContentCategories string `json:"content_categories"`
					PublishTime string `json:"publish_time"`
					ContentId int64 `json:"content_id"`
					PromotionTag string `json:"promotion_tag"`
					ItemIds []int64 `json:"item_ids"`
					UpdateTime string `json:"update_time"`
				} `json:"map_data"`
			} `json:"content"`
			LastTimestamp string `json:"last_timestamp"`
		}
	} `json:"result"`
}

func NewContentGetRequest() *ContentGetRequest {
	return &ContentGetRequest{&services.BaseRequest{}}
}

