package service

import (
	"context"

	"github.com/joaoianuci/go-grpc/internal/database"
	"github.com/joaoianuci/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	newCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return newCategory, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.CategoryListRequest) (*pb.CategoryListResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	pbCategories := make([]*pb.Category, 0)
	for _, category := range categories {
		pbCategory := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &pb.CategoryListResponse{Categories: pbCategories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	searchedCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return searchedCategory, nil
}
