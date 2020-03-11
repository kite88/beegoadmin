package models

import "beegoadmin/utils/page"

type Role struct {
	Id     uint
	Name   string
	IsRoot int
}

//创建数据
func AddRole(value *Role) error {
	err := Db.Create(value).Error
	return err
}

//更新数据
func UpdateRole(value map[string]interface{}, query interface{}, args ...interface{}) error {
	err := Db.Model(&Role{}).Where(query, args...).Updates(value).Error
	return err
}

//获取数据
func GetRoleData(viewRows uint, p uint, query interface{}, args ...interface{}) page.Data {
	var count uint
	var list []Role
	Db.Where(query, args...).Find(&list).Count(&count)
	data, limit, offset := page.Page(count, viewRows, p)
	Db.Where(query, args...).Limit(limit).Offset(offset).Order("id desc").Find(&list)
	data.List = list
	return data
}

//获取单条数据
func FindRoleDataByRole(role Role) (Role, error) {
	var result Role
	err := Db.Where(role).First(&result).Error
	return result, err
}

//删除
func DelRole(query interface{}, args ...interface{}) int64 {
	n := Db.Where(query, args...).Delete(Role{}).RowsAffected
	return n
}

func (m *Role) TableName() string {
	return TableName("role")
}
