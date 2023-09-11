package main

import (
	"context"
	"crud_blog/controllers"
	"crud_blog/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var(
	users_controller controllers.UserController
	posts_controller controllers.PostController
)

func Run()*gin.Engine{
	env, err := utils.LoadEnv(".")
	
	if err != nil {
		log.Fatal("cannot load env", err)
	}
	
	ctx := context.TODO()
	
	mongoConn := options.Client().ApplyURI(env.MongoURI)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	
	if err != nil{
		log.Panic((err.Error()))
	}
	
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil{
		panic(err)
	}
	
	fmt.Println("MongoDB connection successfull!")
	
	users_controller, posts_controller = initCols(mongoClient, ctx)
	
	server := gin.Default()
	
	// Initialize routes
	
	server.POST("/register", users_controller.RegisterNewUser())
	server.GET("/login/:id", users_controller.LoginUser())
	return server
}


func initCols(client *mongo.Client, ctx context.Context)(controllers.UserController, controllers.PostController){
	user_col := client.Database("crud_blog").Collection("users")
	posts_col := client.Database("crud_blog").Collection("posts")
	
	users_controller = controllers.InitUserController(user_col, ctx)
	posts_controller = controllers.InitPostController(posts_col, ctx)
	
	return users_controller, posts_controller
}

func main(){
	server := Run()
	log.Fatal(server.Run(":" + "5001"))
}