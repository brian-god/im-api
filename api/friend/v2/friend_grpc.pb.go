// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: friend/v2/friend.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Friend_GetFriendList_FullMethodName      = "/api.friend.v2.Friend/GetFriendList"
	Friend_GetFriendsInfo_FullMethodName     = "/api.friend.v2.Friend/GetFriendsInfo"
	Friend_AddFriend_FullMethodName          = "/api.friend.v2.Friend/AddFriend"
	Friend_GetFriendApplyList_FullMethodName = "/api.friend.v2.Friend/GetFriendApplyList"
	Friend_AddBlacklist_FullMethodName       = "/api.friend.v2.Friend/AddBlacklist"
	Friend_RemoveBlacklist_FullMethodName    = "/api.friend.v2.Friend/RemoveBlacklist"
	Friend_IsFriend_FullMethodName           = "/api.friend.v2.Friend/IsFriend"
	Friend_IsInBlackList_FullMethodName      = "/api.friend.v2.Friend/IsInBlackList"
	Friend_GetBlacklist_FullMethodName       = "/api.friend.v2.Friend/GetBlacklist"
	Friend_DeleteFriend_FullMethodName       = "/api.friend.v2.Friend/DeleteFriend"
	Friend_AddFriendResponse_FullMethodName  = "/api.friend.v2.Friend/AddFriendResponse"
	Friend_SetFriendRemark_FullMethodName    = "/api.friend.v2.Friend/SetFriendRemark"
)

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	//获取好友列表仅仅返回id昵称和头像
	GetFriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListReply, error)
	// 获取好友信息
	GetFriendsInfo(ctx context.Context, in *GetFriendsInfoReq, opts ...grpc.CallOption) (*GetFriendInfoReply, error)
	// 添加好友
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendReply, error)
	// 获取好友申请列表
	GetFriendApplyList(ctx context.Context, in *GetFriendApplyListReq, opts ...grpc.CallOption) (*GetFriendApplyListReply, error)
	// 添加黑名单
	AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistReply, error)
	// 删除黑名单
	RemoveBlacklist(ctx context.Context, in *RemoveBlacklistReq, opts ...grpc.CallOption) (*RemoveBlacklistReply, error)
	// 是否好友
	IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendReply, error)
	// 是否黑名单
	IsInBlackList(ctx context.Context, in *IsInBlackListReq, opts ...grpc.CallOption) (*IsInBlackListReply, error)
	// 获取黑名单列表
	GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistReply, error)
	// 删除好友
	DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendReply, error)
	// 处理好友请求
	AddFriendResponse(ctx context.Context, in *AddFriendResponseReq, opts ...grpc.CallOption) (*AddFriendResponseReply, error)
	// 设置好友备注
	SetFriendRemark(ctx context.Context, in *SetFriendRemarkReq, opts ...grpc.CallOption) (*SetFriendRemarkReply, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) GetFriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListReply, error) {
	out := new(FriendListReply)
	err := c.cc.Invoke(ctx, Friend_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriendsInfo(ctx context.Context, in *GetFriendsInfoReq, opts ...grpc.CallOption) (*GetFriendInfoReply, error) {
	out := new(GetFriendInfoReply)
	err := c.cc.Invoke(ctx, Friend_GetFriendsInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendReply, error) {
	out := new(AddFriendReply)
	err := c.cc.Invoke(ctx, Friend_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriendApplyList(ctx context.Context, in *GetFriendApplyListReq, opts ...grpc.CallOption) (*GetFriendApplyListReply, error) {
	out := new(GetFriendApplyListReply)
	err := c.cc.Invoke(ctx, Friend_GetFriendApplyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistReply, error) {
	out := new(AddBlacklistReply)
	err := c.cc.Invoke(ctx, Friend_AddBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) RemoveBlacklist(ctx context.Context, in *RemoveBlacklistReq, opts ...grpc.CallOption) (*RemoveBlacklistReply, error) {
	out := new(RemoveBlacklistReply)
	err := c.cc.Invoke(ctx, Friend_RemoveBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendReply, error) {
	out := new(IsFriendReply)
	err := c.cc.Invoke(ctx, Friend_IsFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) IsInBlackList(ctx context.Context, in *IsInBlackListReq, opts ...grpc.CallOption) (*IsInBlackListReply, error) {
	out := new(IsInBlackListReply)
	err := c.cc.Invoke(ctx, Friend_IsInBlackList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistReply, error) {
	out := new(GetBlacklistReply)
	err := c.cc.Invoke(ctx, Friend_GetBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendReply, error) {
	out := new(DeleteFriendReply)
	err := c.cc.Invoke(ctx, Friend_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) AddFriendResponse(ctx context.Context, in *AddFriendResponseReq, opts ...grpc.CallOption) (*AddFriendResponseReply, error) {
	out := new(AddFriendResponseReply)
	err := c.cc.Invoke(ctx, Friend_AddFriendResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) SetFriendRemark(ctx context.Context, in *SetFriendRemarkReq, opts ...grpc.CallOption) (*SetFriendRemarkReply, error) {
	out := new(SetFriendRemarkReply)
	err := c.cc.Invoke(ctx, Friend_SetFriendRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	//获取好友列表仅仅返回id昵称和头像
	GetFriendList(context.Context, *FriendListReq) (*FriendListReply, error)
	// 获取好友信息
	GetFriendsInfo(context.Context, *GetFriendsInfoReq) (*GetFriendInfoReply, error)
	// 添加好友
	AddFriend(context.Context, *AddFriendReq) (*AddFriendReply, error)
	// 获取好友申请列表
	GetFriendApplyList(context.Context, *GetFriendApplyListReq) (*GetFriendApplyListReply, error)
	// 添加黑名单
	AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistReply, error)
	// 删除黑名单
	RemoveBlacklist(context.Context, *RemoveBlacklistReq) (*RemoveBlacklistReply, error)
	// 是否好友
	IsFriend(context.Context, *IsFriendReq) (*IsFriendReply, error)
	// 是否黑名单
	IsInBlackList(context.Context, *IsInBlackListReq) (*IsInBlackListReply, error)
	// 获取黑名单列表
	GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistReply, error)
	// 删除好友
	DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendReply, error)
	// 处理好友请求
	AddFriendResponse(context.Context, *AddFriendResponseReq) (*AddFriendResponseReply, error)
	// 设置好友备注
	SetFriendRemark(context.Context, *SetFriendRemarkReq) (*SetFriendRemarkReply, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) GetFriendList(context.Context, *FriendListReq) (*FriendListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedFriendServer) GetFriendsInfo(context.Context, *GetFriendsInfoReq) (*GetFriendInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendsInfo not implemented")
}
func (UnimplementedFriendServer) AddFriend(context.Context, *AddFriendReq) (*AddFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendServer) GetFriendApplyList(context.Context, *GetFriendApplyListReq) (*GetFriendApplyListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendApplyList not implemented")
}
func (UnimplementedFriendServer) AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedFriendServer) RemoveBlacklist(context.Context, *RemoveBlacklistReq) (*RemoveBlacklistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlacklist not implemented")
}
func (UnimplementedFriendServer) IsFriend(context.Context, *IsFriendReq) (*IsFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedFriendServer) IsInBlackList(context.Context, *IsInBlackListReq) (*IsInBlackListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInBlackList not implemented")
}
func (UnimplementedFriendServer) GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklist not implemented")
}
func (UnimplementedFriendServer) DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendServer) AddFriendResponse(context.Context, *AddFriendResponseReq) (*AddFriendResponseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriendResponse not implemented")
}
func (UnimplementedFriendServer) SetFriendRemark(context.Context, *SetFriendRemarkReq) (*SetFriendRemarkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriendRemark not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendList(ctx, req.(*FriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriendsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriendsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendsInfo(ctx, req.(*GetFriendsInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriendApplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendApplyListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendApplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriendApplyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendApplyList(ctx, req.(*GetFriendApplyListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddBlacklist(ctx, req.(*AddBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_RemoveBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).RemoveBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_RemoveBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).RemoveBlacklist(ctx, req.(*RemoveBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_IsFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).IsFriend(ctx, req.(*IsFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_IsInBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).IsInBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_IsInBlackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).IsInBlackList(ctx, req.(*IsInBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetBlacklist(ctx, req.(*GetBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).DeleteFriend(ctx, req.(*DeleteFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_AddFriendResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendResponseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddFriendResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddFriendResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddFriendResponse(ctx, req.(*AddFriendResponseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_SetFriendRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendRemarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).SetFriendRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_SetFriendRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).SetFriendRemark(ctx, req.(*SetFriendRemarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.friend.v2.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriendList",
			Handler:    _Friend_GetFriendList_Handler,
		},
		{
			MethodName: "GetFriendsInfo",
			Handler:    _Friend_GetFriendsInfo_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _Friend_AddFriend_Handler,
		},
		{
			MethodName: "GetFriendApplyList",
			Handler:    _Friend_GetFriendApplyList_Handler,
		},
		{
			MethodName: "AddBlacklist",
			Handler:    _Friend_AddBlacklist_Handler,
		},
		{
			MethodName: "RemoveBlacklist",
			Handler:    _Friend_RemoveBlacklist_Handler,
		},
		{
			MethodName: "IsFriend",
			Handler:    _Friend_IsFriend_Handler,
		},
		{
			MethodName: "IsInBlackList",
			Handler:    _Friend_IsInBlackList_Handler,
		},
		{
			MethodName: "GetBlacklist",
			Handler:    _Friend_GetBlacklist_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _Friend_DeleteFriend_Handler,
		},
		{
			MethodName: "AddFriendResponse",
			Handler:    _Friend_AddFriendResponse_Handler,
		},
		{
			MethodName: "SetFriendRemark",
			Handler:    _Friend_SetFriendRemark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend/v2/friend.proto",
}
