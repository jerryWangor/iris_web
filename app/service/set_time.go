package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"errors"
)

var SetTime = new(SetTimeService)

type SetTimeService struct{}

func (s *SetTimeService) GetList() ([]model.SetTime, error) {
	// 创建查询实例
	query := utils.XormDb.Where("1=1")
	query = query.OrderBy("id desc")
	// 查询列表
	var list []model.SetTime
	err := query.Find(&list)
	return list, err
}

func (s *SetTimeService) Add(req dto.SetTimeReq, userId int) (int64, error) {

	// 实例化对象
	var entity model.SetTime
	entity.Time = utils.TimeStringToTimeTime(req.Time)
	entity.CreateUser = userId
	entity.CreateTime = utils.GetNowTimeTime()

	// 插入数据
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	return rows, nil
}
