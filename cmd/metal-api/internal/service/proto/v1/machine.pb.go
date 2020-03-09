// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.11.4
// source: v1/machine.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common   *Common                 `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Url      *wrappers.StringValue   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`           // the url of this machine
	Features []*wrappers.StringValue `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"` // features of this machine
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{0}
}

func (x *Machine) GetCommon() *Common {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *Machine) GetUrl() *wrappers.StringValue {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Machine) GetFeatures() []*wrappers.StringValue {
	if x != nil {
		return x.Features
	}
	return nil
}

type MachineCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *MachineCreateRequest) Reset() {
	*x = MachineCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineCreateRequest) ProtoMessage() {}

func (x *MachineCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineCreateRequest.ProtoReflect.Descriptor instead.
func (*MachineCreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{1}
}

func (x *MachineCreateRequest) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *MachineUpdateRequest) Reset() {
	*x = MachineUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineUpdateRequest) ProtoMessage() {}

func (x *MachineUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineUpdateRequest.ProtoReflect.Descriptor instead.
func (*MachineUpdateRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{2}
}

func (x *MachineUpdateRequest) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiable *Identifiable `protobuf:"bytes,1,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
}

func (x *MachineDeleteRequest) Reset() {
	*x = MachineDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineDeleteRequest) ProtoMessage() {}

func (x *MachineDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineDeleteRequest.ProtoReflect.Descriptor instead.
func (*MachineDeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{3}
}

func (x *MachineDeleteRequest) GetIdentifiable() *Identifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

type MachineGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiable *Identifiable `protobuf:"bytes,1,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
}

func (x *MachineGetRequest) Reset() {
	*x = MachineGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineGetRequest) ProtoMessage() {}

func (x *MachineGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineGetRequest.ProtoReflect.Descriptor instead.
func (*MachineGetRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{4}
}

func (x *MachineGetRequest) GetIdentifiable() *Identifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

type MachineFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FindCriteria *FindCriteria `protobuf:"bytes,1,opt,name=findCriteria,proto3" json:"findCriteria,omitempty"`
}

func (x *MachineFindRequest) Reset() {
	*x = MachineFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineFindRequest) ProtoMessage() {}

func (x *MachineFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineFindRequest.ProtoReflect.Descriptor instead.
func (*MachineFindRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{5}
}

func (x *MachineFindRequest) GetFindCriteria() *FindCriteria {
	if x != nil {
		return x.FindCriteria
	}
	return nil
}

type MachineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *MachineResponse) Reset() {
	*x = MachineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineResponse) ProtoMessage() {}

func (x *MachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineResponse.ProtoReflect.Descriptor instead.
func (*MachineResponse) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{6}
}

func (x *MachineResponse) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MachineListRequest) Reset() {
	*x = MachineListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListRequest) ProtoMessage() {}

func (x *MachineListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListRequest.ProtoReflect.Descriptor instead.
func (*MachineListRequest) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{7}
}

type MachineListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *MachineListResponse) Reset() {
	*x = MachineListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_machine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListResponse) ProtoMessage() {}

func (x *MachineListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_machine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListResponse.ProtoReflect.Descriptor instead.
func (*MachineListResponse) Descriptor() ([]byte, []int) {
	return file_v1_machine_proto_rawDescGZIP(), []int{8}
}

func (x *MachineListResponse) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

var File_v1_machine_proto protoreflect.FileDescriptor

var file_v1_machine_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x14,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0x56,
	0x0a, 0x14, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x22, 0x42, 0x0a, 0x0f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x13, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x32, 0xd8, 0x03, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x22, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x20,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_machine_proto_rawDescOnce sync.Once
	file_v1_machine_proto_rawDescData = file_v1_machine_proto_rawDesc
)

func file_v1_machine_proto_rawDescGZIP() []byte {
	file_v1_machine_proto_rawDescOnce.Do(func() {
		file_v1_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_machine_proto_rawDescData)
	})
	return file_v1_machine_proto_rawDescData
}

