// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package goproto

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

// TxDBClient is the client API for TxDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxDBClient interface {
	GetTxn(ctx context.Context, in *TxnHash, opts ...grpc.CallOption) (*TxnResponse, error)
	SetTxn(ctx context.Context, in *SignedTxn, opts ...grpc.CallOption) (*Err, error)
	GetTxns(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*TxnsResponse, error)
	SetTxns(ctx context.Context, in *TxnsRequest, opts ...grpc.CallOption) (*Err, error)
	GetEvents(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*EventsResponse, error)
	SetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*Err, error)
	GetErrors(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*ErrorsResponse, error)
	SetError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Err, error)
}

type txDBClient struct {
	cc grpc.ClientConnInterface
}

func NewTxDBClient(cc grpc.ClientConnInterface) TxDBClient {
	return &txDBClient{cc}
}

func (c *txDBClient) GetTxn(ctx context.Context, in *TxnHash, opts ...grpc.CallOption) (*TxnResponse, error) {
	out := new(TxnResponse)
	err := c.cc.Invoke(ctx, "/TxDB/GetTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) SetTxn(ctx context.Context, in *SignedTxn, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/TxDB/SetTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) GetTxns(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*TxnsResponse, error) {
	out := new(TxnsResponse)
	err := c.cc.Invoke(ctx, "/TxDB/GetTxns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) SetTxns(ctx context.Context, in *TxnsRequest, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/TxDB/SetTxns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) GetEvents(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/TxDB/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) SetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/TxDB/SetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) GetErrors(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*ErrorsResponse, error) {
	out := new(ErrorsResponse)
	err := c.cc.Invoke(ctx, "/TxDB/GetErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txDBClient) SetError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/TxDB/SetError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxDBServer is the server API for TxDB service.
// All implementations should embed UnimplementedTxDBServer
// for forward compatibility
type TxDBServer interface {
	GetTxn(context.Context, *TxnHash) (*TxnResponse, error)
	SetTxn(context.Context, *SignedTxn) (*Err, error)
	GetTxns(context.Context, *BlockHash) (*TxnsResponse, error)
	SetTxns(context.Context, *TxnsRequest) (*Err, error)
	GetEvents(context.Context, *BlockHash) (*EventsResponse, error)
	SetEvents(context.Context, *EventsRequest) (*Err, error)
	GetErrors(context.Context, *BlockHash) (*ErrorsResponse, error)
	SetError(context.Context, *Error) (*Err, error)
}

// UnimplementedTxDBServer should be embedded to have forward compatible implementations.
type UnimplementedTxDBServer struct {
}

func (UnimplementedTxDBServer) GetTxn(context.Context, *TxnHash) (*TxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxn not implemented")
}
func (UnimplementedTxDBServer) SetTxn(context.Context, *SignedTxn) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTxn not implemented")
}
func (UnimplementedTxDBServer) GetTxns(context.Context, *BlockHash) (*TxnsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxns not implemented")
}
func (UnimplementedTxDBServer) SetTxns(context.Context, *TxnsRequest) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTxns not implemented")
}
func (UnimplementedTxDBServer) GetEvents(context.Context, *BlockHash) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedTxDBServer) SetEvents(context.Context, *EventsRequest) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEvents not implemented")
}
func (UnimplementedTxDBServer) GetErrors(context.Context, *BlockHash) (*ErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetErrors not implemented")
}
func (UnimplementedTxDBServer) SetError(context.Context, *Error) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetError not implemented")
}

// UnsafeTxDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxDBServer will
// result in compilation errors.
type UnsafeTxDBServer interface {
	mustEmbedUnimplementedTxDBServer()
}

func RegisterTxDBServer(s grpc.ServiceRegistrar, srv TxDBServer) {
	s.RegisterService(&TxDB_ServiceDesc, srv)
}

func _TxDB_GetTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).GetTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/GetTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).GetTxn(ctx, req.(*TxnHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_SetTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedTxn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).SetTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/SetTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).SetTxn(ctx, req.(*SignedTxn))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_GetTxns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).GetTxns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/GetTxns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).GetTxns(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_SetTxns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).SetTxns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/SetTxns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).SetTxns(ctx, req.(*TxnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).GetEvents(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_SetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).SetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/SetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).SetEvents(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_GetErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).GetErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/GetErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).GetErrors(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxDB_SetError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Error)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxDBServer).SetError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TxDB/SetError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxDBServer).SetError(ctx, req.(*Error))
	}
	return interceptor(ctx, in, info, handler)
}

// TxDB_ServiceDesc is the grpc.ServiceDesc for TxDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TxDB",
	HandlerType: (*TxDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxn",
			Handler:    _TxDB_GetTxn_Handler,
		},
		{
			MethodName: "SetTxn",
			Handler:    _TxDB_SetTxn_Handler,
		},
		{
			MethodName: "GetTxns",
			Handler:    _TxDB_GetTxns_Handler,
		},
		{
			MethodName: "SetTxns",
			Handler:    _TxDB_SetTxns_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _TxDB_GetEvents_Handler,
		},
		{
			MethodName: "SetEvents",
			Handler:    _TxDB_SetEvents_Handler,
		},
		{
			MethodName: "GetErrors",
			Handler:    _TxDB_GetErrors_Handler,
		},
		{
			MethodName: "SetError",
			Handler:    _TxDB_SetError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "txdb.proto",
}
