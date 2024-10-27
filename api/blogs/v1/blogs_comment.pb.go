// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/blogs/v1/blogs_comment.proto

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

type CreateBlogsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId   int64  `protobuf:"varint,1,opt,name=BlogId,proto3" json:"BlogId,omitempty"`     //博客的ID
	TopId    int64  `protobuf:"varint,2,opt,name=TopId,proto3" json:"TopId,omitempty"`       //顶级评论
	ParentId int64  `protobuf:"varint,3,opt,name=ParentId,proto3" json:"ParentId,omitempty"` //回复的那个评论
	Level    int64  `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`       //第几楼
	Total    int64  `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`       //总数
	Status   int64  `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`     //是否隐藏
	Content  string `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`    //评论内容
}

func (x *CreateBlogsCommentRequest) Reset() {
	*x = CreateBlogsCommentRequest{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogsCommentRequest) ProtoMessage() {}

func (x *CreateBlogsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogsCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogsCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlogsCommentRequest) GetBlogId() int64 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetTopId() int64 {
	if x != nil {
		return x.TopId
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateBlogsCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateBlogsCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBlogsCommentReply) Reset() {
	*x = CreateBlogsCommentReply{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogsCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogsCommentReply) ProtoMessage() {}

func (x *CreateBlogsCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogsCommentReply.ProtoReflect.Descriptor instead.
func (*CreateBlogsCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{1}
}

type UpdateBlogsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`          //评论的ID
	Status  int64  `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`  //是否隐藏
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"` //评论内容
}

func (x *UpdateBlogsCommentRequest) Reset() {
	*x = UpdateBlogsCommentRequest{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogsCommentRequest) ProtoMessage() {}

func (x *UpdateBlogsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogsCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogsCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBlogsCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBlogsCommentRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateBlogsCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateBlogsCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBlogsCommentReply) Reset() {
	*x = UpdateBlogsCommentReply{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogsCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogsCommentReply) ProtoMessage() {}

func (x *UpdateBlogsCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogsCommentReply.ProtoReflect.Descriptor instead.
func (*UpdateBlogsCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{3}
}

type DeleteBlogsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` //评论的ID
}

func (x *DeleteBlogsCommentRequest) Reset() {
	*x = DeleteBlogsCommentRequest{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogsCommentRequest) ProtoMessage() {}

func (x *DeleteBlogsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogsCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlogsCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBlogsCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBlogsCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBlogsCommentReply) Reset() {
	*x = DeleteBlogsCommentReply{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogsCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogsCommentReply) ProtoMessage() {}

func (x *DeleteBlogsCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogsCommentReply.ProtoReflect.Descriptor instead.
func (*DeleteBlogsCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{5}
}

type GetBlogsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` //评论的ID
}

func (x *GetBlogsCommentRequest) Reset() {
	*x = GetBlogsCommentRequest{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsCommentRequest) ProtoMessage() {}

func (x *GetBlogsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsCommentRequest.ProtoReflect.Descriptor instead.
func (*GetBlogsCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetBlogsCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBlogsCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`             //评论的ID
	BlogId   int64  `protobuf:"varint,2,opt,name=BlogId,proto3" json:"BlogId,omitempty"`     //博客的ID
	TopId    int64  `protobuf:"varint,3,opt,name=TopId,proto3" json:"TopId,omitempty"`       //顶级评论
	ParentId int64  `protobuf:"varint,4,opt,name=ParentId,proto3" json:"ParentId,omitempty"` //回复的那个评论
	Level    int64  `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty"`       //第几楼
	Total    int64  `protobuf:"varint,6,opt,name=Total,proto3" json:"Total,omitempty"`       //总数
	Status   int64  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`     //是否隐藏
	Content  string `protobuf:"bytes,8,opt,name=Content,proto3" json:"Content,omitempty"`    //评论内容
}

func (x *GetBlogsCommentReply) Reset() {
	*x = GetBlogsCommentReply{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogsCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogsCommentReply) ProtoMessage() {}

func (x *GetBlogsCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogsCommentReply.ProtoReflect.Descriptor instead.
func (*GetBlogsCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlogsCommentReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBlogsCommentReply) GetBlogId() int64 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *GetBlogsCommentReply) GetTopId() int64 {
	if x != nil {
		return x.TopId
	}
	return 0
}

func (x *GetBlogsCommentReply) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetBlogsCommentReply) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetBlogsCommentReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetBlogsCommentReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetBlogsCommentReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ListBlogsCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     *global.PageIndex `protobuf:"bytes,1,opt,name=Page,proto3" json:"Page,omitempty"`          //分页信息
	BlogId   int64             `protobuf:"varint,2,opt,name=BlogId,proto3" json:"BlogId,omitempty"`     //博客的ID
	ParentId int64             `protobuf:"varint,3,opt,name=ParentId,proto3" json:"ParentId,omitempty"` //回复的那个评论的ID
	Content  string            `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`    //评论内容
}

func (x *ListBlogsCommentRequest) Reset() {
	*x = ListBlogsCommentRequest{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsCommentRequest) ProtoMessage() {}

func (x *ListBlogsCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsCommentRequest.ProtoReflect.Descriptor instead.
func (*ListBlogsCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{8}
}

func (x *ListBlogsCommentRequest) GetPage() *global.PageIndex {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListBlogsCommentRequest) GetBlogId() int64 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *ListBlogsCommentRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListBlogsCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ListBlogsCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	End  int64                   `protobuf:"varint,1,opt,name=End,proto3" json:"End,omitempty"`
	List []*GetBlogsCommentReply `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *ListBlogsCommentReply) Reset() {
	*x = ListBlogsCommentReply{}
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsCommentReply) ProtoMessage() {}

func (x *ListBlogsCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blogs_v1_blogs_comment_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsCommentReply.ProtoReflect.Descriptor instead.
func (*ListBlogsCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blogs_v1_blogs_comment_proto_rawDescGZIP(), []int{9}
}

func (x *ListBlogsCommentReply) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ListBlogsCommentReply) GetList() []*GetBlogsCommentReply {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_blogs_v1_blogs_comment_proto protoreflect.FileDescriptor

var file_api_blogs_v1_blogs_comment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f,
	0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x19, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5d, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x42,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x42, 0x6c,
	0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xfd, 0x03,
	0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x64,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x5b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x26, 0x0a,
	0x0c, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x14, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blogs_v1_blogs_comment_proto_rawDescOnce sync.Once
	file_api_blogs_v1_blogs_comment_proto_rawDescData = file_api_blogs_v1_blogs_comment_proto_rawDesc
)

func file_api_blogs_v1_blogs_comment_proto_rawDescGZIP() []byte {
	file_api_blogs_v1_blogs_comment_proto_rawDescOnce.Do(func() {
		file_api_blogs_v1_blogs_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blogs_v1_blogs_comment_proto_rawDescData)
	})
	return file_api_blogs_v1_blogs_comment_proto_rawDescData
}

var file_api_blogs_v1_blogs_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_blogs_v1_blogs_comment_proto_goTypes = []any{
	(*CreateBlogsCommentRequest)(nil), // 0: api.blogs.v1.CreateBlogsCommentRequest
	(*CreateBlogsCommentReply)(nil),   // 1: api.blogs.v1.CreateBlogsCommentReply
	(*UpdateBlogsCommentRequest)(nil), // 2: api.blogs.v1.UpdateBlogsCommentRequest
	(*UpdateBlogsCommentReply)(nil),   // 3: api.blogs.v1.UpdateBlogsCommentReply
	(*DeleteBlogsCommentRequest)(nil), // 4: api.blogs.v1.DeleteBlogsCommentRequest
	(*DeleteBlogsCommentReply)(nil),   // 5: api.blogs.v1.DeleteBlogsCommentReply
	(*GetBlogsCommentRequest)(nil),    // 6: api.blogs.v1.GetBlogsCommentRequest
	(*GetBlogsCommentReply)(nil),      // 7: api.blogs.v1.GetBlogsCommentReply
	(*ListBlogsCommentRequest)(nil),   // 8: api.blogs.v1.ListBlogsCommentRequest
	(*ListBlogsCommentReply)(nil),     // 9: api.blogs.v1.ListBlogsCommentReply
	(*global.PageIndex)(nil),          // 10: api.global.PageIndex
}
var file_api_blogs_v1_blogs_comment_proto_depIdxs = []int32{
	10, // 0: api.blogs.v1.ListBlogsCommentRequest.Page:type_name -> api.global.PageIndex
	7,  // 1: api.blogs.v1.ListBlogsCommentReply.List:type_name -> api.blogs.v1.GetBlogsCommentReply
	0,  // 2: api.blogs.v1.BlogsComment.CreateBlogsComment:input_type -> api.blogs.v1.CreateBlogsCommentRequest
	2,  // 3: api.blogs.v1.BlogsComment.UpdateBlogsComment:input_type -> api.blogs.v1.UpdateBlogsCommentRequest
	4,  // 4: api.blogs.v1.BlogsComment.DeleteBlogsComment:input_type -> api.blogs.v1.DeleteBlogsCommentRequest
	6,  // 5: api.blogs.v1.BlogsComment.GetBlogsComment:input_type -> api.blogs.v1.GetBlogsCommentRequest
	8,  // 6: api.blogs.v1.BlogsComment.ListBlogsComment:input_type -> api.blogs.v1.ListBlogsCommentRequest
	1,  // 7: api.blogs.v1.BlogsComment.CreateBlogsComment:output_type -> api.blogs.v1.CreateBlogsCommentReply
	3,  // 8: api.blogs.v1.BlogsComment.UpdateBlogsComment:output_type -> api.blogs.v1.UpdateBlogsCommentReply
	5,  // 9: api.blogs.v1.BlogsComment.DeleteBlogsComment:output_type -> api.blogs.v1.DeleteBlogsCommentReply
	7,  // 10: api.blogs.v1.BlogsComment.GetBlogsComment:output_type -> api.blogs.v1.GetBlogsCommentReply
	9,  // 11: api.blogs.v1.BlogsComment.ListBlogsComment:output_type -> api.blogs.v1.ListBlogsCommentReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_blogs_v1_blogs_comment_proto_init() }
func file_api_blogs_v1_blogs_comment_proto_init() {
	if File_api_blogs_v1_blogs_comment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_blogs_v1_blogs_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blogs_v1_blogs_comment_proto_goTypes,
		DependencyIndexes: file_api_blogs_v1_blogs_comment_proto_depIdxs,
		MessageInfos:      file_api_blogs_v1_blogs_comment_proto_msgTypes,
	}.Build()
	File_api_blogs_v1_blogs_comment_proto = out.File
	file_api_blogs_v1_blogs_comment_proto_rawDesc = nil
	file_api_blogs_v1_blogs_comment_proto_goTypes = nil
	file_api_blogs_v1_blogs_comment_proto_depIdxs = nil
}
