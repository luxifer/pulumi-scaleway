// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecretVersion(ctx *pulumi.Context, args *LookupSecretVersionArgs, opts ...pulumi.InvokeOption) (*LookupSecretVersionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSecretVersionResult
	err := ctx.Invoke("scaleway:index/getSecretVersion:getSecretVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionArgs struct {
	// `region`) The region
	// in which the resource exists.
	Region *string `pulumi:"region"`
	// The revision for this Secret Version.
	Revision *string `pulumi:"revision"`
	// The Secret ID associated wit the secret version.
	// Only one of `secretId` and `secretName` should be specified.
	SecretId *string `pulumi:"secretId"`
	// The Name of Secret associated wit the secret version.
	// Only one of `secretId` and `secretName` should be specified.
	SecretName *string `pulumi:"secretName"`
}

// A collection of values returned by getSecretVersion.
type LookupSecretVersionResult struct {
	// Date and time of secret version's creation (RFC 3339 format).
	CreatedAt string `pulumi:"createdAt"`
	// The data payload of the secret version. more on the data section
	Data string `pulumi:"data"`
	// (Optional) Description of the secret version (e.g. `my-new-description`).
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Region     *string `pulumi:"region"`
	Revision   *string `pulumi:"revision"`
	SecretId   *string `pulumi:"secretId"`
	SecretName *string `pulumi:"secretName"`
	// The status of the Secret Version.
	Status string `pulumi:"status"`
	// Date and time of secret version's last update (RFC 3339 format).
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupSecretVersionOutput(ctx *pulumi.Context, args LookupSecretVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSecretVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretVersionResult, error) {
			args := v.(LookupSecretVersionArgs)
			r, err := LookupSecretVersion(ctx, &args, opts...)
			var s LookupSecretVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretVersionResultOutput)
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionOutputArgs struct {
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The revision for this Secret Version.
	Revision pulumi.StringPtrInput `pulumi:"revision"`
	// The Secret ID associated wit the secret version.
	// Only one of `secretId` and `secretName` should be specified.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
	// The Name of Secret associated wit the secret version.
	// Only one of `secretId` and `secretName` should be specified.
	SecretName pulumi.StringPtrInput `pulumi:"secretName"`
}

func (LookupSecretVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretVersionArgs)(nil)).Elem()
}

// A collection of values returned by getSecretVersion.
type LookupSecretVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSecretVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretVersionResult)(nil)).Elem()
}

func (o LookupSecretVersionResultOutput) ToLookupSecretVersionResultOutput() LookupSecretVersionResultOutput {
	return o
}

func (o LookupSecretVersionResultOutput) ToLookupSecretVersionResultOutputWithContext(ctx context.Context) LookupSecretVersionResultOutput {
	return o
}

// Date and time of secret version's creation (RFC 3339 format).
func (o LookupSecretVersionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The data payload of the secret version. more on the data section
func (o LookupSecretVersionResultOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Data }).(pulumi.StringOutput)
}

// (Optional) Description of the secret version (e.g. `my-new-description`).
func (o LookupSecretVersionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretVersionResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupSecretVersionResultOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) *string { return v.Revision }).(pulumi.StringPtrOutput)
}

func (o LookupSecretVersionResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSecretVersionResultOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) *string { return v.SecretName }).(pulumi.StringPtrOutput)
}

// The status of the Secret Version.
func (o LookupSecretVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Date and time of secret version's last update (RFC 3339 format).
func (o LookupSecretVersionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretVersionResultOutput{})
}