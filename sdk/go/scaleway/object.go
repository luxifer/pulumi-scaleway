// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway object storage objects.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
//
// ## Import
//
// Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/object:Object some_object fr-par/some-bucket/some-file
// ```
type Object struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The name of the file to upload, defaults to an empty file
	File pulumi.StringPtrOutput `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrOutput `pulumi:"hash"`
	// The path of the object.
	Key pulumi.StringOutput `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// Map of tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewObject registers a new resource with the given unique name, arguments, and options.
func NewObject(ctx *pulumi.Context,
	name string, args *ObjectArgs, opts ...pulumi.ResourceOption) (*Object, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Object
	err := ctx.RegisterResource("scaleway:index/object:Object", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObject gets an existing Object resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectState, opts ...pulumi.ResourceOption) (*Object, error) {
	var resource Object
	err := ctx.ReadResource("scaleway:index/object:Object", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Object resources.
type objectState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The name of the file to upload, defaults to an empty file
	File *string `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash *string `pulumi:"hash"`
	// The path of the object.
	Key *string `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata map[string]string `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region *string `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass *string `pulumi:"storageClass"`
	// Map of tags
	Tags map[string]string `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility *string `pulumi:"visibility"`
}

type ObjectState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The name of the file to upload, defaults to an empty file
	File pulumi.StringPtrInput
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrInput
	// The path of the object.
	Key pulumi.StringPtrInput
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The Scaleway region this bucket resides in.
	Region pulumi.StringPtrInput
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrInput
	// Map of tags
	Tags pulumi.StringMapInput
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringPtrInput
}

func (ObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectState)(nil)).Elem()
}

type objectArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The name of the file to upload, defaults to an empty file
	File *string `pulumi:"file"`
	// Hash of the file, used to trigger upload on file change
	Hash *string `pulumi:"hash"`
	// The path of the object.
	Key string `pulumi:"key"`
	// Map of metadata used for the object, keys must be lowercase
	Metadata map[string]string `pulumi:"metadata"`
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The Scaleway region this bucket resides in.
	Region *string `pulumi:"region"`
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass *string `pulumi:"storageClass"`
	// Map of tags
	Tags map[string]string `pulumi:"tags"`
	// Visibility of the object, `public-read` or `private`
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Object resource.
type ObjectArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The name of the file to upload, defaults to an empty file
	File pulumi.StringPtrInput
	// Hash of the file, used to trigger upload on file change
	Hash pulumi.StringPtrInput
	// The path of the object.
	Key pulumi.StringInput
	// Map of metadata used for the object, keys must be lowercase
	Metadata pulumi.StringMapInput
	// `projectId`) The ID of the project the bucket is associated with.
	ProjectId pulumi.StringPtrInput
	// The Scaleway region this bucket resides in.
	Region pulumi.StringPtrInput
	// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
	StorageClass pulumi.StringPtrInput
	// Map of tags
	Tags pulumi.StringMapInput
	// Visibility of the object, `public-read` or `private`
	Visibility pulumi.StringPtrInput
}

func (ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectArgs)(nil)).Elem()
}

type ObjectInput interface {
	pulumi.Input

	ToObjectOutput() ObjectOutput
	ToObjectOutputWithContext(ctx context.Context) ObjectOutput
}

func (*Object) ElementType() reflect.Type {
	return reflect.TypeOf((**Object)(nil)).Elem()
}

func (i *Object) ToObjectOutput() ObjectOutput {
	return i.ToObjectOutputWithContext(context.Background())
}

func (i *Object) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectOutput)
}

// ObjectArrayInput is an input type that accepts ObjectArray and ObjectArrayOutput values.
// You can construct a concrete instance of `ObjectArrayInput` via:
//
//          ObjectArray{ ObjectArgs{...} }
type ObjectArrayInput interface {
	pulumi.Input

	ToObjectArrayOutput() ObjectArrayOutput
	ToObjectArrayOutputWithContext(context.Context) ObjectArrayOutput
}

type ObjectArray []ObjectInput

func (ObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Object)(nil)).Elem()
}

func (i ObjectArray) ToObjectArrayOutput() ObjectArrayOutput {
	return i.ToObjectArrayOutputWithContext(context.Background())
}

func (i ObjectArray) ToObjectArrayOutputWithContext(ctx context.Context) ObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectArrayOutput)
}

// ObjectMapInput is an input type that accepts ObjectMap and ObjectMapOutput values.
// You can construct a concrete instance of `ObjectMapInput` via:
//
//          ObjectMap{ "key": ObjectArgs{...} }
type ObjectMapInput interface {
	pulumi.Input

	ToObjectMapOutput() ObjectMapOutput
	ToObjectMapOutputWithContext(context.Context) ObjectMapOutput
}

type ObjectMap map[string]ObjectInput

func (ObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Object)(nil)).Elem()
}

func (i ObjectMap) ToObjectMapOutput() ObjectMapOutput {
	return i.ToObjectMapOutputWithContext(context.Background())
}

func (i ObjectMap) ToObjectMapOutputWithContext(ctx context.Context) ObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectMapOutput)
}

type ObjectOutput struct{ *pulumi.OutputState }

func (ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Object)(nil)).Elem()
}

func (o ObjectOutput) ToObjectOutput() ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return o
}

// The name of the bucket.
func (o ObjectOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Object) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The name of the file to upload, defaults to an empty file
func (o ObjectOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Object) pulumi.StringPtrOutput { return v.File }).(pulumi.StringPtrOutput)
}

// Hash of the file, used to trigger upload on file change
func (o ObjectOutput) Hash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Object) pulumi.StringPtrOutput { return v.Hash }).(pulumi.StringPtrOutput)
}

// The path of the object.
func (o ObjectOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Object) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Map of metadata used for the object, keys must be lowercase
func (o ObjectOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Object) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// `projectId`) The ID of the project the bucket is associated with.
func (o ObjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Object) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The Scaleway region this bucket resides in.
func (o ObjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Object) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
func (o ObjectOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Object) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

// Map of tags
func (o ObjectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Object) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Visibility of the object, `public-read` or `private`
func (o ObjectOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Object) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ObjectArrayOutput struct{ *pulumi.OutputState }

func (ObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Object)(nil)).Elem()
}

func (o ObjectArrayOutput) ToObjectArrayOutput() ObjectArrayOutput {
	return o
}

func (o ObjectArrayOutput) ToObjectArrayOutputWithContext(ctx context.Context) ObjectArrayOutput {
	return o
}

func (o ObjectArrayOutput) Index(i pulumi.IntInput) ObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Object {
		return vs[0].([]*Object)[vs[1].(int)]
	}).(ObjectOutput)
}

type ObjectMapOutput struct{ *pulumi.OutputState }

func (ObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Object)(nil)).Elem()
}

func (o ObjectMapOutput) ToObjectMapOutput() ObjectMapOutput {
	return o
}

func (o ObjectMapOutput) ToObjectMapOutputWithContext(ctx context.Context) ObjectMapOutput {
	return o
}

func (o ObjectMapOutput) MapIndex(k pulumi.StringInput) ObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Object {
		return vs[0].(map[string]*Object)[vs[1].(string)]
	}).(ObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectInput)(nil)).Elem(), &Object{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectArrayInput)(nil)).Elem(), ObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectMapInput)(nil)).Elem(), ObjectMap{})
	pulumi.RegisterOutputType(ObjectOutput{})
	pulumi.RegisterOutputType(ObjectArrayOutput{})
	pulumi.RegisterOutputType(ObjectMapOutput{})
}
