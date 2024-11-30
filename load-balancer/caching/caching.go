package caching

import (
	"fmt"
	inMemory "load-balancer/caching/inmemory"
	redisCaching "load-balancer/caching/redis"
	"load-balancer/caching/structure"
	"load-balancer/global"
	"net/http"
)

var SetCache func(string, []byte, *http.Response)
var GetCachedResponse func(string) (*structure.Cache, bool)

func InitCaching() {
	if global.Data["caching"] == "redis" {
		SetCache = redisCaching.SetCache
		GetCachedResponse = redisCaching.GetCachedResponse
	} else if global.Data["caching"] == "in-memory" {
		SetCache = inMemory.SetCache
		GetCachedResponse = inMemory.GetCachedResponse
	} else {
		SetCache = func(string, []byte, *http.Response) {}
		GetCachedResponse = func(string) (*structure.Cache, bool) { return nil, false }
		fmt.Println("No caching mechanism selected.")
	}
}
