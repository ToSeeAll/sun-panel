package cUserToken

import (
	"sun-panel/global"
	"sun-panel/lib/cache"

	"time"
)

func InitCUserToken() cache.Cacher[string] {
	return global.NewCache[string](876000*time.Hour, 876000*time.Hour, "CUserToken")
}

// func InitVerifyCodeCachePool() {
// 	global.VerifyCodeCachePool = cache.NewGoCache(10*time.Minute, 60*time.Second)
// }
