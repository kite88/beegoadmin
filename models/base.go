package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseModel struct{}

var Db *gorm.DB

func init() {
	fmt.Print("gorm init ")
	DbInit()
}

func DbInit() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	var err error
	Db, err = gorm.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败", err)
	} else {
		fmt.Println("数据库连接成功", DbVersion())
	}
	Db.LogMode(true)
	Db.SingularTable(true)
}

func DbVersion() string {
	var rows string
	result := Db.Raw("select version() as version").Row()
	result.Scan(&rows)
	return "mysql " + rows
}

func DbClose() {
	defer Db.Close()
}

func TableName(tableName string) string {
	return beego.AppConfig.String("dbprefix") + tableName
}