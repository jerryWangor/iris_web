package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"encoding/json"
	"errors"
)

var SendMail = new(sendMailService)

type sendMailService struct{}

func (s *sendMailService) GetList() ([]model.SendMail, error) {
	// 创建查询实例
	query := utils.XormDb.Where("1=1")
	query = query.OrderBy("id desc")
	// 查询列表
	var list []model.SendMail
	err := query.Find(&list)
	return list, err
}

func (s *sendMailService) Add(req dto.SendMailReq, userId int) (int, error) {

	// 实例化对象
	var entity model.SendMail
	entity.Title = req.Title
	entity.Content = req.Content
	entity.Type = req.Type
	entity.RoleList = req.RoleList
	// itemlist转成切片
	var itemlist []model.ItemList
	err := json.Unmarshal([]byte(req.ItemList), &itemlist)
	if err != nil {
		return 0, err
	}
	entity.ItemList = itemlist
	entity.CreateUser = userId
	entity.CreateTime = utils.GetNowTimeTime()

	// 插入数据
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	return entity.Id, nil
}

func (s *sendMailService) Update(req dto.SendMailReq) (int64, error) {

	// 实例化对象
	var entity model.SendMail
	entity.Id = req.Id
	entity.Title = req.Title
	entity.Content = req.Content
	entity.Type = req.Type
	entity.RoleList = req.RoleList
	// itemlist转成切片
	var itemlist []model.ItemList
	err := json.Unmarshal([]byte(req.ItemList), &itemlist)
	if err != nil {
		return 0, err
	}
	entity.ItemList = itemlist
	entity.Status = req.Status
	entity.ReturnInfo = req.ReturnInfo

	// 插入数据
	return entity.Update()
}
