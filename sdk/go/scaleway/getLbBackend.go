// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Scaleway Load-Balancer Backends.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#backends-cbf4eb).
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
// 		mainLbIp, err := scaleway.NewLbIp(ctx, "mainLbIp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		mainLb, err := scaleway.NewLb(ctx, "mainLb", &scaleway.LbArgs{
// 			IpId: mainLbIp.ID(),
// 			Type: pulumi.String("LB-S"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainLbBackend, err := scaleway.NewLbBackend(ctx, "mainLbBackend", &scaleway.LbBackendArgs{
// 			LbId:            mainLb.ID(),
// 			ForwardProtocol: pulumi.String("http"),
// 			ForwardPort:     pulumi.Int(80),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = scaleway.LookupLbBackendOutput(ctx, GetLbBackendOutputArgs{
// 			BackendId: mainLbBackend.ID(),
// 		}, nil)
// 		_ = scaleway.LookupLbBackendOutput(ctx, GetLbBackendOutputArgs{
// 			Name: mainLbBackend.Name,
// 			LbId: mainLb.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func LookupLbBackend(ctx *pulumi.Context, args *LookupLbBackendArgs, opts ...pulumi.InvokeOption) (*LookupLbBackendResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLbBackendResult
	err := ctx.Invoke("scaleway:index/getLbBackend:getLbBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbBackend.
type LookupLbBackendArgs struct {
	// The backend id.
	// - Only one of `name` and `backendId` should be specified.
	BackendId *string `pulumi:"backendId"`
	// The load-balancer ID this backend is attached to.
	LbId *string `pulumi:"lbId"`
	// The name of the backend.
	// - When using the `name` you should specify the `lb-id`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getLbBackend.
type LookupLbBackendResult struct {
	BackendId             *string                       `pulumi:"backendId"`
	FailoverHost          string                        `pulumi:"failoverHost"`
	ForwardPort           int                           `pulumi:"forwardPort"`
	ForwardPortAlgorithm  string                        `pulumi:"forwardPortAlgorithm"`
	ForwardProtocol       string                        `pulumi:"forwardProtocol"`
	HealthCheckDelay      string                        `pulumi:"healthCheckDelay"`
	HealthCheckHttps      []GetLbBackendHealthCheckHttp `pulumi:"healthCheckHttps"`
	HealthCheckMaxRetries int                           `pulumi:"healthCheckMaxRetries"`
	HealthCheckPort       int                           `pulumi:"healthCheckPort"`
	HealthCheckTcps       []GetLbBackendHealthCheckTcp  `pulumi:"healthCheckTcps"`
	HealthCheckTimeout    string                        `pulumi:"healthCheckTimeout"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string   `pulumi:"id"`
	IgnoreSslServerVerify    bool     `pulumi:"ignoreSslServerVerify"`
	LbId                     *string  `pulumi:"lbId"`
	Name                     *string  `pulumi:"name"`
	OnMarkedDownAction       string   `pulumi:"onMarkedDownAction"`
	ProxyProtocol            string   `pulumi:"proxyProtocol"`
	SendProxyV2              bool     `pulumi:"sendProxyV2"`
	ServerIps                []string `pulumi:"serverIps"`
	SslBridging              bool     `pulumi:"sslBridging"`
	StickySessions           string   `pulumi:"stickySessions"`
	StickySessionsCookieName string   `pulumi:"stickySessionsCookieName"`
	TimeoutConnect           string   `pulumi:"timeoutConnect"`
	TimeoutServer            string   `pulumi:"timeoutServer"`
	TimeoutTunnel            string   `pulumi:"timeoutTunnel"`
}

func LookupLbBackendOutput(ctx *pulumi.Context, args LookupLbBackendOutputArgs, opts ...pulumi.InvokeOption) LookupLbBackendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbBackendResult, error) {
			args := v.(LookupLbBackendArgs)
			r, err := LookupLbBackend(ctx, &args, opts...)
			var s LookupLbBackendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbBackendResultOutput)
}

// A collection of arguments for invoking getLbBackend.
type LookupLbBackendOutputArgs struct {
	// The backend id.
	// - Only one of `name` and `backendId` should be specified.
	BackendId pulumi.StringPtrInput `pulumi:"backendId"`
	// The load-balancer ID this backend is attached to.
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The name of the backend.
	// - When using the `name` you should specify the `lb-id`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupLbBackendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbBackendArgs)(nil)).Elem()
}

// A collection of values returned by getLbBackend.
type LookupLbBackendResultOutput struct{ *pulumi.OutputState }

func (LookupLbBackendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbBackendResult)(nil)).Elem()
}

func (o LookupLbBackendResultOutput) ToLookupLbBackendResultOutput() LookupLbBackendResultOutput {
	return o
}

func (o LookupLbBackendResultOutput) ToLookupLbBackendResultOutputWithContext(ctx context.Context) LookupLbBackendResultOutput {
	return o
}

func (o LookupLbBackendResultOutput) BackendId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbBackendResult) *string { return v.BackendId }).(pulumi.StringPtrOutput)
}

func (o LookupLbBackendResultOutput) FailoverHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.FailoverHost }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) ForwardPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbBackendResult) int { return v.ForwardPort }).(pulumi.IntOutput)
}

func (o LookupLbBackendResultOutput) ForwardPortAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.ForwardPortAlgorithm }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) ForwardProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.ForwardProtocol }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckDelay() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.HealthCheckDelay }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckHttps() GetLbBackendHealthCheckHttpArrayOutput {
	return o.ApplyT(func(v LookupLbBackendResult) []GetLbBackendHealthCheckHttp { return v.HealthCheckHttps }).(GetLbBackendHealthCheckHttpArrayOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckMaxRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbBackendResult) int { return v.HealthCheckMaxRetries }).(pulumi.IntOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbBackendResult) int { return v.HealthCheckPort }).(pulumi.IntOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckTcps() GetLbBackendHealthCheckTcpArrayOutput {
	return o.ApplyT(func(v LookupLbBackendResult) []GetLbBackendHealthCheckTcp { return v.HealthCheckTcps }).(GetLbBackendHealthCheckTcpArrayOutput)
}

func (o LookupLbBackendResultOutput) HealthCheckTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.HealthCheckTimeout }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbBackendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) IgnoreSslServerVerify() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbBackendResult) bool { return v.IgnoreSslServerVerify }).(pulumi.BoolOutput)
}

func (o LookupLbBackendResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbBackendResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLbBackendResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbBackendResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLbBackendResultOutput) OnMarkedDownAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.OnMarkedDownAction }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) ProxyProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.ProxyProtocol }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) SendProxyV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbBackendResult) bool { return v.SendProxyV2 }).(pulumi.BoolOutput)
}

func (o LookupLbBackendResultOutput) ServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLbBackendResult) []string { return v.ServerIps }).(pulumi.StringArrayOutput)
}

func (o LookupLbBackendResultOutput) SslBridging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbBackendResult) bool { return v.SslBridging }).(pulumi.BoolOutput)
}

func (o LookupLbBackendResultOutput) StickySessions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.StickySessions }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) StickySessionsCookieName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.StickySessionsCookieName }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) TimeoutConnect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.TimeoutConnect }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) TimeoutServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.TimeoutServer }).(pulumi.StringOutput)
}

func (o LookupLbBackendResultOutput) TimeoutTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbBackendResult) string { return v.TimeoutTunnel }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbBackendResultOutput{})
}