// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: proto/courier.proto

package courier_proto

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

type NewCourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *NewCourierRequest) Reset() {
	*x = NewCourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCourierRequest) ProtoMessage() {}

func (x *NewCourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCourierRequest.ProtoReflect.Descriptor instead.
func (*NewCourierRequest) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{0}
}

func (x *NewCourierRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NewCourierRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NewCourierRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *NewCourierRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *NewCourierRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewCourierRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UpdateCourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  string `protobuf:"bytes,5,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone     string `protobuf:"bytes,7,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *UpdateCourierRequest) Reset() {
	*x = UpdateCourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourierRequest) ProtoMessage() {}

func (x *UpdateCourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourierRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourierRequest) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCourierRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCourierRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateCourierRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UpdateCourierRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UpdateCourierRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateCourierRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CourierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Available bool   `protobuf:"varint,7,opt,name=Available,proto3" json:"Available,omitempty"`
}

func (x *CourierResponse) Reset() {
	*x = CourierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierResponse) ProtoMessage() {}

func (x *CourierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierResponse.ProtoReflect.Descriptor instead.
func (*CourierResponse) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{2}
}

func (x *CourierResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CourierResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CourierResponse) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CourierResponse) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CourierResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CourierResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CourierResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type CourierListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierListResponse []*CourierResponse `protobuf:"bytes,1,rep,name=CourierListResponse,proto3" json:"CourierListResponse,omitempty"`
}

func (x *CourierListResponse) Reset() {
	*x = CourierListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierListResponse) ProtoMessage() {}

func (x *CourierListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierListResponse.ProtoReflect.Descriptor instead.
func (*CourierListResponse) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{3}
}

func (x *CourierListResponse) GetCourierListResponse() []*CourierResponse {
	if x != nil {
		return x.CourierListResponse
	}
	return nil
}

type ParamCourier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available *string `protobuf:"bytes,1,opt,name=available,proto3,oneof" json:"available,omitempty"`
	Id        *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *ParamCourier) Reset() {
	*x = ParamCourier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamCourier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamCourier) ProtoMessage() {}

func (x *ParamCourier) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamCourier.ProtoReflect.Descriptor instead.
func (*ParamCourier) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{4}
}

func (x *ParamCourier) GetAvailable() string {
	if x != nil && x.Available != nil {
		return *x.Available
	}
	return ""
}

func (x *ParamCourier) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type CourierID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierID int64 `protobuf:"varint,1,opt,name=CourierID,proto3" json:"CourierID,omitempty"`
}

func (x *CourierID) Reset() {
	*x = CourierID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierID) ProtoMessage() {}

func (x *CourierID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierID.ProtoReflect.Descriptor instead.
func (*CourierID) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{5}
}

func (x *CourierID) GetCourierID() int64 {
	if x != nil {
		return x.CourierID
	}
	return 0
}

type CourierDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierDeleteResponse string `protobuf:"bytes,1,opt,name=CourierDeleteResponse,proto3" json:"CourierDeleteResponse,omitempty"`
}

func (x *CourierDeleteResponse) Reset() {
	*x = CourierDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierDeleteResponse) ProtoMessage() {}

func (x *CourierDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierDeleteResponse.ProtoReflect.Descriptor instead.
func (*CourierDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{6}
}

func (x *CourierDeleteResponse) GetCourierDeleteResponse() string {
	if x != nil {
		return x.CourierDeleteResponse
	}
	return ""
}

type ParamLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City *string `protobuf:"bytes,1,opt,name=city,proto3,oneof" json:"city,omitempty"`
}

func (x *ParamLocation) Reset() {
	*x = ParamLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamLocation) ProtoMessage() {}

func (x *ParamLocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamLocation.ProtoReflect.Descriptor instead.
func (*ParamLocation) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{7}
}

