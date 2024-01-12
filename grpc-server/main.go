package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"grpc-server/user/handler"
	userRepo "grpc-server/user/repository"
	userUsecase "grpc-server/user/usecase"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	var port = "8957"
	server := grpc.NewServer()

	db, err := gorm.Open("mysql", "user:user@tcp(127.0.0.1:3306)/user_grpc?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := userRepo.CreateUserRepoImpl(db)
	userUseCase := userUsecase.NewUserUseCase(userRepository)

	handler.NewUserHandler(server, userUseCase)

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Starting at Port: ", port)
	log.Fatal(server.Serve(conn))
}
