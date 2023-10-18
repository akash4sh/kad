// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: capten_plugins.proto

package captenpluginspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CaptenPlugins_AddGitProject_FullMethodName             = "/captenpluginspb.capten_plugins/AddGitProject"
	CaptenPlugins_UpdateGitProject_FullMethodName          = "/captenpluginspb.capten_plugins/UpdateGitProject"
	CaptenPlugins_DeleteGitProject_FullMethodName          = "/captenpluginspb.capten_plugins/DeleteGitProject"
	CaptenPlugins_GetGitProjects_FullMethodName            = "/captenpluginspb.capten_plugins/GetGitProjects"
	CaptenPlugins_GetGitProjectsForLabels_FullMethodName   = "/captenpluginspb.capten_plugins/GetGitProjectsForLabels"
	CaptenPlugins_AddCloudProvider_FullMethodName          = "/captenpluginspb.capten_plugins/AddCloudProvider"
	CaptenPlugins_UpdateCloudProvider_FullMethodName       = "/captenpluginspb.capten_plugins/UpdateCloudProvider"
	CaptenPlugins_DeleteCloudProvider_FullMethodName       = "/captenpluginspb.capten_plugins/DeleteCloudProvider"
	CaptenPlugins_GetCloudProviders_FullMethodName         = "/captenpluginspb.capten_plugins/GetCloudProviders"
	CaptenPlugins_GetCloudProvidersForLabel_FullMethodName = "/captenpluginspb.capten_plugins/GetCloudProvidersForLabel"
	CaptenPlugins_RegisterArgoCDProject_FullMethodName     = "/captenpluginspb.capten_plugins/RegisterArgoCDProject"
	CaptenPlugins_GetArgoCDProjects_FullMethodName         = "/captenpluginspb.capten_plugins/GetArgoCDProjects"
	CaptenPlugins_UnRegisterArgoCDProject_FullMethodName   = "/captenpluginspb.capten_plugins/UnRegisterArgoCDProject"
	CaptenPlugins_RegisterTektonProject_FullMethodName     = "/captenpluginspb.capten_plugins/RegisterTektonProject"
	CaptenPlugins_GetTektonProjects_FullMethodName         = "/captenpluginspb.capten_plugins/GetTektonProjects"
	CaptenPlugins_UnRegisterTektonProject_FullMethodName   = "/captenpluginspb.capten_plugins/UnRegisterTektonProject"
)

// CaptenPluginsClient is the client API for CaptenPlugins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptenPluginsClient interface {
	AddGitProject(ctx context.Context, in *AddGitProjectRequest, opts ...grpc.CallOption) (*AddGitProjectResponse, error)
	UpdateGitProject(ctx context.Context, in *UpdateGitProjectRequest, opts ...grpc.CallOption) (*UpdateGitProjectResponse, error)
	DeleteGitProject(ctx context.Context, in *DeleteGitProjectRequest, opts ...grpc.CallOption) (*DeleteGitProjectResponse, error)
	GetGitProjects(ctx context.Context, in *GetGitProjectsRequest, opts ...grpc.CallOption) (*GetGitProjectsResponse, error)
	GetGitProjectsForLabels(ctx context.Context, in *GetGitProjectsForLabelsRequest, opts ...grpc.CallOption) (*GetGitProjectsForLabelsResponse, error)
	AddCloudProvider(ctx context.Context, in *AddCloudProviderRequest, opts ...grpc.CallOption) (*AddCloudProviderResponse, error)
	UpdateCloudProvider(ctx context.Context, in *UpdateCloudProviderRequest, opts ...grpc.CallOption) (*UpdateCloudProviderResponse, error)
	DeleteCloudProvider(ctx context.Context, in *DeleteCloudProviderRequest, opts ...grpc.CallOption) (*DeleteCloudProviderResponse, error)
	GetCloudProviders(ctx context.Context, in *GetCloudProvidersRequest, opts ...grpc.CallOption) (*GetCloudProvidersResponse, error)
	GetCloudProvidersForLabel(ctx context.Context, in *GetCloudProvidersForLabelsRequest, opts ...grpc.CallOption) (*GetCloudProvidersForLabelsResponse, error)
	RegisterArgoCDProject(ctx context.Context, in *RegisterArgoCDProjectRequest, opts ...grpc.CallOption) (*RegisterArgoCDProjectResponse, error)
	GetArgoCDProjects(ctx context.Context, in *GetArgoCDProjectsRequest, opts ...grpc.CallOption) (*GetArgoCDProjectsResponse, error)
	UnRegisterArgoCDProject(ctx context.Context, in *UnRegisterArgoCDProjectRequest, opts ...grpc.CallOption) (*UnRegisterArgoCDProjectResponse, error)
	RegisterTektonProject(ctx context.Context, in *RegisterTektonProjectRequest, opts ...grpc.CallOption) (*RegisterTektonProjectResponse, error)
	GetTektonProjects(ctx context.Context, in *GetTektonProjectsRequest, opts ...grpc.CallOption) (*GetTektonProjectsResponse, error)
	UnRegisterTektonProject(ctx context.Context, in *UnRegisterTektonProjectRequest, opts ...grpc.CallOption) (*UnRegisterTektonProjectResponse, error)
}

type captenPluginsClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptenPluginsClient(cc grpc.ClientConnInterface) CaptenPluginsClient {
	return &captenPluginsClient{cc}
}

func (c *captenPluginsClient) AddGitProject(ctx context.Context, in *AddGitProjectRequest, opts ...grpc.CallOption) (*AddGitProjectResponse, error) {
	out := new(AddGitProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_AddGitProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) UpdateGitProject(ctx context.Context, in *UpdateGitProjectRequest, opts ...grpc.CallOption) (*UpdateGitProjectResponse, error) {
	out := new(UpdateGitProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_UpdateGitProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) DeleteGitProject(ctx context.Context, in *DeleteGitProjectRequest, opts ...grpc.CallOption) (*DeleteGitProjectResponse, error) {
	out := new(DeleteGitProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_DeleteGitProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetGitProjects(ctx context.Context, in *GetGitProjectsRequest, opts ...grpc.CallOption) (*GetGitProjectsResponse, error) {
	out := new(GetGitProjectsResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetGitProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetGitProjectsForLabels(ctx context.Context, in *GetGitProjectsForLabelsRequest, opts ...grpc.CallOption) (*GetGitProjectsForLabelsResponse, error) {
	out := new(GetGitProjectsForLabelsResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetGitProjectsForLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) AddCloudProvider(ctx context.Context, in *AddCloudProviderRequest, opts ...grpc.CallOption) (*AddCloudProviderResponse, error) {
	out := new(AddCloudProviderResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_AddCloudProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) UpdateCloudProvider(ctx context.Context, in *UpdateCloudProviderRequest, opts ...grpc.CallOption) (*UpdateCloudProviderResponse, error) {
	out := new(UpdateCloudProviderResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_UpdateCloudProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) DeleteCloudProvider(ctx context.Context, in *DeleteCloudProviderRequest, opts ...grpc.CallOption) (*DeleteCloudProviderResponse, error) {
	out := new(DeleteCloudProviderResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_DeleteCloudProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetCloudProviders(ctx context.Context, in *GetCloudProvidersRequest, opts ...grpc.CallOption) (*GetCloudProvidersResponse, error) {
	out := new(GetCloudProvidersResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetCloudProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetCloudProvidersForLabel(ctx context.Context, in *GetCloudProvidersForLabelsRequest, opts ...grpc.CallOption) (*GetCloudProvidersForLabelsResponse, error) {
	out := new(GetCloudProvidersForLabelsResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetCloudProvidersForLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) RegisterArgoCDProject(ctx context.Context, in *RegisterArgoCDProjectRequest, opts ...grpc.CallOption) (*RegisterArgoCDProjectResponse, error) {
	out := new(RegisterArgoCDProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_RegisterArgoCDProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetArgoCDProjects(ctx context.Context, in *GetArgoCDProjectsRequest, opts ...grpc.CallOption) (*GetArgoCDProjectsResponse, error) {
	out := new(GetArgoCDProjectsResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetArgoCDProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) UnRegisterArgoCDProject(ctx context.Context, in *UnRegisterArgoCDProjectRequest, opts ...grpc.CallOption) (*UnRegisterArgoCDProjectResponse, error) {
	out := new(UnRegisterArgoCDProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_UnRegisterArgoCDProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) RegisterTektonProject(ctx context.Context, in *RegisterTektonProjectRequest, opts ...grpc.CallOption) (*RegisterTektonProjectResponse, error) {
	out := new(RegisterTektonProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_RegisterTektonProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) GetTektonProjects(ctx context.Context, in *GetTektonProjectsRequest, opts ...grpc.CallOption) (*GetTektonProjectsResponse, error) {
	out := new(GetTektonProjectsResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_GetTektonProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captenPluginsClient) UnRegisterTektonProject(ctx context.Context, in *UnRegisterTektonProjectRequest, opts ...grpc.CallOption) (*UnRegisterTektonProjectResponse, error) {
	out := new(UnRegisterTektonProjectResponse)
	err := c.cc.Invoke(ctx, CaptenPlugins_UnRegisterTektonProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptenPluginsServer is the server API for CaptenPlugins service.
// All implementations must embed UnimplementedCaptenPluginsServer
// for forward compatibility
type CaptenPluginsServer interface {
	AddGitProject(context.Context, *AddGitProjectRequest) (*AddGitProjectResponse, error)
	UpdateGitProject(context.Context, *UpdateGitProjectRequest) (*UpdateGitProjectResponse, error)
	DeleteGitProject(context.Context, *DeleteGitProjectRequest) (*DeleteGitProjectResponse, error)
	GetGitProjects(context.Context, *GetGitProjectsRequest) (*GetGitProjectsResponse, error)
	GetGitProjectsForLabels(context.Context, *GetGitProjectsForLabelsRequest) (*GetGitProjectsForLabelsResponse, error)
	AddCloudProvider(context.Context, *AddCloudProviderRequest) (*AddCloudProviderResponse, error)
	UpdateCloudProvider(context.Context, *UpdateCloudProviderRequest) (*UpdateCloudProviderResponse, error)
	DeleteCloudProvider(context.Context, *DeleteCloudProviderRequest) (*DeleteCloudProviderResponse, error)
	GetCloudProviders(context.Context, *GetCloudProvidersRequest) (*GetCloudProvidersResponse, error)
	GetCloudProvidersForLabel(context.Context, *GetCloudProvidersForLabelsRequest) (*GetCloudProvidersForLabelsResponse, error)
	RegisterArgoCDProject(context.Context, *RegisterArgoCDProjectRequest) (*RegisterArgoCDProjectResponse, error)
	GetArgoCDProjects(context.Context, *GetArgoCDProjectsRequest) (*GetArgoCDProjectsResponse, error)
	UnRegisterArgoCDProject(context.Context, *UnRegisterArgoCDProjectRequest) (*UnRegisterArgoCDProjectResponse, error)
	RegisterTektonProject(context.Context, *RegisterTektonProjectRequest) (*RegisterTektonProjectResponse, error)
	GetTektonProjects(context.Context, *GetTektonProjectsRequest) (*GetTektonProjectsResponse, error)
	UnRegisterTektonProject(context.Context, *UnRegisterTektonProjectRequest) (*UnRegisterTektonProjectResponse, error)
	mustEmbedUnimplementedCaptenPluginsServer()
}

// UnimplementedCaptenPluginsServer must be embedded to have forward compatible implementations.
type UnimplementedCaptenPluginsServer struct {
}

func (UnimplementedCaptenPluginsServer) AddGitProject(context.Context, *AddGitProjectRequest) (*AddGitProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGitProject not implemented")
}
func (UnimplementedCaptenPluginsServer) UpdateGitProject(context.Context, *UpdateGitProjectRequest) (*UpdateGitProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGitProject not implemented")
}
func (UnimplementedCaptenPluginsServer) DeleteGitProject(context.Context, *DeleteGitProjectRequest) (*DeleteGitProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGitProject not implemented")
}
func (UnimplementedCaptenPluginsServer) GetGitProjects(context.Context, *GetGitProjectsRequest) (*GetGitProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGitProjects not implemented")
}
func (UnimplementedCaptenPluginsServer) GetGitProjectsForLabels(context.Context, *GetGitProjectsForLabelsRequest) (*GetGitProjectsForLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGitProjectsForLabels not implemented")
}
func (UnimplementedCaptenPluginsServer) AddCloudProvider(context.Context, *AddCloudProviderRequest) (*AddCloudProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCloudProvider not implemented")
}
func (UnimplementedCaptenPluginsServer) UpdateCloudProvider(context.Context, *UpdateCloudProviderRequest) (*UpdateCloudProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCloudProvider not implemented")
}
func (UnimplementedCaptenPluginsServer) DeleteCloudProvider(context.Context, *DeleteCloudProviderRequest) (*DeleteCloudProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCloudProvider not implemented")
}
func (UnimplementedCaptenPluginsServer) GetCloudProviders(context.Context, *GetCloudProvidersRequest) (*GetCloudProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudProviders not implemented")
}
func (UnimplementedCaptenPluginsServer) GetCloudProvidersForLabel(context.Context, *GetCloudProvidersForLabelsRequest) (*GetCloudProvidersForLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudProvidersForLabel not implemented")
}
func (UnimplementedCaptenPluginsServer) RegisterArgoCDProject(context.Context, *RegisterArgoCDProjectRequest) (*RegisterArgoCDProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterArgoCDProject not implemented")
}
func (UnimplementedCaptenPluginsServer) GetArgoCDProjects(context.Context, *GetArgoCDProjectsRequest) (*GetArgoCDProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArgoCDProjects not implemented")
}
func (UnimplementedCaptenPluginsServer) UnRegisterArgoCDProject(context.Context, *UnRegisterArgoCDProjectRequest) (*UnRegisterArgoCDProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterArgoCDProject not implemented")
}
func (UnimplementedCaptenPluginsServer) RegisterTektonProject(context.Context, *RegisterTektonProjectRequest) (*RegisterTektonProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTektonProject not implemented")
}
func (UnimplementedCaptenPluginsServer) GetTektonProjects(context.Context, *GetTektonProjectsRequest) (*GetTektonProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTektonProjects not implemented")
}
func (UnimplementedCaptenPluginsServer) UnRegisterTektonProject(context.Context, *UnRegisterTektonProjectRequest) (*UnRegisterTektonProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterTektonProject not implemented")
}
func (UnimplementedCaptenPluginsServer) mustEmbedUnimplementedCaptenPluginsServer() {}

// UnsafeCaptenPluginsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptenPluginsServer will
// result in compilation errors.
type UnsafeCaptenPluginsServer interface {
	mustEmbedUnimplementedCaptenPluginsServer()
}

func RegisterCaptenPluginsServer(s grpc.ServiceRegistrar, srv CaptenPluginsServer) {
	s.RegisterService(&CaptenPlugins_ServiceDesc, srv)
}

func _CaptenPlugins_AddGitProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGitProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).AddGitProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_AddGitProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).AddGitProject(ctx, req.(*AddGitProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_UpdateGitProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGitProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).UpdateGitProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_UpdateGitProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).UpdateGitProject(ctx, req.(*UpdateGitProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_DeleteGitProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGitProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).DeleteGitProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_DeleteGitProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).DeleteGitProject(ctx, req.(*DeleteGitProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetGitProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGitProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetGitProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetGitProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetGitProjects(ctx, req.(*GetGitProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetGitProjectsForLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGitProjectsForLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetGitProjectsForLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetGitProjectsForLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetGitProjectsForLabels(ctx, req.(*GetGitProjectsForLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_AddCloudProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCloudProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).AddCloudProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_AddCloudProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).AddCloudProvider(ctx, req.(*AddCloudProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_UpdateCloudProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).UpdateCloudProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_UpdateCloudProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).UpdateCloudProvider(ctx, req.(*UpdateCloudProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_DeleteCloudProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCloudProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).DeleteCloudProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_DeleteCloudProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).DeleteCloudProvider(ctx, req.(*DeleteCloudProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetCloudProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetCloudProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetCloudProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetCloudProviders(ctx, req.(*GetCloudProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetCloudProvidersForLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudProvidersForLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetCloudProvidersForLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetCloudProvidersForLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetCloudProvidersForLabel(ctx, req.(*GetCloudProvidersForLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_RegisterArgoCDProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterArgoCDProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).RegisterArgoCDProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_RegisterArgoCDProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).RegisterArgoCDProject(ctx, req.(*RegisterArgoCDProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetArgoCDProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArgoCDProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetArgoCDProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetArgoCDProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetArgoCDProjects(ctx, req.(*GetArgoCDProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_UnRegisterArgoCDProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterArgoCDProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).UnRegisterArgoCDProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_UnRegisterArgoCDProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).UnRegisterArgoCDProject(ctx, req.(*UnRegisterArgoCDProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_RegisterTektonProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTektonProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).RegisterTektonProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_RegisterTektonProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).RegisterTektonProject(ctx, req.(*RegisterTektonProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_GetTektonProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTektonProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).GetTektonProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_GetTektonProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).GetTektonProjects(ctx, req.(*GetTektonProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptenPlugins_UnRegisterTektonProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterTektonProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptenPluginsServer).UnRegisterTektonProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptenPlugins_UnRegisterTektonProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptenPluginsServer).UnRegisterTektonProject(ctx, req.(*UnRegisterTektonProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaptenPlugins_ServiceDesc is the grpc.ServiceDesc for CaptenPlugins service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaptenPlugins_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "captenpluginspb.capten_plugins",
	HandlerType: (*CaptenPluginsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGitProject",
			Handler:    _CaptenPlugins_AddGitProject_Handler,
		},
		{
			MethodName: "UpdateGitProject",
			Handler:    _CaptenPlugins_UpdateGitProject_Handler,
		},
		{
			MethodName: "DeleteGitProject",
			Handler:    _CaptenPlugins_DeleteGitProject_Handler,
		},
		{
			MethodName: "GetGitProjects",
			Handler:    _CaptenPlugins_GetGitProjects_Handler,
		},
		{
			MethodName: "GetGitProjectsForLabels",
			Handler:    _CaptenPlugins_GetGitProjectsForLabels_Handler,
		},
		{
			MethodName: "AddCloudProvider",
			Handler:    _CaptenPlugins_AddCloudProvider_Handler,
		},
		{
			MethodName: "UpdateCloudProvider",
			Handler:    _CaptenPlugins_UpdateCloudProvider_Handler,
		},
		{
			MethodName: "DeleteCloudProvider",
			Handler:    _CaptenPlugins_DeleteCloudProvider_Handler,
		},
		{
			MethodName: "GetCloudProviders",
			Handler:    _CaptenPlugins_GetCloudProviders_Handler,
		},
		{
			MethodName: "GetCloudProvidersForLabel",
			Handler:    _CaptenPlugins_GetCloudProvidersForLabel_Handler,
		},
		{
			MethodName: "RegisterArgoCDProject",
			Handler:    _CaptenPlugins_RegisterArgoCDProject_Handler,
		},
		{
			MethodName: "GetArgoCDProjects",
			Handler:    _CaptenPlugins_GetArgoCDProjects_Handler,
		},
		{
			MethodName: "UnRegisterArgoCDProject",
			Handler:    _CaptenPlugins_UnRegisterArgoCDProject_Handler,
		},
		{
			MethodName: "RegisterTektonProject",
			Handler:    _CaptenPlugins_RegisterTektonProject_Handler,
		},
		{
			MethodName: "GetTektonProjects",
			Handler:    _CaptenPlugins_GetTektonProjects_Handler,
		},
		{
			MethodName: "UnRegisterTektonProject",
			Handler:    _CaptenPlugins_UnRegisterTektonProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "capten_plugins.proto",
}
