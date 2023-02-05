package repository

import (
	"github.com/golineshop/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	// InitTable 初始化数据表
	InitTable() (err error)
	// FindCategoryByID 根据ID查找分类
	FindCategoryByID(id int64) (category *model.Category, err error)
	// CreateCategory 创建分类
	CreateCategory(category *model.Category) (id int64, err error)
	// DeleteCategoryByID 根据ID删除分类
	DeleteCategoryByID(id int64) (err error)
	// UpdateCategory 根据ID更新分类
	UpdateCategory(category *model.Category) (err error)
	// FindAll 查找所有分类
	FindAll() (categoryList []model.Category, err error)
	// FindCategoryByName 根据名字查找分类
	FindCategoryByName(name string) (category *model.Category, err error)
	// FindCategoryByLevel 根据级别查找分类
	FindCategoryByLevel(level uint32) (categoryList []model.Category, err error)
	// FindCategoryByParent 根据父分类查找分类
	FindCategoryByParent(parentID int64) (categoryList []model.Category, err error)
}

// NewCategoryRepository 创建categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化数据表
func (c *CategoryRepository) InitTable() (err error) {
	return c.mysqlDb.CreateTable(&model.Category{}).Error
}

// FindCategoryByID 根据ID查找分类
func (c *CategoryRepository) FindCategoryByID(id int64) (category *model.Category, err error) {
	return category, c.mysqlDb.First(category, id).Error
}

// CreateCategory 创建分类
func (c *CategoryRepository) CreateCategory(category *model.Category) (id int64, err error) {
	return category.ID, c.mysqlDb.Create(category).Error
}

// DeleteCategoryByID 根据ID删除分类
func (c *CategoryRepository) DeleteCategoryByID(id int64) (err error) {
	return c.mysqlDb.Where("id = ?", id).Delete(&model.Category{}).Error
}

// UpdateCategory 根据ID更新分类
func (c *CategoryRepository) UpdateCategory(category *model.Category) (err error) {
	return c.mysqlDb.Model(category).Update(category).Error
}

// FindAll 查找所有分类
func (c *CategoryRepository) FindAll() (categoryList []model.Category, err error) {
	return categoryList, c.mysqlDb.Find(&categoryList).Error
}

// FindCategoryByName 根据名字查找分类
func (c *CategoryRepository) FindCategoryByName(name string) (category *model.Category, err error) {
	return category, c.mysqlDb.Where("category_name = ?", name).Find(category).Error
}

// FindCategoryByLevel 根据级别查找分类
func (c *CategoryRepository) FindCategoryByLevel(level uint32) (categoryList []model.Category, err error) {
	return categoryList, c.mysqlDb.Where("category_level = ?", level).Find(categoryList).Error
}

// FindCategoryByParent 根据父分类查找分类
func (c *CategoryRepository) FindCategoryByParent(parentID int64) (categoryList []model.Category, err error) {
	return categoryList, c.mysqlDb.Where("category_parent = ?", parentID).Find(categoryList).Error
}
