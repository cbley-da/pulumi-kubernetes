# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import meta as _meta

__all__ = [
    'AggregationRuleArgs',
    'ClusterRoleBindingArgs',
    'ClusterRoleArgs',
    'PolicyRuleArgs',
    'RoleBindingArgs',
    'RoleRefArgs',
    'RoleArgs',
    'SubjectArgs',
]

@pulumi.input_type
class AggregationRuleArgs:
    def __init__(__self__, *,
                 cluster_role_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.LabelSelectorArgs']]]] = None):
        """
        AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
        :param pulumi.Input[Sequence[pulumi.Input['_meta.v1.LabelSelectorArgs']]] cluster_role_selectors: ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
        """
        if cluster_role_selectors is not None:
            pulumi.set(__self__, "cluster_role_selectors", cluster_role_selectors)

    @property
    @pulumi.getter(name="clusterRoleSelectors")
    def cluster_role_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.LabelSelectorArgs']]]]:
        """
        ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
        """
        return pulumi.get(self, "cluster_role_selectors")

    @cluster_role_selectors.setter
    def cluster_role_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_meta.v1.LabelSelectorArgs']]]]):
        pulumi.set(self, "cluster_role_selectors", value)


@pulumi.input_type
class ClusterRoleBindingArgs:
    def __init__(__self__, *,
                 role_ref: pulumi.Input['RoleRefArgs'],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 subjects: Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]] = None):
        """
        ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
        :param pulumi.Input['RoleRefArgs'] role_ref: RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]] subjects: Subjects holds references to the objects the role applies to.
        """
        pulumi.set(__self__, "role_ref", role_ref)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'rbac.authorization.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'ClusterRoleBinding')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if subjects is not None:
            pulumi.set(__self__, "subjects", subjects)

    @property
    @pulumi.getter(name="roleRef")
    def role_ref(self) -> pulumi.Input['RoleRefArgs']:
        """
        RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        """
        return pulumi.get(self, "role_ref")

    @role_ref.setter
    def role_ref(self, value: pulumi.Input['RoleRefArgs']):
        pulumi.set(self, "role_ref", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def subjects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]]:
        """
        Subjects holds references to the objects the role applies to.
        """
        return pulumi.get(self, "subjects")

    @subjects.setter
    def subjects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]]):
        pulumi.set(self, "subjects", value)


@pulumi.input_type
class ClusterRoleArgs:
    def __init__(__self__, *,
                 aggregation_rule: Optional[pulumi.Input['AggregationRuleArgs']] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]] = None):
        """
        ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
        :param pulumi.Input['AggregationRuleArgs'] aggregation_rule: AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]] rules: Rules holds all the PolicyRules for this ClusterRole
        """
        if aggregation_rule is not None:
            pulumi.set(__self__, "aggregation_rule", aggregation_rule)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'rbac.authorization.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'ClusterRole')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="aggregationRule")
    def aggregation_rule(self) -> Optional[pulumi.Input['AggregationRuleArgs']]:
        """
        AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
        """
        return pulumi.get(self, "aggregation_rule")

    @aggregation_rule.setter
    def aggregation_rule(self, value: Optional[pulumi.Input['AggregationRuleArgs']]):
        pulumi.set(self, "aggregation_rule", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]]:
        """
        Rules holds all the PolicyRules for this ClusterRole
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class PolicyRuleArgs:
    def __init__(__self__, *,
                 verbs: pulumi.Input[Sequence[pulumi.Input[str]]],
                 api_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 non_resource_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] verbs: Verbs is a list of Verbs that apply to ALL the ResourceKinds contained in this rule. '*' represents all verbs.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_groups: APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] non_resource_urls: NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_names: ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: Resources is a list of resources this rule applies to. '*' represents all resources.
        """
        pulumi.set(__self__, "verbs", verbs)
        if api_groups is not None:
            pulumi.set(__self__, "api_groups", api_groups)
        if non_resource_urls is not None:
            pulumi.set(__self__, "non_resource_urls", non_resource_urls)
        if resource_names is not None:
            pulumi.set(__self__, "resource_names", resource_names)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def verbs(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Verbs is a list of Verbs that apply to ALL the ResourceKinds contained in this rule. '*' represents all verbs.
        """
        return pulumi.get(self, "verbs")

    @verbs.setter
    def verbs(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "verbs", value)

    @property
    @pulumi.getter(name="apiGroups")
    def api_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
        """
        return pulumi.get(self, "api_groups")

    @api_groups.setter
    def api_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_groups", value)

    @property
    @pulumi.getter(name="nonResourceURLs")
    def non_resource_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
        """
        return pulumi.get(self, "non_resource_urls")

    @non_resource_urls.setter
    def non_resource_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "non_resource_urls", value)

    @property
    @pulumi.getter(name="resourceNames")
    def resource_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
        """
        return pulumi.get(self, "resource_names")

    @resource_names.setter
    def resource_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_names", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Resources is a list of resources this rule applies to. '*' represents all resources.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resources", value)


@pulumi.input_type
class RoleBindingArgs:
    def __init__(__self__, *,
                 role_ref: pulumi.Input['RoleRefArgs'],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 subjects: Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]] = None):
        """
        RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
        :param pulumi.Input['RoleRefArgs'] role_ref: RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]] subjects: Subjects holds references to the objects the role applies to.
        """
        pulumi.set(__self__, "role_ref", role_ref)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'rbac.authorization.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'RoleBinding')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if subjects is not None:
            pulumi.set(__self__, "subjects", subjects)

    @property
    @pulumi.getter(name="roleRef")
    def role_ref(self) -> pulumi.Input['RoleRefArgs']:
        """
        RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        """
        return pulumi.get(self, "role_ref")

    @role_ref.setter
    def role_ref(self, value: pulumi.Input['RoleRefArgs']):
        pulumi.set(self, "role_ref", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def subjects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]]:
        """
        Subjects holds references to the objects the role applies to.
        """
        return pulumi.get(self, "subjects")

    @subjects.setter
    def subjects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SubjectArgs']]]]):
        pulumi.set(self, "subjects", value)


@pulumi.input_type
class RoleRefArgs:
    def __init__(__self__, *,
                 api_group: pulumi.Input[str],
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        RoleRef contains information that points to the role being used
        :param pulumi.Input[str] api_group: APIGroup is the group for the resource being referenced
        :param pulumi.Input[str] kind: Kind is the type of resource being referenced
        :param pulumi.Input[str] name: Name is the name of resource being referenced
        """
        pulumi.set(__self__, "api_group", api_group)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiGroup")
    def api_group(self) -> pulumi.Input[str]:
        """
        APIGroup is the group for the resource being referenced
        """
        return pulumi.get(self, "api_group")

    @api_group.setter
    def api_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_group", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Kind is the type of resource being referenced
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name is the name of resource being referenced
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]] = None):
        """
        Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]] rules: Rules holds all the PolicyRules for this Role
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'rbac.authorization.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'Role')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]]:
        """
        Rules holds all the PolicyRules for this Role
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class SubjectArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str],
                 api_group: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
        :param pulumi.Input[str] kind: Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
        :param pulumi.Input[str] name: Name of the object being referenced.
        :param pulumi.Input[str] api_group: APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
        :param pulumi.Input[str] namespace: Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        if api_group is not None:
            pulumi.set(__self__, "api_group", api_group)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the object being referenced.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="apiGroup")
    def api_group(self) -> Optional[pulumi.Input[str]]:
        """
        APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
        """
        return pulumi.get(self, "api_group")

    @api_group.setter
    def api_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_group", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


