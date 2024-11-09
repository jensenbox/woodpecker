// Copyright 2021 Woodpecker Authors
// Copyright 2011 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.4
// source: woodpecker.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Woodpecker_Version_FullMethodName         = "/proto.Woodpecker/Version"
	Woodpecker_Next_FullMethodName            = "/proto.Woodpecker/Next"
	Woodpecker_Init_FullMethodName            = "/proto.Woodpecker/Init"
	Woodpecker_Wait_FullMethodName            = "/proto.Woodpecker/Wait"
	Woodpecker_Done_FullMethodName            = "/proto.Woodpecker/Done"
	Woodpecker_Extend_FullMethodName          = "/proto.Woodpecker/Extend"
	Woodpecker_Update_FullMethodName          = "/proto.Woodpecker/Update"
	Woodpecker_Log_FullMethodName             = "/proto.Woodpecker/Log"
	Woodpecker_RegisterAgent_FullMethodName   = "/proto.Woodpecker/RegisterAgent"
	Woodpecker_UnregisterAgent_FullMethodName = "/proto.Woodpecker/UnregisterAgent"
	Woodpecker_ReportHealth_FullMethodName    = "/proto.Woodpecker/ReportHealth"
)

// WoodpeckerClient is the client API for Woodpecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Woodpecker Server Service
type WoodpeckerClient interface {
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error)
	Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error)
	RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error)
	UnregisterAgent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ReportHealth(ctx context.Context, in *ReportHealthRequest, opts ...grpc.CallOption) (*Empty, error)
}

type woodpeckerClient struct {
	cc grpc.ClientConnInterface
}

func NewWoodpeckerClient(cc grpc.ClientConnInterface) WoodpeckerClient {
	return &woodpeckerClient{cc}
}

func (c *woodpeckerClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, Woodpecker_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NextResponse)
	err := c.cc.Invoke(ctx, Woodpecker_Next_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Wait_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Done_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Extend(ctx context.Context, in *ExtendRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Extend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_Log_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterAgentResponse)
	err := c.cc.Invoke(ctx, Woodpecker_RegisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) UnregisterAgent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_UnregisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *woodpeckerClient) ReportHealth(ctx context.Context, in *ReportHealthRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Woodpecker_ReportHealth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WoodpeckerServer is the server API for Woodpecker service.
// All implementations must embed UnimplementedWoodpeckerServer
// for forward compatibility
//
// Woodpecker Server Service
type WoodpeckerServer interface {
	Version(context.Context, *Empty) (*VersionResponse, error)
	Next(context.Context, *NextRequest) (*NextResponse, error)
	Init(context.Context, *InitRequest) (*Empty, error)
	Wait(context.Context, *WaitRequest) (*Empty, error)
	Done(context.Context, *DoneRequest) (*Empty, error)
	Extend(context.Context, *ExtendRequest) (*Empty, error)
	Update(context.Context, *UpdateRequest) (*Empty, error)
	Log(context.Context, *LogRequest) (*Empty, error)
	RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error)
	UnregisterAgent(context.Context, *Empty) (*Empty, error)
	ReportHealth(context.Context, *ReportHealthRequest) (*Empty, error)
	mustEmbedUnimplementedWoodpeckerServer()
}

// UnimplementedWoodpeckerServer must be embedded to have forward compatible implementations.
type UnimplementedWoodpeckerServer struct {
}

func (UnimplementedWoodpeckerServer) Version(context.Context, *Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedWoodpeckerServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedWoodpeckerServer) Init(context.Context, *InitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedWoodpeckerServer) Wait(context.Context, *WaitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wait not implemented")
}
func (UnimplementedWoodpeckerServer) Done(context.Context, *DoneRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}
func (UnimplementedWoodpeckerServer) Extend(context.Context, *ExtendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extend not implemented")
}
func (UnimplementedWoodpeckerServer) Update(context.Context, *UpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWoodpeckerServer) Log(context.Context, *LogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedWoodpeckerServer) RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedWoodpeckerServer) UnregisterAgent(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterAgent not implemented")
}
func (UnimplementedWoodpeckerServer) ReportHealth(context.Context, *ReportHealthRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportHealth not implemented")
}
func (UnimplementedWoodpeckerServer) mustEmbedUnimplementedWoodpeckerServer() {}

// UnsafeWoodpeckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WoodpeckerServer will
// result in compilation errors.
type UnsafeWoodpeckerServer interface {
	mustEmbedUnimplementedWoodpeckerServer()
}

func RegisterWoodpeckerServer(s grpc.ServiceRegistrar, srv WoodpeckerServer) {
	s.RegisterService(&Woodpecker_ServiceDesc, srv)
}

func _Woodpecker_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Next_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Next(ctx, req.(*NextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Wait_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Done_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Done(ctx, req.(*DoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Extend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Extend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Extend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Extend(ctx, req.(*ExtendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_Log_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_RegisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).RegisterAgent(ctx, req.(*RegisterAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_UnregisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).UnregisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_UnregisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).UnregisterAgent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Woodpecker_ReportHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerServer).ReportHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Woodpecker_ReportHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerServer).ReportHealth(ctx, req.(*ReportHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Woodpecker_ServiceDesc is the grpc.ServiceDesc for Woodpecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Woodpecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Woodpecker",
	HandlerType: (*WoodpeckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Woodpecker_Version_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Woodpecker_Next_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Woodpecker_Init_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Woodpecker_Wait_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Woodpecker_Done_Handler,
		},
		{
			MethodName: "Extend",
			Handler:    _Woodpecker_Extend_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Woodpecker_Update_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Woodpecker_Log_Handler,
		},
		{
			MethodName: "RegisterAgent",
			Handler:    _Woodpecker_RegisterAgent_Handler,
		},
		{
			MethodName: "UnregisterAgent",
			Handler:    _Woodpecker_UnregisterAgent_Handler,
		},
		{
			MethodName: "ReportHealth",
			Handler:    _Woodpecker_ReportHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "woodpecker.proto",
}

const (
	WoodpeckerAuth_Auth_FullMethodName = "/proto.WoodpeckerAuth/Auth"
)

// WoodpeckerAuthClient is the client API for WoodpeckerAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WoodpeckerAuthClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type woodpeckerAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewWoodpeckerAuthClient(cc grpc.ClientConnInterface) WoodpeckerAuthClient {
	return &woodpeckerAuthClient{cc}
}

func (c *woodpeckerAuthClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, WoodpeckerAuth_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WoodpeckerAuthServer is the server API for WoodpeckerAuth service.
// All implementations must embed UnimplementedWoodpeckerAuthServer
// for forward compatibility
type WoodpeckerAuthServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedWoodpeckerAuthServer()
}

// UnimplementedWoodpeckerAuthServer must be embedded to have forward compatible implementations.
type UnimplementedWoodpeckerAuthServer struct {
}

func (UnimplementedWoodpeckerAuthServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedWoodpeckerAuthServer) mustEmbedUnimplementedWoodpeckerAuthServer() {}

// UnsafeWoodpeckerAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WoodpeckerAuthServer will
// result in compilation errors.
type UnsafeWoodpeckerAuthServer interface {
	mustEmbedUnimplementedWoodpeckerAuthServer()
}

func RegisterWoodpeckerAuthServer(s grpc.ServiceRegistrar, srv WoodpeckerAuthServer) {
	s.RegisterService(&WoodpeckerAuth_ServiceDesc, srv)
}

func _WoodpeckerAuth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WoodpeckerAuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WoodpeckerAuth_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WoodpeckerAuthServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WoodpeckerAuth_ServiceDesc is the grpc.ServiceDesc for WoodpeckerAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WoodpeckerAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WoodpeckerAuth",
	HandlerType: (*WoodpeckerAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _WoodpeckerAuth_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "woodpecker.proto",
}
