// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is the list of ConfigMaps.
	Items ConfigMapTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewConfigMapList registers a new resource with the given unique name, arguments, and options.
func NewConfigMapList(ctx *pulumi.Context,
	name string, args *ConfigMapListArgs, opts ...pulumi.ResourceOption) (*ConfigMapList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ConfigMapList")
	var resource ConfigMapList
	err := ctx.RegisterResource("kubernetes:core/v1:ConfigMapList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMapList gets an existing ConfigMapList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigMapList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigMapListState, opts ...pulumi.ResourceOption) (*ConfigMapList, error) {
	var resource ConfigMapList
	err := ctx.ReadResource("kubernetes:core/v1:ConfigMapList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigMapList resources.
type configMapListState struct {
}

type ConfigMapListState struct {
}

func (ConfigMapListState) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapListState)(nil)).Elem()
}

type configMapListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of ConfigMaps.
	Items []ConfigMapType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ConfigMapList resource.
type ConfigMapListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of ConfigMaps.
	Items ConfigMapTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (ConfigMapListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapListArgs)(nil)).Elem()
}

type ConfigMapListInput interface {
	pulumi.Input

	ToConfigMapListOutput() ConfigMapListOutput
	ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput
}

func (*ConfigMapList) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMapList)(nil)).Elem()
}

func (i *ConfigMapList) ToConfigMapListOutput() ConfigMapListOutput {
	return i.ToConfigMapListOutputWithContext(context.Background())
}

func (i *ConfigMapList) ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapListOutput)
}

// ConfigMapListArrayInput is an input type that accepts ConfigMapListArray and ConfigMapListArrayOutput values.
// You can construct a concrete instance of `ConfigMapListArrayInput` via:
//
//	ConfigMapListArray{ ConfigMapListArgs{...} }
type ConfigMapListArrayInput interface {
	pulumi.Input

	ToConfigMapListArrayOutput() ConfigMapListArrayOutput
	ToConfigMapListArrayOutputWithContext(context.Context) ConfigMapListArrayOutput
}

type ConfigMapListArray []ConfigMapListInput

func (ConfigMapListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMapList)(nil)).Elem()
}

func (i ConfigMapListArray) ToConfigMapListArrayOutput() ConfigMapListArrayOutput {
	return i.ToConfigMapListArrayOutputWithContext(context.Background())
}

func (i ConfigMapListArray) ToConfigMapListArrayOutputWithContext(ctx context.Context) ConfigMapListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapListArrayOutput)
}

// ConfigMapListMapInput is an input type that accepts ConfigMapListMap and ConfigMapListMapOutput values.
// You can construct a concrete instance of `ConfigMapListMapInput` via:
//
//	ConfigMapListMap{ "key": ConfigMapListArgs{...} }
type ConfigMapListMapInput interface {
	pulumi.Input

	ToConfigMapListMapOutput() ConfigMapListMapOutput
	ToConfigMapListMapOutputWithContext(context.Context) ConfigMapListMapOutput
}

type ConfigMapListMap map[string]ConfigMapListInput

func (ConfigMapListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMapList)(nil)).Elem()
}

func (i ConfigMapListMap) ToConfigMapListMapOutput() ConfigMapListMapOutput {
	return i.ToConfigMapListMapOutputWithContext(context.Background())
}

func (i ConfigMapListMap) ToConfigMapListMapOutputWithContext(ctx context.Context) ConfigMapListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapListMapOutput)
}

type ConfigMapListOutput struct{ *pulumi.OutputState }

func (ConfigMapListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMapList)(nil)).Elem()
}

func (o ConfigMapListOutput) ToConfigMapListOutput() ConfigMapListOutput {
	return o
}

func (o ConfigMapListOutput) ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ConfigMapListOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigMapList) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Items is the list of ConfigMaps.
func (o ConfigMapListOutput) Items() ConfigMapTypeArrayOutput {
	return o.ApplyT(func(v *ConfigMapList) ConfigMapTypeArrayOutput { return v.Items }).(ConfigMapTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ConfigMapListOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigMapList) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ConfigMapListOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v *ConfigMapList) metav1.ListMetaPtrOutput { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

type ConfigMapListArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMapList)(nil)).Elem()
}

func (o ConfigMapListArrayOutput) ToConfigMapListArrayOutput() ConfigMapListArrayOutput {
	return o
}

func (o ConfigMapListArrayOutput) ToConfigMapListArrayOutputWithContext(ctx context.Context) ConfigMapListArrayOutput {
	return o
}

func (o ConfigMapListArrayOutput) Index(i pulumi.IntInput) ConfigMapListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigMapList {
		return vs[0].([]*ConfigMapList)[vs[1].(int)]
	}).(ConfigMapListOutput)
}

type ConfigMapListMapOutput struct{ *pulumi.OutputState }

func (ConfigMapListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMapList)(nil)).Elem()
}

func (o ConfigMapListMapOutput) ToConfigMapListMapOutput() ConfigMapListMapOutput {
	return o
}

func (o ConfigMapListMapOutput) ToConfigMapListMapOutputWithContext(ctx context.Context) ConfigMapListMapOutput {
	return o
}

func (o ConfigMapListMapOutput) MapIndex(k pulumi.StringInput) ConfigMapListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigMapList {
		return vs[0].(map[string]*ConfigMapList)[vs[1].(string)]
	}).(ConfigMapListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapListInput)(nil)).Elem(), &ConfigMapList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapListArrayInput)(nil)).Elem(), ConfigMapListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapListMapInput)(nil)).Elem(), ConfigMapListMap{})
	pulumi.RegisterOutputType(ConfigMapListOutput{})
	pulumi.RegisterOutputType(ConfigMapListArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapListMapOutput{})
}
