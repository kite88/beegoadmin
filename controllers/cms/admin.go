package cms

import (
	"beegoadmin/models"
	"beegoadmin/utils"
	"fmt"
	"strconv"
	"strings"
)

type AdminController struct {
	BaseController
}

//获取单条数据
func (c *AdminController) FindData() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{} = nil
	Id := c.Input().Get("id")
	IntId, _ := strconv.Atoi(Id)
	if strings.Trim(Id, " ") == "" || IntId == 0 {
		code = 400
		msg = "参数有误"
	} else {
		r, err := models.FindAdminByUserInfo(models.Admin{Id: IntId})
		if err == nil {
			code = 200
			msg = "获取成功"
			data = r
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

//获取数据
func (c *AdminController) GetData() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{} = nil

	//分页参数
	p := c.Input().Get("p")
	IntP, _ := strconv.Atoi(p)
	//分页显示记录条数
	viewRows := c.Input().Get("view_rows")
	IntViewRows, _ := strconv.Atoi(viewRows)

	data = models.GetAdminData(uint(IntViewRows), uint(IntP))

	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}

//列表页
func (c *AdminController) Index() {
	p := strings.Trim(c.Input().Get("p"), " ")
	c.Data["p"] = p
	c.TplName = c.moduleName + "/" + c.controllerName + "/index.html"
}

//添加账号
func (c *AdminController) Add() {
	if c.Ctx.Input.Method() == "POST" {
		var code int = 200
		var msg string = "请求成功"
		var data interface{} = nil

		username := c.GetString("username")
		password := c.GetString("password")
		role := c.GetString("role")

		if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
			code = 400
			msg = "用户名或者密码不能为空"
		} else if _, err := models.FindAdminByUserInfo(models.Admin{Username: username}); err == nil {
			code = 400
			msg = "该用户已经存在，换一个试试"
		} else if strings.Trim(role, " ") == "" {
			code = 400
			msg = "角色不能为空"
		} else {
			RoleIntId, _ := strconv.Atoi(role)
			save := models.Admin{
				Username: username,
				Password: utils.Md5(password),
				RoleId:   RoleIntId,
			}
			if err := models.AddAdmin(&save); err != nil {
				code = 400
				msg = "添加失败"
			} else {
				code = 200
				msg = "添加成功"
			}
		}
		c.Data["json"] = ReturnJson{
			Code: code,
			Msg:  msg,
			Data: data,
		}
		c.ServeJSON()
	} else {
		c.TplName = c.moduleName + "/" + c.controllerName + "/add.html"
	}
}

//编辑账号
func (c *AdminController) Edit() {
	if c.Ctx.Input.Method() == "POST" {
		var code int = 200
		var msg string = "请求成功"
		var data interface{} = nil

		id := c.Input().Get("id")
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		role := c.Input().Get("role")
		intId, _ := strconv.Atoi(id)

		if strings.Trim(id, " ") == "" || intId == 0 {
			code = 400
			msg = "参数错误"
		} else if strings.Trim(username, " ") == "" {
			code = 400
			msg = "用户名不能为空"
		} else if count := models.FindAdminCount("id <> ? and username = ?", intId, username); count > 0 {
			code = 400
			msg = "用户名已经存在,换其他的试试"
			data = count
		} else {
			save := make(map[string]interface{})
			save["username"] = username
			password = strings.Trim(password, " ")
			if password != "" {
				save["password"] = utils.Md5(password)
			}
			if role := strings.Trim(role, " "); role != "" {
				IntRoleId, _ := strconv.Atoi(role)
				save["role_id"] = IntRoleId
			}
			err := models.UpdateAdmin(save, "id = ?", intId)
			if err == nil {
				code = 200
				msg = "更新成功"
			} else {
				code = 400
				msg = "更新失败"
			}
		}
		c.Data["json"] = ReturnJson{
			Code: code,
			Msg:  msg,
			Data: data,
		}
		c.ServeJSON()
	} else {
		Id := c.Input().Get("id")
		c.Data["id"] = Id
		P := c.Input().Get("p")
		c.Data["p"] = P
		c.TplName = c.moduleName + "/" + c.controllerName + "/edit.html"
	}
}

//删除
func (c *AdminController) Del() {
	var code = 200
	var msg = "请求成功"

	Id := c.Input().Get("id")
	IntId, _ := strconv.Atoi(Id)
	if strings.Trim(Id, " ") == "" || IntId == 0 {
		code = 200
		msg = "参数错误"
	} else {
		if DelNumber := models.DelAdmin("id = ?", IntId); DelNumber > 0 {
			code = 200
			msg = "删除成功"
		} else {
			code = 400
			msg = "删除失败"
		}
	}
	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
	}
	c.ServeJSON()
}

//批量删除
func (c *AdminController) BatchDel() {
	var code = 200
	var msg = "请求成功"
	var data interface{}

	ids := strings.Trim(c.Input().Get("idString"), " ")
	if "" == ids {
		code = 400
		msg = "参数不能为空"
	} else {
		if n := models.DelAdmin("id in (" + ids + ")"); n > 0 {
			code = 200
			msg = fmt.Sprintf("成功删除了%d条记录", n)
		} else {
			code = 400
			msg = "删除失败"
		}
	}
	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}