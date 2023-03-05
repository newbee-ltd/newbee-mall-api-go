package mall

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallShopCartService struct {
}

// GetMyShoppingCartItems 不分页
func (m *MallShopCartService) GetMyShoppingCartItems(token string) (err error, cartItems []mallRes.CartItemResponse) {
	var userToken mall.MallUserToken
	var shopCartItems []mall.MallShoppingCartItem
	var goodsInfos []manage.MallGoodsInfo
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItems
	}
	global.GVA_DB.Where("user_id=? and is_deleted = 0", userToken.UserId).Find(&shopCartItems)
	var goodsIds []int
	for _, shopcartItem := range shopCartItems {
		goodsIds = append(goodsIds, shopcartItem.GoodsId)
	}
	global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&goodsInfos)
	goodsMap := make(map[int]manage.MallGoodsInfo)
	for _, goodsInfo := range goodsInfos {
		goodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, v := range shopCartItems {
		var cartItem mallRes.CartItemResponse
		copier.Copy(&cartItem, &v)
		if _, ok := goodsMap[v.GoodsId]; ok {
			goodsInfo := goodsMap[v.GoodsId]
			cartItem.GoodsName = goodsInfo.GoodsName
			cartItem.GoodsCoverImg = goodsInfo.GoodsCoverImg
			cartItem.SellingPrice = goodsInfo.SellingPrice
		}
		cartItems = append(cartItems, cartItem)
	}

	return
}

func (m *MallShopCartService) SaveMallCartItem(token string, req mallReq.SaveCartItemParam) (err error) {
	if req.GoodsCount < 1 {
		return errors.New("商品数量不能小于 1 ！")

	}
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItems []mall.MallShoppingCartItem
	// 是否已存在商品
	err = global.GVA_DB.Where("user_id = ? and goods_id = ? and is_deleted = 0", userToken.UserId, req.GoodsId).Find(&shopCartItems).Error
	if err != nil {
		return errors.New("已存在！无需重复添加！")
	}
	err = global.GVA_DB.Where("goods_id = ? ", req.GoodsId).First(&manage.MallGoodsInfo{}).Error
	if err != nil {
		return errors.New(" 商品为空")
	}
	var total int64
	global.GVA_DB.Where("user_id =?  and is_deleted = 0", userToken.UserId).Count(&total)
	if total > 20 {
		return errors.New("超出购物车最大容量！")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = copier.Copy(&shopCartItem, &req); err != nil {
		return err
	}
	shopCartItem.UserId = userToken.UserId
	shopCartItem.CreateTime = common.JSONTime{Time: time.Now()}
	shopCartItem.UpdateTime = common.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) UpdateMallCartItem(token string, req mallReq.UpdateCartItemParam) (err error) {
	//超出单个商品的最大数量
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = global.GVA_DB.Where("cart_item_id=? and is_deleted = 0", req.CartItemId).First(&shopCartItem).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if shopCartItem.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	shopCartItem.GoodsCount = req.GoodsCount
	shopCartItem.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) DeleteMallCartItem(token string, id int) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItem mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).First(&shopCartItem).Error
	if err != nil {
		return
	}
	if userToken.UserId != shopCartItem.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).UpdateColumns(&mall.MallShoppingCartItem{IsDeleted: 1}).Error
	return
}

func (m *MallShopCartService) GetCartItemsForSettle(token string, cartItemIds []int) (err error, cartItemRes []mallRes.CartItemResponse) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItemRes
	}
	var shopCartItems []mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id in (?) and user_id = ? and is_deleted = 0", cartItemIds, userToken.UserId).Find(&shopCartItems).Error
	if err != nil {
		return
	}
	_, cartItemRes = getMallShoppingCartItemVOS(shopCartItems)
	//购物车算价
	priceTotal := 0
	for _, cartItem := range cartItemRes {
		priceTotal = priceTotal + cartItem.GoodsCount*cartItem.SellingPrice
	}
	return
}

// 购物车数据转换
func getMallShoppingCartItemVOS(cartItems []mall.MallShoppingCartItem) (err error, cartItemsRes []mallRes.CartItemResponse) {
	var goodsIds []int
	for _, cartItem := range cartItems {
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newBeeMallGoods []manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&newBeeMallGoods).Error
	if err != nil {
		return
	}

	newBeeMallGoodsMap := make(map[int]manage.MallGoodsInfo)
	for _, goodsInfo := range newBeeMallGoods {
		newBeeMallGoodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, cartItem := range cartItems {
		var cartItemRes mallRes.CartItemResponse
		copier.Copy(&cartItemRes, &cartItem)
		// 是否包含key
		if _, ok := newBeeMallGoodsMap[cartItemRes.GoodsId]; ok {
			newBeeMallGoodsTemp := newBeeMallGoodsMap[cartItemRes.GoodsId]
			cartItemRes.GoodsCoverImg = newBeeMallGoodsTemp.GoodsCoverImg
			goodsName := utils.SubStrLen(newBeeMallGoodsTemp.GoodsName, 28)
			cartItemRes.GoodsName = goodsName
			cartItemRes.SellingPrice = newBeeMallGoodsTemp.SellingPrice
			cartItemsRes = append(cartItemsRes, cartItemRes)
		}
	}
	return
}
