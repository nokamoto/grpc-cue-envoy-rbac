// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/rbac.proto

package api

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *RoleBinding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *RoleBinding) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type CreateRoleBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleBinding *RoleBinding `protobuf:"bytes,1,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty"`
}

func (x *CreateRoleBindingRequest) Reset() {
	*x = CreateRoleBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleBindingRequest) ProtoMessage() {}

func (x *CreateRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoleBindingRequest) GetRoleBinding() *RoleBinding {
	if x != nil {
		return x.RoleBinding
	}
	return nil
}

type AuthorizeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          string         `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Authorization *Authorization `protobuf:"bytes,2,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *AuthorizeUserRequest) Reset() {
	*x = AuthorizeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserRequest) ProtoMessage() {}

func (x *AuthorizeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeUserRequest) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizeUserRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthorizeUserRequest) GetAuthorization() *Authorization {
	if x != nil {
		return x.Authorization
	}
	return nil
}

type AuthorizeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorizeUserResponse) Reset() {
	*x = AuthorizeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rbac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserResponse) ProtoMessage() {}

func (x *AuthorizeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rbac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeUserResponse) Descriptor() ([]byte, []int) {
	return file_api_rbac_proto_rawDescGZIP(), []int{5}
}

var File_api_rbac_proto protoreflect.FileDescriptor

var file_api_rbac_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x3e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x5b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x70, 0x0a, 0x14,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17,
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcb, 0x02, 0x0a, 0x0b, 0x52, 0x42, 0x41, 0x43,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x6f, 0x6b, 0x61,
	0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x22, 0x16, 0x82, 0xb5, 0x18, 0x12, 0x0a, 0x10, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x7b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e,
	0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x1d, 0x82, 0xb5, 0x18, 0x19, 0x0a, 0x17, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x63, 0x75, 0x65, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x72, 0x62, 0x61, 0x63,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rbac_proto_rawDescOnce sync.Once
	file_api_rbac_proto_rawDescData = file_api_rbac_proto_rawDesc
)

func file_api_rbac_proto_rawDescGZIP() []byte {
	file_api_rbac_proto_rawDescOnce.Do(func() {
		file_api_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rbac_proto_rawDescData)
	})
	return file_api_rbac_proto_rawDescData
}

var file_api_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_rbac_proto_goTypes = []interface{}{
	(*Role)(nil),                     // 0: nokamoto.github.Role
	(*RoleBinding)(nil),              // 1: nokamoto.github.RoleBinding
	(*CreateRoleRequest)(nil),        // 2: nokamoto.github.CreateRoleRequest
	(*CreateRoleBindingRequest)(nil), // 3: nokamoto.github.CreateRoleBindingRequest
	(*AuthorizeUserRequest)(nil),     // 4: nokamoto.github.AuthorizeUserRequest
	(*AuthorizeUserResponse)(nil),    // 5: nokamoto.github.AuthorizeUserResponse
	(*Authorization)(nil),            // 6: nokamoto.github.Authorization
}
var file_api_rbac_proto_depIdxs = []int32{
	0, // 0: nokamoto.github.CreateRoleRequest.role:type_name -> nokamoto.github.Role
	1, // 1: nokamoto.github.CreateRoleBindingRequest.role_binding:type_name -> nokamoto.github.RoleBinding
	6, // 2: nokamoto.github.AuthorizeUserRequest.authorization:type_name -> nokamoto.github.Authorization
	2, // 3: nokamoto.github.RBACService.CreateRole:input_type -> nokamoto.github.CreateRoleRequest
	3, // 4: nokamoto.github.RBACService.CreateRoleBinding:input_type -> nokamoto.github.CreateRoleBindingRequest
	4, // 5: nokamoto.github.RBACService.AuthorizeUser:input_type -> nokamoto.github.AuthorizeUserRequest
	0, // 6: nokamoto.github.RBACService.CreateRole:output_type -> nokamoto.github.Role
	1, // 7: nokamoto.github.RBACService.CreateRoleBinding:output_type -> nokamoto.github.RoleBinding
	5, // 8: nokamoto.github.RBACService.AuthorizeUser:output_type -> nokamoto.github.AuthorizeUserResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_rbac_proto_init() }
func file_api_rbac_proto_init() {
	if File_api_rbac_proto != nil {
		return
	}
	file_api_authorization_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_api_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_api_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_api_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleBindingRequest); i {
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
		file_api_rbac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserRequest); i {
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
		file_api_rbac_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserResponse); i {
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
			RawDescriptor: file_api_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rbac_proto_goTypes,
		DependencyIndexes: file_api_rbac_proto_depIdxs,
		MessageInfos:      file_api_rbac_proto_msgTypes,
	}.Build()
	File_api_rbac_proto = out.File
	file_api_rbac_proto_rawDesc = nil
	file_api_rbac_proto_goTypes = nil
	file_api_rbac_proto_depIdxs = nil
}
