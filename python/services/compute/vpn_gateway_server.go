// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for VpnGateway.
type VpnGatewayServer struct{}

// ProtoToVpnGatewayVpnInterface converts a VpnGatewayVpnInterface resource from its proto representation.
func ProtoToComputeVpnGatewayVpnInterface(p *computepb.ComputeVpnGatewayVpnInterface) *compute.VpnGatewayVpnInterface {
	if p == nil {
		return nil
	}
	obj := &compute.VpnGatewayVpnInterface{
		Id:        dcl.Int64OrNil(p.Id),
		IPAddress: dcl.StringOrNil(p.IpAddress),
	}
	return obj
}

// ProtoToVpnGateway converts a VpnGateway resource from its proto representation.
func ProtoToVpnGateway(p *computepb.ComputeVpnGateway) *compute.VpnGateway {
	obj := &compute.VpnGateway{
		Id:          dcl.Int64OrNil(p.Id),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Region:      dcl.StringOrNil(p.Region),
		Network:     dcl.StringOrNil(p.Network),
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetVpnInterface() {
		obj.VpnInterface = append(obj.VpnInterface, *ProtoToComputeVpnGatewayVpnInterface(r))
	}
	return obj
}

// VpnGatewayVpnInterfaceToProto converts a VpnGatewayVpnInterface resource to its proto representation.
func ComputeVpnGatewayVpnInterfaceToProto(o *compute.VpnGatewayVpnInterface) *computepb.ComputeVpnGatewayVpnInterface {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeVpnGatewayVpnInterface{
		Id:        dcl.ValueOrEmptyInt64(o.Id),
		IpAddress: dcl.ValueOrEmptyString(o.IPAddress),
	}
	return p
}

// VpnGatewayToProto converts a VpnGateway resource to its proto representation.
func VpnGatewayToProto(resource *compute.VpnGateway) *computepb.ComputeVpnGateway {
	p := &computepb.ComputeVpnGateway{
		Id:          dcl.ValueOrEmptyInt64(resource.Id),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Region:      dcl.ValueOrEmptyString(resource.Region),
		Network:     dcl.ValueOrEmptyString(resource.Network),
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.VpnInterface {
		p.VpnInterface = append(p.VpnInterface, ComputeVpnGatewayVpnInterfaceToProto(&r))
	}

	return p
}

// ApplyVpnGateway handles the gRPC request by passing it to the underlying VpnGateway Apply() method.
func (s *VpnGatewayServer) applyVpnGateway(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeVpnGatewayRequest) (*computepb.ComputeVpnGateway, error) {
	p := ProtoToVpnGateway(request.GetResource())
	res, err := c.ApplyVpnGateway(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VpnGatewayToProto(res)
	return r, nil
}

// ApplyVpnGateway handles the gRPC request by passing it to the underlying VpnGateway Apply() method.
func (s *VpnGatewayServer) ApplyComputeVpnGateway(ctx context.Context, request *computepb.ApplyComputeVpnGatewayRequest) (*computepb.ComputeVpnGateway, error) {
	cl, err := createConfigVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVpnGateway(ctx, cl, request)
}

// DeleteVpnGateway handles the gRPC request by passing it to the underlying VpnGateway Delete() method.
func (s *VpnGatewayServer) DeleteComputeVpnGateway(ctx context.Context, request *computepb.DeleteComputeVpnGatewayRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVpnGateway(ctx, ProtoToVpnGateway(request.GetResource()))

}

// ListVpnGateway handles the gRPC request by passing it to the underlying VpnGatewayList() method.
func (s *VpnGatewayServer) ListComputeVpnGateway(ctx context.Context, request *computepb.ListComputeVpnGatewayRequest) (*computepb.ListComputeVpnGatewayResponse, error) {
	cl, err := createConfigVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVpnGateway(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeVpnGateway
	for _, r := range resources.Items {
		rp := VpnGatewayToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeVpnGatewayResponse{Items: protos}, nil
}

func createConfigVpnGateway(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
