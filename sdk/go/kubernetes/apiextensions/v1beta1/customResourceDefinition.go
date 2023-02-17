// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format <.spec.name>.<.spec.group>. Deprecated in v1.16, planned for removal in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
type CustomResourceDefinition struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecOutput `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPtrOutput `pulumi:"status"`
}

// NewCustomResourceDefinition registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinition(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("apiextensions.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CustomResourceDefinition")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomResourceDefinition
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinition gets an existing CustomResourceDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionState, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	var resource CustomResourceDefinition
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinition resources.
type customResourceDefinitionState struct {
}

type CustomResourceDefinitionState struct {
}

func (CustomResourceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionState)(nil)).Elem()
}

type customResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CustomResourceDefinition resource.
type CustomResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecInput
}

func (CustomResourceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionArgs)(nil)).Elem()
}

type CustomResourceDefinitionInput interface {
	pulumi.Input

	ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput
	ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput
}

func (*CustomResourceDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinition)(nil)).Elem()
}

func (i *CustomResourceDefinition) ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput {
	return i.ToCustomResourceDefinitionOutputWithContext(context.Background())
}

func (i *CustomResourceDefinition) ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionOutput)
}

// CustomResourceDefinitionArrayInput is an input type that accepts CustomResourceDefinitionArray and CustomResourceDefinitionArrayOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionArrayInput` via:
//
//	CustomResourceDefinitionArray{ CustomResourceDefinitionArgs{...} }
type CustomResourceDefinitionArrayInput interface {
	pulumi.Input

	ToCustomResourceDefinitionArrayOutput() CustomResourceDefinitionArrayOutput
	ToCustomResourceDefinitionArrayOutputWithContext(context.Context) CustomResourceDefinitionArrayOutput
}

type CustomResourceDefinitionArray []CustomResourceDefinitionInput

func (CustomResourceDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinition)(nil)).Elem()
}

func (i CustomResourceDefinitionArray) ToCustomResourceDefinitionArrayOutput() CustomResourceDefinitionArrayOutput {
	return i.ToCustomResourceDefinitionArrayOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionArray) ToCustomResourceDefinitionArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionArrayOutput)
}

// CustomResourceDefinitionMapInput is an input type that accepts CustomResourceDefinitionMap and CustomResourceDefinitionMapOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionMapInput` via:
//
//	CustomResourceDefinitionMap{ "key": CustomResourceDefinitionArgs{...} }
type CustomResourceDefinitionMapInput interface {
	pulumi.Input

	ToCustomResourceDefinitionMapOutput() CustomResourceDefinitionMapOutput
	ToCustomResourceDefinitionMapOutputWithContext(context.Context) CustomResourceDefinitionMapOutput
}

type CustomResourceDefinitionMap map[string]CustomResourceDefinitionInput

func (CustomResourceDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinition)(nil)).Elem()
}

func (i CustomResourceDefinitionMap) ToCustomResourceDefinitionMapOutput() CustomResourceDefinitionMapOutput {
	return i.ToCustomResourceDefinitionMapOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionMap) ToCustomResourceDefinitionMapOutputWithContext(ctx context.Context) CustomResourceDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionMapOutput)
}

type CustomResourceDefinitionOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinition)(nil)).Elem()
}

func (o CustomResourceDefinitionOutput) ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput {
	return o
}

func (o CustomResourceDefinitionOutput) ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CustomResourceDefinitionOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinition) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CustomResourceDefinitionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinition) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CustomResourceDefinitionOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinition) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// spec describes how the user wants the resources to appear
func (o CustomResourceDefinitionOutput) Spec() CustomResourceDefinitionSpecOutput {
	return o.ApplyT(func(v *CustomResourceDefinition) CustomResourceDefinitionSpecOutput { return v.Spec }).(CustomResourceDefinitionSpecOutput)
}

// status indicates the actual state of the CustomResourceDefinition
func (o CustomResourceDefinitionOutput) Status() CustomResourceDefinitionStatusPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinition) CustomResourceDefinitionStatusPtrOutput { return v.Status }).(CustomResourceDefinitionStatusPtrOutput)
}

type CustomResourceDefinitionArrayOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinition)(nil)).Elem()
}

func (o CustomResourceDefinitionArrayOutput) ToCustomResourceDefinitionArrayOutput() CustomResourceDefinitionArrayOutput {
	return o
}

func (o CustomResourceDefinitionArrayOutput) ToCustomResourceDefinitionArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionArrayOutput {
	return o
}

func (o CustomResourceDefinitionArrayOutput) Index(i pulumi.IntInput) CustomResourceDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomResourceDefinition {
		return vs[0].([]*CustomResourceDefinition)[vs[1].(int)]
	}).(CustomResourceDefinitionOutput)
}

type CustomResourceDefinitionMapOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinition)(nil)).Elem()
}

func (o CustomResourceDefinitionMapOutput) ToCustomResourceDefinitionMapOutput() CustomResourceDefinitionMapOutput {
	return o
}

func (o CustomResourceDefinitionMapOutput) ToCustomResourceDefinitionMapOutputWithContext(ctx context.Context) CustomResourceDefinitionMapOutput {
	return o
}

func (o CustomResourceDefinitionMapOutput) MapIndex(k pulumi.StringInput) CustomResourceDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomResourceDefinition {
		return vs[0].(map[string]*CustomResourceDefinition)[vs[1].(string)]
	}).(CustomResourceDefinitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionInput)(nil)).Elem(), &CustomResourceDefinition{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionArrayInput)(nil)).Elem(), CustomResourceDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionMapInput)(nil)).Elem(), CustomResourceDefinitionMap{})
	pulumi.RegisterOutputType(CustomResourceDefinitionOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionArrayOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionMapOutput{})
}
