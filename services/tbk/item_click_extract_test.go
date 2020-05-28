package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestItemClickExtract(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewItemClickExtractRequest()
	req.SetClickUrl("https://s.click.taobao.com/t?e=m%3D2%26s%3Djq7wKv1HMgAcQipKwQzePOeEDrYVVa64K7Vc7tFgwiHjf2vlNIV67hldFvpXfyjNsmcYjUfw1pLT6y8UW0x6Ul%2FZMKpynZSgUyMoWRC37ckebvs5ByqxLPAy%2Fay3dFHhPplz8soEI3Cs607XS4XlaiyzR7DYgLslEiM%2FlSG%2FbZSYNdUBCVZhTaLPESru7OP3xiXvDf8DaRs%3D&scm=null&pvid=null&app_pvid=59590_11.26.37.227_534_1583335591696&ptl=floorId%3A17741&originalFloorId%3A17741&app_pvid%3A59590_11.26.37.227_534_1583335591696&union_lens=lensId%3APUB%401583335586%40da652561-53fa-4947-ad94-6ab158209d40_605696975243%40026ZqIg9M9sSG5TnZQzgq9ei")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.ItemClickExtract(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
