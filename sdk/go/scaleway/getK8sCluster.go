// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Kubernetes Cluster.
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
// 		_, err := scaleway.LookupK8sCluster(ctx, &GetK8sClusterArgs{
// 			ClusterId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupK8sCluster(ctx *pulumi.Context, args *LookupK8sClusterArgs, opts ...pulumi.InvokeOption) (*LookupK8sClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupK8sClusterResult
	err := ctx.Invoke("scaleway:index/getK8sCluster:getK8sCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getK8sCluster.
type LookupK8sClusterArgs struct {
	// The cluster ID. Only one of `name` and `clusterId` should be specified.
	ClusterId *string `pulumi:"clusterId"`
	// The cluster name. Only one of `name` and `clusterId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The region in which the cluster exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getK8sCluster.
type LookupK8sClusterResult struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) enabled on the cluster.
	AdmissionPlugins  []string `pulumi:"admissionPlugins"`
	ApiserverCertSans []string `pulumi:"apiserverCertSans"`
	// The URL of the Kubernetes API server.
	ApiserverUrl string `pulumi:"apiserverUrl"`
	// The auto upgrade configuration.
	AutoUpgrades []GetK8sClusterAutoUpgrade `pulumi:"autoUpgrades"`
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfigs []GetK8sClusterAutoscalerConfig `pulumi:"autoscalerConfigs"`
	ClusterId         *string                         `pulumi:"clusterId"`
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	Cni string `pulumi:"cni"`
	// The creation date of the cluster.
	CreatedAt string `pulumi:"createdAt"`
	// A description for the Kubernetes cluster.
	Description string `pulumi:"description"`
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) enabled on the cluster.
	FeatureGates []string `pulumi:"featureGates"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                             `pulumi:"id"`
	Kubeconfigs          []GetK8sClusterKubeconfig          `pulumi:"kubeconfigs"`
	Name                 *string                            `pulumi:"name"`
	OpenIdConnectConfigs []GetK8sClusterOpenIdConnectConfig `pulumi:"openIdConnectConfigs"`
	// The ID of the organization the cluster is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the private network of the cluster.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The ID of the project the cluster is associated with.
	ProjectId string `pulumi:"projectId"`
	// The region in which the cluster is.
	Region *string `pulumi:"region"`
	// The status of the Kubernetes cluster.
	Status string `pulumi:"status"`
	// The tags associated with the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
	// The type of the Kubernetes cluster.
	Type string `pulumi:"type"`
	// The last update date of the cluster.
	UpdatedAt string `pulumi:"updatedAt"`
	// True if a newer Kubernetes version is available.
	UpgradeAvailable bool `pulumi:"upgradeAvailable"`
	// The version of the Kubernetes cluster.
	Version string `pulumi:"version"`
	// The DNS wildcard that points to all ready nodes.
	WildcardDns string `pulumi:"wildcardDns"`
}

func LookupK8sClusterOutput(ctx *pulumi.Context, args LookupK8sClusterOutputArgs, opts ...pulumi.InvokeOption) LookupK8sClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupK8sClusterResult, error) {
			args := v.(LookupK8sClusterArgs)
			r, err := LookupK8sCluster(ctx, &args, opts...)
			var s LookupK8sClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupK8sClusterResultOutput)
}

// A collection of arguments for invoking getK8sCluster.
type LookupK8sClusterOutputArgs struct {
	// The cluster ID. Only one of `name` and `clusterId` should be specified.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The cluster name. Only one of `name` and `clusterId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the cluster exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupK8sClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sClusterArgs)(nil)).Elem()
}

// A collection of values returned by getK8sCluster.
type LookupK8sClusterResultOutput struct{ *pulumi.OutputState }

func (LookupK8sClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sClusterResult)(nil)).Elem()
}

func (o LookupK8sClusterResultOutput) ToLookupK8sClusterResultOutput() LookupK8sClusterResultOutput {
	return o
}

func (o LookupK8sClusterResultOutput) ToLookupK8sClusterResultOutputWithContext(ctx context.Context) LookupK8sClusterResultOutput {
	return o
}

// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) enabled on the cluster.
func (o LookupK8sClusterResultOutput) AdmissionPlugins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.AdmissionPlugins }).(pulumi.StringArrayOutput)
}

func (o LookupK8sClusterResultOutput) ApiserverCertSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.ApiserverCertSans }).(pulumi.StringArrayOutput)
}

// The URL of the Kubernetes API server.
func (o LookupK8sClusterResultOutput) ApiserverUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.ApiserverUrl }).(pulumi.StringOutput)
}

// The auto upgrade configuration.
func (o LookupK8sClusterResultOutput) AutoUpgrades() GetK8sClusterAutoUpgradeArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterAutoUpgrade { return v.AutoUpgrades }).(GetK8sClusterAutoUpgradeArrayOutput)
}

// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
func (o LookupK8sClusterResultOutput) AutoscalerConfigs() GetK8sClusterAutoscalerConfigArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterAutoscalerConfig { return v.AutoscalerConfigs }).(GetK8sClusterAutoscalerConfigArrayOutput)
}

func (o LookupK8sClusterResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The Container Network Interface (CNI) for the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Cni() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Cni }).(pulumi.StringOutput)
}

// The creation date of the cluster.
func (o LookupK8sClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description for the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) enabled on the cluster.
func (o LookupK8sClusterResultOutput) FeatureGates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.FeatureGates }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupK8sClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupK8sClusterResultOutput) Kubeconfigs() GetK8sClusterKubeconfigArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterKubeconfig { return v.Kubeconfigs }).(GetK8sClusterKubeconfigArrayOutput)
}

func (o LookupK8sClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupK8sClusterResultOutput) OpenIdConnectConfigs() GetK8sClusterOpenIdConnectConfigArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterOpenIdConnectConfig { return v.OpenIdConnectConfigs }).(GetK8sClusterOpenIdConnectConfigArrayOutput)
}

// The ID of the organization the cluster is associated with.
func (o LookupK8sClusterResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the private network of the cluster.
func (o LookupK8sClusterResultOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// The ID of the project the cluster is associated with.
func (o LookupK8sClusterResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which the cluster is.
func (o LookupK8sClusterResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The status of the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The last update date of the cluster.
func (o LookupK8sClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// True if a newer Kubernetes version is available.
func (o LookupK8sClusterResultOutput) UpgradeAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) bool { return v.UpgradeAvailable }).(pulumi.BoolOutput)
}

// The version of the Kubernetes cluster.
func (o LookupK8sClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

// The DNS wildcard that points to all ready nodes.
func (o LookupK8sClusterResultOutput) WildcardDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.WildcardDns }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupK8sClusterResultOutput{})
}