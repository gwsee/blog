// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/bff/v1/user.proto

package v1

import (
	global "blog/api/global"
	v1 "blog/api/user/v1"
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

var File_api_bff_v1_user_proto protoreflect.FileDescriptor

var file_api_bff_v1_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66,
	0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x55, 0x0a,
	0x08, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x73, 0x61, 0x76, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x5e, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x12, 0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x6a,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x0e, 0x53, 0x61,
	0x76, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x61, 0x76, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x12, 0x5f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x76, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x51, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x55, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x22, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x12, 0x62, 0x6c,
	0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_bff_v1_user_proto_goTypes = []any{
	(*v1.SaveUserRequest)(nil),       // 0: api.user.v1.SaveUserRequest
	(*global.Empty)(nil),             // 1: api.global.Empty
	(*v1.SaveProjectRequest)(nil),    // 2: api.user.v1.SaveProjectRequest
	(*global.ID)(nil),                // 3: api.global.ID
	(*v1.ListProjectRequest)(nil),    // 4: api.user.v1.ListProjectRequest
	(*v1.SaveExperienceRequest)(nil), // 5: api.user.v1.SaveExperienceRequest
	(*v1.ListExperienceRequest)(nil), // 6: api.user.v1.ListExperienceRequest
	(*v1.PhotosReq)(nil),             // 7: api.user.v1.PhotosReq
	(*global.PageInfo)(nil),          // 8: api.global.PageInfo
	(*v1.GetUserReply)(nil),          // 9: api.user.v1.GetUserReply
	(*v1.GetProjectReply)(nil),       // 10: api.user.v1.GetProjectReply
	(*v1.ListProjectReply)(nil),      // 11: api.user.v1.ListProjectReply
	(*v1.GetExperienceReply)(nil),    // 12: api.user.v1.GetExperienceReply
	(*v1.ListExperienceReply)(nil),   // 13: api.user.v1.ListExperienceReply
	(*v1.PhotosReply)(nil),           // 14: api.user.v1.PhotosReply
	(*v1.MessagesReply)(nil),         // 15: api.user.v1.MessagesReply
}
var file_api_bff_v1_user_proto_depIdxs = []int32{
	0,  // 0: api.bff.v1.User.SaveUser:input_type -> api.user.v1.SaveUserRequest
	1,  // 1: api.bff.v1.User.GetUser:input_type -> api.global.Empty
	2,  // 2: api.bff.v1.User.SaveProject:input_type -> api.user.v1.SaveProjectRequest
	3,  // 3: api.bff.v1.User.DeleteProject:input_type -> api.global.ID
	3,  // 4: api.bff.v1.User.GetProject:input_type -> api.global.ID
	4,  // 5: api.bff.v1.User.ListProject:input_type -> api.user.v1.ListProjectRequest
	5,  // 6: api.bff.v1.User.SaveExperience:input_type -> api.user.v1.SaveExperienceRequest
	3,  // 7: api.bff.v1.User.DeleteExperience:input_type -> api.global.ID
	3,  // 8: api.bff.v1.User.GetExperience:input_type -> api.global.ID
	6,  // 9: api.bff.v1.User.ListExperience:input_type -> api.user.v1.ListExperienceRequest
	7,  // 10: api.bff.v1.User.Photos:input_type -> api.user.v1.PhotosReq
	8,  // 11: api.bff.v1.User.Messages:input_type -> api.global.PageInfo
	1,  // 12: api.bff.v1.User.SaveUser:output_type -> api.global.Empty
	9,  // 13: api.bff.v1.User.GetUser:output_type -> api.user.v1.GetUserReply
	1,  // 14: api.bff.v1.User.SaveProject:output_type -> api.global.Empty
	1,  // 15: api.bff.v1.User.DeleteProject:output_type -> api.global.Empty
	10, // 16: api.bff.v1.User.GetProject:output_type -> api.user.v1.GetProjectReply
	11, // 17: api.bff.v1.User.ListProject:output_type -> api.user.v1.ListProjectReply
	1,  // 18: api.bff.v1.User.SaveExperience:output_type -> api.global.Empty
	1,  // 19: api.bff.v1.User.DeleteExperience:output_type -> api.global.Empty
	12, // 20: api.bff.v1.User.GetExperience:output_type -> api.user.v1.GetExperienceReply
	13, // 21: api.bff.v1.User.ListExperience:output_type -> api.user.v1.ListExperienceReply
	14, // 22: api.bff.v1.User.Photos:output_type -> api.user.v1.PhotosReply
	15, // 23: api.bff.v1.User.Messages:output_type -> api.user.v1.MessagesReply
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_bff_v1_user_proto_init() }
func file_api_bff_v1_user_proto_init() {
	if File_api_bff_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_bff_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bff_v1_user_proto_goTypes,
		DependencyIndexes: file_api_bff_v1_user_proto_depIdxs,
	}.Build()
	File_api_bff_v1_user_proto = out.File
	file_api_bff_v1_user_proto_rawDesc = nil
	file_api_bff_v1_user_proto_goTypes = nil
	file_api_bff_v1_user_proto_depIdxs = nil
}
