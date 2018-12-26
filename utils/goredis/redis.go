package goredis

import (
	"ProjModules/utils/setting"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConnection *redis.Pool

func Setup()  {
	RedisConnection = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         setting.RedisSetting.MaxIdle,
		MaxActive:       setting.RedisSetting.MaxActive,
		IdleTimeout:     setting.RedisSetting.IdleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func Set(key string, data interface{}, time int) (bool, error) {
	connection := RedisConnection.Get()
	defer connection.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	reply, err := redis.Bool(connection.Do("SET", key, value))
	connection.Do("EXPIRE", key, time)
	return reply, err
}

func Exist(key string) bool {
	connection := RedisConnection.Get()
	defer connection.Close()
	exist, err := redis.Bool(connection.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exist
}

func Get(key string) ([]byte, error) {
	connection := RedisConnection.Get()
	defer connection.Close()
	reply, err := redis.Bytes(connection.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Delete(key string) (bool, error) {
	connection := RedisConnection.Get()
	defer connection.Close()
	reply, err := redis.Bool(connection.Do("DEL", key))
	return reply, err
}

func LikeDeletes(key string) error {
	connection := RedisConnection.Get()
	defer connection.Close()

	keys, err := redis.Strings(connection.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err := Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}