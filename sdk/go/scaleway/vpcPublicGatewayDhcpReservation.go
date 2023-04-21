// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages the [Scaleway DHCP Reservations](https://www.scaleway.com/en/docs/network/vpc/concepts/#dhcp).
//
// The static associations are used to assign IP addresses based on the MAC addresses of the Instance.
//
// Statically assigned IP addresses should fall within the configured subnet, but be outside of the dynamic range.
//
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544) and [configuration guide](https://www.scaleway.com/en/docs/network/vpc/how-to/configure-a-public-gateway/#how-to-review-and-configure-dhcp).
//
// [DHCP reservations](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6) hold both dynamic DHCP leases (IP addresses dynamically assigned by the gateway to instances) and static user-created DHCP reservations.
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
// 		mainVpcPrivateNetwork, err := scaleway.NewVpcPrivateNetwork(ctx, "mainVpcPrivateNetwork", nil)
// 		if err != nil {
// 			return err
// 		}
// 		mainInstanceServer, err := scaleway.NewInstanceServer(ctx, "mainInstanceServer", &scaleway.InstanceServerArgs{
// 			Image: pulumi.String("ubuntu_jammy"),
// 			Type:  pulumi.String("DEV1-S"),
// 			Zone:  pulumi.String("fr-par-1"),
// 			PrivateNetworks: InstanceServerPrivateNetworkArray{
// 				&InstanceServerPrivateNetworkArgs{
// 					PnId: mainVpcPrivateNetwork.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainVpcPublicGatewayIp, err := scaleway.NewVpcPublicGatewayIp(ctx, "mainVpcPublicGatewayIp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		mainVpcPublicGatewayDhcp, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "mainVpcPublicGatewayDhcp", &scaleway.VpcPublicGatewayDhcpArgs{
// 			Subnet: pulumi.String("192.168.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainVpcPublicGateway, err := scaleway.NewVpcPublicGateway(ctx, "mainVpcPublicGateway", &scaleway.VpcPublicGatewayArgs{
// 			Type: pulumi.String("VPC-GW-S"),
// 			IpId: mainVpcPublicGatewayIp.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		mainVpcGatewayNetwork, err := scaleway.NewVpcGatewayNetwork(ctx, "mainVpcGatewayNetwork", &scaleway.VpcGatewayNetworkArgs{
// 			GatewayId:        mainVpcPublicGateway.ID(),
// 			PrivateNetworkId: mainVpcPrivateNetwork.ID(),
// 			DhcpId:           mainVpcPublicGatewayDhcp.ID(),
// 			CleanupDhcp:      pulumi.Bool(true),
// 			EnableMasquerade: pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			mainVpcPublicGatewayIp,
// 			mainVpcPrivateNetwork,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewVpcPublicGatewayDhcpReservation(ctx, "mainVpcPublicGatewayDhcpReservation", &scaleway.VpcPublicGatewayDhcpReservationArgs{
// 			GatewayNetworkId: mainVpcGatewayNetwork.ID(),
// 			MacAddress: mainInstanceServer.PrivateNetworks.ApplyT(func(privateNetworks []InstanceServerPrivateNetwork) (string, error) {
// 				return privateNetworks[0].MacAddress, nil
// 			}).(pulumi.StringOutput),
// 			IpAddress: pulumi.String("192.168.1.1"),
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
// Public gateway DHCP Reservation config can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGatewayDhcpReservation struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId pulumi.StringOutput `pulumi:"gatewayNetworkId"`
	// The Hostname of the client machine.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The IP address to give to the machine (IP address).
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayDhcpReservation registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayDhcpReservation(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayDhcpReservationArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcpReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayNetworkId'")
	}
	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.MacAddress == nil {
		return nil, errors.New("invalid value for required argument 'MacAddress'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayDhcpReservation
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayDhcpReservation gets an existing VpcPublicGatewayDhcpReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayDhcpReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayDhcpReservationState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcpReservation, error) {
	var resource VpcPublicGatewayDhcpReservation
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayDhcpReservation resources.
type vpcPublicGatewayDhcpReservationState struct {
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The Hostname of the client machine.
	Hostname *string `pulumi:"hostname"`
	// The IP address to give to the machine (IP address).
	IpAddress *string `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress *string `pulumi:"macAddress"`
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
	Type *string `pulumi:"type"`
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayDhcpReservationState struct {
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt pulumi.StringPtrInput
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId pulumi.StringPtrInput
	// The Hostname of the client machine.
	Hostname pulumi.StringPtrInput
	// The IP address to give to the machine (IP address).
	IpAddress pulumi.StringPtrInput
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringPtrInput
	// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
	Type pulumi.StringPtrInput
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpReservationState)(nil)).Elem()
}

type vpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId string `pulumi:"gatewayNetworkId"`
	// The IP address to give to the machine (IP address).
	IpAddress string `pulumi:"ipAddress"`
	// The MAC address to give a static entry to.
	MacAddress string `pulumi:"macAddress"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayDhcpReservation resource.
type VpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork.
	GatewayNetworkId pulumi.StringInput
	// The IP address to give to the machine (IP address).
	IpAddress pulumi.StringInput
	// The MAC address to give a static entry to.
	MacAddress pulumi.StringInput
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpReservationArgs)(nil)).Elem()
}

type VpcPublicGatewayDhcpReservationInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput
	ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput
}

func (*VpcPublicGatewayDhcpReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (i *VpcPublicGatewayDhcpReservation) ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput {
	return i.ToVpcPublicGatewayDhcpReservationOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayDhcpReservation) ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpReservationOutput)
}

// VpcPublicGatewayDhcpReservationArrayInput is an input type that accepts VpcPublicGatewayDhcpReservationArray and VpcPublicGatewayDhcpReservationArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayDhcpReservationArrayInput` via:
//
//          VpcPublicGatewayDhcpReservationArray{ VpcPublicGatewayDhcpReservationArgs{...} }
type VpcPublicGatewayDhcpReservationArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpReservationArrayOutput() VpcPublicGatewayDhcpReservationArrayOutput
	ToVpcPublicGatewayDhcpReservationArrayOutputWithContext(context.Context) VpcPublicGatewayDhcpReservationArrayOutput
}

type VpcPublicGatewayDhcpReservationArray []VpcPublicGatewayDhcpReservationInput

func (VpcPublicGatewayDhcpReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (i VpcPublicGatewayDhcpReservationArray) ToVpcPublicGatewayDhcpReservationArrayOutput() VpcPublicGatewayDhcpReservationArrayOutput {
	return i.ToVpcPublicGatewayDhcpReservationArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayDhcpReservationArray) ToVpcPublicGatewayDhcpReservationArrayOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpReservationArrayOutput)
}

// VpcPublicGatewayDhcpReservationMapInput is an input type that accepts VpcPublicGatewayDhcpReservationMap and VpcPublicGatewayDhcpReservationMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayDhcpReservationMapInput` via:
//
//          VpcPublicGatewayDhcpReservationMap{ "key": VpcPublicGatewayDhcpReservationArgs{...} }
type VpcPublicGatewayDhcpReservationMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpReservationMapOutput() VpcPublicGatewayDhcpReservationMapOutput
	ToVpcPublicGatewayDhcpReservationMapOutputWithContext(context.Context) VpcPublicGatewayDhcpReservationMapOutput
}

type VpcPublicGatewayDhcpReservationMap map[string]VpcPublicGatewayDhcpReservationInput

func (VpcPublicGatewayDhcpReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (i VpcPublicGatewayDhcpReservationMap) ToVpcPublicGatewayDhcpReservationMapOutput() VpcPublicGatewayDhcpReservationMapOutput {
	return i.ToVpcPublicGatewayDhcpReservationMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayDhcpReservationMap) ToVpcPublicGatewayDhcpReservationMapOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpReservationMapOutput)
}

type VpcPublicGatewayDhcpReservationOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpReservationOutput) ToVpcPublicGatewayDhcpReservationOutput() VpcPublicGatewayDhcpReservationOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationOutput) ToVpcPublicGatewayDhcpReservationOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationOutput {
	return o
}

// The date and time of the creation of the public gateway DHCP config.
func (o VpcPublicGatewayDhcpReservationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the owning GatewayNetwork.
func (o VpcPublicGatewayDhcpReservationOutput) GatewayNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.GatewayNetworkId }).(pulumi.StringOutput)
}

// The Hostname of the client machine.
func (o VpcPublicGatewayDhcpReservationOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The IP address to give to the machine (IP address).
func (o VpcPublicGatewayDhcpReservationOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The MAC address to give a static entry to.
func (o VpcPublicGatewayDhcpReservationOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
func (o VpcPublicGatewayDhcpReservationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the public gateway DHCP config.
func (o VpcPublicGatewayDhcpReservationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the public gateway DHCP config should be created.
func (o VpcPublicGatewayDhcpReservationOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcpReservation) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayDhcpReservationArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpReservationArrayOutput) ToVpcPublicGatewayDhcpReservationArrayOutput() VpcPublicGatewayDhcpReservationArrayOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationArrayOutput) ToVpcPublicGatewayDhcpReservationArrayOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationArrayOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayDhcpReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGatewayDhcpReservation {
		return vs[0].([]*VpcPublicGatewayDhcpReservation)[vs[1].(int)]
	}).(VpcPublicGatewayDhcpReservationOutput)
}

type VpcPublicGatewayDhcpReservationMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayDhcpReservation)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpReservationMapOutput) ToVpcPublicGatewayDhcpReservationMapOutput() VpcPublicGatewayDhcpReservationMapOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationMapOutput) ToVpcPublicGatewayDhcpReservationMapOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpReservationMapOutput {
	return o
}

func (o VpcPublicGatewayDhcpReservationMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayDhcpReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGatewayDhcpReservation {
		return vs[0].(map[string]*VpcPublicGatewayDhcpReservation)[vs[1].(string)]
	}).(VpcPublicGatewayDhcpReservationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpReservationInput)(nil)).Elem(), &VpcPublicGatewayDhcpReservation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpReservationArrayInput)(nil)).Elem(), VpcPublicGatewayDhcpReservationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpReservationMapInput)(nil)).Elem(), VpcPublicGatewayDhcpReservationMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpReservationOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpReservationArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpReservationMapOutput{})
}
