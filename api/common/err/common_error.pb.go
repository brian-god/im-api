// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/common/err/common_error.proto

package err

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type CommonErrorReason int32

const (
	// 没有权限
	CommonErrorReason_NoPermission CommonErrorReason = 0
	// 没有数据
	CommonErrorReason_NoData CommonErrorReason = 1
)

// Enum value maps for CommonErrorReason.
var (
	CommonErrorReason_name = map[int32]string{
		0: "NoPermission",
		1: "NoData",
	}
	CommonErrorReason_value = map[string]int32{
		"NoPermission": 0,
		"NoData":       1,
	}
)

func (x CommonErrorReason) Enum() *CommonErrorReason {
	p := new(CommonErrorReason)
	*p = x
	return p
}

func (x CommonErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_err_common_error_proto_enumTypes[0].Descriptor()
}

func (CommonErrorReason) Type() protoreflect.EnumType {
	return &file_api_common_err_common_error_proto_enumTypes[0]
}

func (x CommonErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonErrorReason.Descriptor instead.
func (CommonErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_common_err_common_error_proto_rawDescGZIP(), []int{0}
}

var File_api_common_err_common_error_proto protoreflect.FileDescriptor

var file_api_common_err_common_error_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x72, 0x72, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x43, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x0c, 0x4e, 0x6f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x1a,
	0x04, 0xa8, 0x45, 0x92, 0x03, 0x12, 0x10, 0x0a, 0x06, 0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x10,
	0x01, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x43, 0x0a,
	0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x50,
	0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72,
	0x69, 0x61, 0x6e, 0x2d, 0x67, 0x6f, 0x64, 0x2f, 0x78, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x3b, 0x65,
	0x72, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_err_common_error_proto_rawDescOnce sync.Once
	file_api_common_err_common_error_proto_rawDescData = file_api_common_err_common_error_proto_rawDesc
)

func file_api_common_err_common_error_proto_rawDescGZIP() []byte {
	file_api_common_err_common_error_proto_rawDescOnce.Do(func() {
		file_api_common_err_common_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_err_common_error_proto_rawDescData)
	})
	return file_api_common_err_common_error_proto_rawDescData
}

var file_api_common_err_common_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_common_err_common_error_proto_goTypes = []interface{}{
	(CommonErrorReason)(0), // 0: api.common.err.CommonErrorReason
}
var file_api_common_err_common_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_common_err_common_error_proto_init() }
func file_api_common_err_common_error_proto_init() {
	if File_api_common_err_common_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_common_err_common_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_err_common_error_proto_goTypes,
		DependencyIndexes: file_api_common_err_common_error_proto_depIdxs,
		EnumInfos:         file_api_common_err_common_error_proto_enumTypes,
	}.Build()
	File_api_common_err_common_error_proto = out.File
	file_api_common_err_common_error_proto_rawDesc = nil
	file_api_common_err_common_error_proto_goTypes = nil
	file_api_common_err_common_error_proto_depIdxs = nil
}
