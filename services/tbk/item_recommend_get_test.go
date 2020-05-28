package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestItemRecommendGet(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewItemRecommendRequest()
	req.SetFields("num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url")
	req.SetNumIid("586114788080")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.ItemRecommendGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
