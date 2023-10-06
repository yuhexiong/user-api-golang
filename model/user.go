package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"id,omitempty" json:"id" example:"623853b9503ce2ecdd221c94"`
	Status    *uint8             `bson:"status,omitempty" json:"status" example:"0"` // 0: 正常, 1: 停用, 2: 暫存, 9: 刪除
	Sort      *uint64            `bson:"sort,omitempty" json:"sort" example:"0"`     // 排序
	CreatedAt *time.Time         `bson:"createdAt,omitempty" json:"createdAt" example:"2022-03-21T10:30:17.711Z"`
	UpdatedAt *time.Time         `bson:"updatedAt,omitempty" json:"updatedAt" example:"2022-03-21T10:30:17.711Z"`
	Email     string             `bson:"email,omitempty" json:"email" example:"hygeai@hixcare.com"` // 電子郵件
	SidNo     string             `bson:"sidNo,omitempty" json:"sidNo" example:"A1234567890"`        // 身分證字號
	Name      string             `bson:"name,omitempty" json:"name" example:"艾瑞克"`                  // 姓名
	Gender    string             `bson:"gender,omitempty" json:"gender" example:"M"`                // M=男, F=女 (大寫)
	Birthday  string             `bson:"birthday,omitempty" json:"birthday" example:"2000-02-29"`   // 生日
	Phone     string             `bson:"phone,omitempty" json:"phone" example:"0900000000"`         // 電話
}
