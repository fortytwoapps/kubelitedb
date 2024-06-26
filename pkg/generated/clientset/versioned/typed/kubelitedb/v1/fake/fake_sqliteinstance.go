/*
Copyright 2024 Forty Two Apps.

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

	v1 "github.com/fortytwoapps/kubelitedb/pkg/apis/kubelitedb/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSQLiteInstances implements SQLiteInstanceInterface
type FakeSQLiteInstances struct {
	Fake *FakeKubelitedbV1
	ns   string
}

var sqliteinstancesResource = v1.SchemeGroupVersion.WithResource("sqliteinstances")

var sqliteinstancesKind = v1.SchemeGroupVersion.WithKind("SQLiteInstance")

// Get takes name of the sQLiteInstance, and returns the corresponding sQLiteInstance object, and an error if there is any.
func (c *FakeSQLiteInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SQLiteInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqliteinstancesResource, c.ns, name), &v1.SQLiteInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SQLiteInstance), err
}

// List takes label and field selectors, and returns the list of SQLiteInstances that match those selectors.
func (c *FakeSQLiteInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SQLiteInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqliteinstancesResource, sqliteinstancesKind, c.ns, opts), &v1.SQLiteInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SQLiteInstanceList{ListMeta: obj.(*v1.SQLiteInstanceList).ListMeta}
	for _, item := range obj.(*v1.SQLiteInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sQLiteInstances.
func (c *FakeSQLiteInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqliteinstancesResource, c.ns, opts))

}

// Create takes the representation of a sQLiteInstance and creates it.  Returns the server's representation of the sQLiteInstance, and an error, if there is any.
func (c *FakeSQLiteInstances) Create(ctx context.Context, sQLiteInstance *v1.SQLiteInstance, opts metav1.CreateOptions) (result *v1.SQLiteInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqliteinstancesResource, c.ns, sQLiteInstance), &v1.SQLiteInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SQLiteInstance), err
}

// Update takes the representation of a sQLiteInstance and updates it. Returns the server's representation of the sQLiteInstance, and an error, if there is any.
func (c *FakeSQLiteInstances) Update(ctx context.Context, sQLiteInstance *v1.SQLiteInstance, opts metav1.UpdateOptions) (result *v1.SQLiteInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqliteinstancesResource, c.ns, sQLiteInstance), &v1.SQLiteInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SQLiteInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSQLiteInstances) UpdateStatus(ctx context.Context, sQLiteInstance *v1.SQLiteInstance, opts metav1.UpdateOptions) (*v1.SQLiteInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqliteinstancesResource, "status", c.ns, sQLiteInstance), &v1.SQLiteInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SQLiteInstance), err
}

// Delete takes name of the sQLiteInstance and deletes it. Returns an error if one occurs.
func (c *FakeSQLiteInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sqliteinstancesResource, c.ns, name, opts), &v1.SQLiteInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSQLiteInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqliteinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SQLiteInstanceList{})
	return err
}

// Patch applies the patch and returns the patched sQLiteInstance.
func (c *FakeSQLiteInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SQLiteInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqliteinstancesResource, c.ns, name, pt, data, subresources...), &v1.SQLiteInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SQLiteInstance), err
}
