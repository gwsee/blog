// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.4
// source: api/global/cfg.proto

package global

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts              []string `protobuf:"bytes,1,rep,name=Hosts,proto3" json:"Hosts,omitempty"`
	Key                string   `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Id                 int64    `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	User               string   `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
	Pass               string   `protobuf:"bytes,5,opt,name=Pass,proto3" json:"Pass,omitempty"`
	CertFile           string   `protobuf:"bytes,6,opt,name=CertFile,proto3" json:"CertFile,omitempty"`
	CertKeyFile        string   `protobuf:"bytes,7,opt,name=CertKeyFile,proto3" json:"CertKeyFile,omitempty"`
	CACertFile         string   `protobuf:"bytes,8,opt,name=CACertFile,proto3" json:"CACertFile,omitempty"`
	InsecureSkipVerify bool     `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
}

func (x *Etcd) Reset() {
	*x = Etcd{}
	mi := &file_api_global_cfg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etcd) ProtoMessage() {}

func (x *Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etcd.ProtoReflect.Descriptor instead.
func (*Etcd) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{0}
}

func (x *Etcd) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *Etcd) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Etcd) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Etcd) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Etcd) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *Etcd) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *Etcd) GetCertKeyFile() string {
	if x != nil {
		return x.CertKeyFile
	}
	return ""
}

func (x *Etcd) GetCACertFile() string {
	if x != nil {
		return x.CACertFile
	}
	return ""
}

func (x *Etcd) GetInsecureSkipVerify() bool {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return false
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceKey string `protobuf:"bytes,1,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	ApiKey     string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_api_global_cfg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{1}
}

func (x *Auth) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *Auth) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_api_global_cfg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *GRPC) Reset() {
	*x = GRPC{}
	mi := &file_api_global_cfg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPC) ProtoMessage() {}

func (x *GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPC.ProtoReflect.Descriptor instead.
func (*GRPC) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{3}
}

func (x *GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HTTP) Reset() {
	*x = HTTP{}
	mi := &file_api_global_cfg_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTP) ProtoMessage() {}

func (x *HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTP.ProtoReflect.Descriptor instead.
func (*HTTP) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{4}
}

func (x *HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	mi := &file_api_global_cfg_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{5}
}

func (x *Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_api_global_cfg_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{6}
}

func (x *Data) GetDatabase() *Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	mi := &file_api_global_cfg_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{7}
}

func (x *Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	UserName     string               `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password     string               `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Redis) Reset() {
	*x = Redis{}
	mi := &file_api_global_cfg_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redis) ProtoMessage() {}

func (x *Redis) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redis.ProtoReflect.Descriptor instead.
func (*Redis) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{8}
}

func (x *Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Redis) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Oss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Secret string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Site   string `protobuf:"bytes,5,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *Oss) Reset() {
	*x = Oss{}
	mi := &file_api_global_cfg_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Oss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oss) ProtoMessage() {}

func (x *Oss) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_cfg_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oss.ProtoReflect.Descriptor instead.
func (*Oss) Descriptor() ([]byte, []int) {
	return file_api_global_cfg_proto_rawDescGZIP(), []int{9}
}

func (x *Oss) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Oss) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Oss) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Oss) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Oss) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

var File_api_global_cfg_proto protoreflect.FileDescriptor

var file_api_global_cfg_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x1a, 0x2a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4,
	0x01, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x50, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x65, 0x72, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x65, 0x72, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x4b, 0x65,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x41, 0x43, 0x65, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x40, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x54, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x69, 0x0a,
	0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x69, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x22, 0x23, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x30, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x22, 0x3a, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a,
	0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x73, 0x0a, 0x03, 0x4f, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x42, 0x1f, 0x0a, 0x0a, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x01, 0x5a, 0x0f, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_global_cfg_proto_rawDescOnce sync.Once
	file_api_global_cfg_proto_rawDescData = file_api_global_cfg_proto_rawDesc
)

func file_api_global_cfg_proto_rawDescGZIP() []byte {
	file_api_global_cfg_proto_rawDescOnce.Do(func() {
		file_api_global_cfg_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_global_cfg_proto_rawDescData)
	})
	return file_api_global_cfg_proto_rawDescData
}

var file_api_global_cfg_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_global_cfg_proto_goTypes = []any{
	(*Etcd)(nil),                // 0: api.global.Etcd
	(*Auth)(nil),                // 1: api.global.Auth
	(*Server)(nil),              // 2: api.global.Server
	(*GRPC)(nil),                // 3: api.global.GRPC
	(*HTTP)(nil),                // 4: api.global.HTTP
	(*Trace)(nil),               // 5: api.global.Trace
	(*Data)(nil),                // 6: api.global.Data
	(*Database)(nil),            // 7: api.global.Database
	(*Redis)(nil),               // 8: api.global.Redis
	(*Oss)(nil),                 // 9: api.global.Oss
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_api_global_cfg_proto_depIdxs = []int32{
	4,  // 0: api.global.Server.http:type_name -> api.global.HTTP
	3,  // 1: api.global.Server.grpc:type_name -> api.global.GRPC
	10, // 2: api.global.GRPC.timeout:type_name -> google.protobuf.Duration
	10, // 3: api.global.HTTP.timeout:type_name -> google.protobuf.Duration
	7,  // 4: api.global.Data.database:type_name -> api.global.Database
	8,  // 5: api.global.Data.redis:type_name -> api.global.Redis
	10, // 6: api.global.Redis.read_timeout:type_name -> google.protobuf.Duration
	10, // 7: api.global.Redis.write_timeout:type_name -> google.protobuf.Duration
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_global_cfg_proto_init() }
func file_api_global_cfg_proto_init() {
	if File_api_global_cfg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_global_cfg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_global_cfg_proto_goTypes,
		DependencyIndexes: file_api_global_cfg_proto_depIdxs,
		MessageInfos:      file_api_global_cfg_proto_msgTypes,
	}.Build()
	File_api_global_cfg_proto = out.File
	file_api_global_cfg_proto_rawDesc = nil
	file_api_global_cfg_proto_goTypes = nil
	file_api_global_cfg_proto_depIdxs = nil
}
