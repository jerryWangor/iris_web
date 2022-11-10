package dto

import (
	"github.com/gookit/validate"
)

// 添加菜单
type SendMailReq struct {
	Id         int    `form:"id"`
	Title      string `form:"title"`
	Content    string `form:"content"`
	Type       int    `form:"type"`
	RoleList   string `form:"role_list" validate:"required"`
	ItemList   string `form:"item_list" validate:"required"`
	Status     int    `form:"status"`
	ReturnInfo string `form:"return_info"`
}

// 添加菜单表单验证
func (v SendMailReq) Messages() map[string]string {
	return validate.MS{
		"RoleList.required": "角色列表不能为空.",
		"ItemList.required": "道具列表不能为空.",
	}
}
