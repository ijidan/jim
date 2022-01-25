package request

type MessageSendUserTextMessageRequest struct {
	ToUserId uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Content  string `form:"content" json:"content" binding:"omitempty,min=1,max=100"`
}

type MessageSendUserLocationMessageRequest struct {
	ToUserId   uint64  `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	CoverImage string  `form:"cover_image" json:"cover_image" binding:"required,url,endswith=jpg|jpeg|png" `
	Lat        float64 `form:"lat" json:"lat" binding:"required"`
	Lng        float64 `form:"lng" json:"lng" binding:"required"`
	MapLink    string  `form:"map_link" json:"map_link" binding:"required,url" `
	Desc       string  `form:"desc" json:"desc" binding:"required,url,max=100" `
}

type MessageSendUserFaceMessageRequest struct {
	ToUserId uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Symbol   string `form:"symbol" json:"symbol" binding:"required,max=20" `
}

type MessageSendUserSoundMessageRequest struct {
	ToUserId uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Url      string `form:"url" json:"url" binding:"required,url" `
	Size     uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Seconds  uint64 `form:"seconds" json:"seconds" binding:"required,numeric,min=1,gte=1" `
}

type MessageSendUserVideoMessageRequest struct {
	ToUserId    uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Size        uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Seconds     uint64 `form:"seconds" json:"seconds" binding:"required,numeric,min=1,gte=1" `
	Url         string `form:"url" json:"url" binding:"required,url" `
	Format      string `form:"format" json:"format" binding:"required" `
	ThumbUrl    string `form:"thumb_url" json:"thumb_url" binding:"required,url" `
	ThumbSize   uint64 `form:"thumb_size" json:"thumb_size" binding:"required,numeric,min=1,gte=1" `
	ThumbWidth  float64 `form:"thumb_width" json:"thumb_width" binding:"required,numeric,min=1,gte=1" `
	ThumbHeight float64 `form:"thumb_height" json:"thumb_height" binding:"required,numeric,min=1,gte=1" `
	ThumbFormat uint64 `form:"thumb_format" json:"thumb_format" binding:"required,oneof=0 1 2 3 255" `
}

type MessageSendUserImageMessageRequest struct {
	ToUserId uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Type     uint64 `form:"type" json:"type" binding:"required,oneof=0 1 2 " `
	Format   uint64 `form:"format" json:"format" binding:"required,oneof=0 1 2 3 255" `
	Size     uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Width    uint64 `form:"width" json:"width" binding:"required,numeric,min=1,gte=1" `
	Height   uint64 `form:"height" json:"height" binding:"required,numeric,min=1,gte=1" `
	Url      string `form:"url" json:"url" binding:"required,url,endswith=jpg|jpeg|png" `
}

type MessageSendUserFileMessageRequest struct {
	ToUserId uint64 `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1,gte=1" `
	Size     uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Name     string `form:"name" json:"name" binding:"required,min=2,max=50" `
	Url      string `form:"url" json:"url" binding:"required,url" `
}

type MessageSendGroupTextMessageRequest struct {
	ToGroupId uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	Content   string `form:"content" json:"content" binding:"omitempty,min=1,max=100"`
}

type MessageSendGroupLocationMessageRequest struct {
	ToGroupId  uint64  `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	CoverImage string  `form:"cover_image" json:"cover_image" binding:"required,url,endswith=jpg|jpeg|png" `
	Lat        float64 `form:"lat" json:"lat" binding:"required"`
	Lng        float64 `form:"lng" json:"lng" binding:"required"`
	MapLink    string  `form:"map_link" json:"map_link" binding:"required,url" `
	Desc       string  `form:"desc" json:"desc" binding:"required,url,max=100" `
}

type MessageSendGroupFceMessageRequest struct {
	ToGroupId uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	Symbol    string `form:"symbol" json:"symbol" binding:"required,max=20" `
}

type MessageSendGroupSoundMessageRequest struct {
	ToGroupId uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  uint64 `form:"at_user_id" json:"at_user_id" binding:"required,numeric,min=1,gte=1" `
	Url       string `form:"url" json:"url" binding:"required,url" `
	Size      uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Seconds   uint64 `form:"seconds" json:"seconds" binding:"required,numeric,min=1,gte=1" `
}

type MessageSendGroupVideoMessageRequest struct {
	ToGroupId   uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	Size        uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Url         string `form:"url" json:"url" binding:"required,url" `
	Format      string `form:"format" json:"format" binding:"required" `
	ThumbUrl    string `form:"thumb_url" json:"thumb_url" binding:"required,url" `
	ThumbSize   uint64 `form:"thumb_size" json:"thumb_size" binding:"required,numeric,min=1,gte=1" `
	ThumbWidth  uint64 `form:"thumb_width" json:"thumb_width" binding:"required,numeric,min=1,gte=1" `
	ThumbHeight uint64 `form:"thumb_height" json:"thumb_height" binding:"required,numeric,min=1,gte=1" `
	ThumbFormat uint64 `form:"thumb_format" json:"thumb_format" binding:"required,oneof=0 1 2 3 255" `
}

type MessageSendGroupImageMessageRequest struct {
	ToGroupId uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	Type      uint64 `form:"type" json:"type" binding:"required,oneof=0 1 2 " `
	Format    uint64 `form:"format" json:"format" binding:"required,oneof=0 1 2 3 255" `
	Size      uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Width     uint64 `form:"width" json:"width" binding:"required,numeric,min=1,gte=1" `
	Height    uint64 `form:"height" json:"height" binding:"required,numeric,min=1,gte=1" `
	Url       string `form:"url" json:"url" binding:"required,url,endswith=jpg|jpeg|png" `
}

type MessageSendGroupFileMessageRequest struct {
	ToGroupId uint64 `form:"to_group_id" json:"to_group_id" binding:"required,numeric,min=1,gte=1" `
	AtUserId  []uint64 `form:"at_user_id" json:"at_user_id" binding:"required" `
	Size      uint64 `form:"size" json:"size" binding:"required,numeric,min=1,gte=1" `
	Name      string `form:"name" json:"name" binding:"required,min=2,max=50" `
	Url       string `form:"url" json:"url" binding:"required,url" `
}
