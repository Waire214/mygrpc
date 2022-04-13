package servers

import (
	"context"
	"product/server/service/pb"
)

type Server struct {
	pb.UnimplementedProductInfoServer
}

func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	return &pb.ProductID{
		Value: "hello",
	}, nil
}
