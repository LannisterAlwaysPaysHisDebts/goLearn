package services

import (
	"go_iris/datamodels"
	"go_iris/repositories"
)

type IProductService interface {
	GetById(int64) (*datamodels.Product, error)
	GetAll() ([]*datamodels.Product, error)
	DeleteById(int64) bool
	Insert(product *datamodels.Product) (int64, error)
	Update(product *datamodels.Product) error
	SubNumberOne(productId int64) error
}

type ProductService struct {
	Repository repositories.IProduct
}

func (p *ProductService) GetById(id int64) (*datamodels.Product, error) {
	return p.Repository.SelectByKey(id)
}

func (p *ProductService) GetAll() ([]*datamodels.Product, error) {
	return p.Repository.SelectAll()
}

func (p *ProductService) DeleteById(id int64) bool {
	return p.Repository.Delete(id)
}

func (p *ProductService) Insert(product *datamodels.Product) (int64, error) {
	return p.Repository.Insert(product)
}

func (p *ProductService) Update(product *datamodels.Product) error {
	return p.Repository.Update(product)
}

func (p *ProductService) SubNumberOne(productId int64) error {
	return p.Repository.SubProductNum(productId)
}

func NewProductService(product repositories.IProduct) IProductService {
	return &ProductService{product}
}
