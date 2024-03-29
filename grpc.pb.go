// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: grpc.proto

package test

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

type AmountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       int64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Id         int64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	BidderName string  `protobuf:"bytes,3,opt,name=bidderName,proto3" json:"bidderName,omitempty"`
	Amount     float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AmountMessage) Reset() {
	*x = AmountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmountMessage) ProtoMessage() {}

func (x *AmountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmountMessage.ProtoReflect.Descriptor instead.
func (*AmountMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *AmountMessage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AmountMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AmountMessage) GetBidderName() string {
	if x != nil {
		return x.BidderName
	}
	return ""
}

func (x *AmountMessage) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AcknowledgementMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Id   int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AcknowledgementMessage) Reset() {
	*x = AcknowledgementMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcknowledgementMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcknowledgementMessage) ProtoMessage() {}

func (x *AcknowledgementMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcknowledgementMessage.ProtoReflect.Descriptor instead.
func (*AcknowledgementMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *AcknowledgementMessage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AcknowledgementMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OutcomeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time              int64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	HighestBid        float64 `protobuf:"fixed64,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	HighestBidderID   int64   `protobuf:"varint,3,opt,name=highestBidderID,proto3" json:"highestBidderID,omitempty"`
	HighestBidderName string  `protobuf:"bytes,4,opt,name=highestBidderName,proto3" json:"highestBidderName,omitempty"`
	Ended             bool    `protobuf:"varint,5,opt,name=ended,proto3" json:"ended,omitempty"`
	EndTime           int64   `protobuf:"varint,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *OutcomeMessage) Reset() {
	*x = OutcomeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutcomeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutcomeMessage) ProtoMessage() {}

func (x *OutcomeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutcomeMessage.ProtoReflect.Descriptor instead.
func (*OutcomeMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *OutcomeMessage) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *OutcomeMessage) GetHighestBid() float64 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *OutcomeMessage) GetHighestBidderID() int64 {
	if x != nil {
		return x.HighestBidderID
	}
	return 0
}

func (x *OutcomeMessage) GetHighestBidderName() string {
	if x != nil {
		return x.HighestBidderName
	}
	return ""
}

func (x *OutcomeMessage) GetEnded() bool {
	if x != nil {
		return x.Ended
	}
	return false
}

func (x *OutcomeMessage) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type Bidder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	HighestBid float64 `protobuf:"fixed64,3,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *Bidder) Reset() {
	*x = Bidder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bidder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bidder) ProtoMessage() {}

func (x *Bidder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bidder.ProtoReflect.Descriptor instead.
func (*Bidder) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Bidder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bidder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bidder) GetHighestBid() float64 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type AuctionHouseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bidders         []*Bidder `protobuf:"bytes,1,rep,name=bidders,proto3" json:"bidders,omitempty"`
	HighestBidderID int64     `protobuf:"varint,2,opt,name=highestBidderID,proto3" json:"highestBidderID,omitempty"`
	EndTime         int64     `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Ended           bool      `protobuf:"varint,4,opt,name=ended,proto3" json:"ended,omitempty"`
}

func (x *AuctionHouseMessage) Reset() {
	*x = AuctionHouseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionHouseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionHouseMessage) ProtoMessage() {}

func (x *AuctionHouseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionHouseMessage.ProtoReflect.Descriptor instead.
func (*AuctionHouseMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *AuctionHouseMessage) GetBidders() []*Bidder {
	if x != nil {
		return x.Bidders
	}
	return nil
}

func (x *AuctionHouseMessage) GetHighestBidderID() int64 {
	if x != nil {
		return x.HighestBidderID
	}
	return 0
}

func (x *AuctionHouseMessage) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *AuctionHouseMessage) GetEnded() bool {
	if x != nil {
		return x.Ended
	}
	return false
}

type BooleanMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BooleanMessage) Reset() {
	*x = BooleanMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanMessage) ProtoMessage() {}

func (x *BooleanMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanMessage.ProtoReflect.Descriptor instead.
func (*BooleanMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *BooleanMessage) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type ReplicationManagerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsLeader bool  `protobuf:"varint,2,opt,name=isLeader,proto3" json:"isLeader,omitempty"`
}

func (x *ReplicationManagerMessage) Reset() {
	*x = ReplicationManagerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationManagerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationManagerMessage) ProtoMessage() {}

func (x *ReplicationManagerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationManagerMessage.ProtoReflect.Descriptor instead.
func (*ReplicationManagerMessage) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *ReplicationManagerMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReplicationManagerMessage) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{7}
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x47, 0x52,
	0x50, 0x43, 0x65, 0x78, 0x22, 0x6b, 0x0a, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x69, 0x64, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x3c, 0x0a, 0x16, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xcc, 0x01, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x11, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4c,
	0x0a, 0x06, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x99, 0x01, 0x0a,
	0x13, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x42,
	0x69, 0x64, 0x64, 0x65, 0x72, 0x52, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x47, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x32, 0xb6, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x47, 0x52,
	0x50, 0x43, 0x65, 0x78, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x41, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0c,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x0c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x1a, 0x16, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0xa2, 0x01, 0x0a, 0x12, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x21,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x65, 0x78, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_proto_goTypes = []interface{}{
	(*AmountMessage)(nil),             // 0: GRPCex.AmountMessage
	(*AcknowledgementMessage)(nil),    // 1: GRPCex.AcknowledgementMessage
	(*OutcomeMessage)(nil),            // 2: GRPCex.OutcomeMessage
	(*Bidder)(nil),                    // 3: GRPCex.Bidder
	(*AuctionHouseMessage)(nil),       // 4: GRPCex.AuctionHouseMessage
	(*BooleanMessage)(nil),            // 5: GRPCex.BooleanMessage
	(*ReplicationManagerMessage)(nil), // 6: GRPCex.ReplicationManagerMessage
	(*Void)(nil),                      // 7: GRPCex.Void
}
var file_grpc_proto_depIdxs = []int32{
	3, // 0: GRPCex.AuctionHouseMessage.bidders:type_name -> GRPCex.Bidder
	0, // 1: GRPCex.AuctionService.bid:input_type -> GRPCex.AmountMessage
	7, // 2: GRPCex.AuctionService.Result:input_type -> GRPCex.Void
	7, // 3: GRPCex.AuctionService.IsLeader:input_type -> GRPCex.Void
	7, // 4: GRPCex.ReplicationService.RequestInfo:input_type -> GRPCex.Void
	4, // 5: GRPCex.ReplicationService.Replicate:input_type -> GRPCex.AuctionHouseMessage
	1, // 6: GRPCex.AuctionService.bid:output_type -> GRPCex.AcknowledgementMessage
	2, // 7: GRPCex.AuctionService.Result:output_type -> GRPCex.OutcomeMessage
	5, // 8: GRPCex.AuctionService.IsLeader:output_type -> GRPCex.BooleanMessage
	6, // 9: GRPCex.ReplicationService.RequestInfo:output_type -> GRPCex.ReplicationManagerMessage
	1, // 10: GRPCex.ReplicationService.Replicate:output_type -> GRPCex.AcknowledgementMessage
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmountMessage); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcknowledgementMessage); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutcomeMessage); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bidder); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionHouseMessage); i {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanMessage); i {
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
		file_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationManagerMessage); i {
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
		file_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