func (x *ParamLocation) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Latitude   *string `protobuf:"bytes,2,opt,name=Latitude,proto3,oneof" json:"Latitude,omitempty"`
	Longitude  *string `protobuf:"bytes,3,opt,name=Longitude,proto3,oneof" json:"Longitude,omitempty"`
	Country    *string `protobuf:"bytes,4,opt,name=Country,proto3,oneof" json:"Country,omitempty"`
	City       *string `protobuf:"bytes,5,opt,name=City,proto3,oneof" json:"City,omitempty"`
	Region     *string `protobuf:"bytes,6,opt,name=Region,proto3,oneof" json:"Region,omitempty"`
	Street     *string `protobuf:"bytes,7,opt,name=Street,proto3,oneof" json:"Street,omitempty"`
	HomeNumber *string `protobuf:"bytes,8,opt,name=HomeNumber,proto3,oneof" json:"HomeNumber,omitempty"`
	Floor      *string `protobuf:"bytes,9,opt,name=Floor,proto3,oneof" json:"Floor,omitempty"`
	Door       *string `protobuf:"bytes,10,opt,name=Door,proto3,oneof" json:"Door,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{8}
}

func (x *Location) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Location) GetLatitude() string {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return ""
}

func (x *Location) GetLongitude() string {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return ""
}

func (x *Location) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *Location) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *Location) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *Location) GetStreet() string {
	if x != nil && x.Street != nil {
		return *x.Street
	}
	return ""
}

func (x *Location) GetHomeNumber() string {
	if x != nil && x.HomeNumber != nil {
		return *x.HomeNumber
	}
	return ""
}

func (x *Location) GetFloor() string {
	if x != nil && x.Floor != nil {
		return *x.Floor
	}
	return ""
}

func (x *Location) GetDoor() string {
	if x != nil && x.Door != nil {
		return *x.Door
	}
	return ""
}

type LocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationList []*Location `protobuf:"bytes,1,rep,name=LocationList,proto3" json:"LocationList,omitempty"`
}

func (x *LocationList) Reset() {
	*x = LocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationList) ProtoMessage() {}

func (x *LocationList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationList.ProtoReflect.Descriptor instead.
func (*LocationList) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{9}
}

func (x *LocationList) GetLocationList() []*Location {
	if x != nil {
		return x.LocationList
	}
	return nil
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_courier_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_courier_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_proto_courier_proto_rawDescGZIP(), []int{10}
}

func (x *UserID) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_proto_courier_proto protoreflect.FileDescriptor

var file_proto_courier_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0xb1,
	0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xc1, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x61, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x13, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x22, 0x29, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x15,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x0d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x22, 0x99,
	0x03, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x48, 0x6f, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x0a, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52,
	0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x44, 0x6f, 0x6f,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43, 0x69,
	0x74, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x48, 0x6f, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x46, 0x6c, 0x6f, 0x6f,
	0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x44, 0x6f, 0x6f, 0x72, 0x22, 0x45, 0x0a, 0x0c, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x20, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x32, 0xab, 0x05, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12,
	0x4a, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x4e, 0x65,
	0x77, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x67,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4e, 0x65, 0x77, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0f, 0x67, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_courier_proto_rawDescOnce sync.Once
	file_proto_courier_proto_rawDescData = file_proto_courier_proto_rawDesc
)

func file_proto_courier_proto_rawDescGZIP() []byte {
	file_proto_courier_proto_rawDescOnce.Do(func() {
		file_proto_courier_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_courier_proto_rawDescData)
	})
	return file_proto_courier_proto_rawDescData
}

var file_proto_courier_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_courier_proto_goTypes = []interface{}{
	(*NewCourierRequest)(nil),     // 0: courier.NewCourierRequest
	(*UpdateCourierRequest)(nil),  // 1: courier.UpdateCourierRequest
	(*CourierResponse)(nil),       // 2: courier.CourierResponse
	(*CourierListResponse)(nil),   // 3: courier.CourierListResponse
	(*ParamCourier)(nil),          // 4: courier.ParamCourier
	(*CourierID)(nil),             // 5: courier.CourierID
	(*CourierDeleteResponse)(nil), // 6: courier.CourierDeleteResponse
	(*ParamLocation)(nil),         // 7: courier.ParamLocation
	(*Location)(nil),              // 8: courier.Location
	(*LocationList)(nil),          // 9: courier.LocationList
	(*UserID)(nil),                // 10: courier.userID
}
var file_proto_courier_proto_depIdxs = []int32{
	2,  // 0: courier.CourierListResponse.CourierListResponse:type_name -> courier.CourierResponse
	8,  // 1: courier.LocationList.LocationList:type_name -> courier.Location
	0,  // 2: courier.Courier.insertNewCourier:input_type -> courier.NewCourierRequest
	4,  // 3: courier.Courier.getAllCourier:input_type -> courier.ParamCourier
	5,  // 4: courier.Courier.deleteCourier:input_type -> courier.CourierID
	1,  // 5: courier.Courier.updateCourier:input_type -> courier.UpdateCourierRequest
	4,  // 6: courier.Courier.updateCourierAvailable:input_type -> courier.ParamCourier
	5,  // 7: courier.Courier.getCourier:input_type -> courier.CourierID
	8,  // 8: courier.Courier.insertNewLocation:input_type -> courier.Location
	8,  // 9: courier.Courier.updateLocation:input_type -> courier.Location
	10, // 10: courier.Courier.getLocation:input_type -> courier.userID
	7,  // 11: courier.Courier.getLocationList:input_type -> courier.ParamLocation
	2,  // 12: courier.Courier.insertNewCourier:output_type -> courier.CourierResponse
	3,  // 13: courier.Courier.getAllCourier:output_type -> courier.CourierListResponse
	6,  // 14: courier.Courier.deleteCourier:output_type -> courier.CourierDeleteResponse
	2,  // 15: courier.Courier.updateCourier:output_type -> courier.CourierResponse
	2,  // 16: courier.Courier.updateCourierAvailable:output_type -> courier.CourierResponse
	2,  // 17: courier.Courier.getCourier:output_type -> courier.CourierResponse
	8,  // 18: courier.Courier.insertNewLocation:output_type -> courier.Location
	8,  // 19: courier.Courier.updateLocation:output_type -> courier.Location
	8,  // 20: courier.Courier.getLocation:output_type -> courier.Location
	9,  // 21: courier.Courier.getLocationList:output_type -> courier.LocationList
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_courier_proto_init() }
func file_proto_courier_proto_init() {
	if File_proto_courier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_courier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCourierRequest); i {
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
		file_proto_courier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourierRequest); i {
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
		file_proto_courier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierResponse); i {
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
		file_proto_courier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierListResponse); i {
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
		file_proto_courier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamCourier); i {
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
		file_proto_courier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierID); i {
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
		file_proto_courier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierDeleteResponse); i {
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
		file_proto_courier_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamLocation); i {
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
		file_proto_courier_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_proto_courier_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationList); i {
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
		file_proto_courier_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
	file_proto_courier_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_proto_courier_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_proto_courier_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_courier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_courier_proto_goTypes,
		DependencyIndexes: file_proto_courier_proto_depIdxs,
		MessageInfos:      file_proto_courier_proto_msgTypes,
	}.Build()
	File_proto_courier_proto = out.File
	file_proto_courier_proto_rawDesc = nil
	file_proto_courier_proto_goTypes = nil
	file_proto_courier_proto_depIdxs = nil
}
