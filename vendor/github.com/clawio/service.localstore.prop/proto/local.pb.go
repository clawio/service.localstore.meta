// Code generated by protoc-gen-go.
// source: local.proto
// DO NOT EDIT!

/*
Package localstore is a generated protocol buffer package.

It is generated from these files:
	local.proto

It has these top-level messages:
	Void
	PutReq
	GetReq
	RmReq
	MvReq
	Record
*/
package localstore

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

type Void struct {
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}

type PutReq struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Checksum    string `protobuf:"bytes,3,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *PutReq) Reset()         { *m = PutReq{} }
func (m *PutReq) String() string { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()    {}

type GetReq struct {
	AccessToken   string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Path          string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	ForceCreation bool   `protobuf:"varint,3,opt,name=force_creation" json:"force_creation,omitempty"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}

type RmReq struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *RmReq) Reset()         { *m = RmReq{} }
func (m *RmReq) String() string { return proto.CompactTextString(m) }
func (*RmReq) ProtoMessage()    {}

type MvReq struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	Src         string `protobuf:"bytes,2,opt,name=src" json:"src,omitempty"`
	Dst         string `protobuf:"bytes,3,opt,name=dst" json:"dst,omitempty"`
}

func (m *MvReq) Reset()         { *m = MvReq{} }
func (m *MvReq) String() string { return proto.CompactTextString(m) }
func (*MvReq) ProtoMessage()    {}

type Record struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Checksum string `protobuf:"bytes,3,opt,name=checksum" json:"checksum,omitempty"`
	Modified uint32 `protobuf:"varint,4,opt,name=modified" json:"modified,omitempty"`
	Etag     string `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Prop service

type PropClient interface {
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Record, error)
	// rpc Cp(CpReq) returns (Void) {}
	Mv(ctx context.Context, in *MvReq, opts ...grpc.CallOption) (*Void, error)
	Rm(ctx context.Context, in *RmReq, opts ...grpc.CallOption) (*Void, error)
}

type propClient struct {
	cc *grpc.ClientConn
}

func NewPropClient(cc *grpc.ClientConn) PropClient {
	return &propClient{cc}
}

func (c *propClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Prop/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := grpc.Invoke(ctx, "/localstore.Prop/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propClient) Mv(ctx context.Context, in *MvReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Prop/Mv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propClient) Rm(ctx context.Context, in *RmReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Prop/Rm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Prop service

type PropServer interface {
	Put(context.Context, *PutReq) (*Void, error)
	Get(context.Context, *GetReq) (*Record, error)
	// rpc Cp(CpReq) returns (Void) {}
	Mv(context.Context, *MvReq) (*Void, error)
	Rm(context.Context, *RmReq) (*Void, error)
}

func RegisterPropServer(s *grpc.Server, srv PropServer) {
	s.RegisterService(&_Prop_serviceDesc, srv)
}

func _Prop_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PropServer).Put(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Prop_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PropServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Prop_Mv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PropServer).Mv(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Prop_Rm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PropServer).Rm(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Prop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "localstore.Prop",
	HandlerType: (*PropServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Prop_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Prop_Get_Handler,
		},
		{
			MethodName: "Mv",
			Handler:    _Prop_Mv_Handler,
		},
		{
			MethodName: "Rm",
			Handler:    _Prop_Rm_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}