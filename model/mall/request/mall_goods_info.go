package request

type GoodsSearchParams struct {
	Keyword         string `form:"keyword"`
	GoodsCategoryId int    `form:"goodsCategoryId"`
	OrderBy         string `form:"orderBy"`
	PageNumber      int    `form:"pageNumber"`
}
