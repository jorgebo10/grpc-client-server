package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "grpc/server/ecommercepb"
	"log"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*Product
	UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, product *Product) (*ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	product.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*Product)
	}
	s.productMap[product.Id] = product
	log.Printf("Product %v : %v - Added.", product.Id, product.Name)
	return &ProductID{Value: product.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *ProductID) (*Product, error) {
	product, exists := s.productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
