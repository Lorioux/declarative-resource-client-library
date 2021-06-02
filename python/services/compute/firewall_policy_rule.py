# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.compute import firewall_policy_rule_pb2
from google3.cloud.graphite.mmv2.services.google.compute import (
    firewall_policy_rule_pb2_grpc,
)

from typing import List


class FirewallPolicyRule(object):
    def __init__(
        self,
        description: str = None,
        priority: int = None,
        match: dict = None,
        action: str = None,
        direction: str = None,
        target_resources: list = None,
        enable_logging: bool = None,
        rule_tuple_count: int = None,
        target_service_accounts: list = None,
        disabled: bool = None,
        kind: str = None,
        firewall_policy: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.priority = priority
        self.match = match
        self.action = action
        self.direction = direction
        self.target_resources = target_resources
        self.enable_logging = enable_logging
        self.target_service_accounts = target_service_accounts
        self.disabled = disabled
        self.kind = kind
        self.firewall_policy = firewall_policy
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_policy_rule_pb2_grpc.ComputeFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = firewall_policy_rule_pb2.ApplyComputeFirewallPolicyRuleRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if FirewallPolicyRuleMatch.to_proto(self.match):
            request.resource.match.CopyFrom(
                FirewallPolicyRuleMatch.to_proto(self.match)
            )
        else:
            request.resource.ClearField("match")
        if Primitive.to_proto(self.action):
            request.resource.action = Primitive.to_proto(self.action)

        if FirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            request.resource.direction = FirewallPolicyRuleDirectionEnum.to_proto(
                self.direction
            )

        if Primitive.to_proto(self.target_resources):
            request.resource.target_resources.extend(
                Primitive.to_proto(self.target_resources)
            )
        if Primitive.to_proto(self.enable_logging):
            request.resource.enable_logging = Primitive.to_proto(self.enable_logging)

        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeFirewallPolicyRule(request)
        self.description = Primitive.from_proto(response.description)
        self.priority = Primitive.from_proto(response.priority)
        self.match = FirewallPolicyRuleMatch.from_proto(response.match)
        self.action = Primitive.from_proto(response.action)
        self.direction = FirewallPolicyRuleDirectionEnum.from_proto(response.direction)
        self.target_resources = Primitive.from_proto(response.target_resources)
        self.enable_logging = Primitive.from_proto(response.enable_logging)
        self.rule_tuple_count = Primitive.from_proto(response.rule_tuple_count)
        self.target_service_accounts = Primitive.from_proto(
            response.target_service_accounts
        )
        self.disabled = Primitive.from_proto(response.disabled)
        self.kind = Primitive.from_proto(response.kind)
        self.firewall_policy = Primitive.from_proto(response.firewall_policy)

    def delete(self):
        stub = firewall_policy_rule_pb2_grpc.ComputeFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = firewall_policy_rule_pb2.DeleteComputeFirewallPolicyRuleRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.priority):
            request.resource.priority = Primitive.to_proto(self.priority)

        if FirewallPolicyRuleMatch.to_proto(self.match):
            request.resource.match.CopyFrom(
                FirewallPolicyRuleMatch.to_proto(self.match)
            )
        else:
            request.resource.ClearField("match")
        if Primitive.to_proto(self.action):
            request.resource.action = Primitive.to_proto(self.action)

        if FirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            request.resource.direction = FirewallPolicyRuleDirectionEnum.to_proto(
                self.direction
            )

        if Primitive.to_proto(self.target_resources):
            request.resource.target_resources.extend(
                Primitive.to_proto(self.target_resources)
            )
        if Primitive.to_proto(self.enable_logging):
            request.resource.enable_logging = Primitive.to_proto(self.enable_logging)

        if Primitive.to_proto(self.target_service_accounts):
            request.resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if Primitive.to_proto(self.firewall_policy):
            request.resource.firewall_policy = Primitive.to_proto(self.firewall_policy)

        response = stub.DeleteComputeFirewallPolicyRule(request)

    @classmethod
    def list(self, firewallPolicy, service_account_file=""):
        stub = firewall_policy_rule_pb2_grpc.ComputeFirewallPolicyRuleServiceStub(
            channel.Channel()
        )
        request = firewall_policy_rule_pb2.ListComputeFirewallPolicyRuleRequest()
        request.service_account_file = service_account_file
        request.FirewallPolicy = firewallPolicy

        return stub.ListComputeFirewallPolicyRule(request).items

    def to_proto(self):
        resource = firewall_policy_rule_pb2.ComputeFirewallPolicyRule()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.priority):
            resource.priority = Primitive.to_proto(self.priority)
        if FirewallPolicyRuleMatch.to_proto(self.match):
            resource.match.CopyFrom(FirewallPolicyRuleMatch.to_proto(self.match))
        else:
            resource.ClearField("match")
        if Primitive.to_proto(self.action):
            resource.action = Primitive.to_proto(self.action)
        if FirewallPolicyRuleDirectionEnum.to_proto(self.direction):
            resource.direction = FirewallPolicyRuleDirectionEnum.to_proto(
                self.direction
            )
        if Primitive.to_proto(self.target_resources):
            resource.target_resources.extend(Primitive.to_proto(self.target_resources))
        if Primitive.to_proto(self.enable_logging):
            resource.enable_logging = Primitive.to_proto(self.enable_logging)
        if Primitive.to_proto(self.target_service_accounts):
            resource.target_service_accounts.extend(
                Primitive.to_proto(self.target_service_accounts)
            )
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if Primitive.to_proto(self.kind):
            resource.kind = Primitive.to_proto(self.kind)
        if Primitive.to_proto(self.firewall_policy):
            resource.firewall_policy = Primitive.to_proto(self.firewall_policy)
        return resource


