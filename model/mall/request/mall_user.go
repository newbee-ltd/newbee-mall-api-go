package request

//用户注册
type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

//更新用户信息
type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}
