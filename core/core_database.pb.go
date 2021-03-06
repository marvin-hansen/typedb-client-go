// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: core/core_database.proto

package core

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

type CoreDatabaseManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager) Reset() {
	*x = CoreDatabaseManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager) ProtoMessage() {}

func (x *CoreDatabaseManager) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0}
}

type CoreDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabase) Reset() {
	*x = CoreDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase) ProtoMessage() {}

func (x *CoreDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase.ProtoReflect.Descriptor instead.
func (*CoreDatabase) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1}
}

type CoreDatabaseManager_Contains struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager_Contains) Reset() {
	*x = CoreDatabaseManager_Contains{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Contains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Contains) ProtoMessage() {}

func (x *CoreDatabaseManager_Contains) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Contains.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Contains) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 0}
}

type CoreDatabaseManager_Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager_Create) Reset() {
	*x = CoreDatabaseManager_Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Create) ProtoMessage() {}

func (x *CoreDatabaseManager_Create) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Create.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Create) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 1}
}

type CoreDatabaseManager_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager_All) Reset() {
	*x = CoreDatabaseManager_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_All) ProtoMessage() {}

func (x *CoreDatabaseManager_All) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_All.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_All) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 2}
}

type CoreDatabaseManager_Contains_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoreDatabaseManager_Contains_Req) Reset() {
	*x = CoreDatabaseManager_Contains_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Contains_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Contains_Req) ProtoMessage() {}

func (x *CoreDatabaseManager_Contains_Req) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Contains_Req.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Contains_Req) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CoreDatabaseManager_Contains_Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CoreDatabaseManager_Contains_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contains bool `protobuf:"varint,1,opt,name=contains,proto3" json:"contains,omitempty"`
}

func (x *CoreDatabaseManager_Contains_Res) Reset() {
	*x = CoreDatabaseManager_Contains_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Contains_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Contains_Res) ProtoMessage() {}

func (x *CoreDatabaseManager_Contains_Res) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Contains_Res.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Contains_Res) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *CoreDatabaseManager_Contains_Res) GetContains() bool {
	if x != nil {
		return x.Contains
	}
	return false
}

type CoreDatabaseManager_Create_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoreDatabaseManager_Create_Req) Reset() {
	*x = CoreDatabaseManager_Create_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Create_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Create_Req) ProtoMessage() {}

func (x *CoreDatabaseManager_Create_Req) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Create_Req.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Create_Req) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *CoreDatabaseManager_Create_Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CoreDatabaseManager_Create_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager_Create_Res) Reset() {
	*x = CoreDatabaseManager_Create_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_Create_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_Create_Res) ProtoMessage() {}

func (x *CoreDatabaseManager_Create_Res) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_Create_Res.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_Create_Res) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 1, 1}
}

type CoreDatabaseManager_All_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabaseManager_All_Req) Reset() {
	*x = CoreDatabaseManager_All_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_All_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_All_Req) ProtoMessage() {}

func (x *CoreDatabaseManager_All_Req) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_All_Req.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_All_Req) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 2, 0}
}

type CoreDatabaseManager_All_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *CoreDatabaseManager_All_Res) Reset() {
	*x = CoreDatabaseManager_All_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabaseManager_All_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabaseManager_All_Res) ProtoMessage() {}

func (x *CoreDatabaseManager_All_Res) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabaseManager_All_Res.ProtoReflect.Descriptor instead.
func (*CoreDatabaseManager_All_Res) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *CoreDatabaseManager_All_Res) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type CoreDatabase_Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabase_Schema) Reset() {
	*x = CoreDatabase_Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Schema) ProtoMessage() {}

func (x *CoreDatabase_Schema) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Schema.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Schema) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 0}
}

type CoreDatabase_Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabase_Delete) Reset() {
	*x = CoreDatabase_Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Delete) ProtoMessage() {}

func (x *CoreDatabase_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Delete.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Delete) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 1}
}

type CoreDatabase_Schema_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoreDatabase_Schema_Req) Reset() {
	*x = CoreDatabase_Schema_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Schema_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Schema_Req) ProtoMessage() {}

func (x *CoreDatabase_Schema_Req) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Schema_Req.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Schema_Req) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *CoreDatabase_Schema_Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CoreDatabase_Schema_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *CoreDatabase_Schema_Res) Reset() {
	*x = CoreDatabase_Schema_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Schema_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Schema_Res) ProtoMessage() {}

func (x *CoreDatabase_Schema_Res) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Schema_Res.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Schema_Res) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *CoreDatabase_Schema_Res) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type CoreDatabase_Delete_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CoreDatabase_Delete_Req) Reset() {
	*x = CoreDatabase_Delete_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Delete_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Delete_Req) ProtoMessage() {}

func (x *CoreDatabase_Delete_Req) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Delete_Req.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Delete_Req) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *CoreDatabase_Delete_Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CoreDatabase_Delete_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreDatabase_Delete_Res) Reset() {
	*x = CoreDatabase_Delete_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_core_database_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDatabase_Delete_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDatabase_Delete_Res) ProtoMessage() {}

