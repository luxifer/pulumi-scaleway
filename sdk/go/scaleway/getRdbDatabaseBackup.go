// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an RDB backup.
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
// 		_, err := scaleway.LookupRdbDatabaseBackup(ctx, &GetRdbDatabaseBackupArgs{
// 			Name: pulumi.StringRef("mybackup"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.LookupRdbDatabaseBackup(ctx, &GetRdbDatabaseBackupArgs{
// 			InstanceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
// 			Name:       pulumi.StringRef("mybackup"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = scaleway.LookupRdbDatabaseBackup(ctx, &GetRdbDatabaseBackupArgs{
// 			BackupId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRdbDatabaseBackup(ctx *pulumi.Context, args *LookupRdbDatabaseBackupArgs, opts ...pulumi.InvokeOption) (*LookupRdbDatabaseBackupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRdbDatabaseBackupResult
	err := ctx.Invoke("scaleway:index/getRdbDatabaseBackup:getRdbDatabaseBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdbDatabaseBackup.
type LookupRdbDatabaseBackupArgs struct {
	// The RDB backup ID.
	// Only one of the `name` and `backupId` should be specified.
	BackupId *string `pulumi:"backupId"`
	// The RDB instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the RDB instance.
	// Only one of the `name` and `backupId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getRdbDatabaseBackup.
type LookupRdbDatabaseBackupResult struct {
	BackupId     *string `pulumi:"backupId"`
	CreatedAt    string  `pulumi:"createdAt"`
	DatabaseName string  `pulumi:"databaseName"`
	ExpiresAt    string  `pulumi:"expiresAt"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	InstanceId   *string `pulumi:"instanceId"`
	InstanceName string  `pulumi:"instanceName"`
	Name         *string `pulumi:"name"`
	Region       *string `pulumi:"region"`
	Size         int     `pulumi:"size"`
	UpdatedAt    string  `pulumi:"updatedAt"`
}

func LookupRdbDatabaseBackupOutput(ctx *pulumi.Context, args LookupRdbDatabaseBackupOutputArgs, opts ...pulumi.InvokeOption) LookupRdbDatabaseBackupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRdbDatabaseBackupResult, error) {
			args := v.(LookupRdbDatabaseBackupArgs)
			r, err := LookupRdbDatabaseBackup(ctx, &args, opts...)
			var s LookupRdbDatabaseBackupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRdbDatabaseBackupResultOutput)
}

// A collection of arguments for invoking getRdbDatabaseBackup.
type LookupRdbDatabaseBackupOutputArgs struct {
	// The RDB backup ID.
	// Only one of the `name` and `backupId` should be specified.
	BackupId pulumi.StringPtrInput `pulumi:"backupId"`
	// The RDB instance ID.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the RDB instance.
	// Only one of the `name` and `backupId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupRdbDatabaseBackupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbDatabaseBackupArgs)(nil)).Elem()
}

// A collection of values returned by getRdbDatabaseBackup.
type LookupRdbDatabaseBackupResultOutput struct{ *pulumi.OutputState }

func (LookupRdbDatabaseBackupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbDatabaseBackupResult)(nil)).Elem()
}

func (o LookupRdbDatabaseBackupResultOutput) ToLookupRdbDatabaseBackupResultOutput() LookupRdbDatabaseBackupResultOutput {
	return o
}

func (o LookupRdbDatabaseBackupResultOutput) ToLookupRdbDatabaseBackupResultOutputWithContext(ctx context.Context) LookupRdbDatabaseBackupResultOutput {
	return o
}

func (o LookupRdbDatabaseBackupResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRdbDatabaseBackupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o LookupRdbDatabaseBackupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseBackupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdbDatabaseBackupResultOutput{})
}
