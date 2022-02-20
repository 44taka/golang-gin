package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	mockUserUseCase := mocks.NewMockUserUseCase(ctrl)

	var user model.User
	mockUserUseCase.EXPECT().FindById(ctx, 1).Return(user, nil)

	userHandler := NewUserHandler(mockUserUseCase)
	fmt.Println("********************")
	userHandler.FindById(ctx)
	fmt.Println("********************")

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
