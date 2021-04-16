/**
 * @Author pibing
 * @create 2020/12/31 6:12 PM
 */

package go_redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go-pkg/pkg/go-redis/redislock"
	"strings"
	"time"
)

var client *redis.Client
var clusterClient *redis.ClusterClient

func RedisInit(host, password string, db, poolsize int) error {

	redisOptions := &redis.Options{}
	redisOptions.Addr = host
	if password != "" {
		redisOptions.Password = password
	}
	redisOptions.DB = db
	if poolsize > 0 {
		redisOptions.PoolSize = poolsize
	}
	client = redis.NewClient(redisOptions)

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("redis RedisInit err: %v", err)
		return err
	}

	return nil
}

//redis集群
func RedisInitForCluster(hosts, password string, poolsize int) error {

	ips := strings.Split(hosts, ",")

	clusterOptions := &redis.ClusterOptions{}

	clusterOptions.Addrs = ips

	if password != "" {
		clusterOptions.Password = password
	}

	if poolsize > 0 {
		clusterOptions.PoolSize = poolsize
	}

	clusterClient = redis.NewClusterClient(clusterOptions)

	_, err := clusterClient.Ping().Result()
	if err != nil {
		fmt.Printf("redis RedisInitForCluster err: %v", err)
		return err
	}
	return nil
}

func Hset(key string, field string, value interface{}) error {
	return client.HSet(key, field, value).Err()
}

func HSet(key, field string, value interface{}) error {
	btr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.HSet(key, field, string(btr)).Err()
}

func HMSet(key string, fields map[string]interface{}) error {
	return client.HMSet(key, fields).Err()
}

func SetTTL(key string, extime int64) error {
	return client.Expire(key, time.Duration(extime)*time.Duration(time.Second)).Err()
}

func Del(key string) error {
	return client.Del(key).Err()
}

func Hdel(key string, fields string) error {
	return client.HDel(key, fields).Err()
}

func Hget(key string, field string) (string, error) {
	return client.HGet(key, field).Result()
}

func HGet(key, field string, T interface{}) error {
	str, err := client.HGet(key, field).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), T)
	if err != nil {
		return err
	}
	return nil
}

func HgetAll(key string) map[string]interface{} {
	return ConvertStringToMap(client.HGetAll(key).Val())
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func GetTtl(key string) time.Duration {
	return client.TTL(key).Val()
}

func SetByTtl(key string, value string, extime int64) error {
	return client.Set(key, value, time.Duration(extime)*time.Duration(time.Second)).Err()
}

func Set(key string, value string) error {
	return client.Set(key, value, time.Duration(-1)*time.Second).Err()
}

func IncrKey(key string, value int64) error {
	return client.IncrBy(key, value).Err()
}

func Keys(pattern string) ([]string, error) {
	return client.Keys(pattern).Result()
}

func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.Scan(cursor, match, count).Result()
}

func Expire(key string, extime int64) error {
	return client.Expire(key, time.Duration(extime)*time.Second).Err()
}

func ExpireAt(key string, ex time.Time) error {
	return client.ExpireAt(key, ex).Err()
}

func Hincrby(key string, field string, incr int64) error {
	return client.HIncrBy(key, field, incr).Err()
}

func HincrbyWithResult(key string, field string, incr int64) (error, int64) {
	result := client.HIncrBy(key, field, incr)
	return result.Err(), result.Val()
}

/**
 * redis 分布式锁
 * @param  key  {string} 			 初始化锁的互斥标示
 * @param  opts {*redislock.Options} 初始化锁的配置信息（详见该机构提的配置注释）
 * @return *    {*redislock.Options} 锁的实例（业务逻辑完成之后需要解锁）
 * @return *    {error} 			 异常
 */
func Lock(key string, opts *redislock.Options) (*redislock.Locker, error) {
	locker, err := redislock.Obtain(client, key, opts)

	if err != nil {
		return locker, err
	} else if locker == nil {
		return locker, errors.New("could not obtain lock!")
	}
	//defer locker.Unlock()

	return locker, nil
}

/**
 * redis 分布式锁 -（解锁）
 * @param  locker {*redislock.Locker} 锁的实例（初始化的时候得到）
 */
func UnLock(locker *redislock.Locker) {
	locker.Unlock()
}

// 转换从redis获取的数据
func ConvertStringToMap(base map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for k, v := range base {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(v), &dat); err == nil {
			resultMap[k] = dat
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}
