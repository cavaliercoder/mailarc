// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: config.proto

package config_proto

import (
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

type S3Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *S3Store) Reset() {
	*x = S3Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Store) ProtoMessage() {}

func (x *S3Store) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Store.ProtoReflect.Descriptor instead.
func (*S3Store) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *S3Store) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FileStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileStore) Reset() {
	*x = FileStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStore) ProtoMessage() {}

func (x *FileStore) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStore.ProtoReflect.Descriptor instead.
func (*FileStore) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *FileStore) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Pop3Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr       string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	DeleteRead bool   `protobuf:"varint,4,opt,name=delete_read,json=deleteRead,proto3" json:"delete_read,omitempty"`
}

func (x *Pop3Server) Reset() {
	*x = Pop3Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pop3Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pop3Server) ProtoMessage() {}

func (x *Pop3Server) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pop3Server.ProtoReflect.Descriptor instead.
func (*Pop3Server) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *Pop3Server) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Pop3Server) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Pop3Server) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Pop3Server) GetDeleteRead() bool {
	if x != nil {
		return x.DeleteRead
	}
	return false
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Pop3Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Types that are assignable to Store:
	//	*Config_FileStore
	//	*Config_S3Store
	Store isConfig_Store `protobuf_oneof:"store"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetServer() *Pop3Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (m *Config) GetStore() isConfig_Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (x *Config) GetFileStore() *FileStore {
	if x, ok := x.GetStore().(*Config_FileStore); ok {
		return x.FileStore
	}
	return nil
}

func (x *Config) GetS3Store() *S3Store {
	if x, ok := x.GetStore().(*Config_S3Store); ok {
		return x.S3Store
	}
	return nil
}

type isConfig_Store interface {
	isConfig_Store()
}

type Config_FileStore struct {
	FileStore *FileStore `protobuf:"bytes,2,opt,name=file_store,json=fileStore,proto3,oneof"`
}

type Config_S3Store struct {
	S3Store *S3Store `protobuf:"bytes,3,opt,name=s3_store,json=s3Store,proto3,oneof"`
}

func (*Config_FileStore) isConfig_Store() {}

func (*Config_S3Store) isConfig_Store() {}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x61, 0x69, 0x6c, 0x61, 0x72, 0x63, 0x22, 0x1b, 0x0a, 0x07, 0x53, 0x33, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x1f, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x79, 0x0a, 0x0a, 0x50, 0x6f, 0x70, 0x33, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x22, 0xa2, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x61, 0x72, 0x63, 0x2e, 0x50, 0x6f, 0x70, 0x33, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x61, 0x72, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x73, 0x33, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x61, 0x72, 0x63, 0x2e, 0x53, 0x33, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x73, 0x33, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x61, 0x72, 0x63, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_proto_goTypes = []interface{}{
	(*S3Store)(nil),    // 0: mailarc.S3Store
	(*FileStore)(nil),  // 1: mailarc.FileStore
	(*Pop3Server)(nil), // 2: mailarc.Pop3Server
	(*Config)(nil),     // 3: mailarc.Config
}
var file_config_proto_depIdxs = []int32{
	2, // 0: mailarc.Config.server:type_name -> mailarc.Pop3Server
	1, // 1: mailarc.Config.file_store:type_name -> mailarc.FileStore
	0, // 2: mailarc.Config.s3_store:type_name -> mailarc.S3Store
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Store); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStore); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pop3Server); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
	file_config_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Config_FileStore)(nil),
		(*Config_S3Store)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}