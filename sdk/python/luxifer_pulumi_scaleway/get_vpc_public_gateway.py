# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVpcPublicGatewayResult',
    'AwaitableGetVpcPublicGatewayResult',
    'get_vpc_public_gateway',
    'get_vpc_public_gateway_output',
]

@pulumi.output_type
class GetVpcPublicGatewayResult:
    """
    A collection of values returned by getVpcPublicGateway.
    """
    def __init__(__self__, bastion_enabled=None, bastion_port=None, created_at=None, enable_smtp=None, id=None, ip_id=None, name=None, organization_id=None, project_id=None, public_gateway_id=None, tags=None, type=None, updated_at=None, upstream_dns_servers=None, zone=None):
        if bastion_enabled and not isinstance(bastion_enabled, bool):
            raise TypeError("Expected argument 'bastion_enabled' to be a bool")
        pulumi.set(__self__, "bastion_enabled", bastion_enabled)
        if bastion_port and not isinstance(bastion_port, int):
            raise TypeError("Expected argument 'bastion_port' to be a int")
        pulumi.set(__self__, "bastion_port", bastion_port)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if enable_smtp and not isinstance(enable_smtp, bool):
            raise TypeError("Expected argument 'enable_smtp' to be a bool")
        pulumi.set(__self__, "enable_smtp", enable_smtp)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_id and not isinstance(ip_id, str):
            raise TypeError("Expected argument 'ip_id' to be a str")
        pulumi.set(__self__, "ip_id", ip_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public_gateway_id and not isinstance(public_gateway_id, str):
            raise TypeError("Expected argument 'public_gateway_id' to be a str")
        pulumi.set(__self__, "public_gateway_id", public_gateway_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if upstream_dns_servers and not isinstance(upstream_dns_servers, list):
            raise TypeError("Expected argument 'upstream_dns_servers' to be a list")
        pulumi.set(__self__, "upstream_dns_servers", upstream_dns_servers)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="bastionEnabled")
    def bastion_enabled(self) -> bool:
        return pulumi.get(self, "bastion_enabled")

    @property
    @pulumi.getter(name="bastionPort")
    def bastion_port(self) -> int:
        return pulumi.get(self, "bastion_port")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="enableSmtp")
    def enable_smtp(self) -> bool:
        return pulumi.get(self, "enable_smtp")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> str:
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicGatewayId")
    def public_gateway_id(self) -> Optional[str]:
        return pulumi.get(self, "public_gateway_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="upstreamDnsServers")
    def upstream_dns_servers(self) -> Sequence[str]:
        return pulumi.get(self, "upstream_dns_servers")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetVpcPublicGatewayResult(GetVpcPublicGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPublicGatewayResult(
            bastion_enabled=self.bastion_enabled,
            bastion_port=self.bastion_port,
            created_at=self.created_at,
            enable_smtp=self.enable_smtp,
            id=self.id,
            ip_id=self.ip_id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            public_gateway_id=self.public_gateway_id,
            tags=self.tags,
            type=self.type,
            updated_at=self.updated_at,
            upstream_dns_servers=self.upstream_dns_servers,
            zone=self.zone)


def get_vpc_public_gateway(name: Optional[str] = None,
                           public_gateway_id: Optional[str] = None,
                           zone: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPublicGatewayResult:
    """
    Gets information about a public gateway.

    ## Example Usage

    ```python
    import pulumi
    import luxifer_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.VpcPublicGateway("main",
        type="VPC-GW-S",
        zone="nl-ams-1")
    pg_test_by_name = scaleway.get_vpc_public_gateway_output(name=main.name,
        zone="nl-ams-1")
    pg_test_by_id = scaleway.get_vpc_public_gateway_output(public_gateway_id=main.id)
    ```


    :param str name: Exact name of the public gateway.
    :param str zone: `zone`) The zone in which
           the public gateway should be created.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['publicGatewayId'] = public_gateway_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPublicGateway:getVpcPublicGateway', __args__, opts=opts, typ=GetVpcPublicGatewayResult).value

    return AwaitableGetVpcPublicGatewayResult(
        bastion_enabled=__ret__.bastion_enabled,
        bastion_port=__ret__.bastion_port,
        created_at=__ret__.created_at,
        enable_smtp=__ret__.enable_smtp,
        id=__ret__.id,
        ip_id=__ret__.ip_id,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        project_id=__ret__.project_id,
        public_gateway_id=__ret__.public_gateway_id,
        tags=__ret__.tags,
        type=__ret__.type,
        updated_at=__ret__.updated_at,
        upstream_dns_servers=__ret__.upstream_dns_servers,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_vpc_public_gateway)
def get_vpc_public_gateway_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                  public_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  zone: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcPublicGatewayResult]:
    """
    Gets information about a public gateway.

    ## Example Usage

    ```python
    import pulumi
    import luxifer_pulumi_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.VpcPublicGateway("main",
        type="VPC-GW-S",
        zone="nl-ams-1")
    pg_test_by_name = scaleway.get_vpc_public_gateway_output(name=main.name,
        zone="nl-ams-1")
    pg_test_by_id = scaleway.get_vpc_public_gateway_output(public_gateway_id=main.id)
    ```


    :param str name: Exact name of the public gateway.
    :param str zone: `zone`) The zone in which
           the public gateway should be created.
    """
    ...
