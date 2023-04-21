// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a dhcp entries. For further information please check the
// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)
func LookupVpcPublicGatewayDhcpReservation(ctx *pulumi.Context, args *LookupVpcPublicGatewayDhcpReservationArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayDhcpReservationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayDhcpReservationResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the Reservation to retrieve
	ReservationId *string `pulumi:"reservationId"`
	// Boolean to wait for macAddress to exist in dhcp
	WaitForDhcp *bool `pulumi:"waitForDhcp"`
	// `zone`) The zone in which
	// the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationResult struct {
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt string `pulumi:"createdAt"`
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The Hostname of the client machine.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IP address to give to the machine (IP address).
	IpAddress     string  `pulumi:"ipAddress"`
	MacAddress    *string `pulumi:"macAddress"`
	ReservationId *string `pulumi:"reservationId"`
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
	Type string `pulumi:"type"`
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt   string  `pulumi:"updatedAt"`
	WaitForDhcp *bool   `pulumi:"waitForDhcp"`
	Zone        *string `pulumi:"zone"`
}

func LookupVpcPublicGatewayDhcpReservationOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayDhcpReservationOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayDhcpReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayDhcpReservationResult, error) {
			args := v.(LookupVpcPublicGatewayDhcpReservationArgs)
			r, err := LookupVpcPublicGatewayDhcpReservation(ctx, &args, opts...)
			var s LookupVpcPublicGatewayDhcpReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcPublicGatewayDhcpReservationResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationOutputArgs struct {
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId pulumi.StringPtrInput `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the Reservation to retrieve
	ReservationId pulumi.StringPtrInput `pulumi:"reservationId"`
	// Boolean to wait for macAddress to exist in dhcp
	WaitForDhcp pulumi.BoolPtrInput `pulumi:"waitForDhcp"`
	// `zone`) The zone in which
	// the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupVpcPublicGatewayDhcpReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpReservationArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayDhcpReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpReservationResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ToLookupVpcPublicGatewayDhcpReservationResultOutput() LookupVpcPublicGatewayDhcpReservationResultOutput {
	return o
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ToLookupVpcPublicGatewayDhcpReservationResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayDhcpReservationResultOutput {
	return o
}

// The date and time of the creation of the public gateway DHCP config.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the owning GatewayNetwork.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) GatewayNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.GatewayNetworkId }).(pulumi.StringPtrOutput)
}

// The Hostname of the client machine.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IP address to give to the machine (IP address).
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ReservationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.ReservationId }).(pulumi.StringPtrOutput)
}

// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the public gateway DHCP config.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) WaitForDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *bool { return v.WaitForDhcp }).(pulumi.BoolPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayDhcpReservationResultOutput{})
}
