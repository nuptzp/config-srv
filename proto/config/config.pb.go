// Code generated by protoc-gen-go.
// source: github.com/micro/config-srv/proto/config/config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/config-srv/proto/config/config.proto

It has these top-level messages:
	Change
	ChangeLog
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	SearchRequest
	SearchResponse
	AuditLogRequest
	AuditLogResponse
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import config1 "github.com/micro/go-platform/config/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Change struct {
	Id        string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path      string             `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Author    string             `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Comment   string             `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	Timestamp int64              `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	ChangeSet *config1.ChangeSet `protobuf:"bytes,6,opt,name=change_set" json:"change_set,omitempty"`
}

func (m *Change) Reset()                    { *m = Change{} }
func (m *Change) String() string            { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()               {}
func (*Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Change) GetChangeSet() *config1.ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type ChangeLog struct {
	Action string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Change *Change `protobuf:"bytes,2,opt,name=change" json:"change,omitempty"`
}

func (m *ChangeLog) Reset()                    { *m = ChangeLog{} }
func (m *ChangeLog) String() string            { return proto.CompactTextString(m) }
func (*ChangeLog) ProtoMessage()               {}
func (*ChangeLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChangeLog) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type UpdateRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ReadRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ReadResponse struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadResponse) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type SearchRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type SearchResponse struct {
	Configs []*Change `protobuf:"bytes,1,rep,name=configs" json:"configs,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SearchResponse) GetConfigs() []*Change {
	if m != nil {
		return m.Configs
	}
	return nil
}

type AuditLogRequest struct {
	// from unix timestamp
	From int64 `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	// to unix timestamp
	To int64 `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	// limit number items
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// the offset
	Offset int64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *AuditLogRequest) Reset()                    { *m = AuditLogRequest{} }
func (m *AuditLogRequest) String() string            { return proto.CompactTextString(m) }
func (*AuditLogRequest) ProtoMessage()               {}
func (*AuditLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type AuditLogResponse struct {
	Changes []*ChangeLog `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
}

func (m *AuditLogResponse) Reset()                    { *m = AuditLogResponse{} }
func (m *AuditLogResponse) String() string            { return proto.CompactTextString(m) }
func (*AuditLogResponse) ProtoMessage()               {}
func (*AuditLogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AuditLogResponse) GetChanges() []*ChangeLog {
	if m != nil {
		return m.Changes
	}
	return nil
}

func init() {
	proto.RegisterType((*Change)(nil), "Change")
	proto.RegisterType((*ChangeLog)(nil), "ChangeLog")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*AuditLogRequest)(nil), "AuditLogRequest")
	proto.RegisterType((*AuditLogResponse)(nil), "AuditLogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Config service

type ConfigClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	AuditLog(ctx context.Context, in *AuditLogRequest, opts ...client.CallOption) (*AuditLogResponse, error)
}

type configClient struct {
	c           client.Client
	serviceName string
}

func NewConfigClient(serviceName string, c client.Client) ConfigClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "config"
	}
	return &configClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *configClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) AuditLog(ctx context.Context, in *AuditLogRequest, opts ...client.CallOption) (*AuditLogResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.AuditLog", in)
	out := new(AuditLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Config service

type ConfigHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	AuditLog(context.Context, *AuditLogRequest, *AuditLogResponse) error
}

func RegisterConfigHandler(s server.Server, hdlr ConfigHandler) {
	s.Handle(s.NewHandler(&Config{hdlr}))
}

type Config struct {
	ConfigHandler
}

func (h *Config) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ConfigHandler.Create(ctx, in, out)
}

func (h *Config) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ConfigHandler.Update(ctx, in, out)
}

func (h *Config) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ConfigHandler.Delete(ctx, in, out)
}

func (h *Config) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.ConfigHandler.Search(ctx, in, out)
}

func (h *Config) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ConfigHandler.Read(ctx, in, out)
}

