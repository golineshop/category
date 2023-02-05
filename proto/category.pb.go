// Code generated by protoc-gen-go. DO NOT EDIT.
// source: category.proto

/*
Package go_micro_service_category is a generated protocol buffer package.

It is generated from these files:

	category.proto

It has these top-level messages:

	CreateCategoryRequest
	CreateCategoryResponse
	UpdateCategoryRequest
	UpdateCategoryResponse
	DeleteCategoryRequest
	DeleteCategoryResponse
	FindByNameRequest
	FindByNameResponse
	FindByIDRequest
	FindByIDResponse
	FindByLevelRequest
	FindByLevelResponse
	FindByParentRequest
	FindByParentResponse
	FindAllRequest
	FindAllResponse
	CategoryResponse
*/
package go_micro_service_category

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateCategoryRequest struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CreateCategoryRequest) Reset()                    { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()               {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateCategoryRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CreateCategoryRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CreateCategoryRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CreateCategoryRequest) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *CreateCategoryRequest) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type CreateCategoryResponse struct {
	Message    string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	CategoryId int64  `protobuf:"varint,2,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *CreateCategoryResponse) Reset()                    { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()               {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateCategoryResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type UpdateCategoryRequest struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *UpdateCategoryRequest) Reset()                    { *m = UpdateCategoryRequest{} }
func (m *UpdateCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoryRequest) ProtoMessage()               {}
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateCategoryRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *UpdateCategoryRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *UpdateCategoryRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *UpdateCategoryRequest) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *UpdateCategoryRequest) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type UpdateCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateCategoryResponse) Reset()                    { *m = UpdateCategoryResponse{} }
func (m *UpdateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoryResponse) ProtoMessage()               {}
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteCategoryRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *DeleteCategoryRequest) Reset()                    { *m = DeleteCategoryRequest{} }
func (m *DeleteCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryRequest) ProtoMessage()               {}
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteCategoryRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type DeleteCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteCategoryResponse) Reset()                    { *m = DeleteCategoryResponse{} }
func (m *DeleteCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryResponse) ProtoMessage()               {}
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FindByNameRequest struct {
	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
}

func (m *FindByNameRequest) Reset()                    { *m = FindByNameRequest{} }
func (m *FindByNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByNameRequest) ProtoMessage()               {}
func (*FindByNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FindByNameRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

type FindByNameResponse struct {
	CategoryId          int64  `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,5,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *FindByNameResponse) Reset()                    { *m = FindByNameResponse{} }
