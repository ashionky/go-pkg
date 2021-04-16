/**
 * @Author pibing
 * @create 2021/3/31 2:30 PM
 */

package scripts

func PopRedisQueueTask() string {
	script :=
		`if #KEYS < 2 or (not #ARGV) then
          return false
         end
         local total = tonumber(ARGV[1])
         if total <= 0  then
           return false
         end
         local remain = redis.call('lLen',KEYS[1])
         local maxPer = tonumber((total - remain) / (#KEYS - 1))
         for id,key in pairs(KEYS) do
           if not total then
              break
           end
           local curPer = maxPer
           while id >=2 and total and curPer > 0 do
             local popValue = redis.call('rPopLpush',KEYS[id],KEYS[1])
              if popValue then
                 curPer = curPer - 1
                 total = total - 1
              else
                 break
              end
           end
         end
         return redis.call('lrange',KEYS[1],0, -1)`
	return script
}
