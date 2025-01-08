package service

type LoginPassWordRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginPassWordResponse struct {
	Toekn        string `json:"toekn"`
	RefreshToken string `json:"refreshToken"`
}
