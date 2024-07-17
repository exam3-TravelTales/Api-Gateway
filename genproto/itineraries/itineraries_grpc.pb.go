// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: itineraries.proto

package itineraries

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

// ItinerariesClient is the client API for Itineraries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItinerariesClient interface {
	Itineraries(ctx context.Context, in *ItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error)
	UpdateItineraries(ctx context.Context, in *UpdateItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error)
	DeleteItineraries(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error)
	GetItineraries(ctx context.Context, in *GetItinerariesReq, opts ...grpc.CallOption) (*GetItinerariesRes, error)
	GetItinerariesById(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetItinerariesByIdRes, error)
	CommentItineraries(ctx context.Context, in *CommentItinerariesReq, opts ...grpc.CallOption) (*CommentItinerariesRes, error)
}

type itinerariesClient struct {
	cc grpc.ClientConnInterface
}

func NewItinerariesClient(cc grpc.ClientConnInterface) ItinerariesClient {
	return &itinerariesClient{cc}
}

func (c *itinerariesClient) Itineraries(ctx context.Context, in *ItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error) {
	out := new(ItinerariesRes)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/Itineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) UpdateItineraries(ctx context.Context, in *UpdateItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error) {
	out := new(ItinerariesRes)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/UpdateItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) DeleteItineraries(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/DeleteItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetItineraries(ctx context.Context, in *GetItinerariesReq, opts ...grpc.CallOption) (*GetItinerariesRes, error) {
	out := new(GetItinerariesRes)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/GetItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) GetItinerariesById(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetItinerariesByIdRes, error) {
	out := new(GetItinerariesByIdRes)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/GetItinerariesById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesClient) CommentItineraries(ctx context.Context, in *CommentItinerariesReq, opts ...grpc.CallOption) (*CommentItinerariesRes, error) {
	out := new(CommentItinerariesRes)
	err := c.cc.Invoke(ctx, "/itineraries.Itineraries/CommentItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItinerariesServer is the server API for Itineraries service.
// All implementations must embed UnimplementedItinerariesServer
// for forward compatibility
type ItinerariesServer interface {
	Itineraries(context.Context, *ItinerariesReq) (*ItinerariesRes, error)
	UpdateItineraries(context.Context, *UpdateItinerariesReq) (*ItinerariesRes, error)
	DeleteItineraries(context.Context, *StoryId) (*Void, error)
	GetItineraries(context.Context, *GetItinerariesReq) (*GetItinerariesRes, error)
	GetItinerariesById(context.Context, *StoryId) (*GetItinerariesByIdRes, error)
	CommentItineraries(context.Context, *CommentItinerariesReq) (*CommentItinerariesRes, error)
	mustEmbedUnimplementedItinerariesServer()
}

// UnimplementedItinerariesServer must be embedded to have forward compatible implementations.
type UnimplementedItinerariesServer struct {
}

func (UnimplementedItinerariesServer) Itineraries(context.Context, *ItinerariesReq) (*ItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Itineraries not implemented")
}
func (UnimplementedItinerariesServer) UpdateItineraries(context.Context, *UpdateItinerariesReq) (*ItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItineraries not implemented")
}
func (UnimplementedItinerariesServer) DeleteItineraries(context.Context, *StoryId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItineraries not implemented")
}
func (UnimplementedItinerariesServer) GetItineraries(context.Context, *GetItinerariesReq) (*GetItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItineraries not implemented")
}
func (UnimplementedItinerariesServer) GetItinerariesById(context.Context, *StoryId) (*GetItinerariesByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItinerariesById not implemented")
}
func (UnimplementedItinerariesServer) CommentItineraries(context.Context, *CommentItinerariesReq) (*CommentItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentItineraries not implemented")
}
func (UnimplementedItinerariesServer) mustEmbedUnimplementedItinerariesServer() {}

// UnsafeItinerariesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItinerariesServer will
// result in compilation errors.
type UnsafeItinerariesServer interface {
	mustEmbedUnimplementedItinerariesServer()
}

func RegisterItinerariesServer(s grpc.ServiceRegistrar, srv ItinerariesServer) {
	s.RegisterService(&Itineraries_ServiceDesc, srv)
}

func _Itineraries_Itineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).Itineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/Itineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).Itineraries(ctx, req.(*ItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_UpdateItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).UpdateItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/UpdateItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).UpdateItineraries(ctx, req.(*UpdateItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_DeleteItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).DeleteItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/DeleteItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).DeleteItineraries(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/GetItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetItineraries(ctx, req.(*GetItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_GetItinerariesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).GetItinerariesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/GetItinerariesById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).GetItinerariesById(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Itineraries_CommentItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServer).CommentItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itineraries.Itineraries/CommentItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServer).CommentItineraries(ctx, req.(*CommentItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Itineraries_ServiceDesc is the grpc.ServiceDesc for Itineraries service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Itineraries_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itineraries.Itineraries",
	HandlerType: (*ItinerariesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Itineraries",
			Handler:    _Itineraries_Itineraries_Handler,
		},
		{
			MethodName: "UpdateItineraries",
			Handler:    _Itineraries_UpdateItineraries_Handler,
		},
		{
			MethodName: "DeleteItineraries",
			Handler:    _Itineraries_DeleteItineraries_Handler,
		},
		{
			MethodName: "GetItineraries",
			Handler:    _Itineraries_GetItineraries_Handler,
		},
		{
			MethodName: "GetItinerariesById",
			Handler:    _Itineraries_GetItinerariesById_Handler,
		},
		{
			MethodName: "CommentItineraries",
			Handler:    _Itineraries_CommentItineraries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "itineraries.proto",
}