func (x *CoreDatabase_Delete_Res) ProtoReflect() protoreflect.Message {
	mi := &file_core_core_database_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDatabase_Delete_Res.ProtoReflect.Descriptor instead.
func (*CoreDatabase_Delete_Res) Descriptor() ([]byte, []int) {
	return file_core_core_database_proto_rawDescGZIP(), []int{1, 1, 1}
}

var File_core_core_database_proto protoreflect.FileDescriptor

var file_core_core_database_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x13,
	0x43, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x1a, 0x48, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x1a,
	0x19, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x21, 0x0a, 0x03, 0x52, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x2a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x19, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x1a, 0x29, 0x0a, 0x03, 0x41, 0x6c, 0x6c,
	0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x0c, 0x43, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x1a, 0x42, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19,
	0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1d, 0x0a, 0x03, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x19, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x05, 0x0a,
	0x03, 0x52, 0x65, 0x73, 0x42, 0x30, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x42, 0x11, 0x43, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_core_database_proto_rawDescOnce sync.Once
	file_core_core_database_proto_rawDescData = file_core_core_database_proto_rawDesc
)

func file_core_core_database_proto_rawDescGZIP() []byte {
	file_core_core_database_proto_rawDescOnce.Do(func() {
		file_core_core_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_core_database_proto_rawDescData)
	})
	return file_core_core_database_proto_rawDescData
}

var file_core_core_database_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_core_core_database_proto_goTypes = []interface{}{
	(*CoreDatabaseManager)(nil),              // 0: typedb.protocol.CoreDatabaseManager
	(*CoreDatabase)(nil),                     // 1: typedb.protocol.CoreDatabase
	(*CoreDatabaseManager_Contains)(nil),     // 2: typedb.protocol.CoreDatabaseManager.Contains
	(*CoreDatabaseManager_Create)(nil),       // 3: typedb.protocol.CoreDatabaseManager.Create
	(*CoreDatabaseManager_All)(nil),          // 4: typedb.protocol.CoreDatabaseManager.All
	(*CoreDatabaseManager_Contains_Req)(nil), // 5: typedb.protocol.CoreDatabaseManager.Contains.Req
	(*CoreDatabaseManager_Contains_Res)(nil), // 6: typedb.protocol.CoreDatabaseManager.Contains.Res
	(*CoreDatabaseManager_Create_Req)(nil),   // 7: typedb.protocol.CoreDatabaseManager.Create.Req
	(*CoreDatabaseManager_Create_Res)(nil),   // 8: typedb.protocol.CoreDatabaseManager.Create.Res
	(*CoreDatabaseManager_All_Req)(nil),      // 9: typedb.protocol.CoreDatabaseManager.All.Req
	(*CoreDatabaseManager_All_Res)(nil),      // 10: typedb.protocol.CoreDatabaseManager.All.Res
	(*CoreDatabase_Schema)(nil),              // 11: typedb.protocol.CoreDatabase.Schema
	(*CoreDatabase_Delete)(nil),              // 12: typedb.protocol.CoreDatabase.Delete
	(*CoreDatabase_Schema_Req)(nil),          // 13: typedb.protocol.CoreDatabase.Schema.Req
	(*CoreDatabase_Schema_Res)(nil),          // 14: typedb.protocol.CoreDatabase.Schema.Res
	(*CoreDatabase_Delete_Req)(nil),          // 15: typedb.protocol.CoreDatabase.Delete.Req
	(*CoreDatabase_Delete_Res)(nil),          // 16: typedb.protocol.CoreDatabase.Delete.Res
}
var file_core_core_database_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_core_database_proto_init() }
func file_core_core_database_proto_init() {
	if File_core_core_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_core_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager); i {
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
		file_core_core_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase); i {
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
		file_core_core_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Contains); i {
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
		file_core_core_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Create); i {
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
		file_core_core_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_All); i {
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
		file_core_core_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Contains_Req); i {
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
		file_core_core_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Contains_Res); i {
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
		file_core_core_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Create_Req); i {
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
		file_core_core_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_Create_Res); i {
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
		file_core_core_database_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_All_Req); i {
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
		file_core_core_database_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabaseManager_All_Res); i {
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
		file_core_core_database_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Schema); i {
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
		file_core_core_database_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Delete); i {
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
		file_core_core_database_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Schema_Req); i {
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
		file_core_core_database_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Schema_Res); i {
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
		file_core_core_database_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Delete_Req); i {
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
		file_core_core_database_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDatabase_Delete_Res); i {
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
			RawDescriptor: file_core_core_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_core_database_proto_goTypes,
		DependencyIndexes: file_core_core_database_proto_depIdxs,
		MessageInfos:      file_core_core_database_proto_msgTypes,
	}.Build()
	File_core_core_database_proto = out.File
	file_core_core_database_proto_rawDesc = nil
	file_core_core_database_proto_goTypes = nil
	file_core_core_database_proto_depIdxs = nil
}
