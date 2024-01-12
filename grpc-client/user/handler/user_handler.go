package handle

import (
	"context"
	protoData "grpc-client/proto"
	"grpc-client/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserController struct {
	userClient protoData.UsersClient
}

func NewUserController(r *gin.Engine, userClient protoData.UsersClient) {
	userHandler := &UserController{userClient}

	r.GET("/user", userHandler.GetUser)
	r.POST("/user", userHandler.InsertUser)
	r.GET("/user/:id", userHandler.FindUserById)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
}

func (e *UserController) GetUser(ctx *gin.Context) {
	userList, err := e.userClient.GetUserList(ctx, new(empty.Empty))
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(ctx, userList.List)
}

func (e *UserController) FindUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userid := protoData.UserId{
		Id: id,
	}
	user, err := e.userClient.GetUserById(ctx, &userid)
	if err != nil {
		utils.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccess(ctx, user)
}

func (e *UserController) InsertUser(ctx *gin.Context) {
	var user protoData.User
	err := ctx.Bind(&user)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = e.userClient.InsertUser(context.Background(), &user)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(ctx, "Success Insert Data")
}

func (e *UserController) UpdateUser(ctx *gin.Context) {
	var user protoData.User
	id := ctx.Param("id")
	err := ctx.Bind(&user)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var up = protoData.UserUpdate{
		Id:   id,
		User: &user,
	}

	_, err = e.userClient.UpdateUser(ctx, &up)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(ctx, "Update User Success")
}

func (e *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var up = protoData.UserId{
		Id: id,
	}

	_, err := e.userClient.DeleteUser(ctx, &up)
	if err != nil {
		utils.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccess(ctx, "Success delete data")
}
