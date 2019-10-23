# azure.crossplane.io/v1alpha2 API Reference

Package v1alpha2 contains core resources for Microsoft Azure.

This API group contains the following Crossplane resources:

* [Provider](#Provider)
* [ResourceGroup](#ResourceGroup)

## Provider

A Provider configures an Azure &#39;provider&#39;, i.e. a connection to a particular Azure account using a particular Azure Service Principal.


Name | Type | Description
-----|------|------------
`apiVersion` | string | `azure.crossplane.io/v1alpha2`
`kind` | string | `Provider`
`metadata` | [meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#objectmeta-v1-meta) | Kubernetes object metadata.
`spec` | [ProviderSpec](#ProviderSpec) | A ProviderSpec defines the desired state of a Provider.



## ResourceGroup

A ResourceGroup is a managed resource that represents an Azure Resource Group.


Name | Type | Description
-----|------|------------
`apiVersion` | string | `azure.crossplane.io/v1alpha2`
`kind` | string | `ResourceGroup`
`metadata` | [meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#objectmeta-v1-meta) | Kubernetes object metadata.
`spec` | [ResourceGroupSpec](#ResourceGroupSpec) | A ResourceGroupSpec defines the desired state of a ResourceGroup.
`status` | [ResourceGroupStatus](#ResourceGroupStatus) | A ResourceGroupStatus represents the observed status of a ResourceGroup.



## ProviderSpec

A ProviderSpec defines the desired state of a Provider.

Appears in:

* [Provider](#Provider)


Name | Type | Description
-----|------|------------
`credentialsSecretRef` | [core/v1.SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#secretkeyselector-v1-core) | Azure service principal credentials json secret key reference A Secret containing JSON encoded credentials for an Azure Service Principal that will be used to authenticate to this Azure Provider.



## ResourceGroupSpec

A ResourceGroupSpec defines the desired state of a ResourceGroup.

Appears in:

* [ResourceGroup](#ResourceGroup)


Name | Type | Description
-----|------|------------
`name` | string | Name of the resource group.
`location` | string | Location of the resource group. See the  official list of valid regions - https://azure.microsoft.com/en-us/global-infrastructure/regions/


ResourceGroupSpec supports all fields of:

* [v1alpha1.ResourceSpec](../crossplane-runtime/core-crossplane-io-v1alpha1.md#resourcespec)


## ResourceGroupStatus

A ResourceGroupStatus represents the observed status of a ResourceGroup.

Appears in:

* [ResourceGroup](#ResourceGroup)


Name | Type | Description
-----|------|------------
`name` | string | Name of the resource group.


ResourceGroupStatus supports all fields of:

* [v1alpha1.ResourceStatus](../crossplane-runtime/core-crossplane-io-v1alpha1.md#resourcestatus)


This API documentation was generated by `crossdocs`.