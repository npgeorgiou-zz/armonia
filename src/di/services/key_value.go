package services

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type KeyValue interface {
	SetForEver(key string, value string)
	Set(key string, value string, keepFor int)
	Get(key string) string
}

type RedisKeyValue struct {
	driver *redis.Client
}

func NewRedisKeyValue(driver *redis.Client) RedisKeyValue {
	return RedisKeyValue{driver:driver}
}

func (kv *RedisKeyValue) SetForEver(key string, value string) {
	kv.Set(key, value, 0)
}

func (kv *RedisKeyValue) Set(key string, value string, keepFor int) {
	kv.driver.Set(
		key,
		value,
		time.Duration(keepFor)*time.Second,
	)
}

func (kv *RedisKeyValue) Get(key string) string {
	return kv.driver.Get(key).String()
}