class FirewallPolicyRuleMatch(object):
    def __init__(
        self,
        src_ip_ranges: list = None,
        dest_ip_ranges: list = None,
        layer4_configs: list = None,
    ):
        self.src_ip_ranges = src_ip_ranges
        self.dest_ip_ranges = dest_ip_ranges
        self.layer4_configs = layer4_configs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firewall_policy_rule_pb2.ComputeFirewallPolicyRuleMatch()
        if Primitive.to_proto(resource.src_ip_ranges):
            res.src_ip_ranges.extend(Primitive.to_proto(resource.src_ip_ranges))
        if Primitive.to_proto(resource.dest_ip_ranges):
            res.dest_ip_ranges.extend(Primitive.to_proto(resource.dest_ip_ranges))
        if FirewallPolicyRuleMatchLayer4ConfigsArray.to_proto(resource.layer4_configs):
            res.layer4_configs.extend(
                FirewallPolicyRuleMatchLayer4ConfigsArray.to_proto(
                    resource.layer4_configs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirewallPolicyRuleMatch(
            src_ip_ranges=resource.src_ip_ranges,
            dest_ip_ranges=resource.dest_ip_ranges,
            layer4_configs=resource.layer4_configs,
        )


class FirewallPolicyRuleMatchArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirewallPolicyRuleMatch.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirewallPolicyRuleMatch.from_proto(i) for i in resources]


class FirewallPolicyRuleMatchLayer4Configs(object):
    def __init__(self, ip_protocol: str = None, ports: list = None):
        self.ip_protocol = ip_protocol
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = firewall_policy_rule_pb2.ComputeFirewallPolicyRuleMatchLayer4Configs()
        if Primitive.to_proto(resource.ip_protocol):
            res.ip_protocol = Primitive.to_proto(resource.ip_protocol)
        if Primitive.to_proto(resource.ports):
            res.ports.extend(Primitive.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FirewallPolicyRuleMatchLayer4Configs(
            ip_protocol=resource.ip_protocol, ports=resource.ports,
        )


class FirewallPolicyRuleMatchLayer4ConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FirewallPolicyRuleMatchLayer4Configs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FirewallPolicyRuleMatchLayer4Configs.from_proto(i) for i in resources]


class FirewallPolicyRuleDirectionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return firewall_policy_rule_pb2.ComputeFirewallPolicyRuleDirectionEnum.Value(
            "ComputeFirewallPolicyRuleDirectionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return firewall_policy_rule_pb2.ComputeFirewallPolicyRuleDirectionEnum.Name(
            resource
        )[len("ComputeFirewallPolicyRuleDirectionEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
