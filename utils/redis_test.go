package utils

import (
	"encoding/json"
	"fmt"
	"ginweb/config"
	"ginweb/model"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func Test_InitRedisPool(t *testing.T) {
	config.SetUpConfig()
	pool = InitRedisPool()

	c := pool.Get()
	defer c.Close()

	_, err := c.Do("Set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}

	fmt.Println(r)
	pool.Close() //关闭连接池
}

func TestSetRedis(t *testing.T) {
	err := SetRedis("key", []byte("123456"), 60)
	if err != nil {
		println("err=====")
	}
	println("info=====")
}

func TestGetRedis(t *testing.T) {

	user := model.SysUser{
		Phone:    "123",
		Username: "kit",
		Nickname: "kitshy",
	}
	tokenBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	str := string(tokenBytes)
	fmt.Println("json====", str)

	var dao model.SysUser
	json.Unmarshal([]byte(str), &dao)

	fmt.Println("object=====", dao)

	key, err := GetRedis("key")
	if err != nil {
		println("err=====")
	}
	println("info======{}==", key)
}
