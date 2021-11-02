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

	akashnetworkv1 "github.com/ovrclk/akash/pkg/apis/akash.network/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInventoryRequests implements InventoryRequestInterface
type FakeInventoryRequests struct {
	Fake *FakeAkashV1
}

var inventoryrequestsResource = schema.GroupVersionResource{Group: "akash.network", Version: "v1", Resource: "inventoryrequests"}

var inventoryrequestsKind = schema.GroupVersionKind{Group: "akash.network", Version: "v1", Kind: "InventoryRequest"}

// Get takes name of the inventoryRequest, and returns the corresponding inventoryRequest object, and an error if there is any.
func (c *FakeInventoryRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *akashnetworkv1.InventoryRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(inventoryrequestsResource, name), &akashnetworkv1.InventoryRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*akashnetworkv1.InventoryRequest), err
}

// List takes label and field selectors, and returns the list of InventoryRequests that match those selectors.
func (c *FakeInventoryRequests) List(ctx context.Context, opts v1.ListOptions) (result *akashnetworkv1.InventoryRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(inventoryrequestsResource, inventoryrequestsKind, opts), &akashnetworkv1.InventoryRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &akashnetworkv1.InventoryRequestList{ListMeta: obj.(*akashnetworkv1.InventoryRequestList).ListMeta}
	for _, item := range obj.(*akashnetworkv1.InventoryRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inventoryRequests.
func (c *FakeInventoryRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(inventoryrequestsResource, opts))
}

// Create takes the representation of a inventoryRequest and creates it.  Returns the server's representation of the inventoryRequest, and an error, if there is any.
func (c *FakeInventoryRequests) Create(ctx context.Context, inventoryRequest *akashnetworkv1.InventoryRequest, opts v1.CreateOptions) (result *akashnetworkv1.InventoryRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(inventoryrequestsResource, inventoryRequest), &akashnetworkv1.InventoryRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*akashnetworkv1.InventoryRequest), err
}

// Update takes the representation of a inventoryRequest and updates it. Returns the server's representation of the inventoryRequest, and an error, if there is any.
func (c *FakeInventoryRequests) Update(ctx context.Context, inventoryRequest *akashnetworkv1.InventoryRequest, opts v1.UpdateOptions) (result *akashnetworkv1.InventoryRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(inventoryrequestsResource, inventoryRequest), &akashnetworkv1.InventoryRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*akashnetworkv1.InventoryRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInventoryRequests) UpdateStatus(ctx context.Context, inventoryRequest *akashnetworkv1.InventoryRequest, opts v1.UpdateOptions) (*akashnetworkv1.InventoryRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(inventoryrequestsResource, "status", inventoryRequest), &akashnetworkv1.InventoryRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*akashnetworkv1.InventoryRequest), err
}

// Delete takes name of the inventoryRequest and deletes it. Returns an error if one occurs.
func (c *FakeInventoryRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(inventoryrequestsResource, name), &akashnetworkv1.InventoryRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInventoryRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(inventoryrequestsResource, listOpts)

	_, err := c.Fake.Invokes(action, &akashnetworkv1.InventoryRequestList{})
	return err
}

// Patch applies the patch and returns the patched inventoryRequest.
func (c *FakeInventoryRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *akashnetworkv1.InventoryRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(inventoryrequestsResource, name, pt, data, subresources...), &akashnetworkv1.InventoryRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*akashnetworkv1.InventoryRequest), err
}