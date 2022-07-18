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
// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
type ValidatingWebhookConfigurationPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookPatchArrayOutput `pulumi:"webhooks"`
}

// NewValidatingWebhookConfigurationPatch registers a new resource with the given unique name, arguments, and options.
func NewValidatingWebhookConfigurationPatch(ctx *pulumi.Context,
	name string, args *ValidatingWebhookConfigurationPatchArgs, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfigurationPatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1")
	args.Kind = pulumi.StringPtr("ValidatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource ValidatingWebhookConfigurationPatch
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfigurationPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidatingWebhookConfigurationPatch gets an existing ValidatingWebhookConfigurationPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidatingWebhookConfigurationPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidatingWebhookConfigurationPatchState, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfigurationPatch, error) {
	var resource ValidatingWebhookConfigurationPatch
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfigurationPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidatingWebhookConfigurationPatch resources.
type validatingWebhookConfigurationPatchState struct {
}

type ValidatingWebhookConfigurationPatchState struct {
}

func (ValidatingWebhookConfigurationPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationPatchState)(nil)).Elem()
}

type validatingWebhookConfigurationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []ValidatingWebhookPatch `pulumi:"webhooks"`
}

// The set of arguments for constructing a ValidatingWebhookConfigurationPatch resource.
type ValidatingWebhookConfigurationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookPatchArrayInput
}

func (ValidatingWebhookConfigurationPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationPatchArgs)(nil)).Elem()
}

type ValidatingWebhookConfigurationPatchInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationPatchOutput() ValidatingWebhookConfigurationPatchOutput
	ToValidatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchOutput
}

func (*ValidatingWebhookConfigurationPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i *ValidatingWebhookConfigurationPatch) ToValidatingWebhookConfigurationPatchOutput() ValidatingWebhookConfigurationPatchOutput {
	return i.ToValidatingWebhookConfigurationPatchOutputWithContext(context.Background())
}

func (i *ValidatingWebhookConfigurationPatch) ToValidatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationPatchOutput)
}

// ValidatingWebhookConfigurationPatchArrayInput is an input type that accepts ValidatingWebhookConfigurationPatchArray and ValidatingWebhookConfigurationPatchArrayOutput values.
// You can construct a concrete instance of `ValidatingWebhookConfigurationPatchArrayInput` via:
//
//          ValidatingWebhookConfigurationPatchArray{ ValidatingWebhookConfigurationPatchArgs{...} }
type ValidatingWebhookConfigurationPatchArrayInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationPatchArrayOutput() ValidatingWebhookConfigurationPatchArrayOutput
	ToValidatingWebhookConfigurationPatchArrayOutputWithContext(context.Context) ValidatingWebhookConfigurationPatchArrayOutput
}

type ValidatingWebhookConfigurationPatchArray []ValidatingWebhookConfigurationPatchInput

func (ValidatingWebhookConfigurationPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i ValidatingWebhookConfigurationPatchArray) ToValidatingWebhookConfigurationPatchArrayOutput() ValidatingWebhookConfigurationPatchArrayOutput {
	return i.ToValidatingWebhookConfigurationPatchArrayOutputWithContext(context.Background())
}

func (i ValidatingWebhookConfigurationPatchArray) ToValidatingWebhookConfigurationPatchArrayOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationPatchArrayOutput)
}

// ValidatingWebhookConfigurationPatchMapInput is an input type that accepts ValidatingWebhookConfigurationPatchMap and ValidatingWebhookConfigurationPatchMapOutput values.
// You can construct a concrete instance of `ValidatingWebhookConfigurationPatchMapInput` via:
//
//          ValidatingWebhookConfigurationPatchMap{ "key": ValidatingWebhookConfigurationPatchArgs{...} }
type ValidatingWebhookConfigurationPatchMapInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationPatchMapOutput() ValidatingWebhookConfigurationPatchMapOutput
	ToValidatingWebhookConfigurationPatchMapOutputWithContext(context.Context) ValidatingWebhookConfigurationPatchMapOutput
}

type ValidatingWebhookConfigurationPatchMap map[string]ValidatingWebhookConfigurationPatchInput

func (ValidatingWebhookConfigurationPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i ValidatingWebhookConfigurationPatchMap) ToValidatingWebhookConfigurationPatchMapOutput() ValidatingWebhookConfigurationPatchMapOutput {
	return i.ToValidatingWebhookConfigurationPatchMapOutputWithContext(context.Background())
}

func (i ValidatingWebhookConfigurationPatchMap) ToValidatingWebhookConfigurationPatchMapOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationPatchMapOutput)
}

type ValidatingWebhookConfigurationPatchOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o ValidatingWebhookConfigurationPatchOutput) ToValidatingWebhookConfigurationPatchOutput() ValidatingWebhookConfigurationPatchOutput {
	return o
}

func (o ValidatingWebhookConfigurationPatchOutput) ToValidatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ValidatingWebhookConfigurationPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatingWebhookConfigurationPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ValidatingWebhookConfigurationPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatingWebhookConfigurationPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
func (o ValidatingWebhookConfigurationPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ValidatingWebhookConfigurationPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Webhooks is a list of webhooks and the affected resources and operations.
func (o ValidatingWebhookConfigurationPatchOutput) Webhooks() ValidatingWebhookPatchArrayOutput {
	return o.ApplyT(func(v *ValidatingWebhookConfigurationPatch) ValidatingWebhookPatchArrayOutput { return v.Webhooks }).(ValidatingWebhookPatchArrayOutput)
}

type ValidatingWebhookConfigurationPatchArrayOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o ValidatingWebhookConfigurationPatchArrayOutput) ToValidatingWebhookConfigurationPatchArrayOutput() ValidatingWebhookConfigurationPatchArrayOutput {
	return o
}

func (o ValidatingWebhookConfigurationPatchArrayOutput) ToValidatingWebhookConfigurationPatchArrayOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchArrayOutput {
	return o
}

func (o ValidatingWebhookConfigurationPatchArrayOutput) Index(i pulumi.IntInput) ValidatingWebhookConfigurationPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ValidatingWebhookConfigurationPatch {
		return vs[0].([]*ValidatingWebhookConfigurationPatch)[vs[1].(int)]
	}).(ValidatingWebhookConfigurationPatchOutput)
}

type ValidatingWebhookConfigurationPatchMapOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o ValidatingWebhookConfigurationPatchMapOutput) ToValidatingWebhookConfigurationPatchMapOutput() ValidatingWebhookConfigurationPatchMapOutput {
	return o
}

func (o ValidatingWebhookConfigurationPatchMapOutput) ToValidatingWebhookConfigurationPatchMapOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPatchMapOutput {
	return o
}

func (o ValidatingWebhookConfigurationPatchMapOutput) MapIndex(k pulumi.StringInput) ValidatingWebhookConfigurationPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ValidatingWebhookConfigurationPatch {
		return vs[0].(map[string]*ValidatingWebhookConfigurationPatch)[vs[1].(string)]
	}).(ValidatingWebhookConfigurationPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationPatchInput)(nil)).Elem(), &ValidatingWebhookConfigurationPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationPatchArrayInput)(nil)).Elem(), ValidatingWebhookConfigurationPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationPatchMapInput)(nil)).Elem(), ValidatingWebhookConfigurationPatchMap{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationPatchOutput{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationPatchArrayOutput{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationPatchMapOutput{})
}
