package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type Article struct {
    Id int64
    ArticleTitle string
    ArticleContent string
    Hot int32 `orm:"default(0)"`
    Category int8
    Status int8
    IsDel int8 `orm:"default(0)"`
    Creator int64
    Reviser int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

type ArticleCategory struct {
    Id int64
    ArticleId int64
    CategoryId int32
    Creator int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

type ArticleTag struct {
    Id int64
    ArticleId int64
    TagId int64
    Creator int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

type Dict struct {
    Id int32
    Title string `orm:"size(128)"`
    Type string `orm:"size(128)"`
    Val string `orm:"size(128)"`
    Order int16 `orm:"default(0)"`
    IsDel int8 `orm:"default(0)"`
    Creator int64
    Reviser int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

type User struct {
    Id int64
    Username string `orm:"size(128)"`
    Email string `orm:"size(128)"`
    Account string `orm:"size(128)"`
    Pwd string `orm:"size(45)"`
    Status int8 `orm:"default(1)"`
    IsDel int8 `orm:"default(0)"`
    Creator int64
    Reviser int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

type UserTag struct {
    Id int64
    TagName string `orm:"size(64)"`
    UserId int64
    IsDel int8 `orm:"default(0)"`
    Creator int64
    Reviser int64
    CreateTime time.Time
    ModifyTime time.Time `orm:"null"`
}

func init()  {
    // register model
    orm.RegisterModelWithPrefix("ks_", new(Article), new(ArticleCategory), new(ArticleTag), new(Dict), new(User), new(UserTag))
}