package e

const (
	// redis 存储 token前缀
	REDIS_USER_TOKEN = "USER_TOKEN_"
	// redis 存储 刷新token前缀
	REDIS_USER_REFRESH_TOKEN = "USER_REFRESH_TOKEN_"
	// redis存储 发送登录短信前缀
	REDIS_SMS_LOGIN_PHONE = "SME_LOGIN_PHONE_"
	// redis存储 发送邮箱前缀
	REDIS_SMS_LOGIN_EMAIL = "SMS_LOGIN_EMAIL_"
	// 首页 分类导航列表 缓存
	REDIS_HOME_CATEGORY_NAV_LIST_CACHE = "REDIS_HOME_CATEGORY_NAV_LIST_CACHE_"
)
