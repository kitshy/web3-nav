package dto

/*
*
用户
*/
type UserDto struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"NewPassword"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	ImageUrl    string `json:"imageUrl"`
	Type        string `json:"type"`
}
