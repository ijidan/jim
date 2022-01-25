package request

type UserRegisterRequest struct {
	Nickname    string `form:"nickname" json:"nickname" binding:"required,min=2,max=30" `
	Password    string `form:"password" json:"password" binding:"required,min=6,max=18,alphanum"`
	PasswordRpt string `form:"password_rpt" json:"password_rpt" binding:"required,min=6,max=18,alphanum,eqfield=password"`
}

type UserLoginRequest struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30" `
	Password string `form:"password" json:"password" binding:"required,min=6,max=18,alphanum"`
}

type UserGetRequest struct {
	Id uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
}

type UserQueryRequest struct {
	Keyword  string `form:"keyword" json:"keyword" binding:"omitempty,required,min=1,max=30" `
	Page     uint64 `form:"page" json:"page" binding:"omitempty,required,min=1,gte=1" `
	PageSize uint64 `form:"page_size" json:"page_size" binding:"omitempty,required,min=1,gte=1" `
}

type UserUpdateAvatar struct {
	AvatarUrl  string `form:"avatar_url" json:"avatar_url" binding:"required,url,endswith=jpg|jpeg|png" `
}

type UserUpdatePassword struct {
	Password    string `form:"password" json:"password" binding:"required,min=6,max=18,alphanum"`
	PasswordRpt string `form:"password_rpt" json:"password_rpt" binding:"required,min=6,max=18,alphanum,eqfield=password"`
}