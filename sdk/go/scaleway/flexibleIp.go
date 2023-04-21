// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway flexible IPs.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).
//
// ## Examples
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
// 		_, err := scaleway.NewFlexibleIp(ctx, "main", &scaleway.FlexibleIpArgs{
// 			Reverse: pulumi.String("my-reverse.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### With zone
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
// 		_, err := scaleway.NewFlexibleIp(ctx, "main", &scaleway.FlexibleIpArgs{
// 			Zone: pulumi.String("fr-par-2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### With baremetal server
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
// 		mainAccountSshKey, err := scaleway.NewAccountSshKey(ctx, "mainAccountSshKey", &scaleway.AccountSshKeyArgs{
// 			PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILHy/M5FVm5ydLGcal3e5LNcfTalbeN7QL/ZGCvDEdqJ foobar@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		byId, err := scaleway.GetBaremetalOs(ctx, &GetBaremetalOsArgs{
// 			Zone:    pulumi.StringRef("fr-par-2"),
// 			Name:    pulumi.StringRef("Ubuntu"),
// 			Version: pulumi.StringRef("20.04 LTS (Focal Fossa)"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		myOffer, err := scaleway.GetBaremetalOffer(ctx, &GetBaremetalOfferArgs{
// 			Zone: pulumi.StringRef("fr-par-2"),
// 			Name: pulumi.StringRef("EM-A210R-HDD"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		base, err := scaleway.NewBaremetalServer(ctx, "base", &scaleway.BaremetalServerArgs{
// 			Zone:      pulumi.String("fr-par-2"),
// 			Offer:     pulumi.String(myOffer.OfferId),
// 			Os:        pulumi.String(byId.OsId),
// 			SshKeyIds: mainAccountSshKey.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewFlexibleIp(ctx, "mainFlexibleIp", &scaleway.FlexibleIpArgs{
// 			ServerId: base.ID(),
// 			Zone:     pulumi.String("fr-par-2"),
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
// Flexible IPs can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/flexibleIp:FlexibleIp main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type FlexibleIp struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the Flexible IP (Format ISO 8601)
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the flexible IP
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IPv4 address of the Flexible IP
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The MAC address of the server associated with this flexible IP
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The organization of the Flexible IP
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The project of the Flexible IP
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The ID of the associated server
	ServerId pulumi.StringPtrOutput `pulumi:"serverId"`
	// The tags associated with the flexible IP
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the Flexible IP (Format ISO 8601)
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone of the Flexible IP
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFlexibleIp registers a new resource with the given unique name, arguments, and options.
func NewFlexibleIp(ctx *pulumi.Context,
	name string, args *FlexibleIpArgs, opts ...pulumi.ResourceOption) (*FlexibleIp, error) {
	if args == nil {
		args = &FlexibleIpArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FlexibleIp
	err := ctx.RegisterResource("scaleway:index/flexibleIp:FlexibleIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleIp gets an existing FlexibleIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleIpState, opts ...pulumi.ResourceOption) (*FlexibleIp, error) {
	var resource FlexibleIp
	err := ctx.ReadResource("scaleway:index/flexibleIp:FlexibleIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleIp resources.
type flexibleIpState struct {
	// The date and time of the creation of the Flexible IP (Format ISO 8601)
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the flexible IP
	Description *string `pulumi:"description"`
	// The IPv4 address of the Flexible IP
	IpAddress *string `pulumi:"ipAddress"`
	// The MAC address of the server associated with this flexible IP
	MacAddress *string `pulumi:"macAddress"`
	// The organization of the Flexible IP
	OrganizationId *string `pulumi:"organizationId"`
	// The project of the Flexible IP
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse *string `pulumi:"reverse"`
	// The ID of the associated server
	ServerId *string `pulumi:"serverId"`
	// The tags associated with the flexible IP
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the Flexible IP (Format ISO 8601)
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone of the Flexible IP
	Zone *string `pulumi:"zone"`
}

type FlexibleIpState struct {
	// The date and time of the creation of the Flexible IP (Format ISO 8601)
	CreatedAt pulumi.StringPtrInput
	// Description of the flexible IP
	Description pulumi.StringPtrInput
	// The IPv4 address of the Flexible IP
	IpAddress pulumi.StringPtrInput
	// The MAC address of the server associated with this flexible IP
	MacAddress pulumi.StringPtrInput
	// The organization of the Flexible IP
	OrganizationId pulumi.StringPtrInput
	// The project of the Flexible IP
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringPtrInput
	// The ID of the associated server
	ServerId pulumi.StringPtrInput
	// The tags associated with the flexible IP
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the Flexible IP (Format ISO 8601)
	UpdatedAt pulumi.StringPtrInput
	// The zone of the Flexible IP
	Zone pulumi.StringPtrInput
}

func (FlexibleIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIpState)(nil)).Elem()
}

type flexibleIpArgs struct {
	// Description of the flexible IP
	Description *string `pulumi:"description"`
	// The project of the Flexible IP
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse *string `pulumi:"reverse"`
	// The ID of the associated server
	ServerId *string `pulumi:"serverId"`
	// The tags associated with the flexible IP
	Tags []string `pulumi:"tags"`
	// The zone of the Flexible IP
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a FlexibleIp resource.
type FlexibleIpArgs struct {
	// Description of the flexible IP
	Description pulumi.StringPtrInput
	// The project of the Flexible IP
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringPtrInput
	// The ID of the associated server
	ServerId pulumi.StringPtrInput
	// The tags associated with the flexible IP
	Tags pulumi.StringArrayInput
	// The zone of the Flexible IP
	Zone pulumi.StringPtrInput
}

func (FlexibleIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIpArgs)(nil)).Elem()
}

type FlexibleIpInput interface {
	pulumi.Input

	ToFlexibleIpOutput() FlexibleIpOutput
	ToFlexibleIpOutputWithContext(ctx context.Context) FlexibleIpOutput
}

func (*FlexibleIp) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIp)(nil)).Elem()
}

