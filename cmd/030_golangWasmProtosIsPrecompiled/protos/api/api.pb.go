// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Instance struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InstanceType         string   `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Ecu                  float32  `protobuf:"fixed32,3,opt,name=ecu,proto3" json:"ecu,omitempty"`
	Memory               float32  `protobuf:"fixed32,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Network              string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	Price                string   `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_6165d98f41fd7615, []int{0}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *Instance) GetEcu() float32 {
	if m != nil {
		return m.Ecu
	}
	return 0
}

func (m *Instance) GetMemory() float32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Instance) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Instance) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type Instances struct {
	Instances            []*Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Instances) Reset()         { *m = Instances{} }
func (m *Instances) String() string { return proto.CompactTextString(m) }
func (*Instances) ProtoMessage()    {}
func (*Instances) Descriptor() ([]byte, []int) {
	return fileDescriptor_6165d98f41fd7615, []int{1}
}

func (m *Instances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instances.Unmarshal(m, b)
}
func (m *Instances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instances.Marshal(b, m, deterministic)
}
func (m *Instances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instances.Merge(m, src)
}
func (m *Instances) XXX_Size() int {
	return xxx_messageInfo_Instances.Size(m)
}
func (m *Instances) XXX_DiscardUnknown() {
	xxx_messageInfo_Instances.DiscardUnknown(m)
}

var xxx_messageInfo_Instances proto.InternalMessageInfo

func (m *Instances) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6165d98f41fd7615, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*Instance)(nil), "api.Instance")
	proto.RegisterType((*Instances)(nil), "api.Instances")
	proto.RegisterType((*SearchRequest)(nil), "api.SearchRequest")
}

func init() { proto.RegisterFile("protos/api.proto", fileDescriptor_6165d98f41fd7615) }

var fileDescriptor_6165d98f41fd7615 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x26, 0x4d, 0x9b, 0xf7, 0xed, 0x68, 0xa5, 0x0c, 0x22, 0x8b, 0xa7, 0x52, 0x11, 0x0a, 0xda,
	0x06, 0xda, 0x8b, 0x20, 0x1e, 0xf4, 0xe6, 0x35, 0x7a, 0xf2, 0x22, 0xdb, 0x65, 0x48, 0x57, 0xbb,
	0x1f, 0xdd, 0xdd, 0x50, 0xe2, 0x7f, 0xf1, 0xbf, 0x4a, 0x37, 0x09, 0xa5, 0xb7, 0xe7, 0x63, 0x76,
	0x9e, 0x9d, 0x07, 0xc6, 0xd6, 0x99, 0x60, 0x7c, 0xce, 0xad, 0x5c, 0x44, 0x88, 0x29, 0xb7, 0x72,
	0xfa, 0x9b, 0xc0, 0xff, 0x57, 0xed, 0x03, 0xd7, 0x82, 0x10, 0xa1, 0xaf, 0xb9, 0x22, 0x96, 0x4c,
	0x92, 0xd9, 0xb0, 0x88, 0x18, 0x6f, 0x60, 0x24, 0x5b, 0xff, 0x33, 0xd4, 0x96, 0x58, 0x2f, 0x9a,
	0xe7, 0x9d, 0xf8, 0x5e, 0x5b, 0xc2, 0x31, 0xa4, 0x24, 0x2a, 0x96, 0x4e, 0x92, 0x59, 0xaf, 0x38,
	0x40, 0xbc, 0x82, 0x4c, 0x91, 0x32, 0xae, 0x66, 0xfd, 0x28, 0xb6, 0x0c, 0x19, 0xfc, 0xd3, 0x14,
	0xf6, 0xc6, 0x7d, 0xb3, 0x41, 0x5c, 0xd4, 0x51, 0xbc, 0x84, 0x81, 0x75, 0x52, 0x10, 0xcb, 0xa2,
	0xde, 0x90, 0xe9, 0x03, 0x0c, 0xbb, 0xef, 0x79, 0xbc, 0x83, 0x61, 0x17, 0xeb, 0x59, 0x32, 0x49,
	0x67, 0x67, 0xcb, 0xd1, 0xe2, 0x70, 0x50, 0x37, 0x52, 0x1c, 0xfd, 0xe9, 0x2d, 0x8c, 0xde, 0x88,
	0x3b, 0xb1, 0x29, 0x68, 0x57, 0x91, 0x0f, 0x87, 0x80, 0x5d, 0x45, 0xae, 0x6e, 0xcf, 0x6b, 0xc8,
	0x72, 0x05, 0xe9, 0xb3, 0x95, 0x78, 0x0f, 0x59, 0x33, 0x8d, 0x18, 0x37, 0x9e, 0x3c, 0xbd, 0xbe,
	0x38, 0x49, 0xf1, 0x2f, 0x4f, 0x1f, 0x8f, 0xa5, 0x0c, 0x9b, 0x6a, 0xbd, 0x10, 0x46, 0xe5, 0xa5,
	0xe3, 0x1b, 0xae, 0xbe, 0x48, 0x7b, 0xa3, 0xf3, 0x35, 0xff, 0xa1, 0xed, 0xbc, 0x34, 0x5b, 0xae,
	0xcb, 0xf9, 0x9e, 0x7b, 0x35, 0x8f, 0x6d, 0x8b, 0xfc, 0xd8, 0xff, 0x3a, 0x8b, 0x78, 0xf5, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x85, 0xc3, 0x5b, 0x94, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Instances, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Instances, error) {
	out := new(Instances)
	err := c.cc.Invoke(ctx, "/api.Api/SearchInDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	Search(context.Context, *SearchRequest) (*Instances, error)
}

// UnimplementedApiServer can be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (*UnimplementedApiServer) Search(ctx context.Context, req *SearchRequest) (*Instances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchInDB not implemented")
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/SearchInDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchInDB",
			Handler:    _Api_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/api.proto",
}
