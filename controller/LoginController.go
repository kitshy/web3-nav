package controller

import (
	"ginweb/common"
	"ginweb/dto"
	"ginweb/model"
	"ginweb/service"
	"ginweb/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*
*
用户登录
*/
func Login(c *gin.Context) {
	rs := common.Gin{C: c}

	var authDto dto.LoginAuth
	c.ShouldBindJSON(&authDto)

	// 账户密码登录
	if authDto.Method == "1" {
		if authDto.Password == "" || (authDto.Phone == "" && authDto.Username == "" && authDto.Email == "") {
			rs.ResponseFail("请输入账户密码", nil)
			return
		}
		auth, err := service.UserPassword(authDto)
		if err != nil {
			rs.ResponseFail("稍后在试", nil)
			return
		}
		rs.ResponseSuccess(auth)
		return
	}

	// 手机号登录
	if authDto.Method == "2" {
		if authDto.Code == "" || authDto.Phone == "" {
			rs.ResponseFail("请输入手机号验证码", nil)
			return
		}
		auth, err := service.PhoneCode(authDto)
		if err != nil {
			rs.ResponseFail("稍后在试", nil)
		}
		rs.ResponseSuccess(auth)
		return
	}

	// 邮箱登录
	if authDto.Method == "3" {
		if authDto.Code == "" || authDto.Email == "" {
			rs.ResponseFail("请输入邮箱验证码", nil)
			return
		}
		auth, err := service.EmailCode(authDto)
		if err != nil {
			rs.ResponseFail("稍后在试", nil)
		}
		rs.ResponseSuccess(auth)
		return
	}

	rs.ResponseFail("非法登录", nil)

}

/*
*
获取用户信息
*/
func GetUserInfo(c *gin.Context) {
	rs := common.Gin{C: c}
	user, ok := utils.GetCurrentUserInContext(c.Request.Context())
	if !ok {
		rs.ResponseFail("获取用户信息错误", nil)
		return
	}
	user, err := service.GetUserById(user.Id)
	if err != nil || user == nil {
		rs.ResponseFail("获取用户信息错误", nil)
		return
	}
	rs.ResponseSuccess(dto.UserDto{
		Id:       user.Id,
		Nickname: user.Nickname,
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
	})
}

/*
*
新增用户
*/
func AddUser(c *gin.Context) {
	rs := common.Gin{C: c}

	var user model.SysUser
	c.ShouldBindJSON(&user)
	if user.Nickname == "" || user.Password == "" || user.Username == "" {
		rs.ResponseFail("昵称-用户名-密码不能为空", nil)
		return
	}

	/**
	检查用户名是否存在
	*/
	if !service.CheckIsNotExitUsername(user.Username) {
		rs.ResponseFail("用户名已被注册", nil)
		return
	}

	// 检查邮箱，手机号是否存在
	ok, msg := service.CheckIsNotExitPhoneAndEmail(user.Phone, user.Email)
	if !ok {
		rs.ResponseFail(msg, nil)
		return
	}

	//密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		rs.ResponseFail("fail", nil)
		return
	}
	user.Password = string(password)
	user.Type = "0"
	ok = service.AddUser(user)
	if !ok {
		rs.ResponseFail("新增失败", nil)
		return
	}
	rs.ResponseSuccess("success")
	return
}

/*
*
编辑用户  , 密码 - 昵称 - 图像
*/
func EditUser(c *gin.Context) {
	rs := common.Gin{C: c}

	var dto dto.UserDto
	c.ShouldBindJSON(&dto)

	// 获取当前登录用户信息 编辑
	sysUser, ok := utils.GetCurrentUserInContext(c.Request.Context())
	if !ok {
		rs.ResponseFail("获取用户信息错误", nil)
		return
	}
	var user = model.SysUser{
		Nickname: dto.Nickname,
		ImageUrl: dto.ImageUrl,
	}
	user.Id = sysUser.Id

	// 更新密码   需要原密码和新密码
	if dto.NewPassword != "" {

		// 验证原密码
		err := bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(dto.Password))
		if err != nil {
			rs.ResponseFail("原密码错误", nil)
			return
		}

		// 加密新密码
		password, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}
		user.Password = string(password)
	}

	ok = service.EditUser(user)
	if !ok {
		rs.ResponseFail("请重新编辑信息", nil)
		return
	}
	rs.ResponseSuccess("success")
	return
}

/*
*
新增编辑用户
*/
func AddEditUser(c *gin.Context) {

	rs := common.Gin{C: c}

	var dto dto.UserDto
	c.ShouldBindJSON(&dto)

	if dto.Id == 0 {
		if dto.Nickname == "" || dto.Password == "" || dto.Username == "" {
			rs.ResponseFail("昵称-用户名-密码不能为空", nil)
			return
		}

		/**
		检查用户名是否存在
		*/
		if !service.CheckIsNotExitUsername(dto.Username) {
			rs.ResponseFail("用户名已被注册", nil)
			return
		}

		// 检查邮箱，手机号是否存在
		ok, msg := service.CheckIsNotExitPhoneAndEmail(dto.Phone, dto.Email)
		if !ok {
			rs.ResponseFail(msg, nil)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			rs.ResponseFail("fail", nil)
			return
		}

		var user = model.SysUser{
			Nickname: dto.Nickname,
			Username: dto.Username,
			Password: string(password),
			ImageUrl: dto.ImageUrl,
			Phone:    dto.Phone,
			Email:    dto.Email,
			Type:     "1",
		}
		ok = service.AddUser(user)
		if !ok {
			rs.ResponseFail("新增失败", nil)
			return
		}
	} else {

		var user = model.SysUser{
			Nickname: dto.Nickname,
			ImageUrl: dto.ImageUrl,
		}
		// 更新密码
		if dto.NewPassword != "" {
			password, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
			if err != nil {
				rs.ResponseFail("fail", nil)
				return
			}
			user.Password = string(password)
		}
		user.Id = dto.Id

		ok := service.EditUser(user)
		if !ok {
			rs.ResponseFail("请重新编辑信息", nil)
			return
		}
	}

	rs.ResponseSuccess("success")
	return
}
