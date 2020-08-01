package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestPrivilege(t *testing.T) {

	credential := &auth.Credentials{AppKey: "25254678", AppSecret: []byte("5e5e16ddb6c2d780f91401a9c990b7d6")}

	c := tbk.NewClientWith(credential)
	/* 620291377030a4a8608b5ZZ68cdbb018c32ac162afd52141954274556 马姗姗
	// pid mm_464100192_651000075_109921200152
	// 渠道ID 2544959300
	// eid gsid2412
	*/
	/* 6201f07cccfb341ZZadebde93ea2c7a95384e6c7d481996133167095545 自营
	// pid 45828539 92383100212
	// 渠道ID 561721281
	// sid 521177352
	// eid gsid2412
	*/

	//自营
	c.SetSession("620291377030a4a8608b5ZZ68cdbb018c32ac162afd52141954274556")
	req := tbk.NewPrivilegeGetRequest()
	req.SetItemId("622911258398")
	req.SetSiteId("651000075") //mm_54103422_45828539_99513750139
	req.SetAdzoneId("109921200152")
	req.SetRelationId("2544959300") //
	//req.SetSpecialId("521177352")
	//req.SetExternalId("gsid2412")
	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.PrivilegeGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)
		return
	}
	fmt.Printf("正常响应:%#v\n", resp)
	fmt.Println("淘口令：", taoCode(resp.TbkPrivilegeGetResponse.Result.Data.CouponClickUrl))
	fmt.Println("min_commission_rate: ", resp.TbkPrivilegeGetResponse.Result.Data.MinCommissionRate)
	fmt.Println("max_commission_rate: ", resp.TbkPrivilegeGetResponse.Result.Data.MaxCommissionRate)
}

func taoCode(purl string) string {
	credential := &auth.Credentials{AppKey: "25254678", AppSecret: []byte("5e5e16ddb6c2d780f91401a9c990b7d6")}

	c := tbk.NewClientWith(credential)
	req := tbk.NewTpwdCreateRequest()
	req.SetText("测试淘口令")
	req.SetUrl(purl)

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.TpwdCreate(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return ""
	}
	fmt.Printf("正常响应:%#v\n", resp)
	return resp.TbkTpwdCreateResponse.Data.Model
}