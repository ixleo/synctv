// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/provider/plugin.proto

package providerpb

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

type InitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RedirectUrl  string `protobuf:"bytes,3,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *InitReq) Reset() {
	*x = InitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitReq) ProtoMessage() {}

func (x *InitReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitReq.ProtoReflect.Descriptor instead.
func (*InitReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *InitReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *InitReq) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *InitReq) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type GetTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetTokenReq) Reset() {
	*x = GetTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenReq) ProtoMessage() {}

func (x *GetTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenReq.ProtoReflect.Descriptor instead.
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Expiry       int64  `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

type RefreshTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenReq) Reset() {
	*x = RefreshTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReq) ProtoMessage() {}

func (x *RefreshTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RefreshTokenResp) Reset() {
	*x = RefreshTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResp) ProtoMessage() {}

func (x *RefreshTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResp.ProtoReflect.Descriptor instead.
func (*RefreshTokenResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type ProviderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProviderResp) Reset() {
	*x = ProviderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderResp) ProtoMessage() {}

func (x *ProviderResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderResp.ProtoReflect.Descriptor instead.
func (*ProviderResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *ProviderResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewAuthURLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *NewAuthURLReq) Reset() {
	*x = NewAuthURLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAuthURLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAuthURLReq) ProtoMessage() {}

func (x *NewAuthURLReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAuthURLReq.ProtoReflect.Descriptor instead.
func (*NewAuthURLReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *NewAuthURLReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type NewAuthURLResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *NewAuthURLResp) Reset() {
	*x = NewAuthURLResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAuthURLResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAuthURLResp) ProtoMessage() {}

func (x *NewAuthURLResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAuthURLResp.ProtoReflect.Descriptor instead.
func (*NewAuthURLResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *NewAuthURLResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserInfoReq) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ProviderUserId uint64 `protobuf:"varint,2,opt,name=provider_user_id,json=providerUserId,proto3" json:"provider_user_id,omitempty"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserInfoResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserInfoResp) GetProviderUserId() uint64 {
	if x != nil {
		return x.ProviderUserId
	}
	return 0
}

type Enpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Enpty) Reset() {
	*x = Enpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provider_plugin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enpty) ProtoMessage() {}

func (x *Enpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enpty.ProtoReflect.Descriptor instead.
func (*Enpty) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{10}
}

var File_proto_provider_plugin_proto protoreflect.FileDescriptor

var file_proto_provider_plugin_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x22, 0x36, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x4e,
	0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x32, 0xd7, 0x02, 0x0a, 0x0c, 0x4f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_provider_plugin_proto_rawDescOnce sync.Once
	file_proto_provider_plugin_proto_rawDescData = file_proto_provider_plugin_proto_rawDesc
)

func file_proto_provider_plugin_proto_rawDescGZIP() []byte {
	file_proto_provider_plugin_proto_rawDescOnce.Do(func() {
		file_proto_provider_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_provider_plugin_proto_rawDescData)
	})
	return file_proto_provider_plugin_proto_rawDescData
}

var file_proto_provider_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_provider_plugin_proto_goTypes = []interface{}{
	(*InitReq)(nil),          // 0: proto.InitReq
	(*GetTokenReq)(nil),      // 1: proto.GetTokenReq
	(*Token)(nil),            // 2: proto.Token
	(*RefreshTokenReq)(nil),  // 3: proto.RefreshTokenReq
	(*RefreshTokenResp)(nil), // 4: proto.RefreshTokenResp
	(*ProviderResp)(nil),     // 5: proto.ProviderResp
	(*NewAuthURLReq)(nil),    // 6: proto.NewAuthURLReq
	(*NewAuthURLResp)(nil),   // 7: proto.NewAuthURLResp
	(*GetUserInfoReq)(nil),   // 8: proto.GetUserInfoReq
	(*GetUserInfoResp)(nil),  // 9: proto.GetUserInfoResp
	(*Enpty)(nil),            // 10: proto.Enpty
}
var file_proto_provider_plugin_proto_depIdxs = []int32{
	2,  // 0: proto.RefreshTokenResp.token:type_name -> proto.Token
	2,  // 1: proto.GetUserInfoReq.token:type_name -> proto.Token
	0,  // 2: proto.Oauth2Plugin.Init:input_type -> proto.InitReq
	10, // 3: proto.Oauth2Plugin.Provider:input_type -> proto.Enpty
	6,  // 4: proto.Oauth2Plugin.NewAuthURL:input_type -> proto.NewAuthURLReq
	1,  // 5: proto.Oauth2Plugin.GetToken:input_type -> proto.GetTokenReq
	3,  // 6: proto.Oauth2Plugin.RefreshToken:input_type -> proto.RefreshTokenReq
	8,  // 7: proto.Oauth2Plugin.GetUserInfo:input_type -> proto.GetUserInfoReq
	10, // 8: proto.Oauth2Plugin.Init:output_type -> proto.Enpty
	5,  // 9: proto.Oauth2Plugin.Provider:output_type -> proto.ProviderResp
	7,  // 10: proto.Oauth2Plugin.NewAuthURL:output_type -> proto.NewAuthURLResp
	2,  // 11: proto.Oauth2Plugin.GetToken:output_type -> proto.Token
	4,  // 12: proto.Oauth2Plugin.RefreshToken:output_type -> proto.RefreshTokenResp
	9,  // 13: proto.Oauth2Plugin.GetUserInfo:output_type -> proto.GetUserInfoResp
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_provider_plugin_proto_init() }
func file_proto_provider_plugin_proto_init() {
	if File_proto_provider_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_provider_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitReq); i {
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
		file_proto_provider_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenReq); i {
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
		file_proto_provider_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_proto_provider_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenReq); i {
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
		file_proto_provider_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenResp); i {
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
		file_proto_provider_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderResp); i {
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
		file_proto_provider_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAuthURLReq); i {
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
		file_proto_provider_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAuthURLResp); i {
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
		file_proto_provider_plugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_proto_provider_plugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_proto_provider_plugin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enpty); i {
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
			RawDescriptor: file_proto_provider_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_provider_plugin_proto_goTypes,
		DependencyIndexes: file_proto_provider_plugin_proto_depIdxs,
		MessageInfos:      file_proto_provider_plugin_proto_msgTypes,
	}.Build()
	File_proto_provider_plugin_proto = out.File
	file_proto_provider_plugin_proto_rawDesc = nil
	file_proto_provider_plugin_proto_goTypes = nil
	file_proto_provider_plugin_proto_depIdxs = nil
}
