package widget

import (
	"html/template"
	"strings"
)

func Query(s string) template.HTML {
	str := "<button class=\"layui-btn\" lay-submit=\"\" lay-filter=\"searchForm\" id=\"search\"><i class=\"layui-icon\">\uE615</i>" + s + "</button>"
	return unescapeHTML(str)
}

func Select(s string, list string, level int) template.HTML {
	str := "<select name=\"city\" lay-verify=\"\">\n  <option value=\"\">请选择一个城市</option>\n  <option value=\"010\">北京</option>\n  <option value=\"021\">上海</option>\n  <option value=\"0571\">杭州</option>\n</select>"
	return unescapeHTML(str)
}

func Submit(s string, t int, o string) template.HTML {
	//{{submit "submit|立即保存,close|关闭" 1 ""}}
	str := ""
	str += "<div class=\"layui-form-item text-center model-form-footer\">"

	tarr := strings.Split(s, ",")
	for _, v := range tarr {
		narr := strings.Split(v, "|")
		switch narr[0] {
		case "submit":
			str += "<button class=\"layui-btn\" lay-filter=\"submitForm\" lay-submit=\"\">" + narr[1] + "</button>"
		case "close":
			str += "<button class=\"layui-btn layui-btn-primary\" type=\"button\" ew-event=\"closeDialog\">" + narr[1] + "</button>"
		}
	}

	str += "</div>"
	return unescapeHTML(str)
}

func Add(s string, o string) template.HTML {
	str := "<a href=\"javascript:\" class=\"layui-btn btnOption  layui-btn-small btnadd\" id=\"add\" data-param=\"" + o + "\" lay-event=\"add\"><i class=\"layui-icon layui-icon-add-1\"></i> " + s + "</a>"
	return unescapeHTML(str)
}

func Expand(s string) template.HTML {
	return ""
}

func Collapse(s string) template.HTML {
	return ""
}

func Addz(s string) template.HTML {
	return ""
}

func Edit(s string) template.HTML {
	str := "<a class=\"layui-btn layui-btn-xs btnEdit\" lay-event=\"edit\" title=\"编辑\"><i class=\"layui-icon\">\uE642</i>" + s + "</a>"
	return unescapeHTML(str)
}

func Delete(s string) template.HTML {
	str := "<a class=\"layui-btn layui-btn-danger layui-btn-xs btnDel\" lay-event=\"del\" title=\"删除\"><i class=\"layui-icon\">\uE640</i>" + s + "</a>"
	return unescapeHTML(str)
}

func Widget(id string, icon string, name string, btn string, n int, o string) template.HTML {
	//{{widget "resetPwd" "layui-icon-password" "重置密码" "layui-btn-warm" 2 ""}}
	str := "<a href=\"javascript:\" class=\"layui-btn btnOption " + btn + " layui-btn-xs btnresetPwd\" id=\"" + id + "\" data-param=\"{}\" lay-event=\"" + id + "\"><i class=\"layui-icon " + icon + "\"></i> " + name + "</a>"
	return unescapeHTML(str)
}

// 批量删除
func Dall(s string) template.HTML {
	str := "<a href=\"javascript:\" class=\"layui-btn btnOption layui-btn-danger layui-btn-small btndall\" id=\"dall\" data-param=\"{}\" lay-event=\"dall\"><i class=\"layui-icon layui-icon-delete\"></i> " + s + "</a>"
	return unescapeHTML(str)
}

func Switch(s string, n string, val int) template.HTML {
	//{{switch "status" "在用|禁用" .info.Status}}
	return ""
}

func Safe(s string) template.HTML {
	return ""
}

func Date(s string, b string) template.HTML {
	//{{date "birthday|1|出生日期|date" .info.Birthday}}
	return ""
}

func Icon(s string) template.HTML {
	return ""
}

func Transfer(s string) template.HTML {
	return ""
}

func UploadImage(s string, avatar string, o string, t int) template.HTML {
	//{{upload_image "avatar|头像|90x90|建议上传尺寸450x450|450x450" .info.Avatar "" 0}}
	return ""
}

func Album(s string) template.HTML {
	return ""
}

func Item(s string) template.HTML {
	return ""
}

func Kindeditor(s string) template.HTML {
	return ""
}

func Checkbox(s string, data map[string]interface{}, r []int) template.HTML {
	//{{checkbox "roleIds|name|id" .roleList .info.RoleIds}}
	return ""
}

func Radio(s string) template.HTML {
	return ""
}

func City(c int, a int, b int) template.HTML {
	//{{city .info.DistrictCode 3 1}}
	return ""
}

// 转义html
func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}
