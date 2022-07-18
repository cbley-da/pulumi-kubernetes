// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/installation-configuration/#server-side-apply) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// DaemonSet represents the configuration of a daemon set.
type DaemonSetPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecPatchPtrOutput `pulumi:"spec"`
	// The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status DaemonSetStatusPatchPtrOutput `pulumi:"status"`
}

// NewDaemonSetPatch registers a new resource with the given unique name, arguments, and options.
func NewDaemonSetPatch(ctx *pulumi.Context,
	name string, args *DaemonSetPatchArgs, opts ...pulumi.ResourceOption) (*DaemonSetPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1")
	args.Kind = pulumi.StringPtr("DaemonSet")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1beta2:DaemonSetPatch"),
		},
		{
			Type: pulumi.String("kubernetes:extensions/v1beta1:DaemonSetPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource DaemonSetPatch
	err := ctx.RegisterResource("kubernetes:apps/v1:DaemonSetPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDaemonSetPatch gets an existing DaemonSetPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDaemonSetPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DaemonSetPatchState, opts ...pulumi.ResourceOption) (*DaemonSetPatch, error) {
	var resource DaemonSetPatch
	err := ctx.ReadResource("kubernetes:apps/v1:DaemonSetPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DaemonSetPatch resources.
type daemonSetPatchState struct {
}

type DaemonSetPatchState struct {
}

func (DaemonSetPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetPatchState)(nil)).Elem()
}

type daemonSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DaemonSetSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a DaemonSetPatch resource.
type DaemonSetPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchInput
	// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecPatchPtrInput
}

func (DaemonSetPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*daemonSetPatchArgs)(nil)).Elem()
}

type DaemonSetPatchInput interface {
	pulumi.Input

	ToDaemonSetPatchOutput() DaemonSetPatchOutput
	ToDaemonSetPatchOutputWithContext(ctx context.Context) DaemonSetPatchOutput
}

func (*DaemonSetPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSetPatch)(nil)).Elem()
}

func (i *DaemonSetPatch) ToDaemonSetPatchOutput() DaemonSetPatchOutput {
	return i.ToDaemonSetPatchOutputWithContext(context.Background())
}

func (i *DaemonSetPatch) ToDaemonSetPatchOutputWithContext(ctx context.Context) DaemonSetPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetPatchOutput)
}

// DaemonSetPatchArrayInput is an input type that accepts DaemonSetPatchArray and DaemonSetPatchArrayOutput values.
// You can construct a concrete instance of `DaemonSetPatchArrayInput` via:
//
//          DaemonSetPatchArray{ DaemonSetPatchArgs{...} }
type DaemonSetPatchArrayInput interface {
	pulumi.Input

	ToDaemonSetPatchArrayOutput() DaemonSetPatchArrayOutput
	ToDaemonSetPatchArrayOutputWithContext(context.Context) DaemonSetPatchArrayOutput
}

type DaemonSetPatchArray []DaemonSetPatchInput

func (DaemonSetPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSetPatch)(nil)).Elem()
}

func (i DaemonSetPatchArray) ToDaemonSetPatchArrayOutput() DaemonSetPatchArrayOutput {
	return i.ToDaemonSetPatchArrayOutputWithContext(context.Background())
}

func (i DaemonSetPatchArray) ToDaemonSetPatchArrayOutputWithContext(ctx context.Context) DaemonSetPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetPatchArrayOutput)
}

// DaemonSetPatchMapInput is an input type that accepts DaemonSetPatchMap and DaemonSetPatchMapOutput values.
// You can construct a concrete instance of `DaemonSetPatchMapInput` via:
//
//          DaemonSetPatchMap{ "key": DaemonSetPatchArgs{...} }
type DaemonSetPatchMapInput interface {
	pulumi.Input

	ToDaemonSetPatchMapOutput() DaemonSetPatchMapOutput
	ToDaemonSetPatchMapOutputWithContext(context.Context) DaemonSetPatchMapOutput
}

type DaemonSetPatchMap map[string]DaemonSetPatchInput

func (DaemonSetPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSetPatch)(nil)).Elem()
}

func (i DaemonSetPatchMap) ToDaemonSetPatchMapOutput() DaemonSetPatchMapOutput {
	return i.ToDaemonSetPatchMapOutputWithContext(context.Background())
}

func (i DaemonSetPatchMap) ToDaemonSetPatchMapOutputWithContext(ctx context.Context) DaemonSetPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaemonSetPatchMapOutput)
}

type DaemonSetPatchOutput struct{ *pulumi.OutputState }

func (DaemonSetPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaemonSetPatch)(nil)).Elem()
}

func (o DaemonSetPatchOutput) ToDaemonSetPatchOutput() DaemonSetPatchOutput {
	return o
}

func (o DaemonSetPatchOutput) ToDaemonSetPatchOutputWithContext(ctx context.Context) DaemonSetPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o DaemonSetPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaemonSetPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o DaemonSetPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaemonSetPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o DaemonSetPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *DaemonSetPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o DaemonSetPatchOutput) Spec() DaemonSetSpecPatchPtrOutput {
	return o.ApplyT(func(v *DaemonSetPatch) DaemonSetSpecPatchPtrOutput { return v.Spec }).(DaemonSetSpecPatchPtrOutput)
}

// The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o DaemonSetPatchOutput) Status() DaemonSetStatusPatchPtrOutput {
	return o.ApplyT(func(v *DaemonSetPatch) DaemonSetStatusPatchPtrOutput { return v.Status }).(DaemonSetStatusPatchPtrOutput)
}

type DaemonSetPatchArrayOutput struct{ *pulumi.OutputState }

func (DaemonSetPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DaemonSetPatch)(nil)).Elem()
}

func (o DaemonSetPatchArrayOutput) ToDaemonSetPatchArrayOutput() DaemonSetPatchArrayOutput {
	return o
}

func (o DaemonSetPatchArrayOutput) ToDaemonSetPatchArrayOutputWithContext(ctx context.Context) DaemonSetPatchArrayOutput {
	return o
}

func (o DaemonSetPatchArrayOutput) Index(i pulumi.IntInput) DaemonSetPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DaemonSetPatch {
		return vs[0].([]*DaemonSetPatch)[vs[1].(int)]
	}).(DaemonSetPatchOutput)
}

type DaemonSetPatchMapOutput struct{ *pulumi.OutputState }

func (DaemonSetPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DaemonSetPatch)(nil)).Elem()
}

func (o DaemonSetPatchMapOutput) ToDaemonSetPatchMapOutput() DaemonSetPatchMapOutput {
	return o
}

func (o DaemonSetPatchMapOutput) ToDaemonSetPatchMapOutputWithContext(ctx context.Context) DaemonSetPatchMapOutput {
	return o
}

func (o DaemonSetPatchMapOutput) MapIndex(k pulumi.StringInput) DaemonSetPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DaemonSetPatch {
		return vs[0].(map[string]*DaemonSetPatch)[vs[1].(string)]
	}).(DaemonSetPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetPatchInput)(nil)).Elem(), &DaemonSetPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetPatchArrayInput)(nil)).Elem(), DaemonSetPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DaemonSetPatchMapInput)(nil)).Elem(), DaemonSetPatchMap{})
	pulumi.RegisterOutputType(DaemonSetPatchOutput{})
	pulumi.RegisterOutputType(DaemonSetPatchArrayOutput{})
	pulumi.RegisterOutputType(DaemonSetPatchMapOutput{})
}
