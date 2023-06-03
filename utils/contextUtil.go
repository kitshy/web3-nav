package utils

import (
	"encoding/json"
	"ginweb/model"
	"golang.org/x/net/context"
)

const ContextKey = "ContextKey"

/*
*
创建上下文
*/
func NewContext(ctx context.Context, user *model.SysUser) context.Context {
	return context.WithValue(ctx, ContextKey, user)
}

/*
*
从上下文获取登录用户
*/
func GetCurrentUserInContext(ctx context.Context) (*model.SysUser, bool) {
	u, ok := ctx.Value(ContextKey).(*model.SysUser)
	return u, ok
}

/*
*
从redis获取登录用户
*/
func GetCurrentUserInRedis(token string) (*model.SysUser, error) {
	str, err := GetRedis(token)
	if err != nil {
		return nil, err
	}
	var user model.SysUser
	json.Unmarshal([]byte(str), &user)
	return &user, err
}
