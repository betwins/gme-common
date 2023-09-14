package order

type QueryListReq struct {
	OrderId   string `form:"orderId" `  // 订单id
	PageIndex int    `form:"pageIndex"` // "页码"
	PageSize  int    `form:"pageSize"`  // "每页个数"
}
