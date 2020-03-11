package cms

import "time"

type ApiController struct {
	BaseController
}

//上传图片
func (c *ApiController) UpImages() {
	var code = 200
	var msg = "请求成功"
	var data = ""

	var file string
	var size int64
	var allowExtMap map[string]bool
	var uploadDir string

	file = "file"
	size = 4194304
	allowExtMap = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
	}
	uploadDir = "static/upload_path/images/" + time.Now().Format("2006/01/02/")

	FCode, FError, FPath := c.UpFile(file, size, allowExtMap, uploadDir)
	if FCode == 1 {
		code = 200
	} else {
		code = 400
	}
	msg = FError
	data = FPath

	c.Data["json"] = ReturnJson{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}
