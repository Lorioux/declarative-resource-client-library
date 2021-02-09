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
package monitoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *UptimeCheckConfig) validate() error {

	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "timeout"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MonitoredResource) {
		if err := r.MonitoredResource.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ResourceGroup) {
		if err := r.ResourceGroup.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpCheck) {
		if err := r.HttpCheck.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TcpCheck) {
		if err := r.TcpCheck.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UptimeCheckConfigMonitoredResource) validate() error {
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "filterLabels"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigResourceGroup) validate() error {
	return nil
}
func (r *UptimeCheckConfigHttpCheck) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AuthInfo) {
		if err := r.AuthInfo.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *UptimeCheckConfigHttpCheckAuthInfo) validate() error {
	if err := dcl.Required(r, "username"); err != nil {
		return err
	}
	if err := dcl.Required(r, "password"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigTcpCheck) validate() error {
	if err := dcl.Required(r, "port"); err != nil {
		return err
	}
	return nil
}
func (r *UptimeCheckConfigContentMatchers) validate() error {
	if err := dcl.Required(r, "content"); err != nil {
		return err
	}
	return nil
}

func uptimeCheckConfigGetURL(userBasePath string, r *UptimeCheckConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

func uptimeCheckConfigListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/uptimeCheckConfigs", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func uptimeCheckConfigCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("v3/projects/{{project}}/uptimeCheckConfigs", "https://monitoring.googleapis.com/", userBasePath, params), nil

}

func uptimeCheckConfigDeleteURL(userBasePath string, r *UptimeCheckConfig) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("v3/projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/", userBasePath, params), nil
}

// uptimeCheckConfigApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type uptimeCheckConfigApiOperation interface {
	do(context.Context, *UptimeCheckConfig, *Client) error
}

