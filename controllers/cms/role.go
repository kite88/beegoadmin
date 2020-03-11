package cms

import (
	"beegoadmin/models"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

//获取单条数据
func (c *RoleController) FindData() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{} = nil
	Id := c.Input().Get("id")
	IntId, _ := strconv.Atoi(Id)
	if strings.Trim(Id, " ") == "" || IntId == 0 {
		code = 400
		msg = "参数错误"
	} else {
		if ret, err := models.FindRoleDataByRole(models.Role{Id: uint(IntId)}); err == nil {
			code = 200
			msg = "获取成功"
			data = ret
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
func (c *RoleController) GetData() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{} = nil

	//分页参数
	p := c.Input().Get("p")
	IntP, _ := strconv.Atoi(p)
	//分页显示记录条数
	viewRows := c.Input().Get("view_rows")
	IntViewRows, _ := strconv.Atoi(viewRows)
	//获取搜索时的关键字
	RoleName := strings.Trim(c.Input().Get("role_name"), " ")
	if RoleName != "" {
		data = models.GetRoleData(uint(IntViewRows), uint(IntP), "name LIKE ?", "%"+RoleName+"%")
	} else {
		data = models.GetRoleData(uint(IntViewRows), uint(IntP), "")
	}

	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}

//列表页面
func (c *RoleController) Index() {
	p := strings.Trim(c.Input().Get("p"), " ")
	c.Data["p"] = p
	c.TplName = c.moduleName + "/" + c.controllerName + "/index.html"
}

//添加数据
func (c *RoleController) Add() {
	if c.Ctx.Input.Method() == "POST" {
		var code int = 200
		var msg string = "请求成功"
		var data interface{} = nil

		name := c.Input().Get("name")
		if strings.Trim(name, " ") == "" {
			code = 400
			msg = "角色名称不能为空"
		} else {
			if err := models.AddRole(&models.Role{Name: name}); err == nil {
				code = 200
				msg = "创建成功"
			} else {
				code = 400
				msg = "创建失败"
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

//编辑数据
func (c *RoleController) Edit() {
	if c.Ctx.Input.Method() == "POST" {
		var code = 200
		var msg = "请求成功"

		Id := c.Input().Get("id")
		Name := c.Input().Get("name")
		IntId, _ := strconv.Atoi(Id)
		if strings.Trim(Id, " ") == "" || IntId == 0 {
			code = 400
			msg = "参数错误"
		} else if strings.Trim(Name, " ") == "" {
			code = 400
			msg = "角色名称不能为空"
		} else {
			save := make(map[string]interface{})
			save["name"] = Name
			if err := models.UpdateRole(save, "id = ?", IntId); err == nil {
				code = 200
				msg = "修改成功"
			} else {
				code = 400
				msg = "修改失败"
			}
		}
		c.Data["json"] = ReturnJson{
			Code: code,
			Msg:  msg,
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

//删除数据
func (c *RoleController) Del() {
	var code = 200
	var msg = "请求成功"

	Id := c.Input().Get("id")
	if strings.Trim(Id, " ") == "" {
		code = 400
		msg = "参数错误"
	} else {
		IntId, _ := strconv.Atoi(Id)
		if DelN := models.DelRole("id = ?", IntId); DelN > 0 {
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