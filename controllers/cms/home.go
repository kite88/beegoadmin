package cms

import (
	"beegoadmin/models"
	"beegoadmin/utils"
	"github.com/astaxie/beego"
	"os"
	"runtime"
	"strings"
	"time"
)

type HomeController struct {
	BaseController
	NavController
}

//主页
func (c *HomeController) Index() {
	//获取一级栏目菜单
	c.Data["TopMenu"], _ = models.FindNavByWhere("pid = ?", 0)
	c.Data["AdminUser"] = c.AdminUser
	c.Data["MessageCount"] = models.GetMessageCount("admin_id = ? AND is_read <> ?", c.AdminId, 1)
	c.Data["Navpath"] = c.URLFor("HomeController.Nav")
	c.TplName = c.moduleName + "/" + c.controllerName + "/index.html"
}

//右侧导航栏
func (c *HomeController) Nav() {
	MenuList := c.RecursionMenuList(0)
	c.Data["MenuList"] = MenuList
	c.TplName = c.moduleName + "/" + c.controllerName + "/nav.html"
}

//我的资料
func (c *HomeController) MyInfo() {
	if c.Ctx.Input.Method() == "POST" {
		var code = 200
		var msg = "请求成功"
		var data interface{}

		AdminId := utils.GetDataValue("Id", c.AdminUser)

		Password := strings.Trim(c.GetString("Password"), " ")
		HeadPic := strings.Trim(c.GetString("HeadPic"), " ")

		if Password == "" && HeadPic == "" {
			code = 400
			msg = "没有修改到任何东西"
		} else {
			save := make(map[string]interface{})
			if Password != "" {
				save["password"] = utils.Md5(Password)
			}
			if HeadPic != "" {
				save["head_pic"] = HeadPic
			}
			err := models.UpdateAdmin(save, "id = ?", AdminId.(int))
			if err == nil {
				code = 200
				msg = "资料修改成功"
				//更新session
				if adminR, errR := models.FindAdminByUserInfo(models.Admin{Id: AdminId.(int)}); errR == nil {
					c.SetSession("AdminUser", adminR)
					data = adminR
				}
			} else {
				code = 400
				msg = "修改失败"
			}
		}

		c.Data["json"] = ReturnJson{
			Code: code,
			Msg:  msg,
			Data: data,
		}
		c.ServeJSON()
	} else {
		AdminId := utils.GetDataValue("Id", c.AdminUser)
		Info, _ := models.FindAdminByUserInfo(models.Admin{Id: AdminId.(int)})
		c.Data["info"] = Info
		c.TplName = c.moduleName + "/" + c.controllerName + "/my_info.html"
	}
}

//消息
func (c *HomeController) Message() {
	if c.Ctx.Input.Method() == "POST" {
		var code int = 200
		var msg string = "请求成功"
		var data interface{} = nil

		AdminId := utils.GetDataValue("Id", c.AdminUser)

		viewRows, _ := c.GetInt("view_rows", 10)
		p, _ := c.GetInt("p", 1)
		key, _ := c.GetInt("key", 4)
		//今天的
		if key == 1 {
			BeginTime, EndTime := utils.GetSameDayUnix()
			data = models.GetMessageData(uint(viewRows), uint(p), "admin_id = ? AND (create_time BETWEEN ? AND ?) ", AdminId, BeginTime, EndTime)
		}
		//未读
		if key == 2 {
			data = models.GetMessageData(uint(viewRows), uint(p), "admin_id = ? AND is_read = ?", AdminId, 0)
		}
		//已读
		if key == 3 {
			data = models.GetMessageData(uint(viewRows), uint(p), "admin_id = ? AND is_read = ?", AdminId, 1)
		}
		//全部
		if key == 4 {
			data = models.GetMessageData(uint(viewRows), uint(p), "admin_id = ?", AdminId)
		}
		c.Data["json"] = ReturnJson{
			Code: code,
			Msg:  msg,
			Data: data,
		}
		c.ServeJSON()
	} else {
		c.TplName = c.moduleName + "/" + c.controllerName + "/message.html"
	}
}

//获取消息
func (c *HomeController) GetMessage() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{} = nil

	MessageId, _ := c.GetInt("id", 0)
	if MessageId == 0 {
		code = 400
		msg = "参数错误"
	} else {
		r, e := models.FindMessageByMessage(models.Message{Id: uint(MessageId), AdminId: uint(c.AdminId)})
		if e == nil {
			code = 200
			msg = "获取成功"
			//把信息标为 已读
			if r.IsRead != 1 {
				save := map[string]interface{}{
					"is_read":   1,
					"read_time": time.Now().Unix(),
				}
				models.UpdateMessage(save, "id = ?", r.Id)
			}
			//返回未读数量
			res := utils.StructToJsonToMap(r)
			res["UnreadMessageCount"] = models.GetMessageCount("admin_id = ? AND is_read <> ?", c.AdminId, 1)
			data = res
		} else {
			code = 400
			msg = "获取失败"
		}
	}
	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}

//退出登陆
func (c *HomeController) Logout() {
	c.DelSession("AdminUser")
	c.Redirect(c.URLFor("IndexController.Login"), 302)
}

type Sys struct {
	Name  string
	Value interface{}
}

//服务器信息
func (c *HomeController) Sys() {
	if c.Ctx.Input.IsPost() {
		var data []Sys
		data = append(data, Sys{
			Name:  "系统",
			Value: runtime.GOOS,
		})
		data = append(data, Sys{
			Name:  "系统架构",
			Value: runtime.GOARCH,
		})
		Hostname, _ := os.Hostname()
		data = append(data, Sys{
			Name:  "服务器名",
			Value: Hostname,
		})
		data = append(data, Sys{
			Name:  "Go版本",
			Value: runtime.Version(),
		})
		data = append(data, Sys{
			Name:  "主机名",
			Value: c.Ctx.Input.Host(),
		})
		LocalIP, _ := utils.GetLocalIP()
		data = append(data, Sys{
			Name:  "服务器IP(内网)",
			Value: LocalIP,
		})
		PubNetIP, _ := utils.GetPublicNetIP()
		data = append(data, Sys{
			Name:  "服务器IP(公网)",
			Value: PubNetIP,
		})
		data = append(data, Sys{
			Name:  "域名",
			Value: c.Ctx.Input.Domain(),
		})
		data = append(data, Sys{
			Name:  "端口",
			Value: beego.AppConfig.String("httpport"),
		})
		data = append(data, Sys{
			Name:  "数据库",
			Value: models.DbVersion(),
		})
		data = append(data, Sys{
			Name:  "数据库地址",
			Value: beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport"),
		})
		data = append(data, Sys{
			Name:  "服务器时间",
			Value: time.Now().Format("2006-01-02 15:04:05"),
		})
		ExecutePath, _ := utils.GetExecutePath()
		data = append(data, Sys{
			Name:  "web目录",
			Value: ExecutePath,
		})
		data = append(data, Sys{
			Name:  "服务器存储空间",
			Value: utils.GetDisk(),
		})

		c.Data["json"] = ReturnJson{
			Code: 200,
			Msg:  "请求成功",
			Data: data,
		}
		c.ServeJSON()
	} else {
		c.TplName = c.moduleName + "/" + c.controllerName + "/sys.html"
	}
}
