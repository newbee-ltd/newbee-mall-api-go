package manage

import (
	"errors"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
	"time"
)

type ManageCarouselService struct {
}

func (m *ManageCarouselService) CreateCarousel(req manageReq.MallCarouselAddParam) (err error) {
	carouseRank, _ := strconv.Atoi(req.CarouselRank)
	mallCarousel := manage.MallCarousel{
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: carouseRank,
		CreateTime:   common.JSONTime{Time: time.Now()},
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}
	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = utils.Verify(mallCarousel, utils.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Create(&mallCarousel).Error
	return err
}

func (m *ManageCarouselService) DeleteCarousel(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&manage.MallCarousel{}, "carousel_id in ?", ids.Ids).Error
	return err
}

func (m *ManageCarouselService) UpdateCarousel(req manageReq.MallCarouselUpdateParam) (err error) {
	if errors.Is(global.GVA_DB.Where("carousel_id = ?", req.CarouselId).First(&manage.MallCarousel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("未查询到记录！")
	}
	carouseRank, _ := strconv.Atoi(req.CarouselRank)
	mallCarousel := manage.MallCarousel{
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: carouseRank,
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}
	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = utils.Verify(mallCarousel, utils.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Where("carousel_id = ?", req.CarouselId).UpdateColumns(&mallCarousel).Error
	return err
}

func (m *ManageCarouselService) GetCarousel(id int) (err error, mallCarousel manage.MallCarousel) {
	err = global.GVA_DB.Where("carousel_id = ?", id).First(&mallCarousel).Error
	return
}

func (m *ManageCarouselService) GetCarouselInfoList(info manageReq.MallCarouselSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallCarousel{})
	var mallCarousels []manage.MallCarousel
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("carousel_rank desc").Find(&mallCarousels).Error
	return err, mallCarousels, total
}
