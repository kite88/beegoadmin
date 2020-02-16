package routers

import "github.com/astaxie/beego"
import "beegoadmin/controllers/admin"

func init() {
    //beego.Router("/", &controllers.MainController{})

	//后台路由
	beego.Router("/admin",&admin.HomeController{},"*:Index")//主页
	beego.Router("/admin/home/index",&admin.HomeController{},"*:Index")//主页
	beego.Router("/admin/home/nav",&admin.HomeController{},"get:Nav")//右侧导航页
	beego.Router("/admin/index/login",&admin.IndexController{},"get:Login;post:CheckLogin")//登陆
	beego.Router("/admin/index/captcha",&admin.IndexController{},"*:GetCaptchaData")
	beego.Router("/admin/home/logout",&admin.HomeController{},"*:Logout")//退出登陆
	beego.Router("/admin/nav/list",&admin.NavController{},"get:List")//获取菜单
}
