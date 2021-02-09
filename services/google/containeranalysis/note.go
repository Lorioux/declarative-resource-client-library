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
package containeranalysis

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Note struct {
	Name             *string            `json:"name"`
	ShortDescription *string            `json:"shortDescription"`
	LongDescription  *string            `json:"longDescription"`
	RelatedUrl       []NoteRelatedUrl   `json:"relatedUrl"`
	ExpirationTime   *string            `json:"expirationTime"`
	CreateTime       *string            `json:"createTime"`
	UpdateTime       *string            `json:"updateTime"`
	RelatedNoteNames []string           `json:"relatedNoteNames"`
	Vulnerability    *NoteVulnerability `json:"vulnerability"`
	Build            *NoteBuild         `json:"build"`
	Image            *NoteImage         `json:"image"`
	Package          *NotePackage       `json:"package"`
	Discovery        *NoteDiscovery     `json:"discovery"`
	Project          *string            `json:"project"`
}

func (r *Note) String() string {
	return dcl.SprintResource(r)
}

// The enum NoteVulnerabilitySeverityEnum.
type NoteVulnerabilitySeverityEnum string

// NoteVulnerabilitySeverityEnumRef returns a *NoteVulnerabilitySeverityEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilitySeverityEnumRef(s string) *NoteVulnerabilitySeverityEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilitySeverityEnum(s)
	return &v
}

