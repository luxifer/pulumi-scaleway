// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Load-Balancer Frontends. For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api).
//
// ## Examples Usage
//
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/luxifer/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewLbFrontend(ctx, "frontend01", &scaleway.LbFrontendArgs{
// 			LbId:        pulumi.Any(scaleway_lb.Lb01.Id),
// 			BackendId:   pulumi.Any(scaleway_lb_backend.Backend01.Id),
// 			InboundPort: pulumi.Int(80),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/lbFrontend:LbFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type LbFrontend struct {
	pulumi.CustomResourceState

	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LbFrontendAclArrayOutput `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	BackendId pulumi.StringOutput `pulumi:"backendId"`
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// List of Certificate IDs that should be used by the frontend.
	CertificateIds pulumi.StringArrayOutput `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrOutput `pulumi:"enableHttp3"`
	// TCP port to listen on the front side.
	InboundPort pulumi.IntOutput `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrOutput `pulumi:"timeoutClient"`
}

// NewLbFrontend registers a new resource with the given unique name, arguments, and options.
func NewLbFrontend(ctx *pulumi.Context,
	name string, args *LbFrontendArgs, opts ...pulumi.ResourceOption) (*LbFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendId == nil {
		return nil, errors.New("invalid value for required argument 'BackendId'")
	}
	if args.InboundPort == nil {
		return nil, errors.New("invalid value for required argument 'InboundPort'")
	}
	if args.LbId == nil {
		return nil, errors.New("invalid value for required argument 'LbId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LbFrontend
	err := ctx.RegisterResource("scaleway:index/lbFrontend:LbFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbFrontend gets an existing LbFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbFrontendState, opts ...pulumi.ResourceOption) (*LbFrontend, error) {
	var resource LbFrontend
	err := ctx.ReadResource("scaleway:index/lbFrontend:LbFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbFrontend resources.
type lbFrontendState struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls []LbFrontendAcl `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	BackendId *string `pulumi:"backendId"`
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId *string `pulumi:"certificateId"`
	// List of Certificate IDs that should be used by the frontend.
	CertificateIds []string `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 *bool `pulumi:"enableHttp3"`
	// TCP port to listen on the front side.
	InboundPort *int `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId *string `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient *string `pulumi:"timeoutClient"`
}

type LbFrontendState struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LbFrontendAclArrayInput
	// The load-balancer backend ID this frontend is attached to.
	BackendId pulumi.StringPtrInput
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId pulumi.StringPtrInput
	// List of Certificate IDs that should be used by the frontend.
	CertificateIds pulumi.StringArrayInput
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrInput
	// TCP port to listen on the front side.
	InboundPort pulumi.IntPtrInput
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrInput
}

func (LbFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbFrontendState)(nil)).Elem()
}

type lbFrontendArgs struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls []LbFrontendAcl `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	BackendId string `pulumi:"backendId"`
	// List of Certificate IDs that should be used by the frontend.
	CertificateIds []string `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 *bool `pulumi:"enableHttp3"`
	// TCP port to listen on the front side.
	InboundPort int `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId string `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient *string `pulumi:"timeoutClient"`
}

// The set of arguments for constructing a LbFrontend resource.
type LbFrontendArgs struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LbFrontendAclArrayInput
	// The load-balancer backend ID this frontend is attached to.
	BackendId pulumi.StringInput
	// List of Certificate IDs that should be used by the frontend.
	CertificateIds pulumi.StringArrayInput
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrInput
	// TCP port to listen on the front side.
	InboundPort pulumi.IntInput
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrInput
}

func (LbFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbFrontendArgs)(nil)).Elem()
}

type LbFrontendInput interface {
	pulumi.Input

	ToLbFrontendOutput() LbFrontendOutput
	ToLbFrontendOutputWithContext(ctx context.Context) LbFrontendOutput
}

func (*LbFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**LbFrontend)(nil)).Elem()
}

func (i *LbFrontend) ToLbFrontendOutput() LbFrontendOutput {
	return i.ToLbFrontendOutputWithContext(context.Background())
}

func (i *LbFrontend) ToLbFrontendOutputWithContext(ctx context.Context) LbFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbFrontendOutput)
}

// LbFrontendArrayInput is an input type that accepts LbFrontendArray and LbFrontendArrayOutput values.
// You can construct a concrete instance of `LbFrontendArrayInput` via:
//
//          LbFrontendArray{ LbFrontendArgs{...} }
type LbFrontendArrayInput interface {
	pulumi.Input

	ToLbFrontendArrayOutput() LbFrontendArrayOutput
	ToLbFrontendArrayOutputWithContext(context.Context) LbFrontendArrayOutput
}

type LbFrontendArray []LbFrontendInput

func (LbFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbFrontend)(nil)).Elem()
}

func (i LbFrontendArray) ToLbFrontendArrayOutput() LbFrontendArrayOutput {
	return i.ToLbFrontendArrayOutputWithContext(context.Background())
}

func (i LbFrontendArray) ToLbFrontendArrayOutputWithContext(ctx context.Context) LbFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbFrontendArrayOutput)
}

// LbFrontendMapInput is an input type that accepts LbFrontendMap and LbFrontendMapOutput values.
// You can construct a concrete instance of `LbFrontendMapInput` via:
//
//          LbFrontendMap{ "key": LbFrontendArgs{...} }
type LbFrontendMapInput interface {
	pulumi.Input

	ToLbFrontendMapOutput() LbFrontendMapOutput
	ToLbFrontendMapOutputWithContext(context.Context) LbFrontendMapOutput
}

type LbFrontendMap map[string]LbFrontendInput

func (LbFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbFrontend)(nil)).Elem()
}

func (i LbFrontendMap) ToLbFrontendMapOutput() LbFrontendMapOutput {
	return i.ToLbFrontendMapOutputWithContext(context.Background())
}

func (i LbFrontendMap) ToLbFrontendMapOutputWithContext(ctx context.Context) LbFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbFrontendMapOutput)
}

type LbFrontendOutput struct{ *pulumi.OutputState }

func (LbFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbFrontend)(nil)).Elem()
}

func (o LbFrontendOutput) ToLbFrontendOutput() LbFrontendOutput {
	return o
}

func (o LbFrontendOutput) ToLbFrontendOutputWithContext(ctx context.Context) LbFrontendOutput {
	return o
}

// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
func (o LbFrontendOutput) Acls() LbFrontendAclArrayOutput {
	return o.ApplyT(func(v *LbFrontend) LbFrontendAclArrayOutput { return v.Acls }).(LbFrontendAclArrayOutput)
}

// The load-balancer backend ID this frontend is attached to.
func (o LbFrontendOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringOutput { return v.BackendId }).(pulumi.StringOutput)
}

// (Deprecated) first certificate ID used by the frontend.
//
// Deprecated: Please use certificate_ids
func (o LbFrontendOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// List of Certificate IDs that should be used by the frontend.
func (o LbFrontendOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringArrayOutput { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

// Activates HTTP/3 protocol.
func (o LbFrontendOutput) EnableHttp3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.BoolPtrOutput { return v.EnableHttp3 }).(pulumi.BoolPtrOutput)
}

// TCP port to listen on the front side.
func (o LbFrontendOutput) InboundPort() pulumi.IntOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.IntOutput { return v.InboundPort }).(pulumi.IntOutput)
}

// The load-balancer ID this frontend is attached to.
func (o LbFrontendOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// The ACL name. If not provided it will be randomly generated.
func (o LbFrontendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Maximum inactivity time on the client side. (e.g.: `1s`)
func (o LbFrontendOutput) TimeoutClient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbFrontend) pulumi.StringPtrOutput { return v.TimeoutClient }).(pulumi.StringPtrOutput)
}

type LbFrontendArrayOutput struct{ *pulumi.OutputState }

func (LbFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbFrontend)(nil)).Elem()
}

func (o LbFrontendArrayOutput) ToLbFrontendArrayOutput() LbFrontendArrayOutput {
	return o
}

func (o LbFrontendArrayOutput) ToLbFrontendArrayOutputWithContext(ctx context.Context) LbFrontendArrayOutput {
	return o
}

func (o LbFrontendArrayOutput) Index(i pulumi.IntInput) LbFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbFrontend {
		return vs[0].([]*LbFrontend)[vs[1].(int)]
	}).(LbFrontendOutput)
}

type LbFrontendMapOutput struct{ *pulumi.OutputState }

func (LbFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbFrontend)(nil)).Elem()
}

func (o LbFrontendMapOutput) ToLbFrontendMapOutput() LbFrontendMapOutput {
	return o
}

func (o LbFrontendMapOutput) ToLbFrontendMapOutputWithContext(ctx context.Context) LbFrontendMapOutput {
	return o
}

func (o LbFrontendMapOutput) MapIndex(k pulumi.StringInput) LbFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbFrontend {
		return vs[0].(map[string]*LbFrontend)[vs[1].(string)]
	}).(LbFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbFrontendInput)(nil)).Elem(), &LbFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbFrontendArrayInput)(nil)).Elem(), LbFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbFrontendMapInput)(nil)).Elem(), LbFrontendMap{})
	pulumi.RegisterOutputType(LbFrontendOutput{})
	pulumi.RegisterOutputType(LbFrontendArrayOutput{})
	pulumi.RegisterOutputType(LbFrontendMapOutput{})
}