// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agentpb

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	StoreCred(ctx context.Context, in *StoreCredRequest, opts ...grpc.CallOption) (*StoreCredResponse, error)
	StoreSecret(ctx context.Context, in *StoreSecretRequest, opts ...grpc.CallOption) (*StoreSecretResponse, error)
	ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StoreCred(ctx context.Context, in *StoreCredRequest, opts ...grpc.CallOption) (*StoreCredResponse, error) {
	out := new(StoreCredResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/StoreCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StoreSecret(ctx context.Context, in *StoreSecretRequest, opts ...grpc.CallOption) (*StoreSecretResponse, error) {
	out := new(StoreSecretResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/StoreSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClimonAppInstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClimonAppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/DeployerAppInstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/DeployerAppDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClusterAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ClusterDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/RepositoryAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/RepositoryDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ProjectAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, "/agentpb.Agent/ProjectDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	SubmitJob(context.Context, *JobRequest) (*JobResponse, error)
	StoreCred(context.Context, *StoreCredRequest) (*StoreCredResponse, error)
	StoreSecret(context.Context, *StoreSecretRequest) (*StoreSecretResponse, error)
	ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error)
	ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error)
	DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error)
	DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error)
	ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error)
	ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error)
	RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error)
	RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error)
	ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error)
	ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) SubmitJob(context.Context, *JobRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedAgentServer) StoreCred(context.Context, *StoreCredRequest) (*StoreCredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCred not implemented")
}
func (UnimplementedAgentServer) StoreSecret(context.Context, *StoreSecretRequest) (*StoreSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreSecret not implemented")
}
func (UnimplementedAgentServer) ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppInstall not implemented")
}
func (UnimplementedAgentServer) ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppDelete not implemented")
}
func (UnimplementedAgentServer) DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppInstall not implemented")
}
func (UnimplementedAgentServer) DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppDelete not implemented")
}
func (UnimplementedAgentServer) ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterAdd not implemented")
}
func (UnimplementedAgentServer) ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterDelete not implemented")
}
func (UnimplementedAgentServer) RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryAdd not implemented")
}
func (UnimplementedAgentServer) RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryDelete not implemented")
}
func (UnimplementedAgentServer) ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectAdd not implemented")
}
func (UnimplementedAgentServer) ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectDelete not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SubmitJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StoreCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StoreCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/StoreCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StoreCred(ctx, req.(*StoreCredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StoreSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StoreSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/StoreSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StoreSecret(ctx, req.(*StoreSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClimonAppInstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppInstall(ctx, req.(*ClimonInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClimonAppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppDelete(ctx, req.(*ClimonDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/DeployerAppInstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppInstall(ctx, req.(*ApplicationInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/DeployerAppDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppDelete(ctx, req.(*ApplicationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClusterAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterAdd(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ClusterDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterDelete(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/RepositoryAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryAdd(ctx, req.(*RepositoryAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/RepositoryDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryDelete(ctx, req.(*RepositoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ProjectAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectAdd(ctx, req.(*ProjectAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Agent/ProjectDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectDelete(ctx, req.(*ProjectDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _Agent_SubmitJob_Handler,
		},
		{
			MethodName: "StoreCred",
			Handler:    _Agent_StoreCred_Handler,
		},
		{
			MethodName: "StoreSecret",
			Handler:    _Agent_StoreSecret_Handler,
		},
		{
			MethodName: "ClimonAppInstall",
			Handler:    _Agent_ClimonAppInstall_Handler,
		},
		{
			MethodName: "ClimonAppDelete",
			Handler:    _Agent_ClimonAppDelete_Handler,
		},
		{
			MethodName: "DeployerAppInstall",
			Handler:    _Agent_DeployerAppInstall_Handler,
		},
		{
			MethodName: "DeployerAppDelete",
			Handler:    _Agent_DeployerAppDelete_Handler,
		},
		{
			MethodName: "ClusterAdd",
			Handler:    _Agent_ClusterAdd_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Agent_ClusterDelete_Handler,
		},
		{
			MethodName: "RepositoryAdd",
			Handler:    _Agent_RepositoryAdd_Handler,
		},
		{
			MethodName: "RepositoryDelete",
			Handler:    _Agent_RepositoryDelete_Handler,
		},
		{
			MethodName: "ProjectAdd",
			Handler:    _Agent_ProjectAdd_Handler,
		},
		{
			MethodName: "ProjectDelete",
			Handler:    _Agent_ProjectDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
