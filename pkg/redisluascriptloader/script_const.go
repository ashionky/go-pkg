package redisluascriptloader

const (

	// 加载状态
	ScriptStatusUnLoaded = 0 // 未加载
	ScriptStatusReady    = 1 // 已加载可用
	ScriptStatusLoading  = 2 // 正在加载

	LoadWaitingQueueLength = 200 // 最大等待加载脚本数量

	// 脚本名称

	PopRedisQueueTask = "popRedisQueueTask" //队列
)
