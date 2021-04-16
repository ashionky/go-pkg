- 文件结构

    - `lua_script_definition`文件，定义脚本对应map

    - `scripts`目录存放需要用到的lua脚本函数，文件名对应脚本名称，对应函数返回lua脚本字符串

    - `script_const`文件，对应脚本名字等常量定义

- 代码使用

```golang
scriptLoader := redisluascriptloader.NewRedisScriptLoader(deps.CacheRedis)

res, err := scriptLoader.ExecScript(redisluascriptloader.PopRedisQueueTask, keys, args)
```
