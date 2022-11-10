package controller

import (
	"easygoadmin/app/constant"
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gmd5"
	"encoding/json"
	"fmt"
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var SendMail = new(SendMailController)

type SendMailController struct{}

func (c *SendMailController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("send_mail/index.html")
}

func (c *SendMailController) List(ctx iris.Context) {
	// 调用获取列表方法
	lists, err := service.SendMail.GetList()
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	//for k, _ := range lists {
	//	itemlist, _ := json.Marshal(lists[k].ItemList)
	//	lists[k].ItemList = itemlist.(string)
	//}

	// 返回结果集
	ctx.JSON(common.JsonResult{
		Code: 0,
		Data: lists,
		Msg:  "操作成功",
	})
}

func (c *SendMailController) Mail(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("send_mail/mail.html")
}

func (c *SendMailController) SendMail(ctx iris.Context) {
	// 添加对象
	var req dto.SendMailReq
	// 参数绑定
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

	// 先入库
	req.Title = "测试"
	req.Content = "测试"
	req.Type = 1
	req.Status = 0

	id, err := service.SendMail.Add(req, utils.Uid(ctx))
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用GM服务
	var items []model.GmMailItem
	err = json.Unmarshal([]byte(req.ItemList), &items)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	var target []model.GmMailTarget
	// 逗号分割role_list，循环
	roleArr := strings.Split(req.RoleList, "，")
	target = append(target, model.GmMailTarget{Type: "list_role", Args: roleArr})

	mail := model.GmMail{
		Title:   req.Title,
		Content: req.Content,
		Items:   items,
		Target:  target,
		Other:   "",
	}

	message, err := json.Marshal(mail)
	pass, err := gmd5.Encrypt(constant.GMKEY + string(message))
	query := url.Values{}
	query.Add("pass", pass)
	query.Add("message", string(message))
	fmt.Println("message", string(message))
	url := constant.GMURL + "sendmail?" + query.Encode()

	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -2,
			Msg:  err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// 读取数据
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -3,
			Msg:  err.Error(),
		})
		return
	}
	var jsonResp model.JsonResp
	json.Unmarshal(bds, &jsonResp)
	fmt.Println("接口返回值", string(bds))
	req.Id = id
	req.ReturnInfo = string(bds)
	if jsonResp.Code == 0 && jsonResp.Message == "success" {
		// 更新状态
		req.Status = 1
		service.SendMail.Update(req)
	} else {
		req.Status = 2
		service.SendMail.Update(req)
		ctx.JSON(common.JsonResult{
			Code: -4,
			Msg:  string(bds),
		})
		return
	}

	// 设置成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "发送成功",
	})
}
