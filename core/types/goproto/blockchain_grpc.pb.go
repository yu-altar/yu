// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: blockchain.proto

package goproto

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

// BlockChainClient is the client API for BlockChain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockChainClient interface {
	GetGenesis(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error)
	SetGenesis(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error)
	AppendBlock(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error)
	GetBlock(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*BlockResponse, error)
	ExistsBlock(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*Bool, error)
	UpdateBlock(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error)
	Children(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*BlocksResponse, error)
	Finalize(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*Err, error)
	GetFinalizedBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error)
	GetEndBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error)
	GetRangeBlocks(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*BlocksResponse, error)
}

type blockChainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockChainClient(cc grpc.ClientConnInterface) BlockChainClient {
	return &blockChainClient{cc}
}

func (c *blockChainClient) GetGenesis(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/GetGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) SetGenesis(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/BlockChain/SetGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) AppendBlock(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/BlockChain/AppendBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) GetBlock(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) ExistsBlock(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/BlockChain/ExistsBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) UpdateBlock(ctx context.Context, in *CompactBlock, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/BlockChain/UpdateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) Children(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*BlocksResponse, error) {
	out := new(BlocksResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/Children", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) Finalize(ctx context.Context, in *BlockHash, opts ...grpc.CallOption) (*Err, error) {
	out := new(Err)
	err := c.cc.Invoke(ctx, "/BlockChain/Finalize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) GetFinalizedBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/GetFinalizedBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) GetEndBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/GetEndBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainClient) GetRangeBlocks(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*BlocksResponse, error) {
	out := new(BlocksResponse)
	err := c.cc.Invoke(ctx, "/BlockChain/GetRangeBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockChainServer is the server API for BlockChain service.
// All implementations should embed UnimplementedBlockChainServer
// for forward compatibility
type BlockChainServer interface {
	GetGenesis(context.Context, *emptypb.Empty) (*BlockResponse, error)
	SetGenesis(context.Context, *CompactBlock) (*Err, error)
	AppendBlock(context.Context, *CompactBlock) (*Err, error)
	GetBlock(context.Context, *BlockHash) (*BlockResponse, error)
	ExistsBlock(context.Context, *BlockHash) (*Bool, error)
	UpdateBlock(context.Context, *CompactBlock) (*Err, error)
	Children(context.Context, *BlockHash) (*BlocksResponse, error)
	Finalize(context.Context, *BlockHash) (*Err, error)
	GetFinalizedBlock(context.Context, *emptypb.Empty) (*BlockResponse, error)
	GetEndBlock(context.Context, *emptypb.Empty) (*BlockResponse, error)
	GetRangeBlocks(context.Context, *RangeRequest) (*BlocksResponse, error)
}

// UnimplementedBlockChainServer should be embedded to have forward compatible implementations.
type UnimplementedBlockChainServer struct {
}

func (UnimplementedBlockChainServer) GetGenesis(context.Context, *emptypb.Empty) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenesis not implemented")
}
func (UnimplementedBlockChainServer) SetGenesis(context.Context, *CompactBlock) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGenesis not implemented")
}
func (UnimplementedBlockChainServer) AppendBlock(context.Context, *CompactBlock) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendBlock not implemented")
}
func (UnimplementedBlockChainServer) GetBlock(context.Context, *BlockHash) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockChainServer) ExistsBlock(context.Context, *BlockHash) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsBlock not implemented")
}
func (UnimplementedBlockChainServer) UpdateBlock(context.Context, *CompactBlock) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlock not implemented")
}
func (UnimplementedBlockChainServer) Children(context.Context, *BlockHash) (*BlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Children not implemented")
}
func (UnimplementedBlockChainServer) Finalize(context.Context, *BlockHash) (*Err, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finalize not implemented")
}
func (UnimplementedBlockChainServer) GetFinalizedBlock(context.Context, *emptypb.Empty) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinalizedBlock not implemented")
}
func (UnimplementedBlockChainServer) GetEndBlock(context.Context, *emptypb.Empty) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndBlock not implemented")
}
func (UnimplementedBlockChainServer) GetRangeBlocks(context.Context, *RangeRequest) (*BlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRangeBlocks not implemented")
}

// UnsafeBlockChainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockChainServer will
// receipt in compilation errors.
type UnsafeBlockChainServer interface {
	mustEmbedUnimplementedBlockChainServer()
}

func RegisterBlockChainServer(s grpc.ServiceRegistrar, srv BlockChainServer) {
	s.RegisterService(&BlockChain_ServiceDesc, srv)
}

func _BlockChain_GetGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).GetGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/GetGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).GetGenesis(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_SetGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompactBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).SetGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/SetGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).SetGenesis(ctx, req.(*CompactBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_AppendBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompactBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).AppendBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/AppendBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).AppendBlock(ctx, req.(*CompactBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).GetBlock(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_ExistsBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).ExistsBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/ExistsBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).ExistsBlock(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_UpdateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompactBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).UpdateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/UpdateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).UpdateBlock(ctx, req.(*CompactBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_Children_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).Children(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/Children",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).Children(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_Finalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).Finalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/Finalize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).Finalize(ctx, req.(*BlockHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_GetFinalizedBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).GetFinalizedBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/GetFinalizedBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).GetFinalizedBlock(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_GetEndBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).GetEndBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/GetEndBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).GetEndBlock(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChain_GetRangeBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainServer).GetRangeBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChain/GetRangeBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainServer).GetRangeBlocks(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockChain_ServiceDesc is the grpc.ServiceDesc for BlockChain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockChain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlockChain",
	HandlerType: (*BlockChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenesis",
			Handler:    _BlockChain_GetGenesis_Handler,
		},
		{
			MethodName: "SetGenesis",
			Handler:    _BlockChain_SetGenesis_Handler,
		},
		{
			MethodName: "AppendBlock",
			Handler:    _BlockChain_AppendBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BlockChain_GetBlock_Handler,
		},
		{
			MethodName: "ExistsBlock",
			Handler:    _BlockChain_ExistsBlock_Handler,
		},
		{
			MethodName: "UpdateBlock",
			Handler:    _BlockChain_UpdateBlock_Handler,
		},
		{
			MethodName: "Children",
			Handler:    _BlockChain_Children_Handler,
		},
		{
			MethodName: "Finalize",
			Handler:    _BlockChain_Finalize_Handler,
		},
		{
			MethodName: "GetFinalizedBlock",
			Handler:    _BlockChain_GetFinalizedBlock_Handler,
		},
		{
			MethodName: "GetEndBlock",
			Handler:    _BlockChain_GetEndBlock_Handler,
		},
		{
			MethodName: "GetRangeBlocks",
			Handler:    _BlockChain_GetRangeBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain.proto",
}
