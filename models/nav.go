package models

type Nav struct {
	Id int
	Name string
	Pid int
}

func (m *Nav) TableName() string {
	return TableName("nav")
}