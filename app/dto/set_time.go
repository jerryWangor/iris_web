package dto

import (
	"github.com/gookit/validate"
)

// 添加菜单
type SetTimeReq struct {
	Id   string `form:"id"`
	Time string `form:"time" validate:"required"` // 时间
}

// 添加菜单表单验证
func (v SetTimeReq) Messages() map[string]string {
	return validate.MS{
		"Time.required": "时间不能为空.",
	}
}
