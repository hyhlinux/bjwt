package main

import (
	"bjwt/config"
	"bjwt/controllers"
	_ "bjwt/routers"

	pb "bjwt/protos"
	"net"
	//"github.com/astaxie/beego"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//beego.Run()
	grpcStart()
}

func grpcStart() {
	lis, err := net.Listen("tcp", config.AppConf.GrpcListen)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	fmt.Printf("Jwt servic grpc-port(%v)\n", config.AppConf.GrpcListen)
	//grpc token
	pb.RegisterTokenServiceServer(s, &controllers.JWTSerController{})
	//grpc email
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
