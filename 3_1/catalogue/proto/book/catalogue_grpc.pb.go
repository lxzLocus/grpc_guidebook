// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/book/catalogue.proto

package book

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

// CatalogueClient is the client API for Catalogue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogueClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	ListBooks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBooksResponse, error)
}

type catalogueClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogueClient(cc grpc.ClientConnInterface) CatalogueClient {
	return &catalogueClient{cc}
}

func (c *catalogueClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/book.Catalogue/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueClient) ListBooks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/book.Catalogue/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogueServer is the server API for Catalogue service.
// All implementations must embed UnimplementedCatalogueServer
// for forward compatibility
type CatalogueServer interface {
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	ListBooks(context.Context, *emptypb.Empty) (*ListBooksResponse, error)
	mustEmbedUnimplementedCatalogueServer()
}

// UnimplementedCatalogueServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogueServer struct {
}

func (UnimplementedCatalogueServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedCatalogueServer) ListBooks(context.Context, *emptypb.Empty) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedCatalogueServer) mustEmbedUnimplementedCatalogueServer() {}

// UnsafeCatalogueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogueServer will
// result in compilation errors.
type UnsafeCatalogueServer interface {
	mustEmbedUnimplementedCatalogueServer()
}

func RegisterCatalogueServer(s grpc.ServiceRegistrar, srv CatalogueServer) {
	s.RegisterService(&Catalogue_ServiceDesc, srv)
}

func _Catalogue_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.Catalogue/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalogue_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.Catalogue/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServer).ListBooks(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalogue_ServiceDesc is the grpc.ServiceDesc for Catalogue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalogue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.Catalogue",
	HandlerType: (*CatalogueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _Catalogue_GetBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _Catalogue_ListBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book/catalogue.proto",
}
