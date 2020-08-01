package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestTpwdCreate(t *testing.T) {
	credential := &auth.Credentials{AppKey: "25254678", AppSecret: []byte("5e5e16ddb6c2d780f91401a9c990b7d6")}

	c := tbk.NewClientWith(credential)
	u := "https://mos.m.taobao.com/union/accurate-return?targetUrl=https%3A%2F%2Fuland.taobao.com%2Fcoupon%2Fedetail%3Fe%3DhoLFan7AtgcGQASttHIRqb0EIlytAlj%252FYqXW9vejEUYTqwA3YD77rCn0KRp0PQNXER5xndNmRk3V4pnBJPLz9zEhJpUUrcnYisiADW1vMjwKb5nDc53rb8Rlqw%252FTjW%252BQLspxGy3zBjaiuj3CzrHp9chouCqZLg4hT%252F1mTtwPHxQZwz%252BY5VeKWY369hICPECOe6L%252Bf9DtnlUCMISgOiKUD9tN7m9cS4TU%26traceId%3D0b105ada15944368879286166ee9b5%26union_lens%3DlensId%3ATAPI%401594436887%400b0fe08b_0e19_1733bd825b1_35fe%4001%26xId%3D1gHtXAPOz9mYKgom14UQKxN2kiPWC75vhpx6msj62Zxs1T3DMoU9qqgl55j92JjikU72ypIBz80iPUlvAyIbwZGiwtIYqitsahUkSZoWqjdg&tkId=54103422&externalId=gsid2412"
	req := tbk.NewTpwdCreateRequest()
	req.SetText("测试淘口令")
	req.SetUrl(u)

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.TpwdCreate(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
