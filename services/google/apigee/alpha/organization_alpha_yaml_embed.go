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
// gen_go_data -package alpha -var YAML_organization blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/alpha/organization.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/alpha/organization.yaml
var YAML_organization = []byte("info:\n  title: Apigee/Organization\n  description: The Apigee Organization resource\n  x-dcl-struct-name: Organization\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Organization\n    parameters:\n    - name: Organization\n      required: true\n      description: A full instance of a Organization\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a Organization\n    parameters:\n    - name: Organization\n      required: true\n      description: A full instance of a Organization\n    timeoutSecs: 4800\n  delete:\n    description: The function used to delete a Organization\n    parameters:\n    - name: Organization\n      required: true\n      description: A full instance of a Organization\n    timeoutSecs: 4800\n  deleteAll:\n    description: The function used to delete all Organization\n    parameters: []\n    timeoutSecs: 4800\n  list:\n    description: The function used to list information about many Organization\n    parameters: []\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    Organization:\n      title: Organization\n      x-dcl-id: organizations/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - analyticsRegion\n      - runtimeType\n      properties:\n        addonsConfig:\n          type: object\n          x-dcl-go-name: AddonsConfig\n          x-dcl-go-type: OrganizationAddonsConfig\n          description: Addon configurations of the Apigee organization.\n          properties:\n            advancedApiOpsConfig:\n              type: object\n              x-dcl-go-name: AdvancedApiOpsConfig\n              x-dcl-go-type: OrganizationAddonsConfigAdvancedApiOpsConfig\n              description: Configuration for the Advanced API Ops add-on.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Flag that specifies whether the Advanced API Ops add-on\n                    is enabled.\n            monetizationConfig:\n              type: object\n              x-dcl-go-name: MonetizationConfig\n              x-dcl-go-type: OrganizationAddonsConfigMonetizationConfig\n              description: Configuration for the Monetization add-on.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Flag that specifies whether the Monetization add-on\n                    is enabled.\n        analyticsRegion:\n          type: string\n          x-dcl-go-name: AnalyticsRegion\n          description: Required. Primary GCP region for analytics data storage. For\n            valid values, see (https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).\n          x-kubernetes-immutable: true\n        authorizedNetwork:\n          type: string\n          x-dcl-go-name: AuthorizedNetwork\n          description: 'Compute Engine network used for Service Networking to be peered\n            with Apigee runtime instances. See (https://cloud.google.com/vpc/docs/shared-vpc).\n            To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`.\n            For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:**\n            Not supported for Apigee hybrid.'\n          x-dcl-references:\n          - resource: Compute/Network\n            field: name\n        billingType:\n          type: string\n          x-dcl-go-name: BillingType\n          x-dcl-go-type: OrganizationBillingTypeEnum\n          readOnly: true\n          description: 'Output only. Billing type of the Apigee organization. See\n            (https://cloud.google.com/apigee/pricing). Possible values: BILLING_TYPE_UNSPECIFIED,\n            SUBSCRIPTION, EVALUATION'\n          x-kubernetes-immutable: true\n          enum:\n          - BILLING_TYPE_UNSPECIFIED\n          - SUBSCRIPTION\n          - EVALUATION\n        caCertificate:\n          type: string\n          x-dcl-go-name: CaCertificate\n          readOnly: true\n          description: Output only. Base64-encoded public certificate for the root\n            CA of the Apigee organization. Valid only when (#RuntimeType) is `CLOUD`.\n          x-kubernetes-immutable: true\n        createdAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreatedAt\n          readOnly: true\n          description: Output only. Time that the Apigee organization was created\n            in milliseconds since epoch.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the Apigee organization.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Display name for the Apigee organization.\n        environments:\n          type: array\n          x-dcl-go-name: Environments\n          readOnly: true\n          description: Output only. List of environments in the Apigee organization.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        expiresAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: ExpiresAt\n          readOnly: true\n          description: Output only. Time that the Apigee organization is scheduled\n            for deletion.\n          x-kubernetes-immutable: true\n        lastModifiedAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedAt\n          readOnly: true\n          description: Output only. Time that the Apigee organization was last modified\n            in milliseconds since epoch.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. Name of the Apigee organization.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: 'Required. Name of the GCP project in which to associate the\n            Apigee organization. Pass the information as a query parameter using the\n            following structure in your request: projects/<project> Authorization\n            requires the following IAM permission on the specified resource parent:\n            apigee.organizations.create'\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        projectId:\n          type: string\n          x-dcl-go-name: ProjectId\n          readOnly: true\n          description: Output only. Project ID associated with the Apigee organization.\n          x-kubernetes-immutable: true\n        properties:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Properties\n          description: Properties defined in the Apigee organization profile.\n        runtimeDatabaseEncryptionKeyName:\n          type: string\n          x-dcl-go-name: RuntimeDatabaseEncryptionKeyName\n          description: 'Cloud KMS key name used for encrypting the data that is stored\n            and replicated across runtime instances. Update is not allowed after the\n            organization is created. Required when (#RuntimeType) is `TRIAL`, a Google-Managed\n            encryption key will be used. For example: \"projects/foo/locations/us/keyRings/bar/cryptoKeys/baz\".\n            **Note:** Not supported for Apigee hybrid.'\n          x-dcl-references:\n          - resource: Cloudkms/CryptoKey\n            field: name\n        runtimeType:\n          type: string\n          x-dcl-go-name: RuntimeType\n          x-dcl-go-type: OrganizationRuntimeTypeEnum\n          description: 'Required. Runtime type of the Apigee organization based on\n            the Apigee subscription purchased. Possible values: RUNTIME_TYPE_UNSPECIFIED,\n            CLOUD, HYBRID'\n          x-kubernetes-immutable: true\n          enum:\n          - RUNTIME_TYPE_UNSPECIFIED\n          - CLOUD\n          - HYBRID\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: OrganizationStateEnum\n          readOnly: true\n          description: 'Output only. State of the organization. Values other than\n            ACTIVE means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED,\n            MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - SNAPSHOT_STATE_UNSPECIFIED\n          - MISSING\n          - OK_DOCSTORE\n          - OK_SUBMITTED\n          - OK_EXTERNAL\n          - DELETED\n        subscriptionType:\n          type: string\n          x-dcl-go-name: SubscriptionType\n          x-dcl-go-type: OrganizationSubscriptionTypeEnum\n          readOnly: true\n          description: 'Output only. DEPRECATED: This will eventually be replaced\n            by BillingType. Subscription type of the Apigee organization. Valid values\n            include trial (free, limited, and for evaluation purposes only) or paid\n            (full subscription has been purchased). See (https://cloud.google.com/apigee/pricing/).\n            Possible values: SUBSCRIPTION_TYPE_UNSPECIFIED, PAID, TRIAL'\n          x-kubernetes-immutable: true\n          enum:\n          - SUBSCRIPTION_TYPE_UNSPECIFIED\n          - PAID\n          - TRIAL\n")

// 9607 bytes
// MD5: b274c92486a09f031f71845ccf0c5efe
