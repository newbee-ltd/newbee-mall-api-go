package request

import (
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallGoodsInfoSearch struct {
	manage.MallGoodsInfo
	request.PageInfo
}

type GoodsInfoAddParam struct {
	GoodsName          string `json:"goodsName"`
	GoodsIntro         string `json:"goodsIntro"`
	GoodsCategoryId    int    `json:"goodsCategoryId"`
	GoodsCoverImg      string `json:"goodsCoverImg"`
	GoodsCarousel      string `json:"goodsCarousel"`
	GoodsDetailContent string `json:"goodsDetailContent"`
	OriginalPrice      string `json:"originalPrice"`
	SellingPrice       string `json:"sellingPrice"`
	StockNum           string `json:"stockNum"`
	Tag                string `json:"tag"`
	GoodsSellStatus    string `json:"goodsSellStatus"`
}

// GoodsInfoUpdateParam 更新商品信息的入参
type GoodsInfoUpdateParam struct {
	GoodsId            string          `json:"goodsId"`
	GoodsName          string          `json:"goodsName"`
	GoodsIntro         string          `json:"goodsIntro"`
	GoodsCategoryId    int             `json:"goodsCategoryId"`
	GoodsCoverImg      string          `json:"goodsCoverImg"`
	GoodsCarousel      string          `json:"goodsCarousel"`
	GoodsDetailContent string          `json:"goodsDetailContent"`
	OriginalPrice      string          `json:"originalPrice"`
	SellingPrice       int             `json:"sellingPrice"`
	StockNum           string          `json:"stockNum"`
	Tag                string          `json:"tag"`
	GoodsSellStatus    int             `json:"goodsSellStatus"`
	UpdateUser         int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime         common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改时间;type:datetime"`
}

type StockNumDTO struct {
	GoodsId    int `json:"goodsId"`
	GoodsCount int `json:"goodsCount"`
}
