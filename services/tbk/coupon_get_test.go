package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestCouponGet(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewCouponGetRequest()
	req.SetItemId("608466498831")
	req.SetActivityId("26eb53ad3bed4d33b9331dd2b4e55def")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.CouponGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
