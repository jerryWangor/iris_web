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
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gstr"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var Menu = new(menuService)

type menuService struct{}

func (s *menuService) GetList(req dto.MenuPageReq) ([]model.Menu, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 菜单名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询列表
	var list []model.Menu
	err := query.Find(&list)
	return list, err
}

func (s *menuService) Add(req dto.MenuAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Menu
	entity.Pid = gconv.Int(req.Pid)
	entity.Name = req.Name
	entity.Icon = req.Icon
	entity.Url = req.Url
	entity.Target = gconv.Int(req.Target)
	entity.Permission = req.Permission
	entity.Type = gconv.Int(req.Type)
	if req.Status == "on" {
		entity.Status = 1
	} else {
		entity.Status = 2
	}
	if req.Hide == "on" {
		entity.Hide = 1
	} else {
		entity.Hide = 2
	}
	entity.Note = req.Note
	entity.Sort = gconv.Int(req.Sort)
	entity.CreateUser = userId
	entity.CreateTime = utils.GetNowTimeTime()
	entity.UpdateUser = userId
	entity.UpdateTime = utils.GetNowTimeTime()
	entity.Mark = 1
	// 插入数据
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	// 添加节点
	setPermission(entity.Type, req.Func, req.Name, req.Url, gconv.Int(entity.Id), userId)
	return rows, nil
}

func (s *menuService) Update(req dto.MenuUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.Menu{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	entity.Pid = gconv.Int(req.Pid)
	entity.Name = req.Name
	entity.Icon = req.Icon
	entity.Url = req.Url
	entity.Target = gconv.Int(req.Target)
	entity.Permission = req.Permission
	entity.Type = gconv.Int(req.Type)
	if req.Status == "on" {
		entity.Status = 1
	} else {
		entity.Status = 2
	}
	if req.Hide == "on" {
		entity.Hide = 1
	} else {
		entity.Hide = 2
	}
	entity.Note = req.Note
	entity.Sort = gconv.Int(req.Sort)
	entity.UpdateUser = userId
	entity.UpdateTime = utils.GetNowTimeTime()
	// 更新数据
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, errors.New("更新失败")
	}

	// 添加节点
	setPermission(entity.Type, req.Func, req.Name, req.Url, entity.Id, userId)
	return rows, nil
}

