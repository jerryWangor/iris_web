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

package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var Index = new(IndexController)

type IndexController struct{}

func (c *IndexController) Index(ctx iris.Context) {
	// 获取用户信息
	userInfo := service.Login.GetProfile(utils.Uid(ctx))
	// 获取菜单列表
	menuList := service.Menu.GetPermissionMenuTreeList(userInfo.Id)
	ctx.ViewData("userInfo", userInfo)
	ctx.ViewData("menuList", menuList)
	// 渲染模板
	ctx.View("index.html")
}

func (c *IndexController) Main(ctx iris.Context) {
	// 渲染模板
	ctx.View("welcome.html")
}

func (c *IndexController) UserInfo(ctx iris.Context) {
	if ctx.Method() == "POST" {
		// 参数验证
		var req dto.UserInfoReq
		if err := ctx.ReadForm(&req); err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		// 参数校验
		v := validate.Struct(&req)
		if !v.Validate() {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  v.Errors.One(),
			})
			return
		}
		// 更新信息
		_, err := service.User.UpdateUserInfo(req, utils.Uid(ctx))
		if err != nil {
			ctx.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// 返回结果
		ctx.JSON(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
		return
	}
	// 获取用户信息
	userInfo := service.Login.GetProfile(utils.Uid(ctx))
	// 绑定数据
	ctx.ViewData("userInfo", userInfo)
	// 渲染模板
	ctx.View("user_info/index.html")
}

func (c *IndexController) UpdatePwd(ctx iris.Context) {
	// 参数验证
	var req dto.UpdatePwd
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 参数校验
	v := validate.Struct(&req)
	if !v.Validate() {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
		return
	}
	// 调用更新密码方法
	rows, err := service.User.UpdatePwd(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新密码成功",
	})
}

func (c *IndexController) Logout(ctx iris.Context) {
	// 清除全部SESSION
	sessions.Get(ctx).Clear()
	//// 删除指定的SESSION
	//sessions.Get(ctx).Delete(conf.USER_ID)
	// 跳转登录页
	ctx.Redirect("/login")
}
