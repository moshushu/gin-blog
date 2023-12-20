package gredis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/moshushu/gin-blog/pkg/setting"
)

var RedisConn *redis.Pool

func Setup() error {
	// 设置 RedisConn 为 redis.Pool（连接池）
	RedisConn = &redis.Pool{
		// MaxIdle：最大空闲连接数
		MaxIdle: setting.RedisSetting.MaxIdle,
		// MaxActive：在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		MaxActive: setting.RedisSetting.MaxActive,
		// IdleTimeout：在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		// Dial：提供创建和配置应用程序连接的一个函数
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
			return c, err
		},
		// TestOnBorrow：可选的应用程序检查健康功能
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	// RedisConn.Get()：在连接池中获取一个活跃连接
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// conn.Do(commandName string, args ...interface{})：向 Redis 服务器发送命令并返回收到的答复
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	// redis.Bool 将命令返回转为布尔值
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
