package model

import (
	"easygoadmin/utils"
	"time"
)

type SetTime struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Time       time.Time `json:"time" xorm:"not null comment('时间') DATETIME"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
}

// 根据条件查询单条数据
func (r *SetTime) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *SetTime) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *SetTime) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *SetTime) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Role{})
}

//批量删除
func (r *SetTime) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Role{})
}
