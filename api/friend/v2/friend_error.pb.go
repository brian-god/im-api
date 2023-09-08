// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/friend/v2/friend_error.proto

package v2

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

type FriendErrorReason int32

const (
	// 获取好友信息失败
	FriendErrorReason_GetFriendList_FAIL FriendErrorReason = 0
	// 获取好友信息失败
	FriendErrorReason_GetFriendsInfo_FAIL FriendErrorReason = 1
	// 添加好友失败
	FriendErrorReason_AddFriend_FAIL FriendErrorReason = 2
	// 获取好友申请列表
	FriendErrorReason_GetFriendApplyList_FAIL FriendErrorReason = 3
	// 添加黑名单
	FriendErrorReason_AddBlacklist_FAIL FriendErrorReason = 4
	// 删除黑名单
	FriendErrorReason_RemoveBlacklist_FAIL FriendErrorReason = 5
	// 是否好友
	FriendErrorReason_IsFriend_FAIL FriendErrorReason = 6
	// 是否黑名单
	FriendErrorReason_IsInBlackList_FAIL FriendErrorReason = 7
	// 获取黑名单列表
	FriendErrorReason_GetBlacklist_FAIL FriendErrorReason = 8
	// 删除好友
	FriendErrorReason_DeleteFriend_FAIL FriendErrorReason = 9
	// 删除好友
	FriendErrorReason_AddFriendResponse_FAIL FriendErrorReason = 10
	// 删除好友
	FriendErrorReason_SetFriendRemark_FAIL FriendErrorReason = 11
)

// Enum value maps for FriendErrorReason.
var (
	FriendErrorReason_name = map[int32]string{
		0:  "GetFriendList_FAIL",
		1:  "GetFriendsInfo_FAIL",
		2:  "AddFriend_FAIL",
		3:  "GetFriendApplyList_FAIL",
		4:  "AddBlacklist_FAIL",
		5:  "RemoveBlacklist_FAIL",
		6:  "IsFriend_FAIL",
		7:  "IsInBlackList_FAIL",
		8:  "GetBlacklist_FAIL",
		9:  "DeleteFriend_FAIL",
		10: "AddFriendResponse_FAIL",
		11: "SetFriendRemark_FAIL",
	}
	FriendErrorReason_value = map[string]int32{
		"GetFriendList_FAIL":      0,
		"GetFriendsInfo_FAIL":     1,
		"AddFriend_FAIL":          2,
		"GetFriendApplyList_FAIL": 3,
		"AddBlacklist_FAIL":       4,
		"RemoveBlacklist_FAIL":    5,
		"IsFriend_FAIL":           6,
		"IsInBlackList_FAIL":      7,
		"GetBlacklist_FAIL":       8,
		"DeleteFriend_FAIL":       9,
		"AddFriendResponse_FAIL":  10,
		"SetFriendRemark_FAIL":    11,
	}
)

func (x FriendErrorReason) Enum() *FriendErrorReason {
	p := new(FriendErrorReason)
	*p = x
	return p
}

func (x FriendErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FriendErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_friend_v2_friend_error_proto_enumTypes[0].Descriptor()
}

func (FriendErrorReason) Type() protoreflect.EnumType {
	return &file_api_friend_v2_friend_error_proto_enumTypes[0]
}

func (x FriendErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FriendErrorReason.Descriptor instead.
func (FriendErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_friend_v2_friend_error_proto_rawDescGZIP(), []int{0}
}

var File_api_friend_v2_friend_error_proto protoreflect.FileDescriptor

var file_api_friend_v2_friend_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x32, 0x2f,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x83, 0x03, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x1a, 0x04, 0xa8,
	0x45, 0x9a, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x03,
	0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x04, 0x1a, 0x04, 0xa8,
	0x45, 0x9a, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x05, 0x1a, 0x04, 0xa8,
	0x45, 0x9a, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x49, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1c, 0x0a, 0x12,
	0x49, 0x73, 0x49, 0x6e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x08, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x09, 0x1a, 0x04,
	0xa8, 0x45, 0x9a, 0x03, 0x12, 0x20, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x0a,
	0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x0b,
	0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x40, 0x0a, 0x0d,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x69, 0x61,
	0x6e, 0x2d, 0x67, 0x6f, 0x64, 0x2f, 0x78, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_friend_v2_friend_error_proto_rawDescOnce sync.Once
	file_api_friend_v2_friend_error_proto_rawDescData = file_api_friend_v2_friend_error_proto_rawDesc
)

func file_api_friend_v2_friend_error_proto_rawDescGZIP() []byte {
	file_api_friend_v2_friend_error_proto_rawDescOnce.Do(func() {
		file_api_friend_v2_friend_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_friend_v2_friend_error_proto_rawDescData)
	})
	return file_api_friend_v2_friend_error_proto_rawDescData
}

var file_api_friend_v2_friend_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_friend_v2_friend_error_proto_goTypes = []interface{}{
	(FriendErrorReason)(0), // 0: api.friend.v1.FriendErrorReason
}
var file_api_friend_v2_friend_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_friend_v2_friend_error_proto_init() }
func file_api_friend_v2_friend_error_proto_init() {
	if File_api_friend_v2_friend_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_friend_v2_friend_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_friend_v2_friend_error_proto_goTypes,
		DependencyIndexes: file_api_friend_v2_friend_error_proto_depIdxs,
		EnumInfos:         file_api_friend_v2_friend_error_proto_enumTypes,
	}.Build()
	File_api_friend_v2_friend_error_proto = out.File
	file_api_friend_v2_friend_error_proto_rawDesc = nil
	file_api_friend_v2_friend_error_proto_goTypes = nil
	file_api_friend_v2_friend_error_proto_depIdxs = nil
}
