package service

import (
	"ginweb/common"
	"ginweb/model"
)

/*
*
获取账户通过id
*/
func GetUserById(id int) (sysUser *model.SysUser, err error) {
	return model.GetUserById(id)
}

/*
*
新增用户
*/
func AddUser(user model.SysUser) bool {
	err := model.AddUser(user)
	if err != nil {
		return false
	}
	return true
}

/*
*
编辑用户
*/
func EditUser(user model.SysUser) bool {
	err := model.EditUser(user)
	if err != nil {
		return false
	}
	return true
}

/*
*
检查用户名没有重复
*/
func CheckIsNotExitUsername(username string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		panic(&common.Response{})
	}
	if user == nil {
		return true
	}
	return false
}

/*
*
检查手机号 是否存在 不存在返回 1
*/
func CheckIsNotExitPhone(phone string) bool {
	user, err := model.GetUserByPhone(phone)
	if err != nil {
		panic(&common.Response{})
	}
	if user == nil {
		return true
	}
	return false
}

/*
*
检查邮箱是否存在 不存在 返回 1
*/
func CheckIsNotExitEmail(email string) bool {
	user, err := model.GetUserByEmail(email)
	if err != nil {
		panic(&common.Response{})
	}
	if user == nil {
		return true
	}
	return false
}

/*
*
检查是否存在手机号邮箱
*/
func CheckIsNotExitPhoneAndEmail(phone string, email string) (bool, string) {
	if phone != "" {
		return CheckIsNotExitPhone(phone), "手机号已存在"
	}
	if email != "" {
		return CheckIsNotExitEmail(email), "邮箱已存在"
	}
	return true, ""
}
