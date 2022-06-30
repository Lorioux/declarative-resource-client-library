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
// gen_go_data -package alpha -var YAML_brand blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iap/alpha/brand.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iap/alpha/brand.yaml
var YAML_brand = []byte("info:\n  title: Iap/Brand\n  description: The Iap Brand resource\n  x-dcl-struct-name: Brand\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Brand\n    parameters:\n    - name: Brand\n      required: true\n      description: A full instance of a Brand\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a Brand\n    parameters:\n    - name: Brand\n      required: true\n      description: A full instance of a Brand\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many Brand\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    Brand:\n      title: Brand\n      x-dcl-id: projects/{{project}}/brands/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      properties:\n        applicationTitle:\n          type: string\n          x-dcl-go-name: ApplicationTitle\n          description: Application name displayed on OAuth consent screen.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. Identifier of the brand. NOTE: GCP project number\n            achieves the same brand identification purpose as only one brand per project\n            can be created.'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        orgInternalOnly:\n          type: boolean\n          x-dcl-go-name: OrgInternalOnly\n          readOnly: true\n          description: Output only. Whether the brand is only intended for usage inside\n            the G Suite organization only.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: GCP Project id under which the brand is to be created.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        supportEmail:\n          type: string\n          x-dcl-go-name: SupportEmail\n          description: Support email displayed on the OAuth consent screen.\n          x-kubernetes-immutable: true\n")

// 2308 bytes
// MD5: db06ac24708b1c4c9b59c6bf0a633f6e
