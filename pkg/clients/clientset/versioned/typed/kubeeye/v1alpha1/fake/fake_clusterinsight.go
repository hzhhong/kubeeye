/*
Copyright 2022.

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

	v1alpha1 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterInsights implements ClusterInsightInterface
type FakeClusterInsights struct {
	Fake *FakeKubeeyeV1alpha1
}

var clusterinsightsResource = schema.GroupVersionResource{Group: "kubeeye.kubesphere.io", Version: "v1alpha1", Resource: "clusterinsights"}

var clusterinsightsKind = schema.GroupVersionKind{Group: "kubeeye.kubesphere.io", Version: "v1alpha1", Kind: "ClusterInsight"}

// Get takes name of the clusterInsight, and returns the corresponding clusterInsight object, and an error if there is any.
func (c *FakeClusterInsights) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterInsight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterinsightsResource, name), &v1alpha1.ClusterInsight{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterInsight), err
}

// List takes label and field selectors, and returns the list of ClusterInsights that match those selectors.
func (c *FakeClusterInsights) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterInsightList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterinsightsResource, clusterinsightsKind, opts), &v1alpha1.ClusterInsightList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterInsightList{ListMeta: obj.(*v1alpha1.ClusterInsightList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterInsightList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterInsights.
func (c *FakeClusterInsights) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterinsightsResource, opts))
}

// Create takes the representation of a clusterInsight and creates it.  Returns the server's representation of the clusterInsight, and an error, if there is any.
func (c *FakeClusterInsights) Create(ctx context.Context, clusterInsight *v1alpha1.ClusterInsight, opts v1.CreateOptions) (result *v1alpha1.ClusterInsight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterinsightsResource, clusterInsight), &v1alpha1.ClusterInsight{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterInsight), err
}

// Update takes the representation of a clusterInsight and updates it. Returns the server's representation of the clusterInsight, and an error, if there is any.
func (c *FakeClusterInsights) Update(ctx context.Context, clusterInsight *v1alpha1.ClusterInsight, opts v1.UpdateOptions) (result *v1alpha1.ClusterInsight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterinsightsResource, clusterInsight), &v1alpha1.ClusterInsight{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterInsight), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterInsights) UpdateStatus(ctx context.Context, clusterInsight *v1alpha1.ClusterInsight, opts v1.UpdateOptions) (*v1alpha1.ClusterInsight, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterinsightsResource, "status", clusterInsight), &v1alpha1.ClusterInsight{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterInsight), err
}

// Delete takes name of the clusterInsight and deletes it. Returns an error if one occurs.
func (c *FakeClusterInsights) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterinsightsResource, name, opts), &v1alpha1.ClusterInsight{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterInsights) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterinsightsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterInsightList{})
	return err
}

// Patch applies the patch and returns the patched clusterInsight.
func (c *FakeClusterInsights) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterInsight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterinsightsResource, name, pt, data, subresources...), &v1alpha1.ClusterInsight{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterInsight), err
}
