package repos

import (
	"context"
	"log"

	"github.com/fnxr21/brand/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func Repository(db *gorm.DB) *repository {
	return &repository{db: db}
}

type BrandService interface {
	Create(ctx context.Context, brand entities.Brand) (entities.Brand, error)
}

func (r *repository) Create(ctx context.Context, brand entities.Brand) (entities.Brand, error) {
	log.Println(brand, "brand")
	log.Println(ctx, "context")

	// err := r.db.WithContext(ctx).Create(&brand).Error
	return entities.Brand{}, nil
}

// func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
// 	err := r.db.Create(&cart).Error

// 	return cart, err
// }
