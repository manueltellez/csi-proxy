// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/api/disk/v1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server/disk/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1.RegisterDiskServer(grpcServer, s)
}

func (s *versionedAPI) DiskStats(context context.Context, versionedRequest *v1.DiskStatsRequest) (*v1.DiskStatsResponse, error) {
	request := &internal.DiskStatsRequest{}
	if err := Convert_v1_DiskStatsRequest_To_internal_DiskStatsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.DiskStats(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.DiskStatsResponse{}
	if err := Convert_internal_DiskStatsResponse_To_v1_DiskStatsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) GetAttachState(context context.Context, versionedRequest *v1.GetAttachStateRequest) (*v1.GetAttachStateResponse, error) {
	request := &internal.GetAttachStateRequest{}
	if err := Convert_v1_GetAttachStateRequest_To_internal_GetAttachStateRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.GetAttachState(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.GetAttachStateResponse{}
	if err := Convert_internal_GetAttachStateResponse_To_v1_GetAttachStateResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) ListDiskIDs(context context.Context, versionedRequest *v1.ListDiskIDsRequest) (*v1.ListDiskIDsResponse, error) {
	request := &internal.ListDiskIDsRequest{}
	if err := Convert_v1_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ListDiskIDs(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.ListDiskIDsResponse{}
	if err := Convert_internal_ListDiskIDsResponse_To_v1_ListDiskIDsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) ListDiskLocations(context context.Context, versionedRequest *v1.ListDiskLocationsRequest) (*v1.ListDiskLocationsResponse, error) {
	request := &internal.ListDiskLocationsRequest{}
	if err := Convert_v1_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ListDiskLocations(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.ListDiskLocationsResponse{}
	if err := Convert_internal_ListDiskLocationsResponse_To_v1_ListDiskLocationsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) PartitionDisk(context context.Context, versionedRequest *v1.PartitionDiskRequest) (*v1.PartitionDiskResponse, error) {
	request := &internal.PartitionDiskRequest{}
	if err := Convert_v1_PartitionDiskRequest_To_internal_PartitionDiskRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.PartitionDisk(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.PartitionDiskResponse{}
	if err := Convert_internal_PartitionDiskResponse_To_v1_PartitionDiskResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) Rescan(context context.Context, versionedRequest *v1.RescanRequest) (*v1.RescanResponse, error) {
	request := &internal.RescanRequest{}
	if err := Convert_v1_RescanRequest_To_internal_RescanRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Rescan(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.RescanResponse{}
	if err := Convert_internal_RescanResponse_To_v1_RescanResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) SetAttachState(context context.Context, versionedRequest *v1.SetAttachStateRequest) (*v1.SetAttachStateResponse, error) {
	request := &internal.SetAttachStateRequest{}
	if err := Convert_v1_SetAttachStateRequest_To_internal_SetAttachStateRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.SetAttachState(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.SetAttachStateResponse{}
	if err := Convert_internal_SetAttachStateResponse_To_v1_SetAttachStateResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