func (i *FlexibleIp) ToFlexibleIpOutput() FlexibleIpOutput {
	return i.ToFlexibleIpOutputWithContext(context.Background())
}

func (i *FlexibleIp) ToFlexibleIpOutputWithContext(ctx context.Context) FlexibleIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpOutput)
}

// FlexibleIpArrayInput is an input type that accepts FlexibleIpArray and FlexibleIpArrayOutput values.
// You can construct a concrete instance of `FlexibleIpArrayInput` via:
//
//          FlexibleIpArray{ FlexibleIpArgs{...} }
type FlexibleIpArrayInput interface {
	pulumi.Input

	ToFlexibleIpArrayOutput() FlexibleIpArrayOutput
	ToFlexibleIpArrayOutputWithContext(context.Context) FlexibleIpArrayOutput
}

type FlexibleIpArray []FlexibleIpInput

func (FlexibleIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIp)(nil)).Elem()
}

func (i FlexibleIpArray) ToFlexibleIpArrayOutput() FlexibleIpArrayOutput {
	return i.ToFlexibleIpArrayOutputWithContext(context.Background())
}

func (i FlexibleIpArray) ToFlexibleIpArrayOutputWithContext(ctx context.Context) FlexibleIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpArrayOutput)
}

// FlexibleIpMapInput is an input type that accepts FlexibleIpMap and FlexibleIpMapOutput values.
// You can construct a concrete instance of `FlexibleIpMapInput` via:
//
//          FlexibleIpMap{ "key": FlexibleIpArgs{...} }
type FlexibleIpMapInput interface {
	pulumi.Input

	ToFlexibleIpMapOutput() FlexibleIpMapOutput
	ToFlexibleIpMapOutputWithContext(context.Context) FlexibleIpMapOutput
}

type FlexibleIpMap map[string]FlexibleIpInput

func (FlexibleIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIp)(nil)).Elem()
}

func (i FlexibleIpMap) ToFlexibleIpMapOutput() FlexibleIpMapOutput {
	return i.ToFlexibleIpMapOutputWithContext(context.Background())
}

func (i FlexibleIpMap) ToFlexibleIpMapOutputWithContext(ctx context.Context) FlexibleIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIpMapOutput)
}

type FlexibleIpOutput struct{ *pulumi.OutputState }

func (FlexibleIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIp)(nil)).Elem()
}

func (o FlexibleIpOutput) ToFlexibleIpOutput() FlexibleIpOutput {
	return o
}

func (o FlexibleIpOutput) ToFlexibleIpOutputWithContext(ctx context.Context) FlexibleIpOutput {
	return o
}

// The date and time of the creation of the Flexible IP (Format ISO 8601)
func (o FlexibleIpOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the flexible IP
func (o FlexibleIpOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IPv4 address of the Flexible IP
func (o FlexibleIpOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The MAC address of the server associated with this flexible IP
func (o FlexibleIpOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The organization of the Flexible IP
func (o FlexibleIpOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The project of the Flexible IP
func (o FlexibleIpOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain associated with this flexible IP.
func (o FlexibleIpOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The ID of the associated server
func (o FlexibleIpOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringPtrOutput { return v.ServerId }).(pulumi.StringPtrOutput)
}

// The tags associated with the flexible IP
func (o FlexibleIpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the Flexible IP (Format ISO 8601)
func (o FlexibleIpOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone of the Flexible IP
func (o FlexibleIpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type FlexibleIpArrayOutput struct{ *pulumi.OutputState }

func (FlexibleIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIp)(nil)).Elem()
}

func (o FlexibleIpArrayOutput) ToFlexibleIpArrayOutput() FlexibleIpArrayOutput {
	return o
}

func (o FlexibleIpArrayOutput) ToFlexibleIpArrayOutputWithContext(ctx context.Context) FlexibleIpArrayOutput {
	return o
}

func (o FlexibleIpArrayOutput) Index(i pulumi.IntInput) FlexibleIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlexibleIp {
		return vs[0].([]*FlexibleIp)[vs[1].(int)]
	}).(FlexibleIpOutput)
}

type FlexibleIpMapOutput struct{ *pulumi.OutputState }

func (FlexibleIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIp)(nil)).Elem()
}

func (o FlexibleIpMapOutput) ToFlexibleIpMapOutput() FlexibleIpMapOutput {
	return o
}

func (o FlexibleIpMapOutput) ToFlexibleIpMapOutputWithContext(ctx context.Context) FlexibleIpMapOutput {
	return o
}

func (o FlexibleIpMapOutput) MapIndex(k pulumi.StringInput) FlexibleIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlexibleIp {
		return vs[0].(map[string]*FlexibleIp)[vs[1].(string)]
	}).(FlexibleIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpInput)(nil)).Elem(), &FlexibleIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpArrayInput)(nil)).Elem(), FlexibleIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIpMapInput)(nil)).Elem(), FlexibleIpMap{})
	pulumi.RegisterOutputType(FlexibleIpOutput{})
	pulumi.RegisterOutputType(FlexibleIpArrayOutput{})
	pulumi.RegisterOutputType(FlexibleIpMapOutput{})
}
