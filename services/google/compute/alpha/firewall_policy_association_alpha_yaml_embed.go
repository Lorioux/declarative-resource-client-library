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
// gen_go_data -package alpha -var YAML_firewall_policy_association blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/firewall_policy_association.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/firewall_policy_association.yaml
var YAML_firewall_policy_association = []byte("info:\n  title: Compute/FirewallPolicyAssociation\n  description: The Compute FirewallPolicyAssociation resource\n  x-dcl-struct-name: FirewallPolicyAssociation\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a FirewallPolicyAssociation\n    parameters:\n    - name: FirewallPolicyAssociation\n      required: true\n      description: A full instance of a FirewallPolicyAssociation\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a FirewallPolicyAssociation\n    parameters:\n    - name: FirewallPolicyAssociation\n      required: true\n      description: A full instance of a FirewallPolicyAssociation\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a FirewallPolicyAssociation\n    parameters:\n    - name: FirewallPolicyAssociation\n      required: true\n      description: A full instance of a FirewallPolicyAssociation\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all FirewallPolicyAssociation\n    parameters:\n    - name: firewallpolicy\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many FirewallPolicyAssociation\n    parameters:\n    - name: firewallpolicy\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    FirewallPolicyAssociation:\n      title: FirewallPolicyAssociation\n      x-dcl-id: locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - attachmentTarget\n      - firewallPolicy\n      properties:\n        attachmentTarget:\n          type: string\n          x-dcl-go-name: AttachmentTarget\n          description: The target that the firewall policy is attached to.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Folder\n            field: name\n          - resource: Cloudresourcemanager/Organization\n            field: name\n        firewallPolicy:\n          type: string\n          x-dcl-go-name: FirewallPolicy\n          description: The firewall policy ID of the association.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/FirewallPolicy\n            field: name\n            parent: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name for an association.\n          x-kubernetes-immutable: true\n        shortName:\n          type: string\n          x-dcl-go-name: ShortName\n          readOnly: true\n          description: The short name of the firewall policy of the association.\n          x-kubernetes-immutable: true\n")

// 2843 bytes
// MD5: 39f13031b1a4c9c149143b50f812eacc
