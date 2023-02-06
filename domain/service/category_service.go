package service

import (
	"github.com/golineshop/category/domain/model"
	"github.com/golineshop/category/domain/repository"
)

type ICategoryService interface {
	AddCategory(category *model.Category) (id int64, err error)
	DeleteCategoryByID(ID int64) (err error)
	UpdateCategory(category *model.Category) (err error)
	FindCategoryByID(ID int64) (category *model.Category, err error)
	FindAllCategory() (categoryList []model.Category, err error)
	FindCategoryByName(name string) (category *model.Category, err error)
	FindCategoryByLevel(level uint32) (categoryList []model.Category, err error)
	FindCategoryByParent(parentID int64) (categoryList []model.Category, err error)
}

func NewCategoryService(categoryRepository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func (c *CategoryService) AddCategory(category *model.Category) (id int64, err error) {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryService) DeleteCategoryByID(ID int64) (err error) {
	return c.CategoryRepository.DeleteCategoryByID(ID)
}

func (c *CategoryService) UpdateCategory(category *model.Category) (err error) {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryService) FindCategoryByID(ID int64) (category *model.Category, err error) {
	return c.CategoryRepository.FindCategoryByID(ID)
}

func (c *CategoryService) FindAllCategory() (categoryList []model.Category, err error) {
	return c.CategoryRepository.FindAll()
}

func (c *CategoryService) FindCategoryByName(name string) (category *model.Category, err error) {
	return c.CategoryRepository.FindCategoryByName(name)
}

func (c *CategoryService) FindCategoryByLevel(level uint32) (categoryList []model.Category, err error) {
	return c.CategoryRepository.FindCategoryByLevel(level)
}

func (c *CategoryService) FindCategoryByParent(parentID int64) (categoryList []model.Category, err error) {
	return c.CategoryRepository.FindCategoryByParent(parentID)
}
