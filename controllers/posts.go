package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostController interface {
	
}

type postController struct{
	col*mongo.Collection
	ctx context.Context
}


func InitPostController(col*mongo.Collection, ctx context.Context) PostController{
	return &postController{
		col: col,
		ctx: ctx,
	}
}

