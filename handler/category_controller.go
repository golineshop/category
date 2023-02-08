package handler

import (
	"context"
	"github.com/golineshop/category/common"
	"github.com/golineshop/category/domain/model"
	"github.com/golineshop/category/domain/service"
	proto "github.com/golineshop/category/proto"
	"github.com/prometheus/common/log"
)

type CategoryController struct {
	CategoryService service.ICategoryService
}

// CreateCategory 创建分类
func (c *CategoryController) CreateCategory(ctx context.Context, request *proto.CreateCategoryRequest, response *proto.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(category, request)
	if err != nil {
		return err
	}
	ID, err := c.CategoryService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "添加信息成功"
	response.CategoryId = ID
	return nil
}

// UpdateCategory 更新分类信息
func (c *CategoryController) UpdateCategory(ctx context.Context, request *proto.UpdateCategoryRequest, response *proto.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CategoryService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "类别更新成功"
	return nil
}

// DeleteCategory 分类删除
func (c *CategoryController) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest, response *proto.DeleteCategoryResponse) error {
	err := c.CategoryService.DeleteCategoryByID(request.CategoryId)
	if err != nil {
		return nil
	}
	response.Message = "删除成功"
	return nil
}

// FindCategoryByName 根据分类名称查找分类
func (c *CategoryController) FindCategoryByName(ctx context.Context, request *proto.FindByNameRequest, response *proto.FindByNameResponse) error {
	category, err := c.CategoryService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

// FindCategoryByID 根据分类ID查找分类
func (c *CategoryController) FindCategoryByID(ctx context.Context, request *proto.FindByIDRequest, response *proto.FindByIDResponse) error {
	category, err := c.CategoryService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

// FindCategoryByLevel 根据层级查找分类
func (c *CategoryController) FindCategoryByLevel(ctx context.Context, request *proto.FindByLevelRequest, response *proto.FindByLevelResponse) error {
	categoryList, err := c.CategoryService.FindCategoryByLevel(request.CategoryLevel)
	if err != nil {
		return err
	}
	for _, category := range categoryList {
		singleCategoryResponse := &proto.CategoryResponse{}
		err := common.SwapTo(category, singleCategoryResponse)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, singleCategoryResponse)
	}
	return nil
}

// FindCategoryByParent 根据父分类查找分类
func (c *CategoryController) FindCategoryByParent(ctx context.Context, request *proto.FindByParentRequest, response *proto.FindByParentResponse) error {
	categoryList, err := c.CategoryService.FindCategoryByParent(request.CategoryParent)
	if err != nil {
		return err
	}
	for _, category := range categoryList {
		singleCategoryResponse := &proto.CategoryResponse{}
		err := common.SwapTo(category, singleCategoryResponse)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, singleCategoryResponse)
	}
	return nil
}

// FindAllCategory 查询所有分类
func (c *CategoryController) FindAllCategory(ctx context.Context, request *proto.FindAllRequest, response *proto.FindAllResponse) error {
	categoryList, err := c.CategoryService.FindAllCategory()
	if err != nil {
		return err
	}
	for _, category := range categoryList {
		singleCategoryResponse := &proto.CategoryResponse{}
		err := common.SwapTo(category, singleCategoryResponse)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, singleCategoryResponse)
	}
	return nil
}
