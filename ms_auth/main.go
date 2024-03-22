package main

import (
	"context"
	"fmt"
	"log"
	"ms_auth/infrastructure"
	"ms_auth/model"
	"ms_auth/pb"
	"ms_auth/service"
	"net"

	"google.golang.org/grpc"
)

// Implements the Server gRPC
type Server_GRPC_MS_Auth struct {
	userService service.UserService
	pb.UnimplementedAuthenServer
}

func (s *Server_GRPC_MS_Auth) Login(ctx context.Context, in *pb.LoginMessage) (*pb.LoginResponse, error) {
	// log.Printf("=> Request Login From: %v\n", in.GetEmail())
	result, err := s.userService.Login(in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		User: &pb.User{
			Name:  result.Name,
			Email: result.Email,
		},
		SessionId:   fmt.Sprintf("%d", result.ID),
		AccessToken: fmt.Sprintf(">%d - Access: %s connect to system", result.ID, result.Email),
	}, nil
}

func (s *Server_GRPC_MS_Auth) CreateUser(ctx context.Context, in *pb.CreateUserMessage) (*pb.CreateUserResponse, error) {
	log.Printf("=> Request register: %v\n", in.GetName())
	result, err := s.userService.CreateUser(in.GetName(), in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("=> Error: %v", err)
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Name:     result.Name,
			Email:    result.Email,
			Password: result.Password,
		},
	}, nil
}

func main() {

	defer func() {
		caches := infrastructure.GetCache()
		db := infrastructure.GetDB()

		// Clear database cache
		if err := db.Delete(&model.CacheMem{}, "id IS NOT NULL").Error; err != nil {
			log.Println("==> Err clear db cache")
		}

		temp := []model.CacheMem{}
		for cache := range caches {
			temp = append(temp, model.CacheMem{
				Email: cache,
			})
		}

		if err := db.Model(&model.CacheMem{}).Create(&temp).Error; err != nil {
			log.Println("=====> Missing in-memory")
		}
	}()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthenServer(s, &Server_GRPC_MS_Auth{
		userService: service.NewUserService(),
	})
	log.Printf("Starting micro_auth: port - %s\n", "9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
