// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta2/service.proto

package v1beta2

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("akash/deployment/v1beta2/service.proto", fileDescriptor_0e37360c059968cc)
}

var fileDescriptor_0e37360c059968cc = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd3, 0x3f, 0x4f, 0xfa, 0x40,
	0x18, 0xc0, 0x71, 0xc8, 0x2f, 0x61, 0xb8, 0xe5, 0x87, 0x9d, 0xcc, 0x0d, 0x37, 0xca, 0x02, 0x77,
	0x11, 0xff, 0xac, 0x26, 0x42, 0x74, 0x22, 0x31, 0x1a, 0x07, 0xdd, 0x0e, 0xfa, 0x70, 0x34, 0x40,
	0x9f, 0xcb, 0xdd, 0x95, 0x60, 0x7c, 0x13, 0xbe, 0x09, 0xdf, 0x8b, 0x23, 0xa3, 0xa3, 0x81, 0x37,
	0x62, 0xac, 0xda, 0xab, 0x28, 0x94, 0xae, 0xed, 0xa7, 0xcf, 0xb7, 0xd7, 0xf4, 0x21, 0x07, 0x72,
	0x2c, 0xed, 0x48, 0x84, 0xa0, 0x27, 0xf8, 0x30, 0x85, 0xd8, 0x89, 0xd9, 0x61, 0x1f, 0x9c, 0x6c,
	0x0b, 0x0b, 0x66, 0x16, 0x0d, 0x80, 0x6b, 0x83, 0x0e, 0x83, 0xfd, 0xd4, 0x71, 0xef, 0xf8, 0x97,
	0xa3, 0xcd, 0x8d, 0x13, 0xfc, 0xa5, 0xa9, 0x55, 0x9f, 0x73, 0x68, 0x63, 0xa3, 0x56, 0x06, 0x13,
	0x9d, 0xc1, 0xf6, 0x73, 0x8d, 0xfc, 0xeb, 0x59, 0x15, 0xcc, 0x49, 0xbd, 0x63, 0x40, 0x3a, 0xe8,
	0x66, 0x8f, 0x04, 0x2d, 0xbe, 0xe9, 0x6d, 0x78, 0xcf, 0xaa, 0x75, 0x4e, 0x4f, 0x4a, 0xf1, 0x6b,
	0xb0, 0x1a, 0x63, 0x0b, 0xc1, 0x23, 0xd9, 0xeb, 0x82, 0x46, 0x1b, 0xb9, 0x5c, 0x9a, 0x6f, 0x9d,
	0xf5, 0xcb, 0xd3, 0xd3, 0x72, 0x3e, 0x8b, 0xcf, 0x49, 0xfd, 0x56, 0x87, 0x65, 0x8e, 0xbd, 0xce,
	0x0b, 0x8e, 0xbd, 0xce, 0xb3, 0x72, 0x42, 0xfe, 0x77, 0x26, 0x68, 0xf3, 0xe1, 0xe6, 0xf6, 0x0f,
	0xf8, 0x53, 0xd3, 0xe3, 0x32, 0x3a, 0xcb, 0x0e, 0x09, 0x49, 0x6f, 0x5d, 0x7e, 0xfc, 0x06, 0x41,
	0xa3, 0x78, 0x46, 0x0a, 0xa9, 0xd8, 0x11, 0xe6, 0x3b, 0x57, 0x32, 0xd9, 0xad, 0xe3, 0x61, 0x41,
	0xc7, 0xc3, 0x7c, 0xe7, 0xc6, 0x49, 0xe3, 0x76, 0xe9, 0x78, 0x58, 0xd0, 0xf1, 0xf0, 0xbb, 0x73,
	0x7e, 0xf7, 0xb2, 0x64, 0xd5, 0xc5, 0x92, 0x55, 0xdf, 0x96, 0xac, 0xfa, 0xb4, 0x62, 0x95, 0xc5,
	0x8a, 0x55, 0x5e, 0x57, 0xac, 0x72, 0x7f, 0xa6, 0x22, 0x37, 0x4a, 0xfa, 0x7c, 0x80, 0x53, 0x61,
	0xf5, 0x08, 0x0c, 0xc6, 0x17, 0x61, 0x2c, 0xd2, 0xf9, 0x2d, 0xa9, 0xa3, 0xd6, 0x10, 0xcd, 0x58,
	0x28, 0x14, 0x31, 0x86, 0xf0, 0xc7, 0x46, 0xf6, 0x6b, 0xe9, 0x26, 0x1e, 0xbd, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0xba, 0xe6, 0xd4, 0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateDeployment defines a method to create new deployment given proper inputs.
	CreateDeployment(ctx context.Context, in *MsgCreateDeployment, opts ...grpc.CallOption) (*MsgCreateDeploymentResponse, error)
	// DepositDeployment deposits more funds into the deployment account
	DepositDeployment(ctx context.Context, in *MsgDepositDeployment, opts ...grpc.CallOption) (*MsgDepositDeploymentResponse, error)
	// UpdateDeployment defines a method to update a deployment given proper inputs.
	UpdateDeployment(ctx context.Context, in *MsgUpdateDeployment, opts ...grpc.CallOption) (*MsgUpdateDeploymentResponse, error)
	// CloseDeployment defines a method to close a deployment given proper inputs.
	CloseDeployment(ctx context.Context, in *MsgCloseDeployment, opts ...grpc.CallOption) (*MsgCloseDeploymentResponse, error)
	// CloseGroup defines a method to close a group of a deployment given proper inputs.
	CloseGroup(ctx context.Context, in *MsgCloseGroup, opts ...grpc.CallOption) (*MsgCloseGroupResponse, error)
	// PauseGroup defines a method to close a group of a deployment given proper inputs.
	PauseGroup(ctx context.Context, in *MsgPauseGroup, opts ...grpc.CallOption) (*MsgPauseGroupResponse, error)
	// StartGroup defines a method to close a group of a deployment given proper inputs.
	StartGroup(ctx context.Context, in *MsgStartGroup, opts ...grpc.CallOption) (*MsgStartGroupResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDeployment(ctx context.Context, in *MsgCreateDeployment, opts ...grpc.CallOption) (*MsgCreateDeploymentResponse, error) {
	out := new(MsgCreateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepositDeployment(ctx context.Context, in *MsgDepositDeployment, opts ...grpc.CallOption) (*MsgDepositDeploymentResponse, error) {
	out := new(MsgDepositDeploymentResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/DepositDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDeployment(ctx context.Context, in *MsgUpdateDeployment, opts ...grpc.CallOption) (*MsgUpdateDeploymentResponse, error) {
	out := new(MsgUpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseDeployment(ctx context.Context, in *MsgCloseDeployment, opts ...grpc.CallOption) (*MsgCloseDeploymentResponse, error) {
	out := new(MsgCloseDeploymentResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/CloseDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseGroup(ctx context.Context, in *MsgCloseGroup, opts ...grpc.CallOption) (*MsgCloseGroupResponse, error) {
	out := new(MsgCloseGroupResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/CloseGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PauseGroup(ctx context.Context, in *MsgPauseGroup, opts ...grpc.CallOption) (*MsgPauseGroupResponse, error) {
	out := new(MsgPauseGroupResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/PauseGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartGroup(ctx context.Context, in *MsgStartGroup, opts ...grpc.CallOption) (*MsgStartGroupResponse, error) {
	out := new(MsgStartGroupResponse)
	err := c.cc.Invoke(ctx, "/akash.deployment.v1beta2.Msg/StartGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateDeployment defines a method to create new deployment given proper inputs.
	CreateDeployment(context.Context, *MsgCreateDeployment) (*MsgCreateDeploymentResponse, error)
	// DepositDeployment deposits more funds into the deployment account
	DepositDeployment(context.Context, *MsgDepositDeployment) (*MsgDepositDeploymentResponse, error)
	// UpdateDeployment defines a method to update a deployment given proper inputs.
	UpdateDeployment(context.Context, *MsgUpdateDeployment) (*MsgUpdateDeploymentResponse, error)
	// CloseDeployment defines a method to close a deployment given proper inputs.
	CloseDeployment(context.Context, *MsgCloseDeployment) (*MsgCloseDeploymentResponse, error)
	// CloseGroup defines a method to close a group of a deployment given proper inputs.
	CloseGroup(context.Context, *MsgCloseGroup) (*MsgCloseGroupResponse, error)
	// PauseGroup defines a method to close a group of a deployment given proper inputs.
	PauseGroup(context.Context, *MsgPauseGroup) (*MsgPauseGroupResponse, error)
	// StartGroup defines a method to close a group of a deployment given proper inputs.
	StartGroup(context.Context, *MsgStartGroup) (*MsgStartGroupResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateDeployment(ctx context.Context, req *MsgCreateDeployment) (*MsgCreateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployment not implemented")
}
func (*UnimplementedMsgServer) DepositDeployment(ctx context.Context, req *MsgDepositDeployment) (*MsgDepositDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositDeployment not implemented")
}
func (*UnimplementedMsgServer) UpdateDeployment(ctx context.Context, req *MsgUpdateDeployment) (*MsgUpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployment not implemented")
}
func (*UnimplementedMsgServer) CloseDeployment(ctx context.Context, req *MsgCloseDeployment) (*MsgCloseDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDeployment not implemented")
}
func (*UnimplementedMsgServer) CloseGroup(ctx context.Context, req *MsgCloseGroup) (*MsgCloseGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseGroup not implemented")
}
func (*UnimplementedMsgServer) PauseGroup(ctx context.Context, req *MsgPauseGroup) (*MsgPauseGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseGroup not implemented")
}
func (*UnimplementedMsgServer) StartGroup(ctx context.Context, req *MsgStartGroup) (*MsgStartGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGroup not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDeployment(ctx, req.(*MsgCreateDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepositDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/DepositDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositDeployment(ctx, req.(*MsgDepositDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/UpdateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDeployment(ctx, req.(*MsgUpdateDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/CloseDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseDeployment(ctx, req.(*MsgCloseDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/CloseGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseGroup(ctx, req.(*MsgCloseGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PauseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPauseGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PauseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/PauseGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PauseGroup(ctx, req.(*MsgPauseGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.deployment.v1beta2.Msg/StartGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartGroup(ctx, req.(*MsgStartGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.deployment.v1beta2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeployment",
			Handler:    _Msg_CreateDeployment_Handler,
		},
		{
			MethodName: "DepositDeployment",
			Handler:    _Msg_DepositDeployment_Handler,
		},
		{
			MethodName: "UpdateDeployment",
			Handler:    _Msg_UpdateDeployment_Handler,
		},
		{
			MethodName: "CloseDeployment",
			Handler:    _Msg_CloseDeployment_Handler,
		},
		{
			MethodName: "CloseGroup",
			Handler:    _Msg_CloseGroup_Handler,
		},
		{
			MethodName: "PauseGroup",
			Handler:    _Msg_PauseGroup_Handler,
		},
		{
			MethodName: "StartGroup",
			Handler:    _Msg_StartGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/deployment/v1beta2/service.proto",
}
