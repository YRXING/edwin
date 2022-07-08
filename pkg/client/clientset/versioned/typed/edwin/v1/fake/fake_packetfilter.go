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
	"context"

	edwinv1 "github.com/YRXING/edwin/pkg/apis/edwin/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePacketFilters implements PacketFilterInterface
type FakePacketFilters struct {
	Fake *FakeEdwinV1
	ns   string
}

var packetfiltersResource = schema.GroupVersionResource{Group: "edwin.k8s.io", Version: "v1", Resource: "packetfilters"}

var packetfiltersKind = schema.GroupVersionKind{Group: "edwin.k8s.io", Version: "v1", Kind: "PacketFilter"}

// Get takes name of the packetFilter, and returns the corresponding packetFilter object, and an error if there is any.
func (c *FakePacketFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *edwinv1.PacketFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packetfiltersResource, c.ns, name), &edwinv1.PacketFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edwinv1.PacketFilter), err
}

// List takes label and field selectors, and returns the list of PacketFilters that match those selectors.
func (c *FakePacketFilters) List(ctx context.Context, opts v1.ListOptions) (result *edwinv1.PacketFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packetfiltersResource, packetfiltersKind, c.ns, opts), &edwinv1.PacketFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &edwinv1.PacketFilterList{ListMeta: obj.(*edwinv1.PacketFilterList).ListMeta}
	for _, item := range obj.(*edwinv1.PacketFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packetFilters.
func (c *FakePacketFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packetfiltersResource, c.ns, opts))

}

// Create takes the representation of a packetFilter and creates it.  Returns the server's representation of the packetFilter, and an error, if there is any.
func (c *FakePacketFilters) Create(ctx context.Context, packetFilter *edwinv1.PacketFilter, opts v1.CreateOptions) (result *edwinv1.PacketFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packetfiltersResource, c.ns, packetFilter), &edwinv1.PacketFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edwinv1.PacketFilter), err
}

// Update takes the representation of a packetFilter and updates it. Returns the server's representation of the packetFilter, and an error, if there is any.
func (c *FakePacketFilters) Update(ctx context.Context, packetFilter *edwinv1.PacketFilter, opts v1.UpdateOptions) (result *edwinv1.PacketFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packetfiltersResource, c.ns, packetFilter), &edwinv1.PacketFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edwinv1.PacketFilter), err
}

// Delete takes name of the packetFilter and deletes it. Returns an error if one occurs.
func (c *FakePacketFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(packetfiltersResource, c.ns, name, opts), &edwinv1.PacketFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePacketFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packetfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &edwinv1.PacketFilterList{})
	return err
}

// Patch applies the patch and returns the patched packetFilter.
func (c *FakePacketFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *edwinv1.PacketFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packetfiltersResource, c.ns, name, pt, data, subresources...), &edwinv1.PacketFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*edwinv1.PacketFilter), err
}
