// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/account/v1/account.proto

package v1

import (
	global "blog/api/global"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	NickName    string `protobuf:"bytes,4,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Avatar      string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_api_account_v1_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateAccountRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CreateAccountRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	mi := &file_api_account_v1_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *ResetPasswordRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ResetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginByAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginByAccountRequest) Reset() {
	*x = LoginByAccountRequest{}
	mi := &file_api_account_v1_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginByAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByAccountRequest) ProtoMessage() {}

func (x *LoginByAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByAccountRequest.ProtoReflect.Descriptor instead.
func (*LoginByAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *LoginByAccountRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginByAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginByAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *LoginByAccountReply) Reset() {
	*x = LoginByAccountReply{}
	mi := &file_api_account_v1_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginByAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByAccountReply) ProtoMessage() {}

func (x *LoginByAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByAccountReply.ProtoReflect.Descriptor instead.
func (*LoginByAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *LoginByAccountReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName    string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Avatar      string `protobuf:"bytes,2,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	mi := &file_api_account_v1_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAccountRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UpdateAccountRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateAccountRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AccountInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NickName    string `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Account     string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Avatar      string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	BlogNum     int64  `protobuf:"varint,7,opt,name=BlogNum,proto3" json:"BlogNum,omitempty"`
}

func (x *AccountInfoReply) Reset() {
	*x = AccountInfoReply{}
	mi := &file_api_account_v1_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfoReply) ProtoMessage() {}

func (x *AccountInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfoReply.ProtoReflect.Descriptor instead.
func (*AccountInfoReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *AccountInfoReply) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AccountInfoReply) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *AccountInfoReply) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountInfoReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountInfoReply) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AccountInfoReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AccountInfoReply) GetBlogNum() int64 {
	if x != nil {
		return x.BlogNum
	}
	return 0
}

var File_api_account_v1_account_proto protoreflect.FileDescriptor

var file_api_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x14, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x67, 0x4e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x42, 0x6c, 0x6f, 0x67, 0x4e, 0x75, 0x6d, 0x32, 0xb0, 0x03, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x09,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x09, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x12, 0x67,
	0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x09, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x09, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x42, 0x2a, 0x0a, 0x0e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x16, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_v1_account_proto_rawDescOnce sync.Once
	file_api_account_v1_account_proto_rawDescData = file_api_account_v1_account_proto_rawDesc
)

func file_api_account_v1_account_proto_rawDescGZIP() []byte {
	file_api_account_v1_account_proto_rawDescOnce.Do(func() {
		file_api_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_v1_account_proto_rawDescData)
	})
	return file_api_account_v1_account_proto_rawDescData
}

var file_api_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_account_v1_account_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),  // 0: api.account.v1.CreateAccountRequest
	(*ResetPasswordRequest)(nil),  // 1: api.account.v1.ResetPasswordRequest
	(*LoginByAccountRequest)(nil), // 2: api.account.v1.LoginByAccountRequest
	(*LoginByAccountReply)(nil),   // 3: api.account.v1.LoginByAccountReply
	(*UpdateAccountRequest)(nil),  // 4: api.account.v1.UpdateAccountRequest
	(*AccountInfoReply)(nil),      // 5: api.account.v1.AccountInfoReply
	(*global.Empty)(nil),          // 6: api.global.Empty
}
var file_api_account_v1_account_proto_depIdxs = []int32{
	0, // 0: api.account.v1.Account.CreateAccount:input_type -> api.account.v1.CreateAccountRequest
	1, // 1: api.account.v1.Account.ResetPassword:input_type -> api.account.v1.ResetPasswordRequest
	2, // 2: api.account.v1.Account.LoginByAccount:input_type -> api.account.v1.LoginByAccountRequest
	6, // 3: api.account.v1.Account.Info:input_type -> api.global.Empty
	4, // 4: api.account.v1.Account.UpdateAccount:input_type -> api.account.v1.UpdateAccountRequest
	6, // 5: api.account.v1.Account.CreateAccount:output_type -> api.global.Empty
	6, // 6: api.account.v1.Account.ResetPassword:output_type -> api.global.Empty
	3, // 7: api.account.v1.Account.LoginByAccount:output_type -> api.account.v1.LoginByAccountReply
	5, // 8: api.account.v1.Account.Info:output_type -> api.account.v1.AccountInfoReply
	6, // 9: api.account.v1.Account.UpdateAccount:output_type -> api.global.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_account_v1_account_proto_init() }
func file_api_account_v1_account_proto_init() {
	if File_api_account_v1_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_account_v1_account_proto_goTypes,
		DependencyIndexes: file_api_account_v1_account_proto_depIdxs,
		MessageInfos:      file_api_account_v1_account_proto_msgTypes,
	}.Build()
	File_api_account_v1_account_proto = out.File
	file_api_account_v1_account_proto_rawDesc = nil
	file_api_account_v1_account_proto_goTypes = nil
	file_api_account_v1_account_proto_depIdxs = nil
}
