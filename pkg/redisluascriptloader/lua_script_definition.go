package redisluascriptloader

import (
	"go-pkg/pkg/redisluascriptloader/scripts"
)

type shas struct {
	Sha    string // redis
	Script string // 脚本SHA1校验值
	Status int64  // 脚本加载状态，0：未加载；1：已加载可用；2：正在加载
	Name   string // 脚本名称
}

var (
	ScriptPopRedisQueue = &shas{
		Sha:    "",
		Script: scripts.PopRedisQueueTask(),
		Status: ScriptStatusUnLoaded,
		Name:   "PopRedisQueueTask",
	}
)

var scriptsDefine = map[string]*shas{
	PopRedisQueueTask: ScriptPopRedisQueue,
}
