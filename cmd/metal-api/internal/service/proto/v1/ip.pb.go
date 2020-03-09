// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.11.4
// source: v1/ip.proto

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

type IP_Type int32

const (
	IP_STATIC    IP_Type = 0
	IP_EPHEMERAL IP_Type = 1
)

// Enum value maps for IP_Type.
var (
	IP_Type_name = map[int32]string{
		0: "STATIC",
		1: "EPHEMERAL",
	}
	IP_Type_value = map[string]int32{
		"STATIC":    0,
		"EPHEMERAL": 1,
	}
)

func (x IP_Type) Enum() *IP_Type {
	p := new(IP_Type)
	*p = x
	return p
}

func (x IP_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IP_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_ip_proto_enumTypes[0].Descriptor()
}

func (IP_Type) Type() protoreflect.EnumType {
	return &file_v1_ip_proto_enumTypes[0]
}

func (x IP_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IP_Type.Descriptor instead.
func (IP_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{0, 0}
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common    *Common                 `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	ProjectID string                  `protobuf:"bytes,2,opt,name=projectID,proto3" json:"projectID,omitempty"`                  // the project this ip address belongs to
	NetworkID string                  `protobuf:"bytes,3,opt,name=networkID,proto3" json:"networkID,omitempty"`                  // the network this ip allocate request address belongs to
	Type      IP_Type                 `protobuf:"varint,4,opt,name=type,proto3,enum=metal.api.v1.IP_Type" json:"type,omitempty"` // the ip type, EPHEMERAL leads to automatic cleanup of the ip address, STATIC will enable re-use of the ip at a later point in time
	Tags      []*wrappers.StringValue `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`                            // free tags that you associate with this ip
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{0}
}

func (x *IP) GetCommon() *Common {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *IP) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *IP) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *IP) GetType() IP_Type {
	if x != nil {
		return x.Type
	}
	return IP_STATIC
}

func (x *IP) GetTags() []*wrappers.StringValue {
	if x != nil {
		return x.Tags
	}
	return nil
}

type IPIdentifiable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPAddress string `protobuf:"bytes,1,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
}

func (x *IPIdentifiable) Reset() {
	*x = IPIdentifiable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPIdentifiable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPIdentifiable) ProtoMessage() {}

func (x *IPIdentifiable) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPIdentifiable.ProtoReflect.Descriptor instead.
func (*IPIdentifiable) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{1}
}

func (x *IPIdentifiable) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

type IPAllocateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP        *IP                   `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	MachineID *wrappers.StringValue `protobuf:"bytes,2,opt,name=machineID,proto3" json:"machineID,omitempty"` // the machine id this ip should be associated with
}

func (x *IPAllocateRequest) Reset() {
	*x = IPAllocateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAllocateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAllocateRequest) ProtoMessage() {}

func (x *IPAllocateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAllocateRequest.ProtoReflect.Descriptor instead.
func (*IPAllocateRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{2}
}

func (x *IPAllocateRequest) GetIP() *IP {
	if x != nil {
		return x.IP
	}
	return nil
}

func (x *IPAllocateRequest) GetMachineID() *wrappers.StringValue {
	if x != nil {
		return x.MachineID
	}
	return nil
}

type IPUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common       *Common                 `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Identifiable *IPIdentifiable         `protobuf:"bytes,2,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
	Type         IP_Type                 `protobuf:"varint,3,opt,name=type,proto3,enum=metal.api.v1.IP_Type" json:"type,omitempty"` // the ip type, EPHEMERAL leads to automatic cleanup of the ip address, STATIC will enable re-use of the ip at a later point in time
	Tags         []*wrappers.StringValue `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`                            // free tags that you associate with this ip
}

func (x *IPUpdateRequest) Reset() {
	*x = IPUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPUpdateRequest) ProtoMessage() {}

func (x *IPUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPUpdateRequest.ProtoReflect.Descriptor instead.
func (*IPUpdateRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{3}
}

func (x *IPUpdateRequest) GetCommon() *Common {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *IPUpdateRequest) GetIdentifiable() *IPIdentifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

func (x *IPUpdateRequest) GetType() IP_Type {
	if x != nil {
		return x.Type
	}
	return IP_STATIC
}

func (x *IPUpdateRequest) GetTags() []*wrappers.StringValue {
	if x != nil {
		return x.Tags
	}
	return nil
}

type IPDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiable *Identifiable `protobuf:"bytes,1,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
}

