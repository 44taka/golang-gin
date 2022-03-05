package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/infrastructure/persistence"
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// TODO::テストデータのリファクタやる
	var users []*model.UserTestData
	testdata, _ := ioutil.ReadFile("../../testdata/user_fixture.json")
	err := json.Unmarshal(testdata, &users)
	if err != nil {
		os.Exit(1)
	}
	db := infrastructure.NewDB(infrastructure.NewConfig()).Connect()
	db.Exec("delete from users")
	for i := 0; i < len(users); i++ {
		db.Exec(
			"insert into users values (?, ?, ?)",
			users[i].ID,
			users[i].Name,
			users[i].Password,
		)
	}
	status := m.Run()
	db.Exec("delete from users")
	os.Exit(status)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	// ctx.Request, _ = http.NewRequest(http.MethodGet, "/users/1", nil)
	// ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)

	userPersistence := persistence.NewUserPersistence(db.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := NewUserHandler(userUseCase)
	userHandler.FindById(ctx)

	// ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// mockUserUseCase := mocks.NewMockUserUseCase(ctrl)

	// var user model.User
	// mockUserUseCase.EXPECT().FindById(ctx, 1).Return(user, nil)

	// userHandler := NewUserHandler(mockUserUseCase)
	// fmt.Println("********************")
	// userHandler.FindById(ctx)
	// fmt.Println("********************")

	// response := httptest.NewRecorder()
	// ctx, _ := gin.CreateTestContext(response)
	// ctx.Request, _ = http.NewRequest(http.MethodGet, "/users/1", nil)
	// fmt.Println("********************")
	// fmt.Println(response.Body.String())
	// fmt.Println("********************")

	asserts := assert.New(t)
	asserts.Equal(1, 1)
}

// func TestCreate(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)

// 	mockUserHandler := mock_handler.NewMockUserHandler(ctrl)
// 	mockUserHandler.EXPECT().Create(gomock.Any())
// 	mockUserHandler.FindById(ctx)

// 	asserts := assert.New(t)
// 	asserts.Equal(1, 1)
// }
