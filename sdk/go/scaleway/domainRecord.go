// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Domain record.\
// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
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
// 		_, err := scaleway.NewDomainRecord(ctx, "www", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.4"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			Ttl:     pulumi.Int(3600),
// 			Type:    pulumi.String("A"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "www2", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.5"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			Ttl:     pulumi.Int(3600),
// 			Type:    pulumi.String("A"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "mx", &scaleway.DomainRecordArgs{
// 			Data:     pulumi.String("mx.online.net."),
// 			DnsZone:  pulumi.String("domain.tld"),
// 			Priority: pulumi.Int(10),
// 			Ttl:      pulumi.Int(3600),
// 			Type:     pulumi.String("MX"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "mx2", &scaleway.DomainRecordArgs{
// 			Data:     pulumi.String("mx-cache.online.net."),
// 			DnsZone:  pulumi.String("domain.tld"),
// 			Priority: pulumi.Int(20),
// 			Ttl:      pulumi.Int(3600),
// 			Type:     pulumi.String("MX"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### With dynamic records
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
// 		_, err := scaleway.NewDomainRecord(ctx, "geoIp", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.4"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			GeoIp: &DomainRecordGeoIpArgs{
// 				Matches: DomainRecordGeoIpMatchArray{
// 					&DomainRecordGeoIpMatchArgs{
// 						Continents: pulumi.StringArray{
// 							pulumi.String("EU"),
// 						},
// 						Countries: pulumi.StringArray{
// 							pulumi.String("FR"),
// 						},
// 						Data: pulumi.String("1.2.3.5"),
// 					},
// 					&DomainRecordGeoIpMatchArgs{
// 						Continents: pulumi.StringArray{
// 							pulumi.String("NA"),
// 						},
// 						Data: pulumi.String("4.3.2.1"),
// 					},
// 				},
// 			},
// 			Ttl:  pulumi.Int(3600),
// 			Type: pulumi.String("A"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "httpService", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.4"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			HttpService: &DomainRecordHttpServiceArgs{
// 				Ips: pulumi.StringArray{
// 					pulumi.String("1.2.3.5"),
// 					pulumi.String("1.2.3.6"),
// 				},
// 				MustContain: pulumi.String("up"),
// 				Strategy:    pulumi.String("hashed"),
// 				Url:         pulumi.String("http://mywebsite.com/health"),
// 				UserAgent:   pulumi.String("scw_service_up"),
// 			},
// 			Ttl:  pulumi.Int(3600),
// 			Type: pulumi.String("A"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "view", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.4"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			Ttl:     pulumi.Int(3600),
// 			Type:    pulumi.String("A"),
// 			Views: DomainRecordViewArray{
// 				&DomainRecordViewArgs{
// 					Data:   pulumi.String("1.2.3.5"),
// 					Subnet: pulumi.String("100.0.0.0/16"),
// 				},
// 				&DomainRecordViewArgs{
// 					Data:   pulumi.String("1.2.3.6"),
// 					Subnet: pulumi.String("100.1.0.0/16"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "weighted", &scaleway.DomainRecordArgs{
// 			Data:    pulumi.String("1.2.3.4"),
// 			DnsZone: pulumi.String("domain.tld"),
// 			Ttl:     pulumi.Int(3600),
// 			Type:    pulumi.String("A"),
// 			Weighteds: DomainRecordWeightedArray{
// 				&DomainRecordWeightedArgs{
// 					Ip:     pulumi.String("1.2.3.5"),
// 					Weight: pulumi.Int(1),
// 				},
// 				&DomainRecordWeightedArgs{
// 					Ip:     pulumi.String("1.2.3.6"),
// 					Weight: pulumi.Int(2),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Create an instance and add records with the new instance IP
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/luxifer/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		projectId := cfg.Require("projectId")
// 		dnsZone := cfg.Require("dnsZone")
// 		publicIp, err := scaleway.NewInstanceIp(ctx, "publicIp", &scaleway.InstanceIpArgs{
// 			ProjectId: pulumi.String(projectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		web, err := scaleway.NewInstanceServer(ctx, "web", &scaleway.InstanceServerArgs{
// 			ProjectId: pulumi.String(projectId),
// 			Type:      pulumi.String("DEV1-S"),
// 			Image:     pulumi.String("ubuntu_jammy"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("front"),
// 				pulumi.String("web"),
// 			},
// 			IpId: publicIp.ID(),
// 			RootVolume: &InstanceServerRootVolumeArgs{
// 				SizeInGb: pulumi.Int(20),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "webA", &scaleway.DomainRecordArgs{
// 			DnsZone: pulumi.String(dnsZone),
// 			Type:    pulumi.String("A"),
// 			Data:    web.PublicIp,
// 			Ttl:     pulumi.Int(3600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "webCname", &scaleway.DomainRecordArgs{
// 			DnsZone: pulumi.String(dnsZone),
// 			Type:    pulumi.String("CNAME"),
// 			Data:    pulumi.String(fmt.Sprintf("web.%v.", dnsZone)),
// 			Ttl:     pulumi.Int(3600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.NewDomainRecord(ctx, "webAlias", &scaleway.DomainRecordArgs{
// 			DnsZone: pulumi.String(dnsZone),
// 			Type:    pulumi.String("ALIAS"),
// 			Data:    pulumi.String(fmt.Sprintf("web.%v.", dnsZone)),
// 			Ttl:     pulumi.Int(3600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Multiple records
//
// Some record types can have multiple `data` with the same `name` (eg: `A`, `AAAA`, `MX`, `NS`...).\
// You can duplicate a resource `DomainRecord` with the same `name`, the records will be added.
//
// Please note, some record (eg: `CNAME`, Multiple dynamic records of different types...) has to be unique.
//
// ## Import
//
// Record can be imported using the `{dns_zone}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/domainRecord:DomainRecord www subdomain.domain.tld/11111111-1111-1111-1111-111111111111
// ```
type DomainRecord struct {
	pulumi.CustomResourceState

	// The data of the view record
	Data pulumi.StringOutput `pulumi:"data"`
	// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
	DnsZone pulumi.StringOutput `pulumi:"dnsZone"`
	// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
	GeoIp DomainRecordGeoIpPtrOutput `pulumi:"geoIp"`
	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
	HttpService DomainRecordHttpServicePtrOutput `pulumi:"httpService"`
	// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
	KeepEmptyZone pulumi.BoolPtrOutput `pulumi:"keepEmptyZone"`
	// The name of the record (can be an empty string for a root record).
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the record (mostly used with an `MX` record)
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Does the DNS zone is the root zone or not
	RootZone pulumi.BoolOutput `pulumi:"rootZone"`
	// Time To Live of the record in seconds.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	Type pulumi.StringOutput `pulumi:"type"`
	// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
	Views DomainRecordViewArrayOutput `pulumi:"views"`
	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
	Weighteds DomainRecordWeightedArrayOutput `pulumi:"weighteds"`
}

// NewDomainRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainRecord(ctx *pulumi.Context,
	name string, args *DomainRecordArgs, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.DnsZone == nil {
		return nil, errors.New("invalid value for required argument 'DnsZone'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DomainRecord
	err := ctx.RegisterResource("scaleway:index/domainRecord:DomainRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainRecord gets an existing DomainRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainRecordState, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	var resource DomainRecord
	err := ctx.ReadResource("scaleway:index/domainRecord:DomainRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainRecord resources.
type domainRecordState struct {
	// The data of the view record
	Data *string `pulumi:"data"`
	// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
	DnsZone *string `pulumi:"dnsZone"`
	// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
	GeoIp *DomainRecordGeoIp `pulumi:"geoIp"`
	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
	HttpService *DomainRecordHttpService `pulumi:"httpService"`
	// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
	KeepEmptyZone *bool `pulumi:"keepEmptyZone"`
	// The name of the record (can be an empty string for a root record).
	Name *string `pulumi:"name"`
	// The priority of the record (mostly used with an `MX` record)
	Priority *int `pulumi:"priority"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Does the DNS zone is the root zone or not
	RootZone *bool `pulumi:"rootZone"`
	// Time To Live of the record in seconds.
	Ttl *int `pulumi:"ttl"`
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	Type *string `pulumi:"type"`
	// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
	Views []DomainRecordView `pulumi:"views"`
	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
	Weighteds []DomainRecordWeighted `pulumi:"weighteds"`
}

type DomainRecordState struct {
	// The data of the view record
	Data pulumi.StringPtrInput
	// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
	DnsZone pulumi.StringPtrInput
	// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
	GeoIp DomainRecordGeoIpPtrInput
	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
	HttpService DomainRecordHttpServicePtrInput
	// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
	KeepEmptyZone pulumi.BoolPtrInput
	// The name of the record (can be an empty string for a root record).
	Name pulumi.StringPtrInput
	// The priority of the record (mostly used with an `MX` record)
	Priority pulumi.IntPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Does the DNS zone is the root zone or not
	RootZone pulumi.BoolPtrInput
	// Time To Live of the record in seconds.
	Ttl pulumi.IntPtrInput
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	Type pulumi.StringPtrInput
	// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
	Views DomainRecordViewArrayInput
	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
	Weighteds DomainRecordWeightedArrayInput
}

func (DomainRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordState)(nil)).Elem()
}

type domainRecordArgs struct {
	// The data of the view record
	Data string `pulumi:"data"`
	// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
	DnsZone string `pulumi:"dnsZone"`
	// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
	GeoIp *DomainRecordGeoIp `pulumi:"geoIp"`
	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
	HttpService *DomainRecordHttpService `pulumi:"httpService"`
	// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
	KeepEmptyZone *bool `pulumi:"keepEmptyZone"`
	// The name of the record (can be an empty string for a root record).
	Name *string `pulumi:"name"`
	// The priority of the record (mostly used with an `MX` record)
	Priority *int `pulumi:"priority"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Time To Live of the record in seconds.
	Ttl *int `pulumi:"ttl"`
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	Type string `pulumi:"type"`
	// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
	Views []DomainRecordView `pulumi:"views"`
	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
	Weighteds []DomainRecordWeighted `pulumi:"weighteds"`
}

// The set of arguments for constructing a DomainRecord resource.
type DomainRecordArgs struct {
	// The data of the view record
	Data pulumi.StringInput
	// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
	DnsZone pulumi.StringInput
	// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
	GeoIp DomainRecordGeoIpPtrInput
	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
	HttpService DomainRecordHttpServicePtrInput
	// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
	KeepEmptyZone pulumi.BoolPtrInput
	// The name of the record (can be an empty string for a root record).
	Name pulumi.StringPtrInput
	// The priority of the record (mostly used with an `MX` record)
	Priority pulumi.IntPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Time To Live of the record in seconds.
	Ttl pulumi.IntPtrInput
	// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
	Type pulumi.StringInput
	// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
	Views DomainRecordViewArrayInput
	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
	Weighteds DomainRecordWeightedArrayInput
}

func (DomainRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordArgs)(nil)).Elem()
}

type DomainRecordInput interface {
	pulumi.Input

	ToDomainRecordOutput() DomainRecordOutput
	ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput
}

func (*DomainRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecord)(nil)).Elem()
}

