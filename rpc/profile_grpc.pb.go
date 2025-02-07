// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: profile.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProfilesService_GetProfileByUserId_FullMethodName = "/microblog.rpc.v1.ProfilesService/GetProfileByUserId"
	ProfilesService_CreateProfile_FullMethodName      = "/microblog.rpc.v1.ProfilesService/CreateProfile"
)

// ProfilesServiceClient is the client API for ProfilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesServiceClient interface {
	GetProfileByUserId(ctx context.Context, in *GetProfileByUserIdRequest, opts ...grpc.CallOption) (*GetProfileByUserIdResponse, error)
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error)
}

type profilesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesServiceClient(cc grpc.ClientConnInterface) ProfilesServiceClient {
	return &profilesServiceClient{cc}
}

func (c *profilesServiceClient) GetProfileByUserId(ctx context.Context, in *GetProfileByUserIdRequest, opts ...grpc.CallOption) (*GetProfileByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileByUserIdResponse)
	err := c.cc.Invoke(ctx, ProfilesService_GetProfileByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProfileResponse)
	err := c.cc.Invoke(ctx, ProfilesService_CreateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServiceServer is the server API for ProfilesService service.
// All implementations must embed UnimplementedProfilesServiceServer
// for forward compatibility.
type ProfilesServiceServer interface {
	GetProfileByUserId(context.Context, *GetProfileByUserIdRequest) (*GetProfileByUserIdResponse, error)
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
	mustEmbedUnimplementedProfilesServiceServer()
}

// UnimplementedProfilesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfilesServiceServer struct{}

func (UnimplementedProfilesServiceServer) GetProfileByUserId(context.Context, *GetProfileByUserIdRequest) (*GetProfileByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileByUserId not implemented")
}
func (UnimplementedProfilesServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfilesServiceServer) mustEmbedUnimplementedProfilesServiceServer() {}
func (UnimplementedProfilesServiceServer) testEmbeddedByValue()                         {}

// UnsafeProfilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServiceServer will
// result in compilation errors.
type UnsafeProfilesServiceServer interface {
	mustEmbedUnimplementedProfilesServiceServer()
}

func RegisterProfilesServiceServer(s grpc.ServiceRegistrar, srv ProfilesServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfilesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfilesService_ServiceDesc, srv)
}

func _ProfilesService_GetProfileByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).GetProfileByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesService_GetProfileByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).GetProfileByUserId(ctx, req.(*GetProfileByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesService_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilesService_ServiceDesc is the grpc.ServiceDesc for ProfilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microblog.rpc.v1.ProfilesService",
	HandlerType: (*ProfilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileByUserId",
			Handler:    _ProfilesService_GetProfileByUserId_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _ProfilesService_CreateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
