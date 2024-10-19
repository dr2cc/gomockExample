package mocks

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/zmey56/gomock/internal/mocks"
	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

func TestGetUserByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the UserRepository interface
	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Set the expected behavior: when GetUserByID is called with "1", return a user.
	mockRepo.EXPECT().
		GetUserByID("1").
		Return(&models.User{ID: "1", Name: "Alex"}, nil)

	service := service.NewUserService(mockRepo)

	user, err := service.GetUser("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if user.ID != "1" {
		t.Errorf("Expected user ID to be '1', got %s", user.ID)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Ожидаем, что метод DeleteUser будет вызван один раз с параметром "1"
	mockRepo.EXPECT().
		DeleteUser("1").
		Return(nil).
		Times(1)

	service := service.NewUserService(mockRepo)
	err := service.DeleteUser("1")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteUserWorkflow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	gomock.InOrder(
		mockRepo.EXPECT().GetUserByID("1").Return(&models.User{ID: "1", Name: "Alex"}, nil),
		mockRepo.EXPECT().DeleteUser("1").Return(nil),
	)

	service := service.NewUserService(mockRepo)

	user, err := service.GetUser("1")
	if err != nil || user.ID != "1" {
		t.Fatalf("Expected to find user with ID 1, got error %v", err)
	}

	err = service.DeleteUser("1")
	if err != nil {
		t.Fatalf("Expected no error when deleting user, got %v", err)
	}
}
