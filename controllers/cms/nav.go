package cms

import (
	//debug "github.com/favframework/debug"
	"beegoadmin/models"
)

/*
导航栏目管理
*/
type NavController struct {
	BaseController
}

//递归菜单
func (c *NavController) RecursionMenuList(pid int) []*models.NavTree {
	var lists []*models.NavTree
	list, _ := models.FindNavByWhere("pid = ?", pid)
	for _, value := range list {
		child := c.RecursionMenuList(value.Id)
		node := &models.NavTree{
			Id:     value.Id,
			Name:   value.Name,
			Pid:    value.Pid,
			Router: value.Router,
			Sort:   value.Sort,
		}
		node.Children = child
		lists = append(lists, node)
	}
	return lists
}

//菜单栏目接口
func (c *NavController) Menu() {
	menuList := c.RecursionMenuList(0)
	ret := ReturnJson{
		Code: 200,
		Msg:  "获取成功",
		Data: menuList,
	}
	c.Data["json"] = ret
	c.ServeJSON()
}
