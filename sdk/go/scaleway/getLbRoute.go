// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Scaleway Load-Balancer Routes.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#route-ff94b7).
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
// 		ip01, err := scaleway.NewLbIp(ctx, "ip01", nil)
// 		if err != nil {
// 			return err
// 		}
// 		lb01, err := scaleway.NewLb(ctx, "lb01", &scaleway.LbArgs{
// 			IpId: ip01.ID(),
// 			Type: pulumi.String("lb-s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bkd01, err := scaleway.NewLbBackend(ctx, "bkd01", &scaleway.LbBackendArgs{
// 			LbId:            lb01.ID(),
// 			ForwardProtocol: pulumi.String("tcp"),
// 			ForwardPort:     pulumi.Int(80),
// 			ProxyProtocol:   pulumi.String("none"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		frt01, err := scaleway.NewLbFrontend(ctx, "frt01", &scaleway.LbFrontendArgs{
// 			LbId:        lb01.ID(),
// 			BackendId:   bkd01.ID(),
// 			InboundPort: pulumi.Int(80),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		rt01, err := scaleway.NewLbRoute(ctx, "rt01", &scaleway.LbRouteArgs{
// 			FrontendId: frt01.ID(),
// 			BackendId:  bkd01.ID(),
// 			MatchSni:   pulumi.String("sni.scaleway.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = scaleway.LookupLbRouteOutput(ctx, GetLbRouteOutputArgs{
// 			RouteId: rt01.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func LookupLbRoute(ctx *pulumi.Context, args *LookupLbRouteArgs, opts ...pulumi.InvokeOption) (*LookupLbRouteResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLbRouteResult
	err := ctx.Invoke("scaleway:index/getLbRoute:getLbRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbRoute.
type LookupLbRouteArgs struct {
	// The route id.
	RouteId string `pulumi:"routeId"`
}

// A collection of values returned by getLbRoute.
type LookupLbRouteResult struct {
	BackendId  string `pulumi:"backendId"`
	CreatedAt  string `pulumi:"createdAt"`
	FrontendId string `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	MatchHostHeader string `pulumi:"matchHostHeader"`
	MatchSni        string `pulumi:"matchSni"`
	RouteId         string `pulumi:"routeId"`
	UpdatedAt       string `pulumi:"updatedAt"`
}

func LookupLbRouteOutput(ctx *pulumi.Context, args LookupLbRouteOutputArgs, opts ...pulumi.InvokeOption) LookupLbRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbRouteResult, error) {
			args := v.(LookupLbRouteArgs)
			r, err := LookupLbRoute(ctx, &args, opts...)
			var s LookupLbRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbRouteResultOutput)
}

// A collection of arguments for invoking getLbRoute.
type LookupLbRouteOutputArgs struct {
	// The route id.
	RouteId pulumi.StringInput `pulumi:"routeId"`
}

func (LookupLbRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbRouteArgs)(nil)).Elem()
}

// A collection of values returned by getLbRoute.
type LookupLbRouteResultOutput struct{ *pulumi.OutputState }

func (LookupLbRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbRouteResult)(nil)).Elem()
}

func (o LookupLbRouteResultOutput) ToLookupLbRouteResultOutput() LookupLbRouteResultOutput {
	return o
}

func (o LookupLbRouteResultOutput) ToLookupLbRouteResultOutputWithContext(ctx context.Context) LookupLbRouteResultOutput {
	return o
}

func (o LookupLbRouteResultOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.BackendId }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) FrontendId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.FrontendId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbRouteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) MatchHostHeader() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.MatchHostHeader }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) MatchSni() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.MatchSni }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.RouteId }).(pulumi.StringOutput)
}

func (o LookupLbRouteResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRouteResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbRouteResultOutput{})
}
