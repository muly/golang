// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.9.1
// source: proto/uid.proto

package uid

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qty int64 `protobuf:"varint,1,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_proto_uid_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

var File_proto_uid_proto protoreflect.FileDescriptor

var file_proto_uid_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x75, 0x69, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79, 0x32, 0x37,
	0x0a, 0x03, 0x55, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x0a, 0x2e, 0x75, 0x69, 0x64, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_uid_proto_rawDescOnce sync.Once
	file_proto_uid_proto_rawDescData = file_proto_uid_proto_rawDesc
)

func file_proto_uid_proto_rawDescGZIP() []byte {
	file_proto_uid_proto_rawDescOnce.Do(func() {
		file_proto_uid_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_uid_proto_rawDescData)
	})
	return file_proto_uid_proto_rawDescData
}

var file_proto_uid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_uid_proto_goTypes = []interface{}{
	(*Input)(nil),       // 0: uid.input
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_proto_uid_proto_depIdxs = []int32{
	0, // 0: uid.Uid.Generate:input_type -> uid.input
	1, // 1: uid.Uid.Generate:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_uid_proto_init() }
func file_proto_uid_proto_init() {
	if File_proto_uid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_uid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_uid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_uid_proto_goTypes,
		DependencyIndexes: file_proto_uid_proto_depIdxs,
		MessageInfos:      file_proto_uid_proto_msgTypes,
	}.Build()
	File_proto_uid_proto = out.File
	file_proto_uid_proto_rawDesc = nil
	file_proto_uid_proto_goTypes = nil
	file_proto_uid_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UidClient is the client API for Uid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UidClient interface {
	Generate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*empty.Empty, error)
}

type uidClient struct {
	cc grpc.ClientConnInterface
}

func NewUidClient(cc grpc.ClientConnInterface) UidClient {
	return &uidClient{cc}
}

func (c *uidClient) Generate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/uid.Uid/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UidServer is the server API for Uid service.
type UidServer interface {
	Generate(context.Context, *Input) (*empty.Empty, error)
}

// UnimplementedUidServer can be embedded to have forward compatible implementations.
type UnimplementedUidServer struct {
}

func (*UnimplementedUidServer) Generate(context.Context, *Input) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterUidServer(s *grpc.Server, srv UidServer) {
	s.RegisterService(&_Uid_serviceDesc, srv)
}

func _Uid_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UidServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uid.Uid/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UidServer).Generate(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _Uid_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uid.Uid",
	HandlerType: (*UidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Uid_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/uid.proto",
}