// newUpdateUptimeCheckConfigUpdateRequest creates a request for an
// UptimeCheckConfig resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateUptimeCheckConfigUpdateRequest(ctx context.Context, f *UptimeCheckConfig, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheck(c, f.HttpCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpCheck into httpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["httpCheck"] = v
	}
	if v, err := expandUptimeCheckConfigTcpCheck(c, f.TcpCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpCheck into tcpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["tcpCheck"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		req["timeout"] = v
	}
	if v, err := expandUptimeCheckConfigContentMatchersSlice(c, f.ContentMatchers); err != nil {
		return nil, fmt.Errorf("error expanding ContentMatchers into contentMatchers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["contentMatchers"] = v
	}
	if v := f.PrivateCheckers; !dcl.IsEmptyValueIndirect(v) {
		req["privateCheckers"] = v
	}
	if v := f.SelectedRegions; !dcl.IsEmptyValueIndirect(v) {
		req["selectedRegions"] = v
	}
	req["name"] = fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", *f.Project, *f.Name)

	return req, nil
}

// marshalUpdateUptimeCheckConfigUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateUptimeCheckConfigUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateUptimeCheckConfigUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateUptimeCheckConfigUpdateOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {
	_, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"displayName", "httpCheck", "tcpCheck", "timeout", "contentMatchers", "privateCheckers", "selectedRegions"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateUptimeCheckConfigUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateUptimeCheckConfigUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.Retry)
	if err != nil {
		return err
	}

	var o operations.MonitoringOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listUptimeCheckConfigRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := uptimeCheckConfigListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != UptimeCheckConfigMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listUptimeCheckConfigOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listUptimeCheckConfig(ctx context.Context, project, pageToken string, pageSize int32) ([]*UptimeCheckConfig, string, error) {
	b, err := c.listUptimeCheckConfigRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listUptimeCheckConfigOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*UptimeCheckConfig
	for _, v := range m.Items {
		res := flattenUptimeCheckConfig(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllUptimeCheckConfig(ctx context.Context, f func(*UptimeCheckConfig) bool, resources []*UptimeCheckConfig) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteUptimeCheckConfig(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteUptimeCheckConfigOperation struct{}

func (op *deleteUptimeCheckConfigOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {

	_, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("UptimeCheckConfig not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetUptimeCheckConfig checking for existence. error: %v", err)
		return err
	}

	u, err := uptimeCheckConfigDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.Retry)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.MonitoringOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET"); err != nil {
		return err
	}
	_, err = c.GetUptimeCheckConfig(ctx, r.urlNormalized())
	if !dcl.IsNotFound(err) {
		return dcl.NotDeletedError{ExistingResource: r}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createUptimeCheckConfigOperation struct{}

func (op *createUptimeCheckConfigOperation) do(ctx context.Context, r *UptimeCheckConfig, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := uptimeCheckConfigCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.Retry)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.MonitoringOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://monitoring.googleapis.com/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	r.Name, err = o.FetchName()
	if err != nil {
		return fmt.Errorf("error trying to retrieve Name: %w", err)
	}

	if _, err := c.GetUptimeCheckConfig(ctx, r.urlNormalized()); err != nil {
		return err
	}

	return nil
}

func (c *Client) getUptimeCheckConfigRaw(ctx context.Context, r *UptimeCheckConfig) ([]byte, error) {
	if dcl.IsZeroValue(r.Period) {
		r.Period = dcl.String("60s")
	}

	u, err := uptimeCheckConfigGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.Retry)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) uptimeCheckConfigDiffsForRawDesired(ctx context.Context, rawDesired *UptimeCheckConfig, opts ...dcl.ApplyOption) (initial, desired *UptimeCheckConfig, diffs []uptimeCheckConfigDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *UptimeCheckConfig
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*UptimeCheckConfig); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected UptimeCheckConfig, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeUptimeCheckConfigDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetUptimeCheckConfig(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a UptimeCheckConfig resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve UptimeCheckConfig resource: %v", err)
		}
		c.Config.Logger.Info("Found that UptimeCheckConfig resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for UptimeCheckConfig: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for UptimeCheckConfig: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeUptimeCheckConfigInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for UptimeCheckConfig: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for UptimeCheckConfig: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffUptimeCheckConfig(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeUptimeCheckConfigInitialState(rawInitial, rawDesired *UptimeCheckConfig) (*UptimeCheckConfig, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeUptimeCheckConfigDesiredState(rawDesired, rawInitial *UptimeCheckConfig, opts ...dcl.ApplyOption) (*UptimeCheckConfig, error) {

	if dcl.IsZeroValue(rawDesired.Period) {
		rawDesired.Period = dcl.String("60s")
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*UptimeCheckConfig); !ok {
			return nil, fmt.Errorf("Initial state hint was of the wrong type; expected UptimeCheckConfig, got %T", sh)
		} else {
			_ = r
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MonitoredResource = canonicalizeUptimeCheckConfigMonitoredResource(rawDesired.MonitoredResource, nil, opts...)
		rawDesired.ResourceGroup = canonicalizeUptimeCheckConfigResourceGroup(rawDesired.ResourceGroup, nil, opts...)
		rawDesired.HttpCheck = canonicalizeUptimeCheckConfigHttpCheck(rawDesired.HttpCheck, nil, opts...)
		rawDesired.TcpCheck = canonicalizeUptimeCheckConfigTcpCheck(rawDesired.TcpCheck, nil, opts...)

		return rawDesired, nil
	}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	rawDesired.MonitoredResource = canonicalizeUptimeCheckConfigMonitoredResource(rawDesired.MonitoredResource, rawInitial.MonitoredResource, opts...)
	rawDesired.ResourceGroup = canonicalizeUptimeCheckConfigResourceGroup(rawDesired.ResourceGroup, rawInitial.ResourceGroup, opts...)
	rawDesired.HttpCheck = canonicalizeUptimeCheckConfigHttpCheck(rawDesired.HttpCheck, rawInitial.HttpCheck, opts...)
	rawDesired.TcpCheck = canonicalizeUptimeCheckConfigTcpCheck(rawDesired.TcpCheck, rawInitial.TcpCheck, opts...)
	if dcl.IsZeroValue(rawDesired.Period) {
		rawDesired.Period = rawInitial.Period
	}
	if dcl.IsZeroValue(rawDesired.Timeout) {
		rawDesired.Timeout = rawInitial.Timeout
	}
	if dcl.IsZeroValue(rawDesired.ContentMatchers) {
		rawDesired.ContentMatchers = rawInitial.ContentMatchers
	}
	if dcl.IsZeroValue(rawDesired.PrivateCheckers) {
		rawDesired.PrivateCheckers = rawInitial.PrivateCheckers
	}
	if dcl.IsZeroValue(rawDesired.SelectedRegions) {
		rawDesired.SelectedRegions = rawInitial.SelectedRegions
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeUptimeCheckConfigNewState(c *Client, rawNew, rawDesired *UptimeCheckConfig) (*UptimeCheckConfig, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.NameToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MonitoredResource) && dcl.IsEmptyValueIndirect(rawDesired.MonitoredResource) {
		rawNew.MonitoredResource = rawDesired.MonitoredResource
	} else {
		rawNew.MonitoredResource = canonicalizeNewUptimeCheckConfigMonitoredResource(c, rawDesired.MonitoredResource, rawNew.MonitoredResource)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ResourceGroup) && dcl.IsEmptyValueIndirect(rawDesired.ResourceGroup) {
		rawNew.ResourceGroup = rawDesired.ResourceGroup
	} else {
		rawNew.ResourceGroup = canonicalizeNewUptimeCheckConfigResourceGroup(c, rawDesired.ResourceGroup, rawNew.ResourceGroup)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpCheck) && dcl.IsEmptyValueIndirect(rawDesired.HttpCheck) {
		rawNew.HttpCheck = rawDesired.HttpCheck
	} else {
		rawNew.HttpCheck = canonicalizeNewUptimeCheckConfigHttpCheck(c, rawDesired.HttpCheck, rawNew.HttpCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.TcpCheck) && dcl.IsEmptyValueIndirect(rawDesired.TcpCheck) {
		rawNew.TcpCheck = rawDesired.TcpCheck
	} else {
		rawNew.TcpCheck = canonicalizeNewUptimeCheckConfigTcpCheck(c, rawDesired.TcpCheck, rawNew.TcpCheck)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Period) && dcl.IsEmptyValueIndirect(rawDesired.Period) {
		rawNew.Period = rawDesired.Period
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Timeout) && dcl.IsEmptyValueIndirect(rawDesired.Timeout) {
		rawNew.Timeout = rawDesired.Timeout
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ContentMatchers) && dcl.IsEmptyValueIndirect(rawDesired.ContentMatchers) {
		rawNew.ContentMatchers = rawDesired.ContentMatchers
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PrivateCheckers) && dcl.IsEmptyValueIndirect(rawDesired.PrivateCheckers) {
		rawNew.PrivateCheckers = rawDesired.PrivateCheckers
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelectedRegions) && dcl.IsEmptyValueIndirect(rawDesired.SelectedRegions) {
		rawNew.SelectedRegions = rawDesired.SelectedRegions
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeUptimeCheckConfigMonitoredResource(des, initial *UptimeCheckConfigMonitoredResource, opts ...dcl.ApplyOption) *UptimeCheckConfigMonitoredResource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.IsZeroValue(des.FilterLabels) {
		des.FilterLabels = initial.FilterLabels
	}

	return des
}

func canonicalizeNewUptimeCheckConfigMonitoredResource(c *Client, des, nw *UptimeCheckConfigMonitoredResource) *UptimeCheckConfigMonitoredResource {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigMonitoredResourceSet(c *Client, des, nw []UptimeCheckConfigMonitoredResource) []UptimeCheckConfigMonitoredResource {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigMonitoredResource
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigMonitoredResource(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeUptimeCheckConfigResourceGroup(des, initial *UptimeCheckConfigResourceGroup, opts ...dcl.ApplyOption) *UptimeCheckConfigResourceGroup {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.GroupId, initial.GroupId) || dcl.IsZeroValue(des.GroupId) {
		des.GroupId = initial.GroupId
	}
	if dcl.IsZeroValue(des.ResourceType) {
		des.ResourceType = initial.ResourceType
	}

	return des
}

func canonicalizeNewUptimeCheckConfigResourceGroup(c *Client, des, nw *UptimeCheckConfigResourceGroup) *UptimeCheckConfigResourceGroup {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.GroupId, nw.GroupId) || dcl.IsZeroValue(des.GroupId) {
		nw.GroupId = des.GroupId
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigResourceGroupSet(c *Client, des, nw []UptimeCheckConfigResourceGroup) []UptimeCheckConfigResourceGroup {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigResourceGroup
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigResourceGroup(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeUptimeCheckConfigHttpCheck(des, initial *UptimeCheckConfigHttpCheck, opts ...dcl.ApplyOption) *UptimeCheckConfigHttpCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.RequestMethod) {
		des.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}

	if dcl.IsZeroValue(des.Path) {
		des.Path = dcl.String("/")
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RequestMethod) {
		des.RequestMethod = initial.RequestMethod
	}
	if dcl.IsZeroValue(des.UseSsl) {
		des.UseSsl = initial.UseSsl
	}
	if dcl.IsZeroValue(des.Path) {
		des.Path = initial.Path
	}
	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}
	des.AuthInfo = canonicalizeUptimeCheckConfigHttpCheckAuthInfo(des.AuthInfo, initial.AuthInfo, opts...)
	if dcl.IsZeroValue(des.MaskHeaders) {
		des.MaskHeaders = initial.MaskHeaders
	}
	if dcl.IsZeroValue(des.Headers) {
		des.Headers = initial.Headers
	}
	if dcl.IsZeroValue(des.ContentType) {
		des.ContentType = initial.ContentType
	}
	if dcl.IsZeroValue(des.ValidateSsl) {
		des.ValidateSsl = initial.ValidateSsl
	}
	if dcl.IsZeroValue(des.Body) {
		des.Body = initial.Body
	}

	return des
}

func canonicalizeNewUptimeCheckConfigHttpCheck(c *Client, des, nw *UptimeCheckConfigHttpCheck) *UptimeCheckConfigHttpCheck {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.RequestMethod) {
		nw.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}

	if dcl.IsZeroValue(nw.Path) {
		nw.Path = dcl.String("/")
	}

	nw.AuthInfo = canonicalizeNewUptimeCheckConfigHttpCheckAuthInfo(c, des.AuthInfo, nw.AuthInfo)

	return nw
}

func canonicalizeNewUptimeCheckConfigHttpCheckSet(c *Client, des, nw []UptimeCheckConfigHttpCheck) []UptimeCheckConfigHttpCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigHttpCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigHttpCheck(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeUptimeCheckConfigHttpCheckAuthInfo(des, initial *UptimeCheckConfigHttpCheckAuthInfo, opts ...dcl.ApplyOption) *UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}

	return des
}

func canonicalizeNewUptimeCheckConfigHttpCheckAuthInfo(c *Client, des, nw *UptimeCheckConfigHttpCheckAuthInfo) *UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigHttpCheckAuthInfoSet(c *Client, des, nw []UptimeCheckConfigHttpCheckAuthInfo) []UptimeCheckConfigHttpCheckAuthInfo {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigHttpCheckAuthInfo
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigHttpCheckAuthInfo(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeUptimeCheckConfigTcpCheck(des, initial *UptimeCheckConfigTcpCheck, opts ...dcl.ApplyOption) *UptimeCheckConfigTcpCheck {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Port) {
		des.Port = initial.Port
	}

	return des
}

func canonicalizeNewUptimeCheckConfigTcpCheck(c *Client, des, nw *UptimeCheckConfigTcpCheck) *UptimeCheckConfigTcpCheck {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigTcpCheckSet(c *Client, des, nw []UptimeCheckConfigTcpCheck) []UptimeCheckConfigTcpCheck {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigTcpCheck
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigTcpCheck(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeUptimeCheckConfigContentMatchers(des, initial *UptimeCheckConfigContentMatchers, opts ...dcl.ApplyOption) *UptimeCheckConfigContentMatchers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if dcl.IsZeroValue(des.Matcher) {
		des.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	if sh := dcl.FetchStateHint(opts); sh != nil {
		r := sh.(*UptimeCheckConfig)
		_ = r
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Content) {
		des.Content = initial.Content
	}
	if dcl.IsZeroValue(des.Matcher) {
		des.Matcher = initial.Matcher
	}

	return des
}

func canonicalizeNewUptimeCheckConfigContentMatchers(c *Client, des, nw *UptimeCheckConfigContentMatchers) *UptimeCheckConfigContentMatchers {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.Matcher) {
		nw.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	return nw
}

func canonicalizeNewUptimeCheckConfigContentMatchersSet(c *Client, des, nw []UptimeCheckConfigContentMatchers) []UptimeCheckConfigContentMatchers {
	if des == nil {
		return nw
	}
	var reorderedNew []UptimeCheckConfigContentMatchers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareUptimeCheckConfigContentMatchers(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

type uptimeCheckConfigDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         uptimeCheckConfigApiOperation
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffUptimeCheckConfig(c *Client, desired, actual *UptimeCheckConfig, opts ...dcl.ApplyOption) ([]uptimeCheckConfigDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []uptimeCheckConfigDiff
	if !dcl.IsZeroValue(desired.Name) && !dcl.NameToSelfLink(desired.Name, actual.Name) {
		c.Config.Logger.Infof("Detected diff in Name.\nDESIRED: %v\nACTUAL: %v", desired.Name, actual.Name)
		diffs = append(diffs, uptimeCheckConfigDiff{
			RequiresRecreate: true,
			FieldName:        "Name",
		})
	}
	if !dcl.IsZeroValue(desired.DisplayName) && (dcl.IsZeroValue(actual.DisplayName) || !reflect.DeepEqual(*desired.DisplayName, *actual.DisplayName)) {
		c.Config.Logger.Infof("Detected diff in DisplayName.\nDESIRED: %v\nACTUAL: %v", desired.DisplayName, actual.DisplayName)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "DisplayName",
		})

	}
	if compareUptimeCheckConfigMonitoredResource(c, desired.MonitoredResource, actual.MonitoredResource) {
		c.Config.Logger.Infof("Detected diff in MonitoredResource.\nDESIRED: %v\nACTUAL: %v", desired.MonitoredResource, actual.MonitoredResource)
		diffs = append(diffs, uptimeCheckConfigDiff{
			RequiresRecreate: true,
			FieldName:        "MonitoredResource",
		})
	}
	if compareUptimeCheckConfigResourceGroup(c, desired.ResourceGroup, actual.ResourceGroup) {
		c.Config.Logger.Infof("Detected diff in ResourceGroup.\nDESIRED: %v\nACTUAL: %v", desired.ResourceGroup, actual.ResourceGroup)
		diffs = append(diffs, uptimeCheckConfigDiff{
			RequiresRecreate: true,
			FieldName:        "ResourceGroup",
		})
	}
	if compareUptimeCheckConfigHttpCheck(c, desired.HttpCheck, actual.HttpCheck) {
		c.Config.Logger.Infof("Detected diff in HttpCheck.\nDESIRED: %v\nACTUAL: %v", desired.HttpCheck, actual.HttpCheck)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "HttpCheck",
		})

	}
	if compareUptimeCheckConfigTcpCheck(c, desired.TcpCheck, actual.TcpCheck) {
		c.Config.Logger.Infof("Detected diff in TcpCheck.\nDESIRED: %v\nACTUAL: %v", desired.TcpCheck, actual.TcpCheck)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "TcpCheck",
		})

	}
	if !dcl.IsZeroValue(desired.Period) && (dcl.IsZeroValue(actual.Period) || !reflect.DeepEqual(*desired.Period, *actual.Period)) {
		c.Config.Logger.Infof("Detected diff in Period.\nDESIRED: %v\nACTUAL: %v", desired.Period, actual.Period)
		diffs = append(diffs, uptimeCheckConfigDiff{
			RequiresRecreate: true,
			FieldName:        "Period",
		})
	}
	if !dcl.IsZeroValue(desired.Timeout) && (dcl.IsZeroValue(actual.Timeout) || !reflect.DeepEqual(*desired.Timeout, *actual.Timeout)) {
		c.Config.Logger.Infof("Detected diff in Timeout.\nDESIRED: %v\nACTUAL: %v", desired.Timeout, actual.Timeout)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "Timeout",
		})

	}
	if compareUptimeCheckConfigContentMatchersSlice(c, desired.ContentMatchers, actual.ContentMatchers) {
		c.Config.Logger.Infof("Detected diff in ContentMatchers.\nDESIRED: %v\nACTUAL: %v", desired.ContentMatchers, actual.ContentMatchers)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "ContentMatchers",
		})

	}
	if !dcl.SliceEquals(desired.PrivateCheckers, actual.PrivateCheckers) {
		c.Config.Logger.Infof("Detected diff in PrivateCheckers.\nDESIRED: %v\nACTUAL: %v", desired.PrivateCheckers, actual.PrivateCheckers)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "PrivateCheckers",
		})

	}
	if !dcl.SliceEquals(desired.SelectedRegions, actual.SelectedRegions) {
		c.Config.Logger.Infof("Detected diff in SelectedRegions.\nDESIRED: %v\nACTUAL: %v", desired.SelectedRegions, actual.SelectedRegions)

		diffs = append(diffs, uptimeCheckConfigDiff{
			UpdateOp:  &updateUptimeCheckConfigUpdateOperation{},
			FieldName: "SelectedRegions",
		})

	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []uptimeCheckConfigDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareUptimeCheckConfigMonitoredResourceSlice(c *Client, desired, actual []UptimeCheckConfigMonitoredResource) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigMonitoredResource, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigMonitoredResource(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigMonitoredResource, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigMonitoredResource(c *Client, desired, actual *UptimeCheckConfigMonitoredResource) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Type == nil && desired.Type != nil && !dcl.IsEmptyValueIndirect(desired.Type) {
		c.Config.Logger.Infof("desired Type %s - but actually nil", dcl.SprintResource(desired.Type))
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) && !(dcl.IsEmptyValueIndirect(desired.Type) && dcl.IsZeroValue(actual.Type)) {
		c.Config.Logger.Infof("Diff in Type. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if actual.FilterLabels == nil && desired.FilterLabels != nil && !dcl.IsEmptyValueIndirect(desired.FilterLabels) {
		c.Config.Logger.Infof("desired FilterLabels %s - but actually nil", dcl.SprintResource(desired.FilterLabels))
		return true
	}
	if !reflect.DeepEqual(desired.FilterLabels, actual.FilterLabels) && !dcl.IsZeroValue(desired.FilterLabels) {
		c.Config.Logger.Infof("Diff in FilterLabels. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FilterLabels), dcl.SprintResource(actual.FilterLabels))
		return true
	}
	return false
}
func compareUptimeCheckConfigResourceGroupSlice(c *Client, desired, actual []UptimeCheckConfigResourceGroup) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigResourceGroup, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigResourceGroup(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigResourceGroup, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigResourceGroup(c *Client, desired, actual *UptimeCheckConfigResourceGroup) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.ResourceType == nil && desired.ResourceType != nil && !dcl.IsEmptyValueIndirect(desired.ResourceType) {
		c.Config.Logger.Infof("desired ResourceType %s - but actually nil", dcl.SprintResource(desired.ResourceType))
		return true
	}
	if !reflect.DeepEqual(desired.ResourceType, actual.ResourceType) && !dcl.IsZeroValue(desired.ResourceType) && !(dcl.IsEmptyValueIndirect(desired.ResourceType) && dcl.IsZeroValue(actual.ResourceType)) {
		c.Config.Logger.Infof("Diff in ResourceType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ResourceType), dcl.SprintResource(actual.ResourceType))
		return true
	}
	return false
}
func compareUptimeCheckConfigHttpCheckSlice(c *Client, desired, actual []UptimeCheckConfigHttpCheck) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigHttpCheck, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigHttpCheck(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigHttpCheck, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigHttpCheck(c *Client, desired, actual *UptimeCheckConfigHttpCheck) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.RequestMethod == nil && desired.RequestMethod != nil && !dcl.IsEmptyValueIndirect(desired.RequestMethod) {
		c.Config.Logger.Infof("desired RequestMethod %s - but actually nil", dcl.SprintResource(desired.RequestMethod))
		return true
	}
	if !reflect.DeepEqual(desired.RequestMethod, actual.RequestMethod) && !dcl.IsZeroValue(desired.RequestMethod) && !(dcl.IsEmptyValueIndirect(desired.RequestMethod) && dcl.IsZeroValue(actual.RequestMethod)) {
		c.Config.Logger.Infof("Diff in RequestMethod. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequestMethod), dcl.SprintResource(actual.RequestMethod))
		return true
	}
	if actual.UseSsl == nil && desired.UseSsl != nil && !dcl.IsEmptyValueIndirect(desired.UseSsl) {
		c.Config.Logger.Infof("desired UseSsl %s - but actually nil", dcl.SprintResource(desired.UseSsl))
		return true
	}
	if !reflect.DeepEqual(desired.UseSsl, actual.UseSsl) && !dcl.IsZeroValue(desired.UseSsl) && !(dcl.IsEmptyValueIndirect(desired.UseSsl) && dcl.IsZeroValue(actual.UseSsl)) {
		c.Config.Logger.Infof("Diff in UseSsl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UseSsl), dcl.SprintResource(actual.UseSsl))
		return true
	}
	if actual.Path == nil && desired.Path != nil && !dcl.IsEmptyValueIndirect(desired.Path) {
		c.Config.Logger.Infof("desired Path %s - but actually nil", dcl.SprintResource(desired.Path))
		return true
	}
	if !reflect.DeepEqual(desired.Path, actual.Path) && !dcl.IsZeroValue(desired.Path) && !(dcl.IsEmptyValueIndirect(desired.Path) && dcl.IsZeroValue(actual.Path)) {
		c.Config.Logger.Infof("Diff in Path. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Path), dcl.SprintResource(actual.Path))
		return true
	}
	if actual.Port == nil && desired.Port != nil && !dcl.IsEmptyValueIndirect(desired.Port) {
		c.Config.Logger.Infof("desired Port %s - but actually nil", dcl.SprintResource(desired.Port))
		return true
	}
	if !reflect.DeepEqual(desired.Port, actual.Port) && !dcl.IsZeroValue(desired.Port) && !(dcl.IsEmptyValueIndirect(desired.Port) && dcl.IsZeroValue(actual.Port)) {
		c.Config.Logger.Infof("Diff in Port. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Port), dcl.SprintResource(actual.Port))
		return true
	}
	if actual.AuthInfo == nil && desired.AuthInfo != nil && !dcl.IsEmptyValueIndirect(desired.AuthInfo) {
		c.Config.Logger.Infof("desired AuthInfo %s - but actually nil", dcl.SprintResource(desired.AuthInfo))
		return true
	}
	if compareUptimeCheckConfigHttpCheckAuthInfo(c, desired.AuthInfo, actual.AuthInfo) && !dcl.IsZeroValue(desired.AuthInfo) {
		c.Config.Logger.Infof("Diff in AuthInfo. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthInfo), dcl.SprintResource(actual.AuthInfo))
		return true
	}
	if actual.MaskHeaders == nil && desired.MaskHeaders != nil && !dcl.IsEmptyValueIndirect(desired.MaskHeaders) {
		c.Config.Logger.Infof("desired MaskHeaders %s - but actually nil", dcl.SprintResource(desired.MaskHeaders))
		return true
	}
	if !reflect.DeepEqual(desired.MaskHeaders, actual.MaskHeaders) && !dcl.IsZeroValue(desired.MaskHeaders) && !(dcl.IsEmptyValueIndirect(desired.MaskHeaders) && dcl.IsZeroValue(actual.MaskHeaders)) {
		c.Config.Logger.Infof("Diff in MaskHeaders. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaskHeaders), dcl.SprintResource(actual.MaskHeaders))
		return true
	}
	if actual.Headers == nil && desired.Headers != nil && !dcl.IsEmptyValueIndirect(desired.Headers) {
		c.Config.Logger.Infof("desired Headers %s - but actually nil", dcl.SprintResource(desired.Headers))
		return true
	}
	if !reflect.DeepEqual(desired.Headers, actual.Headers) && !dcl.IsZeroValue(desired.Headers) {
		c.Config.Logger.Infof("Diff in Headers. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Headers), dcl.SprintResource(actual.Headers))
		return true
	}
	if actual.ContentType == nil && desired.ContentType != nil && !dcl.IsEmptyValueIndirect(desired.ContentType) {
		c.Config.Logger.Infof("desired ContentType %s - but actually nil", dcl.SprintResource(desired.ContentType))
		return true
	}
	if !reflect.DeepEqual(desired.ContentType, actual.ContentType) && !dcl.IsZeroValue(desired.ContentType) && !(dcl.IsEmptyValueIndirect(desired.ContentType) && dcl.IsZeroValue(actual.ContentType)) {
		c.Config.Logger.Infof("Diff in ContentType. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ContentType), dcl.SprintResource(actual.ContentType))
		return true
	}
	if actual.ValidateSsl == nil && desired.ValidateSsl != nil && !dcl.IsEmptyValueIndirect(desired.ValidateSsl) {
		c.Config.Logger.Infof("desired ValidateSsl %s - but actually nil", dcl.SprintResource(desired.ValidateSsl))
		return true
	}
	if !reflect.DeepEqual(desired.ValidateSsl, actual.ValidateSsl) && !dcl.IsZeroValue(desired.ValidateSsl) && !(dcl.IsEmptyValueIndirect(desired.ValidateSsl) && dcl.IsZeroValue(actual.ValidateSsl)) {
		c.Config.Logger.Infof("Diff in ValidateSsl. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ValidateSsl), dcl.SprintResource(actual.ValidateSsl))
		return true
	}
	if actual.Body == nil && desired.Body != nil && !dcl.IsEmptyValueIndirect(desired.Body) {
		c.Config.Logger.Infof("desired Body %s - but actually nil", dcl.SprintResource(desired.Body))
		return true
	}
	if !reflect.DeepEqual(desired.Body, actual.Body) && !dcl.IsZeroValue(desired.Body) && !(dcl.IsEmptyValueIndirect(desired.Body) && dcl.IsZeroValue(actual.Body)) {
		c.Config.Logger.Infof("Diff in Body. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Body), dcl.SprintResource(actual.Body))
		return true
	}
	return false
}
func compareUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, desired, actual []UptimeCheckConfigHttpCheckAuthInfo) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigHttpCheckAuthInfo, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigHttpCheckAuthInfo(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigHttpCheckAuthInfo, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigHttpCheckAuthInfo(c *Client, desired, actual *UptimeCheckConfigHttpCheckAuthInfo) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Username == nil && desired.Username != nil && !dcl.IsEmptyValueIndirect(desired.Username) {
		c.Config.Logger.Infof("desired Username %s - but actually nil", dcl.SprintResource(desired.Username))
		return true
	}
	if !reflect.DeepEqual(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) && !(dcl.IsEmptyValueIndirect(desired.Username) && dcl.IsZeroValue(actual.Username)) {
		c.Config.Logger.Infof("Diff in Username. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	return false
}
func compareUptimeCheckConfigTcpCheckSlice(c *Client, desired, actual []UptimeCheckConfigTcpCheck) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigTcpCheck, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigTcpCheck(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigTcpCheck, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigTcpCheck(c *Client, desired, actual *UptimeCheckConfigTcpCheck) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Port == nil && desired.Port != nil && !dcl.IsEmptyValueIndirect(desired.Port) {
		c.Config.Logger.Infof("desired Port %s - but actually nil", dcl.SprintResource(desired.Port))
		return true
	}
	if !reflect.DeepEqual(desired.Port, actual.Port) && !dcl.IsZeroValue(desired.Port) && !(dcl.IsEmptyValueIndirect(desired.Port) && dcl.IsZeroValue(actual.Port)) {
		c.Config.Logger.Infof("Diff in Port. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Port), dcl.SprintResource(actual.Port))
		return true
	}
	return false
}
func compareUptimeCheckConfigContentMatchersSlice(c *Client, desired, actual []UptimeCheckConfigContentMatchers) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigContentMatchers, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigContentMatchers(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigContentMatchers, element %d. \nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigContentMatchers(c *Client, desired, actual *UptimeCheckConfigContentMatchers) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if actual.Content == nil && desired.Content != nil && !dcl.IsEmptyValueIndirect(desired.Content) {
		c.Config.Logger.Infof("desired Content %s - but actually nil", dcl.SprintResource(desired.Content))
		return true
	}
	if !reflect.DeepEqual(desired.Content, actual.Content) && !dcl.IsZeroValue(desired.Content) && !(dcl.IsEmptyValueIndirect(desired.Content) && dcl.IsZeroValue(actual.Content)) {
		c.Config.Logger.Infof("Diff in Content. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Content), dcl.SprintResource(actual.Content))
		return true
	}
	if actual.Matcher == nil && desired.Matcher != nil && !dcl.IsEmptyValueIndirect(desired.Matcher) {
		c.Config.Logger.Infof("desired Matcher %s - but actually nil", dcl.SprintResource(desired.Matcher))
		return true
	}
	if !reflect.DeepEqual(desired.Matcher, actual.Matcher) && !dcl.IsZeroValue(desired.Matcher) && !(dcl.IsEmptyValueIndirect(desired.Matcher) && dcl.IsZeroValue(actual.Matcher)) {
		c.Config.Logger.Infof("Diff in Matcher. \nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Matcher), dcl.SprintResource(actual.Matcher))
		return true
	}
	return false
}
func compareUptimeCheckConfigResourceGroupResourceTypeEnumSlice(c *Client, desired, actual []UptimeCheckConfigResourceGroupResourceTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigResourceGroupResourceTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigResourceGroupResourceTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigResourceGroupResourceTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigResourceGroupResourceTypeEnum(c *Client, desired, actual *UptimeCheckConfigResourceGroupResourceTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUptimeCheckConfigHttpCheckRequestMethodEnumSlice(c *Client, desired, actual []UptimeCheckConfigHttpCheckRequestMethodEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigHttpCheckRequestMethodEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigHttpCheckRequestMethodEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigHttpCheckRequestMethodEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigHttpCheckRequestMethodEnum(c *Client, desired, actual *UptimeCheckConfigHttpCheckRequestMethodEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUptimeCheckConfigHttpCheckContentTypeEnumSlice(c *Client, desired, actual []UptimeCheckConfigHttpCheckContentTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigHttpCheckContentTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigHttpCheckContentTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigHttpCheckContentTypeEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigHttpCheckContentTypeEnum(c *Client, desired, actual *UptimeCheckConfigHttpCheckContentTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareUptimeCheckConfigContentMatchersMatcherEnumSlice(c *Client, desired, actual []UptimeCheckConfigContentMatchersMatcherEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in UptimeCheckConfigContentMatchersMatcherEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareUptimeCheckConfigContentMatchersMatcherEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in UptimeCheckConfigContentMatchersMatcherEnum, element %d. \nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareUptimeCheckConfigContentMatchersMatcherEnum(c *Client, desired, actual *UptimeCheckConfigContentMatchersMatcherEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *UptimeCheckConfig) urlNormalized() *UptimeCheckConfig {
	normalized := deepcopy.Copy(*r).(UptimeCheckConfig)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *UptimeCheckConfig) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UptimeCheckConfig) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *UptimeCheckConfig) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *UptimeCheckConfig) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("v3/projects/{{project}}/uptimeCheckConfigs/{{name}}", "https://monitoring.googleapis.com/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the UptimeCheckConfig resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *UptimeCheckConfig) marshal(c *Client) ([]byte, error) {
	m, err := expandUptimeCheckConfig(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling UptimeCheckConfig: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalUptimeCheckConfig decodes JSON responses into the UptimeCheckConfig resource schema.
func unmarshalUptimeCheckConfig(b []byte, c *Client) (*UptimeCheckConfig, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenUptimeCheckConfig(c, m), nil
}

// expandUptimeCheckConfig expands UptimeCheckConfig into a JSON request object.
func expandUptimeCheckConfig(c *Client, f *UptimeCheckConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v, err := expandUptimeCheckConfigMonitoredResource(c, f.MonitoredResource); err != nil {
		return nil, fmt.Errorf("error expanding MonitoredResource into monitoredResource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["monitoredResource"] = v
	}
	if v, err := expandUptimeCheckConfigResourceGroup(c, f.ResourceGroup); err != nil {
		return nil, fmt.Errorf("error expanding ResourceGroup into resourceGroup: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroup"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheck(c, f.HttpCheck); err != nil {
		return nil, fmt.Errorf("error expanding HttpCheck into httpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["httpCheck"] = v
	}
	if v, err := expandUptimeCheckConfigTcpCheck(c, f.TcpCheck); err != nil {
		return nil, fmt.Errorf("error expanding TcpCheck into tcpCheck: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tcpCheck"] = v
	}
	if v := f.Period; !dcl.IsEmptyValueIndirect(v) {
		m["period"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v, err := expandUptimeCheckConfigContentMatchersSlice(c, f.ContentMatchers); err != nil {
		return nil, fmt.Errorf("error expanding ContentMatchers into contentMatchers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["contentMatchers"] = v
	}
	if v := f.PrivateCheckers; !dcl.IsEmptyValueIndirect(v) {
		m["privateCheckers"] = v
	}
	if v := f.SelectedRegions; !dcl.IsEmptyValueIndirect(v) {
		m["selectedRegions"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfig flattens UptimeCheckConfig from a JSON request object into the
// UptimeCheckConfig type.
func flattenUptimeCheckConfig(c *Client, i interface{}) *UptimeCheckConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &UptimeCheckConfig{}
	r.Name = dcl.FlattenSecretValue(m["name"])
	r.DisplayName = dcl.FlattenString(m["displayName"])
	r.MonitoredResource = flattenUptimeCheckConfigMonitoredResource(c, m["monitoredResource"])
	r.ResourceGroup = flattenUptimeCheckConfigResourceGroup(c, m["resourceGroup"])
	r.HttpCheck = flattenUptimeCheckConfigHttpCheck(c, m["httpCheck"])
	r.TcpCheck = flattenUptimeCheckConfigTcpCheck(c, m["tcpCheck"])
	r.Period = dcl.FlattenString(m["period"])
	if _, ok := m["period"]; !ok {
		c.Config.Logger.Info("Using default value for period")
		r.Period = dcl.String("60s")
	}
	r.Timeout = dcl.FlattenString(m["timeout"])
	r.ContentMatchers = flattenUptimeCheckConfigContentMatchersSlice(c, m["contentMatchers"])
	r.PrivateCheckers = dcl.FlattenStringSlice(m["privateCheckers"])
	r.SelectedRegions = dcl.FlattenStringSlice(m["selectedRegions"])
	r.Project = dcl.FlattenString(m["project"])

	return r
}

// expandUptimeCheckConfigMonitoredResourceMap expands the contents of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResourceMap(c *Client, f map[string]UptimeCheckConfigMonitoredResource) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigMonitoredResource(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigMonitoredResourceSlice expands the contents of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResourceSlice(c *Client, f []UptimeCheckConfigMonitoredResource) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigMonitoredResource(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigMonitoredResourceMap flattens the contents of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResourceMap(c *Client, i interface{}) map[string]UptimeCheckConfigMonitoredResource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigMonitoredResource{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigMonitoredResource{}
	}

	items := make(map[string]UptimeCheckConfigMonitoredResource)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigMonitoredResource(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigMonitoredResourceSlice flattens the contents of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResourceSlice(c *Client, i interface{}) []UptimeCheckConfigMonitoredResource {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigMonitoredResource{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigMonitoredResource{}
	}

	items := make([]UptimeCheckConfigMonitoredResource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigMonitoredResource(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigMonitoredResource expands an instance of UptimeCheckConfigMonitoredResource into a JSON
// request object.
func expandUptimeCheckConfigMonitoredResource(c *Client, f *UptimeCheckConfigMonitoredResource) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.FilterLabels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigMonitoredResource flattens an instance of UptimeCheckConfigMonitoredResource from a JSON
// response object.
func flattenUptimeCheckConfigMonitoredResource(c *Client, i interface{}) *UptimeCheckConfigMonitoredResource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigMonitoredResource{}
	r.Type = dcl.FlattenString(m["type"])
	r.FilterLabels = dcl.FlattenKeyValuePairs(m["labels"])

	return r
}

// expandUptimeCheckConfigResourceGroupMap expands the contents of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroupMap(c *Client, f map[string]UptimeCheckConfigResourceGroup) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigResourceGroup(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigResourceGroupSlice expands the contents of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroupSlice(c *Client, f []UptimeCheckConfigResourceGroup) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigResourceGroup(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigResourceGroupMap flattens the contents of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupMap(c *Client, i interface{}) map[string]UptimeCheckConfigResourceGroup {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigResourceGroup{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigResourceGroup{}
	}

	items := make(map[string]UptimeCheckConfigResourceGroup)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigResourceGroup(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigResourceGroupSlice flattens the contents of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupSlice(c *Client, i interface{}) []UptimeCheckConfigResourceGroup {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigResourceGroup{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigResourceGroup{}
	}

	items := make([]UptimeCheckConfigResourceGroup, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigResourceGroup(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigResourceGroup expands an instance of UptimeCheckConfigResourceGroup into a JSON
// request object.
func expandUptimeCheckConfigResourceGroup(c *Client, f *UptimeCheckConfigResourceGroup) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding GroupId into groupId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["groupId"] = v
	}
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigResourceGroup flattens an instance of UptimeCheckConfigResourceGroup from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroup(c *Client, i interface{}) *UptimeCheckConfigResourceGroup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigResourceGroup{}
	r.GroupId = dcl.FlattenString(m["groupId"])
	r.ResourceType = flattenUptimeCheckConfigResourceGroupResourceTypeEnum(m["resourceType"])

	return r
}

// expandUptimeCheckConfigHttpCheckMap expands the contents of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckMap(c *Client, f map[string]UptimeCheckConfigHttpCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigHttpCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigHttpCheckSlice expands the contents of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckSlice(c *Client, f []UptimeCheckConfigHttpCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigHttpCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigHttpCheckMap flattens the contents of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckMap(c *Client, i interface{}) map[string]UptimeCheckConfigHttpCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigHttpCheck{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigHttpCheck{}
	}

	items := make(map[string]UptimeCheckConfigHttpCheck)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigHttpCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckSlice flattens the contents of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheck{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheck{}
	}

	items := make([]UptimeCheckConfigHttpCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigHttpCheck expands an instance of UptimeCheckConfigHttpCheck into a JSON
// request object.
func expandUptimeCheckConfigHttpCheck(c *Client, f *UptimeCheckConfigHttpCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RequestMethod; !dcl.IsEmptyValueIndirect(v) {
		m["requestMethod"] = v
	}
	if v := f.UseSsl; !dcl.IsEmptyValueIndirect(v) {
		m["useSsl"] = v
	}
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, f.AuthInfo); err != nil {
		return nil, fmt.Errorf("error expanding AuthInfo into authInfo: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authInfo"] = v
	}
	if v := f.MaskHeaders; !dcl.IsEmptyValueIndirect(v) {
		m["maskHeaders"] = v
	}
	if v := f.Headers; !dcl.IsEmptyValueIndirect(v) {
		m["headers"] = v
	}
	if v := f.ContentType; !dcl.IsEmptyValueIndirect(v) {
		m["contentType"] = v
	}
	if v := f.ValidateSsl; !dcl.IsEmptyValueIndirect(v) {
		m["validateSsl"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigHttpCheck flattens an instance of UptimeCheckConfigHttpCheck from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheck(c *Client, i interface{}) *UptimeCheckConfigHttpCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigHttpCheck{}
	r.RequestMethod = flattenUptimeCheckConfigHttpCheckRequestMethodEnum(m["requestMethod"])
	if dcl.IsEmptyValueIndirect(m["requestMethod"]) {
		c.Config.Logger.Info("Using default value for requestMethod.")
		r.RequestMethod = UptimeCheckConfigHttpCheckRequestMethodEnumRef("GET")
	}
	r.UseSsl = dcl.FlattenBool(m["useSsl"])
	r.Path = dcl.FlattenString(m["path"])
	if dcl.IsEmptyValueIndirect(m["path"]) {
		c.Config.Logger.Info("Using default value for path.")
		r.Path = dcl.String("/")
	}
	r.Port = dcl.FlattenInteger(m["port"])
	r.AuthInfo = flattenUptimeCheckConfigHttpCheckAuthInfo(c, m["authInfo"])
	r.MaskHeaders = dcl.FlattenBool(m["maskHeaders"])
	r.Headers = dcl.FlattenKeyValuePairs(m["headers"])
	r.ContentType = flattenUptimeCheckConfigHttpCheckContentTypeEnum(m["contentType"])
	r.ValidateSsl = dcl.FlattenBool(m["validateSsl"])
	r.Body = dcl.FlattenString(m["body"])

	return r
}

// expandUptimeCheckConfigHttpCheckAuthInfoMap expands the contents of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfoMap(c *Client, f map[string]UptimeCheckConfigHttpCheckAuthInfo) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigHttpCheckAuthInfoSlice expands the contents of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, f []UptimeCheckConfigHttpCheckAuthInfo) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigHttpCheckAuthInfo(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigHttpCheckAuthInfoMap flattens the contents of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfoMap(c *Client, i interface{}) map[string]UptimeCheckConfigHttpCheckAuthInfo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigHttpCheckAuthInfo{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigHttpCheckAuthInfo{}
	}

	items := make(map[string]UptimeCheckConfigHttpCheckAuthInfo)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigHttpCheckAuthInfo(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckAuthInfoSlice flattens the contents of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfoSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckAuthInfo {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckAuthInfo{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckAuthInfo{}
	}

	items := make([]UptimeCheckConfigHttpCheckAuthInfo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckAuthInfo(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigHttpCheckAuthInfo expands an instance of UptimeCheckConfigHttpCheckAuthInfo into a JSON
// request object.
func expandUptimeCheckConfigHttpCheckAuthInfo(c *Client, f *UptimeCheckConfigHttpCheckAuthInfo) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigHttpCheckAuthInfo flattens an instance of UptimeCheckConfigHttpCheckAuthInfo from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckAuthInfo(c *Client, i interface{}) *UptimeCheckConfigHttpCheckAuthInfo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigHttpCheckAuthInfo{}
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenSecretValue(m["password"])

	return r
}

// expandUptimeCheckConfigTcpCheckMap expands the contents of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheckMap(c *Client, f map[string]UptimeCheckConfigTcpCheck) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigTcpCheck(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigTcpCheckSlice expands the contents of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheckSlice(c *Client, f []UptimeCheckConfigTcpCheck) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigTcpCheck(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigTcpCheckMap flattens the contents of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheckMap(c *Client, i interface{}) map[string]UptimeCheckConfigTcpCheck {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigTcpCheck{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigTcpCheck{}
	}

	items := make(map[string]UptimeCheckConfigTcpCheck)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigTcpCheck(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigTcpCheckSlice flattens the contents of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheckSlice(c *Client, i interface{}) []UptimeCheckConfigTcpCheck {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigTcpCheck{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigTcpCheck{}
	}

	items := make([]UptimeCheckConfigTcpCheck, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigTcpCheck(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigTcpCheck expands an instance of UptimeCheckConfigTcpCheck into a JSON
// request object.
func expandUptimeCheckConfigTcpCheck(c *Client, f *UptimeCheckConfigTcpCheck) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigTcpCheck flattens an instance of UptimeCheckConfigTcpCheck from a JSON
// response object.
func flattenUptimeCheckConfigTcpCheck(c *Client, i interface{}) *UptimeCheckConfigTcpCheck {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigTcpCheck{}
	r.Port = dcl.FlattenInteger(m["port"])

	return r
}

// expandUptimeCheckConfigContentMatchersMap expands the contents of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchersMap(c *Client, f map[string]UptimeCheckConfigContentMatchers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandUptimeCheckConfigContentMatchers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandUptimeCheckConfigContentMatchersSlice expands the contents of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchersSlice(c *Client, f []UptimeCheckConfigContentMatchers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandUptimeCheckConfigContentMatchers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenUptimeCheckConfigContentMatchersMap flattens the contents of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersMap(c *Client, i interface{}) map[string]UptimeCheckConfigContentMatchers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]UptimeCheckConfigContentMatchers{}
	}

	if len(a) == 0 {
		return map[string]UptimeCheckConfigContentMatchers{}
	}

	items := make(map[string]UptimeCheckConfigContentMatchers)
	for k, item := range a {
		items[k] = *flattenUptimeCheckConfigContentMatchers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenUptimeCheckConfigContentMatchersSlice flattens the contents of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersSlice(c *Client, i interface{}) []UptimeCheckConfigContentMatchers {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigContentMatchers{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigContentMatchers{}
	}

	items := make([]UptimeCheckConfigContentMatchers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigContentMatchers(c, item.(map[string]interface{})))
	}

	return items
}

// expandUptimeCheckConfigContentMatchers expands an instance of UptimeCheckConfigContentMatchers into a JSON
// request object.
func expandUptimeCheckConfigContentMatchers(c *Client, f *UptimeCheckConfigContentMatchers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Content; !dcl.IsEmptyValueIndirect(v) {
		m["content"] = v
	}
	if v := f.Matcher; !dcl.IsEmptyValueIndirect(v) {
		m["matcher"] = v
	}

	return m, nil
}

// flattenUptimeCheckConfigContentMatchers flattens an instance of UptimeCheckConfigContentMatchers from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchers(c *Client, i interface{}) *UptimeCheckConfigContentMatchers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &UptimeCheckConfigContentMatchers{}
	r.Content = dcl.FlattenString(m["content"])
	r.Matcher = flattenUptimeCheckConfigContentMatchersMatcherEnum(m["matcher"])
	if dcl.IsEmptyValueIndirect(m["matcher"]) {
		c.Config.Logger.Info("Using default value for matcher.")
		r.Matcher = UptimeCheckConfigContentMatchersMatcherEnumRef("CONTAINS_STRING")
	}

	return r
}

// flattenUptimeCheckConfigResourceGroupResourceTypeEnumSlice flattens the contents of UptimeCheckConfigResourceGroupResourceTypeEnum from a JSON
// response object.
func flattenUptimeCheckConfigResourceGroupResourceTypeEnumSlice(c *Client, i interface{}) []UptimeCheckConfigResourceGroupResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigResourceGroupResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigResourceGroupResourceTypeEnum{}
	}

	items := make([]UptimeCheckConfigResourceGroupResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigResourceGroupResourceTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUptimeCheckConfigResourceGroupResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigResourceGroupResourceTypeEnum with the same value as that string.
func flattenUptimeCheckConfigResourceGroupResourceTypeEnum(i interface{}) *UptimeCheckConfigResourceGroupResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigResourceGroupResourceTypeEnumRef("")
	}

	return UptimeCheckConfigResourceGroupResourceTypeEnumRef(s)
}

// flattenUptimeCheckConfigHttpCheckRequestMethodEnumSlice flattens the contents of UptimeCheckConfigHttpCheckRequestMethodEnum from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckRequestMethodEnumSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckRequestMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckRequestMethodEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckRequestMethodEnum{}
	}

	items := make([]UptimeCheckConfigHttpCheckRequestMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckRequestMethodEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckRequestMethodEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigHttpCheckRequestMethodEnum with the same value as that string.
func flattenUptimeCheckConfigHttpCheckRequestMethodEnum(i interface{}) *UptimeCheckConfigHttpCheckRequestMethodEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigHttpCheckRequestMethodEnumRef("")
	}

	return UptimeCheckConfigHttpCheckRequestMethodEnumRef(s)
}

// flattenUptimeCheckConfigHttpCheckContentTypeEnumSlice flattens the contents of UptimeCheckConfigHttpCheckContentTypeEnum from a JSON
// response object.
func flattenUptimeCheckConfigHttpCheckContentTypeEnumSlice(c *Client, i interface{}) []UptimeCheckConfigHttpCheckContentTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigHttpCheckContentTypeEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigHttpCheckContentTypeEnum{}
	}

	items := make([]UptimeCheckConfigHttpCheckContentTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigHttpCheckContentTypeEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUptimeCheckConfigHttpCheckContentTypeEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigHttpCheckContentTypeEnum with the same value as that string.
func flattenUptimeCheckConfigHttpCheckContentTypeEnum(i interface{}) *UptimeCheckConfigHttpCheckContentTypeEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigHttpCheckContentTypeEnumRef("")
	}

	return UptimeCheckConfigHttpCheckContentTypeEnumRef(s)
}

// flattenUptimeCheckConfigContentMatchersMatcherEnumSlice flattens the contents of UptimeCheckConfigContentMatchersMatcherEnum from a JSON
// response object.
func flattenUptimeCheckConfigContentMatchersMatcherEnumSlice(c *Client, i interface{}) []UptimeCheckConfigContentMatchersMatcherEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []UptimeCheckConfigContentMatchersMatcherEnum{}
	}

	if len(a) == 0 {
		return []UptimeCheckConfigContentMatchersMatcherEnum{}
	}

	items := make([]UptimeCheckConfigContentMatchersMatcherEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenUptimeCheckConfigContentMatchersMatcherEnum(item.(map[string]interface{})))
	}

	return items
}

// flattenUptimeCheckConfigContentMatchersMatcherEnum asserts that an interface is a string, and returns a
// pointer to a *UptimeCheckConfigContentMatchersMatcherEnum with the same value as that string.
func flattenUptimeCheckConfigContentMatchersMatcherEnum(i interface{}) *UptimeCheckConfigContentMatchersMatcherEnum {
	s, ok := i.(string)
	if !ok {
		return UptimeCheckConfigContentMatchersMatcherEnumRef("")
	}

	return UptimeCheckConfigContentMatchersMatcherEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *UptimeCheckConfig) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalUptimeCheckConfig(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}
