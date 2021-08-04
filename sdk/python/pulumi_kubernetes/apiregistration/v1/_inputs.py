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
    'APIServiceConditionArgs',
    'APIServiceSpecArgs',
    'APIServiceStatusArgs',
    'APIServiceArgs',
    'ServiceReferenceArgs',
]

@pulumi.input_type
class APIServiceConditionArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 type: pulumi.Input[str],
                 last_transition_time: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None):
        """
        APIServiceCondition describes the state of an APIService at a particular point
        :param pulumi.Input[str] status: Status is the status of the condition. Can be True, False, Unknown.
        :param pulumi.Input[str] type: Type is the type of the condition.
        :param pulumi.Input[str] last_transition_time: Last time the condition transitioned from one status to another.
        :param pulumi.Input[str] message: Human-readable message indicating details about last transition.
        :param pulumi.Input[str] reason: Unique, one-word, CamelCase reason for the condition's last transition.
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)
        if last_transition_time is not None:
            pulumi.set(__self__, "last_transition_time", last_transition_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if reason is not None:
            pulumi.set(__self__, "reason", reason)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Status is the status of the condition. Can be True, False, Unknown.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type is the type of the condition.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="lastTransitionTime")
    def last_transition_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last time the condition transitioned from one status to another.
        """
        return pulumi.get(self, "last_transition_time")

    @last_transition_time.setter
    def last_transition_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_transition_time", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable message indicating details about last transition.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def reason(self) -> Optional[pulumi.Input[str]]:
        """
        Unique, one-word, CamelCase reason for the condition's last transition.
        """
        return pulumi.get(self, "reason")

    @reason.setter
    def reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reason", value)


@pulumi.input_type
class APIServiceSpecArgs:
    def __init__(__self__, *,
                 group_priority_minimum: pulumi.Input[int],
                 version_priority: pulumi.Input[int],
                 ca_bundle: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 insecure_skip_tls_verify: Optional[pulumi.Input[bool]] = None,
                 service: Optional[pulumi.Input['ServiceReferenceArgs']] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
        :param pulumi.Input[int] group_priority_minimum: GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
        :param pulumi.Input[int] version_priority: VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
        :param pulumi.Input[str] ca_bundle: CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
        :param pulumi.Input[str] group: Group is the API group name this server hosts
        :param pulumi.Input[bool] insecure_skip_tls_verify: InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.
        :param pulumi.Input['ServiceReferenceArgs'] service: Service is a reference to the service for this API server.  It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.
        :param pulumi.Input[str] version: Version is the API version this server hosts.  For example, "v1"
        """
        pulumi.set(__self__, "group_priority_minimum", group_priority_minimum)
        pulumi.set(__self__, "version_priority", version_priority)
        if ca_bundle is not None:
            pulumi.set(__self__, "ca_bundle", ca_bundle)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if insecure_skip_tls_verify is not None:
            pulumi.set(__self__, "insecure_skip_tls_verify", insecure_skip_tls_verify)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="groupPriorityMinimum")
    def group_priority_minimum(self) -> pulumi.Input[int]:
        """
        GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
        """
        return pulumi.get(self, "group_priority_minimum")

    @group_priority_minimum.setter
    def group_priority_minimum(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_priority_minimum", value)

    @property
    @pulumi.getter(name="versionPriority")
    def version_priority(self) -> pulumi.Input[int]:
        """
        VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
        """
        return pulumi.get(self, "version_priority")

    @version_priority.setter
    def version_priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "version_priority", value)

    @property
    @pulumi.getter(name="caBundle")
    def ca_bundle(self) -> Optional[pulumi.Input[str]]:
        """
        CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
        """
        return pulumi.get(self, "ca_bundle")

    @ca_bundle.setter
    def ca_bundle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_bundle", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group is the API group name this server hosts
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="insecureSkipTLSVerify")
    def insecure_skip_tls_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.
        """
        return pulumi.get(self, "insecure_skip_tls_verify")

    @insecure_skip_tls_verify.setter
    def insecure_skip_tls_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure_skip_tls_verify", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input['ServiceReferenceArgs']]:
        """
        Service is a reference to the service for this API server.  It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input['ServiceReferenceArgs']]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version is the API version this server hosts.  For example, "v1"
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class APIServiceStatusArgs:
    def __init__(__self__, *,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['APIServiceConditionArgs']]]] = None):
        """
        APIServiceStatus contains derived information about an API server
        :param pulumi.Input[Sequence[pulumi.Input['APIServiceConditionArgs']]] conditions: Current service state of apiService.
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['APIServiceConditionArgs']]]]:
        """
        Current service state of apiService.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['APIServiceConditionArgs']]]]):
        pulumi.set(self, "conditions", value)


@pulumi.input_type
class APIServiceArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['APIServiceSpecArgs']] = None,
                 status: Optional[pulumi.Input['APIServiceStatusArgs']] = None):
        """
        APIService represents a server for a particular GroupVersion. Name must be "version.group".
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['APIServiceSpecArgs'] spec: Spec contains information for locating and communicating with a server
        :param pulumi.Input['APIServiceStatusArgs'] status: Status contains derived information about an API server
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'apiregistration.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'APIService')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

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
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['APIServiceSpecArgs']]:
        """
        Spec contains information for locating and communicating with a server
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['APIServiceSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['APIServiceStatusArgs']]:
        """
        Status contains derived information about an API server
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['APIServiceStatusArgs']]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class ServiceReferenceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        ServiceReference holds a reference to Service.legacy.k8s.io
        :param pulumi.Input[str] name: Name is the name of the service
        :param pulumi.Input[str] namespace: Namespace is the namespace of the service
        :param pulumi.Input[int] port: If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name is the name of the service
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace is the namespace of the service
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


