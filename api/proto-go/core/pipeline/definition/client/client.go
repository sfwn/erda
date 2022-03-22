// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: definition.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/definition/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// DefinitionService definition.proto
	DefinitionService() pb.DefinitionServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		definitionService: pb.NewDefinitionServiceClient(cc),
	}
}

type serviceClients struct {
	definitionService pb.DefinitionServiceClient
}

func (c *serviceClients) DefinitionService() pb.DefinitionServiceClient {
	return c.definitionService
}

type definitionServiceWrapper struct {
	client pb.DefinitionServiceClient
	opts   []grpc1.CallOption
}

func (s *definitionServiceWrapper) Create(ctx context.Context, req *pb.PipelineDefinitionCreateRequest) (*pb.PipelineDefinitionCreateResponse, error) {
	return s.client.Create(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) Update(ctx context.Context, req *pb.PipelineDefinitionUpdateRequest) (*pb.PipelineDefinitionUpdateResponse, error) {
	return s.client.Update(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) Delete(ctx context.Context, req *pb.PipelineDefinitionDeleteRequest) (*pb.PipelineDefinitionDeleteResponse, error) {
	return s.client.Delete(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) Get(ctx context.Context, req *pb.PipelineDefinitionGetRequest) (*pb.PipelineDefinitionGetResponse, error) {
	return s.client.Get(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) List(ctx context.Context, req *pb.PipelineDefinitionListRequest) (*pb.PipelineDefinitionListResponse, error) {
	return s.client.List(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) StatisticsGroupByRemote(ctx context.Context, req *pb.PipelineDefinitionStatisticsRequest) (*pb.PipelineDefinitionStatisticsResponse, error) {
	return s.client.StatisticsGroupByRemote(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) ListUsedRefs(ctx context.Context, req *pb.PipelineDefinitionUsedRefListRequest) (*pb.PipelineDefinitionUsedRefListResponse, error) {
	return s.client.ListUsedRefs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) StatisticsGroupByFilePath(ctx context.Context, req *pb.PipelineDefinitionStatisticsRequest) (*pb.PipelineDefinitionStatisticsResponse, error) {
	return s.client.StatisticsGroupByFilePath(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
