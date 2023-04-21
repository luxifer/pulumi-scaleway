# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqNamespaceArgs', 'MnqNamespace']

@pulumi.input_type
class MnqNamespaceArgs:
    def __init__(__self__, *,
                 protocol: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqNamespace resource.
        :param pulumi.Input[str] protocol: The protocol to apply on your namespace. Please check our
               supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        :param pulumi.Input[str] name: The unique name of the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region
               in which the namespace should be created.
        """
        pulumi.set(__self__, "protocol", protocol)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The protocol to apply on your namespace. Please check our
        supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the
        namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MnqNamespaceState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqNamespace resources.
        :param pulumi.Input[str] created_at: The date and time the Namespace was created.
        :param pulumi.Input[str] endpoint: The endpoint of the service matching the Namespace protocol.
        :param pulumi.Input[str] name: The unique name of the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               namespace is associated with.
        :param pulumi.Input[str] protocol: The protocol to apply on your namespace. Please check our
               supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        :param pulumi.Input[str] region: `region`). The region
               in which the namespace should be created.
        :param pulumi.Input[str] updated_at: The date and time the Namespace was updated.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Namespace was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the service matching the Namespace protocol.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the
        namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol to apply on your namespace. Please check our
        supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Namespace was updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class MnqNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and queuing Namespace.
        For further information please check
        our [documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-index)

        ## Examples

        ### Basic

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        main = scaleway.MnqNamespace("main",
            protocol="nats",
            region="fr-par")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/mnqNamespace:MnqNamespace main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The unique name of the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               namespace is associated with.
        :param pulumi.Input[str] protocol: The protocol to apply on your namespace. Please check our
               supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        :param pulumi.Input[str] region: `region`). The region
               in which the namespace should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and queuing Namespace.
        For further information please check
        our [documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-index)

        ## Examples

        ### Basic

        ```python
        import pulumi
        import luxifer_pulumi_scaleway as scaleway

        main = scaleway.MnqNamespace("main",
            protocol="nats",
            region="fr-par")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/mnqNamespace:MnqNamespace main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqNamespaceArgs.__new__(MnqNamespaceArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["created_at"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["updated_at"] = None
        super(MnqNamespace, __self__).__init__(
            'scaleway:index/mnqNamespace:MnqNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'MnqNamespace':
        """
        Get an existing MnqNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the Namespace was created.
        :param pulumi.Input[str] endpoint: The endpoint of the service matching the Namespace protocol.
        :param pulumi.Input[str] name: The unique name of the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               namespace is associated with.
        :param pulumi.Input[str] protocol: The protocol to apply on your namespace. Please check our
               supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        :param pulumi.Input[str] region: `region`). The region
               in which the namespace should be created.
        :param pulumi.Input[str] updated_at: The date and time the Namespace was updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqNamespaceState.__new__(_MnqNamespaceState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["updated_at"] = updated_at
        return MnqNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the Namespace was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint of the service matching the Namespace protocol.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the namespace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the
        namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol to apply on your namespace. Please check our
        supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time the Namespace was updated.
        """
        return pulumi.get(self, "updated_at")