func (m *FindByNameResponse) String() string            { return proto.CompactTextString(m) }
func (*FindByNameResponse) ProtoMessage()               {}
func (*FindByNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindByNameResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *FindByNameResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *FindByNameResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *FindByNameResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *FindByNameResponse) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *FindByNameResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type FindByIDRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *FindByIDRequest) Reset()                    { *m = FindByIDRequest{} }
func (m *FindByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByIDRequest) ProtoMessage()               {}
func (*FindByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FindByIDRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type FindByIDResponse struct {
	CategoryId          int64  `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,5,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *FindByIDResponse) Reset()                    { *m = FindByIDResponse{} }
func (m *FindByIDResponse) String() string            { return proto.CompactTextString(m) }
func (*FindByIDResponse) ProtoMessage()               {}
func (*FindByIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FindByIDResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *FindByIDResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *FindByIDResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *FindByIDResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *FindByIDResponse) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *FindByIDResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type FindByLevelRequest struct {
	CategoryLevel uint32 `protobuf:"varint,1,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
}

func (m *FindByLevelRequest) Reset()                    { *m = FindByLevelRequest{} }
func (m *FindByLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByLevelRequest) ProtoMessage()               {}
func (*FindByLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FindByLevelRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

type FindByLevelResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindByLevelResponse) Reset()                    { *m = FindByLevelResponse{} }
func (m *FindByLevelResponse) String() string            { return proto.CompactTextString(m) }
func (*FindByLevelResponse) ProtoMessage()               {}
func (*FindByLevelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FindByLevelResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

type FindByParentRequest struct {
	CategoryParent int64 `protobuf:"varint,1,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
}

func (m *FindByParentRequest) Reset()                    { *m = FindByParentRequest{} }
func (m *FindByParentRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByParentRequest) ProtoMessage()               {}
func (*FindByParentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FindByParentRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

type FindByParentResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindByParentResponse) Reset()                    { *m = FindByParentResponse{} }
func (m *FindByParentResponse) String() string            { return proto.CompactTextString(m) }
func (*FindByParentResponse) ProtoMessage()               {}
func (*FindByParentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *FindByParentResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

type FindAllRequest struct {
}

func (m *FindAllRequest) Reset()                    { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()               {}
func (*FindAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type FindAllResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindAllResponse) Reset()                    { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()               {}
func (*FindAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *FindAllResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

type CategoryResponse struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CategoryResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryResponse) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *CategoryResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCategoryRequest)(nil), "go.micro.service.category.CreateCategoryRequest")
	proto.RegisterType((*CreateCategoryResponse)(nil), "go.micro.service.category.CreateCategoryResponse")
	proto.RegisterType((*UpdateCategoryRequest)(nil), "go.micro.service.category.UpdateCategoryRequest")
	proto.RegisterType((*UpdateCategoryResponse)(nil), "go.micro.service.category.UpdateCategoryResponse")
	proto.RegisterType((*DeleteCategoryRequest)(nil), "go.micro.service.category.DeleteCategoryRequest")
	proto.RegisterType((*DeleteCategoryResponse)(nil), "go.micro.service.category.DeleteCategoryResponse")
	proto.RegisterType((*FindByNameRequest)(nil), "go.micro.service.category.FindByNameRequest")
	proto.RegisterType((*FindByNameResponse)(nil), "go.micro.service.category.FindByNameResponse")
	proto.RegisterType((*FindByIDRequest)(nil), "go.micro.service.category.FindByIDRequest")
	proto.RegisterType((*FindByIDResponse)(nil), "go.micro.service.category.FindByIDResponse")
	proto.RegisterType((*FindByLevelRequest)(nil), "go.micro.service.category.FindByLevelRequest")
	proto.RegisterType((*FindByLevelResponse)(nil), "go.micro.service.category.FindByLevelResponse")
	proto.RegisterType((*FindByParentRequest)(nil), "go.micro.service.category.FindByParentRequest")
	proto.RegisterType((*FindByParentResponse)(nil), "go.micro.service.category.FindByParentResponse")
	proto.RegisterType((*FindAllRequest)(nil), "go.micro.service.category.FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "go.micro.service.category.FindAllResponse")
	proto.RegisterType((*CategoryResponse)(nil), "go.micro.service.category.CategoryResponse")
}

func init() { proto.RegisterFile("category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x26, 0xfd, 0x62, 0x4a, 0xd3, 0xb0, 0x6d, 0x23, 0x63, 0x21, 0x88, 0x16, 0x21, 0x42,
	0x0b, 0x2e, 0x0d, 0x17, 0xa4, 0x4a, 0x48, 0xb4, 0x11, 0x28, 0x12, 0x42, 0xc8, 0x88, 0x4b, 0x0f,
	0x58, 0x26, 0x59, 0x45, 0x46, 0xf1, 0x07, 0xb6, 0x09, 0xea, 0x15, 0xf1, 0xa3, 0xf8, 0x3b, 0x5c,
	0xf8, 0x1d, 0xc8, 0xbb, 0xde, 0x75, 0xfd, 0x81, 0xbd, 0x91, 0x7a, 0x6a, 0x8f, 0x99, 0x9d, 0x37,
	0x6f, 0xf6, 0xcd, 0xe6, 0x8d, 0xa1, 0x33, 0xb1, 0x63, 0x3a, 0xf3, 0xc3, 0x0b, 0x23, 0x08, 0xfd,
	0xd8, 0xc7, 0x77, 0x67, 0xbe, 0xe1, 0x3a, 0x93, 0xd0, 0x37, 0x22, 0x1a, 0x2e, 0x9c, 0x09, 0x35,
	0x44, 0x02, 0xf9, 0x8b, 0x60, 0xff, 0x2c, 0xa4, 0x76, 0x4c, 0xcf, 0xd2, 0x90, 0x49, 0xbf, 0x7d,
	0xa7, 0x51, 0x8c, 0x1f, 0xc2, 0xb6, 0xc8, 0xb2, 0x3c, 0xdb, 0xa5, 0x1a, 0xea, 0xa3, 0xc1, 0x2d,
	0xf3, 0xb6, 0x08, 0xbe, 0xb7, 0x5d, 0x8a, 0x1f, 0x65, 0x5c, 0xd6, 0x9c, 0x2e, 0xe8, 0x5c, 0x6b,
	0xf5, 0xd1, 0x60, 0xdb, 0x94, 0xd0, 0x77, 0x49, 0x10, 0x3f, 0x86, 0x1d, 0x99, 0x16, 0xd8, 0x21,
	0xf5, 0x62, 0xad, 0xdd, 0x47, 0x83, 0xb6, 0x29, 0xd1, 0x1f, 0x58, 0x34, 0x57, 0xcf, 0x71, 0xed,
	0x19, 0xd5, 0x56, 0x19, 0xab, 0xac, 0x37, 0x4e, 0x82, 0xf8, 0x18, 0xf6, 0x64, 0xda, 0x94, 0x46,
	0x93, 0xd0, 0x09, 0x62, 0xc7, 0xf7, 0xb4, 0x35, 0x96, 0xbc, 0x2b, 0xce, 0x46, 0xd9, 0x11, 0xf9,
	0x08, 0xbd, 0xe2, 0x3d, 0xa3, 0xc0, 0xf7, 0x22, 0x8a, 0x35, 0xd8, 0x70, 0x69, 0x14, 0x25, 0x64,
	0xfc, 0x8a, 0xe2, 0x27, 0x7e, 0x00, 0x5b, 0x59, 0x37, 0x53, 0x76, 0xb5, 0xb6, 0x09, 0xb2, 0x95,
	0x29, 0x53, 0xef, 0x53, 0x30, 0xbd, 0xfe, 0xea, 0x0d, 0xa1, 0x57, 0xbc, 0x67, 0x93, 0x7a, 0xe4,
	0x25, 0xec, 0x8f, 0xe8, 0x9c, 0x96, 0xb5, 0x29, 0xc8, 0x8a, 0x4a, 0xb2, 0x0e, 0xa1, 0x57, 0x44,
	0x2a, 0xb0, 0xdd, 0x79, 0xe3, 0x78, 0xd3, 0x53, 0xa6, 0xec, 0x32, 0x53, 0x20, 0xbf, 0x5a, 0x80,
	0x2f, 0x43, 0x53, 0xaa, 0xa6, 0x2e, 0xcb, 0xc5, 0x5b, 0x4a, 0x23, 0x6e, 0x2b, 0x8e, 0x78, 0x55,
	0x71, 0xc4, 0x6b, 0xcb, 0x8c, 0x78, 0xbd, 0x6e, 0xc4, 0x3b, 0x5c, 0x85, 0xf1, 0x48, 0x79, 0x50,
	0x3f, 0x5b, 0xd0, 0xcd, 0x40, 0x37, 0x54, 0xb8, 0x13, 0xf1, 0x7c, 0x58, 0x47, 0x42, 0xbb, 0x72,
	0xff, 0xa8, 0xa2, 0x7f, 0xf2, 0x19, 0x76, 0x73, 0xe0, 0x54, 0xc3, 0xb7, 0xb0, 0x29, 0xf2, 0x34,
	0xd4, 0x6f, 0x0f, 0xb6, 0x86, 0x87, 0xc6, 0x7f, 0x4d, 0xdc, 0x28, 0xfe, 0x4d, 0x4c, 0x09, 0x26,
	0xaf, 0x44, 0x7d, 0x2e, 0x83, 0xe8, 0xae, 0x42, 0x36, 0x54, 0x25, 0x1b, 0xb1, 0x60, 0x2f, 0x8f,
	0xbf, 0xea, 0x06, 0xbb, 0xd0, 0x49, 0x08, 0x5e, 0xcf, 0x85, 0x72, 0xe4, 0x9c, 0x3f, 0x44, 0x16,
	0xb9, 0x6a, 0xb6, 0x3f, 0x08, 0xba, 0x25, 0x53, 0xb9, 0x5e, 0x5e, 0x3d, 0xfc, 0xbd, 0x01, 0x9b,
	0xe2, 0x8e, 0xf8, 0x07, 0x74, 0xf2, 0x6b, 0x0f, 0x3f, 0xaf, 0x53, 0xae, 0xea, 0x4b, 0x40, 0x3f,
	0x5e, 0x02, 0xc1, 0x25, 0x25, 0x2b, 0x09, 0x71, 0x7e, 0x63, 0xd4, 0x12, 0x57, 0x2e, 0xd1, 0x5a,
	0xe2, 0xea, 0x75, 0xc4, 0x89, 0xf3, 0xcb, 0xa3, 0x96, 0xb8, 0x72, 0x43, 0xd5, 0x12, 0x57, 0x6f,
	0x26, 0xb2, 0x82, 0x23, 0xee, 0x03, 0xe2, 0x84, 0xaf, 0x13, 0xfc, 0xb4, 0xa6, 0x54, 0x69, 0x61,
	0xe9, 0xcf, 0x14, 0xb3, 0x25, 0xa9, 0xcb, 0x0d, 0x38, 0x23, 0x1d, 0x8f, 0xf0, 0x41, 0x63, 0x11,
	0x69, 0xf1, 0xfa, 0xa1, 0x52, 0xae, 0xa4, 0x5b, 0x70, 0x3b, 0xc9, 0xe8, 0xf8, 0xab, 0x6f, 0x6e,
	0xfb, 0xb2, 0x37, 0xea, 0x86, 0x6a, 0xba, 0xe4, 0xbd, 0xe0, 0x36, 0x94, 0xf1, 0xa6, 0xff, 0xa2,
	0xe6, 0x4a, 0x39, 0xdf, 0xd3, 0x8f, 0x94, 0xf3, 0x25, 0xf5, 0x57, 0x69, 0x47, 0xf2, 0x41, 0x3d,
	0x69, 0xa8, 0x92, 0x99, 0x99, 0x7e, 0xa0, 0x92, 0x2a, 0xb8, 0x4e, 0xef, 0x9f, 0xdf, 0x33, 0x8e,
	0x4e, 0x66, 0xbe, 0xc5, 0x10, 0x56, 0x8a, 0xb0, 0x04, 0xe2, 0xcb, 0x3a, 0xfb, 0x9e, 0x7f, 0xf1,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xea, 0xb8, 0x85, 0xe1, 0x0b, 0x00, 0x00,
}