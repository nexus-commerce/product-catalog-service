package server

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/nexus-commerce/nexus-contracts-go/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"product-catalog-service/internal/service"
)

type Server struct {
	pb.UnimplementedProductCatalogServiceServer
	service *service.Service
}

func NewProductCatalogServer(s *service.Service) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	p, err := s.service.GetProduct(ctx, r.GetId())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("pb not found: %v", err))
		}
		return nil, err
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:            p.ID.Hex(),
			Sku:           p.Sku,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			StockQuantity: p.StockQuantity,
			Category:      p.Category,
			ImageUrl:      p.ImageURL,
			IsActive:      p.IsActive,
			Attributes:    p.Attributes,
		},
	}, err
}

func (s *Server) ListProducts(ctx context.Context, r *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var query string
	if r.GetFilter() != "" {
		query = r.GetFilter()
	}

	products, nextPage, err := s.service.ListProducts(ctx, query, r.GetPage(), r.GetPageSize())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var productList []*pb.Product
	for _, p := range products {
		productList = append(productList, &pb.Product{
			Id:            p.ID.Hex(),
			Sku:           p.Sku,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			StockQuantity: p.StockQuantity,
			Category:      p.Category,
			ImageUrl:      p.ImageURL,
			IsActive:      p.IsActive,
			Attributes:    p.Attributes,
		})
	}

	return &pb.ListProductsResponse{
		Products: productList,
		NextPage: nextPage,
	}, nil
}

func (s *Server) CreateProduct(ctx context.Context, r *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	p, err := s.service.CreateProduct(ctx, r.GetProduct())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:            p.ID.Hex(),
			Sku:           p.Sku,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			StockQuantity: p.StockQuantity,
			Category:      p.Category,
			ImageUrl:      p.ImageURL,
			IsActive:      p.IsActive,
			Attributes:    p.Attributes,
		},
	}, err
}

func (s *Server) UpdateProduct(ctx context.Context, r *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	p, err := s.service.UpdateProduct(ctx, r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateProductResponse{Product: &pb.Product{
		Id:            p.ID.Hex(),
		Sku:           p.Sku,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      p.Category,
		ImageUrl:      p.ImageURL,
		IsActive:      p.IsActive,
		Attributes:    p.Attributes,
	}}, err
}

func (s *Server) DeleteProduct(ctx context.Context, r *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := s.service.DeleteProduct(ctx, r.GetId())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("pb not found: %v", err))
		}
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteProductResponse{}, nil
}

func (s *Server) GetProductBySKU(ctx context.Context, r *pb.GetProductBySKURequest) (*pb.GetProductBySKUResponse, error) {
	p, err := s.service.GetProductBySKU(ctx, r.GetSku())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("p not found: %v", err))
		}
		return nil, err
	}

	return &pb.GetProductBySKUResponse{
		Product: &pb.Product{
			Id:            p.ID.Hex(),
			Sku:           p.Sku,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			StockQuantity: p.StockQuantity,
			Category:      p.Category,
			ImageUrl:      p.ImageURL,
			IsActive:      p.IsActive,
			Attributes:    p.Attributes,
		},
	}, nil
}
