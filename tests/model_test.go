package testmain

import (
    "testing"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
    "gks/models"
    "fmt"
)

func init() {
    // mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // 参数4(可选)  设置最大空闲连接
    // 参数5(可选)  设置最大数据库连接
    maxIdle := 30
    maxConn := 30
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:password@/gks?charset=utf8", maxIdle, maxConn)

    // create table
    orm.RunSyncdb("default", false, false)
}

func TestDBLink(t *testing.T) {
    fmt.Println("\nfunc TestDBLink:")
    o, err := orm.GetDB()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(o)
}

func TestQueryModel(t *testing.T) {
    fmt.Println("\nfunc TestQueryModel:")
    o := orm.NewOrm()
    //o.Using("gks")
    art := &models.Article{Id: 1}
    err := o.Read(art)
    if err == orm.ErrNoRows {
        fmt.Println(err)
    } else {
        fmt.Printf("id %v, title %s\n", art.Id, art.ArticleTitle)
        fmt.Println(art)
        art.ArticleTitle = "update title"

        if num, err := o.Update(art); err == nil {
            fmt.Printf("update id %v\n", num)
        }

        art.Id = art.Id + 1
        art.ArticleTitle = "new art"
        //if num, err := o.Insert(art); err == nil {
        //    fmt.Printf("insert id %v\n", num)
        //}
    }

}

func TestQuerySeter(t *testing.T)  {
    fmt.Println("\nfunc TestQuerySeter:")
    o := orm.NewOrm()
    fmt.Println(o)
    qs := o.QueryTable(&models.Article{})
    qs.Filter("is_del", 0)
    var maps []orm.Params
    num, err := qs.Values(&maps)
    if err == nil {
        fmt.Printf("Result Nums: %d\n", num)
        for _, m := range maps {
            fmt.Println(m["ArticleTitle"], m["Id"])
        }
    }
}

func TestRawQuery(t *testing.T)  {
    fmt.Println("\nfunc TestRawQuery:")
    o := orm.NewOrm()
    fmt.Println(o)

    //o.Using("gks")
    var result orm.RawSeter
    result = o.Raw("select * from ks_article")
    var maps []orm.Params
    num, err := result.Values(&maps)
    if err == nil && num > 0 {
        for _, m := range maps {
            fmt.Println(m["article_title"], m["id"])
        }
    }

}