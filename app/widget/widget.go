package widget

import (
	"easygoadmin/utils"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Query(s string) template.HTML {
	str := `<button class="layui-btn" lay-submit="" lay-filter="searchForm" id="search"><i class="layui-icon"></i>` + s + `</button>`
	return unescapeHTML(str)
}

func Select(s string, data map[int]string, val int) template.HTML {
	// {{select "gender|1|性别|name|id" "1=男,2=女,3=保密" .info.Gender}}
	// {{select "gender|1|性别|name|id" .genderList .info.Gender}}
	//str := `<select name="city" lay-verify="">
	//		  <option value="010">北京</option>
	//		  <option value="021" disabled>上海（禁用效果）</option>
	//		  <option value="0571" selected>杭州</option>
	//		</select>`

	// 解析s
	tarr := strings.Split(s, "|")
	n := tarr[0]
	//ty := tarr[1]
	de := tarr[2]
	//name := tarr[3]
	//id := tarr[4]

	dkey := utils.GetMapIndexOrderInt(data)

	str := `<select name="` + n + `" lay-verify="">`
	str += `<option value="">【请选择` + de + `】</option>`
	// 循环数据
	var selected string
	for _, id := range dkey {
		name := data[id]
		selected = ""
		if id == val {
			selected = "selected"
		}
		str += `<option value="` + strconv.FormatInt(int64(id), 10) + `" ` + selected + `>` + name + `</option>`
	}

	str += `</select>`

	return unescapeHTML(str)
}

func Submit(s string, t int, o string) template.HTML {
	//{{submit "submit|立即保存,close|关闭" 1 ""}}

	if o == "" {
		o = "submitForm"
	}

	str := ""
	str += `<div class="layui-form-item text-center model-form-footer">`

	tarr := strings.Split(s, ",")
	for _, v := range tarr {
		narr := strings.Split(v, "|")
		switch narr[0] {
		case "submit":
			str += `<button class="layui-btn preservation" lay-filter="` + o + `" lay-submit="">` + narr[1] + `</button>`
		case "close":
			str += `<button class="layui-btn layui-btn-primary" type="button" ew-event="closeDialog">` + narr[1] + `</button>`
		}
	}

	str += `</div>`
	return unescapeHTML(str)
}

func Add(s string, o string) template.HTML {
	str := `<a href="javascript:" class="layui-btn btnOption  layui-btn-small btnadd" id="add" data-param="` + o + `" lay-event="add"><i class="layui-icon layui-icon-add-1"></i> ` + s + `</a>`
	return unescapeHTML(str)
}

func Expand(s string) template.HTML {
	// {{expand "全部展开"}}
	str := `<a href="javascript:" class="layui-btn btnOption layui-btn-normal layui-btn-small btnexpand" id="expand" data-param="{}" lay-event="expand"><i class="layui-icon layui-icon-shrink-right"></i> ` + s + `</a>`
	return unescapeHTML(str)
}

func Collapse(s string) template.HTML {
	// {{collapse "全部折叠"}}
	str := `<a href="javascript:" class="layui-btn btnOption layui-btn-warm layui-btn-small btncollapse" id="collapse" data-param="{}" lay-event="collapse"><i class="layui-icon layui-icon-spread-left"></i> ` + s + `</a>`
	return unescapeHTML(str)
}

func Addz(s string) template.HTML {
	str := `<a href="javascript:" class="layui-btn btnOption layui-btn-normal layui-btn-xs btnaddz" id="addz" data-param="{}" lay-event="addz"><i class="layui-icon layui-icon-add-1"></i> ` + s + `</a>`
	return unescapeHTML(str)
}

func Edit(s string) template.HTML {
	str := `<a class="layui-btn layui-btn-xs btnEdit" lay-event="edit" title="编辑"><i class="layui-icon">` + "\uE642" + `</i>` + s + `</a>`
	return unescapeHTML(str)
}

func Delete(s string) template.HTML {
	str := `<a class="layui-btn layui-btn-danger layui-btn-xs btnDel" lay-event="del" title="删除"><i class="layui-icon">` + "\uE640" + `</i>` + s + `</a>`
	return unescapeHTML(str)
}

func Widget(id string, icon string, name string, btn string, n int, o string) template.HTML {
	// {{widget "resetPwd" "layui-icon-password" "重置密码" "layui-btn-warm" 2 ""}}
	// {{widget "permission" "layui-icon-set" "角色权限" "layui-bg-cyan" 2 ""}}
	str := `<a href="javascript:" class="layui-btn btnOption ` + btn + ` layui-btn-xs btnresetPwd" id="` + id + `" data-param="{}" lay-event="` + id + `"><i class="layui-icon ` + icon + `"></i> ` + name + `</a>`
	return unescapeHTML(str)
}

// 批量删除
func Dall(s string) template.HTML {
	str := `<a href="javascript:" class="layui-btn btnOption layui-btn-danger layui-btn-small btndall" id="dall" data-param="{}" lay-event="dall"><i class="layui-icon layui-icon-delete"></i> ` + s + `</a>`
	return unescapeHTML(str)
}

func Switch(s string, n string, val interface{}) template.HTML {
	//{{switch "status" "在用|禁用" .info.Status}}
	checked := "checked"
	if !reflect.DeepEqual(val, 1) {
		checked = ""
	}
	str := `<input type="checkbox" name="` + s + `" lay-skin="switch" lay-text="` + n + `" ` + checked + `>`
	return unescapeHTML(str)
}

func Safe(s string) template.HTML {
	return ""
}

func Date(s string, v time.Time) template.HTML {
	//{{date "birthday|1|出生日期|date" .info.Birthday}}

	tarr := strings.Split(s, "|")
	name := tarr[0]
	//ty := tarr[1]
	de := tarr[2]
	ve := tarr[3]

	// 判断值是否等于0，转成time格式
	val := "2000-11-11"
	if !v.IsZero() {
		val = v.String()
	}

	str := `<input type="text" class="layui-input" value="` + val + `" lay-verify="` + ve + `" placeholder="请选择` + de + `" name="` + name + `" id="` + name + `">`
	return unescapeHTML(str)
}

func Icon(s string) template.HTML {
	return ""
}

func Transfer(s string, o string, f []interface{}) template.HTML {
	// {{transfer "func|0|全部节点,已赋予节点|name|id|220x350" "1=列表,5=添加,10=修改,15=删除,20=详情,25=状态,30=批量删除,35=添加子级,40=全部展开,45=全部折叠,50=导出数据,55=导入数据,60=分配权限,65=重置密码" .funcList}}
	str := ""
	return unescapeHTML(str)
}

func UploadImage(s string, avatar string, o string, t int) template.HTML {
	//{{upload_image "avatar|头像|90x90|建议上传尺寸450x450|450x450" .info.Avatar "" 0}}
	str := `<div class="layui-input-block">
			<div class="layui-upload-drag"><img id="avatar_show_id" src="/static/assets/images/default_upload.png" alt="上传头像" width="90" height="90"><input type="hidden" id="avatar" name="avatar" value="">
			</div>
			<div style="margin-top:10px;">
				<button type="button" class="layui-btn" id="upload_avatar"><i class="layui-icon"></i>上传头像</button>
			</div><div class="layui-form-mid layui-word-aux">建议尺寸：建议上传尺寸450x450</div></div>`
	return unescapeHTML(str)
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

func Checkbox(s string, data map[int]string, r []interface{}) template.HTML {
	//{{checkbox "roleIds|name|id" .roleList .info.RoleIds}}
	str := ``

	tarr := strings.Split(s, "|")
	name := tarr[0]

	// 取得data的key排序
	dkey := utils.GetMapIndexOrderInt(data)
	// 解析r
	var rids []int
	if r != nil {
		for _, v := range r {
			vi, _ := v.(int)
			rids = append(rids, vi)
		}
	}

	for _, k := range dkey {
		v := data[k]
		checked := ""
		kstr := strconv.FormatInt(int64(k), 10)
		if utils.InIntArray(k, rids) {
			checked = "checked"
		}
		str += `<input type="checkbox" value="` + kstr + `" name="` + name + `" title="` + v + `" lay-skin="primary" ` + checked + `>`
	}

	return unescapeHTML(str)
}

func Radio(s string) template.HTML {
	return ""
}

func City(s string, a int, b int) template.HTML {
	//{{city .info.DistrictCode 3 1}}
	return ""
}

// 转义html
func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}
