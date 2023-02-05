// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: category.proto

package go_micro_service_category

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Category service

func NewCategoryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Category service

type CategoryService interface {
	//创建分类
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error)
	//更新分类
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error)
	//删除分类
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error)
	//根据名称查找
	FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*FindByNameResponse, error)
	//根据ID查找
	FindCategoryByID(ctx context.Context, in *FindByIDRequest, opts ...client.CallOption) (*FindByIDResponse, error)
	//根据分类层级查找
	FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*FindByLevelResponse, error)
	//根据父级分类查找
	FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*FindByParentResponse, error)
	//查找所有分类
	FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*FindAllResponse, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.CreateCategory", in)
	out := new(CreateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.UpdateCategory", in)
	out := new(UpdateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Category.DeleteCategory", in)
	out := new(DeleteCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*FindByNameResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByName", in)
	out := new(FindByNameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByID(ctx context.Context, in *FindByIDRequest, opts ...client.CallOption) (*FindByIDResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByID", in)
	out := new(FindByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*FindByLevelResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByLevel", in)
	out := new(FindByLevelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*FindByParentResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindCategoryByParent", in)
	out := new(FindByParentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*FindAllResponse, error) {
	req := c.c.NewRequest(c.name, "Category.FindAllCategory", in)
	out := new(FindAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Category service

type CategoryHandler interface {
	//创建分类
	CreateCategory(context.Context, *CreateCategoryRequest, *CreateCategoryResponse) error
	//更新分类
	UpdateCategory(context.Context, *UpdateCategoryRequest, *UpdateCategoryResponse) error
	//删除分类
	DeleteCategory(context.Context, *DeleteCategoryRequest, *DeleteCategoryResponse) error
	//根据名称查找
	FindCategoryByName(context.Context, *FindByNameRequest, *FindByNameResponse) error
	//根据ID查找
	FindCategoryByID(context.Context, *FindByIDRequest, *FindByIDResponse) error
	//根据分类层级查找
	FindCategoryByLevel(context.Context, *FindByLevelRequest, *FindByLevelResponse) error
	//根据父级分类查找
	FindCategoryByParent(context.Context, *FindByParentRequest, *FindByParentResponse) error
	//查找所有分类
	FindAllCategory(context.Context, *FindAllRequest, *FindAllResponse) error
}

func RegisterCategoryHandler(s server.Server, hdlr CategoryHandler, opts ...server.HandlerOption) error {
	type category interface {
		CreateCategory(ctx context.Context, in *CreateCategoryRequest, out *CreateCategoryResponse) error
		UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, out *UpdateCategoryResponse) error
		DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error
		FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *FindByNameResponse) error
		FindCategoryByID(ctx context.Context, in *FindByIDRequest, out *FindByIDResponse) error
		FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *FindByLevelResponse) error
		FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *FindByParentResponse) error
		FindAllCategory(ctx context.Context, in *FindAllRequest, out *FindAllResponse) error
	}
	type Category struct {
		category
	}
	h := &categoryHandler{hdlr}
	return s.Handle(s.NewHandler(&Category{h}, opts...))
}

type categoryHandler struct {
	CategoryHandler
}

func (h *categoryHandler) CreateCategory(ctx context.Context, in *CreateCategoryRequest, out *CreateCategoryResponse) error {
	return h.CategoryHandler.CreateCategory(ctx, in, out)
}

func (h *categoryHandler) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, out *UpdateCategoryResponse) error {
	return h.CategoryHandler.UpdateCategory(ctx, in, out)
}

func (h *categoryHandler) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error {
	return h.CategoryHandler.DeleteCategory(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *FindByNameResponse) error {
	return h.CategoryHandler.FindCategoryByName(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByID(ctx context.Context, in *FindByIDRequest, out *FindByIDResponse) error {
	return h.CategoryHandler.FindCategoryByID(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *FindByLevelResponse) error {
	return h.CategoryHandler.FindCategoryByLevel(ctx, in, out)
}

func (h *categoryHandler) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *FindByParentResponse) error {
	return h.CategoryHandler.FindCategoryByParent(ctx, in, out)
}

func (h *categoryHandler) FindAllCategory(ctx context.Context, in *FindAllRequest, out *FindAllResponse) error {
	return h.CategoryHandler.FindAllCategory(ctx, in, out)
}
