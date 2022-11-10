package model

import (
	"easygoadmin/utils"
	"time"
)

type Item struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	TypeId     int       `json:"type_id" xorm:"not null default 0 comment('道具模板ID') index INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('道具名字') index VARCHAR(100)"`
	Icon       string    `json:"icon" xorm:"default 'NULL' comment('图标') VARCHAR(50)"`
	Type       int       `json:"type" xorm:"not null default 1 comment('类型：1英雄 2道具') TINYINT(1)"`
	CreateUser int       `json:"create_user" xorm:"not null default 1 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 1 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
}

// 根据条件查询单条数据
func (r *Item) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Item) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Item) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Item) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Item{})
}

//批量删除
func (r *Item) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Item{})
}
