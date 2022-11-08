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
)

var SetTime = new(SetTimeController)

type SetTimeController struct{}

func (c *SetTimeController) Index(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/layout.html")
	// 渲染模板
	ctx.View("set_time/index.html")
}

func (c *SetTimeController) List(ctx iris.Context) {
	// 调用获取列表方法
	lists, err := service.SetTime.GetList()
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回结果集
	ctx.JSON(common.JsonResult{
		Code: 0,
		Data: lists,
		Msg:  "操作成功",
	})
}

func (c *SetTimeController) Set(ctx iris.Context) {
	// 模板布局
	ctx.ViewLayout("public/form.html")
	// 渲染模板
	ctx.View("set_time/set.html")
}

func (c *SetTimeController) SetTime(ctx iris.Context) {
	// 添加对象
	var req dto.SetTimeReq
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

	// 调用GM服务
	time := struct {
		Time string `json:"time"`
	}{
		Time: req.Time,
	}
	message, err := json.Marshal(time)
	pass, err := gmd5.Encrypt(constant.GMKEY + string(message))
	query := url.Values{}
	query.Add("pass", pass)
	query.Add("message", string(message))
	fmt.Println(query.Encode())
	url := constant.GMURL + "/settime?" + query.Encode()

	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -2,
			Msg:  err.Error(),
		})
	}
	defer resp.Body.Close()

	// 读取数据
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(common.JsonResult{
			Code: -3,
			Msg:  err.Error(),
		})
	}
	var jsonResp model.JsonResp
	json.Unmarshal(bds, &jsonResp)
	fmt.Println(string(bds))
	if jsonResp.Code == 0 && jsonResp.Message == "success" {
		// 写入数据库
		service.SetTime.Add(req, utils.Uid(ctx))
	} else {
		ctx.JSON(common.JsonResult{
			Code: -4,
			Msg:  string(bds),
		})
	}

	// 设置成功
	ctx.JSON(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}
