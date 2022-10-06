// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transactional

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

// TransactionalServiceClient is the client API for TransactionalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionalServiceClient interface {
	// Добавить пользователя
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	// Пополнить
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Снять
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Получить баланс
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type transactionalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionalServiceClient(cc grpc.ClientConnInterface) TransactionalServiceClient {
	return &transactionalServiceClient{cc}
}

func (c *transactionalServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/transactional.TransactionalService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionalServiceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transactional.TransactionalService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionalServiceClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transactional.TransactionalService/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionalServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/transactional.TransactionalService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionalServiceServer is the server API for TransactionalService service.
// All implementations must embed UnimplementedTransactionalServiceServer
// for forward compatibility
type TransactionalServiceServer interface {
	// Добавить пользователя
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	// Пополнить
	Deposit(context.Context, *DepositRequest) (*emptypb.Empty, error)
	// Снять
	Withdraw(context.Context, *WithdrawRequest) (*emptypb.Empty, error)
	// Получить баланс
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedTransactionalServiceServer()
}

// UnimplementedTransactionalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionalServiceServer struct {
}

func (UnimplementedTransactionalServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedTransactionalServiceServer) Deposit(context.Context, *DepositRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedTransactionalServiceServer) Withdraw(context.Context, *WithdrawRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedTransactionalServiceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedTransactionalServiceServer) mustEmbedUnimplementedTransactionalServiceServer() {}

// UnsafeTransactionalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionalServiceServer will
// result in compilation errors.
type UnsafeTransactionalServiceServer interface {
	mustEmbedUnimplementedTransactionalServiceServer()
}

func RegisterTransactionalServiceServer(s grpc.ServiceRegistrar, srv TransactionalServiceServer) {
	s.RegisterService(&TransactionalService_ServiceDesc, srv)
}

func _TransactionalService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionalServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactional.TransactionalService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionalServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionalService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionalServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactional.TransactionalService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionalServiceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionalService_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionalServiceServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactional.TransactionalService/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionalServiceServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionalService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionalServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactional.TransactionalService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionalServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionalService_ServiceDesc is the grpc.ServiceDesc for TransactionalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transactional.TransactionalService",
	HandlerType: (*TransactionalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _TransactionalService_AddUser_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _TransactionalService_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _TransactionalService_Withdraw_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _TransactionalService_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/transactional/transactional.proto",
}