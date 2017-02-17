package cache

import (
    "github.com/astaxie/beego/cache"
)

func init() {
    cache, _ := cache.NewCache("memory", `{"interval":600}`)
    //if err == nil {
    //}
    cacheClient = cache
}

var cacheClient cache.Cache

// Get Cache Instance
func GetCClient() cache.Cache  {
    if cacheClient == nil {
        cache, _ := cache.NewCache("memory", `{"interval":600}`)
        cacheClient = cache
    }
    return cacheClient
}