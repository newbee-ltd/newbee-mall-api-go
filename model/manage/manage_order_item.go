package manage

import "main.go/model/common"

type MallOrderItem struct {
	OrderItemId   int             `json:"orderItemId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderId       int             `json:"orderId" form:"orderId" gorm:"column:order_id;;type:bigint"`
	GoodsId       int             `json:"goodsId" form:"goodsId" gorm:"column:goods_id;;type:bigint"`
	GoodsName     string          `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;type:varchar(200);"`
	GoodsCoverImg string          `json:"goodsCoverImg" form:"goodsCoverImg" gorm:"column:goods_cover_img;comment:商品主图;type:varchar(200);"`
	SellingPrice  int             `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品实际售价;type:int"`
	GoodsCount    int             `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;;type:bigint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
}

func (MallOrderItem) TableName() string {
	return "tb_newbee_mall_order_item"
}
