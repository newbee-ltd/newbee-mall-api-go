package service

import (
	"main.go/service/example"
	"main.go/service/mall"
	"main.go/service/manage"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
	MallServiceGroup    mall.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
