package model

import (
	"easygoadmin/utils"
	"time"
)

type ItemList struct {
	Id  int   `json:"id"`
	Num int64 `json:"num"`
}

type SendMail struct {
	Id         int        `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Title      string     `json:"title" xorm:"not null comment('标题') VARCHAR(100)"`
	Content    string     `json:"content" xorm:"not null comment('内容') TEXT"`
	Type       int        `json:"type" xorm:"not null default 1 comment('类型') TINYINT(1)"`
	RoleList   string     `json:"role_list" xorm:"not null comment('角色ID列表') TEXT"`
	ItemList   []ItemList `json:"item_list" xorm:"not null comment('道具列表') TEXT"`
	Status     int        `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	ReturnInfo string     `json:"return_info" xorm:"not null comment('接口返回信息') TEXT"`
	CreateUser int        `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time  `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
}

// 根据条件查询单条数据
func (r *SendMail) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *SendMail) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *SendMail) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *SendMail) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&SendMail{})
}

//批量删除
func (r *SendMail) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&SendMail{})
}

// 定义message struct
type GmMail struct {
	Title   string         `json:"title"`   // 标题
	Content string         `json:"content"` // 内容
	Items   []GmMailItem   `json:"items"`   // 道具列表
	Target  []GmMailTarget `json:"target"`  // 目标用户
	Other   interface{}    `json:"other"`   // 其他参数
}

type GmMailItem struct {
	Id  int32 `json:"id"`  // Id
	Num int64 `json:"num"` // 数量
}

// target参数目前没用，等邮件系统出来
type GmMailTarget struct {
	Type string      `json:"type"` // 类型 例：expire_time  list_role
	Args interface{} `json:"args"` // 参数 例：1667870797   xxxxxxxxx
}
