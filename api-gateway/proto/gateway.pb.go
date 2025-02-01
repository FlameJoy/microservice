// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.27.2
// source: proto/gateway.proto

package proto

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

// Auth-service
type GatewayLoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayLoginReq) Reset() {
	*x = GatewayLoginReq{}
	mi := &file_proto_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLoginReq) ProtoMessage() {}

func (x *GatewayLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLoginReq.ProtoReflect.Descriptor instead.
func (*GatewayLoginReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GatewayLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GatewayLoginResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayLoginResp) Reset() {
	*x = GatewayLoginResp{}
	mi := &file_proto_gateway_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayLoginResp) ProtoMessage() {}

func (x *GatewayLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayLoginResp.ProtoReflect.Descriptor instead.
func (*GatewayLoginResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayLoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GatewayLoginResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GatewayRegisterReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayRegisterReq) Reset() {
	*x = GatewayRegisterReq{}
	mi := &file_proto_gateway_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRegisterReq) ProtoMessage() {}

func (x *GatewayRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRegisterReq.ProtoReflect.Descriptor instead.
func (*GatewayRegisterReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayRegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GatewayRegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GatewayRegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GatewayRegisterResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayRegisterResp) Reset() {
	*x = GatewayRegisterResp{}
	mi := &file_proto_gateway_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayRegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRegisterResp) ProtoMessage() {}

func (x *GatewayRegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRegisterResp.ProtoReflect.Descriptor instead.
func (*GatewayRegisterResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *GatewayRegisterResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Product-service
type GatewayCreateProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SKU           string                 `protobuf:"bytes,1,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price         int64                  `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Category      string                 `protobuf:"bytes,4,opt,name=Category,proto3" json:"Category,omitempty"`
	UOM           string                 `protobuf:"bytes,5,opt,name=UOM,proto3" json:"UOM,omitempty"`
	Brand         string                 `protobuf:"bytes,6,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Stock         int64                  `protobuf:"varint,7,opt,name=Stock,proto3" json:"Stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayCreateProductReq) Reset() {
	*x = GatewayCreateProductReq{}
	mi := &file_proto_gateway_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayCreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayCreateProductReq) ProtoMessage() {}

func (x *GatewayCreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayCreateProductReq.ProtoReflect.Descriptor instead.
func (*GatewayCreateProductReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *GatewayCreateProductReq) GetSKU() string {
	if x != nil {
		return x.SKU
	}
	return ""
}

func (x *GatewayCreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GatewayCreateProductReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GatewayCreateProductReq) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GatewayCreateProductReq) GetUOM() string {
	if x != nil {
		return x.UOM
	}
	return ""
}

func (x *GatewayCreateProductReq) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *GatewayCreateProductReq) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GatewayCreateProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayCreateProductResp) Reset() {
	*x = GatewayCreateProductResp{}
	mi := &file_proto_gateway_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayCreateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayCreateProductResp) ProtoMessage() {}

func (x *GatewayCreateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayCreateProductResp.ProtoReflect.Descriptor instead.
func (*GatewayCreateProductResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *GatewayCreateProductResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GatewayCreateProductResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GatewayUpdateProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SqlQuery      string                 `protobuf:"bytes,1,opt,name=sqlQuery,proto3" json:"sqlQuery,omitempty"`
	Args          []string               `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"` // Аргументы запроса
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayUpdateProductReq) Reset() {
	*x = GatewayUpdateProductReq{}
	mi := &file_proto_gateway_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayUpdateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayUpdateProductReq) ProtoMessage() {}

func (x *GatewayUpdateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayUpdateProductReq.ProtoReflect.Descriptor instead.
func (*GatewayUpdateProductReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *GatewayUpdateProductReq) GetSqlQuery() string {
	if x != nil {
		return x.SqlQuery
	}
	return ""
}

func (x *GatewayUpdateProductReq) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type GatewayUpdateProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Флаг успешного выполнения
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayUpdateProductResp) Reset() {
	*x = GatewayUpdateProductResp{}
	mi := &file_proto_gateway_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayUpdateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayUpdateProductResp) ProtoMessage() {}

func (x *GatewayUpdateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayUpdateProductResp.ProtoReflect.Descriptor instead.
func (*GatewayUpdateProductResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *GatewayUpdateProductResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GatewayUpdateProductResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GatewayDeleteProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayDeleteProductReq) Reset() {
	*x = GatewayDeleteProductReq{}
	mi := &file_proto_gateway_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayDeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayDeleteProductReq) ProtoMessage() {}

func (x *GatewayDeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayDeleteProductReq.ProtoReflect.Descriptor instead.
func (*GatewayDeleteProductReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *GatewayDeleteProductReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GatewayDeleteProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayDeleteProductResp) Reset() {
	*x = GatewayDeleteProductResp{}
	mi := &file_proto_gateway_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayDeleteProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayDeleteProductResp) ProtoMessage() {}

func (x *GatewayDeleteProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayDeleteProductResp.ProtoReflect.Descriptor instead.
func (*GatewayDeleteProductResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *GatewayDeleteProductResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GatewayOrderCreateReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemID        int32                  `protobuf:"varint,1,opt,name=itemID,proto3" json:"itemID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         int32                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayOrderCreateReq) Reset() {
	*x = GatewayOrderCreateReq{}
	mi := &file_proto_gateway_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayOrderCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayOrderCreateReq) ProtoMessage() {}

func (x *GatewayOrderCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayOrderCreateReq.ProtoReflect.Descriptor instead.
func (*GatewayOrderCreateReq) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{10}
}

func (x *GatewayOrderCreateReq) GetItemID() int32 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *GatewayOrderCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GatewayOrderCreateReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GatewayOrderCreateReq) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GatewayOrderCreateResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TotalSum      int32                  `protobuf:"varint,3,opt,name=totalSum,proto3" json:"totalSum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GatewayOrderCreateResp) Reset() {
	*x = GatewayOrderCreateResp{}
	mi := &file_proto_gateway_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayOrderCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayOrderCreateResp) ProtoMessage() {}

func (x *GatewayOrderCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gateway_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayOrderCreateResp.ProtoReflect.Descriptor instead.
func (*GatewayOrderCreateResp) Descriptor() ([]byte, []int) {
	return file_proto_gateway_proto_rawDescGZIP(), []int{11}
}

func (x *GatewayOrderCreateResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GatewayOrderCreateResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GatewayOrderCreateResp) GetTotalSum() int32 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

var File_proto_gateway_proto protoreflect.FileDescriptor

var file_proto_gateway_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x49,
	0x0a, 0x0f, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x62, 0x0a,
	0x12, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x53, 0x4b, 0x55, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x4b, 0x55,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x4f, 0x4d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x4f, 0x4d, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x44, 0x0a, 0x18, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x17, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x4e, 0x0a, 0x18, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x34, 0x0a, 0x18, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x58, 0x0a,
	0x16, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x32, 0xe7, 0x03, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gateway_proto_rawDescOnce sync.Once
	file_proto_gateway_proto_rawDescData = file_proto_gateway_proto_rawDesc
)

func file_proto_gateway_proto_rawDescGZIP() []byte {
	file_proto_gateway_proto_rawDescOnce.Do(func() {
		file_proto_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gateway_proto_rawDescData)
	})
	return file_proto_gateway_proto_rawDescData
}

var file_proto_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_gateway_proto_goTypes = []any{
	(*GatewayLoginReq)(nil),          // 0: gateway.GatewayLoginReq
	(*GatewayLoginResp)(nil),         // 1: gateway.GatewayLoginResp
	(*GatewayRegisterReq)(nil),       // 2: gateway.GatewayRegisterReq
	(*GatewayRegisterResp)(nil),      // 3: gateway.GatewayRegisterResp
	(*GatewayCreateProductReq)(nil),  // 4: gateway.GatewayCreateProductReq
	(*GatewayCreateProductResp)(nil), // 5: gateway.GatewayCreateProductResp
	(*GatewayUpdateProductReq)(nil),  // 6: gateway.GatewayUpdateProductReq
	(*GatewayUpdateProductResp)(nil), // 7: gateway.GatewayUpdateProductResp
	(*GatewayDeleteProductReq)(nil),  // 8: gateway.GatewayDeleteProductReq
	(*GatewayDeleteProductResp)(nil), // 9: gateway.GatewayDeleteProductResp
	(*GatewayOrderCreateReq)(nil),    // 10: gateway.GatewayOrderCreateReq
	(*GatewayOrderCreateResp)(nil),   // 11: gateway.GatewayOrderCreateResp
}
var file_proto_gateway_proto_depIdxs = []int32{
	0,  // 0: gateway.GatewayService.Login:input_type -> gateway.GatewayLoginReq
	2,  // 1: gateway.GatewayService.Register:input_type -> gateway.GatewayRegisterReq
	4,  // 2: gateway.GatewayService.CreateProduct:input_type -> gateway.GatewayCreateProductReq
	6,  // 3: gateway.GatewayService.UpdateProduct:input_type -> gateway.GatewayUpdateProductReq
	8,  // 4: gateway.GatewayService.DeleteProduct:input_type -> gateway.GatewayDeleteProductReq
	10, // 5: gateway.GatewayService.CreateOrder:input_type -> gateway.GatewayOrderCreateReq
	1,  // 6: gateway.GatewayService.Login:output_type -> gateway.GatewayLoginResp
	3,  // 7: gateway.GatewayService.Register:output_type -> gateway.GatewayRegisterResp
	5,  // 8: gateway.GatewayService.CreateProduct:output_type -> gateway.GatewayCreateProductResp
	7,  // 9: gateway.GatewayService.UpdateProduct:output_type -> gateway.GatewayUpdateProductResp
	9,  // 10: gateway.GatewayService.DeleteProduct:output_type -> gateway.GatewayDeleteProductResp
	11, // 11: gateway.GatewayService.CreateOrder:output_type -> gateway.GatewayOrderCreateResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gateway_proto_init() }
func file_proto_gateway_proto_init() {
	if File_proto_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gateway_proto_goTypes,
		DependencyIndexes: file_proto_gateway_proto_depIdxs,
		MessageInfos:      file_proto_gateway_proto_msgTypes,
	}.Build()
	File_proto_gateway_proto = out.File
	file_proto_gateway_proto_rawDesc = nil
	file_proto_gateway_proto_goTypes = nil
	file_proto_gateway_proto_depIdxs = nil
}
