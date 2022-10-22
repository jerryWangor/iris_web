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

package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	"easygoadmin/app/widget"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// 注册路由
func RegisterRouter(app *iris.Application) {
	// 注册SESSION中间件
	session := sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})
	// SESSION中间件
	app.Use(session.Handler())
	// 登录验证中间件
	app.Use(middleware.CheckLogin)

	// 视图文件目录 每次请求时自动重载模板
	tmpl := iris.HTML("./view", ".html").Reload(true)

	// 注册自定义视图函数
	tmpl.AddFunc("safe", widget.Safe)
	tmpl.AddFunc("date", widget.Date)
	tmpl.AddFunc("widget", widget.Widget)
	tmpl.AddFunc("query", widget.Query)
	tmpl.AddFunc("add", widget.Add)
	tmpl.AddFunc("edit", widget.Edit)
	tmpl.AddFunc("delete", widget.Delete)
	tmpl.AddFunc("dall", widget.Dall)
	tmpl.AddFunc("expand", widget.Expand)
	tmpl.AddFunc("collapse", widget.Collapse)
	tmpl.AddFunc("addz", widget.Addz)
	tmpl.AddFunc("switch", widget.Switch)
	tmpl.AddFunc("select", widget.Select)
	tmpl.AddFunc("submit", widget.Submit)
	tmpl.AddFunc("icon", widget.Icon)
	tmpl.AddFunc("transfer", widget.Transfer)
	tmpl.AddFunc("upload_image", widget.UploadImage)
	tmpl.AddFunc("album", widget.Album)
	tmpl.AddFunc("item", widget.Item)
	tmpl.AddFunc("kindeditor", widget.Kindeditor)
	tmpl.AddFunc("checkbox", widget.Checkbox)
	tmpl.AddFunc("radio", widget.Radio)
	tmpl.AddFunc("city", widget.City)

	// 注册视图
	app.RegisterView(tmpl)

	// 静态文件
	app.HandleDir("/static", "./public/static")

	// 错误请求配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("error/404.html")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.View("error/500.html")
	})

	// 登录、主页
	index := app.Party("/")
	{
		index.Get("/", controller.Index.Index)
		index.Any("/login", controller.Login.Login)
		index.Get("/captcha", controller.Login.Captcha)
		index.Get("/index", controller.Index.Index)
		index.Get("/main", controller.Index.Main)
		index.Any("/userInfo", controller.Index.UserInfo)
		index.Any("/updatePwd", controller.Index.UpdatePwd)
		index.Get("/logout", controller.Index.Logout)
	}

	// 文件上传
	upload := app.Party("/upload")
	{
		upload.Post("/uploadImage", controller.Upload.UploadImage)
		upload.Post("/uploadEditImage", controller.Upload.UploadEditImage)
	}

	// 职级管理
	level := app.Party("/level")
	{
		level.Get("/index", controller.Level.Index)
		level.Post("/list", controller.Level.List)
		level.Get("/edit/{id:int}", controller.Level.Edit)
		level.Post("/add", controller.Level.Add)
		level.Post("/update", controller.Level.Update)
		level.Post("/delete/{id:int}", controller.Level.Delete)
		level.Post("/setStatus", controller.Level.Status)
	}

	// 岗位管理
	position := app.Party("/position")
	{
		position.Get("/index", controller.Position.Index)
		position.Post("/list", controller.Position.List)
		position.Get("/edit/{id:int}", controller.Position.Edit)
		position.Post("/add", controller.Position.Add)
		position.Post("/update", controller.Position.Update)
		position.Post("/delete/{id:int}", controller.Position.Delete)
		position.Post("/setStatus", controller.Position.Status)
	}

	/* 角色路由 */
	role := app.Party("role")
	{
		role.Get("/index", controller.Role.Index)
		role.Post("/list", controller.Role.List)
		role.Get("/edit/{id:int}", controller.Role.Edit)
		role.Post("/add", controller.Role.Add)
		role.Post("/update", controller.Role.Update)
		role.Post("/delete/{id:int}", controller.Role.Delete)
		role.Post("/setStatus", controller.Role.Status)
		role.Get("/getRoleList", controller.Role.GetRoleList)
	}

	/* 角色菜单权限 */
	roleMenu := app.Party("rolemenu")
	{
		roleMenu.Get("/index/{roleId:int}", controller.RoleMenu.Index)
		roleMenu.Post("/save", controller.RoleMenu.Save)
	}

	/* 部门管理 */
	dept := app.Party("dept")
	{
		dept.Get("/index", controller.Dept.Index)
		dept.Post("/list", controller.Dept.List)
		dept.Get("/edit/{id:int}", controller.Dept.Edit)
		dept.Get("/edit/{id:int}/{pid:string}", controller.Dept.Edit)
		dept.Post("/add", controller.Dept.Add)
		dept.Post("/update", controller.Dept.Update)
		dept.Post("/delete/{id:int}", controller.Dept.Delete)
		dept.Get("/getDeptList", controller.Dept.GetDeptList)
	}

	/* 用户管理 */
	user := app.Party("user")
	{
		user.Get("/index", controller.User.Index)
		user.Post("/list", controller.User.List)
		user.Get("/edit/{id:int}", controller.User.Edit)
		user.Post("/add", controller.User.Add)
		user.Post("/update", controller.User.Update)
		user.Post("/delete/{id:int}", controller.User.Delete)
		user.Post("/setStatus", controller.User.Status)
		user.Post("/resetPwd", controller.User.ResetPwd)
	}

	/* 菜单管理 */
	menu := app.Party("menu")
	{
		menu.Get("/index", controller.Menu.Index)
		menu.Post("/list", controller.Menu.List)
		menu.Get("/edit/{id:int}", controller.Menu.Edit)
		menu.Get("/edit/{id:int}/{pid:string}", controller.Menu.Edit)
		menu.Post("/add", controller.Menu.Add)
		menu.Post("/update", controller.Menu.Update)
		menu.Post("/delete/{id:int}", controller.Menu.Delete)
	}

	/* 城市管理 */
	city := app.Party("city")
	{
		city.Get("/index", controller.City.Index)
		city.Post("/list", controller.City.List)
		city.Get("/edit/{id:int}", controller.City.Edit)
		city.Get("/edit/{id:int}/{pid:string}", controller.City.Edit)
		city.Post("/add", controller.City.Add)
		city.Post("/update", controller.City.Update)
		city.Post("/delete/{id:int}", controller.City.Delete)
		city.Post("/getChilds", controller.City.GetChilds)
	}

	/* 通知管理 */
	notice := app.Party("notice")
	{
		notice.Get("/index", controller.Notice.Index)
		notice.Post("/list", controller.Notice.List)
		notice.Get("/edit/{id:int}", controller.Notice.Edit)
		notice.Post("/add", controller.Notice.Add)
		notice.Post("/update", controller.Notice.Update)
		notice.Post("/delete/{id:int}", controller.Notice.Delete)
		notice.Post("/setStatus", controller.Notice.Status)
	}

	/* 统计分析 */
	analysis := app.Party("analysis")
	{
		analysis.Get("/index", controller.Analysis.Index)
	}

	/* 代码生成器 */
	generate := app.Party("generate")
	{
		generate.Get("/index", controller.Generate.Index)
		generate.Post("/list", controller.Generate.List)
		generate.Post("/generate", controller.Generate.Generate)
		generate.Post("/batchGenerate", controller.Generate.BatchGenerate)
	}

}
