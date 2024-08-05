// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: tetris.proto

package pb

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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.


// 俄罗斯形状
type TetrisShape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Edges []int32 `protobuf:"varint,1,rep,packed,name=edges,proto3" json:"edges,omitempty"`
}

func (x *TetrisShape) Reset() {
	*x = TetrisShape{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetris_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TetrisShape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TetrisShape) ProtoMessage() {}

func (x *TetrisShape) ProtoReflect() protoreflect.Message {
	mi := &file_tetris_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TetrisShape.ProtoReflect.Descriptor instead.
func (*TetrisShape) Descriptor() ([]byte, []int) {
	return file_tetris_proto_rawDescGZIP(), []int{0}
}

func (x *TetrisShape) GetEdges() []int32 {
	if x != nil {
		return x.Edges
	}
	return nil
}

// 俄罗斯方块
type TetrisBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left   int32          `protobuf:"varint,1,opt,name=left,proto3" json:"left,omitempty"`     // 向左移动
	Right  int32          `protobuf:"varint,2,opt,name=right,proto3" json:"right,omitempty"`   // 向右移动
	Down   int32          `protobuf:"varint,3,opt,name=down,proto3" json:"down,omitempty"`     // 下降
	Up     int32          `protobuf:"varint,4,opt,name=up,proto3" json:"up,omitempty"`         // 变换角度
	Bottom int32          `protobuf:"varint,5,opt,name=bottom,proto3" json:"bottom,omitempty"` // 快速到底部
	Values []*TetrisShape `protobuf:"bytes,6,rep,name=values,proto3" json:"values,omitempty"`  // 俄罗斯形状
}

func (x *TetrisBlock) Reset() {
	*x = TetrisBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetris_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TetrisBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TetrisBlock) ProtoMessage() {}

func (x *TetrisBlock) ProtoReflect() protoreflect.Message {
	mi := &file_tetris_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TetrisBlock.ProtoReflect.Descriptor instead.
func (*TetrisBlock) Descriptor() ([]byte, []int) {
	return file_tetris_proto_rawDescGZIP(), []int{1}
}

func (x *TetrisBlock) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *TetrisBlock) GetRight() int32 {
	if x != nil {
		return x.Right
	}
	return 0
}

func (x *TetrisBlock) GetDown() int32 {
	if x != nil {
		return x.Down
	}
	return 0
}

func (x *TetrisBlock) GetUp() int32 {
	if x != nil {
		return x.Up
	}
	return 0
}

func (x *TetrisBlock) GetBottom() int32 {
	if x != nil {
		return x.Bottom
	}
	return 0
}

func (x *TetrisBlock) GetValues() []*TetrisShape {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_tetris_proto protoreflect.FileDescriptor

var file_tetris_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x74, 0x72, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x23, 0x0a, 0x0b, 0x54, 0x65, 0x74, 0x72, 0x69, 0x73, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x74, 0x72,
	0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x12, 0x27, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x74, 0x72, 0x69, 0x73, 0x53, 0x68, 0x61, 0x70, 0x65, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tetris_proto_rawDescOnce sync.Once
	file_tetris_proto_rawDescData = file_tetris_proto_rawDesc
)

func file_tetris_proto_rawDescGZIP() []byte {
	file_tetris_proto_rawDescOnce.Do(func() {
		file_tetris_proto_rawDescData = protoimpl.X.CompressGZIP(file_tetris_proto_rawDescData)
	})
	return file_tetris_proto_rawDescData
}

var file_tetris_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tetris_proto_goTypes = []interface{}{
	(*TetrisShape)(nil), // 0: pb.TetrisShape
	(*TetrisBlock)(nil), // 1: pb.TetrisBlock
}
var file_tetris_proto_depIdxs = []int32{
	0, // 0: pb.TetrisBlock.values:type_name -> pb.TetrisShape
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tetris_proto_init() }
func file_tetris_proto_init() {
	if File_tetris_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tetris_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TetrisShape); i {
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
		file_tetris_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TetrisBlock); i {
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
			RawDescriptor: file_tetris_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tetris_proto_goTypes,
		DependencyIndexes: file_tetris_proto_depIdxs,
		MessageInfos:      file_tetris_proto_msgTypes,
	}.Build()
	File_tetris_proto = out.File
	file_tetris_proto_rawDesc = nil
	file_tetris_proto_goTypes = nil
	file_tetris_proto_depIdxs = nil
}
