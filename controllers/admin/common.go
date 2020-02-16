package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type CommonController struct {
	beego.Controller
	moduleName string
	controllerName string
	actionName string
	o orm.Ormer
}

func (c *CommonController) Prepare() {
	c.moduleName = "admin"
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	c.o = orm.NewOrm()
}
