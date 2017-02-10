package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"html/template"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

type A struct {
	Website string
	Email  string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.Layout = "index.html"
	c.TplName = "index.tpl"

	//dat := &A{"beego.me", "astaxie@gmail.com"}

	//c.Ctx.WriteString("hello world")
	//renderHTML(c.Ctx.ResponseWriter, "index.html", dat)
	//c.Ctx.Redirect(200, "index")
}

func (c *MainController) HelloTpl() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.Layout = "index.html"
	c.TplName = "index.tpl"

	dat := &A{"beego.me", "astaxie@gmail.com"}
	val := beego.AppConfig.String("autorender")

	//logs.SetLogger(logs.AdapterConsole)
	logs.Info("autorender %s", val)

	renderTemplate(c.Ctx.ResponseWriter, "index", dat)
}

// 渲染页面并输出
func renderTemplate(w http.ResponseWriter, file string, data interface{}) {
	filePath := "./views/" + file + ".tpl";
	b, _ := ioutil.ReadFile(filePath)
	if b == nil {
		logs.Info("file path %s", filePath)
	}
	// 获取页面内容
	t := template.New(file + ".tpl")
	t, err := t.ParseFiles("./views/" + file + ".tpl")
	if err != nil {
		fmt.Println(err)
	}
	//checkErr(err)
	// 将页面渲染后反馈给客户端
	err2 := t.Execute(w, data)
	if err != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}