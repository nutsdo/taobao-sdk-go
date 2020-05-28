package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestTpwdCreate(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewTpwdCreateRequest()
	req.SetText("测试淘口令")
	req.SetUrl("https://uland.taobao.com/coupon/edetail?e=JiCgVr1j%2BG0GQASttHIRqSOR6G7egnNLj1oVGjCLR5fyi1aSyqbpmkpzDMCNLbviAMZjpJtsBIJzfBR45d0O7eSFwsw9whwjDfqEFBOhTcwIEVMQF19gsGCneBWRsmklXV01yT5NCM%2FB0bVmWMCOxWCBLf0Ce8vJQdUCsKS61xH3LTDo28QWQkbYLu1k1GIxF6Y%2FSGkgg%2B%2BMOmm%2BVf6hvw2U0kG5qm9DYwOD23XOnRHymNXcL2pzBIcAHJUIj3zOuBTHTfgZwSk%3D&traceId=0b5218ed15833013628107033e241e&traffic_flag=lm&scm=20140618.1.01010001.s101c6&spm=a21wq.8595780.1000.3&src=tblm_lmapp&union_lens=lensId%3AAPP%401583300638%400b5952e0_0cd4_1709eff9f7b_79f9%40042erumVJGWQTKcxBi2iEZfb&un=e16318b025782788143b731291a8a715&share_crt_v=1&ut_sk=1.utdid_null_1583301363055.TaoPassword-Outside.lianmeng-app&sp_tk=77+leFhPbjFna0JOOTHvv6U=")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.TpwdCreate(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
