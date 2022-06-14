package manage

import (
	"errors"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/enum"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
	"time"
)

type ManageGoodsInfoService struct {
}

// CreateMallGoodsInfo 创建MallGoodsInfo
func (m *ManageGoodsInfoService) CreateMallGoodsInfo(req manageReq.GoodsInfoAddParam) (err error) {
	var goodsCategory manage.MallGoodsCategory
	err = global.GVA_DB.Where("category_id=?  AND is_deleted=0", req.GoodsCategoryId).First(&goodsCategory).Error
	if goodsCategory.CategoryLevel != enum.LevelThree.Code() {
		return errors.New("分类数据异常")
	}
	if !errors.Is(global.GVA_DB.Where("goods_name=? AND goods_category_id=?", req.GoodsName, req.GoodsCategoryId).First(&manage.MallGoodsInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的商品信息")
	}
	originalPrice, _ := strconv.Atoi(req.OriginalPrice)
	sellingPrice, _ := strconv.Atoi(req.SellingPrice)
	stockNum, _ := strconv.Atoi(req.StockNum)
	goodsSellStatus, _ := strconv.Atoi(req.GoodsSellStatus)
	goodsInfo := manage.MallGoodsInfo{
		GoodsName:          req.GoodsName,
		GoodsIntro:         req.GoodsIntro,
		GoodsCategoryId:    req.GoodsCategoryId,
		GoodsCoverImg:      req.GoodsCoverImg,
		GoodsDetailContent: req.GoodsDetailContent,
		OriginalPrice:      originalPrice,
		SellingPrice:       sellingPrice,
		StockNum:           stockNum,
		Tag:                req.Tag,
		GoodsSellStatus:    goodsSellStatus,
		CreateTime:         common.JSONTime{Time: time.Now()},
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(goodsInfo, utils.GoodsAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Create(&goodsInfo).Error
	return err
}

// DeleteMallGoodsInfo 删除MallGoodsInfo记录
func (m *ManageGoodsInfoService) DeleteMallGoodsInfo(mallGoodsInfo manage.MallGoodsInfo) (err error) {
	err = global.GVA_DB.Delete(&mallGoodsInfo).Error
	return err
}

// ChangeMallGoodsInfoByIds 上下架
func (m *ManageGoodsInfoService) ChangeMallGoodsInfoByIds(ids request.IdsReq, sellStatus string) (err error) {
	intSellStatus, _ := strconv.Atoi(sellStatus)
	//更新字段为0时，不能直接UpdateColumns
	err = global.GVA_DB.Model(&manage.MallGoodsInfo{}).Where("goods_id in ?", ids.Ids).Update("goods_sell_status", intSellStatus).Error
	return err
}

// UpdateMallGoodsInfo 更新MallGoodsInfo记录
func (m *ManageGoodsInfoService) UpdateMallGoodsInfo(req manageReq.GoodsInfoUpdateParam) (err error) {
	goodsId, _ := strconv.Atoi(req.GoodsId)
	originalPrice, _ := strconv.Atoi(req.OriginalPrice)
	stockNum, _ := strconv.Atoi(req.StockNum)
	goodsInfo := manage.MallGoodsInfo{
		GoodsId:            goodsId,
		GoodsName:          req.GoodsName,
		GoodsIntro:         req.GoodsIntro,
		GoodsCategoryId:    req.GoodsCategoryId,
		GoodsCoverImg:      req.GoodsCoverImg,
		GoodsDetailContent: req.GoodsDetailContent,
		OriginalPrice:      originalPrice,
		SellingPrice:       req.SellingPrice,
		StockNum:           stockNum,
		Tag:                req.Tag,
		GoodsSellStatus:    req.GoodsSellStatus,
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(goodsInfo, utils.GoodsAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Where("goods_id=?", goodsInfo.GoodsId).Updates(&goodsInfo).Error
	return err
}

// GetMallGoodsInfo 根据id获取MallGoodsInfo记录
func (m *ManageGoodsInfoService) GetMallGoodsInfo(id int) (err error, mallGoodsInfo manage.MallGoodsInfo) {
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	return
}

// GetMallGoodsInfoInfoList 分页获取MallGoodsInfo记录
func (m *ManageGoodsInfoService) GetMallGoodsInfoInfoList(info manageReq.MallGoodsInfoSearch, goodsName string, goodsSellStatus string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
	var mallGoodsInfos []manage.MallGoodsInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if goodsName != "" {
		db.Where("goods_name =?", goodsName)
	}
	if goodsSellStatus != "" {
		db.Where("goods_sell_status =?", goodsSellStatus)
	}
	err = db.Limit(limit).Offset(offset).Order("goods_id desc").Find(&mallGoodsInfos).Error
	return err, mallGoodsInfos, total
}
