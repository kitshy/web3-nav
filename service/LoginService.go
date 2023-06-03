package service

import (
	"encoding/json"
	"ginweb/common"
	"ginweb/common/e"
	"ginweb/dto"
	"ginweb/model"
	"ginweb/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

/*
*
账户密码登录
*/
func UserPassword(auth dto.LoginAuth) (*dto.AuthDto, error) {
	user := GetAccountByUsernamePhoneEmail(auth)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		panic(common.Response{
			Code: e.ERROR,
			Msg:  "密码错误",
		})
	}
	tokenBytes, _ := json.Marshal(user)
	jsonStr := string(tokenBytes)
	token := utils.AesEny(tokenBytes)
	expire := 60 * 60
	utils.SetRedis(token, jsonStr, expire)
	return &dto.AuthDto{
		Token:    token,
		ExpireAt: expire,
	}, nil
}

/*
*
手机号验证码登录
*/
func PhoneCode(auth dto.LoginAuth) (*dto.AuthDto, error) {
	var key strings.Builder
	key.WriteString(e.REDIS_SMS_LOGIN_PHONE)
	key.WriteString(auth.Phone)
	value, err := utils.GetRedis(key.String())
	if err != nil {
		panic(common.Response{
			Code: e.ERROR,
			Msg:  "手机号验证码错误",
		})
	}
	if value == auth.Code {
		user := GetAccountByUsernamePhoneEmail(auth)
		tokenBytes, _ := json.Marshal(user)
		jsonStr := string(tokenBytes)
		token := utils.AesEny(tokenBytes)
		expire := 60 * 60
		utils.SetRedis(token, jsonStr, expire)
		return &dto.AuthDto{
			Token:    token,
			ExpireAt: expire,
		}, nil
	}
	return nil, err
}

/*
*
邮箱验证码登录
*/
func EmailCode(auth dto.LoginAuth) (*dto.AuthDto, error) {
	var key strings.Builder
	key.WriteString(e.REDIS_SMS_LOGIN_EMAIL)
	key.WriteString(auth.Email)
	value, err := utils.GetRedis(key.String())
	if err != nil {
		panic(common.Response{
			Code: e.ERROR,
			Msg:  "邮箱验证码错误",
		})
	}
	if value == auth.Code {
		user := GetAccountByUsernamePhoneEmail(auth)
		tokenBytes, _ := json.Marshal(user)
		jsonStr := string(tokenBytes)
		token := utils.AesEny(tokenBytes)
		expire := 60 * 60
		utils.SetRedis(token, jsonStr, expire)
		return &dto.AuthDto{
			Token:    token,
			ExpireAt: expire,
		}, nil
	}
	return nil, err
}

/*
*
获取账户
*/
func GetAccountByUsernamePhoneEmail(authDto dto.LoginAuth) (user *model.SysUser) {
	//  获取用户
	us, err := model.GetAccountByUsernamePhoneEmail(authDto.Username)
	if err != nil {
		panic(&common.Response{
			Code: e.ERROR,
			Msg:  "稍后在试",
		})
	}
	if us == nil {
		panic(&common.Response{
			Code: e.ERROR,
			Msg:  "账户非法",
		})
	}
	return us
}
