// Code generated by protoc-gen-go.
// source: local.proto
// DO NOT EDIT!

/*
Package localstore is a generated protocol buffer package.

It is generated from these files:
	local.proto

It has these top-level messages:
	Void
	RmReq
	MvReq
	HomeReq
	CpReq
	MkdirReq
	StatReq
	Metadata
	Identity
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

type RmReq struct {
	Idt  *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	Path string    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *RmReq) Reset()         { *m = RmReq{} }
func (m *RmReq) String() string { return proto.CompactTextString(m) }
func (*RmReq) ProtoMessage()    {}

func (m *RmReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type MvReq struct {
	Idt *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	Src string    `protobuf:"bytes,2,opt,name=src" json:"src,omitempty"`
	Dst string    `protobuf:"bytes,3,opt,name=dst" json:"dst,omitempty"`
}

func (m *MvReq) Reset()         { *m = MvReq{} }
func (m *MvReq) String() string { return proto.CompactTextString(m) }
func (*MvReq) ProtoMessage()    {}

func (m *MvReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type HomeReq struct {
	Idt *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
}

func (m *HomeReq) Reset()         { *m = HomeReq{} }
func (m *HomeReq) String() string { return proto.CompactTextString(m) }
func (*HomeReq) ProtoMessage()    {}

func (m *HomeReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type CpReq struct {
	Idt *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	Src string    `protobuf:"bytes,2,opt,name=src" json:"src,omitempty"`
	Dst string    `protobuf:"bytes,3,opt,name=dst" json:"dst,omitempty"`
}

func (m *CpReq) Reset()         { *m = CpReq{} }
func (m *CpReq) String() string { return proto.CompactTextString(m) }
func (*CpReq) ProtoMessage()    {}

func (m *CpReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type MkdirReq struct {
	Idt  *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	Path string    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *MkdirReq) Reset()         { *m = MkdirReq{} }
func (m *MkdirReq) String() string { return proto.CompactTextString(m) }
func (*MkdirReq) ProtoMessage()    {}

func (m *MkdirReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type StatReq struct {
	Idt      *Identity `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	Path     string    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Children bool      `protobuf:"varint,3,opt,name=children" json:"children,omitempty"`
}

func (m *StatReq) Reset()         { *m = StatReq{} }
func (m *StatReq) String() string { return proto.CompactTextString(m) }
func (*StatReq) ProtoMessage()    {}

func (m *StatReq) GetIdt() *Identity {
	if m != nil {
		return m.Idt
	}
	return nil
}

type Metadata struct {
	Id          string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path        string      `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Size        uint32      `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	IsContainer bool        `protobuf:"varint,4,opt,name=is_container" json:"is_container,omitempty"`
	MimeType    string      `protobuf:"bytes,5,opt,name=mime_type" json:"mime_type,omitempty"`
	Checksum    string      `protobuf:"bytes,6,opt,name=checksum" json:"checksum,omitempty"`
	Modified    uint32      `protobuf:"varint,7,opt,name=modified" json:"modified,omitempty"`
	Etag        string      `protobuf:"bytes,8,opt,name=etag" json:"etag,omitempty"`
	Permissions uint32      `protobuf:"varint,9,opt,name=permissions" json:"permissions,omitempty"`
	Children    []*Metadata `protobuf:"bytes,10,rep,name=children" json:"children,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}

func (m *Metadata) GetChildren() []*Metadata {
	if m != nil {
		return m.Children
	}
	return nil
}

type Identity struct {
	Pid         string `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Idp         string `protobuf:"bytes,2,opt,name=idp" json:"idp,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name" json:"display_name,omitempty"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Local service

type LocalClient interface {
	Home(ctx context.Context, in *HomeReq, opts ...grpc.CallOption) (*Void, error)
	Mkdir(ctx context.Context, in *MkdirReq, opts ...grpc.CallOption) (*Void, error)
	Stat(ctx context.Context, in *StatReq, opts ...grpc.CallOption) (*Metadata, error)
	Cp(ctx context.Context, in *CpReq, opts ...grpc.CallOption) (*Void, error)
	Mv(ctx context.Context, in *MvReq, opts ...grpc.CallOption) (*Void, error)
	Rm(ctx context.Context, in *RmReq, opts ...grpc.CallOption) (*Void, error)
}

type localClient struct {
	cc *grpc.ClientConn
}

func NewLocalClient(cc *grpc.ClientConn) LocalClient {
	return &localClient{cc}
}

func (c *localClient) Home(ctx context.Context, in *HomeReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Local/Home", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Mkdir(ctx context.Context, in *MkdirReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Local/Mkdir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Stat(ctx context.Context, in *StatReq, opts ...grpc.CallOption) (*Metadata, error) {
	out := new(Metadata)
	err := grpc.Invoke(ctx, "/localstore.Local/Stat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Cp(ctx context.Context, in *CpReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Local/Cp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Mv(ctx context.Context, in *MvReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Local/Mv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localClient) Rm(ctx context.Context, in *RmReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/localstore.Local/Rm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Local service

type LocalServer interface {
	Home(context.Context, *HomeReq) (*Void, error)
	Mkdir(context.Context, *MkdirReq) (*Void, error)
	Stat(context.Context, *StatReq) (*Metadata, error)
	Cp(context.Context, *CpReq) (*Void, error)
	Mv(context.Context, *MvReq) (*Void, error)
	Rm(context.Context, *RmReq) (*Void, error)
}

func RegisterLocalServer(s *grpc.Server, srv LocalServer) {
	s.RegisterService(&_Local_serviceDesc, srv)
}

func _Local_Home_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Home(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Local_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MkdirReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Mkdir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Local_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Stat(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Local_Cp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Cp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Local_Mv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Mv(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Local_Rm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LocalServer).Rm(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Local_serviceDesc = grpc.ServiceDesc{
	ServiceName: "localstore.Local",
	HandlerType: (*LocalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Home",
			Handler:    _Local_Home_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _Local_Mkdir_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _Local_Stat_Handler,
		},
		{
			MethodName: "Cp",
			Handler:    _Local_Cp_Handler,
		},
		{
			MethodName: "Mv",
			Handler:    _Local_Mv_Handler,
		},
		{
			MethodName: "Rm",
			Handler:    _Local_Rm_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
