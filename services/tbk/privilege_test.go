package tbk_test

import (
"fmt"
"testing"

"github.com/nutsdo/taobao-sdk-go/auth"
"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestPrivilege(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)
	c.SetSession("62004008ZZf5b07ee0629a07373342cd533036fd2f5b44c1866605434")
	req := tbk.NewPrivilegeGetRequest()
	req.SetItemId("556443607855")
	req.SetSiteId("183500025") //mm_48807125_183500025_52280450373
	req.SetAdzoneId("52280450373")
	req.SetRelationId("520045220")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.PrivilegeGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
