package main

import (
"fmt"
"github.com/go-redis/redis" // 实现了redis连接池
"time"
)

// 定义redis链接池
var client *redis.Client

// 初始化redis链接池
func init(){
	client = redis.NewClient(&redis.Options{
	//	Addr:     config.RedisAddr, // Redis地址
	//	Password: config.RedisPwd,  // Redis账号
	//	DB:       config.RedisDB,   // Redis库
	//	PoolSize: config.PoolSize,  // Redis连接池大小
		MaxRetries: 3,              // 最大重试次数
		IdleTimeout: 10*time.Second,            // 空闲链接超时时间
	})
	pong, err := client.Ping().Result()
	if err == redis.Nil {
		//logger.Info("Redis异常")
	} else if err != nil {
		//logger.Info("失败:", err)
	} else {
		//logger.Info(pong)
	}
}

// 向key的hash中添加元素field的值
func HashSet(key, field string, data interface{}) {
	client.ZAdd()
	err := client.HSet(key, field, data)
	if err != nil {
		logger.Error("Redis HSet Error:", err)
	}
}

// 批量向key的hash添加对应元素field的值
func BatchHashSet(key string, fields map[string]interface{}) string {
	val, err := client.HMSet(key, fields).Result()
	if err != nil {
		logger.Error("Redis HMSet Error:", err)
	}
	return val
}

// 通过key获取hash的元素值
func HashGet(key, field string) string {
	result := ""
	val, err := client.HGet(key, field).Result()
	if err == redis.Nil {
		logger.Info("Key Doesn't Exists:", field)
		return result
	}else if err != nil {
		logger.Info("Redis HGet Error:", err)
		return result
	}
	return val
}

// 批量获取key的hash中对应多元素值
func BatchHashGet(key string, fields ...string) map[string]interface{} {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := client.HGet(key, fmt.Sprintf("%s", field)).Result()
		if err == redis.Nil {
			logger.Info("Key Doesn't Exists:", field)
			resMap[field] = result
		}else if err != nil {
			logger.Info("Redis HMGet Error:", err)
			resMap[field] = result
		}
		if val != "" {
			resMap[field] = val
		}else {
			resMap[field] = result
		}
	}
	return resMap
}

// 获取自增唯一ID
func Incr(key string) int {
	val, err := client.Incr(key).Result()
	if err != nil {
		logger.Error("Redis Incr Error:", err)
	}
	return int(val)
}

// 添加集合数据
func SetAdd(key, val string){
	client.SAdd(key, val)
}

// 从集合中获取数据
func SetGet(key string)[]string{
	val, err := client.SMembers(key).Result()
	if err != nil{
		logger.Error("Redis SMembers Error:", err)
	}
	return val
}