func (v NoteVulnerabilitySeverityEnum) Validate() error {
	for _, s := range []string{"SEVERITY_UNSPECIFIED", "MINIMAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilitySeverityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsAffectedVersionStartKindEnum.
type NoteVulnerabilityDetailsAffectedVersionStartKindEnum string

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef returns a *NoteVulnerabilityDetailsAffectedVersionStartKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef(s string) *NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityDetailsAffectedVersionStartKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsAffectedVersionStartKindEnum) Validate() error {
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsAffectedVersionStartKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsAffectedVersionEndKindEnum.
type NoteVulnerabilityDetailsAffectedVersionEndKindEnum string

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef returns a *NoteVulnerabilityDetailsAffectedVersionEndKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef(s string) *NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityDetailsAffectedVersionEndKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsAffectedVersionEndKindEnum) Validate() error {
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsAffectedVersionEndKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsFixedVersionKindEnum.
type NoteVulnerabilityDetailsFixedVersionKindEnum string

// NoteVulnerabilityDetailsFixedVersionKindEnumRef returns a *NoteVulnerabilityDetailsFixedVersionKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsFixedVersionKindEnumRef(s string) *NoteVulnerabilityDetailsFixedVersionKindEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityDetailsFixedVersionKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsFixedVersionKindEnum) Validate() error {
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsFixedVersionKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AttackVectorEnum.
type NoteVulnerabilityCvssV3AttackVectorEnum string

// NoteVulnerabilityCvssV3AttackVectorEnumRef returns a *NoteVulnerabilityCvssV3AttackVectorEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AttackVectorEnumRef(s string) *NoteVulnerabilityCvssV3AttackVectorEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3AttackVectorEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AttackVectorEnum) Validate() error {
	for _, s := range []string{"ATTACK_VECTOR_UNSPECIFIED", "ATTACK_VECTOR_NETWORK", "ATTACK_VECTOR_ADJACENT", "ATTACK_VECTOR_LOCAL", "ATTACK_VECTOR_PHYSICAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AttackVectorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AttackComplexityEnum.
type NoteVulnerabilityCvssV3AttackComplexityEnum string

// NoteVulnerabilityCvssV3AttackComplexityEnumRef returns a *NoteVulnerabilityCvssV3AttackComplexityEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AttackComplexityEnumRef(s string) *NoteVulnerabilityCvssV3AttackComplexityEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3AttackComplexityEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AttackComplexityEnum) Validate() error {
	for _, s := range []string{"ATTACK_COMPLEXITY_UNSPECIFIED", "ATTACK_COMPLEXITY_LOW", "ATTACK_COMPLEXITY_HIGH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AttackComplexityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3PrivilegesRequiredEnum.
type NoteVulnerabilityCvssV3PrivilegesRequiredEnum string

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef returns a *NoteVulnerabilityCvssV3PrivilegesRequiredEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef(s string) *NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3PrivilegesRequiredEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3PrivilegesRequiredEnum) Validate() error {
	for _, s := range []string{"PRIVILEGES_REQUIRED_UNSPECIFIED", "PRIVILEGES_REQUIRED_NONE", "PRIVILEGES_REQUIRED_LOW", "PRIVILEGES_REQUIRED_HIGH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3PrivilegesRequiredEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3UserInteractionEnum.
type NoteVulnerabilityCvssV3UserInteractionEnum string

// NoteVulnerabilityCvssV3UserInteractionEnumRef returns a *NoteVulnerabilityCvssV3UserInteractionEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3UserInteractionEnumRef(s string) *NoteVulnerabilityCvssV3UserInteractionEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3UserInteractionEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3UserInteractionEnum) Validate() error {
	for _, s := range []string{"USER_INTERACTION_UNSPECIFIED", "USER_INTERACTION_NONE", "USER_INTERACTION_REQUIRED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3UserInteractionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3ScopeEnum.
type NoteVulnerabilityCvssV3ScopeEnum string

// NoteVulnerabilityCvssV3ScopeEnumRef returns a *NoteVulnerabilityCvssV3ScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3ScopeEnumRef(s string) *NoteVulnerabilityCvssV3ScopeEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3ScopeEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3ScopeEnum) Validate() error {
	for _, s := range []string{"SCOPE_UNSPECIFIED", "SCOPE_UNCHANGED", "SCOPE_CHANGED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3ScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3ConfidentialityImpactEnum.
type NoteVulnerabilityCvssV3ConfidentialityImpactEnum string

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef returns a *NoteVulnerabilityCvssV3ConfidentialityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef(s string) *NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3ConfidentialityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3ConfidentialityImpactEnum) Validate() error {
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3ConfidentialityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3IntegrityImpactEnum.
type NoteVulnerabilityCvssV3IntegrityImpactEnum string

// NoteVulnerabilityCvssV3IntegrityImpactEnumRef returns a *NoteVulnerabilityCvssV3IntegrityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3IntegrityImpactEnumRef(s string) *NoteVulnerabilityCvssV3IntegrityImpactEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3IntegrityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3IntegrityImpactEnum) Validate() error {
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3IntegrityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AvailabilityImpactEnum.
type NoteVulnerabilityCvssV3AvailabilityImpactEnum string

// NoteVulnerabilityCvssV3AvailabilityImpactEnumRef returns a *NoteVulnerabilityCvssV3AvailabilityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AvailabilityImpactEnumRef(s string) *NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	if s == "" {
		return nil
	}

	v := NoteVulnerabilityCvssV3AvailabilityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AvailabilityImpactEnum) Validate() error {
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AvailabilityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NotePackageDistributionArchitectureEnum.
type NotePackageDistributionArchitectureEnum string

// NotePackageDistributionArchitectureEnumRef returns a *NotePackageDistributionArchitectureEnum with the value of string s
// If the empty string is provided, nil is returned.
func NotePackageDistributionArchitectureEnumRef(s string) *NotePackageDistributionArchitectureEnum {
	if s == "" {
		return nil
	}

	v := NotePackageDistributionArchitectureEnum(s)
	return &v
}

func (v NotePackageDistributionArchitectureEnum) Validate() error {
	for _, s := range []string{"ARCHITECTURE_UNSPECIFIED", "X86", "X64"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NotePackageDistributionArchitectureEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NotePackageDistributionLatestVersionKindEnum.
type NotePackageDistributionLatestVersionKindEnum string

// NotePackageDistributionLatestVersionKindEnumRef returns a *NotePackageDistributionLatestVersionKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NotePackageDistributionLatestVersionKindEnumRef(s string) *NotePackageDistributionLatestVersionKindEnum {
	if s == "" {
		return nil
	}

	v := NotePackageDistributionLatestVersionKindEnum(s)
	return &v
}

func (v NotePackageDistributionLatestVersionKindEnum) Validate() error {
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NotePackageDistributionLatestVersionKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteDiscoveryAnalysisKindEnum.
type NoteDiscoveryAnalysisKindEnum string

// NoteDiscoveryAnalysisKindEnumRef returns a *NoteDiscoveryAnalysisKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteDiscoveryAnalysisKindEnumRef(s string) *NoteDiscoveryAnalysisKindEnum {
	if s == "" {
		return nil
	}

	v := NoteDiscoveryAnalysisKindEnum(s)
	return &v
}

func (v NoteDiscoveryAnalysisKindEnum) Validate() error {
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteDiscoveryAnalysisKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type NoteRelatedUrl struct {
	empty bool    `json:"-"`
	Url   *string `json:"url"`
	Label *string `json:"label"`
}

// This object is used to assert a desired state where this NoteRelatedUrl is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteRelatedUrl *NoteRelatedUrl = &NoteRelatedUrl{empty: true}

func (r *NoteRelatedUrl) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteRelatedUrl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerability struct {
	empty            bool                              `json:"-"`
	CvssScore        *float64                          `json:"cvssScore"`
	Severity         *NoteVulnerabilitySeverityEnum    `json:"severity"`
	Details          []NoteVulnerabilityDetails        `json:"details"`
	CvssV3           *NoteVulnerabilityCvssV3          `json:"cvssV3"`
	WindowsDetails   []NoteVulnerabilityWindowsDetails `json:"windowsDetails"`
	SourceUpdateTime *string                           `json:"sourceUpdateTime"`
}

// This object is used to assert a desired state where this NoteVulnerability is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerability *NoteVulnerability = &NoteVulnerability{empty: true}

func (r *NoteVulnerability) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerability) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetails struct {
	empty                bool                                          `json:"-"`
	SeverityName         *string                                       `json:"severityName"`
	Description          *string                                       `json:"description"`
	PackageType          *string                                       `json:"packageType"`
	AffectedCpeUri       *string                                       `json:"affectedCpeUri"`
	AffectedPackage      *string                                       `json:"affectedPackage"`
	AffectedVersionStart *NoteVulnerabilityDetailsAffectedVersionStart `json:"affectedVersionStart"`
	AffectedVersionEnd   *NoteVulnerabilityDetailsAffectedVersionEnd   `json:"affectedVersionEnd"`
	FixedCpeUri          *string                                       `json:"fixedCpeUri"`
	FixedPackage         *string                                       `json:"fixedPackage"`
	FixedVersion         *NoteVulnerabilityDetailsFixedVersion         `json:"fixedVersion"`
	IsObsolete           *bool                                         `json:"isObsolete"`
	SourceUpdateTime     *string                                       `json:"sourceUpdateTime"`
}

// This object is used to assert a desired state where this NoteVulnerabilityDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetails *NoteVulnerabilityDetails = &NoteVulnerabilityDetails{empty: true}

func (r *NoteVulnerabilityDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsAffectedVersionStart struct {
	empty    bool                                                  `json:"-"`
	Epoch    *int64                                                `json:"epoch"`
	Name     *string                                               `json:"name"`
	Revision *string                                               `json:"revision"`
	Kind     *NoteVulnerabilityDetailsAffectedVersionStartKindEnum `json:"kind"`
	FullName *string                                               `json:"fullName"`
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsAffectedVersionStart is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsAffectedVersionStart *NoteVulnerabilityDetailsAffectedVersionStart = &NoteVulnerabilityDetailsAffectedVersionStart{empty: true}

func (r *NoteVulnerabilityDetailsAffectedVersionStart) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsAffectedVersionStart) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsAffectedVersionEnd struct {
	empty    bool                                                `json:"-"`
	Epoch    *int64                                              `json:"epoch"`
	Name     *string                                             `json:"name"`
	Revision *string                                             `json:"revision"`
	Kind     *NoteVulnerabilityDetailsAffectedVersionEndKindEnum `json:"kind"`
	FullName *string                                             `json:"fullName"`
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsAffectedVersionEnd is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsAffectedVersionEnd *NoteVulnerabilityDetailsAffectedVersionEnd = &NoteVulnerabilityDetailsAffectedVersionEnd{empty: true}

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsFixedVersion struct {
	empty    bool                                          `json:"-"`
	Epoch    *int64                                        `json:"epoch"`
	Name     *string                                       `json:"name"`
	Revision *string                                       `json:"revision"`
	Kind     *NoteVulnerabilityDetailsFixedVersionKindEnum `json:"kind"`
	FullName *string                                       `json:"fullName"`
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsFixedVersion is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsFixedVersion *NoteVulnerabilityDetailsFixedVersion = &NoteVulnerabilityDetailsFixedVersion{empty: true}

func (r *NoteVulnerabilityDetailsFixedVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsFixedVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityCvssV3 struct {
	empty                 bool                                              `json:"-"`
	BaseScore             *float64                                          `json:"baseScore"`
	ExploitabilityScore   *float64                                          `json:"exploitabilityScore"`
	ImpactScore           *float64                                          `json:"impactScore"`
	AttackVector          *NoteVulnerabilityCvssV3AttackVectorEnum          `json:"attackVector"`
	AttackComplexity      *NoteVulnerabilityCvssV3AttackComplexityEnum      `json:"attackComplexity"`
	PrivilegesRequired    *NoteVulnerabilityCvssV3PrivilegesRequiredEnum    `json:"privilegesRequired"`
	UserInteraction       *NoteVulnerabilityCvssV3UserInteractionEnum       `json:"userInteraction"`
	Scope                 *NoteVulnerabilityCvssV3ScopeEnum                 `json:"scope"`
	ConfidentialityImpact *NoteVulnerabilityCvssV3ConfidentialityImpactEnum `json:"confidentialityImpact"`
	IntegrityImpact       *NoteVulnerabilityCvssV3IntegrityImpactEnum       `json:"integrityImpact"`
	AvailabilityImpact    *NoteVulnerabilityCvssV3AvailabilityImpactEnum    `json:"availabilityImpact"`
}

// This object is used to assert a desired state where this NoteVulnerabilityCvssV3 is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityCvssV3 *NoteVulnerabilityCvssV3 = &NoteVulnerabilityCvssV3{empty: true}

func (r *NoteVulnerabilityCvssV3) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityCvssV3) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityWindowsDetails struct {
	empty       bool                                       `json:"-"`
	CpeUri      *string                                    `json:"cpeUri"`
	Name        *string                                    `json:"name"`
	Description *string                                    `json:"description"`
	FixingKbs   []NoteVulnerabilityWindowsDetailsFixingKbs `json:"fixingKbs"`
}

// This object is used to assert a desired state where this NoteVulnerabilityWindowsDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityWindowsDetails *NoteVulnerabilityWindowsDetails = &NoteVulnerabilityWindowsDetails{empty: true}

func (r *NoteVulnerabilityWindowsDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityWindowsDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityWindowsDetailsFixingKbs struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

// This object is used to assert a desired state where this NoteVulnerabilityWindowsDetailsFixingKbs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityWindowsDetailsFixingKbs *NoteVulnerabilityWindowsDetailsFixingKbs = &NoteVulnerabilityWindowsDetailsFixingKbs{empty: true}

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteBuild struct {
	empty          bool    `json:"-"`
	BuilderVersion *string `json:"builderVersion"`
}

// This object is used to assert a desired state where this NoteBuild is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteBuild *NoteBuild = &NoteBuild{empty: true}

func (r *NoteBuild) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteBuild) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteImage struct {
	empty       bool                  `json:"-"`
	ResourceUrl *string               `json:"resourceUrl"`
	Fingerprint *NoteImageFingerprint `json:"fingerprint"`
}

// This object is used to assert a desired state where this NoteImage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteImage *NoteImage = &NoteImage{empty: true}

func (r *NoteImage) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteImage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteImageFingerprint struct {
	empty  bool     `json:"-"`
	V1Name *string  `json:"v1Name"`
	V2Blob []string `json:"v2Blob"`
	V2Name *string  `json:"v2Name"`
}

// This object is used to assert a desired state where this NoteImageFingerprint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteImageFingerprint *NoteImageFingerprint = &NoteImageFingerprint{empty: true}

func (r *NoteImageFingerprint) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteImageFingerprint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackage struct {
	empty        bool                      `json:"-"`
	Name         *string                   `json:"name"`
	Distribution []NotePackageDistribution `json:"distribution"`
}

// This object is used to assert a desired state where this NotePackage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNotePackage *NotePackage = &NotePackage{empty: true}

func (r *NotePackage) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackageDistribution struct {
	empty         bool                                     `json:"-"`
	CpeUri        *string                                  `json:"cpeUri"`
	Architecture  *NotePackageDistributionArchitectureEnum `json:"architecture"`
	LatestVersion *NotePackageDistributionLatestVersion    `json:"latestVersion"`
	Maintainer    *string                                  `json:"maintainer"`
	Url           *string                                  `json:"url"`
	Description   *string                                  `json:"description"`
}

// This object is used to assert a desired state where this NotePackageDistribution is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNotePackageDistribution *NotePackageDistribution = &NotePackageDistribution{empty: true}

func (r *NotePackageDistribution) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackageDistribution) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackageDistributionLatestVersion struct {
	empty    bool                                          `json:"-"`
	Epoch    *int64                                        `json:"epoch"`
	Name     *string                                       `json:"name"`
	Revision *string                                       `json:"revision"`
	Kind     *NotePackageDistributionLatestVersionKindEnum `json:"kind"`
	FullName *string                                       `json:"fullName"`
}

// This object is used to assert a desired state where this NotePackageDistributionLatestVersion is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNotePackageDistributionLatestVersion *NotePackageDistributionLatestVersion = &NotePackageDistributionLatestVersion{empty: true}

func (r *NotePackageDistributionLatestVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackageDistributionLatestVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteDiscovery struct {
	empty        bool                           `json:"-"`
	AnalysisKind *NoteDiscoveryAnalysisKindEnum `json:"analysisKind"`
}

// This object is used to assert a desired state where this NoteDiscovery is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyNoteDiscovery *NoteDiscovery = &NoteDiscovery{empty: true}

func (r *NoteDiscovery) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteDiscovery) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Note) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "containeranalysis",
		Type:    "Note",
		Version: "containeranalysis",
	}
}

const NoteMaxPage = -1

type NoteList struct {
	Items []*Note

	nextToken string

	pageSize int32

	project string
}

func (l *NoteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NoteList) Next(ctx context.Context, c *Client) error {
	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNote(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNote(ctx context.Context, project string) (*NoteList, error) {

	return c.ListNoteWithMaxResults(ctx, project, NoteMaxPage)

}

func (c *Client) ListNoteWithMaxResults(ctx context.Context, project string, pageSize int32) (*NoteList, error) {
	items, token, err := c.listNote(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NoteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetNote(ctx context.Context, r *Note) (*Note, error) {
	b, err := c.getNoteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNote(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNoteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNote(ctx context.Context, r *Note) error {
	if r == nil {
		return fmt.Errorf("Note resource is nil")
	}
	c.Config.Logger.Info("Deleting Note...")
	deleteOp := deleteNoteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNote deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNote(ctx context.Context, project string, filter func(*Note) bool) error {
	listObj, err := c.ListNote(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllNote(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNote(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNote(ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (*Note, error) {
	c.Config.Logger.Info("Beginning ApplyNote...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.noteDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []noteApiOperation
	if create {
		ops = append(ops, &createNoteOperation{})
	} else if recreate {

		ops = append(ops, &deleteNoteOperation{})

		ops = append(ops, &createNoteOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeNoteDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetNote(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNoteNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNoteDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNote(c, newDesired, newState)
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
