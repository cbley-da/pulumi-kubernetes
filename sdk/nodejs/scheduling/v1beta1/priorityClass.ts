// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
 */
export class PriorityClass extends pulumi.CustomResource {
    /**
     * Get an existing PriorityClass resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PriorityClass {
        return new PriorityClass(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:scheduling.k8s.io/v1beta1:PriorityClass';

    /**
     * Returns true if the given object is an instance of PriorityClass.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PriorityClass {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PriorityClass.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"scheduling.k8s.io/v1beta1">;
    /**
     * description is an arbitrary string that usually provides guidelines on when this priority class should be used.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
     */
    public readonly globalDefault!: pulumi.Output<boolean>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"PriorityClass">;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
     */
    public readonly preemptionPolicy!: pulumi.Output<string>;
    /**
     * The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
     */
    public readonly value!: pulumi.Output<number>;

    /**
     * Create a PriorityClass resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PriorityClassArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["apiVersion"] = "scheduling.k8s.io/v1beta1";
            inputs["description"] = args ? args.description : undefined;
            inputs["globalDefault"] = args ? args.globalDefault : undefined;
            inputs["kind"] = "PriorityClass";
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["preemptionPolicy"] = args ? args.preemptionPolicy : undefined;
            inputs["value"] = args ? args.value : undefined;
        } else {
            inputs["apiVersion"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["globalDefault"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["preemptionPolicy"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:scheduling.k8s.io/v1:PriorityClass" }, { type: "kubernetes:scheduling.k8s.io/v1alpha1:PriorityClass" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PriorityClass.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PriorityClass resource.
 */
export interface PriorityClassArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"scheduling.k8s.io/v1beta1">;
    /**
     * description is an arbitrary string that usually provides guidelines on when this priority class should be used.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
     */
    readonly globalDefault?: pulumi.Input<boolean>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"PriorityClass">;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
     */
    readonly preemptionPolicy?: pulumi.Input<string>;
    /**
     * The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
     */
    readonly value: pulumi.Input<number>;
}
