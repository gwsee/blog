// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/bff/v1/palaces.proto

package v1

import (
	global "blog/api/global"
	v1 "blog/api/palaces/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_bff_v1_palaces_proto protoreflect.FileDescriptor

var file_api_bff_v1_palaces_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x0c, 0x0a, 0x07, 0x50, 0x61, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x78, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x6c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x70,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x60, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x73, 0x61,
	0x76, 0x65, 0x12, 0x52, 0x0a, 0x08, 0x44, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x2f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a,
	0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x08, 0x53, 0x61, 0x76,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01,
	0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x44, 0x6f, 0x6e, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a,
	0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x5c,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x6f, 0x6e, 0x65,
	0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2d, 0x64, 0x6f, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x60, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x6c,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x75, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x64, 0x6f, 0x6e, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x22, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x66,
	0x66, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_bff_v1_palaces_proto_goTypes = []any{
	(*v1.CreatePalacesRequest)(nil), // 0: api.palaces.v1.CreatePalacesRequest
	(*v1.UpdatePalacesRequest)(nil), // 1: api.palaces.v1.UpdatePalacesRequest
	(*v1.DeletePalacesRequest)(nil), // 2: api.palaces.v1.DeletePalacesRequest
	(*v1.GetPalacesRequest)(nil),    // 3: api.palaces.v1.GetPalacesRequest
	(*v1.ListPalacesRequest)(nil),   // 4: api.palaces.v1.ListPalacesRequest
	(*v1.SaveMemoRequest)(nil),      // 5: api.palaces.v1.SaveMemoRequest
	(*global.State)(nil),            // 6: api.global.State
	(*global.ID)(nil),               // 7: api.global.ID
	(*v1.ListMemoRequest)(nil),      // 8: api.palaces.v1.ListMemoRequest
	(*v1.SaveTodoItem)(nil),         // 9: api.palaces.v1.SaveTodoItem
	(*v1.ListTodoRequest)(nil),      // 10: api.palaces.v1.ListTodoRequest
	(*v1.CreatePalacesReply)(nil),   // 11: api.palaces.v1.CreatePalacesReply
	(*v1.UpdatePalacesReply)(nil),   // 12: api.palaces.v1.UpdatePalacesReply
	(*v1.DeletePalacesReply)(nil),   // 13: api.palaces.v1.DeletePalacesReply
	(*v1.GetPalacesReply)(nil),      // 14: api.palaces.v1.GetPalacesReply
	(*v1.ListPalacesReply)(nil),     // 15: api.palaces.v1.ListPalacesReply
	(*global.Empty)(nil),            // 16: api.global.Empty
	(*v1.ListMemoReply)(nil),        // 17: api.palaces.v1.ListMemoReply
	(*v1.ListTodoReply)(nil),        // 18: api.palaces.v1.ListTodoReply
}
var file_api_bff_v1_palaces_proto_depIdxs = []int32{
	0,  // 0: api.bff.v1.Palaces.CreatePalaces:input_type -> api.palaces.v1.CreatePalacesRequest
	1,  // 1: api.bff.v1.Palaces.UpdatePalaces:input_type -> api.palaces.v1.UpdatePalacesRequest
	2,  // 2: api.bff.v1.Palaces.DeletePalaces:input_type -> api.palaces.v1.DeletePalacesRequest
	3,  // 3: api.bff.v1.Palaces.GetPalaces:input_type -> api.palaces.v1.GetPalacesRequest
	4,  // 4: api.bff.v1.Palaces.ListPalaces:input_type -> api.palaces.v1.ListPalacesRequest
	5,  // 5: api.bff.v1.Palaces.SaveMemo:input_type -> api.palaces.v1.SaveMemoRequest
	6,  // 6: api.bff.v1.Palaces.DoneMemo:input_type -> api.global.State
	7,  // 7: api.bff.v1.Palaces.DeleteMemo:input_type -> api.global.ID
	8,  // 8: api.bff.v1.Palaces.ListMemo:input_type -> api.palaces.v1.ListMemoRequest
	9,  // 9: api.bff.v1.Palaces.SaveTodo:input_type -> api.palaces.v1.SaveTodoItem
	7,  // 10: api.bff.v1.Palaces.DoneTodo:input_type -> api.global.ID
	7,  // 11: api.bff.v1.Palaces.DeleteTodo:input_type -> api.global.ID
	7,  // 12: api.bff.v1.Palaces.DeleteTodoDone:input_type -> api.global.ID
	7,  // 13: api.bff.v1.Palaces.DeleteTodoRecord:input_type -> api.global.ID
	10, // 14: api.bff.v1.Palaces.ListTodo:input_type -> api.palaces.v1.ListTodoRequest
	10, // 15: api.bff.v1.Palaces.ListTodoDone:input_type -> api.palaces.v1.ListTodoRequest
	11, // 16: api.bff.v1.Palaces.CreatePalaces:output_type -> api.palaces.v1.CreatePalacesReply
	12, // 17: api.bff.v1.Palaces.UpdatePalaces:output_type -> api.palaces.v1.UpdatePalacesReply
	13, // 18: api.bff.v1.Palaces.DeletePalaces:output_type -> api.palaces.v1.DeletePalacesReply
	14, // 19: api.bff.v1.Palaces.GetPalaces:output_type -> api.palaces.v1.GetPalacesReply
	15, // 20: api.bff.v1.Palaces.ListPalaces:output_type -> api.palaces.v1.ListPalacesReply
	16, // 21: api.bff.v1.Palaces.SaveMemo:output_type -> api.global.Empty
	16, // 22: api.bff.v1.Palaces.DoneMemo:output_type -> api.global.Empty
	16, // 23: api.bff.v1.Palaces.DeleteMemo:output_type -> api.global.Empty
	17, // 24: api.bff.v1.Palaces.ListMemo:output_type -> api.palaces.v1.ListMemoReply
	16, // 25: api.bff.v1.Palaces.SaveTodo:output_type -> api.global.Empty
	16, // 26: api.bff.v1.Palaces.DoneTodo:output_type -> api.global.Empty
	16, // 27: api.bff.v1.Palaces.DeleteTodo:output_type -> api.global.Empty
	16, // 28: api.bff.v1.Palaces.DeleteTodoDone:output_type -> api.global.Empty
	16, // 29: api.bff.v1.Palaces.DeleteTodoRecord:output_type -> api.global.Empty
	18, // 30: api.bff.v1.Palaces.ListTodo:output_type -> api.palaces.v1.ListTodoReply
	18, // 31: api.bff.v1.Palaces.ListTodoDone:output_type -> api.palaces.v1.ListTodoReply
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_bff_v1_palaces_proto_init() }
func file_api_bff_v1_palaces_proto_init() {
	if File_api_bff_v1_palaces_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_bff_v1_palaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bff_v1_palaces_proto_goTypes,
		DependencyIndexes: file_api_bff_v1_palaces_proto_depIdxs,
	}.Build()
	File_api_bff_v1_palaces_proto = out.File
	file_api_bff_v1_palaces_proto_rawDesc = nil
	file_api_bff_v1_palaces_proto_goTypes = nil
	file_api_bff_v1_palaces_proto_depIdxs = nil
}
