package request

type PaySuccessParams struct {
	OrderNo string `json:"orderNo"`
	PayType int    `json:"payType"`
}

type OrderSearchParams struct {
	Status     string `form:"status"`
	PageNumber int    `form:"pageNumber"`
}

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds"`
	AddressId   int   `json:"addressId"`
}
