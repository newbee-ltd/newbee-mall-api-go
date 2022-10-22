package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"main.go/global"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallGoodsInfoService struct {
}

// MallGoodsListBySearch 商品搜索分页
func (m *MallGoodsInfoService) MallGoodsListBySearch(pageNumber int, goodsCategoryId int, keyword string, orderBy string) (err error, searchGoodsList []mallRes.GoodsSearchResponse, total int64) {
	// 根据搜索条件查询
	var goodsList []manage.MallGoodsInfo
	db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
	if keyword != "" {
		db.Where("goods_name like ? or goods_intro like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if goodsCategoryId >= 0 {
		db.Where("goods_category_id= ?", goodsCategoryId)
	}
	err = db.Count(&total).Error
	switch orderBy {
	case "new":
		db.Order("goods_id desc")
	case "price":
		db.Order("selling_price asc")
	default:
		db.Order("stock_num desc")
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	err = db.Limit(limit).Offset(offset).Find(&goodsList).Error
	// 返回查询结果
	for _, goods := range goodsList {
		searchGoods := mallRes.GoodsSearchResponse{
			GoodsId:       goods.GoodsId,
			GoodsName:     utils.SubStrLen(goods.GoodsName, 28),
			GoodsIntro:    utils.SubStrLen(goods.GoodsIntro, 28),
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
		}
		searchGoodsList = append(searchGoodsList, searchGoods)
	}
	return
}

// GetMallGoodsInfo 获取商品信息
func (m *MallGoodsInfoService) GetMallGoodsInfo(id int) (err error, res mallRes.GoodsInfoDetailResponse) {
	var mallGoodsInfo manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	if mallGoodsInfo.GoodsSellStatus != 0 {
		return errors.New("商品已下架"), mallRes.GoodsInfoDetailResponse{}
	}
	err = copier.Copy(&res, &mallGoodsInfo)
	if err != nil {
		return err, mallRes.GoodsInfoDetailResponse{}
	}
	var list []string
	list = append(list, mallGoodsInfo.GoodsCarousel)
	res.GoodsCarouselList = list

	return
}
