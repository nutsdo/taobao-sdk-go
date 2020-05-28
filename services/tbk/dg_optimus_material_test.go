package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestDgOptimusMaterial(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewDgOptimusMaterialRequest()
	req.SetAdzoneId("52280450373")
	req.SetPageSize("5")
	req.SetMaterialId("13256")
	req.SetItemId("612887993968")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.DgOptimusMaterial(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}

	fmt.Printf("正常响应:%#v", resp)

}
