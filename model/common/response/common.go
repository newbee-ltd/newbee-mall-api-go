package response

type PageResult struct {
	List       interface{} `json:"list"`
	TotalCount int64       `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	CurrPage   int         `json:"currPage"`
	PageSize   int         `json:"pageSize"`
}
