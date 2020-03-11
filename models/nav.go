package models

type Nav struct {
	Id     int
	Name   string
	Router string
	Pid    int
	Sort   int
}

type NavTree struct {
	Id       int
	Name     string
	Pid      int
	Router   string
	Sort     int
	Children []*NavTree
}

//添加记录
func AddNav(value interface{}) error {
	err := Db.Create(value).Error
	return err
}

//根据条数查询条数
func FindNavCount(query interface{}, args ...interface{}) int {
	var count int
	Db.Model(&Nav{}).Where(query, args...).Count(&count)
	return count
}

//根据条件查询数据
func FindNavByWhere(query interface{}, args ...interface{}) ([]Nav, error) {
	var list []Nav
	err := Db.Where(query, args...).Order("sort asc").Find(&list).Error
	return list, err
}

//使用主键获取记录
func FindNavFirst(id int) (Nav, error) {
	var nav Nav
	err := Db.First(&nav, id).Error
	return nav, err
}

//查询所有数据
func FindNavAll() ([]Nav, error) {
	var nav []Nav
	err := Db.Find(&nav).Error
	return nav, err
}

//删除
func DelNav(query interface{}, args ...interface{}) int64 {
	n := Db.Where(query, args...).Delete(Nav{}).RowsAffected
	return n
}

//更新
func UpdateNav(value map[string]interface{}, query interface{}, args ...interface{}) error {
	err := Db.Model(&Nav{}).Where(query, args...).Updates(value).Error
	return err
}

func (m *Nav) TableName() string {
	return TableName("nav")
}
