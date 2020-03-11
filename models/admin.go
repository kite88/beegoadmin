package models

import (
	"beegoadmin/utils"
	"beegoadmin/utils/page"
	"fmt"
	"time"
)

type Admin struct {
	Id         int `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	Username   string
	Password   string
	HeadPic    string
	RoleId     int

	RoleName         string `gorm:"-"`                   // 忽略这个字段
	CreateTimeFormat string `gorm:"-" json:"CreateTime"` // 忽略这个字段
	UpdateTimeFormat string `gorm:"-" json:"UpdateTime"` // 忽略这个字段
}

//查找用户信息
func FindAdminByUserInfo(UserInfo Admin) (Admin, error) {
	var admin Admin
	var role Role

	adminTable := admin.TableName()
	roleTable := role.TableName()

	adminField := fmt.Sprintf("%s.id,%s.username,%s.role_id,%s.head_pic,", adminTable, adminTable, adminTable, adminTable)
	roleField := fmt.Sprintf("%s.name as role_name", roleTable)
	join := fmt.Sprintf("left join %s on %s.id = %s.role_id", roleTable, roleTable, adminTable)

	var result Admin
	err := Db.Select(adminField + roleField).Where(UserInfo).Joins(join).First(&result).Error
	return result, err
}

//根据条数查询条数
func FindAdminCount(query interface{}, args ...interface{}) int {
	var count int
	Db.Model(&Admin{}).Where(query, args...).Count(&count)
	return count
}

//新增数据
func AddAdmin(value *Admin) error {
	value.CreateTime = time.Now().Unix()
	err := Db.Create(value).Error
	return err
}

//更新数据
func UpdateAdmin(value map[string]interface{}, query interface{}, args ...interface{}) error {
	value["update_time"] = time.Now().Unix()
	err := Db.Model(&Admin{}).Where(query, args...).Updates(value).Error
	return err
}

//删除
func DelAdmin(query interface{}, args ...interface{}) int64 {
	n := Db.Where(query, args...).Delete(Admin{}).RowsAffected
	return n
}

//获取数据
func GetAdminData(viewRows uint, p uint) page.Data {
	var count uint
	var admin Admin
	var role Role
	var list []Admin

	adminTable := admin.TableName()
	roleTable := role.TableName()

	adminField := fmt.Sprintf("%s.id,%s.username,%s.role_id,%s.create_time,%s.update_time,", adminTable, adminTable, adminTable, adminTable, adminTable)
	roleField := fmt.Sprintf("%s.name as role_name", roleTable)
	join := fmt.Sprintf("left join %s on %s.id = %s.role_id", roleTable, roleTable, adminTable)

	Db.Find(&list).Count(&count)
	data, limit, offset := page.Page(count, viewRows, p)

	Db.Model(&admin).
		Select(adminField + roleField).
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Joins(join).
		Find(&list)

	for key, val := range list {
		list[key].CreateTimeFormat = utils.FormatDateTime(val.CreateTime,"")
		list[key].UpdateTimeFormat = utils.FormatDateTime(val.UpdateTime,"")
	}
	data.List = list
	return data
}

func (m *Admin) TableName() string {
	return TableName("admin")
}
