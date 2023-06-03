package dto

type AuthDto struct {
	Token           string `json:"token"`
	ExpireAt        int    `json:"expireAt"`
	RefreshToken    string `json:"refreshToken"`
	RefreshExpireAt int    `json:"refreshExpireAt"`
}
