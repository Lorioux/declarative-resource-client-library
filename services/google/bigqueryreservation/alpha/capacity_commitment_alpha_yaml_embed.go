// Copyright 2022 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_capacity_commitment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/alpha/capacity_commitment.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/alpha/capacity_commitment.yaml
var YAML_capacity_commitment = []byte("info:\n  title: BigqueryReservation/CapacityCommitment\n  description: The BigqueryReservation CapacityCommitment resource\n  x-dcl-struct-name: CapacityCommitment\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a CapacityCommitment\n    parameters:\n    - name: CapacityCommitment\n      required: true\n      description: A full instance of a CapacityCommitment\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a CapacityCommitment\n    parameters:\n    - name: CapacityCommitment\n      required: true\n      description: A full instance of a CapacityCommitment\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a CapacityCommitment\n    parameters:\n    - name: CapacityCommitment\n      required: true\n      description: A full instance of a CapacityCommitment\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all CapacityCommitment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many CapacityCommitment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    CapacityCommitment:\n      title: CapacityCommitment\n      x-dcl-id: projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - project\n      - location\n      properties:\n        commitmentEndTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CommitmentEndTime\n          readOnly: true\n          description: Output only. The end of the current commitment period. It is\n            applicable only for ACTIVE capacity commitments.\n          x-kubernetes-immutable: true\n        commitmentStartTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CommitmentStartTime\n          readOnly: true\n          description: Output only. The start of the current commitment period. It\n            is applicable only for ACTIVE capacity commitments.\n          x-kubernetes-immutable: true\n        failureStatus:\n          type: object\n          x-dcl-go-name: FailureStatus\n          x-dcl-go-type: CapacityCommitmentFailureStatus\n          readOnly: true\n          description: Output only. For FAILED commitment plan, provides the reason\n            of failure.\n          x-kubernetes-immutable: true\n          properties:\n            code:\n              type: integer\n              format: int64\n              x-dcl-go-name: Code\n              description: The status code, which should be an enum value of google.rpc.Code.\n              x-kubernetes-immutable: true\n            details:\n              type: array\n              x-dcl-go-name: Details\n              description: A list of messages that carry the error details. There\n                is a common set of message types for APIs to use.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: CapacityCommitmentFailureStatusDetails\n                properties:\n                  typeUrl:\n                    type: string\n                    x-dcl-go-name: TypeUrl\n                    description: 'A URL/resource name that uniquely identifies the\n                      type of the serialized protocol buffer message. This string\n                      must contain at least one \"/\" character. The last segment of\n                      the URL''s path must represent the fully qualified name of the\n                      type (as in `path/google.protobuf.Duration`). The name should\n                      be in a canonical form (e.g., leading \".\" is not accepted).\n                      In practice, teams usually precompile into the binary all types\n                      that they expect it to use in the context of Any. However, for\n                      URLs which use the scheme `http`, `https`, or no scheme, one\n                      can optionally set up a type server that maps type URLs to message\n                      definitions as follows: * If no scheme is provided, `https`\n                      is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type\n                      value in binary format, or produce an error. * Applications\n                      are allowed to cache lookup results based on the URL, or have\n                      them precompiled into a binary to avoid any lookup. Therefore,\n                      binary compatibility needs to be preserved on changes to types.\n                      (Use versioned type names to manage breaking changes.) Note:\n                      this functionality is not currently available in the official\n                      protobuf release, and it is not used for type URLs beginning\n                      with type.googleapis.com. Schemes other than `http`, `https`\n                      (or the empty scheme) might be used with implementation specific\n                      semantics.'\n                    x-kubernetes-immutable: true\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: Must be a valid serialized protocol buffer of the\n                      above specified type.\n                    x-kubernetes-immutable: true\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: A developer-facing error message, which should be in English.\n                Any user-facing error message should be localized and sent in the\n                google.rpc.Status.details field, or localized by the client.\n              x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123`\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        plan:\n          type: string\n          x-dcl-go-name: Plan\n          x-dcl-go-type: CapacityCommitmentPlanEnum\n          description: 'Capacity commitment commitment plan. Possible values: COMMITMENT_PLAN_UNSPECIFIED,\n            FLEX, TRIAL, MONTHLY, ANNUAL'\n          enum:\n          - COMMITMENT_PLAN_UNSPECIFIED\n          - FLEX\n          - TRIAL\n          - MONTHLY\n          - ANNUAL\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        renewalPlan:\n          type: string\n          x-dcl-go-name: RenewalPlan\n          x-dcl-go-type: CapacityCommitmentRenewalPlanEnum\n          description: 'The plan this capacity commitment is converted to after commitment_end_time\n            passes. Once the plan is changed, committed period is extended according\n            to commitment plan. Only applicable for ANNUAL and TRIAL commitments.\n            Possible values: COMMITMENT_PLAN_UNSPECIFIED, FLEX, TRIAL, MONTHLY, ANNUAL'\n          enum:\n          - COMMITMENT_PLAN_UNSPECIFIED\n          - FLEX\n          - TRIAL\n          - MONTHLY\n          - ANNUAL\n        slotCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: SlotCount\n          description: Number of slots in this commitment.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: CapacityCommitmentStateEnum\n          readOnly: true\n          description: 'Output only. State of the commitment. Possible values: STATE_UNSPECIFIED,\n            PENDING, ACTIVE, FAILED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PENDING\n          - ACTIVE\n          - FAILED\n")

// 8532 bytes
// MD5: 97787519ba1b1d32dc25f64b7d2acc16
