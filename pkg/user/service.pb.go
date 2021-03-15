//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/user/pb/service.proto

package user

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infraboard/mcube/pb/http"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_pkg_user_pb_service_proto protoreflect.FileDescriptor

var file_pkg_user_pb_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xce, 0x08,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x22, 0x2e, 0xa2, 0xa3, 0x8c, 0x4d, 0x29, 0x1a, 0x06, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x42,
	0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x37, 0xa2, 0xa3, 0x8c, 0x4d, 0x32, 0x1a, 0x0f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x3a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x42, 0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x78, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x2f, 0xa2, 0xa3, 0x8c, 0x4d, 0x2a, 0x1a, 0x06, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x2a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01,
	0x42, 0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x3e, 0xa2, 0xa3,
	0x8c, 0x4d, 0x39, 0x1a, 0x15, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x3a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x04, 0x50, 0x4f, 0x53, 0x54,
	0x2a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x42, 0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x83, 0x01, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x3a, 0xa2, 0xa3, 0x8c, 0x4d, 0x35, 0x1a, 0x0f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x3a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x2a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x42, 0x12,
	0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x87, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x37, 0xa2, 0xa3, 0x8c, 0x4d, 0x32, 0x1a, 0x0f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x3a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x03, 0x50, 0x55, 0x54,
	0x2a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x42, 0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x9a, 0x01, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x44, 0xa2, 0xa3, 0x8c, 0x4d, 0x3f, 0x1a, 0x18, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x3a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x03, 0x50, 0x55, 0x54, 0x2a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x30, 0x01, 0x42, 0x12, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x09,
	0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x92, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0xa2,
	0xa3, 0x8c, 0x4d, 0x2a, 0x1a, 0x0a, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x2a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x30, 0x01, 0x42, 0x0a, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x01, 0x2a, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_user_pb_service_proto_goTypes = []interface{}{
	(*QueryAccountRequest)(nil),      // 0: keyauth.user.QueryAccountRequest
	(*DescribeAccountRequest)(nil),   // 1: keyauth.user.DescribeAccountRequest
	(*CreateAccountRequest)(nil),     // 2: keyauth.user.CreateAccountRequest
	(*BlockAccountRequest)(nil),      // 3: keyauth.user.BlockAccountRequest
	(*DeleteAccountRequest)(nil),     // 4: keyauth.user.DeleteAccountRequest
	(*UpdateAccountRequest)(nil),     // 5: keyauth.user.UpdateAccountRequest
	(*UpdatePasswordRequest)(nil),    // 6: keyauth.user.UpdatePasswordRequest
	(*GeneratePasswordRequest)(nil),  // 7: keyauth.user.GeneratePasswordRequest
	(*Set)(nil),                      // 8: keyauth.user.Set
	(*User)(nil),                     // 9: keyauth.user.User
	(*Password)(nil),                 // 10: keyauth.user.Password
	(*GeneratePasswordResponse)(nil), // 11: keyauth.user.GeneratePasswordResponse
}
var file_pkg_user_pb_service_proto_depIdxs = []int32{
	0,  // 0: keyauth.user.UserService.QueryAccount:input_type -> keyauth.user.QueryAccountRequest
	1,  // 1: keyauth.user.UserService.DescribeAccount:input_type -> keyauth.user.DescribeAccountRequest
	2,  // 2: keyauth.user.UserService.CreateAccount:input_type -> keyauth.user.CreateAccountRequest
	3,  // 3: keyauth.user.UserService.BlockAccount:input_type -> keyauth.user.BlockAccountRequest
	4,  // 4: keyauth.user.UserService.DeleteAccount:input_type -> keyauth.user.DeleteAccountRequest
	5,  // 5: keyauth.user.UserService.UpdateAccountProfile:input_type -> keyauth.user.UpdateAccountRequest
	6,  // 6: keyauth.user.UserService.UpdateAccountPassword:input_type -> keyauth.user.UpdatePasswordRequest
	7,  // 7: keyauth.user.UserService.GeneratePassword:input_type -> keyauth.user.GeneratePasswordRequest
	8,  // 8: keyauth.user.UserService.QueryAccount:output_type -> keyauth.user.Set
	9,  // 9: keyauth.user.UserService.DescribeAccount:output_type -> keyauth.user.User
	9,  // 10: keyauth.user.UserService.CreateAccount:output_type -> keyauth.user.User
	9,  // 11: keyauth.user.UserService.BlockAccount:output_type -> keyauth.user.User
	9,  // 12: keyauth.user.UserService.DeleteAccount:output_type -> keyauth.user.User
	9,  // 13: keyauth.user.UserService.UpdateAccountProfile:output_type -> keyauth.user.User
	10, // 14: keyauth.user.UserService.UpdateAccountPassword:output_type -> keyauth.user.Password
	11, // 15: keyauth.user.UserService.GeneratePassword:output_type -> keyauth.user.GeneratePasswordResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_user_pb_service_proto_init() }
func file_pkg_user_pb_service_proto_init() {
	if File_pkg_user_pb_service_proto != nil {
		return
	}
	file_pkg_user_pb_request_proto_init()
	file_pkg_user_pb_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_user_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_user_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_user_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_user_pb_service_proto = out.File
	file_pkg_user_pb_service_proto_rawDesc = nil
	file_pkg_user_pb_service_proto_goTypes = nil
	file_pkg_user_pb_service_proto_depIdxs = nil
}
