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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/serialization"
)

// Server implements the gRPC interface for Subnetwork.
type SubnetworkServer struct{}

// ProtoToSubnetworkPurposeEnum converts a SubnetworkPurposeEnum enum from its proto representation.
func ProtoToComputeSubnetworkPurposeEnum(e computepb.ComputeSubnetworkPurposeEnum) *compute.SubnetworkPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkPurposeEnum_name[int32(e)]; ok {
		e := compute.SubnetworkPurposeEnum(n[len("SubnetworkPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkRoleEnum converts a SubnetworkRoleEnum enum from its proto representation.
func ProtoToComputeSubnetworkRoleEnum(e computepb.ComputeSubnetworkRoleEnum) *compute.SubnetworkRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkRoleEnum_name[int32(e)]; ok {
		e := compute.SubnetworkRoleEnum(n[len("SubnetworkRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigAggregationIntervalEnum converts a SubnetworkLogConfigAggregationIntervalEnum enum from its proto representation.
func ProtoToComputeSubnetworkLogConfigAggregationIntervalEnum(e computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum) *compute.SubnetworkLogConfigAggregationIntervalEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum_name[int32(e)]; ok {
		e := compute.SubnetworkLogConfigAggregationIntervalEnum(n[len("SubnetworkLogConfigAggregationIntervalEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigMetadataEnum converts a SubnetworkLogConfigMetadataEnum enum from its proto representation.
func ProtoToComputeSubnetworkLogConfigMetadataEnum(e computepb.ComputeSubnetworkLogConfigMetadataEnum) *compute.SubnetworkLogConfigMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkLogConfigMetadataEnum_name[int32(e)]; ok {
		e := compute.SubnetworkLogConfigMetadataEnum(n[len("SubnetworkLogConfigMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkSecondaryIPRange converts a SubnetworkSecondaryIPRange resource from its proto representation.
func ProtoToComputeSubnetworkSecondaryIPRange(p *computepb.ComputeSubnetworkSecondaryIPRange) *compute.SubnetworkSecondaryIPRange {
	if p == nil {
		return nil
	}
	obj := &compute.SubnetworkSecondaryIPRange{
		RangeName:   dcl.StringOrNil(p.RangeName),
		IPCidrRange: dcl.StringOrNil(p.IpCidrRange),
	}
	return obj
}

// ProtoToSubnetworkLogConfig converts a SubnetworkLogConfig resource from its proto representation.
func ProtoToComputeSubnetworkLogConfig(p *computepb.ComputeSubnetworkLogConfig) *compute.SubnetworkLogConfig {
	if p == nil {
		return nil
	}
	obj := &compute.SubnetworkLogConfig{
		AggregationInterval: ProtoToComputeSubnetworkLogConfigAggregationIntervalEnum(p.GetAggregationInterval()),
		FlowSampling:        dcl.Float64OrNil(p.FlowSampling),
		Metadata:            ProtoToComputeSubnetworkLogConfigMetadataEnum(p.GetMetadata()),
	}
	return obj
}

// ProtoToSubnetwork converts a Subnetwork resource from its proto representation.
func ProtoToSubnetwork(p *computepb.ComputeSubnetwork) *compute.Subnetwork {
	obj := &compute.Subnetwork{
		CreationTimestamp:     dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:           dcl.StringOrNil(p.Description),
		GatewayAddress:        dcl.StringOrNil(p.GatewayAddress),
		IPCidrRange:           dcl.StringOrNil(p.IpCidrRange),
		Name:                  dcl.StringOrNil(p.Name),
		Network:               dcl.StringOrNil(p.Network),
		Fingerprint:           dcl.StringOrNil(p.Fingerprint),
		Purpose:               ProtoToComputeSubnetworkPurposeEnum(p.GetPurpose()),
		Role:                  ProtoToComputeSubnetworkRoleEnum(p.GetRole()),
		PrivateIPGoogleAccess: dcl.Bool(p.PrivateIpGoogleAccess),
		Region:                dcl.StringOrNil(p.Region),
		LogConfig:             ProtoToComputeSubnetworkLogConfig(p.GetLogConfig()),
		Project:               dcl.StringOrNil(p.Project),
		SelfLink:              dcl.StringOrNil(p.SelfLink),
		EnableFlowLogs:        dcl.Bool(p.EnableFlowLogs),
	}
	for _, r := range p.GetSecondaryIpRange() {
		obj.SecondaryIPRange = append(obj.SecondaryIPRange, *ProtoToComputeSubnetworkSecondaryIPRange(r))
	}
	return obj
}

// SubnetworkPurposeEnumToProto converts a SubnetworkPurposeEnum enum to its proto representation.
func ComputeSubnetworkPurposeEnumToProto(e *compute.SubnetworkPurposeEnum) computepb.ComputeSubnetworkPurposeEnum {
	if e == nil {
		return computepb.ComputeSubnetworkPurposeEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkPurposeEnum_value["SubnetworkPurposeEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkPurposeEnum(v)
	}
	return computepb.ComputeSubnetworkPurposeEnum(0)
}

// SubnetworkRoleEnumToProto converts a SubnetworkRoleEnum enum to its proto representation.
func ComputeSubnetworkRoleEnumToProto(e *compute.SubnetworkRoleEnum) computepb.ComputeSubnetworkRoleEnum {
	if e == nil {
		return computepb.ComputeSubnetworkRoleEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkRoleEnum_value["SubnetworkRoleEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkRoleEnum(v)
	}
	return computepb.ComputeSubnetworkRoleEnum(0)
}

// SubnetworkLogConfigAggregationIntervalEnumToProto converts a SubnetworkLogConfigAggregationIntervalEnum enum to its proto representation.
func ComputeSubnetworkLogConfigAggregationIntervalEnumToProto(e *compute.SubnetworkLogConfigAggregationIntervalEnum) computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum {
	if e == nil {
		return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum_value["SubnetworkLogConfigAggregationIntervalEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(v)
	}
	return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(0)
}

// SubnetworkLogConfigMetadataEnumToProto converts a SubnetworkLogConfigMetadataEnum enum to its proto representation.
func ComputeSubnetworkLogConfigMetadataEnumToProto(e *compute.SubnetworkLogConfigMetadataEnum) computepb.ComputeSubnetworkLogConfigMetadataEnum {
	if e == nil {
		return computepb.ComputeSubnetworkLogConfigMetadataEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkLogConfigMetadataEnum_value["SubnetworkLogConfigMetadataEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkLogConfigMetadataEnum(v)
	}
	return computepb.ComputeSubnetworkLogConfigMetadataEnum(0)
}

// SubnetworkSecondaryIPRangeToProto converts a SubnetworkSecondaryIPRange resource to its proto representation.
func ComputeSubnetworkSecondaryIPRangeToProto(o *compute.SubnetworkSecondaryIPRange) *computepb.ComputeSubnetworkSecondaryIPRange {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSubnetworkSecondaryIPRange{
		RangeName:   dcl.ValueOrEmptyString(o.RangeName),
		IpCidrRange: dcl.ValueOrEmptyString(o.IPCidrRange),
	}
	return p
}

// SubnetworkLogConfigToProto converts a SubnetworkLogConfig resource to its proto representation.
func ComputeSubnetworkLogConfigToProto(o *compute.SubnetworkLogConfig) *computepb.ComputeSubnetworkLogConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSubnetworkLogConfig{
		AggregationInterval: ComputeSubnetworkLogConfigAggregationIntervalEnumToProto(o.AggregationInterval),
		FlowSampling:        dcl.ValueOrEmptyDouble(o.FlowSampling),
		Metadata:            ComputeSubnetworkLogConfigMetadataEnumToProto(o.Metadata),
	}
	return p
}

// SubnetworkToProto converts a Subnetwork resource to its proto representation.
func SubnetworkToProto(resource *compute.Subnetwork) *computepb.ComputeSubnetwork {
	p := &computepb.ComputeSubnetwork{
		CreationTimestamp:     dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:           dcl.ValueOrEmptyString(resource.Description),
		GatewayAddress:        dcl.ValueOrEmptyString(resource.GatewayAddress),
		IpCidrRange:           dcl.ValueOrEmptyString(resource.IPCidrRange),
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Network:               dcl.ValueOrEmptyString(resource.Network),
		Fingerprint:           dcl.ValueOrEmptyString(resource.Fingerprint),
		Purpose:               ComputeSubnetworkPurposeEnumToProto(resource.Purpose),
		Role:                  ComputeSubnetworkRoleEnumToProto(resource.Role),
		PrivateIpGoogleAccess: dcl.ValueOrEmptyBool(resource.PrivateIPGoogleAccess),
		Region:                dcl.ValueOrEmptyString(resource.Region),
		LogConfig:             ComputeSubnetworkLogConfigToProto(resource.LogConfig),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		SelfLink:              dcl.ValueOrEmptyString(resource.SelfLink),
		EnableFlowLogs:        dcl.ValueOrEmptyBool(resource.EnableFlowLogs),
	}
	for _, r := range resource.SecondaryIPRange {
		p.SecondaryIpRange = append(p.SecondaryIpRange, ComputeSubnetworkSecondaryIPRangeToProto(&r))
	}

	return p
}

// ApplySubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) applySubnetwork(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeSubnetworkRequest) (*computepb.ComputeSubnetwork, error) {
	p := ProtoToSubnetwork(request.GetResource())
	res, err := c.ApplySubnetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubnetworkToProto(res)
	return r, nil
}

// ApplySubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) ApplyComputeSubnetwork(ctx context.Context, request *computepb.ApplyComputeSubnetworkRequest) (*computepb.ComputeSubnetwork, error) {
	cl, err := createConfigSubnetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySubnetwork(ctx, cl, request)
}

func (s *SubnetworkServer) ComputeSubnetworkAsHcl(ctx context.Context, request *computepb.ComputeSubnetworkAsHclRequest) (*computepb.ComputeSubnetworkAsHclResponse, error) {
	p := ProtoToSubnetwork(request.GetResource())
	resStr, err := serialization.ComputeSubnetworkAsHCL(*p)
	if err != nil {
		return nil, err
	}

	return &computepb.ComputeSubnetworkAsHclResponse{Hcl: resStr}, nil
}

// DeleteSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Delete() method.
func (s *SubnetworkServer) DeleteComputeSubnetwork(ctx context.Context, request *computepb.DeleteComputeSubnetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubnetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubnetwork(ctx, ProtoToSubnetwork(request.GetResource()))

}

// ListSubnetwork handles the gRPC request by passing it to the underlying SubnetworkList() method.
func (s *SubnetworkServer) ListComputeSubnetwork(ctx context.Context, request *computepb.ListComputeSubnetworkRequest) (*computepb.ListComputeSubnetworkResponse, error) {
	cl, err := createConfigSubnetwork(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubnetwork(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeSubnetwork
	for _, r := range resources.Items {
		rp := SubnetworkToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeSubnetworkResponse{Items: protos}, nil
}

func createConfigSubnetwork(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
