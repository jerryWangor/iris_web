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

package service

import (
	"bytes"
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gfile"
	"easygoadmin/utils/gstr"
	"errors"
	"github.com/kataras/iris/v12"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

var Generate = new(generateService)

type generateService struct{}

func (s *generateService) GetList(req dto.GeneratePageReq) ([]vo.GenerateInfo, error) {
	// 查询SQL
	sql := "SHOW TABLE STATUS"
	// 表名称
	if req.Name != "" {
		sql += " WHERE Name like \"%" + req.Name + "%\""
	}
	// 表描述
	if req.Comment != "" {
		sql += " WHERE Comment like \"%" + req.Comment + "%\""
	}
	// 对象转换
	var list []vo.GenerateInfo
	err := utils.XormDb.SQL(sql).Find(&list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return list, nil
}

func (s *generateService) Generate(req dto.GenerateFileReq, ctx iris.Context) error {
	if utils.AppDebug() {
		return errors.New("演示环境，暂无权限操作")
	}
	// 数据表名
	tableName := req.Name
	// 数据表描述
	moduleTitle := req.Comment
	// 替换“表”
	if gstr.Contains(moduleTitle, "表") {
		moduleTitle = gstr.Replace(moduleTitle, "表", "")
	}
	// 替换“管理”
	if gstr.Contains(moduleTitle, "管理") {
		moduleTitle = gstr.Replace(moduleTitle, "管理", "")
	}
	// 模型名称
	moduleName := gstr.Replace(tableName, "sys_", "")
	// 作者名称
	authorName := "半城风雨"

	// 获取字段列表
	columnList, err := GetColumnList(tableName)
	if err != nil {
		return err
	}

	// 生成控制器
	if err := GenerateController(columnList, authorName, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成控制器
	if err := GenerateDto(columnList, authorName, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成Vo
	if err := GenerateInfoVo(columnList, authorName, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成服务类
	if err := GenerateService(columnList, authorName, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成模块index.html
	if err := GenerateIndex(columnList, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成模块edit.html
	if err := GenerateEdit(columnList, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成模块JS
	if err := GenerateJs(columnList, authorName, moduleName, moduleTitle); err != nil {
		return err
	}

	// 生成菜单权限
	if err := GeneratePermission(moduleName, moduleTitle, utils.Uid(ctx)); err != nil {
		return err
	}

	// 生成路由
	//if err := GenerateRouter(columnList, authorName, moduleName, moduleTitle); err != nil {
	//	return err
	//}

	return nil
}

// 生成控制器
func GenerateController(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 加入数组
		columnList = append(columnList, item)
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("controller.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}, false); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/app/controller/", moduleName, ".go"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成Dto
func GenerateDto(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化查询条件
	queryList := make([]map[string]interface{}, 0)
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 常规字段查询条件
		if columnName == "name" || columnName == "title" {
			// 加入查询条件数组
			queryList = append(queryList, item)
		}
		// 常规下拉选择查询条件
		if _, ok := item["columnValue"]; ok && item["columnValue"] != "" {
			queryList = append(queryList, item)
		}
		// 加入数组
		columnList = append(columnList, item)
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("dto.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
		"queryList":   queryList,
	}, false); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/app/dto/", moduleName, ".go"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成Vo
func GenerateInfoVo(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 加入数组
		columnList = append(columnList, item)
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("vo.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}, false); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/app/vo/", moduleName, ".go"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成服务类
func GenerateService(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化查询条件
	queryList := make([]map[string]interface{}, 0)
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 常规字段查询条件
		if columnName == "name" || columnName == "title" {
			// 加入查询条件数组
			queryList = append(queryList, item)
		}
		// 常规下拉选择查询条件
		if _, ok := item["columnValue"]; ok && item["columnValue"] != "" {
			queryList = append(queryList, item)
		}
		// 加入数组
		columnList = append(columnList, item)
	}
	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("service.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
		"queryList":   queryList,
	}, false); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/app/service/", moduleName, ".go"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成列表页
func GenerateIndex(dataList *common.ArrayList, moduleName string, moduleTitle string) error {
	// 初始化查询条件
	queryList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := item["columnName"]
		if columnName == "name" || columnName == "title" {
			// 加入查询条件数组
			queryList = append(queryList, item)
		}
		// 字段列表格式化，如isVip
		columnName3 := gconv.String(item["columnName3"])
		// 判断是否有columnValue键值
		if _, ok := item["columnValue"]; ok && item["columnValue"] != "" {
			// 加入查询条件数组
			item["columnWidget"] = `{{select "` + gconv.String(columnName3) + `|0|` + gconv.String(item["columnTitle"]) + `|name|id" "` + gconv.String(item["columnValue"]) + `" 0}}`
			queryList = append(queryList, item)
		}
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("index.html", iris.Map{
		"queryList": queryList,
		"funcList1": `{{query "查询"}}
                {{add "添加` + moduleTitle + `" "{}"}}
                {{dall "批量删除"}}`,
		"funcList2": `{{edit "编辑"}}
    {{delete "删除"}}`,
		"defineStart": "{{define \"content\"}}",
		"defineEnd":   "{{end}}",
	}, true); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/view/", moduleName, "/index.html"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成编辑表单
func GenerateEdit(dataList *common.ArrayList, moduleName string, moduleTitle string) error {
	// 初始化表单数组
	formList := make([]map[string]interface{}, 0)
	// 初始化图片数组
	imageList := make([]map[string]interface{}, 0)
	// 初始化多行数组
	rowsList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段类型
		dataType := gconv.String(item["dataType"])
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 字段列表格式化,如IsVip
		columnName2 := gconv.String(item["columnName2"])
		// 字段列表格式化，如isVip
		columnName3 := gconv.String(item["columnName3"])
		// 字段标题
		columnTitle := gconv.String(item["columnTitle"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 图片上传
		if _, ok := item["columnImage"]; ok && item["columnImage"] == true {
			item["columnWidget"] = `{{upload_image "` + columnName3 + `|` + columnTitle + `|90x90|建议上传尺寸450x450" .info.` + columnName2 + ` "" 0}}`
			// 加入数组
			imageList = append(imageList, item)
			continue
		}

		// 多行文本输入
		if _, ok := item["columnText"]; ok && item["columnText"] == true {
			if dataType == "text" {
				item["columnWidget"] = `{{kindeditor "` + columnName3 + `" "default" "80%" 350}}`
			}
			// 加入数组
			rowsList = append(rowsList, item)
			continue
		}
		// 判断是否有columnValue键值
		if _, ok := item["columnValue"]; ok && item["columnValue"] != "" {
			if _, isOk := item["columnSwitch"]; isOk && item["columnSwitch"] == true {
				// 开关组件
				item["columnWidget"] = `{{switch "` + columnName3 + `" "` + gconv.String(item["columnSwitchValue"]) + `" .info.` + columnName2 + `}}`
			} else {
				// 下拉单选组件
				item["columnWidget"] = `{{select "` + columnName3 + `|0|` + columnTitle + `|name|id" "` + gconv.String(item["columnValue"]) + `" .info.` + columnName2 + `}}`
			}
			// 加入数组
			formList = append(formList, item)
			continue
		}
		// 日期组件
		if dataType == "date" || dataType == "datetime" {
			item["columnWidget"] = `{{date "` + columnName + `|1|` + columnTitle + `|` + dataType + `" .info.` + columnName2 + `}}`
			formList = append(formList, item)
			continue
		}
		// 加入数组
		formList = append(formList, item)
	}

	// 初始化数据列数组
	columnList := make([]map[string]interface{}, 0)

	// 根据控制的个数实行分列显示(一行两列)
	if len(formList)+len(imageList)+len(rowsList) > 10 {
		// 一行两列排列
	} else {
		// 单行排列
		columnList = formList
		// 图片
		if len(imageList) > 0 {
			// 遍历
			for _, v := range imageList {
				columnList = append(columnList, v)
			}
		}
		// 多行文本
		if len(rowsList) > 0 {
			// 遍历
			for _, v := range rowsList {
				columnList = append(columnList, v)
			}
		}
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("edit.html", iris.Map{
		"columnList":   columnList,
		"submitWidget": `{{submit "submit|立即保存,close|关闭" 1 ""}}`,
		"defineStart":  "{{define \"form\"}}",
		"defineEnd":    "{{end}}",
	}, true); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/view/", moduleName, "/edit.html"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成模块JS文件
func GenerateJs(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 加入数组
		columnList = append(columnList, item)
	}

	// 读取HTML模板并绑定数据
	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("js.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}, true); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/public/static/module/easygoadmin_", moduleName, ".js"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 生成菜单和权限
func GeneratePermission(modelName string, modelTitle string, userId int) error {
	// 查询记录
	info := &model.Menu{Permission: "sys:" + modelName + ":index"}
	has, err := info.Get()
	if err != nil || !has {
		return err
	}
	// 创建菜单
	var entity model.Menu
	entity.Name = modelTitle
	entity.Icon = "layui-icon-component"
	entity.Url = "/" + modelName + "/index"
	entity.Pid = 225
	entity.Type = 0
	entity.Permission = "sys:" + modelName + ":index"
	entity.Status = 1
	entity.Target = 1
	entity.Sort = 10
	entity.CreateUser = userId
	entity.CreateTime = utils.GetNowTimeTime()
	entity.UpdateUser = userId
	entity.UpdateTime = utils.GetNowTimeTime()
	entity.Mark = 1
	// 记录ID
	menuId := 0
	// 插入或更新记录
	if info == nil {
		// 创建菜单
		_, err := utils.XormDb.Insert(entity)
		if err != nil {
			return err
		}
		// 菜单ID
		menuId = entity.Id
	} else {
		// 更新菜单
		_, err := utils.XormDb.Id(info.Id).Update(entity)
		if err != nil {
			return err
		}
		// 菜单ID
		menuId = info.Id
	}

	// 删除当前菜单的全部节点
	utils.XormDb.Where("pid=?", menuId).Delete(&model.Menu{})

	// 创建节点
	funcList := []int{1, 5, 10, 15, 20, 25, 30}
	for _, v := range funcList {
		// 实例化对象
		var item model.Menu
		item.Pid = menuId
		item.Type = 1
		item.Status = 1
		item.Target = 1
		item.Sort = v
		item.CreateUser = userId
		item.CreateTime = utils.GetNowTimeTime()
		item.UpdateUser = userId
		item.UpdateTime = utils.GetNowTimeTime()
		item.Mark = 1

		// 权限节点
		if v == 1 {
			// 列表
			item.Name = "查询" + modelTitle
			item.Url = "/" + modelName + "/list"
			item.Permission = "sys:" + modelName + ":list"
		} else if v == 5 {
			// 添加
			item.Name = "添加" + modelTitle
			item.Url = "/" + modelName + "/add"
			item.Permission = "sys:" + modelName + ":add"
		} else if v == 10 {
			// 修改
			item.Name = "修改" + modelTitle
			item.Url = "/" + modelName + "/update"
			item.Permission = "sys:" + modelName + ":update"
		} else if v == 15 {
			// 删除
			item.Name = "删除" + modelTitle
			item.Url = "/" + modelName + "/delete"
			item.Permission = "sys:" + modelName + ":delete"
		} else if v == 20 {
			// 详情
			item.Name = modelTitle + "详情"
			item.Url = "/" + modelName + "/detail"
			item.Permission = "sys:" + modelName + ":detail"
		} else if v == 25 {
			// 状态
			item.Name = "设置状态"
			item.Url = "/" + modelName + "/status"
			item.Permission = "sys:" + modelName + ":status"
		} else if v == 30 {
			// 批量删除
			item.Name = "批量删除"
			item.Url = "/" + modelName + "/dall"
			item.Permission = "sys:" + modelName + ":dall"
		}

		// 插入数据
		_, err := utils.XormDb.Insert(item)
		if err != nil {
			break
		}
	}
	return nil
}

// 生成路由文件
func GenerateRouter(dataList *common.ArrayList, authorName string, moduleName string, moduleTitle string) error {
	// 初始化表单数组
	columnList := make([]map[string]interface{}, 0)
	for i := 0; i < dataList.Size(); i++ {
		// 当前元素
		data := dataList.Get(i)
		// 类型转换
		item := data.(map[string]interface{})
		// 字段列名
		columnName := gconv.String(item["columnName"])
		// 移除部分非表单字段
		if columnName == "id" ||
			columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		// 加入数组
		columnList = append(columnList, item)
	}

	// 加载自定义模板绑定数据并写入文件
	if tmp, err := LoadTemplate("router.html", iris.Map{
		"author":      authorName,
		"since":       time.Now().Format("2006-01-02"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}, false); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 文件路径
		fileName := strings.Join([]string{curDir, "/router/", moduleName, ".go"}, "")
		// 删除现有文件
		if err := gfile.Remove(fileName); err != nil {
			return err
		}
		// 写入文件
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				// 写入文件
				f.WriteString(tmp)
			}
			// 关闭
			f.Close()
		}
	}
	return nil
}

// 获取表字段列表
func GetColumnList(tableName string) (*common.ArrayList, error) {
	// 获取数据库名
	DbName := conf.CONFIG.Mysql.Database
	// 获取字段列表
	data, err := utils.XormDb.SQL("SELECT COLUMN_NAME,COLUMN_DEFAULT,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT FROM information_schema.`COLUMNS` where TABLE_SCHEMA = ? AND TABLE_NAME = ?", DbName, tableName).Query()
	if err != nil {
		return nil, err
	}
	// 初始化数组
	result := common.New() //garray.NewArray(true)
	for _, v := range data {
		// 初始化Map
		item := make(map[string]interface{})
		// 字段列名
		columnName := gconv.String(v["COLUMN_NAME"])
		// 系统常规字段直接跳过
		if columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		item["columnName"] = columnName
		// 字段名称驼峰格式一
		columnName2 := gstr.UcWords(columnName)
		if gstr.Contains(columnName, "_") {
			nameArr := gstr.Split(columnName, "_")
			columnName2 = gstr.UcWords(nameArr[0]) + gstr.UcWords(nameArr[1])
		}
		item["columnName2"] = columnName2

		// 字段名称驼峰格式二
		columnName3 := columnName
		if gstr.Contains(columnName, "_") {
			nameArr := gstr.Split(columnName, "_")
			columnName3 = nameArr[0] + gstr.UcWords(nameArr[1])
		}
		item["columnName3"] = columnName3

		// 字段默认值
		item["columnDefault"] = v["COLUMN_DEFAULT"]
		// 字段数据类型
		dataType := gconv.String(v["DATA_TYPE"])
		item["dataType"] = dataType
		if dataType == "int" || dataType == "tinyint" || dataType == "smallint" {
			// 整形
			item["columnType"] = "int"
		} else if dataType == "bigint" {
			item["columnType"] = "int64"
		} else if dataType == "datetime" {
			item["columnType"] = "int"
		} else {
			// 字符串类型
			item["columnType"] = "string"
		}
		// 默认参数
		item["columnSwitch"] = false
		item["columnImage"] = false
		item["columnText"] = false
		item["columnValue"] = ""
		item["columnValueList"] = ""

		// 字段描述
		columnComment := gconv.String(v["COLUMN_COMMENT"])
		item["columnComment"] = columnComment
		// 判断是否有规则描述
		if gstr.Contains(columnComment, ":") || gstr.Contains(columnComment, "：") {
			// 正则根据冒号分裂字符串
			re := regexp.MustCompile("[：；]")
			commentItem := gstr.Split(re.ReplaceAllString(columnComment, "|"), "|")
			// 字段标题
			item["columnTitle"] = commentItem[0]

			// 字段描述数据处理
			commentStr := gstr.Replace(commentItem[1], " ", "|")
			commentArr := gstr.Split(commentStr, "|")

			// 实例化字段描述参数数组
			columnValue := make([]string, 0)
			// 参数值Map列表
			columnValueList := make(map[int]string)
			// 实例化字段描述文字数组
			columnSwitchValue := make([]string, 0)
			// 下拉选择列表解析参数
			columnSelectValue := make([]string, 0)
			for _, v := range commentArr {
				// 正则提取数字键
				regexp := regexp.MustCompile(`[0-9]+`)
				vItem := regexp.FindStringSubmatch(v)
				// 键
				key := vItem[0]
				// 值
				value := gstr.Replace(v, vItem[0], "")
				// 加入数组
				columnValue = append(columnValue, key+"="+value)
				// 参数值Map
				columnValueList[gconv.Int(key)] = value
				// 开关专用参数值
				columnSwitchValue = append(columnSwitchValue, value)
				// 下拉列表解析参数
				columnSelectValue = append(columnSelectValue, value)
			}
			// 字符串逗号拼接
			item["columnValue"] = gstr.Join(columnValue, ",")
			item["columnValueList"] = columnValueList
			item["columnSelectValue"] = gstr.Join(columnSelectValue, "','")

			// 开关判断
			if columnName == "status" || gstr.SubStr(columnName, 0, 3) == "is_" {
				item["columnSwitch"] = true
				item["columnSwitchValue"] = gstr.Join(columnSwitchValue, "|")
				// 方法名处理
				columnSwitchName := ""
				if gstr.Contains(columnName, "_") {
					switchArr := gstr.Split(columnName, "_")
					columnSwitchName = "set" + gstr.UcWords(switchArr[0]) + gstr.UcWords(switchArr[1])
				} else {
					columnSwitchName = "set" + gstr.UcWords(columnName)
				}
				item["columnSwitchName"] = columnSwitchName
			}
		} else {
			// 字段标题
			item["columnTitle"] = columnComment
		}

		// 判断是否是图片
		if gstr.Contains(columnName, "cover") ||
			gstr.Contains(columnName, "avatar") ||
			gstr.Contains(columnName, "image") ||
			gstr.Contains(columnName, "logo") ||
			gstr.Contains(columnName, "pic") {
			item["columnImage"] = true
		}

		// 判断是否多行文本或富文本
		if gstr.Contains(columnName, "note") ||
			gstr.Contains(columnName, "remark") ||
			gstr.Contains(columnName, "content") ||
			gstr.Contains(columnName, "description") ||
			gstr.Contains(columnName, "intro") {
			item["columnText"] = true
		}

		// 加入数组
		result.Add(item)
	}
	return result, nil
}

//读取模板
func LoadTemplate(templateName string, data interface{}, isReplace bool) (string, error) {
	// 获取当前应用根目录
	curDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// 读取模板文件及内容
	b, err := ioutil.ReadFile(curDir + "/templates/" + templateName)
	if err != nil {
		return "", err
	}
	// 创建一个模板
	//tmpl := template.New(templateName)
	//tmpl = tmpl.Funcs(template.FuncMap{
	//	"safe": func(str string) template.HTML {
	//		return template.HTML(str)
	//	},
	//})
	//tmpl, err = tmpl.Parse(string(b))
	tmpl, err := template.New(templateName).Funcs(
		template.FuncMap{
			"safe": func(str string) template.HTML {
				return template.HTML(str)
			},
		},
	).Parse(string(b))
	if err != nil {
		return "", nil
	}
	buffer := bytes.NewBufferString("")
	// 将string与模板合成，变量name的内容会替换掉{{.}}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", nil
	}
	if isReplace {
		// 替换script标签
		result := strings.Replace(buffer.String(), "checkedStart", "'+(d.", -1)
		result = strings.Replace(result, "checkedEnd", "==1 ? 'checked' : '')+'", -1)
		return result, nil
	} else {
		return buffer.String(), nil
	}
}