func (x *IPDeleteRequest) Reset() {
	*x = IPDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPDeleteRequest) ProtoMessage() {}

func (x *IPDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPDeleteRequest.ProtoReflect.Descriptor instead.
func (*IPDeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{4}
}

func (x *IPDeleteRequest) GetIdentifiable() *Identifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

type IPGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiable *Identifiable `protobuf:"bytes,1,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
}

func (x *IPGetRequest) Reset() {
	*x = IPGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPGetRequest) ProtoMessage() {}

func (x *IPGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPGetRequest.ProtoReflect.Descriptor instead.
func (*IPGetRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{5}
}

func (x *IPGetRequest) GetIdentifiable() *Identifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

type IPFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPAddress        *wrappers.StringValue   `protobuf:"bytes,1,opt,name=IPAddress,json=ipaddress,proto3" json:"IPAddress,omitempty"`                   // the address (ipv4 or ipv6) of this ip
	ParentPrefixCidr *wrappers.StringValue   `protobuf:"bytes,2,opt,name=ParentPrefixCidr,json=networkprefix,proto3" json:"ParentPrefixCidr,omitempty"` // the prefix of the network this ip address belongs to
	NetworkID        *wrappers.StringValue   `protobuf:"bytes,3,opt,name=NetworkID,json=networkid,proto3" json:"NetworkID,omitempty"`                   // the network this ip allocate request address belongs to
	Tags             []*wrappers.StringValue `protobuf:"bytes,4,rep,name=Tags,json=tags,proto3" json:"Tags,omitempty"`                                  // the tags that are assigned to this ip address
	ProjectID        *wrappers.StringValue   `protobuf:"bytes,5,opt,name=ProjectID,json=projectid,proto3" json:"ProjectID,omitempty"`                   // the project this ip address belongs to, empty if not strong coupled
	Type             *wrappers.StringValue   `protobuf:"bytes,6,opt,name=Type,json=type,proto3" json:"Type,omitempty"`                                  // the type of the ip address, ephemeral or static
	MachineID        *wrappers.StringValue   `protobuf:"bytes,7,opt,name=MachineID,json=machineid,proto3" json:"MachineID,omitempty"`                   // the machine an ip address is associated to
}

func (x *IPFindRequest) Reset() {
	*x = IPFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPFindRequest) ProtoMessage() {}

func (x *IPFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPFindRequest.ProtoReflect.Descriptor instead.
func (*IPFindRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{6}
}

func (x *IPFindRequest) GetIPAddress() *wrappers.StringValue {
	if x != nil {
		return x.IPAddress
	}
	return nil
}

func (x *IPFindRequest) GetParentPrefixCidr() *wrappers.StringValue {
	if x != nil {
		return x.ParentPrefixCidr
	}
	return nil
}

func (x *IPFindRequest) GetNetworkID() *wrappers.StringValue {
	if x != nil {
		return x.NetworkID
	}
	return nil
}

func (x *IPFindRequest) GetTags() []*wrappers.StringValue {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *IPFindRequest) GetProjectID() *wrappers.StringValue {
	if x != nil {
		return x.ProjectID
	}
	return nil
}

func (x *IPFindRequest) GetType() *wrappers.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *IPFindRequest) GetMachineID() *wrappers.StringValue {
	if x != nil {
		return x.MachineID
	}
	return nil
}

type IPListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IPListRequest) Reset() {
	*x = IPListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPListRequest) ProtoMessage() {}

func (x *IPListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPListRequest.ProtoReflect.Descriptor instead.
func (*IPListRequest) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{7}
}

type IPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP           *IP             `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Identifiable *IPIdentifiable `protobuf:"bytes,2,opt,name=identifiable,proto3" json:"identifiable,omitempty"`
}

func (x *IPResponse) Reset() {
	*x = IPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPResponse) ProtoMessage() {}

func (x *IPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPResponse.ProtoReflect.Descriptor instead.
func (*IPResponse) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{8}
}

func (x *IPResponse) GetIP() *IP {
	if x != nil {
		return x.IP
	}
	return nil
}

