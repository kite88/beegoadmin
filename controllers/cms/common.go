package cms

import (
	"beegoadmin/utils"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type CommonController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
}

type ReturnJson struct {
	Code int
	Msg  string
	Data interface{}
}
type NullObject struct{}
type NullArray []struct{}

func (c *CommonController) Prepare() {
	c.moduleName = "cms"
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
}

/*
	上传文件
	file 上传的file name值 示例
	size 大小限制 字节
	allowExtMap 文件后缀限制
	uploadDir 上传的目录
*/
func (c *CommonController) UpFile(file string, size int64, allowExtMap map[string]bool, uploadDir string) (Code int, Error string, Path string) {
	f, h, e := c.GetFile(file) //获取上传的文件
	if e != nil {
		Code = 0
		Error = fmt.Sprintf("未知错误:%v", e)
		return
	}
	if h.Size > size {
		Code = 0
		Error = fmt.Sprintf("文件大小不能超过%s", utils.FormatByte(uint64(size)))
		return
	}
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	if _, ok := allowExtMap[ext]; !ok {
		Code = 0
		Error = "后缀名不符合上传要求"
		return
	}
	//创建目录
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		Code = 0
		Error = fmt.Sprintf("%v", err)
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ))

	fileName := fmt.Sprintf("%x", hashName) + ext

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile(file, fpath)
	if err != nil {
		Code = 0
		Error = fmt.Sprintf("%v", err)
		return
	}
	Code = 1
	Error = "上传成功~！"
	if fpath != "" {
		Path = "/" + fpath
	}
	return
}
