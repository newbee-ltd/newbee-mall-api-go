package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}
