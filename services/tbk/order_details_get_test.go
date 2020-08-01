package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestOrderDetailsGet(t *testing.T) {
	mycredential := &auth.Credentials{AppKey: "30256862", AppSecret: []byte("701b47bcebbf5350a19045fb58aaa198")}
	c := tbk.NewClientWith(mycredential)

	req := tbk.NewOrderDetailsGetRequest()
	req.SetQueryType("2")
	req.SetStartTime("2020-07-22 16:18:00")
	req.SetEndTime("2020-07-22 17:00:00")
	req.SetOrderScene("2")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.OrderDetailsGet(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}

	//fmt.Printf("正常响应:%#v", resp)
	fmt.Println("订单列表:")
	for _, order := range resp.TbkOrderDetailsGetResponse.Data.Results.PublisherOrderDto {
		//fmt.Printf("订单号: %s, 付款时间: %s , 订单创建时间: %s, 订单状态:%d , 订单类型:%d \n",
		//	order.TradeParentId, order.TbPaidTime, order.TkCreateTime, order.TkStatus, order.RelationId)
		fmt.Printf("%#v\n",order)
		fmt.Println("渠道ID:", order.RelationId)
		fmt.Println("商品ID:", order.ItemId)
		fmt.Println("商品标题:", order.ItemTitle)
		fmt.Println("推广标识:", order.FlowSource)
		fmt.Println("佣金比例:", order.IncomeRate)
		fmt.Println("预估收入:", order.PubSharePreFee)
		fmt.Println("技术服务费:", order.AlimamaShareFee)
		fmt.Println("专项技术服务费:", order.ServiceFeeDtoList)
	}
}
