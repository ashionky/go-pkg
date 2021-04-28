/**
 * @Author pibing
 * @create 2021/3/31 2:30 PM
 */

package scripts

func PopRedisQueueTask() string {
	script :=
		`if #KEYS < 1 or (not #ARGV) then
          return false
         end
         local size = tonumber(ARGV[1])
         if size <= 0  then
           return false
         end
         
         local total = redis.call('lLen',KEYS[1])
         if total == 0  then
           return ''
         end 
         local start = 0
         local endindex   = -size-1
         if total > size then 
            start = total-size
         end
         local list  = redis.call('lrange',KEYS[1],start,-1)
         redis.call('ltrim',KEYS[1],0,endindex)
         return list`
	return script
}
