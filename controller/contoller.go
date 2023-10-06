package controller

import (
	"context"
	"errors"
	"et-practice/config"
	"et-practice/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserBody struct {
	Sort     *uint64 `bson:"sort,omitempty" json:"sort" example:"0"`                    // 排序
	Email    string  `bson:"email,omitempty" json:"email" example:"hygeai@hixcare.com"` // 電子郵件
	SidNo    string  `bson:"sidNo,omitempty" json:"sidNo" example:"A1234567890"`        // 身分證字號
	Name     string  `bson:"name,omitempty" json:"name" example:"艾瑞克"`                  // 姓名
	Gender   string  `bson:"gender,omitempty" json:"gender" example:"M"`                // M=男, F=女 (大寫)
	Birthday string  `bson:"birthday,omitempty" json:"birthday" example:"2000-02-29"`   // 生日
	Phone    string  `bson:"phone,omitempty" json:"phone" example:"0900000000"`         // 電話
}

func GetPointer[T any](value T) *T {
	return &value
}

// 新增使用者
func CreateUser(user CreateUserBody) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser := model.User{
		ID:        primitive.NewObjectID(),
		Status:    GetPointer(uint8(0)),
		Sort:      user.Sort,
		CreatedAt: GetPointer(time.Now()),
		UpdatedAt: GetPointer(time.Now()),
		Email:     user.Email,
		SidNo:     user.SidNo,
		Name:      user.Name,
		Gender:    user.Gender,
		Birthday:  user.Birthday,
		Phone:     user.Phone,
	}

	_, err := config.GetCollection(config.GetDB(), "user").InsertOne(ctx, newUser)

	return newUser.ID.Hex(), err
}

// 取得使用者
func GetUser(userId string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.User
	objId, _ := primitive.ObjectIDFromHex(userId)

	err := config.GetCollection(config.GetDB(), "user").FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	return user, err
}

// 更新使用者
func UpdateUser(userId string, user model.User, updatedUser *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	update := bson.M{
		"updatedAt": GetPointer(time.Now()),
		"status":    user.Status,
		"sort":      user.Sort,
		"email":     user.Email,
		"sidNo":     user.SidNo,
		"name":      user.Name,
		"gender":    user.Gender,
		"birthday":  user.Birthday,
		"phone":     user.Phone,
	}

	result, err := config.GetCollection(config.GetDB(), "user").UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	//get updated user details
	if result.MatchedCount == 1 {
		err := config.GetCollection(config.GetDB(), "user").FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			return err
		}
	}

	return nil
}

// 更新使用者status
func UpdateUserStatus(userId string, status int, updatedUser *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	update := bson.M{"status": status}

	result, err := config.GetCollection(config.GetDB(), "user").UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	//get updated user details
	if result.MatchedCount < 1 {
		return errors.New("user with specified ID not found")
	}

	return nil
}

// 刪除使用者
func DeleteUser(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	result, err := config.GetCollection(config.GetDB(), "user").DeleteOne(ctx, bson.M{"id": objId})

	if result.DeletedCount < 1 {
		return errors.New("user with specified ID not found")
	}

	return err
}
