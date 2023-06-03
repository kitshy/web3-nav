package model

/*
*
用户
*/
type SysUser struct {
	BaseModel
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	ImageUrl string `json:"imageUrl"`
	Type     string `json:"type"` //  0 普通用户  1 管理员
}

func GetAccountByUsernamePhoneEmail(username string) (sysUser *SysUser, err error) {
	err = db.Model(&SysUser{}).Where("is_deleted = 0").Where(
		db.Where("username = ?", username).Or("phone = ?", username).Or("email = ?", username),
	).Scan(&sysUser).Error
	if err != nil {
		return nil, err
	}
	return sysUser, err
}

func GetUserByUsername(username string) (sysUser *SysUser, err error) {
	err = db.Model(&SysUser{}).Where("is_deleted = 0 and username = ?", username).Scan(&sysUser).Error
	return sysUser, err
}

func GetUserByEmail(email string) (sysUser *SysUser, err error) {
	err = db.Model(&SysUser{}).Where("is_deleted = 0 and email = ?", email).Scan(&sysUser).Error
	return sysUser, err
}

func GetUserByPhone(phone string) (sysUser *SysUser, err error) {
	err = db.Model(&SysUser{}).Where("is_deleted = 0 and phone = ?", phone).Scan(&sysUser).Error
	return sysUser, err
}

func GetUserById(id int) (sysUser *SysUser, err error) {
	err = db.Model(&SysUser{}).Where("is_deleted = 0 and id = ?", id).Scan(&sysUser).Error
	if err != nil {
		return nil, err
	}
	return sysUser, err
}

/*
*
新增用户
*/
func AddUser(user SysUser) error {
	return db.Create(&user).Error
}

/*
*
编辑用户
*/
func EditUser(user SysUser) error {
	return db.Model(&user).Updates(&user).Error
}
