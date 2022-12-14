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

package fake

import (
	xiaofengv1 "carrier/pkg/apis/xiaofeng/v1"
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCarriers implements CarrierInterface
type FakeCarriers struct {
	Fake *FakeXiaofengV1
	ns   string
}

var carriersResource = schema.GroupVersionResource{Group: "xiaofeng.k8s.io", Version: "v1", Resource: "carriers"}

var carriersKind = schema.GroupVersionKind{Group: "xiaofeng.k8s.io", Version: "v1", Kind: "Carrier"}

// Get takes name of the carrier, and returns the corresponding carrier object, and an error if there is any.
func (c *FakeCarriers) Get(ctx context.Context, name string, options v1.GetOptions) (result *xiaofengv1.Carrier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(carriersResource, c.ns, name), &xiaofengv1.Carrier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*xiaofengv1.Carrier), err
}

// List takes label and field selectors, and returns the list of Carriers that match those selectors.
func (c *FakeCarriers) List(ctx context.Context, opts v1.ListOptions) (result *xiaofengv1.CarrierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(carriersResource, carriersKind, c.ns, opts), &xiaofengv1.CarrierList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &xiaofengv1.CarrierList{ListMeta: obj.(*xiaofengv1.CarrierList).ListMeta}
	for _, item := range obj.(*xiaofengv1.CarrierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested carriers.
func (c *FakeCarriers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(carriersResource, c.ns, opts))

}

// Create takes the representation of a carrier and creates it.  Returns the server's representation of the carrier, and an error, if there is any.
func (c *FakeCarriers) Create(ctx context.Context, carrier *xiaofengv1.Carrier, opts v1.CreateOptions) (result *xiaofengv1.Carrier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(carriersResource, c.ns, carrier), &xiaofengv1.Carrier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*xiaofengv1.Carrier), err
}

// Update takes the representation of a carrier and updates it. Returns the server's representation of the carrier, and an error, if there is any.
func (c *FakeCarriers) Update(ctx context.Context, carrier *xiaofengv1.Carrier, opts v1.UpdateOptions) (result *xiaofengv1.Carrier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(carriersResource, c.ns, carrier), &xiaofengv1.Carrier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*xiaofengv1.Carrier), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCarriers) UpdateStatus(ctx context.Context, carrier *xiaofengv1.Carrier, opts v1.UpdateOptions) (*xiaofengv1.Carrier, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(carriersResource, "status", c.ns, carrier), &xiaofengv1.Carrier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*xiaofengv1.Carrier), err
}

// Delete takes name of the carrier and deletes it. Returns an error if one occurs.
func (c *FakeCarriers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(carriersResource, c.ns, name, opts), &xiaofengv1.Carrier{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCarriers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(carriersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &xiaofengv1.CarrierList{})
	return err
}

// Patch applies the patch and returns the patched carrier.
func (c *FakeCarriers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *xiaofengv1.Carrier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(carriersResource, c.ns, name, pt, data, subresources...), &xiaofengv1.Carrier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*xiaofengv1.Carrier), err
}
