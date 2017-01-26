package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"html/template"
)

type MainController struct {
	beego.Controller
}
// 渲染页面并输出
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, err := template.New(file).ParseFiles("views/" + file)
	//checkErr(err)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("hello world")
	c.Ctx.Redirect(200, "index")
}
