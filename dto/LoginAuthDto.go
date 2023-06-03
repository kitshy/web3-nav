package dto

/*
*
登录
*/
type LoginAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	// 1 账户密码 2 手机验证 3 邮箱
	Method string `json:"method"`
	Code   string `json:"code"`
}
