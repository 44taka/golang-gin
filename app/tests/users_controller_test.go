package tests

import (
	"fmt"
	"testing"

	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/interfaces/controllers"
	mocks "github.com/44taka/golang-gin/mocks/interfaces/controllers"
)

func TestGetAll(t *testing.T) {
	c := infrastructure.NewConfig()
	db := infrastructure.NewDB(c)
	usersController := controllers.NewUsersController(db)
	ctx := &mocks.Context{}
	usersController.GetAll(ctx)
	// fmt.Println(r)
	fmt.Println("aaaa")
}