func (s *menuService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Menu{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, err
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &model.Menu{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

// 添加节点
func setPermission(menuType int, funcIds string, name string, url string, parentId int, userId int) {
	if menuType != 0 || funcIds == "" || url == "" {
		return
	}
	// 删除现有节点
	utils.XormDb.Where("pid=?", parentId).Delete(&model.Menu{})
	// 模块名称
	moduleTitle := gstr.Replace(name, "管理", "")
	// 创建权限节点
	urlArr := strings.Split(url, "/")

	if len(urlArr) >= 3 {
		// 模块名
		moduleName := urlArr[len(urlArr)-2]
		// 节点处理
		checkedList := strings.Split(funcIds, ",")
		for _, v := range checkedList {
			// 实例化对象
			var entity model.Menu
			// 节点索引
			value := gconv.Int(v)
			if value == 1 {
				entity.Name = "查询" + moduleTitle
				entity.Url = "/" + moduleName + "/list"
				entity.Permission = "sys:" + moduleName + ":list"
			} else if value == 5 {
				entity.Name = "添加" + moduleTitle
				entity.Url = "/" + moduleName + "/add"
				entity.Permission = "sys:" + moduleName + ":add"
			} else if value == 10 {
				entity.Name = "修改" + moduleTitle
				entity.Url = "/" + moduleName + "/update"
				entity.Permission = "sys:" + moduleName + ":update"
			} else if value == 15 {
				entity.Name = "删除" + moduleTitle
				entity.Url = "/" + moduleName + "/delete"
				entity.Permission = "sys:" + moduleName + ":delete"
			} else if value == 20 {
				entity.Name = moduleTitle + "详情"
				entity.Url = "/" + moduleName + "/detail"
				entity.Permission = "sys:" + moduleName + ":detail"
			} else if value == 25 {
				entity.Name = "设置状态"
				entity.Url = "/" + moduleName + "/status"
				entity.Permission = "sys:" + moduleName + ":status"
			} else if value == 30 {
				entity.Name = "批量删除"
				entity.Url = "/" + moduleName + "/dall"
				entity.Permission = "sys:" + moduleName + ":dall"
			} else if value == 35 {
				entity.Name = "添加子级"
				entity.Url = "/" + moduleName + "/addz"
				entity.Permission = "sys:" + moduleName + ":addz"
			} else if value == 40 {
				entity.Name = "全部展开"
				entity.Url = "/" + moduleName + "/expand"
				entity.Permission = "sys:" + moduleName + ":expand"
			} else if value == 45 {
				entity.Name = "全部折叠"
				entity.Url = "/" + moduleName + "/collapse"
				entity.Permission = "sys:" + moduleName + ":collapse"
			} else if value == 50 {
				entity.Name = "导出" + moduleTitle
				entity.Url = "/" + moduleName + "/export"
				entity.Permission = "sys:" + moduleName + ":export"
			} else if value == 55 {
				entity.Name = "导入" + moduleTitle
				entity.Url = "/" + moduleName + "/import"
				entity.Permission = "sys:" + moduleName + ":import"
			} else if value == 60 {
				entity.Name = "分配权限"
				entity.Url = "/" + moduleName + "/permission"
				entity.Permission = "sys:" + moduleName + ":permission"
			} else if value == 65 {
				entity.Name = "重置密码"
				entity.Url = "/" + moduleName + "/resetPwd"
				entity.Permission = "sys:" + moduleName + ":resetPwd"
			}
			entity.Pid = parentId
			entity.Type = 1
			entity.Status = 1
			entity.Target = 1
			entity.Sort = value
			entity.CreateUser = userId
			entity.CreateTime = utils.GetNowTimeTime()
			entity.UpdateUser = userId
			entity.UpdateTime = utils.GetNowTimeTime()
			entity.Mark = 1

			// 插入节点
			entity.Insert()
		}
	}
}

// 获取菜单权限列表
func (s *menuService) GetPermissionList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		menuList, _ := Menu.GetTreeList()
		return menuList
	} else {
		// 非管理员

		// 数据转换
		list := make([]model.Menu, 0)
		// 查询数据
		utils.XormDb.Table("sys_menu").Alias("m").
			Join("INNER", []string{"sys_role_menu", "r"}, "m.id = r.menu_id").
			Join("INNER", []string{"sys_user_role", "ur"}, "ur.role_id=r.role_id").
			Where("ur.user_id=? AND m.type=0 AND m.`status`=1 AND m.mark=1", userId).
			Cols("m.*").
			OrderBy("m.id asc").
			Find(&list)

		// 数据处理
		var menuNode vo.MenuTreeNode
		makeTree(list, &menuNode)
		return menuNode.Children
	}
}

// 获取子级菜单
func (s *menuService) GetTreeList() ([]*vo.MenuTreeNode, error) {
	var menuNode vo.MenuTreeNode
	list := make([]model.Menu, 0)
	err := utils.XormDb.Where("type=0 and mark=1").OrderBy("sort").Find(&list)
	if err != nil {
		return nil, err
	}
	makeTree(list, &menuNode)
	return menuNode.Children, nil
}

//递归生成分类列表
func makeTree(menu []model.Menu, tn *vo.MenuTreeNode) {
	for _, c := range menu {
		if c.Pid == tn.Id {
			child := &vo.MenuTreeNode{}
			child.Menu = c
			tn.Children = append(tn.Children, child)
			makeTree(menu, child)
		}
	}
}

// 获取权限节点列表
func (s *menuService) GetPermissionsList(userId int) []string {
	if userId == 1 {
		// 管理员,管理员拥有全部权限
		permissionList := make([]string, 0)
		utils.XormDb.Table("sys_menu").Cols("permission").Where("type=1").Where("mark=1").Find(&permissionList)
		return permissionList
	} else {
		// 非管理员
		permissionList := make([]string, 0)
		utils.XormDb.Table("sys_menu").Alias("m").
			Join("INNER", []string{"sys_role_menu", "r"}, "m.id = r.menu_id").
			Join("INNER", []string{"sys_user_role", "ur"}, "ur.role_id=r.role_id").
			Where("ur.user_id=? AND m.type=1 AND m.`status`=1 AND m.mark=1", userId).
			Cols("m.permission").
			Find(&permissionList)
		return permissionList
	}
}

// 数据源转换
func (s *menuService) MakeList(data []*vo.MenuTreeNode) map[int]string {
	menuList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			menuList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				menuList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					menuList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return menuList
}

// 获取菜单权限列表（Tree）
func (s *menuService) GetPermissionMenuTreeList(userId int) interface{} {
	if utils.InIntArray(userId, conf.SUPER_ADMIN_USER_IDS) {
		// 管理员(拥有全部权限)
		menuList, _ := Menu.GetTreeList()
		//fmt.Println("菜单列表", menuList)
		return menuList
	} else {
		// 非管理员

		// 数据转换
		list := make([]model.Menu, 0)
		// 查询数据
		//utils.XormDb.Table("sys_menu").Alias("m").
		//	Join("INNER", []string{"sys_role_menu", "r"}, "m.id = r.menu_id").
		//	Join("INNER", []string{"sys_user_role", "ur"}, "ur.role_id=r.role_id").
		//	Where("ur.user_id=? AND m.type=0 AND m.`status`=1 AND m.mark=1", userId).
		//	Cols("m.*").
		//	OrderBy("m.id asc").
		//	Find(&list)

		// 使用原生sql
		rmList := make([]model.RoleMenu, 0)
		utils.XormDb.Table("sys_role_menu").Alias("rm").
			Join("INNER", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
			Where("ur.user_id=?", userId).
			Cols("rm.*").
			OrderBy("rm.id asc").
			Find(&rmList)

		ids := make([]string, 0, 0)
		for _, v := range rmList {
			tids := strings.Split(v.MenuIds, ",")
			for _, id := range tids {
				ids = append(ids, id)
			}
		}
		ids = utils.RemoveDuplicatesAndEmpty(ids)
		idstr := strings.Join(ids, ",")
		// 再查询所有菜单
		utils.XormDb.Where("type=0 AND `status`=1 AND mark=1 and id in (" + idstr + ")").OrderBy("id asc").Find(&list)
		//fmt.Println(list)

		// 数据处理
		var menuNode vo.MenuTreeNode
		makeTree(list, &menuNode)
		return menuNode.Children
	}
}

// 获取菜单权限列表
func (s *menuService) GetPermissionMenuList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		list := make([]model.Menu, 0)
		err := utils.XormDb.Where("mark=1").OrderBy("sort").Find(&list)
		if err != nil {
			return nil
		}
		return list
	} else {
		// 非管理员
		list := make([]model.Menu, 0)
		// 使用原生sql
		rmList := make([]model.RoleMenu, 0)
		utils.XormDb.Table("sys_role_menu").Alias("rm").
			Join("INNER", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
			Where("ur.user_id=?", userId).
			Cols("rm.*").
			OrderBy("rm.id asc").
			Find(&rmList)

		ids := make([]string, 0, 0)
		for _, v := range rmList {
			tids := strings.Split(v.MenuIds, ",")
			for _, id := range tids {
				ids = append(ids, id)
			}
		}
		ids = utils.RemoveDuplicatesAndEmpty(ids)
		idstr := strings.Join(ids, ",")
		// 再查询所有菜单
		utils.XormDb.Where("`status`=1 AND mark=1 and id in (" + idstr + ")").OrderBy("id asc").Find(&list)
		return list
	}
}
