// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an IOT Device.
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
// 		_, err := scaleway.LookupIotDevice(ctx, &GetIotDeviceArgs{
// 			DeviceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupIotDevice(ctx *pulumi.Context, args *LookupIotDeviceArgs, opts ...pulumi.InvokeOption) (*LookupIotDeviceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIotDeviceResult
	err := ctx.Invoke("scaleway:index/getIotDevice:getIotDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIotDevice.
type LookupIotDeviceArgs struct {
	// The device ID.
	// Only one of the `name` and `deviceId` should be specified.
	DeviceId *string `pulumi:"deviceId"`
	// The hub ID.
	HubId *string `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `deviceId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The region in which the hub exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getIotDevice.
type LookupIotDeviceResult struct {
	AllowInsecure            bool                      `pulumi:"allowInsecure"`
	AllowMultipleConnections bool                      `pulumi:"allowMultipleConnections"`
	Certificates             []GetIotDeviceCertificate `pulumi:"certificates"`
	CreatedAt                string                    `pulumi:"createdAt"`
	Description              string                    `pulumi:"description"`
	DeviceId                 *string                   `pulumi:"deviceId"`
	HubId                    string                    `pulumi:"hubId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                      `pulumi:"id"`
	IsConnected    bool                        `pulumi:"isConnected"`
	LastActivityAt string                      `pulumi:"lastActivityAt"`
	MessageFilters []GetIotDeviceMessageFilter `pulumi:"messageFilters"`
	Name           *string                     `pulumi:"name"`
	Region         *string                     `pulumi:"region"`
	Status         string                      `pulumi:"status"`
	UpdatedAt      string                      `pulumi:"updatedAt"`
}

func LookupIotDeviceOutput(ctx *pulumi.Context, args LookupIotDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupIotDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotDeviceResult, error) {
			args := v.(LookupIotDeviceArgs)
			r, err := LookupIotDevice(ctx, &args, opts...)
			var s LookupIotDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotDeviceResultOutput)
}

// A collection of arguments for invoking getIotDevice.
type LookupIotDeviceOutputArgs struct {
	// The device ID.
	// Only one of the `name` and `deviceId` should be specified.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// The hub ID.
	HubId pulumi.StringPtrInput `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `deviceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the hub exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupIotDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDeviceArgs)(nil)).Elem()
}

// A collection of values returned by getIotDevice.
type LookupIotDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupIotDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDeviceResult)(nil)).Elem()
}

func (o LookupIotDeviceResultOutput) ToLookupIotDeviceResultOutput() LookupIotDeviceResultOutput {
	return o
}

func (o LookupIotDeviceResultOutput) ToLookupIotDeviceResultOutputWithContext(ctx context.Context) LookupIotDeviceResultOutput {
	return o
}

func (o LookupIotDeviceResultOutput) AllowInsecure() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) bool { return v.AllowInsecure }).(pulumi.BoolOutput)
}

func (o LookupIotDeviceResultOutput) AllowMultipleConnections() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) bool { return v.AllowMultipleConnections }).(pulumi.BoolOutput)
}

func (o LookupIotDeviceResultOutput) Certificates() GetIotDeviceCertificateArrayOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) []GetIotDeviceCertificate { return v.Certificates }).(GetIotDeviceCertificateArrayOutput)
}

func (o LookupIotDeviceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupIotDeviceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupIotDeviceResultOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

func (o LookupIotDeviceResultOutput) HubId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.HubId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIotDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotDeviceResultOutput) IsConnected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) bool { return v.IsConnected }).(pulumi.BoolOutput)
}

func (o LookupIotDeviceResultOutput) LastActivityAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.LastActivityAt }).(pulumi.StringOutput)
}

func (o LookupIotDeviceResultOutput) MessageFilters() GetIotDeviceMessageFilterArrayOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) []GetIotDeviceMessageFilter { return v.MessageFilters }).(GetIotDeviceMessageFilterArrayOutput)
}

func (o LookupIotDeviceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIotDeviceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupIotDeviceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIotDeviceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDeviceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotDeviceResultOutput{})
}