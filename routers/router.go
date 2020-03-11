package routers

import (
	"github.com/astaxie/beego"
)
import "beegoadmin/controllers/cms"

func init() {
	//beego.Router("/", &controllers.MainController{})

	//管理系统路由
	beego.Router("/cms", &cms.HomeController{}, "*:Index")                                //主页
	beego.Router("/cms/home/index", &cms.HomeController{}, "*:Index")                     //主页
	beego.Router("/cms/home/nav", &cms.HomeController{}, "get:Nav")                       //右侧导航页
	beego.Router("/cms/api/up_images", &cms.ApiController{}, "*:UpImages")                //上传图片接口
	beego.Router("/cms/home/my_info", &cms.HomeController{}, "*:MyInfo")                  //我的资料
	beego.Router("/cms/home/message", &cms.HomeController{}, "*:Message")                 //消息
	beego.Router("/cms/home/get_message", &cms.HomeController{}, "*:GetMessage")          //获取消息
	beego.Router("/cms/home/sys", &cms.HomeController{}, "*:Sys")                         //获取系统信息
	beego.Router("/cms/index/login", &cms.IndexController{}, "get:Login;post:CheckLogin") //登陆
	beego.Router("/cms/index/captcha", &cms.IndexController{}, "*:GetCaptchaData")        //获取验证码
	beego.Router("/cms/home/logout", &cms.HomeController{}, "*:Logout")                   //退出登陆
	beego.Router("/cms/nav/menu", &cms.NavController{}, "*:Menu")                         //获取菜单接口
	beego.Router("/cms/menu/list", &cms.MenuController{}, "*:List")                       //菜单列表
	beego.Router("/cms/menu/del", &cms.MenuController{}, "post:DelMenu")                  //菜单删除
	beego.Router("/cms/menu/edit", &cms.MenuController{}, "*:EditMenu")                   //菜单编辑
	beego.Router("/cms/menu/get", &cms.MenuController{}, "*:GetMenu")                     //菜单获取
	beego.Router("/cms/menu/add", &cms.MenuController{}, "*:AddMenu")                     //菜单添加
	beego.Router("/cms/menu/sort", &cms.MenuController{}, "*:SortMenu")                   //菜单更新排序
	beego.Router("/cms/admin/index", &cms.AdminController{}, "*:Index")                   //管理员添加
	beego.Router("/cms/admin/data", &cms.AdminController{}, "*:GetData")                  //获取管理员数据
	beego.Router("/cms/admin/add", &cms.AdminController{}, "*:Add")                       //添加管理员数据
	beego.Router("/cms/admin/edit", &cms.AdminController{}, "*:Edit")                     //编辑管理员数据
	beego.Router("/cms/admin/find", &cms.AdminController{}, "*:FindData")                 //获取管理员单条数据
	beego.Router("/cms/admin/del", &cms.AdminController{}, "post:Del")                    //删除管理员
	beego.Router("/cms/admin/batch_del", &cms.AdminController{}, "post:BatchDel")         //批量删除管理员
	beego.Router("/cms/role/index", &cms.RoleController{}, "*:Index")                     //角色列表页面
	beego.Router("/cms/role/data", &cms.RoleController{}, "*:GetData")                    //获取角色数据
	beego.Router("/cms/role/add", &cms.RoleController{}, "*:Add")                         //添加角色数据
	beego.Router("/cms/role/edit", &cms.RoleController{}, "*:Edit")                       //编辑角色数据
	beego.Router("/cms/role/find", &cms.RoleController{}, "*:FindData")                   //获取角色单条数据
	beego.Router("/cms/role/del", &cms.RoleController{}, "post:Del")                      //删除角色
}
