package services

import (
    "github.com/astaxie/beego/orm"
    "fmt"
    "gks/utils/cache"
    "time"
)

type ArticleService struct {

}

func QueryArticle(id int) []orm.Params {
    c := cache.GetCClient().Get("xxxxx")
    t := cache.GetCClient().Get("title")
    if c != nil || t != nil {
        fmt.Println(c)
        fmt.Println(t)
        return nil
    }

    o := orm.NewOrm()
    var result orm.RawSeter
    result = o.Raw("select * from ks_article where id = ?", id)
    var maps []orm.Params
    num, err := result.Values(&maps)
    if err == nil && num > 0 {
        for _, m := range maps {
            fmt.Println(m["article_title"], m["id"])
        }
    }

    cache.GetCClient().Put("xxxxx", maps, 60*time.Second)
    cache.GetCClient().Put("title", maps[0]["article_title"], 60*time.Second)
    return maps
}