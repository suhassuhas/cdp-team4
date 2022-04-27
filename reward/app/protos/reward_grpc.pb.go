// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: reward.proto

package protos

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

// RewardClient is the client API for Reward service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RewardClient interface {
	GetRewardPoints(ctx context.Context, in *GetRewardPointsRequest, opts ...grpc.CallOption) (*GetRewardPointsResponse, error)
}

type rewardClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardClient(cc grpc.ClientConnInterface) RewardClient {
	return &rewardClient{cc}
}

func (c *rewardClient) GetRewardPoints(ctx context.Context, in *GetRewardPointsRequest, opts ...grpc.CallOption) (*GetRewardPointsResponse, error) {
	out := new(GetRewardPointsResponse)
	err := c.cc.Invoke(ctx, "/Reward/GetRewardPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardServer is the server API for Reward service.
// All implementations must embed UnimplementedRewardServer
// for forward compatibility
type RewardServer interface {
	GetRewardPoints(context.Context, *GetRewardPointsRequest) (*GetRewardPointsResponse, error)
	mustEmbedUnimplementedRewardServer()
}

// UnimplementedRewardServer must be embedded to have forward compatible implementations.
type UnimplementedRewardServer struct {
}

func (UnimplementedRewardServer) GetRewardPoints(context.Context, *GetRewardPointsRequest) (*GetRewardPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewardPoints not implemented")
}
func (UnimplementedRewardServer) mustEmbedUnimplementedRewardServer() {}

// UnsafeRewardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RewardServer will
// result in compilation errors.
type UnsafeRewardServer interface {
	mustEmbedUnimplementedRewardServer()
}

func RegisterRewardServer(s grpc.ServiceRegistrar, srv RewardServer) {
	s.RegisterService(&Reward_ServiceDesc, srv)
}

func _Reward_GetRewardPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServer).GetRewardPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Reward/GetRewardPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServer).GetRewardPoints(ctx, req.(*GetRewardPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reward_ServiceDesc is the grpc.ServiceDesc for Reward service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reward_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Reward",
	HandlerType: (*RewardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRewardPoints",
			Handler:    _Reward_GetRewardPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reward.proto",
}
