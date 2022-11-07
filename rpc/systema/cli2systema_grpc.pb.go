// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: systema/cli2systema.proto

package systema

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SystemACliClient is the client API for SystemACli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemACliClient interface {
	// GetUnauthenticatedCommandMessages is used by the cli to get messages from
	// Ambassador Cloud that should be raised when specified commands are ran if
	// a user is not authenticated with Ambassador Cloud. These are messages
	// that we want to get to users who aren't currently logged into Ambassador
	// Cloud. Telepresence should cache these messages, since they won't change
	// frequently.
	GetUnauthenticatedCommandMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CommandMessageResponse, error)
}

type systemACliClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemACliClient(cc grpc.ClientConnInterface) SystemACliClient {
	return &systemACliClient{cc}
}

func (c *systemACliClient) GetUnauthenticatedCommandMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CommandMessageResponse, error) {
	out := new(CommandMessageResponse)
	err := c.cc.Invoke(ctx, "/telepresence.systema.SystemACli/GetUnauthenticatedCommandMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemACliServer is the server API for SystemACli service.
// All implementations must embed UnimplementedSystemACliServer
// for forward compatibility
type SystemACliServer interface {
	// GetUnauthenticatedCommandMessages is used by the cli to get messages from
	// Ambassador Cloud that should be raised when specified commands are ran if
	// a user is not authenticated with Ambassador Cloud. These are messages
	// that we want to get to users who aren't currently logged into Ambassador
	// Cloud. Telepresence should cache these messages, since they won't change
	// frequently.
	GetUnauthenticatedCommandMessages(context.Context, *emptypb.Empty) (*CommandMessageResponse, error)
	mustEmbedUnimplementedSystemACliServer()
}

// UnimplementedSystemACliServer must be embedded to have forward compatible implementations.
type UnimplementedSystemACliServer struct {
}

func (UnimplementedSystemACliServer) GetUnauthenticatedCommandMessages(context.Context, *emptypb.Empty) (*CommandMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnauthenticatedCommandMessages not implemented")
}
func (UnimplementedSystemACliServer) mustEmbedUnimplementedSystemACliServer() {}

// UnsafeSystemACliServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemACliServer will
// result in compilation errors.
type UnsafeSystemACliServer interface {
	mustEmbedUnimplementedSystemACliServer()
}

func RegisterSystemACliServer(s grpc.ServiceRegistrar, srv SystemACliServer) {
	s.RegisterService(&SystemACli_ServiceDesc, srv)
}

func _SystemACli_GetUnauthenticatedCommandMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemACliServer).GetUnauthenticatedCommandMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.systema.SystemACli/GetUnauthenticatedCommandMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemACliServer).GetUnauthenticatedCommandMessages(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemACli_ServiceDesc is the grpc.ServiceDesc for SystemACli service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemACli_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.systema.SystemACli",
	HandlerType: (*SystemACliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnauthenticatedCommandMessages",
			Handler:    _SystemACli_GetUnauthenticatedCommandMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "systema/cli2systema.proto",
}
