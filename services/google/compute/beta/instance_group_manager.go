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
package beta

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type InstanceGroupManager struct {
	BaseInstanceName    *string                                   `json:"baseInstanceName"`
	CreationTimestamp   *string                                   `json:"creationTimestamp"`
	DistributionPolicy  *InstanceGroupManagerDistributionPolicy   `json:"distributionPolicy"`
	CurrentActions      *InstanceGroupManagerCurrentActions       `json:"currentActions"`
	Description         *string                                   `json:"description"`
	Versions            []InstanceGroupManagerVersions            `json:"versions"`
	Id                  *int64                                    `json:"id"`
	InstanceGroup       *string                                   `json:"instanceGroup"`
	InstanceTemplate    *string                                   `json:"instanceTemplate"`
	Name                *string                                   `json:"name"`
	NamedPorts          []InstanceGroupManagerNamedPorts          `json:"namedPorts"`
	Status              *InstanceGroupManagerStatus               `json:"status"`
	TargetPools         []string                                  `json:"targetPools"`
	AutohealingPolicies []InstanceGroupManagerAutohealingPolicies `json:"autohealingPolicies"`
	UpdatePolicy        *InstanceGroupManagerUpdatePolicy         `json:"updatePolicy"`
	TargetSize          *int64                                    `json:"targetSize"`
	Zone                *string                                   `json:"zone"`
	Region              *string                                   `json:"region"`
	Project             *string                                   `json:"project"`
	Location            *string                                   `json:"location"`
}

func (r *InstanceGroupManager) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.
type InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum string

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef returns a *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef(s string) *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if s == "" {
		return nil
	}

	v := InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(s)
	return &v
}

