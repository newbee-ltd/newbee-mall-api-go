package router

import (
	"main.go/router/mall"
	"main.go/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
