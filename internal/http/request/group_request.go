package request

type GroupCreateRequest struct {
	Name         string `form:"name" json:"name" binding:"required,min=2,max=30" `
	Introduction string `form:"introduction" json:"introduction" binding:"omitempty,min=1,max=30"`
	AvatarUrl    string `form:"avatar_url" json:"avatar_url" binding:"required,url,max=30,endswith=jpg|jpeg|png" `
}

type GroupUpdateRequest struct {
	Id           uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
	Name         string `form:"name" json:"name" binding:"required,min=2,max=30" `
	Introduction string `form:"introduction" json:"introduction" binding:"omitempty,min=1,max=30"`
	AvatarUrl    string `form:"avatar_url" json:"avatar_url" binding:"required,url,max=30,endswith=jpg|jpeg|png" `
}

type GroupDeleteRequest struct {
	Id uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
}

type GroupJoinRequest struct {
	Id     uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
	UserId uint64 `form:"user_id" json:"user_id" binding:"required,numeric,min=1,gte=1" `
}

type GroupQuitRequest struct {
	Id     uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
	UserId uint64 `form:"user_id" json:"user_id" binding:"required,numeric,min=1,gte=1" `
}

type GroupGetRequest struct {
	Id uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
}

type GroupQueryRequest struct {
	Keyword  string `form:"keyword" json:"keyword" binding:"omitempty,required,min=1,max=30" `
	Page     uint64 `form:"page" json:"page" binding:"omitempty,required,min=1,gte=1" `
	PageSize uint64 `form:"page_size" json:"page_size" binding:"omitempty,required,min=1,gte=1" `
}

type GroupGetUserRequest struct {
	Id       uint64 `form:"id" json:"id" binding:"required,numeric,min=1,gte=1" `
	Page     uint64 `form:"page" json:"page" binding:"omitempty,required,min=1,gte=1" `
	PageSize uint64 `form:"page_size" json:"page_size" binding:"omitempty,required,min=1,gte=1" `
}
