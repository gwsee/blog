// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: app/palaces/internal/conf/conf.proto

package conf

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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *global.Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *global.Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Auth   *global.Auth   `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Trace  *global.Trace  `protobuf:"bytes,4,opt,name=trace,proto3" json:"trace,omitempty"`
	Etcd   *global.Etcd   `protobuf:"bytes,5,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_app_palaces_internal_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_app_palaces_internal_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_app_palaces_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *global.Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *global.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetAuth() *global.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Bootstrap) GetTrace() *global.Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetEtcd() *global.Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

var File_app_palaces_internal_conf_conf_proto protoreflect.FileDescriptor

var file_app_palaces_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x61, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65,
	0x74, 0x63, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_palaces_internal_conf_conf_proto_rawDescOnce sync.Once
	file_app_palaces_internal_conf_conf_proto_rawDescData = file_app_palaces_internal_conf_conf_proto_rawDesc
)

func file_app_palaces_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_app_palaces_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_app_palaces_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_palaces_internal_conf_conf_proto_rawDescData)
	})
	return file_app_palaces_internal_conf_conf_proto_rawDescData
}

var file_app_palaces_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_palaces_internal_conf_conf_proto_goTypes = []any{
	(*Bootstrap)(nil),     // 0: app.palaces.conf.Bootstrap
	(*global.Server)(nil), // 1: api.global.Server
	(*global.Data)(nil),   // 2: api.global.Data
	(*global.Auth)(nil),   // 3: api.global.Auth
	(*global.Trace)(nil),  // 4: api.global.Trace
	(*global.Etcd)(nil),   // 5: api.global.Etcd
}
var file_app_palaces_internal_conf_conf_proto_depIdxs = []int32{
	1, // 0: app.palaces.conf.Bootstrap.server:type_name -> api.global.Server
	2, // 1: app.palaces.conf.Bootstrap.data:type_name -> api.global.Data
	3, // 2: app.palaces.conf.Bootstrap.auth:type_name -> api.global.Auth
	4, // 3: app.palaces.conf.Bootstrap.trace:type_name -> api.global.Trace
	5, // 4: app.palaces.conf.Bootstrap.etcd:type_name -> api.global.Etcd
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_app_palaces_internal_conf_conf_proto_init() }
func file_app_palaces_internal_conf_conf_proto_init() {
	if File_app_palaces_internal_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_palaces_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_palaces_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_app_palaces_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_app_palaces_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_app_palaces_internal_conf_conf_proto = out.File
	file_app_palaces_internal_conf_conf_proto_rawDesc = nil
	file_app_palaces_internal_conf_conf_proto_goTypes = nil
	file_app_palaces_internal_conf_conf_proto_depIdxs = nil
}
