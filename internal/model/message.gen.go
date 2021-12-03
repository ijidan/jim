// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SenderID     int32          `gorm:"column:sender_id;not null;default:0" json:"sender_id"`         // 发送人ID
	SenderType   int32          `gorm:"column:sender_type;not null;default:0" json:"sender_type"`     // 发送人类型
	ReceiverID   int32          `gorm:"column:receiver_id;not null;default:0" json:"receiver_id"`     // 接收人ID，单聊则为user_id,群聊则为group_id
	ReceiverType int32          `gorm:"column:receiver_type;not null;default:0" json:"receiver_type"` // 接收人类型
	ToUserIds    string         `gorm:"column:to_user_ids;not null" json:"to_user_ids"`               // 需要@的用户，多个用,分割
	MessageType  int32          `gorm:"column:message_type;not null;default:0" json:"message_type"`   // 消息类型
	Status       int32          `gorm:"column:status;not null;default:1" json:"status"`               // 状态：1-正常，2-撤回，0-删除
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`                          // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`                          // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                          // 删除时间
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
