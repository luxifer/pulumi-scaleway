// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Scaleway Load-Balancer Certificates.
//
// This data source can prove useful when a module accepts an LB Certificate as an input variable and needs to, for example, determine the security of a certificate for your LB Frontend associated with your domain, etc.
//
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#certificate-330754).
//
// ## Examples
func LookupLbCertificate(ctx *pulumi.Context, args *LookupLbCertificateArgs, opts ...pulumi.InvokeOption) (*LookupLbCertificateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLbCertificateResult
	err := ctx.Invoke("scaleway:index/getLbCertificate:getLbCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbCertificate.
type LookupLbCertificateArgs struct {
	// The certificate id.
	// - Only one of `name` and `certificateId` should be specified.
	CertificateId *string `pulumi:"certificateId"`
	// The load-balancer ID this certificate is attached to.
	LbId *string `pulumi:"lbId"`
	// The name of the certificate backend.
	// - When using a certificate `name` you should specify the `lb-id`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getLbCertificate.
type LookupLbCertificateResult struct {
	CertificateId      *string                             `pulumi:"certificateId"`
	CommonName         string                              `pulumi:"commonName"`
	CustomCertificates []GetLbCertificateCustomCertificate `pulumi:"customCertificates"`
	Fingerprint        string                              `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                        `pulumi:"id"`
	LbId                    *string                       `pulumi:"lbId"`
	Letsencrypts            []GetLbCertificateLetsencrypt `pulumi:"letsencrypts"`
	Name                    *string                       `pulumi:"name"`
	NotValidAfter           string                        `pulumi:"notValidAfter"`
	NotValidBefore          string                        `pulumi:"notValidBefore"`
	Status                  string                        `pulumi:"status"`
	SubjectAlternativeNames []string                      `pulumi:"subjectAlternativeNames"`
}

func LookupLbCertificateOutput(ctx *pulumi.Context, args LookupLbCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupLbCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbCertificateResult, error) {
			args := v.(LookupLbCertificateArgs)
			r, err := LookupLbCertificate(ctx, &args, opts...)
			var s LookupLbCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbCertificateResultOutput)
}

// A collection of arguments for invoking getLbCertificate.
type LookupLbCertificateOutputArgs struct {
	// The certificate id.
	// - Only one of `name` and `certificateId` should be specified.
	CertificateId pulumi.StringPtrInput `pulumi:"certificateId"`
	// The load-balancer ID this certificate is attached to.
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The name of the certificate backend.
	// - When using a certificate `name` you should specify the `lb-id`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupLbCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getLbCertificate.
type LookupLbCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupLbCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbCertificateResult)(nil)).Elem()
}

func (o LookupLbCertificateResultOutput) ToLookupLbCertificateResultOutput() LookupLbCertificateResultOutput {
	return o
}

func (o LookupLbCertificateResultOutput) ToLookupLbCertificateResultOutputWithContext(ctx context.Context) LookupLbCertificateResultOutput {
	return o
}

func (o LookupLbCertificateResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

func (o LookupLbCertificateResultOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.CommonName }).(pulumi.StringOutput)
}

func (o LookupLbCertificateResultOutput) CustomCertificates() GetLbCertificateCustomCertificateArrayOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) []GetLbCertificateCustomCertificate { return v.CustomCertificates }).(GetLbCertificateCustomCertificateArrayOutput)
}

func (o LookupLbCertificateResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbCertificateResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLbCertificateResultOutput) Letsencrypts() GetLbCertificateLetsencryptArrayOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) []GetLbCertificateLetsencrypt { return v.Letsencrypts }).(GetLbCertificateLetsencryptArrayOutput)
}

func (o LookupLbCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLbCertificateResultOutput) NotValidAfter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.NotValidAfter }).(pulumi.StringOutput)
}

func (o LookupLbCertificateResultOutput) NotValidBefore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.NotValidBefore }).(pulumi.StringOutput)
}

func (o LookupLbCertificateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupLbCertificateResultOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLbCertificateResult) []string { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbCertificateResultOutput{})
}
