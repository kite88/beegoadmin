package cms

import "beegoadmin/utils"

type BaseController struct {
	CommonController
	AdminUser interface{}
	AdminId int
}

func (b *BaseController) Prepare() {
	b.CommonController.Prepare()
	b.AdminUser = b.GetSession("AdminUser")
	if b.AdminUser == nil {
		b.Redirect(b.URLFor("IndexController.Login"),302)
	}
	b.AdminId = utils.GetDataValue("Id", b.AdminUser).(int)
}

