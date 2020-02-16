package admin

type BaseController struct {
	CommonController
	AdminUser interface{}
}

func (b *BaseController) Prepare() {
	b.CommonController.Prepare()
	b.AdminUser = b.GetSession("AdminUser")
	if b.AdminUser == nil {
		b.Redirect(b.URLFor("IndexController.Login"),302)
	}
}

