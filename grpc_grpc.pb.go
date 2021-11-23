// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: grpc.proto

package test

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

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionServiceClient interface {
	//given a bid, returns an outcome among {fail, success or exception}
	Bid(ctx context.Context, in *AmountMessage, opts ...grpc.CallOption) (*AcknowledgementMessage, error)
	//Comment: if over, it returns the result, else highest bid.
	Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*OutcomeMessage, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) Bid(ctx context.Context, in *AmountMessage, opts ...grpc.CallOption) (*AcknowledgementMessage, error) {
	out := new(AcknowledgementMessage)
	err := c.cc.Invoke(ctx, "/GRPCex.AuctionService/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*OutcomeMessage, error) {
	out := new(OutcomeMessage)
	err := c.cc.Invoke(ctx, "/GRPCex.AuctionService/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility
type AuctionServiceServer interface {
	//given a bid, returns an outcome among {fail, success or exception}
	Bid(context.Context, *AmountMessage) (*AcknowledgementMessage, error)
	//Comment: if over, it returns the result, else highest bid.
	Result(context.Context, *Void) (*OutcomeMessage, error)
	mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServiceServer struct {
}

func (UnimplementedAuctionServiceServer) Bid(context.Context, *AmountMessage) (*AcknowledgementMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServiceServer) Result(context.Context, *Void) (*OutcomeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmountMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPCex.AuctionService/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Bid(ctx, req.(*AmountMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPCex.AuctionService/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Result(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GRPCex.AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bid",
			Handler:    _AuctionService_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionService_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}