func (h *Config) AuditLog(ctx context.Context, in *AuditLogRequest, out *AuditLogResponse) error {
	return h.ConfigHandler.AuditLog(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x9b, 0x38, 0x75, 0xc9, 0xe4, 0x6f, 0x7d, 0xc1, 0x0a, 0x12, 0x42, 0x96, 0x10, 0x15,
	0xa8, 0x1b, 0x11, 0xfe, 0xdc, 0x51, 0x39, 0x21, 0x4e, 0xad, 0x38, 0x23, 0xd7, 0x59, 0xdb, 0x2b,
	0x65, 0xbd, 0x66, 0xbd, 0xe9, 0x47, 0xe1, 0xf3, 0x32, 0x3b, 0xeb, 0x55, 0x6c, 0x43, 0x45, 0x4f,
	0xc9, 0x8e, 0x7f, 0x6f, 0xde, 0x7a, 0xde, 0x18, 0x3e, 0x15, 0xc2, 0x94, 0xc7, 0x7b, 0x96, 0x29,
	0xb9, 0x95, 0x22, 0xd3, 0x6a, 0x9b, 0xa9, 0x2a, 0x17, 0xc5, 0x75, 0xa3, 0x1f, 0xb6, 0xb5, 0x56,
	0xc6, 0x17, 0xda, 0x1f, 0x46, 0xb5, 0xcd, 0xe7, 0xbf, 0x64, 0x85, 0xba, 0xae, 0x0f, 0xa9, 0xc9,
	0x95, 0x96, 0x5e, 0xd1, 0x95, 0x3b, 0x5d, 0xf2, 0x00, 0xe1, 0x4d, 0x99, 0x56, 0x05, 0x8f, 0x00,
	0xc6, 0x62, 0x1f, 0x8f, 0x5e, 0x8d, 0xae, 0xa6, 0xd1, 0x1c, 0x26, 0x75, 0x6a, 0xca, 0x78, 0x4c,
	0xa7, 0x25, 0x84, 0xe9, 0xd1, 0x94, 0x4a, 0xc7, 0x01, 0x9d, 0x57, 0x70, 0x81, 0x36, 0x92, 0x57,
	0x26, 0x9e, 0x50, 0xe1, 0x12, 0xa6, 0x46, 0x48, 0xde, 0x98, 0x54, 0xd6, 0xf1, 0x39, 0x96, 0x82,
	0xe8, 0x25, 0x40, 0x46, 0x7d, 0x7f, 0x36, 0xdc, 0xc4, 0x21, 0xd6, 0x66, 0x3b, 0x60, 0xce, 0xea,
	0x8e, 0x9b, 0xe4, 0x23, 0x4c, 0xdd, 0xe1, 0xbb, 0x2a, 0xc8, 0x20, 0x33, 0x42, 0x55, 0xad, 0xfd,
	0x73, 0x08, 0x9d, 0x98, 0x2e, 0x30, 0xdb, 0x5d, 0xb4, 0xc2, 0xe4, 0x0a, 0x16, 0x37, 0x9a, 0xa7,
	0x86, 0xdf, 0xf2, 0x5f, 0x47, 0xf4, 0xeb, 0x90, 0xa3, 0x3e, 0xb9, 0x86, 0xa5, 0x27, 0x9b, 0x5a,
	0x55, 0x0d, 0x69, 0x7f, 0xd4, 0xfb, 0x27, 0x6a, 0x3d, 0x79, 0xd2, 0x7e, 0xe5, 0x07, 0xfe, 0x34,
	0xad, 0x27, 0x5b, 0xed, 0x1b, 0x98, 0xdd, 0xf2, 0x74, 0xef, 0x95, 0x8f, 0x8e, 0x19, 0xc1, 0xb9,
	0x03, 0x9d, 0xf0, 0x71, 0x8f, 0x6f, 0xb0, 0xb8, 0xe3, 0xa9, 0xce, 0xca, 0x7f, 0xf5, 0x3c, 0x85,
	0xe5, 0xc2, 0x5b, 0xc0, 0xf9, 0x41, 0x48, 0x61, 0x28, 0xbb, 0xc0, 0x3e, 0x56, 0x79, 0x6e, 0x33,
	0xb1, 0xd1, 0x05, 0xc9, 0x5b, 0x58, 0xfa, 0x5e, 0xad, 0x6d, 0x6c, 0xd3, 0xb5, 0x1b, 0xd2, 0x60,
	0xc7, 0xa0, 0xef, 0xbb, 0xfa, 0x72, 0xdc, 0x0b, 0x83, 0x91, 0x79, 0x67, 0x7c, 0x83, 0x5c, 0x2b,
	0x49, 0xde, 0x81, 0xbd, 0x87, 0x51, 0xe4, 0x1b, 0xfc, 0xcf, 0x77, 0x0b, 0xeb, 0x53, 0xaf, 0xd6,
	0xf9, 0x05, 0x3a, 0x93, 0x93, 0x77, 0xf6, 0x0b, 0x83, 0xd0, 0xee, 0xf7, 0x18, 0x37, 0x95, 0xee,
	0x15, 0xbd, 0xc3, 0x7f, 0x94, 0x6d, 0xb4, 0x64, 0xbd, 0x75, 0xd8, 0xac, 0xd8, 0x20, 0xf4, 0x33,
	0x0b, 0xbb, 0x30, 0x11, 0xee, 0xe5, 0x8f, 0xf0, 0x20, 0x65, 0x82, 0x5d, 0x7a, 0x08, 0xf7, 0x02,
	0x47, 0x78, 0x10, 0x2b, 0xc1, 0x6e, 0x74, 0x08, 0xf7, 0xf2, 0x40, 0xb8, 0x3f, 0x53, 0x84, 0x5f,
	0xc3, 0xc4, 0x86, 0x1b, 0xcd, 0x59, 0x67, 0x19, 0x36, 0x0b, 0xd6, 0x4d, 0x1c, 0xb1, 0xf7, 0xf0,
	0xcc, 0x8f, 0x25, 0x5a, 0xb3, 0xc1, 0xb4, 0x37, 0x97, 0x6c, 0x38, 0xb3, 0xe4, 0xec, 0x3e, 0xa4,
	0x0f, 0xf9, 0xc3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x68, 0x4d, 0xfb, 0x39, 0x04, 0x00,
	0x00,
}
