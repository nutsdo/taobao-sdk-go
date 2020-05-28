package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestTpwdConvert(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewTpwdConvertRequest()
	req.SetPasswordContent("￥o2tl1gpKU9o￥")
	req.SetAdzoneId("65202750048")
	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.TpwdConvert(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
