package main

import (
	"log"

	protoData "grpc-client/proto"
	userHandler "grpc-client/user/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	port := "8827"
	targetPort := "8957"
	conn, err := grpc.Dial("localhost:"+targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cloud not conect to %v %v", targetPort, err)
	}

	user := protoData.NewUsersClient(conn)

	router := gin.Default()

	userHandler.NewUserController(router, user)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
