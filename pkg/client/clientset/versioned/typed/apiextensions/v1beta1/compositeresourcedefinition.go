/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/crossplane/crossplane/apis/apiextensions/v1beta1"
	scheme "github.com/crossplane/crossplane/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CompositeResourceDefinitionsGetter has a method to return a CompositeResourceDefinitionInterface.
// A group's client should implement this interface.
type CompositeResourceDefinitionsGetter interface {
	CompositeResourceDefinitions() CompositeResourceDefinitionInterface
}

// CompositeResourceDefinitionInterface has methods to work with CompositeResourceDefinition resources.
type CompositeResourceDefinitionInterface interface {
	Create(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.CreateOptions) (*v1beta1.CompositeResourceDefinition, error)
	Update(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.UpdateOptions) (*v1beta1.CompositeResourceDefinition, error)
	UpdateStatus(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.UpdateOptions) (*v1beta1.CompositeResourceDefinition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CompositeResourceDefinition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CompositeResourceDefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CompositeResourceDefinition, err error)
	CompositeResourceDefinitionExpansion
}

// compositeResourceDefinitions implements CompositeResourceDefinitionInterface
type compositeResourceDefinitions struct {
	client rest.Interface
}

// newCompositeResourceDefinitions returns a CompositeResourceDefinitions
func newCompositeResourceDefinitions(c *ApiextensionsV1beta1Client) *compositeResourceDefinitions {
	return &compositeResourceDefinitions{
		client: c.RESTClient(),
	}
}

// Get takes name of the compositeResourceDefinition, and returns the corresponding compositeResourceDefinition object, and an error if there is any.
func (c *compositeResourceDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CompositeResourceDefinition, err error) {
	result = &v1beta1.CompositeResourceDefinition{}
	err = c.client.Get().
		Resource("compositeresourcedefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CompositeResourceDefinitions that match those selectors.
func (c *compositeResourceDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CompositeResourceDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CompositeResourceDefinitionList{}
	err = c.client.Get().
		Resource("compositeresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested compositeResourceDefinitions.
func (c *compositeResourceDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("compositeresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a compositeResourceDefinition and creates it.  Returns the server's representation of the compositeResourceDefinition, and an error, if there is any.
func (c *compositeResourceDefinitions) Create(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.CreateOptions) (result *v1beta1.CompositeResourceDefinition, err error) {
	result = &v1beta1.CompositeResourceDefinition{}
	err = c.client.Post().
		Resource("compositeresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositeResourceDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a compositeResourceDefinition and updates it. Returns the server's representation of the compositeResourceDefinition, and an error, if there is any.
func (c *compositeResourceDefinitions) Update(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.UpdateOptions) (result *v1beta1.CompositeResourceDefinition, err error) {
	result = &v1beta1.CompositeResourceDefinition{}
	err = c.client.Put().
		Resource("compositeresourcedefinitions").
		Name(compositeResourceDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositeResourceDefinition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *compositeResourceDefinitions) UpdateStatus(ctx context.Context, compositeResourceDefinition *v1beta1.CompositeResourceDefinition, opts v1.UpdateOptions) (result *v1beta1.CompositeResourceDefinition, err error) {
	result = &v1beta1.CompositeResourceDefinition{}
	err = c.client.Put().
		Resource("compositeresourcedefinitions").
		Name(compositeResourceDefinition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositeResourceDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the compositeResourceDefinition and deletes it. Returns an error if one occurs.
func (c *compositeResourceDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("compositeresourcedefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *compositeResourceDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("compositeresourcedefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched compositeResourceDefinition.
func (c *compositeResourceDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CompositeResourceDefinition, err error) {
	result = &v1beta1.CompositeResourceDefinition{}
	err = c.client.Patch(pt).
		Resource("compositeresourcedefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
