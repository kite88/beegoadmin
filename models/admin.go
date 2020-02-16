package models

type Admin struct {
	Id int
	Username string
	Password string
}

func (m *Admin) TableName() string {
	return TableName("admin")
}