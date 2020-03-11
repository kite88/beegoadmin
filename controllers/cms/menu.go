package cms

import (
	"beegoadmin/models"
	//debug "github.com/favframework/debug"
	"strconv"
	"strings"
)

type MenuController struct {
	BaseController
}

//菜单列表
func (c *MenuController) List() {

	c.TplName = c.moduleName + "/" + c.controllerName + "/list.html"
}

//获取菜单
func (c *MenuController) GetMenu() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{}

	Id := c.GetString("id")
	if strings.Trim(Id, " ") == "" {
		code = 400
		msg = "参数错误"
	} else {
		IntId, _ := strconv.Atoi(Id)
		info, err := models.FindNavFirst(IntId)
		if err == nil {
			code = 200
			msg = "获取成功"
			data = info
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

//菜单添加
func (c *MenuController) AddMenu() {
	if c.Ctx.Input.IsPost() {
		var code int = 200
		var msg string = "请求成功"
		var data interface{}

		Name := strings.Trim(c.GetString("name"), " ")
		Router := strings.Trim(c.GetString("router"), " ")
		Pid, _ := strconv.Atoi(c.GetString("pid"))
		if Name == "" {
			code = 200
			msg = "名称不能为空"
		} else {
			save := models.Nav{
				Name:   Name,
				Router: Router,
				Pid:    Pid,
			}
			err := models.AddNav(&save)
			if err == nil {
				code = 200
				msg = "添加成功"
			} else {
				code = 400
				msg = "添加失败"
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

//编辑菜单
func (c *MenuController) EditMenu() {
	if c.Ctx.Input.IsPost() {
		var code int = 200
		var msg string = "请求成功"
		var data interface{}

		Id, _ := strconv.Atoi(c.GetString("id"))
		Name := strings.Trim(c.GetString("name"), " ")
		Router := strings.Trim(c.GetString("router"), " ")
		Pid, _ := strconv.Atoi(c.GetString("pid"))
		if Name == "" {
			code = 400
			msg = "名称不能为空"
		} else {
			save := make(map[string]interface{})
			save["name"] = Name
			save["router"] = Router
			save["pid"] = Pid
			err := models.UpdateNav(save, "id = ?", Id)
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
		id := strings.Trim(c.GetString("id"), " ")
		c.Data["Id"] = id
		c.TplName = c.moduleName + "/" + c.controllerName + "/edit.html"
	}
}

//删除菜单
func (c *MenuController) DelMenu() {
	var code int = 200
	var msg string = "请求成功"
	var data interface{}

	Id := c.GetString("id")
	if strings.Trim(Id, " ") == "" {
		code = 400
		msg = "参数不能为空"
	} else {
		IntId, _ := strconv.Atoi(Id)
		if _, err := models.FindNavFirst(IntId); err != nil {
			code = 400
			msg = "该菜单不存在"
		} else if count := models.FindNavCount("pid = ?", IntId); count > 0 {
			code = 400
			msg = "该菜单拥有子菜单不能删除"
		} else {
			if delNum := models.DelNav("id = ?", IntId); delNum > 0 {
				code = 200
				msg = "删除成功"
			} else {
				code = 400
				msg = "删除失败"
			}
		}
	}
	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}

//更新排序
func (c *MenuController) SortMenu() {
	var code = 200
	var msg = "请求成功"

	id := strings.Trim(c.Input().Get("id"), " ")
	value := strings.Trim(c.Input().Get("value"), " ")

	if id == "" {
		code = 400
		msg = "参数错误"
	} else {
		IntId, _ := strconv.Atoi(id)
		IntValue, _ := strconv.Atoi(value)
		save := make(map[string]interface{})
		save["sort"] = IntValue
		if err := models.UpdateNav(save, "id = ?", IntId); err == nil {
			code = 200
			msg = "更新成功"
		}else {
			code = 400
			msg = "更新失败"
		}
	}
	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
	}
	c.ServeJSON()
}
