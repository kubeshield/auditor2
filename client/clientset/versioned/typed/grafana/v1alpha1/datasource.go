/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeshield.dev/auditor/apis/grafana/v1alpha1"
	scheme "kubeshield.dev/auditor/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatasourcesGetter has a method to return a DatasourceInterface.
// A group's client should implement this interface.
type DatasourcesGetter interface {
	Datasources(namespace string) DatasourceInterface
}

// DatasourceInterface has methods to work with Datasource resources.
type DatasourceInterface interface {
	Create(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.CreateOptions) (*v1alpha1.Datasource, error)
	Update(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.UpdateOptions) (*v1alpha1.Datasource, error)
	UpdateStatus(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.UpdateOptions) (*v1alpha1.Datasource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Datasource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DatasourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Datasource, err error)
	DatasourceExpansion
}

// datasources implements DatasourceInterface
type datasources struct {
	client rest.Interface
	ns     string
}

// newDatasources returns a Datasources
func newDatasources(c *GrafanaV1alpha1Client, namespace string) *datasources {
	return &datasources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the datasource, and returns the corresponding datasource object, and an error if there is any.
func (c *datasources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Datasource, err error) {
	result = &v1alpha1.Datasource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Datasources that match those selectors.
func (c *datasources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatasourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatasourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datasources.
func (c *datasources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a datasource and creates it.  Returns the server's representation of the datasource, and an error, if there is any.
func (c *datasources) Create(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.CreateOptions) (result *v1alpha1.Datasource, err error) {
	result = &v1alpha1.Datasource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a datasource and updates it. Returns the server's representation of the datasource, and an error, if there is any.
func (c *datasources) Update(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.UpdateOptions) (result *v1alpha1.Datasource, err error) {
	result = &v1alpha1.Datasource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasources").
		Name(datasource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *datasources) UpdateStatus(ctx context.Context, datasource *v1alpha1.Datasource, opts v1.UpdateOptions) (result *v1alpha1.Datasource, err error) {
	result = &v1alpha1.Datasource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasources").
		Name(datasource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the datasource and deletes it. Returns an error if one occurs.
func (c *datasources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datasources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched datasource.
func (c *datasources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Datasource, err error) {
	result = &v1alpha1.Datasource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datasources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
