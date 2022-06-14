package v1

import (
	"main.go/api/v1/mall"
	"main.go/api/v1/manage"
)

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
	MallApiGroup   mall.MallGroup
}

var ApiGroupApp = new(ApiGroup)
