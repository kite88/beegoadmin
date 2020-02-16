package admin

import (
	//debug "github.com/favframework/debug"
	"beegoadmin/utils"
)
type HomeController struct {
	BaseController
}

//主页
func (c *HomeController) Index() {
	c.Data["Username"] = utils.GetDataValue("Username",c.AdminUser)
	c.Data["Navpath"] = c.URLFor("HomeController.Nav")
	c.TplName = c.moduleName + "/" + c.controllerName + "/index.html"
}
//右侧导航栏
func (c *HomeController) Nav() {
	c.TplName = c.moduleName + "/" + c.controllerName + "/nav.html"
}

//退出登陆
func (c *HomeController) Logout()  {
	c.DelSession("AdminUser")
	c.Redirect(c.URLFor("IndexController.Login"),302)
}

