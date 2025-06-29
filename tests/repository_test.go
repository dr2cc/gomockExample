package mocks

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/zmey56/gomock/internal/mocks"
	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"

	_ "github.com/lib/pq"
)

func TestGetUserByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the UserRepository interface
	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Set the expected behavior:
	// when GetUserByID is called with "1", return a user.
	// Установим ожидаемое поведение (и передадим ожидаемые значения):
	// когда GetUserByID вызывается с параметром "1", возвращается имя пользователя.
	mockRepo.EXPECT().
		GetUserByID("1").
		Return(&models.User{ID: "1", Name: "Alex"}, nil)

	// serviceTest := service.NewUserService(mockRepo)
	// user, err := serviceTest.GetUser("1")

	// Я переделал в одну строчку. Мне так понятнее, а сути не меняет
	// Здесь вызов тестируемого метода GetUser
	// Если параметр "1", то все будет в порядке.
	user, err := service.NewUserService(mockRepo).GetUser("1")

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

	mockRepo.EXPECT().
		DeleteUser("1").
		Return(nil) //.Times(1)

	err := service.NewUserService(mockRepo).DeleteUser("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// // Ожидаем, что метод DeleteUser будет вызван один раз с параметром "1"
	// mockRepo.EXPECT().
	// 	DeleteUser("1").
	// 	Return(nil).
	// 	Times(1)

	// // service := service.NewUserService(mockRepo)
	// // err := service.DeleteUser("1")
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

func TestGetUserByID_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().
		GetUserByID("1").
		Return(nil, errors.New("user not found")).
		Times(1)

	service := service.NewUserService(mockRepo)
	user, err := service.GetUser("1")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetUserByID_WithDoCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().
		GetUserByID("1").
		Do(func(id string) {
			fmt.Printf("Getting user with ID: %s\n", id)
		}).
		Return(&models.User{ID: "1", Name: "Alex"}, nil)

	service := service.NewUserService(mockRepo)
	user, err := service.GetUser("1")

	assert.NoError(t, err)
	assert.Equal(t, "1", user.ID)
}

// // Все тесты выше работают!
//
// // 12.06.2025 Видимо не дописанный тест.
// // Хорошая практика- дописать его!!
// // Видимо в type UserRepository interface
// // нужно добавить UpdateUser
// // а в user_service.go надо добавить метод UpdateUserAsync
// func TestUpdateUserAsync(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mocks.NewMockUserRepository(ctrl)
// 	wg := sync.WaitGroup{}

// 	mockRepo.EXPECT().
// 		// видимо нужно добавить метод UpdateUser в интерфейс UserRepository
// 		UpdateUser("1", "Updated Name").
// 		Do(func(id, name string) {
// 			wg.Done()
// 		}).
// 		Return(nil).
// 		Times(1)

// 	service := service.NewUserService(mockRepo)

// 	wg.Add(1)
// 	// видимо в user_service.go надо добавить метод UpdateUserAsync
// 	go service.UpdateUserAsync("1", "Updated Name")

// 	wg.Wait() // ждем завершения goroutine
// }

// // 12.06.2025 Этот не проходит, ругается на Docker
// func TestUserRepositoryWithPostgres(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
// 	defer cancel()

// 	log.Println("Starting PostgreSQL container for test...")

// 	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: testcontainers.ContainerRequest{
// 			Image:        "postgres:latest",
// 			ExposedPorts: []string{"5432/tcp"},
// 			Env: map[string]string{
// 				"POSTGRES_USER":     "user",
// 				"POSTGRES_PASSWORD": "password",
// 				"POSTGRES_DB":       "testdb",
// 			},
// 			WaitingFor: wait.ForLog("database system is ready to accept connections"),
// 		},
// 		Started: true,
// 	})
// 	if err != nil {
// 		t.Fatalf("Failed to start PostgreSQL container: %v", err)
// 	}
// 	defer func() {
// 		log.Println("Stopping PostgreSQL container...")
// 		postgresContainer.Terminate(ctx)
// 	}()

// 	log.Println("PostgreSQL container started successfully.")

// 	// add a delay to make sure that the database is fully ready
// 	time.Sleep(2 * time.Second)

// 	// getting the port to connect to the container
// 	port, err := postgresContainer.MappedPort(ctx, "5432")
// 	if err != nil {
// 		t.Fatalf("Failed to get container port: %v", err)
// 	}

// 	// use 127.0.0.1 to connect
// 	dsn := fmt.Sprintf("postgres://user:password@127.0.0.1:%s/testdb?sslmode=disable", port.Port())

// 	log.Printf("Connecting to PostgreSQL at %s...", dsn)

// 	// connecting to the database using dsn
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
// 	}
// 	defer func() {
// 		log.Println("Closing database connection...")
// 		db.Close()
// 	}()

// 	retries := 5
// 	for retries > 0 {
// 		if err = db.Ping(); err == nil {
// 			break
// 		}
// 		retries--
// 		log.Printf("Retrying to connect to the database. Retries left: %d", retries)
// 		time.Sleep(2 * time.Second)
// 	}

// 	if err != nil {
// 		t.Fatalf("Failed to ping database after retries: %v", err)
// 	}

// 	log.Println("Successfully connected to PostgreSQL database.")
// }

// 12.06.2025 этот тест проходит!
func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name       string
		userID     string
		mockReturn *models.User
		mockError  error
		expected   *models.User
		expectErr  bool
	}{
		{
			name:       "Success",
			userID:     "1",
			mockReturn: &models.User{ID: "1", Name: "Alex"},
			mockError:  nil,
			expected:   &models.User{ID: "1", Name: "Alex"},
			expectErr:  false,
		},
		{
			name:       "UserNotFound",
			userID:     "2",
			mockReturn: nil,
			mockError:  errors.New("user not found"),
			expected:   nil,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				GetUserByID(tt.userID).
				Return(tt.mockReturn, tt.mockError).
				Times(1)

			service := service.NewUserService(mockRepo)
			result, err := service.GetUser(tt.userID)

			if tt.expectErr && err == nil {
				t.Errorf("expected error, got none")
			}

			if !tt.expectErr && err != nil {
				t.Errorf("did not expect error, got %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
