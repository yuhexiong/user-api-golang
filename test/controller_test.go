package test

import (
	"et-practice/controller"
	"et-practice/model"
	"sync"
	"testing"

	goDotENV "github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type argsCreateUser struct {
	user controller.CreateUserBody
}

type argsGetUser struct {
	userId string
}

type argsUpdateUser struct {
	userId      string
	user        model.User
	updatedUser *model.User
}

type argsUpdateUserStatus struct {
	userId      string
	status      int
	updatedUser *model.User
}

type argsDeleteUser struct {
	userId string
}

var newUserId string
var newUser = controller.CreateUserBody{
	Sort:     controller.GetPointer(uint64(0)),
	Email:    "test@test.com",
	SidNo:    "A123456789",
	Name:     "test",
	Gender:   "F",
	Birthday: "2023-10-04",
	Phone:    "0900000000",
}
var updateUser = model.User{
	Status:   controller.GetPointer(uint8(0)),
	Sort:     controller.GetPointer(uint64(0)),
	Email:    "test2@test.com",
	SidNo:    "B123456789",
	Name:     "test2",
	Gender:   "M",
	Birthday: "2023-10-05",
	Phone:    "0999999999",
}

func TestUser(t *testing.T) {
	if err := goDotENV.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	awaitRootElement := sync.WaitGroup{}
	awaitRootElement.Add(1)

	testCreateUser := struct {
		args    argsCreateUser
		wantErr bool
	}{
		args:    argsCreateUser{user: newUser},
		wantErr: false,
	}

	t.Run("測試新增使用者", func(t *testing.T) {
		got, err := controller.CreateUser(testCreateUser.args.user)
		if (err != nil) != testCreateUser.wantErr {
			t.Errorf("CreateUser() error = %v, wantErr %v", err, testCreateUser.wantErr)
			return
		}
		newUserId = got
		awaitRootElement.Done()
	})

	awaitRootElement.Wait()
	awaitRootElement.Add(1)

	testGetUser := struct {
		args    argsGetUser
		wantErr bool
	}{
		args:    argsGetUser{userId: newUserId},
		wantErr: false,
	}

	t.Run("測試取得使用者", func(t *testing.T) {
		_, err := controller.GetUser(testGetUser.args.userId)
		if (err != nil) != testGetUser.wantErr {
			t.Errorf("GetUser() error = %v, wantErr %v", err, testGetUser.wantErr)
			return
		}
		awaitRootElement.Done()
	})

	awaitRootElement.Wait()
	awaitRootElement.Add(1)

	var updatedUser *model.User
	testUpdateUser := struct {
		args    argsUpdateUser
		wantErr bool
	}{
		args: argsUpdateUser{
			userId:      newUserId,
			user:        updateUser,
			updatedUser: updatedUser,
		},
		wantErr: false,
	}
	t.Run("測試更新使用者", func(t *testing.T) {
		if err := controller.UpdateUser(testUpdateUser.args.userId, testUpdateUser.args.user, testUpdateUser.args.updatedUser); (err != nil) != testUpdateUser.wantErr {
			t.Errorf("UpdateUser() error = %v, wantErr %v", err, testUpdateUser.wantErr)
		}
		awaitRootElement.Done()
	})

	awaitRootElement.Wait()
	awaitRootElement.Add(1)

	testUpdateUserStatus := struct {
		args    argsUpdateUserStatus
		wantErr bool
	}{
		args: argsUpdateUserStatus{
			userId:      newUserId,
			status:      9,
			updatedUser: updatedUser,
		},
		wantErr: false,
	}
	t.Run("測試更新使用者狀態", func(t *testing.T) {
		if err := controller.UpdateUserStatus(testUpdateUserStatus.args.userId, testUpdateUserStatus.args.status, testUpdateUserStatus.args.updatedUser); (err != nil) != testUpdateUserStatus.wantErr {
			t.Errorf("UpdateUserStatus() error = %v, wantErr %v", err, testUpdateUserStatus.wantErr)
		}
		awaitRootElement.Done()
	})

	awaitRootElement.Wait()

	testDeleteUser := struct {
		args    argsDeleteUser
		wantErr bool
	}{
		args:    argsDeleteUser{userId: newUserId},
		wantErr: false,
	}
	t.Run("測試刪除使用者", func(t *testing.T) {
		err := controller.DeleteUser(testDeleteUser.args.userId)
		if (err != nil) != testDeleteUser.wantErr {
			t.Errorf("DeleteUser() error = %v, wantErr %v", err, testDeleteUser.wantErr)
			return
		}
	})
}