func (x *IPResponse) GetIdentifiable() *IPIdentifiable {
	if x != nil {
		return x.Identifiable
	}
	return nil
}

type IPListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips []*IP `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
}

func (x *IPListResponse) Reset() {
	*x = IPListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_ip_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPListResponse) ProtoMessage() {}

func (x *IPListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_ip_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPListResponse.ProtoReflect.Descriptor instead.
func (*IPListResponse) Descriptor() ([]byte, []int) {
	return file_v1_ip_proto_rawDescGZIP(), []int{9}
}

func (x *IPListResponse) GetIps() []*IP {
	if x != nil {
		return x.Ips
	}
	return nil
}

var File_v1_ip_proto protoreflect.FileDescriptor

var file_v1_ip_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x31, 0x2f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x50, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x43, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x50, 0x48, 0x45, 0x4d, 0x45, 0x52,
	0x41, 0x4c, 0x10, 0x01, 0x22, 0x2e, 0x0a, 0x0e, 0x49, 0x50, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x71, 0x0a, 0x11, 0x49, 0x50, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x02, 0x49, 0x50, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x02, 0x49, 0x50, 0x12, 0x3a, 0x0a, 0x09, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x22, 0xde, 0x01, 0x0a, 0x0f, 0x49, 0x50, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0c, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x50, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x51, 0x0a, 0x0f, 0x49, 0x50, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x4e, 0x0a, 0x0c, 0x49,
	0x50, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xaa, 0x03, 0x0a, 0x0d,
	0x49, 0x50, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x3a, 0x0a, 0x09, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3a,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x50, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x0a, 0x49, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x02, 0x49, 0x50, 0x12, 0x40, 0x0a, 0x0c, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x50, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0c, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x49,
	0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x03, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x03, 0x69, 0x70,
	0x73, 0x32, 0x9b, 0x03, 0x0a, 0x09, 0x49, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x50, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x50, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_ip_proto_rawDescOnce sync.Once
	file_v1_ip_proto_rawDescData = file_v1_ip_proto_rawDesc
)

func file_v1_ip_proto_rawDescGZIP() []byte {
	file_v1_ip_proto_rawDescOnce.Do(func() {
		file_v1_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_ip_proto_rawDescData)
	})
	return file_v1_ip_proto_rawDescData
}

var file_v1_ip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_ip_proto_goTypes = []interface{}{
	(IP_Type)(0),                 // 0: metal.api.v1.IP.Type
	(*IP)(nil),                   // 1: metal.api.v1.IP
	(*IPIdentifiable)(nil),       // 2: metal.api.v1.IPIdentifiable
	(*IPAllocateRequest)(nil),    // 3: metal.api.v1.IPAllocateRequest
	(*IPUpdateRequest)(nil),      // 4: metal.api.v1.IPUpdateRequest
	(*IPDeleteRequest)(nil),      // 5: metal.api.v1.IPDeleteRequest
	(*IPGetRequest)(nil),         // 6: metal.api.v1.IPGetRequest
	(*IPFindRequest)(nil),        // 7: metal.api.v1.IPFindRequest
	(*IPListRequest)(nil),        // 8: metal.api.v1.IPListRequest
	(*IPResponse)(nil),           // 9: metal.api.v1.IPResponse
	(*IPListResponse)(nil),       // 10: metal.api.v1.IPListResponse
	(*Common)(nil),               // 11: metal.api.v1.Common
	(*wrappers.StringValue)(nil), // 12: google.protobuf.StringValue
	(*Identifiable)(nil),         // 13: metal.api.v1.Identifiable
}
var file_v1_ip_proto_depIdxs = []int32{
	11, // 0: metal.api.v1.IP.common:type_name -> metal.api.v1.Common
	0,  // 1: metal.api.v1.IP.type:type_name -> metal.api.v1.IP.Type
	12, // 2: metal.api.v1.IP.tags:type_name -> google.protobuf.StringValue
	1,  // 3: metal.api.v1.IPAllocateRequest.IP:type_name -> metal.api.v1.IP
	12, // 4: metal.api.v1.IPAllocateRequest.machineID:type_name -> google.protobuf.StringValue
	11, // 5: metal.api.v1.IPUpdateRequest.common:type_name -> metal.api.v1.Common
	2,  // 6: metal.api.v1.IPUpdateRequest.identifiable:type_name -> metal.api.v1.IPIdentifiable
	0,  // 7: metal.api.v1.IPUpdateRequest.type:type_name -> metal.api.v1.IP.Type
	12, // 8: metal.api.v1.IPUpdateRequest.tags:type_name -> google.protobuf.StringValue
	13, // 9: metal.api.v1.IPDeleteRequest.identifiable:type_name -> metal.api.v1.Identifiable
	13, // 10: metal.api.v1.IPGetRequest.identifiable:type_name -> metal.api.v1.Identifiable
	12, // 11: metal.api.v1.IPFindRequest.IPAddress:type_name -> google.protobuf.StringValue
	12, // 12: metal.api.v1.IPFindRequest.ParentPrefixCidr:type_name -> google.protobuf.StringValue
	12, // 13: metal.api.v1.IPFindRequest.NetworkID:type_name -> google.protobuf.StringValue
	12, // 14: metal.api.v1.IPFindRequest.Tags:type_name -> google.protobuf.StringValue
	12, // 15: metal.api.v1.IPFindRequest.ProjectID:type_name -> google.protobuf.StringValue
	12, // 16: metal.api.v1.IPFindRequest.Type:type_name -> google.protobuf.StringValue
	12, // 17: metal.api.v1.IPFindRequest.MachineID:type_name -> google.protobuf.StringValue
	1,  // 18: metal.api.v1.IPResponse.IP:type_name -> metal.api.v1.IP
	2,  // 19: metal.api.v1.IPResponse.identifiable:type_name -> metal.api.v1.IPIdentifiable
	1,  // 20: metal.api.v1.IPListResponse.ips:type_name -> metal.api.v1.IP
	3,  // 21: metal.api.v1.IPService.Allocate:input_type -> metal.api.v1.IPAllocateRequest
	4,  // 22: metal.api.v1.IPService.Update:input_type -> metal.api.v1.IPUpdateRequest
	5,  // 23: metal.api.v1.IPService.Delete:input_type -> metal.api.v1.IPDeleteRequest
	6,  // 24: metal.api.v1.IPService.Get:input_type -> metal.api.v1.IPGetRequest
	7,  // 25: metal.api.v1.IPService.Find:input_type -> metal.api.v1.IPFindRequest
	8,  // 26: metal.api.v1.IPService.List:input_type -> metal.api.v1.IPListRequest
	9,  // 27: metal.api.v1.IPService.Allocate:output_type -> metal.api.v1.IPResponse
	9,  // 28: metal.api.v1.IPService.Update:output_type -> metal.api.v1.IPResponse
	9,  // 29: metal.api.v1.IPService.Delete:output_type -> metal.api.v1.IPResponse
	9,  // 30: metal.api.v1.IPService.Get:output_type -> metal.api.v1.IPResponse
	10, // 31: metal.api.v1.IPService.Find:output_type -> metal.api.v1.IPListResponse
	10, // 32: metal.api.v1.IPService.List:output_type -> metal.api.v1.IPListResponse
	27, // [27:33] is the sub-list for method output_type
	21, // [21:27] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_v1_ip_proto_init() }
func file_v1_ip_proto_init() {
	if File_v1_ip_proto != nil {
		return
	}
	file_v1_common_proto_init()
	file_v1_identifiable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_v1_ip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPIdentifiable); i {
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
		file_v1_ip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAllocateRequest); i {
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
		file_v1_ip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPUpdateRequest); i {
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
		file_v1_ip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPDeleteRequest); i {
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
		file_v1_ip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPGetRequest); i {
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
		file_v1_ip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPFindRequest); i {
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
		file_v1_ip_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPListRequest); i {
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
		file_v1_ip_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPResponse); i {
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
		file_v1_ip_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPListResponse); i {
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
			RawDescriptor: file_v1_ip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_ip_proto_goTypes,
		DependencyIndexes: file_v1_ip_proto_depIdxs,
		EnumInfos:         file_v1_ip_proto_enumTypes,
		MessageInfos:      file_v1_ip_proto_msgTypes,
	}.Build()
	File_v1_ip_proto = out.File
	file_v1_ip_proto_rawDesc = nil
	file_v1_ip_proto_goTypes = nil
	file_v1_ip_proto_depIdxs = nil
}
