package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallIndexConfigSearch struct {
	manage.MallIndexConfig
	request.PageInfo
}

type MallIndexConfigAddParams struct {
	ConfigName  string `json:"configName"`
	ConfigType  int    `json:"configType"`
	GoodsId     string `json:"goodsId"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigRank  string `json:"configRank"`
}

type MallIndexConfigUpdateParams struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigType  int    `json:"configType"`
	GoodsId     int    `json:"goodsId"`
	ConfigRank  string `json:"configRank"`
}
