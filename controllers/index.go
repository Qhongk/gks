package controllers

import (
    "github.com/astaxie/beego"
    "net/http"
    "html/template"
    "fmt"
    "github.com/astaxie/beego/logs"
)

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {
    c.Data["Website"] = "hello.me"
    c.Data["Email"] = "xxx@gmail.com"
    c.Layout = "layout.html"
    c.TplName = "site/index.html"

    c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head.html"
    c.LayoutSections["Scripts"] = "scripts.html"
    c.LayoutSections["FootBar"] = "side_bar.html"

    //services.QueryArticle(2)
}

func (c *MainController) HelloTpl() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"] = "astaxie@gmail.com"
    c.TplName = "site/index.html"

    val := beego.AppConfig.String("autorender")

    //logs.SetLogger(logs.AdapterConsole)
    logs.Info("autorender %s", val)

    renderTemplate(c.Ctx.ResponseWriter, "index", c.Data)
    //c.Ctx.WriteString("hello world")
    //c.Ctx.Redirect(200, "json")
}

// 渲染页面并输出
func renderTemplate(w http.ResponseWriter, file string, data interface{}) {
    rootPath := "/Users/uc/develop/workspace/GoglandProjects/gks"
    //"/Users/uc/develop/workspace/GoglandProjects/gks/views/index.tpl"
    //filePath := "./views/" + file + ".tpl"
    //b, _ := ioutil.ReadFile(filePath)
    //if b == nil {
    //	logs.Info("file path %s", filePath)
    //}
    // 获取页面内容
    t := template.New(file + ".html")
    t, err := t.ParseFiles(rootPath + "/views/site/" + file + ".html")
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
