// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/blogs/v1/blogs.proto

package v1

import (
	global "blog/api/global"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HiddenType int32

const (
	HiddenType_Hidden   HiddenType = 0 //不隐藏
	HiddenType_IsHidden HiddenType = 1 //隐藏
)

// Enum value maps for HiddenType.
var (
	HiddenType_name = map[int32]string{
		0: "Hidden",
		1: "IsHidden",
	}
	HiddenType_value = map[string]int32{
		"Hidden":   0,
		"IsHidden": 1,
	}
)

func (x HiddenType) Enum() *HiddenType {
	p := new(HiddenType)
	*p = x
	return p
}

func (x HiddenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HiddenType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_blogs_v1_blogs_proto_enumTypes[0].Descriptor()
}

func (HiddenType) Type() protoreflect.EnumType {
	return &file_api_blogs_v1_blogs_proto_enumTypes[0]
}

func (x HiddenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HiddenType.Descriptor instead.
func (HiddenType) EnumDescriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{0}
}

type CreateBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string     `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`                                     //标题
	Description string     `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`                         //描述
	IsHidden    HiddenType `protobuf:"varint,3,opt,name=IsHidden,proto3,enum=api.blogs.v1.HiddenType" json:"IsHidden,omitempty"` //是否隐藏
	Tags        []string   `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`                                       //标签
	Cover       string     `protobuf:"bytes,5,opt,name=Cover,proto3" json:"Cover,omitempty"`                                     //封面
	Content     string     `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`                                 //内容
}

