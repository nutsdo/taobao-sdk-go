package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestActivityInfoGet(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)
	req := tbk.NewActivityInfoGetRequest()
	req.SetActivityMaterialId("1579491209717")
	req.SetAdzoneId("52280450373")
	req.SetRelationId("520045220")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.ActivityInfoGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
