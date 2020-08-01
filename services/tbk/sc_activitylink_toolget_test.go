package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestActivityLinkToolGet(t *testing.T) {

	c := tbk.NewClientWith(credential)
	c.SetSession("62004008ZZf5b07ee0629a07373342cd533036fd2f5b44c1866605434")

	req := tbk.NewActivityLinkToolGetRequest()
	req.SetPromotionSceneId("1579491209717")
	req.SetSiteId("183500025")
	req.SetAdzoneId("52280450373")
	req.SetRelationId("520045220")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.ActivityLinkToolGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