func (x *CreateBlogsRequest) Reset() {
	*x = CreateBlogsRequest{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogsRequest) ProtoMessage() {}

func (x *CreateBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogsRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogsRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlogsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBlogsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBlogsRequest) GetIsHidden() HiddenType {
	if x != nil {
		return x.IsHidden
	}
	return HiddenType_Hidden
}

func (x *CreateBlogsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateBlogsRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreateBlogsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateBlogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBlogsReply) Reset() {
	*x = CreateBlogsReply{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogsReply) ProtoMessage() {}

func (x *CreateBlogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogsReply.ProtoReflect.Descriptor instead.
func (*CreateBlogsReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{1}
}

type UpdateBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                                          //根据ID修改
	Title       string     `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`                                     //标题
	Description string     `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`                         //描述
	IsHidden    HiddenType `protobuf:"varint,4,opt,name=IsHidden,proto3,enum=api.blogs.v1.HiddenType" json:"IsHidden,omitempty"` //是否隐藏
	Tags        []string   `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`                                       //标签
	Cover       string     `protobuf:"bytes,6,opt,name=Cover,proto3" json:"Cover,omitempty"`                                     //封面
	Content     string     `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`                                 //内容
}

func (x *UpdateBlogsRequest) Reset() {
	*x = UpdateBlogsRequest{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogsRequest) ProtoMessage() {}

func (x *UpdateBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogsRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogsRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBlogsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBlogsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBlogsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBlogsRequest) GetIsHidden() HiddenType {
	if x != nil {
		return x.IsHidden
	}
	return HiddenType_Hidden
}

func (x *UpdateBlogsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateBlogsRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *UpdateBlogsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateBlogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBlogsReply) Reset() {
	*x = UpdateBlogsReply{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogsReply) ProtoMessage() {}

func (x *UpdateBlogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogsReply.ProtoReflect.Descriptor instead.
func (*UpdateBlogsReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{3}
}

type DeleteBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteBlogsRequest) Reset() {
	*x = DeleteBlogsRequest{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogsRequest) ProtoMessage() {}

func (x *DeleteBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogsRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlogsRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBlogsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBlogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBlogsReply) Reset() {
	*x = DeleteBlogsReply{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogsReply) ProtoMessage() {}

func (x *DeleteBlogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogsReply.ProtoReflect.Descriptor instead.
func (*DeleteBlogsReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{5}
}

type GetBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetBlogsRequest) Reset() {
	*x = GetBlogsRequest{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsRequest) ProtoMessage() {}

func (x *GetBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsRequest.ProtoReflect.Descriptor instead.
func (*GetBlogsRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{6}
}

func (x *GetBlogsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BlogsHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                                          //ID
	Title       string     `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`                                     //标题
	Description string     `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`                         //描述
	IsHidden    HiddenType `protobuf:"varint,4,opt,name=IsHidden,proto3,enum=api.blogs.v1.HiddenType" json:"IsHidden,omitempty"` //是否隐藏
	Tags        []string   `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`                                       //标签
	Cover       string     `protobuf:"bytes,6,opt,name=Cover,proto3" json:"Cover,omitempty"`                                     //封面
	AccountId   int64      `protobuf:"varint,7,opt,name=AccountId,proto3" json:"AccountId,omitempty"`                            //所属帐号ID
	CreatedAt   int64      `protobuf:"varint,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`                            //创建时间
}

func (x *BlogsHeader) Reset() {
	*x = BlogsHeader{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlogsHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogsHeader) ProtoMessage() {}

func (x *BlogsHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogsHeader.ProtoReflect.Descriptor instead.
func (*BlogsHeader) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{7}
}

func (x *BlogsHeader) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogsHeader) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogsHeader) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BlogsHeader) GetIsHidden() HiddenType {
	if x != nil {
		return x.IsHidden
	}
	return HiddenType_Hidden
}

func (x *BlogsHeader) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BlogsHeader) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *BlogsHeader) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *BlogsHeader) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetBlogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *BlogsHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`   //头部内容
	Content string       `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"` //主内容
}

func (x *GetBlogsReply) Reset() {
	*x = GetBlogsReply{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsReply) ProtoMessage() {}

func (x *GetBlogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsReply.ProtoReflect.Descriptor instead.
func (*GetBlogsReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlogsReply) GetHeader() *BlogsHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetBlogsReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ListBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *global.PageInfo `protobuf:"bytes,1,opt,name=Page,proto3" json:"Page,omitempty"`   //分页
	Title string           `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"` //标题
	Tags  []string         `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`   //标签
}

func (x *ListBlogsRequest) Reset() {
	*x = ListBlogsRequest{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsRequest) ProtoMessage() {}

func (x *ListBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsRequest.ProtoReflect.Descriptor instead.
func (*ListBlogsRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{9}
}

func (x *ListBlogsRequest) GetPage() *global.PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListBlogsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListBlogsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ListBlogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"` //总博客数
	List  []*BlogsHeader `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`    //返回的内容
}

func (x *ListBlogsReply) Reset() {
	*x = ListBlogsReply{}
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsReply) ProtoMessage() {}

func (x *ListBlogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsReply.ProtoReflect.Descriptor instead.
func (*ListBlogsReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_proto_rawDescGZIP(), []int{10}
}

func (x *ListBlogsReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListBlogsReply) GetList() []*BlogsHeader {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_blogs_v1_blogs_proto protoreflect.FileDescriptor

var file_api_blogs_v1_blogs_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x08, 0x49, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x49, 0x73, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xd6, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x08, 0x49,
	0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x49, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x08, 0x49, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x49, 0x73, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x55, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x26, 0x0a, 0x0a, 0x48, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x01, 0x32,
	0x8d, 0x03, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x26, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x14, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blogs_v1_blogs_proto_rawDescOnce sync.Once
	file_api_blogs_v1_blogs_proto_rawDescData = file_api_blogs_v1_blogs_proto_rawDesc
)

func file_api_blogs_v1_blogs_proto_rawDescGZIP() []byte {
	file_api_blogs_v1_blogs_proto_rawDescOnce.Do(func() {
		file_api_blogs_v1_blogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blogs_v1_blogs_proto_rawDescData)
	})
	return file_api_blogs_v1_blogs_proto_rawDescData
}

var file_api_blogs_v1_blogs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_blogs_v1_blogs_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_blogs_v1_blogs_proto_goTypes = []any{
	(HiddenType)(0),            // 0: api.blogs.v1.HiddenType
	(*CreateBlogsRequest)(nil), // 1: api.blogs.v1.CreateBlogsRequest
	(*CreateBlogsReply)(nil),   // 2: api.blogs.v1.CreateBlogsReply
	(*UpdateBlogsRequest)(nil), // 3: api.blogs.v1.UpdateBlogsRequest
	(*UpdateBlogsReply)(nil),   // 4: api.blogs.v1.UpdateBlogsReply
	(*DeleteBlogsRequest)(nil), // 5: api.blogs.v1.DeleteBlogsRequest
	(*DeleteBlogsReply)(nil),   // 6: api.blogs.v1.DeleteBlogsReply
	(*GetBlogsRequest)(nil),    // 7: api.blogs.v1.GetBlogsRequest
	(*BlogsHeader)(nil),        // 8: api.blogs.v1.BlogsHeader
	(*GetBlogsReply)(nil),      // 9: api.blogs.v1.GetBlogsReply
	(*ListBlogsRequest)(nil),   // 10: api.blogs.v1.ListBlogsRequest
	(*ListBlogsReply)(nil),     // 11: api.blogs.v1.ListBlogsReply
	(*global.PageInfo)(nil),    // 12: api.global.PageInfo
}
var file_api_blogs_v1_blogs_proto_depIdxs = []int32{
	0,  // 0: api.blogs.v1.CreateBlogsRequest.IsHidden:type_name -> api.blogs.v1.HiddenType
	0,  // 1: api.blogs.v1.UpdateBlogsRequest.IsHidden:type_name -> api.blogs.v1.HiddenType
	0,  // 2: api.blogs.v1.BlogsHeader.IsHidden:type_name -> api.blogs.v1.HiddenType
	8,  // 3: api.blogs.v1.GetBlogsReply.Header:type_name -> api.blogs.v1.BlogsHeader
	12, // 4: api.blogs.v1.ListBlogsRequest.Page:type_name -> api.global.PageInfo
	8,  // 5: api.blogs.v1.ListBlogsReply.List:type_name -> api.blogs.v1.BlogsHeader
	1,  // 6: api.blogs.v1.Blogs.CreateBlogs:input_type -> api.blogs.v1.CreateBlogsRequest
	3,  // 7: api.blogs.v1.Blogs.UpdateBlogs:input_type -> api.blogs.v1.UpdateBlogsRequest
	5,  // 8: api.blogs.v1.Blogs.DeleteBlogs:input_type -> api.blogs.v1.DeleteBlogsRequest
	7,  // 9: api.blogs.v1.Blogs.GetBlogs:input_type -> api.blogs.v1.GetBlogsRequest
	10, // 10: api.blogs.v1.Blogs.ListBlogs:input_type -> api.blogs.v1.ListBlogsRequest
	2,  // 11: api.blogs.v1.Blogs.CreateBlogs:output_type -> api.blogs.v1.CreateBlogsReply
	4,  // 12: api.blogs.v1.Blogs.UpdateBlogs:output_type -> api.blogs.v1.UpdateBlogsReply
	6,  // 13: api.blogs.v1.Blogs.DeleteBlogs:output_type -> api.blogs.v1.DeleteBlogsReply
	9,  // 14: api.blogs.v1.Blogs.GetBlogs:output_type -> api.blogs.v1.GetBlogsReply
	11, // 15: api.blogs.v1.Blogs.ListBlogs:output_type -> api.blogs.v1.ListBlogsReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_blogs_v1_blogs_proto_init() }
func file_api_blogs_v1_blogs_proto_init() {
	if File_api_blogs_v1_blogs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_blogs_v1_blogs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blogs_v1_blogs_proto_goTypes,
		DependencyIndexes: file_api_blogs_v1_blogs_proto_depIdxs,
		EnumInfos:         file_api_blogs_v1_blogs_proto_enumTypes,
		MessageInfos:      file_api_blogs_v1_blogs_proto_msgTypes,
	}.Build()
	File_api_blogs_v1_blogs_proto = out.File
	file_api_blogs_v1_blogs_proto_rawDesc = nil
	file_api_blogs_v1_blogs_proto_goTypes = nil
	file_api_blogs_v1_blogs_proto_depIdxs = nil
}
