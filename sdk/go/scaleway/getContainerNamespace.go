// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a container namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/luxifer/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.LookupContainerNamespace(ctx, &GetContainerNamespaceArgs{
// 			Name: pulumi.StringRef("my-namespace-name"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.LookupContainerNamespace(ctx, &GetContainerNamespaceArgs{
// 			NamespaceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupContainerNamespace(ctx *pulumi.Context, args *LookupContainerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupContainerNamespaceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupContainerNamespaceResult
	err := ctx.Invoke("scaleway:index/getContainerNamespace:getContainerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name *string `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResult struct {
	// The description of the namespace.
	Description     string `pulumi:"description"`
	DestroyRegistry bool   `pulumi:"destroyRegistry"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Name        *string `pulumi:"name"`
	NamespaceId *string `pulumi:"namespaceId"`
	// The organization ID the namespace is associated with.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint string `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId        string            `pulumi:"registryNamespaceId"`
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

func LookupContainerNamespaceOutput(ctx *pulumi.Context, args LookupContainerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupContainerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerNamespaceResult, error) {
			args := v.(LookupContainerNamespaceArgs)
			r, err := LookupContainerNamespace(ctx, &args, opts...)
			var s LookupContainerNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerNamespaceResultOutput)
}

// A collection of arguments for invoking getContainerNamespace.
type LookupContainerNamespaceOutputArgs struct {
	// The namespace name.
	// Only one of `name` and `namespaceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The namespace id.
	// Only one of `name` and `namespaceId` should be specified.
	NamespaceId pulumi.StringPtrInput `pulumi:"namespaceId"`
	// `region`) The region in which the namespace exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupContainerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getContainerNamespace.
type LookupContainerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupContainerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerNamespaceResult)(nil)).Elem()
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutput() LookupContainerNamespaceResultOutput {
	return o
}

func (o LookupContainerNamespaceResultOutput) ToLookupContainerNamespaceResultOutputWithContext(ctx context.Context) LookupContainerNamespaceResultOutput {
	return o
}

// The description of the namespace.
func (o LookupContainerNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) DestroyRegistry() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) bool { return v.DestroyRegistry }).(pulumi.BoolOutput)
}

// The environment variables of the namespace.
func (o LookupContainerNamespaceResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupContainerNamespaceResultOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// The organization ID the namespace is associated with.
func (o LookupContainerNamespaceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The registry endpoint of the namespace.
func (o LookupContainerNamespaceResultOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

// The registry namespace ID of the namespace.
func (o LookupContainerNamespaceResultOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) string { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

func (o LookupContainerNamespaceResultOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerNamespaceResult) map[string]string { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerNamespaceResultOutput{})
}