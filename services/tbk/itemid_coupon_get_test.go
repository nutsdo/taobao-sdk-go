package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestItemidCouponGet(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewItemidCouponGetRequest()
	req.SetNumIids("611032679314")
	req.SetPid("mm_52121231_46736523_1516172508")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.TpwdConvert(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}