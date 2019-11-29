// Code generated by protoc-gen-go. DO NOT EDIT.
// source: torrent-store.proto

/*
Package torrent_store is a generated protocol buffer package.

It is generated from these files:
	torrent-store.proto

It has these top-level messages:
	PushReply
	PushRequest
	PullRequest
	PullReply
*/
package torrent_store

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The push response message containing info hash of the pushed torrent file
type PushReply struct {
	InfoHash string `protobuf:"bytes,1,opt,name=infoHash" json:"infoHash,omitempty"`
}

func (m *PushReply) Reset()                    { *m = PushReply{} }
func (m *PushReply) String() string            { return proto.CompactTextString(m) }
func (*PushReply) ProtoMessage()               {}
func (*PushReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PushReply) GetInfoHash() string {
	if m != nil {
		return m.InfoHash
	}
	return ""
}

// The push request message containing the torrent and expire duration is seconds
type PushRequest struct {
	Torrent []byte `protobuf:"bytes,1,opt,name=torrent,proto3" json:"torrent,omitempty"`
	Expire  int32  `protobuf:"varint,2,opt,name=expire" json:"expire,omitempty"`
}

func (m *PushRequest) Reset()                    { *m = PushRequest{} }
func (m *PushRequest) String() string            { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()               {}
func (*PushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PushRequest) GetTorrent() []byte {
	if m != nil {
		return m.Torrent
	}
	return nil
}

func (m *PushRequest) GetExpire() int32 {
	if m != nil {
		return m.Expire
	}
	return 0
}

// The pull request message containing the infoHash
type PullRequest struct {
	InfoHash string `protobuf:"bytes,1,opt,name=infoHash" json:"infoHash,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PullRequest) GetInfoHash() string {
	if m != nil {
		return m.InfoHash
	}
	return ""
}

// The pull response message containing the torrent
type PullReply struct {
	Torrent []byte `protobuf:"bytes,1,opt,name=torrent,proto3" json:"torrent,omitempty"`
}

func (m *PullReply) Reset()                    { *m = PullReply{} }
func (m *PullReply) String() string            { return proto.CompactTextString(m) }
func (*PullReply) ProtoMessage()               {}
func (*PullReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PullReply) GetTorrent() []byte {
	if m != nil {
		return m.Torrent
	}
	return nil
}

func init() {
	proto.RegisterType((*PushReply)(nil), "PushReply")
	proto.RegisterType((*PushRequest)(nil), "PushRequest")
	proto.RegisterType((*PullRequest)(nil), "PullRequest")
	proto.RegisterType((*PullReply)(nil), "PullReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TorrentStore service

type TorrentStoreClient interface {
	// Pushes torrent to the store
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushReply, error)
	// Pulls torrent from the store
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullReply, error)
}

type torrentStoreClient struct {
	cc *grpc.ClientConn
}

func NewTorrentStoreClient(cc *grpc.ClientConn) TorrentStoreClient {
	return &torrentStoreClient{cc}
}

func (c *torrentStoreClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushReply, error) {
	out := new(PushReply)
	err := grpc.Invoke(ctx, "/TorrentStore/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentStoreClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullReply, error) {
	out := new(PullReply)
	err := grpc.Invoke(ctx, "/TorrentStore/Pull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TorrentStore service

type TorrentStoreServer interface {
	// Pushes torrent to the store
	Push(context.Context, *PushRequest) (*PushReply, error)
	// Pulls torrent from the store
	Pull(context.Context, *PullRequest) (*PullReply, error)
}

func RegisterTorrentStoreServer(s *grpc.Server, srv TorrentStoreServer) {
	s.RegisterService(&_TorrentStore_serviceDesc, srv)
}

func _TorrentStore_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorrentStoreServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TorrentStore/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorrentStoreServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorrentStore_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorrentStoreServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TorrentStore/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorrentStoreServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TorrentStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TorrentStore",
	HandlerType: (*TorrentStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _TorrentStore_Push_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _TorrentStore_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "torrent-store.proto",
}

func init() { proto.RegisterFile("torrent-store.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0x2f, 0x2a,
	0x4a, 0xcd, 0x2b, 0xd1, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0x52, 0xe7, 0xe2, 0x0c, 0x28, 0x2d, 0xce, 0x08, 0x4a, 0x2d, 0xc8, 0xa9, 0x14, 0x92, 0xe2, 0xe2,
	0xc8, 0xcc, 0x4b, 0xcb, 0xf7, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0xf3, 0x95, 0xec, 0xb9, 0xb8, 0x21, 0x0a, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8,
	0xa1, 0xc6, 0x81, 0x55, 0xf2, 0x04, 0xc1, 0xb8, 0x42, 0x62, 0x5c, 0x6c, 0xa9, 0x15, 0x05, 0x99,
	0x45, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x92, 0x26, 0xc8, 0x80, 0x9c,
	0x1c, 0x98, 0x01, 0xf8, 0xec, 0x52, 0x05, 0x39, 0x0a, 0xa4, 0x14, 0xe4, 0x28, 0x9c, 0x36, 0x19,
	0x85, 0x71, 0xf1, 0x84, 0x40, 0x98, 0xc1, 0x20, 0x1f, 0x09, 0x29, 0x71, 0xb1, 0x80, 0x9c, 0x28,
	0xc4, 0xa3, 0x87, 0xe4, 0x52, 0x29, 0x2e, 0x3d, 0xb8, 0x07, 0x95, 0x18, 0x20, 0x6a, 0x72, 0x72,
	0xc0, 0x6a, 0xe0, 0x8e, 0x01, 0xab, 0x81, 0xda, 0xa7, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x1a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x26, 0x9a, 0x8b, 0x31, 0x01, 0x00, 0x00,
}
