// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package dto

import (
	//"easygoadmin/app/model"
	"easygoadmin/app/model"
	"github.com/gookit/validate"
)

// 分页查询条件
type DeptPageReq struct {
	Name  string `form:"name"`  // 部门名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加部门
type DeptAddReq struct {
	Id       int    `form:"id"`
	Name     string `form:"name" validate:"required"`
	Code     string `form:"code" validate:"required"`
	Fullname string `form:"fullname" validate:"required"`
	Type     int    `form:"type" validate:"int"`
	Pid      int    `form:"pid" validate:"int"`
	Sort     int    `form:"sort" validate:"int"`
	Note     string `form:"note"`
}

// 添加部门表单验证
func (v DeptAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":     "部门名称不能为空.",
		"Code.required":     "部门编码不能为空.",
		"Fullname.required": "部门全称不能为空.",
		"Type.int":          "请选择部门类型.",
		"Sort.int":          "排序不能为空.",
	}
}

// 部门编辑
type DeptUpdateReq struct {
	Id       int    `form:"id" validate:"int"`
	Name     string `form:"name" validate:"required"`
	Code     string `form:"code" validate:"required"`
	Fullname string `form:"fullname" validate:"required"`
	Type     int    `form:"type" validate:"int"`
	Pid      int    `form:"pid" validate:"int"`
	Sort     int    `form:"sort" validate:"int"`
	Note     string `form:"note"`
}

// 更新部门表单验证
func (v DeptUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":            "部门ID不能为空.",
		"Name.required":     "部门名称不能为空.",
		"Code.required":     "部门编码不能为空.",
		"Fullname.required": "部门全称不能为空.",
		"Type.int":          "请选择部门类型.",
		"Sort.int":          "排序不能为空.",
	}
}

// 部门树结构
type DeptTreeNode struct {
	model.Dept
	Children []*DeptTreeNode `json:"children"` // 子栏目
}