func (v InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) Validate() error {
	for _, s := range []string{"PROACTIVE", "NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceGroupManagerUpdatePolicyMinimalActionEnum.
type InstanceGroupManagerUpdatePolicyMinimalActionEnum string

// InstanceGroupManagerUpdatePolicyMinimalActionEnumRef returns a *InstanceGroupManagerUpdatePolicyMinimalActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceGroupManagerUpdatePolicyMinimalActionEnumRef(s string) *InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceGroupManagerUpdatePolicyMinimalActionEnum(s)
	return &v
}

func (v InstanceGroupManagerUpdatePolicyMinimalActionEnum) Validate() error {
	for _, s := range []string{"RESTART", "REPLACE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceGroupManagerUpdatePolicyMinimalActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceGroupManagerDistributionPolicy struct {
	empty bool                                          `json:"-"`
	Zones []InstanceGroupManagerDistributionPolicyZones `json:"zones"`
}

// This object is used to assert a desired state where this InstanceGroupManagerDistributionPolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerDistributionPolicy *InstanceGroupManagerDistributionPolicy = &InstanceGroupManagerDistributionPolicy{empty: true}

func (r *InstanceGroupManagerDistributionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerDistributionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerDistributionPolicyZones struct {
	empty bool    `json:"-"`
	Zone  *string `json:"zone"`
}

// This object is used to assert a desired state where this InstanceGroupManagerDistributionPolicyZones is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerDistributionPolicyZones *InstanceGroupManagerDistributionPolicyZones = &InstanceGroupManagerDistributionPolicyZones{empty: true}

func (r *InstanceGroupManagerDistributionPolicyZones) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerDistributionPolicyZones) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerCurrentActions struct {
	empty                  bool   `json:"-"`
	Abandoning             *int64 `json:"abandoning"`
	Creating               *int64 `json:"creating"`
	CreatingWithoutRetries *int64 `json:"creatingWithoutRetries"`
	Deleting               *int64 `json:"deleting"`
	None                   *int64 `json:"none"`
	Recreating             *int64 `json:"recreating"`
	Refreshing             *int64 `json:"refreshing"`
	Restarting             *int64 `json:"restarting"`
}

// This object is used to assert a desired state where this InstanceGroupManagerCurrentActions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerCurrentActions *InstanceGroupManagerCurrentActions = &InstanceGroupManagerCurrentActions{empty: true}

func (r *InstanceGroupManagerCurrentActions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerCurrentActions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerVersions struct {
	empty            bool                                    `json:"-"`
	Name             *string                                 `json:"name"`
	InstanceTemplate *string                                 `json:"instanceTemplate"`
	TargetSize       *InstanceGroupManagerVersionsTargetSize `json:"targetSize"`
}

// This object is used to assert a desired state where this InstanceGroupManagerVersions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerVersions *InstanceGroupManagerVersions = &InstanceGroupManagerVersions{empty: true}

func (r *InstanceGroupManagerVersions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerVersions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerVersionsTargetSize struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

// This object is used to assert a desired state where this InstanceGroupManagerVersionsTargetSize is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerVersionsTargetSize *InstanceGroupManagerVersionsTargetSize = &InstanceGroupManagerVersionsTargetSize{empty: true}

func (r *InstanceGroupManagerVersionsTargetSize) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerVersionsTargetSize) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerNamedPorts struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Port  *int64  `json:"port"`
}

// This object is used to assert a desired state where this InstanceGroupManagerNamedPorts is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerNamedPorts *InstanceGroupManagerNamedPorts = &InstanceGroupManagerNamedPorts{empty: true}

func (r *InstanceGroupManagerNamedPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerNamedPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerStatus struct {
	empty         bool                                     `json:"-"`
	IsStable      *bool                                    `json:"isStable"`
	VersionTarget *InstanceGroupManagerStatusVersionTarget `json:"versionTarget"`
	Autoscalar    *string                                  `json:"autoscalar"`
}

// This object is used to assert a desired state where this InstanceGroupManagerStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerStatus *InstanceGroupManagerStatus = &InstanceGroupManagerStatus{empty: true}

func (r *InstanceGroupManagerStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerStatusVersionTarget struct {
	empty     bool  `json:"-"`
	IsReached *bool `json:"isReached"`
}

// This object is used to assert a desired state where this InstanceGroupManagerStatusVersionTarget is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerStatusVersionTarget *InstanceGroupManagerStatusVersionTarget = &InstanceGroupManagerStatusVersionTarget{empty: true}

func (r *InstanceGroupManagerStatusVersionTarget) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerStatusVersionTarget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerAutohealingPolicies struct {
	empty           bool    `json:"-"`
	HealthCheck     *string `json:"healthCheck"`
	InitialDelaySec *int64  `json:"initialDelaySec"`
}

// This object is used to assert a desired state where this InstanceGroupManagerAutohealingPolicies is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerAutohealingPolicies *InstanceGroupManagerAutohealingPolicies = &InstanceGroupManagerAutohealingPolicies{empty: true}

func (r *InstanceGroupManagerAutohealingPolicies) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerAutohealingPolicies) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicy struct {
	empty                      bool                                                            `json:"-"`
	InstanceRedistributionType *InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum `json:"instanceRedistributionType"`
	MinimalAction              *InstanceGroupManagerUpdatePolicyMinimalActionEnum              `json:"minimalAction"`
	MaxSurge                   *InstanceGroupManagerUpdatePolicyMaxSurge                       `json:"maxSurge"`
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicy *InstanceGroupManagerUpdatePolicy = &InstanceGroupManagerUpdatePolicy{empty: true}

func (r *InstanceGroupManagerUpdatePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicyMaxSurge struct {
	empty          bool                                                    `json:"-"`
	Fixed          *int64                                                  `json:"fixed"`
	Percent        *int64                                                  `json:"percent"`
	Calculated     *int64                                                  `json:"calculated"`
	MaxUnavailable *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable `json:"maxUnavailable"`
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicyMaxSurge is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicyMaxSurge *InstanceGroupManagerUpdatePolicyMaxSurge = &InstanceGroupManagerUpdatePolicyMaxSurge{empty: true}

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicyMaxSurge) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable struct {
	empty      bool   `json:"-"`
	Fixed      *int64 `json:"fixed"`
	Percent    *int64 `json:"percent"`
	Calculated *int64 `json:"calculated"`
}

// This object is used to assert a desired state where this InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable = &InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable{empty: true}

func (r *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *InstanceGroupManager) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "InstanceGroupManager",
		Version: "beta",
	}
}

const InstanceGroupManagerMaxPage = -1

type InstanceGroupManagerList struct {
	Items []*InstanceGroupManager

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *InstanceGroupManagerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceGroupManagerList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstanceGroupManager(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstanceGroupManager(ctx context.Context, project, location string) (*InstanceGroupManagerList, error) {

	return c.ListInstanceGroupManagerWithMaxResults(ctx, project, location, InstanceGroupManagerMaxPage)

}

func (c *Client) ListInstanceGroupManagerWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceGroupManagerList, error) {
	items, token, err := c.listInstanceGroupManager(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceGroupManagerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetInstanceGroupManager(ctx context.Context, r *InstanceGroupManager) (*InstanceGroupManager, error) {
	b, err := c.getInstanceGroupManagerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstanceGroupManager(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceGroupManagerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstanceGroupManager(ctx context.Context, r *InstanceGroupManager) error {
	if r == nil {
		return fmt.Errorf("InstanceGroupManager resource is nil")
	}
	c.Config.Logger.Info("Deleting InstanceGroupManager...")
	deleteOp := deleteInstanceGroupManagerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstanceGroupManager deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstanceGroupManager(ctx context.Context, project, location string, filter func(*InstanceGroupManager) bool) error {
	listObj, err := c.ListInstanceGroupManager(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstanceGroupManager(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstanceGroupManager(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstanceGroupManager(ctx context.Context, rawDesired *InstanceGroupManager, opts ...dcl.ApplyOption) (*InstanceGroupManager, error) {
	c.Config.Logger.Info("Beginning ApplyInstanceGroupManager...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.instanceGroupManagerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []instanceGroupManagerApiOperation
	if create {
		ops = append(ops, &createInstanceGroupManagerOperation{})
	} else if recreate {

		ops = append(ops, &deleteInstanceGroupManagerOperation{})

		ops = append(ops, &createInstanceGroupManagerOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceGroupManagerDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetInstanceGroupManager(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceGroupManagerNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceGroupManagerDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstanceGroupManager(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
