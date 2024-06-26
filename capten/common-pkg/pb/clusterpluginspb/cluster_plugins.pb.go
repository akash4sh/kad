// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: cluster_plugins.proto

package clusterpluginspb

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

type StatusCode int32

const (
	StatusCode_OK               StatusCode = 0
	StatusCode_INTERNRAL_ERROR  StatusCode = 1
	StatusCode_INVALID_ARGUMENT StatusCode = 2
	StatusCode_NOT_FOUND        StatusCode = 3
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "OK",
		1: "INTERNRAL_ERROR",
		2: "INVALID_ARGUMENT",
		3: "NOT_FOUND",
	}
	StatusCode_value = map[string]int32{
		"OK":               0,
		"INTERNRAL_ERROR":  1,
		"INVALID_ARGUMENT": 2,
		"NOT_FOUND":        3,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_plugins_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_cluster_plugins_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{0}
}

type StoreType int32

const (
	StoreType_CENTRAL_CAPTEN_STORE StoreType = 0
	StoreType_LOCAL_CAPTEN_STORE   StoreType = 1
)

// Enum value maps for StoreType.
var (
	StoreType_name = map[int32]string{
		0: "CENTRAL_CAPTEN_STORE",
		1: "LOCAL_CAPTEN_STORE",
	}
	StoreType_value = map[string]int32{
		"CENTRAL_CAPTEN_STORE": 0,
		"LOCAL_CAPTEN_STORE":   1,
	}
)

func (x StoreType) Enum() *StoreType {
	p := new(StoreType)
	*p = x
	return p
}

func (x StoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_plugins_proto_enumTypes[1].Descriptor()
}

func (StoreType) Type() protoreflect.EnumType {
	return &file_cluster_plugins_proto_enumTypes[1]
}

func (x StoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreType.Descriptor instead.
func (StoreType) EnumDescriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{1}
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreType           StoreType `protobuf:"varint,1,opt,name=storeType,proto3,enum=clusterpluginspb.StoreType" json:"storeType,omitempty"`
	PluginName          string    `protobuf:"bytes,2,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
	Description         string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category            string    `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Version             string    `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Icon                []byte    `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	ChartName           string    `protobuf:"bytes,7,opt,name=chartName,proto3" json:"chartName,omitempty"`
	ChartRepo           string    `protobuf:"bytes,8,opt,name=chartRepo,proto3" json:"chartRepo,omitempty"`
	DefaultNamespace    string    `protobuf:"bytes,9,opt,name=defaultNamespace,proto3" json:"defaultNamespace,omitempty"`
	PrivilegedNamespace bool      `protobuf:"varint,10,opt,name=privilegedNamespace,proto3" json:"privilegedNamespace,omitempty"`
	ApiEndpoint         string    `protobuf:"bytes,11,opt,name=apiEndpoint,proto3" json:"apiEndpoint,omitempty"`
	UiEndpoint          string    `protobuf:"bytes,12,opt,name=uiEndpoint,proto3" json:"uiEndpoint,omitempty"`
	UiModuleEndpoint    string    `protobuf:"bytes,13,opt,name=uiModuleEndpoint,proto3" json:"uiModuleEndpoint,omitempty"`
	Capabilities        []string  `protobuf:"bytes,14,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	Values              []byte    `protobuf:"bytes,15,opt,name=values,proto3" json:"values,omitempty"`
	OverrideValues      []byte    `protobuf:"bytes,16,opt,name=overrideValues,proto3" json:"overrideValues,omitempty"`
	InstallStatus       string    `protobuf:"bytes,17,opt,name=installStatus,proto3" json:"installStatus,omitempty"`
	LastUpdateTime      string    `protobuf:"bytes,18,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_CENTRAL_CAPTEN_STORE
}

func (x *Plugin) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *Plugin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plugin) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *Plugin) GetChartName() string {
	if x != nil {
		return x.ChartName
	}
	return ""
}

func (x *Plugin) GetChartRepo() string {
	if x != nil {
		return x.ChartRepo
	}
	return ""
}

func (x *Plugin) GetDefaultNamespace() string {
	if x != nil {
		return x.DefaultNamespace
	}
	return ""
}

func (x *Plugin) GetPrivilegedNamespace() bool {
	if x != nil {
		return x.PrivilegedNamespace
	}
	return false
}

func (x *Plugin) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *Plugin) GetUiEndpoint() string {
	if x != nil {
		return x.UiEndpoint
	}
	return ""
}

func (x *Plugin) GetUiModuleEndpoint() string {
	if x != nil {
		return x.UiModuleEndpoint
	}
	return ""
}

func (x *Plugin) GetCapabilities() []string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *Plugin) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Plugin) GetOverrideValues() []byte {
	if x != nil {
		return x.OverrideValues
	}
	return nil
}

func (x *Plugin) GetInstallStatus() string {
	if x != nil {
		return x.InstallStatus
	}
	return ""
}

func (x *Plugin) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

type DeployClusterPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *Plugin `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *DeployClusterPluginRequest) Reset() {
	*x = DeployClusterPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployClusterPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployClusterPluginRequest) ProtoMessage() {}

func (x *DeployClusterPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployClusterPluginRequest.ProtoReflect.Descriptor instead.
func (*DeployClusterPluginRequest) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{1}
}

func (x *DeployClusterPluginRequest) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type DeployClusterPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=clusterpluginspb.StatusCode" json:"status,omitempty"`
	StatusMessage string     `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *DeployClusterPluginResponse) Reset() {
	*x = DeployClusterPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployClusterPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployClusterPluginResponse) ProtoMessage() {}

func (x *DeployClusterPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployClusterPluginResponse.ProtoReflect.Descriptor instead.
func (*DeployClusterPluginResponse) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{2}
}

func (x *DeployClusterPluginResponse) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

func (x *DeployClusterPluginResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type UnDeployClusterPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreType  StoreType `protobuf:"varint,1,opt,name=storeType,proto3,enum=clusterpluginspb.StoreType" json:"storeType,omitempty"`
	PluginName string    `protobuf:"bytes,2,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
}

func (x *UnDeployClusterPluginRequest) Reset() {
	*x = UnDeployClusterPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnDeployClusterPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnDeployClusterPluginRequest) ProtoMessage() {}

func (x *UnDeployClusterPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnDeployClusterPluginRequest.ProtoReflect.Descriptor instead.
func (*UnDeployClusterPluginRequest) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{3}
}

func (x *UnDeployClusterPluginRequest) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_CENTRAL_CAPTEN_STORE
}

func (x *UnDeployClusterPluginRequest) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

type UnDeployClusterPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=clusterpluginspb.StatusCode" json:"status,omitempty"`
	StatusMessage string     `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *UnDeployClusterPluginResponse) Reset() {
	*x = UnDeployClusterPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnDeployClusterPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnDeployClusterPluginResponse) ProtoMessage() {}

func (x *UnDeployClusterPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnDeployClusterPluginResponse.ProtoReflect.Descriptor instead.
func (*UnDeployClusterPluginResponse) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{4}
}

func (x *UnDeployClusterPluginResponse) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

func (x *UnDeployClusterPluginResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type ClusterPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreType     StoreType `protobuf:"varint,1,opt,name=storeType,proto3,enum=clusterpluginspb.StoreType" json:"storeType,omitempty"`
	PluginName    string    `protobuf:"bytes,2,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
	Description   string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category      string    `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Version       string    `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Icon          []byte    `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	InstallStatus string    `protobuf:"bytes,7,opt,name=installStatus,proto3" json:"installStatus,omitempty"`
}

func (x *ClusterPlugin) Reset() {
	*x = ClusterPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterPlugin) ProtoMessage() {}

func (x *ClusterPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterPlugin.ProtoReflect.Descriptor instead.
func (*ClusterPlugin) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterPlugin) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_CENTRAL_CAPTEN_STORE
}

func (x *ClusterPlugin) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *ClusterPlugin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ClusterPlugin) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ClusterPlugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ClusterPlugin) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *ClusterPlugin) GetInstallStatus() string {
	if x != nil {
		return x.InstallStatus
	}
	return ""
}

type GetClusterPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterPluginsRequest) Reset() {
	*x = GetClusterPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterPluginsRequest) ProtoMessage() {}

func (x *GetClusterPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterPluginsRequest.ProtoReflect.Descriptor instead.
func (*GetClusterPluginsRequest) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{6}
}

type GetClusterPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        StatusCode       `protobuf:"varint,1,opt,name=status,proto3,enum=clusterpluginspb.StatusCode" json:"status,omitempty"`
	StatusMessage string           `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Plugins       []*ClusterPlugin `protobuf:"bytes,3,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *GetClusterPluginsResponse) Reset() {
	*x = GetClusterPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_plugins_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterPluginsResponse) ProtoMessage() {}

func (x *GetClusterPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_plugins_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterPluginsResponse.ProtoReflect.Descriptor instead.
func (*GetClusterPluginsResponse) Descriptor() ([]byte, []int) {
	return file_cluster_plugins_proto_rawDescGZIP(), []int{7}
}

func (x *GetClusterPluginsResponse) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

func (x *GetClusterPluginsResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *GetClusterPluginsResponse) GetPlugins() []*ClusterPlugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

var File_cluster_plugins_proto protoreflect.FileDescriptor

var file_cluster_plugins_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x22, 0x89, 0x05, 0x0a, 0x06, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x69, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x69, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x69, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x75, 0x69, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x1a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x79, 0x0a, 0x1b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x79, 0x0a, 0x1c, 0x55, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x1d, 0x55,
	0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfc, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70,
	0x62, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2a, 0x4e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x52, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x2a, 0x3d, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c,
	0x5f, 0x43, 0x41, 0x50, 0x54, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x45, 0x4e, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x32, 0xf2, 0x02, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x74, 0x0a, 0x13, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x12, 0x2c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x7a, 0x0a, 0x15, 0x55, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_plugins_proto_rawDescOnce sync.Once
	file_cluster_plugins_proto_rawDescData = file_cluster_plugins_proto_rawDesc
)

func file_cluster_plugins_proto_rawDescGZIP() []byte {
	file_cluster_plugins_proto_rawDescOnce.Do(func() {
		file_cluster_plugins_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_plugins_proto_rawDescData)
	})
	return file_cluster_plugins_proto_rawDescData
}

var file_cluster_plugins_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cluster_plugins_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cluster_plugins_proto_goTypes = []interface{}{
	(StatusCode)(0),                       // 0: clusterpluginspb.StatusCode
	(StoreType)(0),                        // 1: clusterpluginspb.StoreType
	(*Plugin)(nil),                        // 2: clusterpluginspb.Plugin
	(*DeployClusterPluginRequest)(nil),    // 3: clusterpluginspb.DeployClusterPluginRequest
	(*DeployClusterPluginResponse)(nil),   // 4: clusterpluginspb.DeployClusterPluginResponse
	(*UnDeployClusterPluginRequest)(nil),  // 5: clusterpluginspb.UnDeployClusterPluginRequest
	(*UnDeployClusterPluginResponse)(nil), // 6: clusterpluginspb.UnDeployClusterPluginResponse
	(*ClusterPlugin)(nil),                 // 7: clusterpluginspb.ClusterPlugin
	(*GetClusterPluginsRequest)(nil),      // 8: clusterpluginspb.GetClusterPluginsRequest
	(*GetClusterPluginsResponse)(nil),     // 9: clusterpluginspb.GetClusterPluginsResponse
}
var file_cluster_plugins_proto_depIdxs = []int32{
	1,  // 0: clusterpluginspb.Plugin.storeType:type_name -> clusterpluginspb.StoreType
	2,  // 1: clusterpluginspb.DeployClusterPluginRequest.plugin:type_name -> clusterpluginspb.Plugin
	0,  // 2: clusterpluginspb.DeployClusterPluginResponse.status:type_name -> clusterpluginspb.StatusCode
	1,  // 3: clusterpluginspb.UnDeployClusterPluginRequest.storeType:type_name -> clusterpluginspb.StoreType
	0,  // 4: clusterpluginspb.UnDeployClusterPluginResponse.status:type_name -> clusterpluginspb.StatusCode
	1,  // 5: clusterpluginspb.ClusterPlugin.storeType:type_name -> clusterpluginspb.StoreType
	0,  // 6: clusterpluginspb.GetClusterPluginsResponse.status:type_name -> clusterpluginspb.StatusCode
	7,  // 7: clusterpluginspb.GetClusterPluginsResponse.plugins:type_name -> clusterpluginspb.ClusterPlugin
	3,  // 8: clusterpluginspb.ClusterPlugins.DeployClusterPlugin:input_type -> clusterpluginspb.DeployClusterPluginRequest
	5,  // 9: clusterpluginspb.ClusterPlugins.UnDeployClusterPlugin:input_type -> clusterpluginspb.UnDeployClusterPluginRequest
	8,  // 10: clusterpluginspb.ClusterPlugins.GetClusterPlugins:input_type -> clusterpluginspb.GetClusterPluginsRequest
	4,  // 11: clusterpluginspb.ClusterPlugins.DeployClusterPlugin:output_type -> clusterpluginspb.DeployClusterPluginResponse
	6,  // 12: clusterpluginspb.ClusterPlugins.UnDeployClusterPlugin:output_type -> clusterpluginspb.UnDeployClusterPluginResponse
	9,  // 13: clusterpluginspb.ClusterPlugins.GetClusterPlugins:output_type -> clusterpluginspb.GetClusterPluginsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cluster_plugins_proto_init() }
func file_cluster_plugins_proto_init() {
	if File_cluster_plugins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_plugins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_cluster_plugins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployClusterPluginRequest); i {
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
		file_cluster_plugins_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployClusterPluginResponse); i {
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
		file_cluster_plugins_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnDeployClusterPluginRequest); i {
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
		file_cluster_plugins_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnDeployClusterPluginResponse); i {
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
		file_cluster_plugins_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterPlugin); i {
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
		file_cluster_plugins_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterPluginsRequest); i {
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
		file_cluster_plugins_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterPluginsResponse); i {
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
			RawDescriptor: file_cluster_plugins_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_plugins_proto_goTypes,
		DependencyIndexes: file_cluster_plugins_proto_depIdxs,
		EnumInfos:         file_cluster_plugins_proto_enumTypes,
		MessageInfos:      file_cluster_plugins_proto_msgTypes,
	}.Build()
	File_cluster_plugins_proto = out.File
	file_cluster_plugins_proto_rawDesc = nil
	file_cluster_plugins_proto_goTypes = nil
	file_cluster_plugins_proto_depIdxs = nil
}
