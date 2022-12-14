/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "carrier/pkg/apis/xiaofeng/v1"
	scheme "carrier/pkg/generated/clientset/versioned/scheme"
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CarriersGetter has a method to return a CarrierInterface.
// A group's client should implement this interface.
type CarriersGetter interface {
	Carriers(namespace string) CarrierInterface
}

// CarrierInterface has methods to work with Carrier resources.
type CarrierInterface interface {
	Create(ctx context.Context, carrier *v1.Carrier, opts metav1.CreateOptions) (*v1.Carrier, error)
	Update(ctx context.Context, carrier *v1.Carrier, opts metav1.UpdateOptions) (*v1.Carrier, error)
	UpdateStatus(ctx context.Context, carrier *v1.Carrier, opts metav1.UpdateOptions) (*v1.Carrier, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Carrier, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CarrierList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Carrier, err error)
	CarrierExpansion
}

// carriers implements CarrierInterface
type carriers struct {
	client rest.Interface
	ns     string
}

// newCarriers returns a Carriers
func newCarriers(c *XiaofengV1Client, namespace string) *carriers {
	return &carriers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the carrier, and returns the corresponding carrier object, and an error if there is any.
func (c *carriers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Carrier, err error) {
	result = &v1.Carrier{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("carriers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Carriers that match those selectors.
func (c *carriers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CarrierList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CarrierList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("carriers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested carriers.
func (c *carriers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("carriers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a carrier and creates it.  Returns the server's representation of the carrier, and an error, if there is any.
func (c *carriers) Create(ctx context.Context, carrier *v1.Carrier, opts metav1.CreateOptions) (result *v1.Carrier, err error) {
	result = &v1.Carrier{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("carriers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(carrier).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a carrier and updates it. Returns the server's representation of the carrier, and an error, if there is any.
func (c *carriers) Update(ctx context.Context, carrier *v1.Carrier, opts metav1.UpdateOptions) (result *v1.Carrier, err error) {
	result = &v1.Carrier{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("carriers").
		Name(carrier.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(carrier).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *carriers) UpdateStatus(ctx context.Context, carrier *v1.Carrier, opts metav1.UpdateOptions) (result *v1.Carrier, err error) {
	result = &v1.Carrier{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("carriers").
		Name(carrier.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(carrier).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the carrier and deletes it. Returns an error if one occurs.
func (c *carriers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("carriers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *carriers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("carriers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched carrier.
func (c *carriers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Carrier, err error) {
	result = &v1.Carrier{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("carriers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
