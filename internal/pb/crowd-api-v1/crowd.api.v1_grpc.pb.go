// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: crowd-api-v1/crowd.api.v1.proto

package crowd_api_v1

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
	CrowdAPIV1_Ping_FullMethodName                  = "/crowd.api.v1.CrowdAPIV1/Ping"
	CrowdAPIV1_ResolveTasksByProject_FullMethodName = "/crowd.api.v1.CrowdAPIV1/ResolveTasksByProject"
	CrowdAPIV1_CreateProject_FullMethodName         = "/crowd.api.v1.CrowdAPIV1/CreateProject"
)

// CrowdAPIV1Client is the client API for CrowdAPIV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrowdAPIV1Client interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Tasks
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Get tasks by project id. This method reserves tasks for the user.
	ResolveTasksByProject(ctx context.Context, in *ResolveTasksByProjectRequest, opts ...grpc.CallOption) (*ResolveTasksByProjectResponse, error)
	// Projects
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
}

type crowdAPIV1Client struct {
	cc grpc.ClientConnInterface
}

func NewCrowdAPIV1Client(cc grpc.ClientConnInterface) CrowdAPIV1Client {
	return &crowdAPIV1Client{cc}
}

func (c *crowdAPIV1Client) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, CrowdAPIV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdAPIV1Client) ResolveTasksByProject(ctx context.Context, in *ResolveTasksByProjectRequest, opts ...grpc.CallOption) (*ResolveTasksByProjectResponse, error) {
	out := new(ResolveTasksByProjectResponse)
	err := c.cc.Invoke(ctx, CrowdAPIV1_ResolveTasksByProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdAPIV1Client) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, CrowdAPIV1_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrowdAPIV1Server is the server API for CrowdAPIV1 service.
// All implementations must embed UnimplementedCrowdAPIV1Server
// for forward compatibility
type CrowdAPIV1Server interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Tasks
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Get tasks by project id. This method reserves tasks for the user.
	ResolveTasksByProject(context.Context, *ResolveTasksByProjectRequest) (*ResolveTasksByProjectResponse, error)
	// Projects
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	mustEmbedUnimplementedCrowdAPIV1Server()
}

// UnimplementedCrowdAPIV1Server must be embedded to have forward compatible implementations.
type UnimplementedCrowdAPIV1Server struct {
}

func (UnimplementedCrowdAPIV1Server) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCrowdAPIV1Server) ResolveTasksByProject(context.Context, *ResolveTasksByProjectRequest) (*ResolveTasksByProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveTasksByProject not implemented")
}
func (UnimplementedCrowdAPIV1Server) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedCrowdAPIV1Server) mustEmbedUnimplementedCrowdAPIV1Server() {}

// UnsafeCrowdAPIV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrowdAPIV1Server will
// result in compilation errors.
type UnsafeCrowdAPIV1Server interface {
	mustEmbedUnimplementedCrowdAPIV1Server()
}

func RegisterCrowdAPIV1Server(s grpc.ServiceRegistrar, srv CrowdAPIV1Server) {
	s.RegisterService(&CrowdAPIV1_ServiceDesc, srv)
}

func _CrowdAPIV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowdAPIV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrowdAPIV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowdAPIV1Server).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrowdAPIV1_ResolveTasksByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveTasksByProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowdAPIV1Server).ResolveTasksByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrowdAPIV1_ResolveTasksByProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowdAPIV1Server).ResolveTasksByProject(ctx, req.(*ResolveTasksByProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrowdAPIV1_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrowdAPIV1Server).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrowdAPIV1_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrowdAPIV1Server).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrowdAPIV1_ServiceDesc is the grpc.ServiceDesc for CrowdAPIV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrowdAPIV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crowd.api.v1.CrowdAPIV1",
	HandlerType: (*CrowdAPIV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CrowdAPIV1_Ping_Handler,
		},
		{
			MethodName: "ResolveTasksByProject",
			Handler:    _CrowdAPIV1_ResolveTasksByProject_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _CrowdAPIV1_CreateProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crowd-api-v1/crowd.api.v1.proto",
}
