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
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strconv"
	"strings"
)

var Item = new(itemService)

type itemService struct{}

func (s *itemService) GetList(req dto.ItemPageReq) ([]model.Item, error) {
	// 创建查询实例
	query := utils.XormDb.Where("1=1")
	// 道具ID
	if req.TypeId != 0 {
		query = query.Where("type_id = ?", req.TypeId)
	}
	// 道具名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.OrderBy("type_id")
	// 查询列表
	var list []model.Item
	err := query.Find(&list)
	return list, err
}

func (s *itemService) Add(req dto.ItemAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Item
	entity.TypeId = gconv.Int(req.TypeId)
	entity.Name = req.Name
	entity.Icon = req.Icon
	entity.Type = gconv.Int(req.Type)
	entity.CreateUser = userId
	entity.CreateTime = utils.GetNowTimeTime()
	entity.UpdateUser = userId
	entity.UpdateTime = utils.GetNowTimeTime()
	// 插入数据
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	return rows, nil
}

func (s *itemService) Update(req dto.ItemUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Item{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	entity.TypeId = gconv.Int(req.TypeId)
	entity.Name = req.Name
	entity.Icon = req.Icon
	entity.Type = gconv.Int(req.Type)
	entity.UpdateUser = userId
	entity.UpdateTime = utils.GetNowTimeTime()
	// 更新数据
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

func (s *itemService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Item{Id: gconv.Int(ids)}
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
			entity := &model.Item{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}