func (i *DomainRecord) ToDomainRecordOutput() DomainRecordOutput {
	return i.ToDomainRecordOutputWithContext(context.Background())
}

func (i *DomainRecord) ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRecordOutput)
}

// DomainRecordArrayInput is an input type that accepts DomainRecordArray and DomainRecordArrayOutput values.
// You can construct a concrete instance of `DomainRecordArrayInput` via:
//
//          DomainRecordArray{ DomainRecordArgs{...} }
type DomainRecordArrayInput interface {
	pulumi.Input

	ToDomainRecordArrayOutput() DomainRecordArrayOutput
	ToDomainRecordArrayOutputWithContext(context.Context) DomainRecordArrayOutput
}

type DomainRecordArray []DomainRecordInput

func (DomainRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainRecord)(nil)).Elem()
}

func (i DomainRecordArray) ToDomainRecordArrayOutput() DomainRecordArrayOutput {
	return i.ToDomainRecordArrayOutputWithContext(context.Background())
}

func (i DomainRecordArray) ToDomainRecordArrayOutputWithContext(ctx context.Context) DomainRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRecordArrayOutput)
}

// DomainRecordMapInput is an input type that accepts DomainRecordMap and DomainRecordMapOutput values.
// You can construct a concrete instance of `DomainRecordMapInput` via:
//
//          DomainRecordMap{ "key": DomainRecordArgs{...} }
type DomainRecordMapInput interface {
	pulumi.Input

	ToDomainRecordMapOutput() DomainRecordMapOutput
	ToDomainRecordMapOutputWithContext(context.Context) DomainRecordMapOutput
}

