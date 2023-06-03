package utils

import (
	"fmt"
	"ginweb/config"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitRedisPool() *redis.Pool {
	host := config.ConfSetting.Redis.Host
	port := config.ConfSetting.Redis.Port
	password := config.ConfSetting.Redis.Password
	dial := redis.DialPassword(password)
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port), dial)
		},
	}
	return pool
}

/*
*
获取redis
*/
func GetRedis(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return result, nil
}

/*
*
设置redis
*/
func SetRedis(key string, value string, expire int) error {
	conn := pool.Get()
	defer conn.Close()
	err := conn.Send("SET", key, value)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, expire)
	if err != nil {
		return err
	}
	err = conn.Flush()
	if err != nil {
		return err
	}
	return nil
}
