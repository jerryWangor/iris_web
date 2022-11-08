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

package middleware

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"reflect"
	"strings"
)

// 登录验证中间件
func CheckAuth(ctx iris.Context) {

	// 放行设置
	urlItem := []string{"/captcha", "/login"}
	whiteItem := []string{"/", "/main", "/index", "/userInfo", "/updatePwd", "/logout"}
	if !utils.InStringArray(ctx.Path(), urlItem) && !strings.Contains(ctx.Path(), "static") {
		// 判断不在白名里，判断不是超级管理员
		if !service.IsSuperAdmin(ctx) && utils.IsLogin(ctx) && !utils.InStringArray(ctx.Path(), whiteItem) {
			// 如果登录了就检查权限
			user := service.GetUserInfo(ctx)
			// 查询该用户的权限列表
			val := utils.RedisClient.Get(utils.GetRedisUidKey(user.Id, conf.USER_MENU_LIST)).Val()
			mlist := make([]model.Menu, 0)
			json.Unmarshal([]byte(val), &mlist)
			flag := false
			for _, v := range mlist {
				if reflect.DeepEqual(v.Url, ctx.Path()) {
					flag = true
				}
			}
			if flag == false {
				ctx.JSON(common.JsonResult{
					Code: -1,
					Msg:  "没有权限",
				})
				return
			}
		}
	}
	ctx.Next()
}