type DomainRecordMap map[string]DomainRecordInput

func (DomainRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainRecord)(nil)).Elem()
}

func (i DomainRecordMap) ToDomainRecordMapOutput() DomainRecordMapOutput {
	return i.ToDomainRecordMapOutputWithContext(context.Background())
}

func (i DomainRecordMap) ToDomainRecordMapOutputWithContext(ctx context.Context) DomainRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainRecordMapOutput)
}

type DomainRecordOutput struct{ *pulumi.OutputState }

func (DomainRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRecord)(nil)).Elem()
}

func (o DomainRecordOutput) ToDomainRecordOutput() DomainRecordOutput {
	return o
}

func (o DomainRecordOutput) ToDomainRecordOutputWithContext(ctx context.Context) DomainRecordOutput {
	return o
}

// The data of the view record
func (o DomainRecordOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
func (o DomainRecordOutput) DnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringOutput { return v.DnsZone }).(pulumi.StringOutput)
}

// The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#geo-ip-records)
func (o DomainRecordOutput) GeoIp() DomainRecordGeoIpPtrOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecordGeoIpPtrOutput { return v.GeoIp }).(DomainRecordGeoIpPtrOutput)
}

// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#healthcheck-records)
func (o DomainRecordOutput) HttpService() DomainRecordHttpServicePtrOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecordHttpServicePtrOutput { return v.HttpService }).(DomainRecordHttpServicePtrOutput)
}

// When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
func (o DomainRecordOutput) KeepEmptyZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolPtrOutput { return v.KeepEmptyZone }).(pulumi.BoolPtrOutput)
}

// The name of the record (can be an empty string for a root record).
func (o DomainRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The priority of the record (mostly used with an `MX` record)
func (o DomainRecordOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The project_id you want to attach the resource to
func (o DomainRecordOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Does the DNS zone is the root zone or not
func (o DomainRecordOutput) RootZone() pulumi.BoolOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.BoolOutput { return v.RootZone }).(pulumi.BoolOutput)
}

// Time To Live of the record in seconds.
func (o DomainRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
func (o DomainRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainRecord) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#views-records)
func (o DomainRecordOutput) Views() DomainRecordViewArrayOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecordViewArrayOutput { return v.Views }).(DomainRecordViewArrayOutput)
}

// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/#weight-records)
func (o DomainRecordOutput) Weighteds() DomainRecordWeightedArrayOutput {
	return o.ApplyT(func(v *DomainRecord) DomainRecordWeightedArrayOutput { return v.Weighteds }).(DomainRecordWeightedArrayOutput)
}

type DomainRecordArrayOutput struct{ *pulumi.OutputState }

func (DomainRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainRecord)(nil)).Elem()
}

func (o DomainRecordArrayOutput) ToDomainRecordArrayOutput() DomainRecordArrayOutput {
	return o
}

func (o DomainRecordArrayOutput) ToDomainRecordArrayOutputWithContext(ctx context.Context) DomainRecordArrayOutput {
	return o
}

func (o DomainRecordArrayOutput) Index(i pulumi.IntInput) DomainRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainRecord {
		return vs[0].([]*DomainRecord)[vs[1].(int)]
	}).(DomainRecordOutput)
}

type DomainRecordMapOutput struct{ *pulumi.OutputState }

func (DomainRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainRecord)(nil)).Elem()
}

func (o DomainRecordMapOutput) ToDomainRecordMapOutput() DomainRecordMapOutput {
	return o
}

func (o DomainRecordMapOutput) ToDomainRecordMapOutputWithContext(ctx context.Context) DomainRecordMapOutput {
	return o
}

func (o DomainRecordMapOutput) MapIndex(k pulumi.StringInput) DomainRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainRecord {
		return vs[0].(map[string]*DomainRecord)[vs[1].(string)]
	}).(DomainRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordInput)(nil)).Elem(), &DomainRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordArrayInput)(nil)).Elem(), DomainRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRecordMapInput)(nil)).Elem(), DomainRecordMap{})
	pulumi.RegisterOutputType(DomainRecordOutput{})
	pulumi.RegisterOutputType(DomainRecordArrayOutput{})
	pulumi.RegisterOutputType(DomainRecordMapOutput{})
}