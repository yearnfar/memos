// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/activity.proto

package store

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

type ActivityMemoCommentPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoId        int32 `protobuf:"varint,1,opt,name=memo_id,json=memoId,proto3" json:"memo_id,omitempty"`
	RelatedMemoId int32 `protobuf:"varint,2,opt,name=related_memo_id,json=relatedMemoId,proto3" json:"related_memo_id,omitempty"`
}

func (x *ActivityMemoCommentPayload) Reset() {
	*x = ActivityMemoCommentPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityMemoCommentPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityMemoCommentPayload) ProtoMessage() {}

func (x *ActivityMemoCommentPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityMemoCommentPayload.ProtoReflect.Descriptor instead.
func (*ActivityMemoCommentPayload) Descriptor() ([]byte, []int) {
	return file_store_activity_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityMemoCommentPayload) GetMemoId() int32 {
	if x != nil {
		return x.MemoId
	}
	return 0
}

func (x *ActivityMemoCommentPayload) GetRelatedMemoId() int32 {
	if x != nil {
		return x.RelatedMemoId
	}
	return 0
}

type ActivityVersionUpdatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ActivityVersionUpdatePayload) Reset() {
	*x = ActivityVersionUpdatePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityVersionUpdatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityVersionUpdatePayload) ProtoMessage() {}

func (x *ActivityVersionUpdatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityVersionUpdatePayload.ProtoReflect.Descriptor instead.
func (*ActivityVersionUpdatePayload) Descriptor() ([]byte, []int) {
	return file_store_activity_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityVersionUpdatePayload) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ActivityPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoComment   *ActivityMemoCommentPayload   `protobuf:"bytes,1,opt,name=memo_comment,json=memoComment,proto3" json:"memo_comment,omitempty"`
	VersionUpdate *ActivityVersionUpdatePayload `protobuf:"bytes,2,opt,name=version_update,json=versionUpdate,proto3" json:"version_update,omitempty"`
}

func (x *ActivityPayload) Reset() {
	*x = ActivityPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityPayload) ProtoMessage() {}

func (x *ActivityPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityPayload.ProtoReflect.Descriptor instead.
func (*ActivityPayload) Descriptor() ([]byte, []int) {
	return file_store_activity_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityPayload) GetMemoComment() *ActivityMemoCommentPayload {
	if x != nil {
		return x.MemoComment
	}
	return nil
}

func (x *ActivityPayload) GetVersionUpdate() *ActivityVersionUpdatePayload {
	if x != nil {
		return x.VersionUpdate
	}
	return nil
}

var File_store_activity_proto protoreflect.FileDescriptor

var file_store_activity_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x22, 0x5d, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4d,
	0x65, 0x6d, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f,
	0x49, 0x64, 0x22, 0x38, 0x0a, 0x1c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x01, 0x0a,
	0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x4a, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x6d,
	0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x0e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x9d,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x42, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x65, 0x61, 0x72, 0x6e, 0x66, 0x61, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_activity_proto_rawDescOnce sync.Once
	file_store_activity_proto_rawDescData = file_store_activity_proto_rawDesc
)

func file_store_activity_proto_rawDescGZIP() []byte {
	file_store_activity_proto_rawDescOnce.Do(func() {
		file_store_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_activity_proto_rawDescData)
	})
	return file_store_activity_proto_rawDescData
}

var file_store_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_activity_proto_goTypes = []any{
	(*ActivityMemoCommentPayload)(nil),   // 0: memos.store.ActivityMemoCommentPayload
	(*ActivityVersionUpdatePayload)(nil), // 1: memos.store.ActivityVersionUpdatePayload
	(*ActivityPayload)(nil),              // 2: memos.store.ActivityPayload
}
var file_store_activity_proto_depIdxs = []int32{
	0, // 0: memos.store.ActivityPayload.memo_comment:type_name -> memos.store.ActivityMemoCommentPayload
	1, // 1: memos.store.ActivityPayload.version_update:type_name -> memos.store.ActivityVersionUpdatePayload
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_activity_proto_init() }
func file_store_activity_proto_init() {
	if File_store_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_activity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityMemoCommentPayload); i {
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
		file_store_activity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityVersionUpdatePayload); i {
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
		file_store_activity_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityPayload); i {
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
			RawDescriptor: file_store_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_activity_proto_goTypes,
		DependencyIndexes: file_store_activity_proto_depIdxs,
		MessageInfos:      file_store_activity_proto_msgTypes,
	}.Build()
	File_store_activity_proto = out.File
	file_store_activity_proto_rawDesc = nil
	file_store_activity_proto_goTypes = nil
	file_store_activity_proto_depIdxs = nil
}
