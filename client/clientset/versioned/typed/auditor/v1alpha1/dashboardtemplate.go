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

	v1alpha1 "kubeshield.dev/auditor/apis/auditor/v1alpha1"
	scheme "kubeshield.dev/auditor/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DashboardTemplatesGetter has a method to return a DashboardTemplateInterface.
// A group's client should implement this interface.
type DashboardTemplatesGetter interface {
	DashboardTemplates(namespace string) DashboardTemplateInterface
}

// DashboardTemplateInterface has methods to work with DashboardTemplate resources.
type DashboardTemplateInterface interface {
	Create(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.CreateOptions) (*v1alpha1.DashboardTemplate, error)
	Update(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.UpdateOptions) (*v1alpha1.DashboardTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DashboardTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DashboardTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DashboardTemplate, err error)
	DashboardTemplateExpansion
}

// dashboardTemplates implements DashboardTemplateInterface
type dashboardTemplates struct {
	client rest.Interface
	ns     string
}

// newDashboardTemplates returns a DashboardTemplates
func newDashboardTemplates(c *AuditorV1alpha1Client, namespace string) *dashboardTemplates {
	return &dashboardTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dashboardTemplate, and returns the corresponding dashboardTemplate object, and an error if there is any.
func (c *dashboardTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DashboardTemplate, err error) {
	result = &v1alpha1.DashboardTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DashboardTemplates that match those selectors.
func (c *dashboardTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DashboardTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DashboardTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dashboardTemplates.
func (c *dashboardTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dashboardTemplate and creates it.  Returns the server's representation of the dashboardTemplate, and an error, if there is any.
func (c *dashboardTemplates) Create(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.CreateOptions) (result *v1alpha1.DashboardTemplate, err error) {
	result = &v1alpha1.DashboardTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dashboardTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dashboardTemplate and updates it. Returns the server's representation of the dashboardTemplate, and an error, if there is any.
func (c *dashboardTemplates) Update(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.UpdateOptions) (result *v1alpha1.DashboardTemplate, err error) {
	result = &v1alpha1.DashboardTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		Name(dashboardTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dashboardTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dashboardTemplate and deletes it. Returns an error if one occurs.
func (c *dashboardTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dashboardTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dashboardtemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dashboardTemplate.
func (c *dashboardTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DashboardTemplate, err error) {
	result = &v1alpha1.DashboardTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dashboardtemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
