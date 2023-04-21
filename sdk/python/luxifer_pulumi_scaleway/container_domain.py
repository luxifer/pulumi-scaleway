# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ContainerDomainArgs', 'ContainerDomain']

@pulumi.input_type
class ContainerDomainArgs:
    def __init__(__self__, *,
                 container_id: pulumi.Input[str],
                 hostname: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ContainerDomain resource.
        :param pulumi.Input[str] container_id: The ID of the container.
        :param pulumi.Input[str] hostname: The hostname with a CNAME record.
        :param pulumi.Input[str] region: `region`) The region in which the container exists
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "hostname", hostname)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[str]:
        """
        The ID of the container.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        """
        The hostname with a CNAME record.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the container exists
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ContainerDomainState:
    def __init__(__self__, *,
                 container_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContainerDomain resources.
        :param pulumi.Input[str] container_id: The ID of the container.
        :param pulumi.Input[str] hostname: The hostname with a CNAME record.
        :param pulumi.Input[str] region: `region`) The region in which the container exists
        :param pulumi.Input[str] url: The URL used to query the container
        """
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname with a CNAME record.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the container exists
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL used to query the container
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ContainerDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Container domain name bindings.
        You can check our [containers guide](https://www.scaleway.com/en/docs/compute/containers/how-to/add-a-custom-domain-to-a-container/) for further information.

        ## Example Usage
        ### Simple

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        app_container = scaleway.Container("appContainer")
        app_container_domain = scaleway.ContainerDomain("appContainerDomain",
            container_id=app_container.id,
            hostname="container.domain.tld")
        ```
        ### Complete example with domain

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        main = scaleway.ContainerNamespace("main", description="test container")
        app_container = scaleway.Container("appContainer",
            namespace_id=main.id,
            registry_image=main.registry_endpoint.apply(lambda registry_endpoint: f"{registry_endpoint}/nginx:alpine"),
            port=80,
            cpu_limit=140,
            memory_limit=256,
            min_scale=1,
            max_scale=1,
            timeout=600,
            max_concurrency=80,
            privacy="public",
            protocol="http1",
            deploy=True)
        app_domain_record = scaleway.DomainRecord("appDomainRecord",
            dns_zone="domain.tld",
            type="CNAME",
            data=app_container.domain_name.apply(lambda domain_name: f"{domain_name}."),
            ttl=3600)
        app_container_domain = scaleway.ContainerDomain("appContainerDomain",
            container_id=app_container.id,
            hostname=pulumi.Output.all(app_domain_record.name, app_domain_record.dns_zone).apply(lambda name, dns_zone: f"{name}.{dns_zone}"))
        ```

        ## Import

        Container domain binding can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/containerDomain:ContainerDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container.
        :param pulumi.Input[str] hostname: The hostname with a CNAME record.
        :param pulumi.Input[str] region: `region`) The region in which the container exists
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Container domain name bindings.
        You can check our [containers guide](https://www.scaleway.com/en/docs/compute/containers/how-to/add-a-custom-domain-to-a-container/) for further information.

        ## Example Usage
        ### Simple

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        app_container = scaleway.Container("appContainer")
        app_container_domain = scaleway.ContainerDomain("appContainerDomain",
            container_id=app_container.id,
            hostname="container.domain.tld")
        ```
        ### Complete example with domain

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        main = scaleway.ContainerNamespace("main", description="test container")
        app_container = scaleway.Container("appContainer",
            namespace_id=main.id,
            registry_image=main.registry_endpoint.apply(lambda registry_endpoint: f"{registry_endpoint}/nginx:alpine"),
            port=80,
            cpu_limit=140,
            memory_limit=256,
            min_scale=1,
            max_scale=1,
            timeout=600,
            max_concurrency=80,
            privacy="public",
            protocol="http1",
            deploy=True)
        app_domain_record = scaleway.DomainRecord("appDomainRecord",
            dns_zone="domain.tld",
            type="CNAME",
            data=app_container.domain_name.apply(lambda domain_name: f"{domain_name}."),
            ttl=3600)
        app_container_domain = scaleway.ContainerDomain("appContainerDomain",
            container_id=app_container.id,
            hostname=pulumi.Output.all(app_domain_record.name, app_domain_record.dns_zone).apply(lambda name, dns_zone: f"{name}.{dns_zone}"))
        ```

        ## Import

        Container domain binding can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/containerDomain:ContainerDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ContainerDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerDomainArgs.__new__(ContainerDomainArgs)

            if container_id is None and not opts.urn:
                raise TypeError("Missing required property 'container_id'")
            __props__.__dict__["container_id"] = container_id
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["region"] = region
            __props__.__dict__["url"] = None
        super(ContainerDomain, __self__).__init__(
            'scaleway:index/containerDomain:ContainerDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_id: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ContainerDomain':
        """
        Get an existing ContainerDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container.
        :param pulumi.Input[str] hostname: The hostname with a CNAME record.
        :param pulumi.Input[str] region: `region`) The region in which the container exists
        :param pulumi.Input[str] url: The URL used to query the container
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerDomainState.__new__(_ContainerDomainState)

        __props__.__dict__["container_id"] = container_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["region"] = region
        __props__.__dict__["url"] = url
        return ContainerDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[str]:
        """
        The ID of the container.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The hostname with a CNAME record.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the container exists
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL used to query the container
        """
        return pulumi.get(self, "url")

