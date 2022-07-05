package server

import (
	"context"
	"grpcUser/proto/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Users *pb.UserRequests
}

func NewServer(users *pb.UserRequests) *Server {
	return &Server{Users: users}
}

var users []*pb.UserRequest

func (s *Server) AddUser(ctx context.Context, req *pb.UserRequest) (*pb.Response, error) {
	users = append(users, req)
	s.Users = &pb.UserRequests{
		Users: users,
	}
	return &pb.Response{Code: 200, Message: "user saved successfully"}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserRequests, error) {
	for _, user := range s.Users.Users {
		if user.Verified == req.Verified {
			users = append(users, user)
		}
	}
	return &pb.UserRequests{Users: users}, nil
}
