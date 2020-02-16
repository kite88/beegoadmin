package admin

import (
	"beegoadmin/models"
	"beegoadmin/utils"
	"github.com/astaxie/beego/cache"

	//"github.com/astaxie/beego/utils/captcha"
	captcha "github.com/kite88/golang/beegocaptcha"
	"strings"
)

type IndexController struct {
	CommonController
}

func (c *IndexController) Prepare() {
	c.CommonController.Prepare()
	//已经登录
	if c.GetSession("AdminUser") != nil {
		c.Redirect(c.URLFor("HomeController.Index"), 302)
	}
}
func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 105
	cpt.StdHeight = 39
}

var cpt *captcha.Captcha

//验证码数据
func (c *IndexController) GetCaptchaData() {


	ret := make(map[string]interface{})
	ret["code"] = 200
	ret["msg"] = "获取成功"
	data := make(map[string]string)
	srcPath,captchaId := cpt.CreateCaptchaData()
	if srcPath != ""  && captchaId != "" {
		data["captchaId"] = captchaId
		data["srcPath"] = srcPath
	}else{
		ret["code"] = 400
		ret["msg"] = "获取失败"
	}
	ret["data"] = data
	c.Data["json"] = ret
	c.ServeJSON()
}

//登陆页面
func (c *IndexController) Login() {
	c.TplName = c.moduleName + "/" + c.controllerName + "/login.html"
}

//登陆校验
func (c *IndexController) CheckLogin() {
	ret := make(map[string]interface{})
	ret["code"] = 200
	ret["msg"] = "请求成功"
	data := make(map[string]interface{})
	username := c.GetString("username")
	password := c.GetString("password")
	captcha := c.GetString("captcha")
	captcha_id := c.GetString("captcha_id")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		ret["code"] = 400
		ret["msg"] = "用户名或者密码为空"
	} else if strings.Trim(captcha," ") == "" {
		ret["code"] = 400
		ret["msg"] = "验证码为空"
	}else if true != cpt.Verify(captcha_id, captcha) {
		ret["code"] = 400
		ret["msg"] = "验证码错误"
	} else{
		password = utils.Md5(password)
		user := models.Admin{Username: username, Password: password}
		args := []string{"username", "password"}
		err := c.o.Read(&user, args...)
		if err != nil {
			ret["code"] = 400
			ret["msg"] = "账户不存在,请检查账号密码是否有误"
		} else {
			ret["code"] = 200
			ret["msg"] = "登陆成功"
			data["url"] = c.URLFor("HomeController.Index")
			c.SetSession("AdminUser", user)
		}
	}
	ret["data"] = data
	c.Data["json"] = ret
	c.ServeJSON()
}