var file_v1_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_machine_proto_goTypes = []interface{}{
	(*Machine)(nil),              // 0: metal.api.v1.Machine
	(*MachineCreateRequest)(nil), // 1: metal.api.v1.MachineCreateRequest
	(*MachineUpdateRequest)(nil), // 2: metal.api.v1.MachineUpdateRequest
	(*MachineDeleteRequest)(nil), // 3: metal.api.v1.MachineDeleteRequest
	(*MachineGetRequest)(nil),    // 4: metal.api.v1.MachineGetRequest
	(*MachineFindRequest)(nil),   // 5: metal.api.v1.MachineFindRequest
	(*MachineResponse)(nil),      // 6: metal.api.v1.MachineResponse
	(*MachineListRequest)(nil),   // 7: metal.api.v1.MachineListRequest
	(*MachineListResponse)(nil),  // 8: metal.api.v1.MachineListResponse
	(*Common)(nil),               // 9: metal.api.v1.Common
	(*wrappers.StringValue)(nil), // 10: google.protobuf.StringValue
	(*Identifiable)(nil),         // 11: metal.api.v1.Identifiable
	(*FindCriteria)(nil),         // 12: metal.api.v1.FindCriteria
}
var file_v1_machine_proto_depIdxs = []int32{
	9,  // 0: metal.api.v1.Machine.common:type_name -> metal.api.v1.Common
	10, // 1: metal.api.v1.Machine.url:type_name -> google.protobuf.StringValue
	10, // 2: metal.api.v1.Machine.features:type_name -> google.protobuf.StringValue
	0,  // 3: metal.api.v1.MachineCreateRequest.machine:type_name -> metal.api.v1.Machine
	0,  // 4: metal.api.v1.MachineUpdateRequest.machine:type_name -> metal.api.v1.Machine
	11, // 5: metal.api.v1.MachineDeleteRequest.identifiable:type_name -> metal.api.v1.Identifiable
	11, // 6: metal.api.v1.MachineGetRequest.identifiable:type_name -> metal.api.v1.Identifiable
	12, // 7: metal.api.v1.MachineFindRequest.findCriteria:type_name -> metal.api.v1.FindCriteria
	0,  // 8: metal.api.v1.MachineResponse.machine:type_name -> metal.api.v1.Machine
	0,  // 9: metal.api.v1.MachineListResponse.machines:type_name -> metal.api.v1.Machine
	1,  // 10: metal.api.v1.MachineService.Create:input_type -> metal.api.v1.MachineCreateRequest
	2,  // 11: metal.api.v1.MachineService.Update:input_type -> metal.api.v1.MachineUpdateRequest
	3,  // 12: metal.api.v1.MachineService.Delete:input_type -> metal.api.v1.MachineDeleteRequest
	4,  // 13: metal.api.v1.MachineService.Get:input_type -> metal.api.v1.MachineGetRequest
	5,  // 14: metal.api.v1.MachineService.Find:input_type -> metal.api.v1.MachineFindRequest
	7,  // 15: metal.api.v1.MachineService.List:input_type -> metal.api.v1.MachineListRequest
	6,  // 16: metal.api.v1.MachineService.Create:output_type -> metal.api.v1.MachineResponse
	6,  // 17: metal.api.v1.MachineService.Update:output_type -> metal.api.v1.MachineResponse
	6,  // 18: metal.api.v1.MachineService.Delete:output_type -> metal.api.v1.MachineResponse
	6,  // 19: metal.api.v1.MachineService.Get:output_type -> metal.api.v1.MachineResponse
	8,  // 20: metal.api.v1.MachineService.Find:output_type -> metal.api.v1.MachineListResponse
	8,  // 21: metal.api.v1.MachineService.List:output_type -> metal.api.v1.MachineListResponse
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_v1_machine_proto_init() }
func file_v1_machine_proto_init() {
	if File_v1_machine_proto != nil {
		return
	}
	file_v1_common_proto_init()
	file_v1_identifiable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_v1_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineCreateRequest); i {
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
		file_v1_machine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineUpdateRequest); i {
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
		file_v1_machine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineDeleteRequest); i {
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
		file_v1_machine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineGetRequest); i {
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
		file_v1_machine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineFindRequest); i {
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
		file_v1_machine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineResponse); i {
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
		file_v1_machine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineListRequest); i {
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
		file_v1_machine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineListResponse); i {
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
			RawDescriptor: file_v1_machine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_machine_proto_goTypes,
		DependencyIndexes: file_v1_machine_proto_depIdxs,
		MessageInfos:      file_v1_machine_proto_msgTypes,
	}.Build()
	File_v1_machine_proto = out.File
	file_v1_machine_proto_rawDesc = nil
	file_v1_machine_proto_goTypes = nil
	file_v1_machine_proto_depIdxs = nil
}