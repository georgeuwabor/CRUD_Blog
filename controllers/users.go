package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController interface {
	RegisterNewUser() gin.HandlerFunc
	LoginUser() gin.HandlerFunc
}

type userController struct{
	col*mongo.Collection
	ctx context.Context
}


func InitUserController(col*mongo.Collection, ctx context.Context) UserController{
	return &userController{
		col: col,
		ctx: ctx,
	}
}

type UserReq struct{
	ID string `uri:"id"`
}

func(u*userController) RegisterNewUser()gin.HandlerFunc{
	return func(c *gin.Context){}
}
func(u*userController) LoginUser()gin.HandlerFunc{
	return func(c *gin.Context){
		var req UserReq
		if err:= c.ShouldBind(&req);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err})
			return
		}
		
		fmt.Println(req.ID)
	return	
	}
}
